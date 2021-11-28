// source: apiKeyServer.proto
/**
 * @fileoverview
 * @enhanceable
 * @suppress {messageConventions} JS Compiler reports an error if a variable or
 *     field starts with 'MSG_' and isn't a translatable message.
 * @public
 */
// GENERATED CODE -- DO NOT EDIT!

goog.provide('proto.apikeyserver.KeyDetailsResponse');

goog.require('jspb.BinaryReader');
goog.require('jspb.BinaryWriter');
goog.require('jspb.Message');

/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.apikeyserver.KeyDetailsResponse = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.apikeyserver.KeyDetailsResponse, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.apikeyserver.KeyDetailsResponse.displayName = 'proto.apikeyserver.KeyDetailsResponse';
}



if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.apikeyserver.KeyDetailsResponse.prototype.toObject = function(opt_includeInstance) {
  return proto.apikeyserver.KeyDetailsResponse.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.apikeyserver.KeyDetailsResponse} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.apikeyserver.KeyDetailsResponse.toObject = function(includeInstance, msg) {
  var f, obj = {
    name: jspb.Message.getFieldWithDefault(msg, 1, ""),
    types: jspb.Message.getFieldWithDefault(msg, 2, ""),
    uses: jspb.Message.getFieldWithDefault(msg, 3, 0),
    kills: jspb.Message.getFieldWithDefault(msg, 4, 0),
    active: jspb.Message.getBooleanFieldWithDefault(msg, 5, false)
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.apikeyserver.KeyDetailsResponse}
 */
proto.apikeyserver.KeyDetailsResponse.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.apikeyserver.KeyDetailsResponse;
  return proto.apikeyserver.KeyDetailsResponse.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.apikeyserver.KeyDetailsResponse} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.apikeyserver.KeyDetailsResponse}
 */
proto.apikeyserver.KeyDetailsResponse.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {string} */ (reader.readString());
      msg.setName(value);
      break;
    case 2:
      var value = /** @type {string} */ (reader.readString());
      msg.setTypes(value);
      break;
    case 3:
      var value = /** @type {number} */ (reader.readUint32());
      msg.setUses(value);
      break;
    case 4:
      var value = /** @type {number} */ (reader.readUint32());
      msg.setKills(value);
      break;
    case 5:
      var value = /** @type {boolean} */ (reader.readBool());
      msg.setActive(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.apikeyserver.KeyDetailsResponse.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.apikeyserver.KeyDetailsResponse.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.apikeyserver.KeyDetailsResponse} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.apikeyserver.KeyDetailsResponse.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getName();
  if (f.length > 0) {
    writer.writeString(
      1,
      f
    );
  }
  f = message.getTypes();
  if (f.length > 0) {
    writer.writeString(
      2,
      f
    );
  }
  f = message.getUses();
  if (f !== 0) {
    writer.writeUint32(
      3,
      f
    );
  }
  f = message.getKills();
  if (f !== 0) {
    writer.writeUint32(
      4,
      f
    );
  }
  f = message.getActive();
  if (f) {
    writer.writeBool(
      5,
      f
    );
  }
};


/**
 * optional string name = 1;
 * @return {string}
 */
proto.apikeyserver.KeyDetailsResponse.prototype.getName = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 1, ""));
};


/**
 * @param {string} value
 * @return {!proto.apikeyserver.KeyDetailsResponse} returns this
 */
proto.apikeyserver.KeyDetailsResponse.prototype.setName = function(value) {
  return jspb.Message.setProto3StringField(this, 1, value);
};


/**
 * optional string types = 2;
 * @return {string}
 */
proto.apikeyserver.KeyDetailsResponse.prototype.getTypes = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 2, ""));
};


/**
 * @param {string} value
 * @return {!proto.apikeyserver.KeyDetailsResponse} returns this
 */
proto.apikeyserver.KeyDetailsResponse.prototype.setTypes = function(value) {
  return jspb.Message.setProto3StringField(this, 2, value);
};


/**
 * optional uint32 uses = 3;
 * @return {number}
 */
proto.apikeyserver.KeyDetailsResponse.prototype.getUses = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 3, 0));
};


/**
 * @param {number} value
 * @return {!proto.apikeyserver.KeyDetailsResponse} returns this
 */
proto.apikeyserver.KeyDetailsResponse.prototype.setUses = function(value) {
  return jspb.Message.setProto3IntField(this, 3, value);
};


/**
 * optional uint32 kills = 4;
 * @return {number}
 */
proto.apikeyserver.KeyDetailsResponse.prototype.getKills = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 4, 0));
};


/**
 * @param {number} value
 * @return {!proto.apikeyserver.KeyDetailsResponse} returns this
 */
proto.apikeyserver.KeyDetailsResponse.prototype.setKills = function(value) {
  return jspb.Message.setProto3IntField(this, 4, value);
};


/**
 * optional bool active = 5;
 * @return {boolean}
 */
proto.apikeyserver.KeyDetailsResponse.prototype.getActive = function() {
  return /** @type {boolean} */ (jspb.Message.getBooleanFieldWithDefault(this, 5, false));
};


/**
 * @param {boolean} value
 * @return {!proto.apikeyserver.KeyDetailsResponse} returns this
 */
proto.apikeyserver.KeyDetailsResponse.prototype.setActive = function(value) {
  return jspb.Message.setProto3BooleanField(this, 5, value);
};

