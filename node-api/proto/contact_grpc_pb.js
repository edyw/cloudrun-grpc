// GENERATED CODE -- DO NOT EDIT!

'use strict';
var grpc = require('@grpc/grpc-js');
var contact_pb = require('./contact_pb.js');

function serialize_contact_ContactReply(arg) {
  if (!(arg instanceof contact_pb.ContactReply)) {
    throw new Error('Expected argument of type contact.ContactReply');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_contact_ContactReply(buffer_arg) {
  return contact_pb.ContactReply.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_contact_ContactRequest(arg) {
  if (!(arg instanceof contact_pb.ContactRequest)) {
    throw new Error('Expected argument of type contact.ContactRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_contact_ContactRequest(buffer_arg) {
  return contact_pb.ContactRequest.deserializeBinary(new Uint8Array(buffer_arg));
}


var ContactService = exports.ContactService = {
  getContact: {
    path: '/contact.Contact/GetContact',
    requestStream: false,
    responseStream: false,
    requestType: contact_pb.ContactRequest,
    responseType: contact_pb.ContactReply,
    requestSerialize: serialize_contact_ContactRequest,
    requestDeserialize: deserialize_contact_ContactRequest,
    responseSerialize: serialize_contact_ContactReply,
    responseDeserialize: deserialize_contact_ContactReply,
  },
};

exports.ContactClient = grpc.makeGenericClientConstructor(ContactService);
