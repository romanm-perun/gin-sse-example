# gin-sse-example

An example of Server-Sent-Events (SSE) in Gin web framework.

This simple example demonstrates how to send messages to clients via SSE. Each client can receive only messages for specific subscription topic.

## Get messages for subscription `topic A`

```bash
curl -N http://localhost:3000/subscription/topic%20A
```

Output:
```
event:topic A
data:the time is 2021-05-10 17:56:15.507041755 +0300 IDT m=+16.008224208

event:topic A
data:the time is 2021-05-10 17:56:17.507907319 +0300 IDT m=+18.009089817

^C
```

## Get messages for subscription `topic B`

```bash
curl -N http://localhost:3000/subscription/topic%20B
```

Output:
```
event:topic B
data:the UTC time is 2021-05-10 14:56:06.008682785 +0000 UTC

event:topic B
data:the UTC time is 2021-05-10 14:56:06.508988191 +0000 UTC

event:topic B
data:the UTC time is 2021-05-10 14:56:07.009543841 +0000 UTC

event:topic B
data:the UTC time is 2021-05-10 14:56:07.510074822 +0000 UTC

event:topic B
data:the UTC time is 2021-05-10 14:56:08.010339837 +0000 UTC

^C
```
