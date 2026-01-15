import 'package:jaspr/jaspr.dart';

import 'state/iris_mode.dart';
import 'ritual/ritual_view.dart';

/// Root application component for Iris.
///
/// Owns global navigation state and orchestrates
/// the high-level experience flow:
/// Ritual → Cards → Observatory.
class IrisApp extends StatefulComponent {
  const IrisApp({super.key});

  @override
  State<IrisApp> createState() => _IrisAppState();
}

class _IrisAppState extends State<IrisApp> {
  IrisMode _mode = IrisMode.ritual;

  void _advance() {
    setState(() {
      _mode = switch (_mode) {
        IrisMode.ritual => IrisMode.cards,
        IrisMode.cards => IrisMode.observatory,
        IrisMode.observatory => IrisMode.observatory,
      };
    });
  }

  @override
  Component build(BuildContext context) {
    return switch (_mode) {
      IrisMode.ritual => RitualView(onComplete: _advance),
      IrisMode.cards => Component.text('CardsView'),
      IrisMode.observatory => Component.text('Observatory'),
    };
  }
}
