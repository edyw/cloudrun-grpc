const grpc = require('@grpc/grpc-js')
const { GoogleAuth } = require('google-auth-library')
const messages = require('./proto/contact_pb')
const services = require('./proto/contact_grpc_pb')

const contactServerHost = 'go-contact-nouqicixbq-as.a.run.app'

const getContactHandler = async (req, res) => {
    // Initialize client connections outside handler in your implementation
    const client = new services.ContactClient(contactServerHost + ':443', grpc.credentials.createSsl())
    const grpcReq = new messages.ContactRequest()
    grpcReq.setPhonenumber(req.params.phone)

    client.getContact(grpcReq, (err, grpcReply) => {
        console.log('grpcReply: ' + JSON.stringify(grpcReply, ' ', 2))
        res.send(grpcReply.getName())
    })
}

const getContactAuthHandler = async (req, res) => {
    var metadata = new grpc.Metadata();
    metadata.add('authorization', 'Bearer ' + await getIdTokenFromMetadataServer())

    const client = new services.ContactClient(contactServerHost + ':443', grpc.credentials.createSsl())
    const grpcReq = new messages.ContactRequest()
    grpcReq.setPhonenumber(req.params.phone)

    client.getContact(grpcReq, metadata, (err, grpcReply) => {
        console.log('grpcReply: ' + JSON.stringify(grpcReply, ' ', 2))
        res.send(grpcReply.getName())
    })
}

const getIdTokenFromMetadataServer = async () => {
    const googleAuth = new GoogleAuth()
    const googleClient = await googleAuth.getClient()

    token = await googleClient.fetchIdToken('https://' + contactServerHost)
    console.log('token: ' + JSON.stringify(token, ' ', 2))
    return token
}

module.exports = {
    getContactHandler,
    getContactAuthHandler
}