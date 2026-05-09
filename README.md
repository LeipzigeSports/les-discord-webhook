# les-discord-webhook
LES Orga Discord Webhooks

SICHERHEITSHINWEIS: Gebe niemals deinen Webhooklink preis!

Um einfache Aufgaben auf Discord zu automatisieren, können Webhooks verwendet werden.

Vorraussetzungen für das Skript:
-> curl
-> Webhook-Link von Discord

Um einen Webhook zu generieren, welcher in der Variable WEBHOOK_URL gespeichert wird, benötigst du in Discord die entsprechenden Rechte dazu. Wenn du diese hast, findest du Webhooks unter den Servereinstellungen -> Integrationen -> Webhooks
Dort kannst du dir einen neuen Webhook erstellen und dann den entsprechenden Channel einstellen wo die Nachricht gesendet werden soll. 

Trage nun die Webhookadresse in WEBHOOK_URL ein.

Das Skript ist grundsätzlich sehr anpassbar und ist eine sehr einfache Fassung. Nachfolgend findest du einige weitere Informationen wie das Skript funktioniert und wie du es erweitern kannst.

Das Skript nutzt curl um HTTP-Daten an den Webhook zu senden, welcher diese Daten dann in den konfigurierten Channel sendet.
Der Curl-Teil des Skripts sollte immer als letztes stehen und nicht verändert werden!

Um Rollen zu pingen, wird die ID der Rolle benötigt. Um diese zu erhalten, klicke Rechtsklick auf die Rolle bei einem Nutzer der sie besitzt (oder in der Rollenübersicht) und kopiere dir die Rollen-ID. 
Sollte der Button nicht angezeigt werden, musst du den Entwicklermodus von Discord aktivieren.

Um nun die Rolle zu pingen, füge folgendes in deine Nachricht ein:
<@&$ROLE_ID>

$ROLE_ID ist dabei die Variable welche auch anders heißen kann.

Um andere Channel in die Nachricht zu integrieren, schreibe folgendes:
<#$CHANNEL_ID>

Dabei ist $CHANNEL_ID wieder die Variable.

Nun da dein Skript fertiggestellt ist, musst du es noch ausführbar machen. Nutze dazu: "chmod +x (Pfad/zum/Skript)"

Um das Skript zu testen, führe es einfach einmal aus. Deine konfigurierte Nachricht sollte nun im Discord erscheinen!

Um das Skript zu automatisieren, kannst du einfach crontab benutzen. Eine Anleitung dafür findest du hier: https://www.stetic.com/developer/cronjob-linux-tutorial-und-crontab-syntax/
Croncommand Generator: https://crontab.guru/
