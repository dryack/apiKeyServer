// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: apikeyserver/apiKeyServer.proto

package main;

public interface RequestPermKillKeyOrBuilder extends
    // @@protoc_insertion_point(interface_extends:apikeyserver.RequestPermKillKey)
    com.google.protobuf.MessageOrBuilder {

  /**
   * <code>string requester = 1;</code>
   * @return The requester.
   */
  java.lang.String getRequester();
  /**
   * <code>string requester = 1;</code>
   * @return The bytes for requester.
   */
  com.google.protobuf.ByteString
      getRequesterBytes();

  /**
   * <code>string key = 2;</code>
   * @return The key.
   */
  java.lang.String getKey();
  /**
   * <code>string key = 2;</code>
   * @return The bytes for key.
   */
  com.google.protobuf.ByteString
      getKeyBytes();

  /**
   * <code>string name = 3;</code>
   * @return The name.
   */
  java.lang.String getName();
  /**
   * <code>string name = 3;</code>
   * @return The bytes for name.
   */
  com.google.protobuf.ByteString
      getNameBytes();

  /**
   * <pre>
   * unixNano
   * </pre>
   *
   * <code>int64 time = 4;</code>
   * @return The time.
   */
  long getTime();
}
