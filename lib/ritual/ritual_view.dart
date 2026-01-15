// ignore_for_file: prefer_html_components

import 'dart:async';

import 'package:jaspr/jaspr.dart';
import 'ritual_phase.dart';

class RitualView extends StatefulComponent {
  final VoidCallback onComplete;

  const RitualView({
    required this.onComplete,
    super.key,
  });

  @override
  State<RitualView> createState() => _RitualViewState();
}

class _RitualViewState extends State<RitualView> {
  RitualPhase phase = RitualPhase.greeting;

  Timer? _autoAdvanceTimer;
  Timer? _handoffTimer;

  int _phaseTick = 0;
  bool _hasInteracted = false;
  bool _handoffStarted = false;
  bool _onCompleteCalled = false;

  @override
  void initState() {
    super.initState();
    _scheduleAutoAdvanceForCurrentPhase();
  }

  @override
  void dispose() {
    _cancelAutoAdvanceTimer();
    _handoffTimer?.cancel();
    super.dispose();
  }

  void _cancelAutoAdvanceTimer() {
    _autoAdvanceTimer?.cancel();
    _autoAdvanceTimer = null;
  }

  void _scheduleAutoAdvanceForCurrentPhase() {
    _cancelAutoAdvanceTimer();

    if (_handoffStarted) return;
    if (phase.isTerminal) return;

    final delayMs = phase.autoAdvanceDelayMs;
    if (delayMs == null) return;

    _autoAdvanceTimer = Timer(Duration(milliseconds: delayMs), () {
      if (!mounted) return;
      _advance(isManual: false);
    });
  }

  void _beginHandoff() {
    if (_handoffStarted) return;
    _handoffStarted = true;

    _cancelAutoAdvanceTimer();
    _handoffTimer?.cancel();

    setState(() {
      // Triggers CSS fade-out via class.
    });

    _handoffTimer = Timer(const Duration(milliseconds: 600), () {
      if (!mounted) return;
      if (_onCompleteCalled) return;
      _onCompleteCalled = true;
      component.onComplete();
    });
  }

  void _advance({required bool isManual}) {
    if (_handoffStarted) return;
    if (isManual) {
      _hasInteracted = true;
    }

    final next = phase.next;

    if (next == null) {
      _beginHandoff();
      return;
    }

    setState(() {
      phase = next;
      _phaseTick++;
    });

    _scheduleAutoAdvanceForCurrentPhase();
  }

  void _handleTap() {
    if (_handoffStarted) return;

    _hasInteracted = true;
    _cancelAutoAdvanceTimer();

    _advance(isManual: true);
  }

  @override
  Component build(BuildContext context) {
    final classNames = <String>{
      'ritual',
      'ritual-phase-${phase.name}',
      if (_handoffStarted) 'ritual--handoff',
      if (_hasInteracted) 'ritual--interacted',
    }.join(' ');

    final showAffordance = phase.showTapAffordance && !_hasInteracted && !_handoffStarted;

    return Component.element(
      tag: 'div',
      attributes: {
        'class': classNames,
        'role': 'button',
        'tabindex': '0',
        'aria-label': 'Ritual step: ${phase.text}',
      },
      events: {
        'click': (_) => _handleTap(),
      },
      children: [
        Component.element(
          tag: 'div',
          attributes: {
            'class': 'ritual-scene',
          },
          children: [
            Component.element(
              tag: 'div',
              key: ValueKey('ritual-text-${phase.name}-$_phaseTick'),
              attributes: {
                'class': 'ritual-text',
              },
              children: [
                Component.text(phase.text),
              ],
            ),

            if (showAffordance)
              Component.element(
                tag: 'div',
                attributes: {
                  'class': 'ritual-affordance',
                },
                children: [
                  Component.element(
                    tag: 'span',
                    attributes: {
                      'class': 'ritual-affordance-dot',
                    },
                  ),
                  Component.element(
                    tag: 'span',
                    attributes: {
                      'class': 'ritual-affordance-text',
                    },
                    children: const [
                      Component.text('Tap to continue'),
                    ],
                  ),
                ],
              ),
          ],
        ),
      ],
    );
  }
}
