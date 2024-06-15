import 'package:flutter/material.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:front/di.dart';

class ChatMessageList extends ConsumerStatefulWidget {
  const ChatMessageList({super.key});

  @override
  ConsumerState<ConsumerStatefulWidget> createState() =>
      _ChatMessageListState();
}

class _ChatMessageListState extends ConsumerState<ChatMessageList> {
  @override
  Widget build(BuildContext context) {
    final messages = ref.watch(chatViewModelNotifier).messages;
    return ListView.builder(
      padding: EdgeInsets.all(8.0),
      itemCount: messages.length,
      itemBuilder: (context, index) {
        return Container(
          margin: EdgeInsets.symmetric(vertical: 4.0),
          padding: EdgeInsets.all(12.0),
          decoration: BoxDecoration(
            color: index % 2 == 0 ? Colors.grey[200] : Colors.blue[200],
            borderRadius: BorderRadius.circular(12.0),
          ),
          child: Text(
            messages[index].content,
            style: TextStyle(fontSize: 16.0),
          ),
        );
      },
    );
  }
}
