// GENERATED CODE - DO NOT MODIFY BY HAND

part of 'socket_message.dart';

// **************************************************************************
// JsonSerializableGenerator
// **************************************************************************

_$SocketMessageImpl _$$SocketMessageImplFromJson(Map<String, dynamic> json) =>
    _$SocketMessageImpl(
      messageType: const SocketMessageTypeConverter()
          .fromJson((json['message_type'] as num).toInt()),
      senderId: json['sender_id'] as String,
      content: json['content'] as String,
      timestamp:
          const CustomDateTimeConverter().fromJson(json['timestamp'] as String),
    );

Map<String, dynamic> _$$SocketMessageImplToJson(_$SocketMessageImpl instance) =>
    <String, dynamic>{
      'message_type':
          const SocketMessageTypeConverter().toJson(instance.messageType),
      'sender_id': instance.senderId,
      'content': instance.content,
      'timestamp': const CustomDateTimeConverter().toJson(instance.timestamp),
    };
