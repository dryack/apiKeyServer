<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: apiKeyServer.proto

namespace Apikeyserver;

use Google\Protobuf\Internal\GPBType;
use Google\Protobuf\Internal\RepeatedField;
use Google\Protobuf\Internal\GPBUtil;

/**
 * request a new key from the server
 *
 * Generated from protobuf message <code>apikeyserver.RequestKey</code>
 */
class RequestKey extends \Google\Protobuf\Internal\Message
{
    /**
     * Generated from protobuf field <code>string requester = 1;</code>
     */
    protected $requester = '';
    /**
     * Generated from protobuf field <code>string type = 2;</code>
     */
    protected $type = '';
    /**
     * requester doesn't want to wait, and instead wishes be told there are no keys available
     *
     * Generated from protobuf field <code>bool acceptExhaustion = 3;</code>
     */
    protected $acceptExhaustion = false;
    /**
     * Generated from protobuf field <code>.google.protobuf.FieldMask field_mask = 4;</code>
     */
    protected $field_mask = null;

    /**
     * Constructor.
     *
     * @param array $data {
     *     Optional. Data for populating the Message object.
     *
     *     @type string $requester
     *     @type string $type
     *     @type bool $acceptExhaustion
     *           requester doesn't want to wait, and instead wishes be told there are no keys available
     *     @type \Google\Protobuf\FieldMask $field_mask
     * }
     */
    public function __construct($data = NULL) {
        \GPBMetadata\ApiKeyServer::initOnce();
        parent::__construct($data);
    }

    /**
     * Generated from protobuf field <code>string requester = 1;</code>
     * @return string
     */
    public function getRequester()
    {
        return $this->requester;
    }

    /**
     * Generated from protobuf field <code>string requester = 1;</code>
     * @param string $var
     * @return $this
     */
    public function setRequester($var)
    {
        GPBUtil::checkString($var, True);
        $this->requester = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>string type = 2;</code>
     * @return string
     */
    public function getType()
    {
        return $this->type;
    }

    /**
     * Generated from protobuf field <code>string type = 2;</code>
     * @param string $var
     * @return $this
     */
    public function setType($var)
    {
        GPBUtil::checkString($var, True);
        $this->type = $var;

        return $this;
    }

    /**
     * requester doesn't want to wait, and instead wishes be told there are no keys available
     *
     * Generated from protobuf field <code>bool acceptExhaustion = 3;</code>
     * @return bool
     */
    public function getAcceptExhaustion()
    {
        return $this->acceptExhaustion;
    }

    /**
     * requester doesn't want to wait, and instead wishes be told there are no keys available
     *
     * Generated from protobuf field <code>bool acceptExhaustion = 3;</code>
     * @param bool $var
     * @return $this
     */
    public function setAcceptExhaustion($var)
    {
        GPBUtil::checkBool($var);
        $this->acceptExhaustion = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>.google.protobuf.FieldMask field_mask = 4;</code>
     * @return \Google\Protobuf\FieldMask
     */
    public function getFieldMask()
    {
        return $this->field_mask;
    }

    /**
     * Generated from protobuf field <code>.google.protobuf.FieldMask field_mask = 4;</code>
     * @param \Google\Protobuf\FieldMask $var
     * @return $this
     */
    public function setFieldMask($var)
    {
        GPBUtil::checkMessage($var, \Google\Protobuf\FieldMask::class);
        $this->field_mask = $var;

        return $this;
    }

}
