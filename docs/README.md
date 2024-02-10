# Protocol Definition Documentation

## User UUIDs
User UUIDs are generated using SHA-512 hashing on the concatenation of the username and a randomly generated number. The resulting UUID is stored in the database for each user.

## Conversations
Users can send messages to each other using UUIDs. Conversations are stored as channels in the database.

## Messages
Messages are stored in the database and can be retrieved using the channel ID, sorted by time. The first message initiates the creation of a conversation, which is then stored.

To send a message, a POST request is made to `/api/chat/message` with the following data:
```json
{
    "uuid": "<UUID>",
    "data": "<message>"
}
```
Messages to clients are streamed using websockets.

## Metadata
Metadata for all users can be accessed via a GET request to `/api/chat/all`. The output is in the following format:
```json
{
    "clients": [
        {
            "user": "<username>",
            "ip": "<ip>",
            "UUID": "<uuid>"
        },
        ...
    ]
}
```

## Security
All conversations are encrypted using SSL to ensure secure communication.

