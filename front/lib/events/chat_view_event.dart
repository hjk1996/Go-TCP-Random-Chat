import 'package:freezed_annotation/freezed_annotation.dart';

part 'chat_view_event.freezed.dart';

@freezed
abstract class ChatViewEvent with _$ChatViewEvent {
  const factory ChatViewEvent.joinRoomConfirm() = ChatViewEventJoinRoomConfirm;
  const factory ChatViewEvent.opponentJoinRoom() =
      ChatViewEventOpponentJoinRoom;
  const factory ChatViewEvent.leaveRoomConfirm() =
      ChatViewEventLeaveRoomConfirm;
  const factory ChatViewEvent.opponentLeaveRoom() =
      ChatViewEventOpponetLeaveRoom;
  const factory ChatViewEvent.opponentSendMessage() =
      ChatViewEventOpponentSendMessage;
  const factory ChatViewEvent.sendMessage() = ChatViewEventSendMessage;
  const factory ChatViewEvent.error(Exception e) = ChatViewEventError;
}
