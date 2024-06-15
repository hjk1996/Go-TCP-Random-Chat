// coverage:ignore-file
// GENERATED CODE - DO NOT MODIFY BY HAND
// ignore_for_file: type=lint
// ignore_for_file: unused_element, deprecated_member_use, deprecated_member_use_from_same_package, use_function_type_syntax_for_parameters, unnecessary_const, avoid_init_to_null, invalid_override_different_default_values_named, prefer_expression_function_bodies, annotate_overrides, invalid_annotation_target, unnecessary_question_mark

part of 'socket_message.dart';

// **************************************************************************
// FreezedGenerator
// **************************************************************************

T _$identity<T>(T value) => value;

final _privateConstructorUsedError = UnsupportedError(
    'It seems like you constructed your class using `MyClass._()`. This constructor is only meant to be used by freezed and you are not supposed to need it nor use it.\nPlease check the documentation here for more information: https://github.com/rrousselGit/freezed#adding-getters-and-methods-to-our-models');

SocketMessage _$SocketMessageFromJson(Map<String, dynamic> json) {
  return _SocketMessage.fromJson(json);
}

/// @nodoc
mixin _$SocketMessage {
  @JsonKey(name: "message_type")
  int get messageType => throw _privateConstructorUsedError;
  @JsonKey(name: "sender_id")
  String get senderId => throw _privateConstructorUsedError;
  String get content => throw _privateConstructorUsedError;
  @CustomDateTimeConverter()
  DateTime get timestamp => throw _privateConstructorUsedError;

  Map<String, dynamic> toJson() => throw _privateConstructorUsedError;
  @JsonKey(ignore: true)
  $SocketMessageCopyWith<SocketMessage> get copyWith =>
      throw _privateConstructorUsedError;
}

/// @nodoc
abstract class $SocketMessageCopyWith<$Res> {
  factory $SocketMessageCopyWith(
          SocketMessage value, $Res Function(SocketMessage) then) =
      _$SocketMessageCopyWithImpl<$Res, SocketMessage>;
  @useResult
  $Res call(
      {@JsonKey(name: "message_type") int messageType,
      @JsonKey(name: "sender_id") String senderId,
      String content,
      @CustomDateTimeConverter() DateTime timestamp});
}

/// @nodoc
class _$SocketMessageCopyWithImpl<$Res, $Val extends SocketMessage>
    implements $SocketMessageCopyWith<$Res> {
  _$SocketMessageCopyWithImpl(this._value, this._then);

  // ignore: unused_field
  final $Val _value;
  // ignore: unused_field
  final $Res Function($Val) _then;

  @pragma('vm:prefer-inline')
  @override
  $Res call({
    Object? messageType = null,
    Object? senderId = null,
    Object? content = null,
    Object? timestamp = null,
  }) {
    return _then(_value.copyWith(
      messageType: null == messageType
          ? _value.messageType
          : messageType // ignore: cast_nullable_to_non_nullable
              as int,
      senderId: null == senderId
          ? _value.senderId
          : senderId // ignore: cast_nullable_to_non_nullable
              as String,
      content: null == content
          ? _value.content
          : content // ignore: cast_nullable_to_non_nullable
              as String,
      timestamp: null == timestamp
          ? _value.timestamp
          : timestamp // ignore: cast_nullable_to_non_nullable
              as DateTime,
    ) as $Val);
  }
}

/// @nodoc
abstract class _$$SocketMessageImplCopyWith<$Res>
    implements $SocketMessageCopyWith<$Res> {
  factory _$$SocketMessageImplCopyWith(
          _$SocketMessageImpl value, $Res Function(_$SocketMessageImpl) then) =
      __$$SocketMessageImplCopyWithImpl<$Res>;
  @override
  @useResult
  $Res call(
      {@JsonKey(name: "message_type") int messageType,
      @JsonKey(name: "sender_id") String senderId,
      String content,
      @CustomDateTimeConverter() DateTime timestamp});
}

/// @nodoc
class __$$SocketMessageImplCopyWithImpl<$Res>
    extends _$SocketMessageCopyWithImpl<$Res, _$SocketMessageImpl>
    implements _$$SocketMessageImplCopyWith<$Res> {
  __$$SocketMessageImplCopyWithImpl(
      _$SocketMessageImpl _value, $Res Function(_$SocketMessageImpl) _then)
      : super(_value, _then);

  @pragma('vm:prefer-inline')
  @override
  $Res call({
    Object? messageType = null,
    Object? senderId = null,
    Object? content = null,
    Object? timestamp = null,
  }) {
    return _then(_$SocketMessageImpl(
      messageType: null == messageType
          ? _value.messageType
          : messageType // ignore: cast_nullable_to_non_nullable
              as int,
      senderId: null == senderId
          ? _value.senderId
          : senderId // ignore: cast_nullable_to_non_nullable
              as String,
      content: null == content
          ? _value.content
          : content // ignore: cast_nullable_to_non_nullable
              as String,
      timestamp: null == timestamp
          ? _value.timestamp
          : timestamp // ignore: cast_nullable_to_non_nullable
              as DateTime,
    ));
  }
}

/// @nodoc
@JsonSerializable()
class _$SocketMessageImpl implements _SocketMessage {
  _$SocketMessageImpl(
      {@JsonKey(name: "message_type") required this.messageType,
      @JsonKey(name: "sender_id") required this.senderId,
      required this.content,
      @CustomDateTimeConverter() required this.timestamp});

  factory _$SocketMessageImpl.fromJson(Map<String, dynamic> json) =>
      _$$SocketMessageImplFromJson(json);

  @override
  @JsonKey(name: "message_type")
  final int messageType;
  @override
  @JsonKey(name: "sender_id")
  final String senderId;
  @override
  final String content;
  @override
  @CustomDateTimeConverter()
  final DateTime timestamp;

  @override
  String toString() {
    return 'SocketMessage(messageType: $messageType, senderId: $senderId, content: $content, timestamp: $timestamp)';
  }

  @override
  bool operator ==(Object other) {
    return identical(this, other) ||
        (other.runtimeType == runtimeType &&
            other is _$SocketMessageImpl &&
            (identical(other.messageType, messageType) ||
                other.messageType == messageType) &&
            (identical(other.senderId, senderId) ||
                other.senderId == senderId) &&
            (identical(other.content, content) || other.content == content) &&
            (identical(other.timestamp, timestamp) ||
                other.timestamp == timestamp));
  }

  @JsonKey(ignore: true)
  @override
  int get hashCode =>
      Object.hash(runtimeType, messageType, senderId, content, timestamp);

  @JsonKey(ignore: true)
  @override
  @pragma('vm:prefer-inline')
  _$$SocketMessageImplCopyWith<_$SocketMessageImpl> get copyWith =>
      __$$SocketMessageImplCopyWithImpl<_$SocketMessageImpl>(this, _$identity);

  @override
  Map<String, dynamic> toJson() {
    return _$$SocketMessageImplToJson(
      this,
    );
  }
}

abstract class _SocketMessage implements SocketMessage {
  factory _SocketMessage(
          {@JsonKey(name: "message_type") required final int messageType,
          @JsonKey(name: "sender_id") required final String senderId,
          required final String content,
          @CustomDateTimeConverter() required final DateTime timestamp}) =
      _$SocketMessageImpl;

  factory _SocketMessage.fromJson(Map<String, dynamic> json) =
      _$SocketMessageImpl.fromJson;

  @override
  @JsonKey(name: "message_type")
  int get messageType;
  @override
  @JsonKey(name: "sender_id")
  String get senderId;
  @override
  String get content;
  @override
  @CustomDateTimeConverter()
  DateTime get timestamp;
  @override
  @JsonKey(ignore: true)
  _$$SocketMessageImplCopyWith<_$SocketMessageImpl> get copyWith =>
      throw _privateConstructorUsedError;
}
