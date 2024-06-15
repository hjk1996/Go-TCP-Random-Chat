// coverage:ignore-file
// GENERATED CODE - DO NOT MODIFY BY HAND
// ignore_for_file: type=lint
// ignore_for_file: unused_element, deprecated_member_use, deprecated_member_use_from_same_package, use_function_type_syntax_for_parameters, unnecessary_const, avoid_init_to_null, invalid_override_different_default_values_named, prefer_expression_function_bodies, annotate_overrides, invalid_annotation_target, unnecessary_question_mark

part of 'chat_view_event.dart';

// **************************************************************************
// FreezedGenerator
// **************************************************************************

T _$identity<T>(T value) => value;

final _privateConstructorUsedError = UnsupportedError(
    'It seems like you constructed your class using `MyClass._()`. This constructor is only meant to be used by freezed and you are not supposed to need it nor use it.\nPlease check the documentation here for more information: https://github.com/rrousselGit/freezed#adding-getters-and-methods-to-our-models');

/// @nodoc
mixin _$ChatViewEvent {
  @optionalTypeArgs
  TResult when<TResult extends Object?>({
    required TResult Function() joinRoom,
    required TResult Function() leaveRoom,
    required TResult Function(Exception e) error,
  }) =>
      throw _privateConstructorUsedError;
  @optionalTypeArgs
  TResult? whenOrNull<TResult extends Object?>({
    TResult? Function()? joinRoom,
    TResult? Function()? leaveRoom,
    TResult? Function(Exception e)? error,
  }) =>
      throw _privateConstructorUsedError;
  @optionalTypeArgs
  TResult maybeWhen<TResult extends Object?>({
    TResult Function()? joinRoom,
    TResult Function()? leaveRoom,
    TResult Function(Exception e)? error,
    required TResult orElse(),
  }) =>
      throw _privateConstructorUsedError;
  @optionalTypeArgs
  TResult map<TResult extends Object?>({
    required TResult Function(ChatViewEventJoinRoom value) joinRoom,
    required TResult Function(ChatViewEventLeaveRoom value) leaveRoom,
    required TResult Function(ChatViewEventError value) error,
  }) =>
      throw _privateConstructorUsedError;
  @optionalTypeArgs
  TResult? mapOrNull<TResult extends Object?>({
    TResult? Function(ChatViewEventJoinRoom value)? joinRoom,
    TResult? Function(ChatViewEventLeaveRoom value)? leaveRoom,
    TResult? Function(ChatViewEventError value)? error,
  }) =>
      throw _privateConstructorUsedError;
  @optionalTypeArgs
  TResult maybeMap<TResult extends Object?>({
    TResult Function(ChatViewEventJoinRoom value)? joinRoom,
    TResult Function(ChatViewEventLeaveRoom value)? leaveRoom,
    TResult Function(ChatViewEventError value)? error,
    required TResult orElse(),
  }) =>
      throw _privateConstructorUsedError;
}

/// @nodoc
abstract class $ChatViewEventCopyWith<$Res> {
  factory $ChatViewEventCopyWith(
          ChatViewEvent value, $Res Function(ChatViewEvent) then) =
      _$ChatViewEventCopyWithImpl<$Res, ChatViewEvent>;
}

/// @nodoc
class _$ChatViewEventCopyWithImpl<$Res, $Val extends ChatViewEvent>
    implements $ChatViewEventCopyWith<$Res> {
  _$ChatViewEventCopyWithImpl(this._value, this._then);

  // ignore: unused_field
  final $Val _value;
  // ignore: unused_field
  final $Res Function($Val) _then;
}

/// @nodoc
abstract class _$$ChatViewEventJoinRoomImplCopyWith<$Res> {
  factory _$$ChatViewEventJoinRoomImplCopyWith(
          _$ChatViewEventJoinRoomImpl value,
          $Res Function(_$ChatViewEventJoinRoomImpl) then) =
      __$$ChatViewEventJoinRoomImplCopyWithImpl<$Res>;
}

/// @nodoc
class __$$ChatViewEventJoinRoomImplCopyWithImpl<$Res>
    extends _$ChatViewEventCopyWithImpl<$Res, _$ChatViewEventJoinRoomImpl>
    implements _$$ChatViewEventJoinRoomImplCopyWith<$Res> {
  __$$ChatViewEventJoinRoomImplCopyWithImpl(_$ChatViewEventJoinRoomImpl _value,
      $Res Function(_$ChatViewEventJoinRoomImpl) _then)
      : super(_value, _then);
}

/// @nodoc

class _$ChatViewEventJoinRoomImpl implements ChatViewEventJoinRoom {
  const _$ChatViewEventJoinRoomImpl();

  @override
  String toString() {
    return 'ChatViewEvent.joinRoom()';
  }

  @override
  bool operator ==(Object other) {
    return identical(this, other) ||
        (other.runtimeType == runtimeType &&
            other is _$ChatViewEventJoinRoomImpl);
  }

  @override
  int get hashCode => runtimeType.hashCode;

  @override
  @optionalTypeArgs
  TResult when<TResult extends Object?>({
    required TResult Function() joinRoom,
    required TResult Function() leaveRoom,
    required TResult Function(Exception e) error,
  }) {
    return joinRoom();
  }

  @override
  @optionalTypeArgs
  TResult? whenOrNull<TResult extends Object?>({
    TResult? Function()? joinRoom,
    TResult? Function()? leaveRoom,
    TResult? Function(Exception e)? error,
  }) {
    return joinRoom?.call();
  }

  @override
  @optionalTypeArgs
  TResult maybeWhen<TResult extends Object?>({
    TResult Function()? joinRoom,
    TResult Function()? leaveRoom,
    TResult Function(Exception e)? error,
    required TResult orElse(),
  }) {
    if (joinRoom != null) {
      return joinRoom();
    }
    return orElse();
  }

  @override
  @optionalTypeArgs
  TResult map<TResult extends Object?>({
    required TResult Function(ChatViewEventJoinRoom value) joinRoom,
    required TResult Function(ChatViewEventLeaveRoom value) leaveRoom,
    required TResult Function(ChatViewEventError value) error,
  }) {
    return joinRoom(this);
  }

  @override
  @optionalTypeArgs
  TResult? mapOrNull<TResult extends Object?>({
    TResult? Function(ChatViewEventJoinRoom value)? joinRoom,
    TResult? Function(ChatViewEventLeaveRoom value)? leaveRoom,
    TResult? Function(ChatViewEventError value)? error,
  }) {
    return joinRoom?.call(this);
  }

  @override
  @optionalTypeArgs
  TResult maybeMap<TResult extends Object?>({
    TResult Function(ChatViewEventJoinRoom value)? joinRoom,
    TResult Function(ChatViewEventLeaveRoom value)? leaveRoom,
    TResult Function(ChatViewEventError value)? error,
    required TResult orElse(),
  }) {
    if (joinRoom != null) {
      return joinRoom(this);
    }
    return orElse();
  }
}

abstract class ChatViewEventJoinRoom implements ChatViewEvent {
  const factory ChatViewEventJoinRoom() = _$ChatViewEventJoinRoomImpl;
}

/// @nodoc
abstract class _$$ChatViewEventLeaveRoomImplCopyWith<$Res> {
  factory _$$ChatViewEventLeaveRoomImplCopyWith(
          _$ChatViewEventLeaveRoomImpl value,
          $Res Function(_$ChatViewEventLeaveRoomImpl) then) =
      __$$ChatViewEventLeaveRoomImplCopyWithImpl<$Res>;
}

/// @nodoc
class __$$ChatViewEventLeaveRoomImplCopyWithImpl<$Res>
    extends _$ChatViewEventCopyWithImpl<$Res, _$ChatViewEventLeaveRoomImpl>
    implements _$$ChatViewEventLeaveRoomImplCopyWith<$Res> {
  __$$ChatViewEventLeaveRoomImplCopyWithImpl(
      _$ChatViewEventLeaveRoomImpl _value,
      $Res Function(_$ChatViewEventLeaveRoomImpl) _then)
      : super(_value, _then);
}

/// @nodoc

class _$ChatViewEventLeaveRoomImpl implements ChatViewEventLeaveRoom {
  const _$ChatViewEventLeaveRoomImpl();

  @override
  String toString() {
    return 'ChatViewEvent.leaveRoom()';
  }

  @override
  bool operator ==(Object other) {
    return identical(this, other) ||
        (other.runtimeType == runtimeType &&
            other is _$ChatViewEventLeaveRoomImpl);
  }

  @override
  int get hashCode => runtimeType.hashCode;

  @override
  @optionalTypeArgs
  TResult when<TResult extends Object?>({
    required TResult Function() joinRoom,
    required TResult Function() leaveRoom,
    required TResult Function(Exception e) error,
  }) {
    return leaveRoom();
  }

  @override
  @optionalTypeArgs
  TResult? whenOrNull<TResult extends Object?>({
    TResult? Function()? joinRoom,
    TResult? Function()? leaveRoom,
    TResult? Function(Exception e)? error,
  }) {
    return leaveRoom?.call();
  }

  @override
  @optionalTypeArgs
  TResult maybeWhen<TResult extends Object?>({
    TResult Function()? joinRoom,
    TResult Function()? leaveRoom,
    TResult Function(Exception e)? error,
    required TResult orElse(),
  }) {
    if (leaveRoom != null) {
      return leaveRoom();
    }
    return orElse();
  }

  @override
  @optionalTypeArgs
  TResult map<TResult extends Object?>({
    required TResult Function(ChatViewEventJoinRoom value) joinRoom,
    required TResult Function(ChatViewEventLeaveRoom value) leaveRoom,
    required TResult Function(ChatViewEventError value) error,
  }) {
    return leaveRoom(this);
  }

  @override
  @optionalTypeArgs
  TResult? mapOrNull<TResult extends Object?>({
    TResult? Function(ChatViewEventJoinRoom value)? joinRoom,
    TResult? Function(ChatViewEventLeaveRoom value)? leaveRoom,
    TResult? Function(ChatViewEventError value)? error,
  }) {
    return leaveRoom?.call(this);
  }

  @override
  @optionalTypeArgs
  TResult maybeMap<TResult extends Object?>({
    TResult Function(ChatViewEventJoinRoom value)? joinRoom,
    TResult Function(ChatViewEventLeaveRoom value)? leaveRoom,
    TResult Function(ChatViewEventError value)? error,
    required TResult orElse(),
  }) {
    if (leaveRoom != null) {
      return leaveRoom(this);
    }
    return orElse();
  }
}

abstract class ChatViewEventLeaveRoom implements ChatViewEvent {
  const factory ChatViewEventLeaveRoom() = _$ChatViewEventLeaveRoomImpl;
}

/// @nodoc
abstract class _$$ChatViewEventErrorImplCopyWith<$Res> {
  factory _$$ChatViewEventErrorImplCopyWith(_$ChatViewEventErrorImpl value,
          $Res Function(_$ChatViewEventErrorImpl) then) =
      __$$ChatViewEventErrorImplCopyWithImpl<$Res>;
  @useResult
  $Res call({Exception e});
}

/// @nodoc
class __$$ChatViewEventErrorImplCopyWithImpl<$Res>
    extends _$ChatViewEventCopyWithImpl<$Res, _$ChatViewEventErrorImpl>
    implements _$$ChatViewEventErrorImplCopyWith<$Res> {
  __$$ChatViewEventErrorImplCopyWithImpl(_$ChatViewEventErrorImpl _value,
      $Res Function(_$ChatViewEventErrorImpl) _then)
      : super(_value, _then);

  @pragma('vm:prefer-inline')
  @override
  $Res call({
    Object? e = null,
  }) {
    return _then(_$ChatViewEventErrorImpl(
      null == e
          ? _value.e
          : e // ignore: cast_nullable_to_non_nullable
              as Exception,
    ));
  }
}

/// @nodoc

class _$ChatViewEventErrorImpl implements ChatViewEventError {
  const _$ChatViewEventErrorImpl(this.e);

  @override
  final Exception e;

  @override
  String toString() {
    return 'ChatViewEvent.error(e: $e)';
  }

  @override
  bool operator ==(Object other) {
    return identical(this, other) ||
        (other.runtimeType == runtimeType &&
            other is _$ChatViewEventErrorImpl &&
            (identical(other.e, e) || other.e == e));
  }

  @override
  int get hashCode => Object.hash(runtimeType, e);

  @JsonKey(ignore: true)
  @override
  @pragma('vm:prefer-inline')
  _$$ChatViewEventErrorImplCopyWith<_$ChatViewEventErrorImpl> get copyWith =>
      __$$ChatViewEventErrorImplCopyWithImpl<_$ChatViewEventErrorImpl>(
          this, _$identity);

  @override
  @optionalTypeArgs
  TResult when<TResult extends Object?>({
    required TResult Function() joinRoom,
    required TResult Function() leaveRoom,
    required TResult Function(Exception e) error,
  }) {
    return error(e);
  }

  @override
  @optionalTypeArgs
  TResult? whenOrNull<TResult extends Object?>({
    TResult? Function()? joinRoom,
    TResult? Function()? leaveRoom,
    TResult? Function(Exception e)? error,
  }) {
    return error?.call(e);
  }

  @override
  @optionalTypeArgs
  TResult maybeWhen<TResult extends Object?>({
    TResult Function()? joinRoom,
    TResult Function()? leaveRoom,
    TResult Function(Exception e)? error,
    required TResult orElse(),
  }) {
    if (error != null) {
      return error(e);
    }
    return orElse();
  }

  @override
  @optionalTypeArgs
  TResult map<TResult extends Object?>({
    required TResult Function(ChatViewEventJoinRoom value) joinRoom,
    required TResult Function(ChatViewEventLeaveRoom value) leaveRoom,
    required TResult Function(ChatViewEventError value) error,
  }) {
    return error(this);
  }

  @override
  @optionalTypeArgs
  TResult? mapOrNull<TResult extends Object?>({
    TResult? Function(ChatViewEventJoinRoom value)? joinRoom,
    TResult? Function(ChatViewEventLeaveRoom value)? leaveRoom,
    TResult? Function(ChatViewEventError value)? error,
  }) {
    return error?.call(this);
  }

  @override
  @optionalTypeArgs
  TResult maybeMap<TResult extends Object?>({
    TResult Function(ChatViewEventJoinRoom value)? joinRoom,
    TResult Function(ChatViewEventLeaveRoom value)? leaveRoom,
    TResult Function(ChatViewEventError value)? error,
    required TResult orElse(),
  }) {
    if (error != null) {
      return error(this);
    }
    return orElse();
  }
}

abstract class ChatViewEventError implements ChatViewEvent {
  const factory ChatViewEventError(final Exception e) =
      _$ChatViewEventErrorImpl;

  Exception get e;
  @JsonKey(ignore: true)
  _$$ChatViewEventErrorImplCopyWith<_$ChatViewEventErrorImpl> get copyWith =>
      throw _privateConstructorUsedError;
}
