import 'dart:async';

import 'package:flutter/material.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:front/di.dart';
import 'package:front/events/chat_view_event.dart';
import 'package:front/screens/chat/widgets/chat_input_field.dart';
import 'package:front/screens/chat/widgets/chat_message_list.dart';
import 'package:top_snackbar_flutter/custom_snack_bar.dart';
import 'package:top_snackbar_flutter/top_snack_bar.dart';

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
        ref.read(chatViewModelNotifier).eventStream.listen((event) {
      event.whenOrNull(
        joinRoomConfirm: () {},
        opponentJoinRoom: () {
          showTopSnackBar(
              Overlay.of(context),
              const CustomSnackBar.success(
                  message: "A user has join the room"));
        },
        leaveRoomConfirm: () {},
        opponentLeaveRoom: () {
          showTopSnackBar(
              Overlay.of(context),
              const CustomSnackBar.success(
                  message: "A user has left the room"));
        },
        error: (e) {
          print(e);
          ScaffoldMessenger.of(context)
              .showSnackBar(SnackBar(content: Text(e.toString())));
        },
      );
    });
  }

  @override
  void dispose() {
    _streamSubscription?.cancel();
    super.dispose();
  }

  @override
  Widget build(BuildContext context) {
    bool inTheRoom =
        ref.watch(chatViewModelNotifier.select((value) => value.inTheRoom));
    return Scaffold(
      appBar: AppBar(
        title: const Text('Random Chat'),
        actions: [
          IconButton(
              onPressed: inTheRoom
                  ? ref.read(chatViewModelNotifier).leaveRoom
                  : ref.read(chatViewModelNotifier).joinRoom,
              icon: Icon(
                inTheRoom ? Icons.exit_to_app : Icons.play_arrow,
              ))
        ],
      ),
      body: Column(
        children: <Widget>[
          const Expanded(
            child: ChatMessageList(),
          ),
          Visibility(
            visible: inTheRoom,
            child: const ChatInputField(),
          ),
          const SizedBox(
            height: 25,
          ),
        ],
      ),
  );
  }
}
