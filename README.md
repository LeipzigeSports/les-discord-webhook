# LES Discord Webhook

Go application to trigger custom Discord webhooks.

> [!NOTE]
> Keep your webhook url's secret.


## Installation
Copy `config.example.json` to `config.json` and ensure it can't be read by anyone.

```bash
cp config.example.json config.json
sudo chown root:root config.json
sudo chmod 640 config.json
```

Adjust the `config.json` file accordingly using a text editor of your choice.
```json
{
    "data": {
        "Event A": {
            "cron": "30 14 * * 0",
            "webhook_url": "https://discord.com/api/webhooks/...",
            "message": "<@&123456> hola!"
        },
        "Event B": {
            "cron": "0 14 * * 0",
            "webhook_url": "https://discord.com/api/webhooks/...",
            "message": "Check out this channel: <#123456>"
        }
    }
}
```

Build the application or start it using the shipped container image.

### Docker
pull and run via cli
```bash
docker pull ghcr.io/leipzigesports/les-discord-webhook:latest
docker run -it --name les-discord-webhook -v ./config.json:/app/config.json -d les-discord-webhook:latest
```

or via `compose.yaml`  
```yaml
services:
  les-discord-webhook:
    ghcr.io/leipzigesports/les-discord-webhook:latest
    container_name: les-discord-webhook
    volumes:
      - ./config.json:/app/config.json
```

```bash
docker compose up -d && docker compose logs -f
```

### Build with go
```bash
# run locally
go run .

# build locally
go build -v -o les-discord-webhook .
```

## Setup Discord webhook
1. select channel
1. open channel settings
1. navigate to integrations
1. create webook
    - select profile picture
    - enter a 'username'
    - copy webhook url
1. create a new section/key in the `config.json` file and insert the webhook url

### Format message
you can embed <&`ROLE_ID`> to ping a role, <@`USER_ID`> to ping a user and <#`CHANNEL_ID`> to link to a channel.