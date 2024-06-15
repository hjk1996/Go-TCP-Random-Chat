import 'dart:async';
import 'dart:convert';
import 'dart:io';

import 'package:flutter/material.dart';
import 'package:front/data/models/message.dart';
import 'package:front/events/chat_view_event.dart';

class ChatViewModel with ChangeNotifier {



  Socket? _socket;
  List<Message> _messages = [];
  List<Message> get messages => _messages;

  final StreamController<ChatViewEvent> _eventController =
      StreamController<ChatViewEvent>.broadcast();

  Stream<ChatViewEvent> get eventStream => _eventController.stream;

  Future<void> _connectToServer() async {
    print("Connecting to server....");
    try {
      _socket = await Socket.connect('localhost', 8888);
      _socket!.listen(
        (data) {
          final parsedData = utf8.decode(data);
          print(parsedData);
        },
        onDone: () {
          _handleDisconnection();
        },
        onError: (error) {},
      );
    } catch (e) {
      print(e);
    }
  }

  void _handleDisconnection() {
    _socket = null;
  }

  Future<void> joinRoom() async {
    print("Joining room...");

    if (_socket == null) {
      await _connectToServer();
    }
    _socket!.write("/join");
  }
}
