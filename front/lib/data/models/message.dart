import 'package:freezed_annotation/freezed_annotation.dart';

part 'message.freezed.dart';

@freezed
class Message with _$Message {
  factory Message(
      {required bool mine,
      required String content,
      required DateTime timestamp}) = _Message;
}
