[Table of contents](README.md#table-of-contents)

# Workers

This page list all the currently available workers on the cozy-stack. It
describes their input arguments object. See the [jobs document](jobs.md) to
know more about the API context in which you can see how to use these
arguments.

## log worker

The `log` worker will just print in the log file the job sent to it. It can
useful for debugging for example.

## push worker

The `push` worker can be used to send push-notifications to a user's device.
The options are:

* `client_id`: the ID of the oauth client to push a notification to.
* `title`: the title of the notification
* `message`: the content of the notification
* `data`: key-value string map for additional metadata (optional)
* `priority`: the notification priority: `high` or `normal` (optional)
* `topic`: the topic identifier of the notification (optional)
* `sound`: a sound associated with the notification (optional)

### Example

```json
{
  "client_id": "abcdef123123123",
  "title": "My Notification",
  "message": "My notification content.",
  "priority": "high",
}
```

### Permissions

To use this worker from a client-side application, you will need to ask the
permission. It is done by adding this to the manifest:

```json
{
  "permissions": {
    "push-notification": {
      "description": "Required to send notifications",
      "type": "io.cozy.jobs",
      "verbs": ["POST"],
      "selector": "worker",
      "values": ["push"]
    }
  }
}
```

## unzip worker

The `unzip` worker can take a zip archive from the VFS, and will unzip the
files inside it to a directory of the VFS. The options are:

* `zip`: the ID of the zip file
* `destination`: the ID of the directory where the files will be unzipped.

### Example

```json
{
  "zip": "8737b5d6-51b6-11e7-9194-bf5b64b3bc9e",
  "destination": "88750a84-51b6-11e7-ba90-4f0b1cb62b7b"
}
```

### Permissions

To use this worker from a client-side application, you will need to ask the
permission. It is done by adding this to the manifest:

```json
{
  "permissions": {
    "unzip-to-a-directory": {
      "description": "Required to unzip a file inside the cozy",
      "type": "io.cozy.jobs",
      "verbs": ["POST"],
      "selector": "worker",
      "values": ["unzip"]
    }
  }
}
```

## sendmail worker

The `sendmail` worker can be used to send mail from the stack. It implies that
the stack has properly configured an access to an SMTP server. You can see an
example of configuration in the [cozy.example.yaml](../cozy.example.yaml) file
at the root of this repository.

`sendmail` options fields are the following:

* `mode`: string specifying the mode of the send:
  * `noreply` to send a notification mail to the user
  * `from` to send a mail from the user
* `to`: list of object `{name, email}` representing the addresses of the
  recipients. (should not be used in `noreply` mode)
* `subject`: string specifying the subject of the mail
* `parts`: list of part objects `{type, body}` listing representing the content
  parts of the
  * `type` string of the content type: either `text/html` or `text/plain`
  * `body` string of the actual body content of the part
* `attachments`: list of objects `{filename, content}` that represent the files
  attached to the email

### Examples

```js
// from mode sending mail from the user to a list of recipients
{
    "mode": "from",
    "to": [
        {"name": "John Doe 1", "email":"john1@doe"},
        {"name": "John Doe 2", "email":"john2@doe"}
    ],
    "subject": "Hey !",
    "parts": [
        {"type":"text/html", "body": "<h1>Hey !</h1>"},
        {"type":"text/plain", "body": "Hey !"}
    ]
}

// noreply mode, sending a notification mail to the user
{
    "mode": "noreply",
    "subject": "You've got a new file !",
    "parts": [
        {"type":"text/html", "body": "<h1>Hey !</h1>"},
        {"type":"text/plain", "body": "Hey !"}
    ]
}
```

### Permissions

To use this worker from a client-side application, you will need to ask the
permission. It is done by adding this to the manifest:

```json
{
  "permissions": {
    "mail-from-the-user": {
      "description": "Required to send mails from the user to his/her friends",
      "type": "io.cozy.jobs",
      "verbs": ["POST"],
      "selector": "worker",
      "values": ["sendmail"]
    }
  }
}
```
