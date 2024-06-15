import 'package:freezed_annotation/freezed_annotation.dart';
import 'package:intl/intl.dart';

part 'socket_message.freezed.dart';
part 'socket_message.g.dart';

enum SocketMessageType {
  createRoomConfirm,
  joinRoomConfirm,
  opponentJoinRoom,
  leaveRoomConfirm,
  opponentLeaveRoom,
  opponentSendMessage,
  error,
}

@freezed
class SocketMessage with _$SocketMessage {
  factory SocketMessage({
    @JsonKey(name: "message_type")
    @SocketMessageTypeConverter()
    required SocketMessageType messageType,
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

class SocketMessageTypeConverter
    implements JsonConverter<SocketMessageType, int> {
  const SocketMessageTypeConverter();

  @override
  SocketMessageType fromJson(int json) {
    return SocketMessageType.values[json];
  }

  @override
  int toJson(SocketMessageType object) {
    return object.index;
  }
}
