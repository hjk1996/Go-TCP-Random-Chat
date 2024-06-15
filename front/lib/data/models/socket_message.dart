import 'package:freezed_annotation/freezed_annotation.dart';
import 'package:intl/intl.dart';

part 'socket_message.freezed.dart';
part 'socket_message.g.dart';

enum SocketMessageType {
  createRoom,
  joinRoom,
  leaveRoom,
  sendMessage,
  error,
}

extension SocketMessageTypeExtension on SocketMessageType {
  int get value {
    switch (this) {
      case SocketMessageType.createRoom:
        return 0;
      case SocketMessageType.joinRoom:
        return 1;
      case SocketMessageType.leaveRoom:
        return 2;
      case SocketMessageType.sendMessage:
        return 3;
      case SocketMessageType.error:
        return 4;
    }
  }

  static SocketMessageType fromValue(int value) {
    switch (value) {
      case 0:
        return SocketMessageType.createRoom;
      case 1:
        return SocketMessageType.joinRoom;
      case 2:
        return SocketMessageType.leaveRoom;
      case 3:
        return SocketMessageType.sendMessage;
      case 4:
        return SocketMessageType.error;
      default:
        throw ArgumentError("Invalid message type value");
    }
  }
}




@freezed
class SocketMessage with _$SocketMessage {
  factory SocketMessage({
    @JsonKey(name: "message_type") required int messageType,
    @JsonKey(name: "sender_id") required String senderId,
    required String content,
    @CustomDateTimeConverter() required DateTime timestamp,
  }) = _SocketMessage;
  factory SocketMessage.fromJson(Map<String, dynamic> json) =>
    _$SocketMessageFromJson(json);
}


class CustomDateTimeConverter implements JsonConverter<DateTime, String> {
  const CustomDateTimeConverter();

  @override
  DateTime fromJson(String json) {
    return DateFormat("yyyy-MM-ddTHH:mm:ss.SSSSSSZ")
        .parse(json, true)
        .toLocal();
  }

  @override
  String toJson(DateTime object) {
    return object.toUtc().toIso8601String();
  }
}
