// lib/ritual/ritual_phase.dart

enum RitualPhase {
  greeting,
  identity,
  purpose,
  invite,
  completion,
}

class RitualPhaseSpec {
  final String text;
  final int? autoAdvanceDelayMs;
  final bool showTapAffordance;
  final bool isTerminal;

  const RitualPhaseSpec({
    required this.text,
    this.autoAdvanceDelayMs,
    this.showTapAffordance = false,
    this.isTerminal = false,
  });
}

extension RitualPhaseCopy on RitualPhase {
  /// Display and behavior metadata for each phase.
  RitualPhaseSpec get spec => switch (this) {
    RitualPhase.greeting => const RitualPhaseSpec(
      text: 'Hi.',
      autoAdvanceDelayMs: 900,
    ),
    RitualPhase.identity => const RitualPhaseSpec(
      text: "I'm Divij.",
      autoAdvanceDelayMs: 1200,
    ),
    RitualPhase.purpose => const RitualPhaseSpec(
      text: 'I build inner worlds.',
      autoAdvanceDelayMs: 1300,
    ),
    RitualPhase.invite => const RitualPhaseSpec(
      text: 'Step closer.',
      autoAdvanceDelayMs: 1100,
    ),
    RitualPhase.completion => const RitualPhaseSpec(
      text: 'Tap to step inside.',
      showTapAffordance: true,
      isTerminal: true,
    ),
  };

  String get text => spec.text;
  int? get autoAdvanceDelayMs => spec.autoAdvanceDelayMs;
  bool get isTerminal => spec.isTerminal;
  bool get showTapAffordance => spec.showTapAffordance;

  /// Explicit state-machine transitions.
  ///
  /// Keep this readable and easy to extend (branching/replay) later.
  RitualPhase? get next => switch (this) {
    RitualPhase.greeting => RitualPhase.identity,
    RitualPhase.identity => RitualPhase.purpose,
    RitualPhase.purpose => RitualPhase.invite,
    RitualPhase.invite => RitualPhase.completion,
    RitualPhase.completion => null,
  };
}
