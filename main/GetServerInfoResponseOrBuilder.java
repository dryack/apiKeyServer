// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: apikeyserver/apiKeyServer.proto

package main;

public interface GetServerInfoResponseOrBuilder extends
    // @@protoc_insertion_point(interface_extends:apikeyserver.GetServerInfoResponse)
    com.google.protobuf.MessageOrBuilder {

  /**
   * <pre>
   * keys.ServerVersion
   * </pre>
   *
   * <code>string serverVersion = 1;</code>
   * @return The serverVersion.
   */
  java.lang.String getServerVersion();
  /**
   * <pre>
   * keys.ServerVersion
   * </pre>
   *
   * <code>string serverVersion = 1;</code>
   * @return The bytes for serverVersion.
   */
  com.google.protobuf.ByteString
      getServerVersionBytes();

  /**
   * <pre>
   * keys.TotalExhaustions
   * </pre>
   *
   * <code>uint32 keyExhaustions = 2;</code>
   * @return The keyExhaustions.
   */
  int getKeyExhaustions();

  /**
   * <pre>
   * keys.TotalPerMinute
   * </pre>
   *
   * <code>uint64 totalAvailableUsesPerMin = 3;</code>
   * @return The totalAvailableUsesPerMin.
   */
  long getTotalAvailableUsesPerMin();

  /**
   * <pre>
   * keys.TotalKeysServed
   * </pre>
   *
   * <code>uint64 totalKeysServed = 4;</code>
   * @return The totalKeysServed.
   */
  long getTotalKeysServed();

  /**
   * <pre>
   * computed
   * </pre>
   *
   * <code>uint64 totalKeysKilled = 5;</code>
   * @return The totalKeysKilled.
   */
  long getTotalKeysKilled();

  /**
   * <code>string keyNamesPermaKilled = 6;</code>
   * @return The keyNamesPermaKilled.
   */
  java.lang.String getKeyNamesPermaKilled();
  /**
   * <code>string keyNamesPermaKilled = 6;</code>
   * @return The bytes for keyNamesPermaKilled.
   */
  com.google.protobuf.ByteString
      getKeyNamesPermaKilledBytes();

  /**
   * <code>repeated .apikeyserver.KeyDetailsResponse items = 7;</code>
   */
  java.util.List<main.KeyDetailsResponse> 
      getItemsList();
  /**
   * <code>repeated .apikeyserver.KeyDetailsResponse items = 7;</code>
   */
  main.KeyDetailsResponse getItems(int index);
  /**
   * <code>repeated .apikeyserver.KeyDetailsResponse items = 7;</code>
   */
  int getItemsCount();
  /**
   * <code>repeated .apikeyserver.KeyDetailsResponse items = 7;</code>
   */
  java.util.List<? extends main.KeyDetailsResponseOrBuilder> 
      getItemsOrBuilderList();
  /**
   * <code>repeated .apikeyserver.KeyDetailsResponse items = 7;</code>
   */
  main.KeyDetailsResponseOrBuilder getItemsOrBuilder(
      int index);

  /**
   * <code>int64 time = 8;</code>
   * @return The time.
   */
  long getTime();

  /**
   * <pre>
   * computed
   * </pre>
   *
   * <code>int64 uptime = 9;</code>
   * @return The uptime.
   */
  long getUptime();

  /**
   * <pre>
   * computed
   * </pre>
   *
   * <code>float avgKeysServedPerMin = 10;</code>
   * @return The avgKeysServedPerMin.
   */
  float getAvgKeysServedPerMin();
}
