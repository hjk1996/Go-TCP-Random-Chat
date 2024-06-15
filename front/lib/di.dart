import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:front/controllers/chat_view_model.dart';

late final ChangeNotifierProvider<ChatViewModel> chatViewModelNotifier;

void init() {
  final chatViewModel = ChatViewModel();
  chatViewModelNotifier = ChangeNotifierProvider<ChatViewModel>(
    (ref) => chatViewModel,
  );
}
