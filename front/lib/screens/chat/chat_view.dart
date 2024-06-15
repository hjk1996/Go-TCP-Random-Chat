import 'dart:async';

import 'package:flutter/material.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:front/di.dart';
import 'package:front/events/chat_view_event.dart';
import 'package:front/screens/chat/widgets/chat_input_field.dart';
import 'package:front/screens/chat/widgets/chat_message_list.dart';

class ChatView extends ConsumerStatefulWidget {
  const ChatView({super.key});

  @override
  ConsumerState<ConsumerStatefulWidget> createState() => _ChatViewState();
}

class _ChatViewState extends ConsumerState<ChatView> {
  StreamSubscription<ChatViewEvent>? _streamSubscription;

  @override
  void initState() {
    super.initState();
    _streamSubscription =
        ref.read(chatViewModelNotifier).eventStream.listen((event) {});
  }

  @override
  void dispose() {
    _streamSubscription?.cancel();
    super.dispose();
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: const Text('Chat App'),
        actions: [
          IconButton(
              onPressed: () async => ref.read(chatViewModelNotifier).joinRoom,
              icon: const Icon(Icons.play_arrow))
        ],
      ),
      body: const Column(
        children: <Widget>[
          Expanded(
            child: ChatMessageList(),
          ),
          Padding(padding: const EdgeInsets.all(10.0), child: ChatInputField()),
          SizedBox(
            height: 20,
          )
        ],
      ),
    );
  }
}
