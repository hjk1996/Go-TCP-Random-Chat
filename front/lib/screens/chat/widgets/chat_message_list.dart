import 'dart:async';

import 'package:flutter/material.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:front/controllers/chat_view_model.dart';
import 'package:front/events/chat_view_event.dart';
import 'package:intl/intl.dart';
import 'package:front/di.dart';
import 'package:front/data/models/message.dart';

class ChatMessageList extends ConsumerStatefulWidget {
  const ChatMessageList({super.key});

  @override
  ConsumerState<ConsumerStatefulWidget> createState() =>
      _ChatMessageListState();
}

class _ChatMessageListState extends ConsumerState<ChatMessageList> {
  final ScrollController _scrollController = ScrollController();

  StreamSubscription<ChatViewEvent>? _streamSubscription;

  @override
  void initState() {
    super.initState();
    _streamSubscription =
        ref.read(chatViewModelNotifier).eventStream.listen((event) {
      event.whenOrNull(
        sendMessage: () => _scrollToBottom(),
      );
    });
  }

  @override
  void dispose() {
    _streamSubscription?.cancel();
    super.dispose();
  }

  void _scrollToBottom() {
    WidgetsBinding.instance.addPostFrameCallback((_) {
      if (_scrollController.hasClients) {
        _scrollController.animateTo(
          _scrollController.position.maxScrollExtent,
          duration: const Duration(milliseconds: 300),
          curve: Curves.easeOut,
        );
      }
    });
  }

  @override
  Widget build(BuildContext context) {
    final messages = ref.watch(chatViewModelNotifier).messages;
    return ListView.builder(
      controller: _scrollController,
      padding: const EdgeInsets.all(8.0),
      itemCount: messages.length,
      itemBuilder: (context, index) {
        final message = messages[index];
        final isMine = message.mine;
        final alignment =
            isMine ? CrossAxisAlignment.end : CrossAxisAlignment.start;
        final bgColor = isMine ? Colors.blue[200] : Colors.grey[200];
        final textColor = isMine ? Colors.white : Colors.black;

        return Align(
          alignment: isMine ? Alignment.centerRight : Alignment.centerLeft,
          child: Container(
            margin: const EdgeInsets.symmetric(vertical: 4.0),
            padding: const EdgeInsets.all(12.0),
            decoration: BoxDecoration(
              color: bgColor,
              borderRadius: BorderRadius.circular(12.0),
            ),
            child: Column(
              crossAxisAlignment: alignment,
              children: [
                Text(
                  message.content,
                  style: TextStyle(fontSize: 16.0, color: textColor),
                ),
              ],
            ),
          ),
        );
      },
    );
  }
}
