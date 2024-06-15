import 'dart:async';
import 'dart:convert';
import 'dart:io';
import 'dart:typed_data';

import 'package:flutter/material.dart';
import 'package:front/data/models/message.dart';
import 'package:front/data/models/socket_message.dart';
import 'package:front/events/chat_view_event.dart';

class ChatViewModel with ChangeNotifier {
  Socket? _socket;

  List<Message> _messages = [];
  List<Message> get messages => _messages;

  bool _inTheRoom = false;
  bool get inTheRoom => _inTheRoom;

  bool _hasOpponent = false;
  bool get hasOpponent => _hasOpponent;

  final StreamController<ChatViewEvent> _eventController =
      StreamController<ChatViewEvent>.broadcast();

  Stream<ChatViewEvent> get eventStream => _eventController.stream;

  Future<void> _connectToServer() async {
    print("Connecting to server....");
    try {
      _socket = await Socket.connect('localhost', 8888);
      _socket!.listen(
        _onDataReceived,
        onDone: () {
          _handleDisconnection();
        },
        onError: (error) {
          _handleDisconnection();
        },
      );
    } on Exception catch (e) {
      _eventController.sink.add(ChatViewEvent.error(e));
      print(e);
    }
  }

  void _handleDisconnection() {
    _socket = null;
    _inTheRoom = false;
    _hasOpponent = false;
    _messages = [];
    notifyListeners();
  }

  void _onDataReceived(Uint8List data) {
    final msg = _parseRawString(data);
    if (msg == null) {
      return;
    }
    switch (msg.messageType) {
      case SocketMessageType.createRoomConfirm:
        _onCreateRoomConfirm(msg);
      case SocketMessageType.joinRoomConfirm:
        _onJoinRoomConfirm(msg);
      case SocketMessageType.opponentJoinRoom:
        _onOpponentJoinRoom(msg);
      case SocketMessageType.leaveRoomConfirm:
        _onLeaveRoomConfirm(msg);
      case SocketMessageType.opponentLeaveRoom:
        _onOpponentLeaveRoom(msg);
      case SocketMessageType.opponentSendMessage:
        _onOpponentSendMessage(msg);
      case SocketMessageType.error:
        _onError(msg);
    }
  }

  void _onCreateRoomConfirm(SocketMessage msg) {
    _eventController.sink.add(const ChatViewEvent.joinRoomConfirm());
  }

  void _onJoinRoomConfirm(SocketMessage msg) {
    _eventController.sink.add(const ChatViewEvent.joinRoomConfirm());
  }

  void _onOpponentJoinRoom(SocketMessage msg) {
    _hasOpponent = true;
    notifyListeners();
    _eventController.sink.add(const ChatViewEvent.opponentJoinRoom());
  }

  void _onLeaveRoomConfirm(SocketMessage msg) {
    _inTheRoom = false;
    _hasOpponent = false;
    _messages = [];
    notifyListeners();
  }

  void _onOpponentLeaveRoom(SocketMessage msg) {
    _inTheRoom = false;
    _hasOpponent = false;
    _messages = [];
    notifyListeners();
    _eventController.sink.add(const ChatViewEvent.opponentLeaveRoom());
  }

  void _onOpponentSendMessage(SocketMessage msg) {
    print("Somebody sent a message to you..");
    final newMessage =
        Message(mine: false, content: msg.content, timestamp: msg.timestamp);
    messages.add(newMessage);
    notifyListeners();
  }

  void _onError(SocketMessage msg) {
    _eventController.sink.add(ChatViewEvent.error(Exception(msg.content)));
  }

  Future<void> joinRoom() async {
    print("Joining room...");

    if (_socket == null) {
      await _connectToServer();
    }

    _inTheRoom = true;
    _socket!.write("/join\n");

    notifyListeners();
  }

  Future<void> leaveRoom() async {
    print("Leaving room...");

    if (_socket == null) {
      print("Not connected to server...");
      return;
    }

    _inTheRoom = false;
    _hasOpponent = false;
    _socket!.write("/leave\n");
    notifyListeners();
  }

  Future<void> sendMessage(String msg) async {
    print("Sending message...");

    if (_socket == null) {
      print("Not connected to server...");
      _inTheRoom = false;
      notifyListeners();
      return;
    }

    final newMessage =
        Message(mine: true, content: msg, timestamp: DateTime.now());
    _messages.add(newMessage);
    _socket!.write("/msg $msg\n");
    notifyListeners();
    _eventController.sink.add(const ChatViewEvent.sendMessage());
  }

  SocketMessage? _parseRawString(Uint8List data) {
    try {
      var parsedData = jsonDecode(utf8.decode(data)) as Map<String, dynamic>;
      return SocketMessage.fromJson(parsedData);
    } catch (e) {
      return null;
    }
  }
}
