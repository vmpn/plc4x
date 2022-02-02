/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

package model

import (
	"github.com/apache/plc4x/plc4go/internal/plc4go/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

type MQTT_ReasonCode uint8

type IMQTT_ReasonCode interface {
	DisconnectReason() bool
	UnsubackResponse() bool
	SubackResponse() bool
	AuthReason() bool
	PubrelPubcompResponse() bool
	PubackPubrecResponse() bool
	ConnackResponse() bool
	Serialize(writeBuffer utils.WriteBuffer) error
}

const(
	MQTT_ReasonCode_SUCCESS MQTT_ReasonCode = 0X00
	MQTT_ReasonCode_GRANTED_QOS_1 MQTT_ReasonCode = 0X01
	MQTT_ReasonCode_GRANTED_QOS_2 MQTT_ReasonCode = 0X02
	MQTT_ReasonCode_DISCONNECT_WITH_WILL_MESSAGE MQTT_ReasonCode = 0X04
	MQTT_ReasonCode_NO_MATCHING_SUBSCRIBERS MQTT_ReasonCode = 0x10
	MQTT_ReasonCode_NO_SUBSCRIPTION_EXISTED MQTT_ReasonCode = 0x11
	MQTT_ReasonCode_CONTINUE_AUTHENTICATION MQTT_ReasonCode = 0X18
	MQTT_ReasonCode_RE_AUTHENTICATE MQTT_ReasonCode = 0X19
	MQTT_ReasonCode_UNSPECIFIED_ERROR MQTT_ReasonCode = 0X80
	MQTT_ReasonCode_MALFORMED_PACKET MQTT_ReasonCode = 0X81
	MQTT_ReasonCode_PROTOCOL_ERROR MQTT_ReasonCode = 0X82
	MQTT_ReasonCode_IMPLEMENTATION_SPECIFIC_ERROR MQTT_ReasonCode = 0X83
	MQTT_ReasonCode_UNSUPPORTED_PROTOCOL_VERSION MQTT_ReasonCode = 0X84
	MQTT_ReasonCode_CLIENT_IDENTIFIER_NOT_VALID MQTT_ReasonCode = 0X85
	MQTT_ReasonCode_BAD_USER_NAME_OR_PASSWORD MQTT_ReasonCode = 0X86
	MQTT_ReasonCode_NOT_AUTHORIZED MQTT_ReasonCode = 0X87
	MQTT_ReasonCode_SERVER_UNAVAILABLE MQTT_ReasonCode = 0X88
	MQTT_ReasonCode_SERVER_BUSY MQTT_ReasonCode = 0X89
	MQTT_ReasonCode_BANNED MQTT_ReasonCode = 0X8A
	MQTT_ReasonCode_SERVER_SHUTTING_DOWN MQTT_ReasonCode = 0X8B
	MQTT_ReasonCode_BAD_AUTHENTICATION_METHOD MQTT_ReasonCode = 0X8C
	MQTT_ReasonCode_KEEP_ALIVE_TIMEOUT MQTT_ReasonCode = 0X8D
	MQTT_ReasonCode_SESSION_TAKEN_OVER MQTT_ReasonCode = 0X8E
	MQTT_ReasonCode_TOPIC_FILTER_INVALID MQTT_ReasonCode = 0X8F
	MQTT_ReasonCode_TOPIC_NAME_INVALID MQTT_ReasonCode = 0X90
	MQTT_ReasonCode_PACKET_IDENTIFIER_IN_USE MQTT_ReasonCode = 0X91
	MQTT_ReasonCode_PACKET_IDENTIFIER_NOT_FOUND MQTT_ReasonCode = 0X92
	MQTT_ReasonCode_RECEIVE_MAXIMUM_EXCEEDED MQTT_ReasonCode = 0X93
	MQTT_ReasonCode_TOPIC_ALIAS_INVALID MQTT_ReasonCode = 0X94
	MQTT_ReasonCode_PACKET_TOO_LARGE MQTT_ReasonCode = 0X95
	MQTT_ReasonCode_MESSAGE_RATE_TO_HIGH MQTT_ReasonCode = 0X96
	MQTT_ReasonCode_QUOTA_EXCEEDED MQTT_ReasonCode = 0X97
	MQTT_ReasonCode_ADMINISTRATIVE_ACTION MQTT_ReasonCode = 0X98
	MQTT_ReasonCode_PAYLOAD_FORMAT_INVALID MQTT_ReasonCode = 0X99
	MQTT_ReasonCode_RETAIN_NOT_SUPPORTED MQTT_ReasonCode = 0X9A
	MQTT_ReasonCode_QOS_NOT_SUPPORTED MQTT_ReasonCode = 0X9B
	MQTT_ReasonCode_USE_ANOTHER_SERVER MQTT_ReasonCode = 0X9C
	MQTT_ReasonCode_SERVER_MOVED MQTT_ReasonCode = 0X9D
	MQTT_ReasonCode_SHARED_SUBSCRIPTIONS_NOT_SUPPORTED MQTT_ReasonCode = 0X9E
	MQTT_ReasonCode_CONNECTION_RATE_EXCEEDED MQTT_ReasonCode = 0X9F
	MQTT_ReasonCode_MAXIMUM_CONNECT_TIME MQTT_ReasonCode = 0XA0
	MQTT_ReasonCode_SUBSCRIPTION_IDENTIFIERS_NOT_SUPPORTED MQTT_ReasonCode = 0XA1
	MQTT_ReasonCode_WILDCARD_SUBSCRIPTIONS_NOT_SUPPORTED MQTT_ReasonCode = 0XA2
)

var MQTT_ReasonCodeValues []MQTT_ReasonCode

func init() {
	_ = errors.New
	MQTT_ReasonCodeValues = []MQTT_ReasonCode {
		MQTT_ReasonCode_SUCCESS,
		MQTT_ReasonCode_GRANTED_QOS_1,
		MQTT_ReasonCode_GRANTED_QOS_2,
		MQTT_ReasonCode_DISCONNECT_WITH_WILL_MESSAGE,
		MQTT_ReasonCode_NO_MATCHING_SUBSCRIBERS,
		MQTT_ReasonCode_NO_SUBSCRIPTION_EXISTED,
		MQTT_ReasonCode_CONTINUE_AUTHENTICATION,
		MQTT_ReasonCode_RE_AUTHENTICATE,
		MQTT_ReasonCode_UNSPECIFIED_ERROR,
		MQTT_ReasonCode_MALFORMED_PACKET,
		MQTT_ReasonCode_PROTOCOL_ERROR,
		MQTT_ReasonCode_IMPLEMENTATION_SPECIFIC_ERROR,
		MQTT_ReasonCode_UNSUPPORTED_PROTOCOL_VERSION,
		MQTT_ReasonCode_CLIENT_IDENTIFIER_NOT_VALID,
		MQTT_ReasonCode_BAD_USER_NAME_OR_PASSWORD,
		MQTT_ReasonCode_NOT_AUTHORIZED,
		MQTT_ReasonCode_SERVER_UNAVAILABLE,
		MQTT_ReasonCode_SERVER_BUSY,
		MQTT_ReasonCode_BANNED,
		MQTT_ReasonCode_SERVER_SHUTTING_DOWN,
		MQTT_ReasonCode_BAD_AUTHENTICATION_METHOD,
		MQTT_ReasonCode_KEEP_ALIVE_TIMEOUT,
		MQTT_ReasonCode_SESSION_TAKEN_OVER,
		MQTT_ReasonCode_TOPIC_FILTER_INVALID,
		MQTT_ReasonCode_TOPIC_NAME_INVALID,
		MQTT_ReasonCode_PACKET_IDENTIFIER_IN_USE,
		MQTT_ReasonCode_PACKET_IDENTIFIER_NOT_FOUND,
		MQTT_ReasonCode_RECEIVE_MAXIMUM_EXCEEDED,
		MQTT_ReasonCode_TOPIC_ALIAS_INVALID,
		MQTT_ReasonCode_PACKET_TOO_LARGE,
		MQTT_ReasonCode_MESSAGE_RATE_TO_HIGH,
		MQTT_ReasonCode_QUOTA_EXCEEDED,
		MQTT_ReasonCode_ADMINISTRATIVE_ACTION,
		MQTT_ReasonCode_PAYLOAD_FORMAT_INVALID,
		MQTT_ReasonCode_RETAIN_NOT_SUPPORTED,
		MQTT_ReasonCode_QOS_NOT_SUPPORTED,
		MQTT_ReasonCode_USE_ANOTHER_SERVER,
		MQTT_ReasonCode_SERVER_MOVED,
		MQTT_ReasonCode_SHARED_SUBSCRIPTIONS_NOT_SUPPORTED,
		MQTT_ReasonCode_CONNECTION_RATE_EXCEEDED,
		MQTT_ReasonCode_MAXIMUM_CONNECT_TIME,
		MQTT_ReasonCode_SUBSCRIPTION_IDENTIFIERS_NOT_SUPPORTED,
		MQTT_ReasonCode_WILDCARD_SUBSCRIPTIONS_NOT_SUPPORTED,
	}
}


func (e MQTT_ReasonCode) DisconnectReason() bool {
	switch e  {
		case 0X00: { /* '0X00' */
            return true
		}
		case 0X01: { /* '0X01' */
            return false
		}
		case 0X02: { /* '0X02' */
            return false
		}
		case 0X04: { /* '0X04' */
            return true
		}
		case 0X18: { /* '0X18' */
            return false
		}
		case 0X19: { /* '0X19' */
            return false
		}
		case 0X80: { /* '0X80' */
            return true
		}
		case 0X81: { /* '0X81' */
            return true
		}
		case 0X82: { /* '0X82' */
            return true
		}
		case 0X83: { /* '0X83' */
            return true
		}
		case 0X84: { /* '0X84' */
            return false
		}
		case 0X85: { /* '0X85' */
            return false
		}
		case 0X86: { /* '0X86' */
            return false
		}
		case 0X87: { /* '0X87' */
            return true
		}
		case 0X88: { /* '0X88' */
            return false
		}
		case 0X89: { /* '0X89' */
            return true
		}
		case 0X8A: { /* '0X8A' */
            return false
		}
		case 0X8B: { /* '0X8B' */
            return true
		}
		case 0X8C: { /* '0X8C' */
            return false
		}
		case 0X8D: { /* '0X8D' */
            return true
		}
		case 0X8E: { /* '0X8E' */
            return true
		}
		case 0X8F: { /* '0X8F' */
            return true
		}
		case 0X90: { /* '0X90' */
            return true
		}
		case 0X91: { /* '0X91' */
            return false
		}
		case 0X92: { /* '0X92' */
            return false
		}
		case 0X93: { /* '0X93' */
            return true
		}
		case 0X94: { /* '0X94' */
            return true
		}
		case 0X95: { /* '0X95' */
            return true
		}
		case 0X96: { /* '0X96' */
            return true
		}
		case 0X97: { /* '0X97' */
            return true
		}
		case 0X98: { /* '0X98' */
            return true
		}
		case 0X99: { /* '0X99' */
            return true
		}
		case 0X9A: { /* '0X9A' */
            return true
		}
		case 0X9B: { /* '0X9B' */
            return true
		}
		case 0X9C: { /* '0X9C' */
            return true
		}
		case 0X9D: { /* '0X9D' */
            return true
		}
		case 0X9E: { /* '0X9E' */
            return true
		}
		case 0X9F: { /* '0X9F' */
            return true
		}
		case 0XA0: { /* '0XA0' */
            return true
		}
		case 0XA1: { /* '0XA1' */
            return true
		}
		case 0XA2: { /* '0XA2' */
            return true
		}
		case 0x10: { /* '0x10' */
            return false
		}
		case 0x11: { /* '0x11' */
            return false
		}
		default: {
			return false
		}
	}
}

func MQTT_ReasonCodeFirstEnumForFieldDisconnectReason(value bool) (MQTT_ReasonCode, error) {
	for _, sizeValue := range MQTT_ReasonCodeValues {
		if sizeValue.DisconnectReason() == value {
			return sizeValue, nil
		}
	}
	return 0, errors.Errorf("enum for %v describing DisconnectReason not found", value)
}

func (e MQTT_ReasonCode) UnsubackResponse() bool {
	switch e  {
		case 0X00: { /* '0X00' */
            return true
		}
		case 0X01: { /* '0X01' */
            return false
		}
		case 0X02: { /* '0X02' */
            return false
		}
		case 0X04: { /* '0X04' */
            return false
		}
		case 0X18: { /* '0X18' */
            return false
		}
		case 0X19: { /* '0X19' */
            return false
		}
		case 0X80: { /* '0X80' */
            return true
		}
		case 0X81: { /* '0X81' */
            return false
		}
		case 0X82: { /* '0X82' */
            return false
		}
		case 0X83: { /* '0X83' */
            return true
		}
		case 0X84: { /* '0X84' */
            return false
		}
		case 0X85: { /* '0X85' */
            return false
		}
		case 0X86: { /* '0X86' */
            return false
		}
		case 0X87: { /* '0X87' */
            return true
		}
		case 0X88: { /* '0X88' */
            return false
		}
		case 0X89: { /* '0X89' */
            return false
		}
		case 0X8A: { /* '0X8A' */
            return false
		}
		case 0X8B: { /* '0X8B' */
            return false
		}
		case 0X8C: { /* '0X8C' */
            return false
		}
		case 0X8D: { /* '0X8D' */
            return false
		}
		case 0X8E: { /* '0X8E' */
            return false
		}
		case 0X8F: { /* '0X8F' */
            return true
		}
		case 0X90: { /* '0X90' */
            return false
		}
		case 0X91: { /* '0X91' */
            return true
		}
		case 0X92: { /* '0X92' */
            return false
		}
		case 0X93: { /* '0X93' */
            return false
		}
		case 0X94: { /* '0X94' */
            return false
		}
		case 0X95: { /* '0X95' */
            return false
		}
		case 0X96: { /* '0X96' */
            return false
		}
		case 0X97: { /* '0X97' */
            return false
		}
		case 0X98: { /* '0X98' */
            return false
		}
		case 0X99: { /* '0X99' */
            return false
		}
		case 0X9A: { /* '0X9A' */
            return false
		}
		case 0X9B: { /* '0X9B' */
            return false
		}
		case 0X9C: { /* '0X9C' */
            return false
		}
		case 0X9D: { /* '0X9D' */
            return false
		}
		case 0X9E: { /* '0X9E' */
            return false
		}
		case 0X9F: { /* '0X9F' */
            return false
		}
		case 0XA0: { /* '0XA0' */
            return false
		}
		case 0XA1: { /* '0XA1' */
            return false
		}
		case 0XA2: { /* '0XA2' */
            return false
		}
		case 0x10: { /* '0x10' */
            return false
		}
		case 0x11: { /* '0x11' */
            return true
		}
		default: {
			return false
		}
	}
}

func MQTT_ReasonCodeFirstEnumForFieldUnsubackResponse(value bool) (MQTT_ReasonCode, error) {
	for _, sizeValue := range MQTT_ReasonCodeValues {
		if sizeValue.UnsubackResponse() == value {
			return sizeValue, nil
		}
	}
	return 0, errors.Errorf("enum for %v describing UnsubackResponse not found", value)
}

func (e MQTT_ReasonCode) SubackResponse() bool {
	switch e  {
		case 0X00: { /* '0X00' */
            return true
		}
		case 0X01: { /* '0X01' */
            return true
		}
		case 0X02: { /* '0X02' */
            return true
		}
		case 0X04: { /* '0X04' */
            return false
		}
		case 0X18: { /* '0X18' */
            return false
		}
		case 0X19: { /* '0X19' */
            return false
		}
		case 0X80: { /* '0X80' */
            return true
		}
		case 0X81: { /* '0X81' */
            return false
		}
		case 0X82: { /* '0X82' */
            return false
		}
		case 0X83: { /* '0X83' */
            return true
		}
		case 0X84: { /* '0X84' */
            return false
		}
		case 0X85: { /* '0X85' */
            return false
		}
		case 0X86: { /* '0X86' */
            return false
		}
		case 0X87: { /* '0X87' */
            return true
		}
		case 0X88: { /* '0X88' */
            return false
		}
		case 0X89: { /* '0X89' */
            return false
		}
		case 0X8A: { /* '0X8A' */
            return false
		}
		case 0X8B: { /* '0X8B' */
            return false
		}
		case 0X8C: { /* '0X8C' */
            return false
		}
		case 0X8D: { /* '0X8D' */
            return false
		}
		case 0X8E: { /* '0X8E' */
            return false
		}
		case 0X8F: { /* '0X8F' */
            return true
		}
		case 0X90: { /* '0X90' */
            return false
		}
		case 0X91: { /* '0X91' */
            return true
		}
		case 0X92: { /* '0X92' */
            return false
		}
		case 0X93: { /* '0X93' */
            return false
		}
		case 0X94: { /* '0X94' */
            return false
		}
		case 0X95: { /* '0X95' */
            return false
		}
		case 0X96: { /* '0X96' */
            return false
		}
		case 0X97: { /* '0X97' */
            return true
		}
		case 0X98: { /* '0X98' */
            return false
		}
		case 0X99: { /* '0X99' */
            return false
		}
		case 0X9A: { /* '0X9A' */
            return false
		}
		case 0X9B: { /* '0X9B' */
            return false
		}
		case 0X9C: { /* '0X9C' */
            return false
		}
		case 0X9D: { /* '0X9D' */
            return false
		}
		case 0X9E: { /* '0X9E' */
            return true
		}
		case 0X9F: { /* '0X9F' */
            return false
		}
		case 0XA0: { /* '0XA0' */
            return false
		}
		case 0XA1: { /* '0XA1' */
            return true
		}
		case 0XA2: { /* '0XA2' */
            return true
		}
		case 0x10: { /* '0x10' */
            return false
		}
		case 0x11: { /* '0x11' */
            return false
		}
		default: {
			return false
		}
	}
}

func MQTT_ReasonCodeFirstEnumForFieldSubackResponse(value bool) (MQTT_ReasonCode, error) {
	for _, sizeValue := range MQTT_ReasonCodeValues {
		if sizeValue.SubackResponse() == value {
			return sizeValue, nil
		}
	}
	return 0, errors.Errorf("enum for %v describing SubackResponse not found", value)
}

func (e MQTT_ReasonCode) AuthReason() bool {
	switch e  {
		case 0X00: { /* '0X00' */
            return true
		}
		case 0X01: { /* '0X01' */
            return false
		}
		case 0X02: { /* '0X02' */
            return false
		}
		case 0X04: { /* '0X04' */
            return false
		}
		case 0X18: { /* '0X18' */
            return true
		}
		case 0X19: { /* '0X19' */
            return true
		}
		case 0X80: { /* '0X80' */
            return false
		}
		case 0X81: { /* '0X81' */
            return false
		}
		case 0X82: { /* '0X82' */
            return false
		}
		case 0X83: { /* '0X83' */
            return false
		}
		case 0X84: { /* '0X84' */
            return false
		}
		case 0X85: { /* '0X85' */
            return false
		}
		case 0X86: { /* '0X86' */
            return false
		}
		case 0X87: { /* '0X87' */
            return false
		}
		case 0X88: { /* '0X88' */
            return false
		}
		case 0X89: { /* '0X89' */
            return false
		}
		case 0X8A: { /* '0X8A' */
            return false
		}
		case 0X8B: { /* '0X8B' */
            return false
		}
		case 0X8C: { /* '0X8C' */
            return false
		}
		case 0X8D: { /* '0X8D' */
            return false
		}
		case 0X8E: { /* '0X8E' */
            return false
		}
		case 0X8F: { /* '0X8F' */
            return false
		}
		case 0X90: { /* '0X90' */
            return false
		}
		case 0X91: { /* '0X91' */
            return false
		}
		case 0X92: { /* '0X92' */
            return false
		}
		case 0X93: { /* '0X93' */
            return false
		}
		case 0X94: { /* '0X94' */
            return false
		}
		case 0X95: { /* '0X95' */
            return false
		}
		case 0X96: { /* '0X96' */
            return false
		}
		case 0X97: { /* '0X97' */
            return false
		}
		case 0X98: { /* '0X98' */
            return false
		}
		case 0X99: { /* '0X99' */
            return false
		}
		case 0X9A: { /* '0X9A' */
            return false
		}
		case 0X9B: { /* '0X9B' */
            return false
		}
		case 0X9C: { /* '0X9C' */
            return false
		}
		case 0X9D: { /* '0X9D' */
            return false
		}
		case 0X9E: { /* '0X9E' */
            return false
		}
		case 0X9F: { /* '0X9F' */
            return false
		}
		case 0XA0: { /* '0XA0' */
            return false
		}
		case 0XA1: { /* '0XA1' */
            return false
		}
		case 0XA2: { /* '0XA2' */
            return false
		}
		case 0x10: { /* '0x10' */
            return false
		}
		case 0x11: { /* '0x11' */
            return false
		}
		default: {
			return false
		}
	}
}

func MQTT_ReasonCodeFirstEnumForFieldAuthReason(value bool) (MQTT_ReasonCode, error) {
	for _, sizeValue := range MQTT_ReasonCodeValues {
		if sizeValue.AuthReason() == value {
			return sizeValue, nil
		}
	}
	return 0, errors.Errorf("enum for %v describing AuthReason not found", value)
}

func (e MQTT_ReasonCode) PubrelPubcompResponse() bool {
	switch e  {
		case 0X00: { /* '0X00' */
            return true
		}
		case 0X01: { /* '0X01' */
            return false
		}
		case 0X02: { /* '0X02' */
            return false
		}
		case 0X04: { /* '0X04' */
            return false
		}
		case 0X18: { /* '0X18' */
            return false
		}
		case 0X19: { /* '0X19' */
            return false
		}
		case 0X80: { /* '0X80' */
            return false
		}
		case 0X81: { /* '0X81' */
            return false
		}
		case 0X82: { /* '0X82' */
            return false
		}
		case 0X83: { /* '0X83' */
            return false
		}
		case 0X84: { /* '0X84' */
            return false
		}
		case 0X85: { /* '0X85' */
            return false
		}
		case 0X86: { /* '0X86' */
            return false
		}
		case 0X87: { /* '0X87' */
            return false
		}
		case 0X88: { /* '0X88' */
            return false
		}
		case 0X89: { /* '0X89' */
            return false
		}
		case 0X8A: { /* '0X8A' */
            return false
		}
		case 0X8B: { /* '0X8B' */
            return false
		}
		case 0X8C: { /* '0X8C' */
            return false
		}
		case 0X8D: { /* '0X8D' */
            return false
		}
		case 0X8E: { /* '0X8E' */
            return false
		}
		case 0X8F: { /* '0X8F' */
            return false
		}
		case 0X90: { /* '0X90' */
            return false
		}
		case 0X91: { /* '0X91' */
            return false
		}
		case 0X92: { /* '0X92' */
            return true
		}
		case 0X93: { /* '0X93' */
            return false
		}
		case 0X94: { /* '0X94' */
            return false
		}
		case 0X95: { /* '0X95' */
            return false
		}
		case 0X96: { /* '0X96' */
            return false
		}
		case 0X97: { /* '0X97' */
            return false
		}
		case 0X98: { /* '0X98' */
            return false
		}
		case 0X99: { /* '0X99' */
            return false
		}
		case 0X9A: { /* '0X9A' */
            return false
		}
		case 0X9B: { /* '0X9B' */
            return false
		}
		case 0X9C: { /* '0X9C' */
            return false
		}
		case 0X9D: { /* '0X9D' */
            return false
		}
		case 0X9E: { /* '0X9E' */
            return false
		}
		case 0X9F: { /* '0X9F' */
            return false
		}
		case 0XA0: { /* '0XA0' */
            return false
		}
		case 0XA1: { /* '0XA1' */
            return false
		}
		case 0XA2: { /* '0XA2' */
            return false
		}
		case 0x10: { /* '0x10' */
            return false
		}
		case 0x11: { /* '0x11' */
            return false
		}
		default: {
			return false
		}
	}
}

func MQTT_ReasonCodeFirstEnumForFieldPubrelPubcompResponse(value bool) (MQTT_ReasonCode, error) {
	for _, sizeValue := range MQTT_ReasonCodeValues {
		if sizeValue.PubrelPubcompResponse() == value {
			return sizeValue, nil
		}
	}
	return 0, errors.Errorf("enum for %v describing PubrelPubcompResponse not found", value)
}

func (e MQTT_ReasonCode) PubackPubrecResponse() bool {
	switch e  {
		case 0X00: { /* '0X00' */
            return true
		}
		case 0X01: { /* '0X01' */
            return false
		}
		case 0X02: { /* '0X02' */
            return false
		}
		case 0X04: { /* '0X04' */
            return false
		}
		case 0X18: { /* '0X18' */
            return false
		}
		case 0X19: { /* '0X19' */
            return false
		}
		case 0X80: { /* '0X80' */
            return true
		}
		case 0X81: { /* '0X81' */
            return false
		}
		case 0X82: { /* '0X82' */
            return false
		}
		case 0X83: { /* '0X83' */
            return true
		}
		case 0X84: { /* '0X84' */
            return false
		}
		case 0X85: { /* '0X85' */
            return false
		}
		case 0X86: { /* '0X86' */
            return false
		}
		case 0X87: { /* '0X87' */
            return true
		}
		case 0X88: { /* '0X88' */
            return false
		}
		case 0X89: { /* '0X89' */
            return false
		}
		case 0X8A: { /* '0X8A' */
            return false
		}
		case 0X8B: { /* '0X8B' */
            return false
		}
		case 0X8C: { /* '0X8C' */
            return false
		}
		case 0X8D: { /* '0X8D' */
            return false
		}
		case 0X8E: { /* '0X8E' */
            return false
		}
		case 0X8F: { /* '0X8F' */
            return false
		}
		case 0X90: { /* '0X90' */
            return true
		}
		case 0X91: { /* '0X91' */
            return true
		}
		case 0X92: { /* '0X92' */
            return false
		}
		case 0X93: { /* '0X93' */
            return false
		}
		case 0X94: { /* '0X94' */
            return false
		}
		case 0X95: { /* '0X95' */
            return false
		}
		case 0X96: { /* '0X96' */
            return false
		}
		case 0X97: { /* '0X97' */
            return true
		}
		case 0X98: { /* '0X98' */
            return false
		}
		case 0X99: { /* '0X99' */
            return true
		}
		case 0X9A: { /* '0X9A' */
            return false
		}
		case 0X9B: { /* '0X9B' */
            return false
		}
		case 0X9C: { /* '0X9C' */
            return false
		}
		case 0X9D: { /* '0X9D' */
            return false
		}
		case 0X9E: { /* '0X9E' */
            return false
		}
		case 0X9F: { /* '0X9F' */
            return false
		}
		case 0XA0: { /* '0XA0' */
            return false
		}
		case 0XA1: { /* '0XA1' */
            return false
		}
		case 0XA2: { /* '0XA2' */
            return false
		}
		case 0x10: { /* '0x10' */
            return true
		}
		case 0x11: { /* '0x11' */
            return false
		}
		default: {
			return false
		}
	}
}

func MQTT_ReasonCodeFirstEnumForFieldPubackPubrecResponse(value bool) (MQTT_ReasonCode, error) {
	for _, sizeValue := range MQTT_ReasonCodeValues {
		if sizeValue.PubackPubrecResponse() == value {
			return sizeValue, nil
		}
	}
	return 0, errors.Errorf("enum for %v describing PubackPubrecResponse not found", value)
}

func (e MQTT_ReasonCode) ConnackResponse() bool {
	switch e  {
		case 0X00: { /* '0X00' */
            return true
		}
		case 0X01: { /* '0X01' */
            return false
		}
		case 0X02: { /* '0X02' */
            return false
		}
		case 0X04: { /* '0X04' */
            return false
		}
		case 0X18: { /* '0X18' */
            return false
		}
		case 0X19: { /* '0X19' */
            return false
		}
		case 0X80: { /* '0X80' */
            return true
		}
		case 0X81: { /* '0X81' */
            return true
		}
		case 0X82: { /* '0X82' */
            return true
		}
		case 0X83: { /* '0X83' */
            return true
		}
		case 0X84: { /* '0X84' */
            return true
		}
		case 0X85: { /* '0X85' */
            return true
		}
		case 0X86: { /* '0X86' */
            return true
		}
		case 0X87: { /* '0X87' */
            return true
		}
		case 0X88: { /* '0X88' */
            return true
		}
		case 0X89: { /* '0X89' */
            return true
		}
		case 0X8A: { /* '0X8A' */
            return true
		}
		case 0X8B: { /* '0X8B' */
            return false
		}
		case 0X8C: { /* '0X8C' */
            return true
		}
		case 0X8D: { /* '0X8D' */
            return false
		}
		case 0X8E: { /* '0X8E' */
            return false
		}
		case 0X8F: { /* '0X8F' */
            return false
		}
		case 0X90: { /* '0X90' */
            return true
		}
		case 0X91: { /* '0X91' */
            return false
		}
		case 0X92: { /* '0X92' */
            return false
		}
		case 0X93: { /* '0X93' */
            return false
		}
		case 0X94: { /* '0X94' */
            return false
		}
		case 0X95: { /* '0X95' */
            return true
		}
		case 0X96: { /* '0X96' */
            return false
		}
		case 0X97: { /* '0X97' */
            return true
		}
		case 0X98: { /* '0X98' */
            return false
		}
		case 0X99: { /* '0X99' */
            return true
		}
		case 0X9A: { /* '0X9A' */
            return true
		}
		case 0X9B: { /* '0X9B' */
            return true
		}
		case 0X9C: { /* '0X9C' */
            return true
		}
		case 0X9D: { /* '0X9D' */
            return true
		}
		case 0X9E: { /* '0X9E' */
            return false
		}
		case 0X9F: { /* '0X9F' */
            return true
		}
		case 0XA0: { /* '0XA0' */
            return false
		}
		case 0XA1: { /* '0XA1' */
            return false
		}
		case 0XA2: { /* '0XA2' */
            return false
		}
		case 0x10: { /* '0x10' */
            return false
		}
		case 0x11: { /* '0x11' */
            return false
		}
		default: {
			return false
		}
	}
}

func MQTT_ReasonCodeFirstEnumForFieldConnackResponse(value bool) (MQTT_ReasonCode, error) {
	for _, sizeValue := range MQTT_ReasonCodeValues {
		if sizeValue.ConnackResponse() == value {
			return sizeValue, nil
		}
	}
	return 0, errors.Errorf("enum for %v describing ConnackResponse not found", value)
}
func MQTT_ReasonCodeByValue(value uint8) MQTT_ReasonCode {
	switch value {
		case 0X00:
			return MQTT_ReasonCode_SUCCESS
		case 0X01:
			return MQTT_ReasonCode_GRANTED_QOS_1
		case 0X02:
			return MQTT_ReasonCode_GRANTED_QOS_2
		case 0X04:
			return MQTT_ReasonCode_DISCONNECT_WITH_WILL_MESSAGE
		case 0X18:
			return MQTT_ReasonCode_CONTINUE_AUTHENTICATION
		case 0X19:
			return MQTT_ReasonCode_RE_AUTHENTICATE
		case 0X80:
			return MQTT_ReasonCode_UNSPECIFIED_ERROR
		case 0X81:
			return MQTT_ReasonCode_MALFORMED_PACKET
		case 0X82:
			return MQTT_ReasonCode_PROTOCOL_ERROR
		case 0X83:
			return MQTT_ReasonCode_IMPLEMENTATION_SPECIFIC_ERROR
		case 0X84:
			return MQTT_ReasonCode_UNSUPPORTED_PROTOCOL_VERSION
		case 0X85:
			return MQTT_ReasonCode_CLIENT_IDENTIFIER_NOT_VALID
		case 0X86:
			return MQTT_ReasonCode_BAD_USER_NAME_OR_PASSWORD
		case 0X87:
			return MQTT_ReasonCode_NOT_AUTHORIZED
		case 0X88:
			return MQTT_ReasonCode_SERVER_UNAVAILABLE
		case 0X89:
			return MQTT_ReasonCode_SERVER_BUSY
		case 0X8A:
			return MQTT_ReasonCode_BANNED
		case 0X8B:
			return MQTT_ReasonCode_SERVER_SHUTTING_DOWN
		case 0X8C:
			return MQTT_ReasonCode_BAD_AUTHENTICATION_METHOD
		case 0X8D:
			return MQTT_ReasonCode_KEEP_ALIVE_TIMEOUT
		case 0X8E:
			return MQTT_ReasonCode_SESSION_TAKEN_OVER
		case 0X8F:
			return MQTT_ReasonCode_TOPIC_FILTER_INVALID
		case 0X90:
			return MQTT_ReasonCode_TOPIC_NAME_INVALID
		case 0X91:
			return MQTT_ReasonCode_PACKET_IDENTIFIER_IN_USE
		case 0X92:
			return MQTT_ReasonCode_PACKET_IDENTIFIER_NOT_FOUND
		case 0X93:
			return MQTT_ReasonCode_RECEIVE_MAXIMUM_EXCEEDED
		case 0X94:
			return MQTT_ReasonCode_TOPIC_ALIAS_INVALID
		case 0X95:
			return MQTT_ReasonCode_PACKET_TOO_LARGE
		case 0X96:
			return MQTT_ReasonCode_MESSAGE_RATE_TO_HIGH
		case 0X97:
			return MQTT_ReasonCode_QUOTA_EXCEEDED
		case 0X98:
			return MQTT_ReasonCode_ADMINISTRATIVE_ACTION
		case 0X99:
			return MQTT_ReasonCode_PAYLOAD_FORMAT_INVALID
		case 0X9A:
			return MQTT_ReasonCode_RETAIN_NOT_SUPPORTED
		case 0X9B:
			return MQTT_ReasonCode_QOS_NOT_SUPPORTED
		case 0X9C:
			return MQTT_ReasonCode_USE_ANOTHER_SERVER
		case 0X9D:
			return MQTT_ReasonCode_SERVER_MOVED
		case 0X9E:
			return MQTT_ReasonCode_SHARED_SUBSCRIPTIONS_NOT_SUPPORTED
		case 0X9F:
			return MQTT_ReasonCode_CONNECTION_RATE_EXCEEDED
		case 0XA0:
			return MQTT_ReasonCode_MAXIMUM_CONNECT_TIME
		case 0XA1:
			return MQTT_ReasonCode_SUBSCRIPTION_IDENTIFIERS_NOT_SUPPORTED
		case 0XA2:
			return MQTT_ReasonCode_WILDCARD_SUBSCRIPTIONS_NOT_SUPPORTED
		case 0x10:
			return MQTT_ReasonCode_NO_MATCHING_SUBSCRIBERS
		case 0x11:
			return MQTT_ReasonCode_NO_SUBSCRIPTION_EXISTED
	}
	return 0
}

func MQTT_ReasonCodeByName(value string) MQTT_ReasonCode {
	switch value {
	case "SUCCESS":
		return MQTT_ReasonCode_SUCCESS
	case "GRANTED_QOS_1":
		return MQTT_ReasonCode_GRANTED_QOS_1
	case "GRANTED_QOS_2":
		return MQTT_ReasonCode_GRANTED_QOS_2
	case "DISCONNECT_WITH_WILL_MESSAGE":
		return MQTT_ReasonCode_DISCONNECT_WITH_WILL_MESSAGE
	case "CONTINUE_AUTHENTICATION":
		return MQTT_ReasonCode_CONTINUE_AUTHENTICATION
	case "RE_AUTHENTICATE":
		return MQTT_ReasonCode_RE_AUTHENTICATE
	case "UNSPECIFIED_ERROR":
		return MQTT_ReasonCode_UNSPECIFIED_ERROR
	case "MALFORMED_PACKET":
		return MQTT_ReasonCode_MALFORMED_PACKET
	case "PROTOCOL_ERROR":
		return MQTT_ReasonCode_PROTOCOL_ERROR
	case "IMPLEMENTATION_SPECIFIC_ERROR":
		return MQTT_ReasonCode_IMPLEMENTATION_SPECIFIC_ERROR
	case "UNSUPPORTED_PROTOCOL_VERSION":
		return MQTT_ReasonCode_UNSUPPORTED_PROTOCOL_VERSION
	case "CLIENT_IDENTIFIER_NOT_VALID":
		return MQTT_ReasonCode_CLIENT_IDENTIFIER_NOT_VALID
	case "BAD_USER_NAME_OR_PASSWORD":
		return MQTT_ReasonCode_BAD_USER_NAME_OR_PASSWORD
	case "NOT_AUTHORIZED":
		return MQTT_ReasonCode_NOT_AUTHORIZED
	case "SERVER_UNAVAILABLE":
		return MQTT_ReasonCode_SERVER_UNAVAILABLE
	case "SERVER_BUSY":
		return MQTT_ReasonCode_SERVER_BUSY
	case "BANNED":
		return MQTT_ReasonCode_BANNED
	case "SERVER_SHUTTING_DOWN":
		return MQTT_ReasonCode_SERVER_SHUTTING_DOWN
	case "BAD_AUTHENTICATION_METHOD":
		return MQTT_ReasonCode_BAD_AUTHENTICATION_METHOD
	case "KEEP_ALIVE_TIMEOUT":
		return MQTT_ReasonCode_KEEP_ALIVE_TIMEOUT
	case "SESSION_TAKEN_OVER":
		return MQTT_ReasonCode_SESSION_TAKEN_OVER
	case "TOPIC_FILTER_INVALID":
		return MQTT_ReasonCode_TOPIC_FILTER_INVALID
	case "TOPIC_NAME_INVALID":
		return MQTT_ReasonCode_TOPIC_NAME_INVALID
	case "PACKET_IDENTIFIER_IN_USE":
		return MQTT_ReasonCode_PACKET_IDENTIFIER_IN_USE
	case "PACKET_IDENTIFIER_NOT_FOUND":
		return MQTT_ReasonCode_PACKET_IDENTIFIER_NOT_FOUND
	case "RECEIVE_MAXIMUM_EXCEEDED":
		return MQTT_ReasonCode_RECEIVE_MAXIMUM_EXCEEDED
	case "TOPIC_ALIAS_INVALID":
		return MQTT_ReasonCode_TOPIC_ALIAS_INVALID
	case "PACKET_TOO_LARGE":
		return MQTT_ReasonCode_PACKET_TOO_LARGE
	case "MESSAGE_RATE_TO_HIGH":
		return MQTT_ReasonCode_MESSAGE_RATE_TO_HIGH
	case "QUOTA_EXCEEDED":
		return MQTT_ReasonCode_QUOTA_EXCEEDED
	case "ADMINISTRATIVE_ACTION":
		return MQTT_ReasonCode_ADMINISTRATIVE_ACTION
	case "PAYLOAD_FORMAT_INVALID":
		return MQTT_ReasonCode_PAYLOAD_FORMAT_INVALID
	case "RETAIN_NOT_SUPPORTED":
		return MQTT_ReasonCode_RETAIN_NOT_SUPPORTED
	case "QOS_NOT_SUPPORTED":
		return MQTT_ReasonCode_QOS_NOT_SUPPORTED
	case "USE_ANOTHER_SERVER":
		return MQTT_ReasonCode_USE_ANOTHER_SERVER
	case "SERVER_MOVED":
		return MQTT_ReasonCode_SERVER_MOVED
	case "SHARED_SUBSCRIPTIONS_NOT_SUPPORTED":
		return MQTT_ReasonCode_SHARED_SUBSCRIPTIONS_NOT_SUPPORTED
	case "CONNECTION_RATE_EXCEEDED":
		return MQTT_ReasonCode_CONNECTION_RATE_EXCEEDED
	case "MAXIMUM_CONNECT_TIME":
		return MQTT_ReasonCode_MAXIMUM_CONNECT_TIME
	case "SUBSCRIPTION_IDENTIFIERS_NOT_SUPPORTED":
		return MQTT_ReasonCode_SUBSCRIPTION_IDENTIFIERS_NOT_SUPPORTED
	case "WILDCARD_SUBSCRIPTIONS_NOT_SUPPORTED":
		return MQTT_ReasonCode_WILDCARD_SUBSCRIPTIONS_NOT_SUPPORTED
	case "NO_MATCHING_SUBSCRIBERS":
		return MQTT_ReasonCode_NO_MATCHING_SUBSCRIBERS
	case "NO_SUBSCRIPTION_EXISTED":
		return MQTT_ReasonCode_NO_SUBSCRIPTION_EXISTED
	}
	return 0
}

func CastMQTT_ReasonCode(structType interface{}) MQTT_ReasonCode {
	castFunc := func(typ interface{}) MQTT_ReasonCode {
		if sMQTT_ReasonCode, ok := typ.(MQTT_ReasonCode); ok {
			return sMQTT_ReasonCode
		}
		return 0
	}
	return castFunc(structType)
}

func (m MQTT_ReasonCode) LengthInBits() uint16 {
	return 8
}

func (m MQTT_ReasonCode) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func MQTT_ReasonCodeParse(readBuffer utils.ReadBuffer) (MQTT_ReasonCode, error) {
	val, err := readBuffer.ReadUint8("MQTT_ReasonCode", 8)
	if err != nil {
		return 0, nil
	}
	return MQTT_ReasonCodeByValue(val), nil
}

func (e MQTT_ReasonCode) Serialize(writeBuffer utils.WriteBuffer) error {
	return writeBuffer.WriteUint8("MQTT_ReasonCode", 8, uint8(e), utils.WithAdditionalStringRepresentation(e.name()))
}

func (e MQTT_ReasonCode) name() string {
	switch e {
	case MQTT_ReasonCode_SUCCESS:
		return "SUCCESS"
	case MQTT_ReasonCode_GRANTED_QOS_1:
		return "GRANTED_QOS_1"
	case MQTT_ReasonCode_GRANTED_QOS_2:
		return "GRANTED_QOS_2"
	case MQTT_ReasonCode_DISCONNECT_WITH_WILL_MESSAGE:
		return "DISCONNECT_WITH_WILL_MESSAGE"
	case MQTT_ReasonCode_CONTINUE_AUTHENTICATION:
		return "CONTINUE_AUTHENTICATION"
	case MQTT_ReasonCode_RE_AUTHENTICATE:
		return "RE_AUTHENTICATE"
	case MQTT_ReasonCode_UNSPECIFIED_ERROR:
		return "UNSPECIFIED_ERROR"
	case MQTT_ReasonCode_MALFORMED_PACKET:
		return "MALFORMED_PACKET"
	case MQTT_ReasonCode_PROTOCOL_ERROR:
		return "PROTOCOL_ERROR"
	case MQTT_ReasonCode_IMPLEMENTATION_SPECIFIC_ERROR:
		return "IMPLEMENTATION_SPECIFIC_ERROR"
	case MQTT_ReasonCode_UNSUPPORTED_PROTOCOL_VERSION:
		return "UNSUPPORTED_PROTOCOL_VERSION"
	case MQTT_ReasonCode_CLIENT_IDENTIFIER_NOT_VALID:
		return "CLIENT_IDENTIFIER_NOT_VALID"
	case MQTT_ReasonCode_BAD_USER_NAME_OR_PASSWORD:
		return "BAD_USER_NAME_OR_PASSWORD"
	case MQTT_ReasonCode_NOT_AUTHORIZED:
		return "NOT_AUTHORIZED"
	case MQTT_ReasonCode_SERVER_UNAVAILABLE:
		return "SERVER_UNAVAILABLE"
	case MQTT_ReasonCode_SERVER_BUSY:
		return "SERVER_BUSY"
	case MQTT_ReasonCode_BANNED:
		return "BANNED"
	case MQTT_ReasonCode_SERVER_SHUTTING_DOWN:
		return "SERVER_SHUTTING_DOWN"
	case MQTT_ReasonCode_BAD_AUTHENTICATION_METHOD:
		return "BAD_AUTHENTICATION_METHOD"
	case MQTT_ReasonCode_KEEP_ALIVE_TIMEOUT:
		return "KEEP_ALIVE_TIMEOUT"
	case MQTT_ReasonCode_SESSION_TAKEN_OVER:
		return "SESSION_TAKEN_OVER"
	case MQTT_ReasonCode_TOPIC_FILTER_INVALID:
		return "TOPIC_FILTER_INVALID"
	case MQTT_ReasonCode_TOPIC_NAME_INVALID:
		return "TOPIC_NAME_INVALID"
	case MQTT_ReasonCode_PACKET_IDENTIFIER_IN_USE:
		return "PACKET_IDENTIFIER_IN_USE"
	case MQTT_ReasonCode_PACKET_IDENTIFIER_NOT_FOUND:
		return "PACKET_IDENTIFIER_NOT_FOUND"
	case MQTT_ReasonCode_RECEIVE_MAXIMUM_EXCEEDED:
		return "RECEIVE_MAXIMUM_EXCEEDED"
	case MQTT_ReasonCode_TOPIC_ALIAS_INVALID:
		return "TOPIC_ALIAS_INVALID"
	case MQTT_ReasonCode_PACKET_TOO_LARGE:
		return "PACKET_TOO_LARGE"
	case MQTT_ReasonCode_MESSAGE_RATE_TO_HIGH:
		return "MESSAGE_RATE_TO_HIGH"
	case MQTT_ReasonCode_QUOTA_EXCEEDED:
		return "QUOTA_EXCEEDED"
	case MQTT_ReasonCode_ADMINISTRATIVE_ACTION:
		return "ADMINISTRATIVE_ACTION"
	case MQTT_ReasonCode_PAYLOAD_FORMAT_INVALID:
		return "PAYLOAD_FORMAT_INVALID"
	case MQTT_ReasonCode_RETAIN_NOT_SUPPORTED:
		return "RETAIN_NOT_SUPPORTED"
	case MQTT_ReasonCode_QOS_NOT_SUPPORTED:
		return "QOS_NOT_SUPPORTED"
	case MQTT_ReasonCode_USE_ANOTHER_SERVER:
		return "USE_ANOTHER_SERVER"
	case MQTT_ReasonCode_SERVER_MOVED:
		return "SERVER_MOVED"
	case MQTT_ReasonCode_SHARED_SUBSCRIPTIONS_NOT_SUPPORTED:
		return "SHARED_SUBSCRIPTIONS_NOT_SUPPORTED"
	case MQTT_ReasonCode_CONNECTION_RATE_EXCEEDED:
		return "CONNECTION_RATE_EXCEEDED"
	case MQTT_ReasonCode_MAXIMUM_CONNECT_TIME:
		return "MAXIMUM_CONNECT_TIME"
	case MQTT_ReasonCode_SUBSCRIPTION_IDENTIFIERS_NOT_SUPPORTED:
		return "SUBSCRIPTION_IDENTIFIERS_NOT_SUPPORTED"
	case MQTT_ReasonCode_WILDCARD_SUBSCRIPTIONS_NOT_SUPPORTED:
		return "WILDCARD_SUBSCRIPTIONS_NOT_SUPPORTED"
	case MQTT_ReasonCode_NO_MATCHING_SUBSCRIBERS:
		return "NO_MATCHING_SUBSCRIBERS"
	case MQTT_ReasonCode_NO_SUBSCRIPTION_EXISTED:
		return "NO_SUBSCRIPTION_EXISTED"
	}
	return ""
}

func (e MQTT_ReasonCode) String() string {
	return e.name()
}
