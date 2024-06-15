import 'package:freezed_annotation/freezed_annotation.dart';

part 'chat_view_event.freezed.dart';

@freezed
abstract class ChatViewEvent with _$ChatViewEvent {
  const factory ChatViewEvent.joinRoom() = ChatViewEventJoinRoom;
  const factory ChatViewEvent.leaveRoom() = ChatViewEventLeaveRoom;
  const factory ChatViewEvent.error(Exception e) = ChatViewEventError;
}
