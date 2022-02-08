# FIM - File Integrity Monitoring
Für einen Test der Applikation wird die Verwendung des Betriebssystems [Ubuntu 21.10](https://releases.ubuntu.com/21.10/) empfohlen.


## Erstellung von TLS-Zertifikaten

Für den Betrieb der Teilapplikationen werden mehrere Zertifikate benötigt. Im Ordner ***tls*** sind einige Zertifikate für die in der Thesis gezeigte Demonstration bereits mitgeliefert. In den folgenden Unterabschnitten wird die Erstellung neuer Zertifikate für Server und Endparteien erklärt. Es sollte immer ein Client-Zertifikat mit dem Common Name ***admin*** erstellt werden, da dieser Benutzer beim Ausführen des Server-Setups standardmäßig erstellt wird. Dieser Benutzer erhält die Rollen viewer, approver und user_admin und kann somit weitere Nutzer anlegen.

### Erstellen eines Stammzertifikats

Mit dem folgenden Befehl wird ein privater Schlüssel auf der elliptischen Kurve secp384r1 mit 384-Bit Schlüssellänge für das Stammzertifikat generiert. 

```
$ openssl ecparam -out ca.key -name secp384r1 -genkey
```

Im Anschluss wird ein Certificate Signing Request erstellt. Dieser wird zur Erstellung des Zertifikats benötigt. Die bei der Ausführung abgefragten Daten können beliebig gewählt werden. Das Challenge Passwort sollte leer gelassen werden.  

```
$ openssl req -new -sha384 -key ca.key -out ca.csr
```

Im letzten Schritt wird aus den vorher erstellten Bestandteilen ein selbst-signiertes Stammzertifikat gebildet. Das fertige Zertifikat befindet sich in der Datei ca.pem und der zugehörige Schlüssel in der Datei ca.key

```
$ openssl x509 -req -sha384 -days 365 -in ca.csr \
 -signkey ca.key -out ca.pem
```

### Erstellen eines selbst-signierten Server-Zertifikats

Zu Beginn wird ein privater Schlüsssel auf der Kurve secp384r1 erzeugt.


```
$ openssl ecparam -out server.key -name secp384r1 -genkey
```

gRPC setzt für den Server die Verwendung eines SAN-Zertifikats voraus. Ein Subject-Alternative-Name-Zertifikat ermöglicht die Nutzung eines einzigen Zertifikats für mehrere IP und DNS-Adressen. Für die Erstellung eines solchen Zertifikats benötigt openssl eine Konfigurationsdatei in der alle alternativen Namen aufgeführt sind. Mit einem Texteditor muss die Datei ***server-cert.cnf*** mit dem folgenden Inhalt angelegt werden.

```
[req]
distinguished_name = req_distinguished_name
req_extensions = req_ext
prompt = no

[req_distinguished_name]
C   = DE
ST  = Schleswig-Holstein
L   = Flensburg
O   = Hochschule Flensburg
OU  = Informatik
CN  = server.local

[req_ext]
subjectAltName = @alt_names

[alt_names]
IP.1 = 10.0.0.50
IP.2 = 127.0.0.1
DNS.1 = localhost
```

Unter dem Abschnitt ***req_distinguished_name*** können die Informationen zur besitzenden Organisation, sowie der Common Name, angepasst werden. Unter ***alt_names*** müssen alle alternativen IP's und DNS-Namen gelistet werden. Zur Ergänzung eines weiteren Eintrags muss lediglich die Zahl hinter der IP beziehungsweise DNS inkrementiert werden. Ein gültige Ergänzung wäre beispielsweise ***DNS.2 = myserver.de***.

Nun muss ein Certificate Signing Request mit der angelegten Konfigurationsdatei gebildet werden. Im Gegensatz zur Erstellung des Stammzertifikats, werden hier keine weiteren Informationen abgefragt, da diese bereits durch die Konfiguration festgelegt wurden.

```
$ openssl req -new -sha384 -key server.key -out server.csr \ 
-config server-cert.cnf
```

Im letzten Schritt wird das vom Stammzertifikat signierte SAN-Zertifikat erzeugt.

```
$ openssl x509 -req -sha384 -days 365 -in server.csr \ 
-CA ca.pem -CAkey ca.key -CAcreateserial -out server.pem \ 
-extensions req_ext -extfile server-cert.cnf
```

### Erstellen eines Zertifikats für einen Agenten oder Client

Ein privater Schlüssel für eine Endpartei wird erzeugt.

```
$ openssl ecparam -out endpoint.key -name secp384r1 -genkey
```

Wie bei den vorangegangen Zertifikaten wird im interaktiven Modus ein Certificate Signing Request gebildet. Der Common Name muss hier zwingend dem Namen der zu erstellenden Endpartei entsprechen. Soll also beispielsweise ein Client mit Ansichtsrechten für Alerts erstellt werden, so könnte der Common Name viewer_alice lauten.

```
$ openssl req -new -sha384 -key endpoint.key -out endpoint.csr
```

Zum Abschluss wird das Zertifikat erstellt und durch das Stammzertifikat signiert.

```
$ openssl x509 -req -sha384 -days 365 -in endpoint.csr \ 
-CA ca.pem -CAkey ca.key -CAcreateserial -out endpoint.pem
```

## Kompilierung und Vorbereitung

Für die einzelnen Bestandteile befindet sich jeweils eine Konfigurationsdatei ***config.yaml*** im entsprechenden Ordner. Die Konfiguration ist standardmäßig für einen Betrieb auf localhost ausgerichtet. Im Bedarfsfall können hier Änderungen vorgenommen werden. Das Feld ***host*** und ***port*** bezieht sich in den Konfigurationsdateien immer auf die Adresse des Servers.

Zum einfachen Kompilieren und Aufsetzen befindet sich das Skript ***prepare.sh*** im Repository. Es kompiliert alle Teilapplikationen, installiert und erstellt eine Datenbank und führt das Setup des Servers aus.

```
$ chmod u+x prepare.sh
$ sudo ./prepare.sh
```

## Starten des Servers

Nach der Ausführung des Skriptes, kann der Server gestartet werden. Standardmäßig versucht der Server die Konfigurationsdatei im Working Directory einzulesen. Dieses Verhalten kann durch die manuelle Angabe eines Pfades überschrieben werden, sollte dies nötig sein.

```
./fimserver --config /path/to/config.yaml
```

## Starten des Agents

Standardmäßig versucht der Agent die Konfigurationsdatei im Working Directory einzulesen. Dieses Verhalten kann durch die manuelle Angabe eines Pfades überschrieben werden, sollte dies nötig sein.

```
sudo ./fimagent --config /path/to/config.yaml
```

## Benutzung des Clients

Der Client bietet eine Vielzahl von Funktionen, die in den folgenden Unterabschnitten erläutert und exemplarisch vorgestellt werden.

### Show-Agents

Der Befehl show-agents kann dazu verwendet werden, alle beim Server registrierten Agenten mit Informationen zu den überwachten Pfaden anzuzeigen.

___Beispiel___

Abfragen aller registrierten Agenten:

```
./fimclient show-agents --config /path/to/config.yaml
```

### Show-Alerts

Der Befehl show-alerts wird dazu verwendet, alle seit dem letzten Aktualisieren der Baseline erstellten Alerts anzuzeigen.

___Beispiel___

Abfragen aller Alerts für den Agenten mit dem Namen agent1

```
./fimclient show-alerts --name agent1 --config /path/to/config.yaml
```

### Create-Agents

Mittels des Befehls create-agents ist es möglich mehrere neue Agenten mit einem einzigen Befehl zu erstellen. Dafür muss dem Client lediglich eine YAML-Datei übergeben werden, die die zu erstellenden Agenten beschreibt.

___Beispiel___

Anlegen der Datei ***agents.yaml*** mit folgendem Inhalt:

```
- name: agent1
  watched_paths:
    - /home/agent1
    - /opt/agent1
- name: agent2
  watched_paths:
    - /home/agent2
    - /usr/bin
```

Anlegen der Agents durch den Client:

```
./fimclient create-agents --file /path/to/agents.yaml --config /path/to/config.yaml
```

### Create-Clients

Mittels des Befehl create-clients können mehrere neue Client-Benutzer erstellt werden. Dazu muss dem Client eine YAML-Datei übergeben werden, die die zu erstellenden Benutzer beschreibt.

___Beispiel___

Anlegen der Datei ***clients.yaml*** mit folgendem Inhalt:

```
- name: alice
  roles:
    - user_admin
    - approver
    - viewer
- name: bob
  roles:
    - viewer
```

Anlegen der neuen Client-Benutzer durch den Client:

```
./fimclient create-clients --file /path/to/clients.yaml --config /path/to/config.yaml
```

### Delete

Mit dem Befehl delete kann ein Agent oder ein Client aus der Datenbank des Servers gelöscht werden. Dieser Befehl löscht restlos alle vorhandenen Daten zu einem Benutzer und ist nicht reversibel.


___Beispiel___

Löschen des Agenten mit dem Namen agent1

```
./fimclient delete --name agent1 --config /path/to/config.yaml
```

### Update-Paths

Der Befehl update-paths kann verwendet werden, um die überwachten Pfade eines Agenten zu überschreiben. Die ursprünglich überwachten Pfade werden damit ungültig und durch die neu angegeben ersetzt.

Im Anschluss an diesen Befehl sollte immer ein approve-Befehl ausgeführt und der betroffene Agent neu gestartet werden, da es sonst zu falschen Alert-Meldungen kommen kann.

___Beispiel___

Anlegen der Datei agent1_paths.yaml mit folgendem Inhalt:

```
name: agent1
watched_paths:
    - /home/agent1
    - /var/lib
```

Änderung der überwachten Dateipfade an den Server übertragen:

```
./fimclient update-paths --file /path/to/agent1_paths.yaml \ 
--config /path/to/config.yaml
```

### Approve

Der Befehl approve kann verwendet werden, um einem Agenten zu erlauben, seine Baseline zu aktualisieren. Dem Agenten wird dazu temporär die Rolle updater zugewiesen, die ihn einmalig zu dieser Aktion berechtigt. Nach dem Ausführen dieses Befehls muss der betroffene Agent neu gestartet werden. 

Ein Aktualisieren der Baseline führt zur Löschung der alten Baseline sowie aller bisher protokollierten Alerts des betreffenden Agents. Die Ausführung des Befehls ist nicht reversibel

___Beispiel___

Erlaubnis zum Aktualisieren der Baseline für den Agenten agent1 erteilen:

```
$ ./fimclient approve --name agent1 --config /path/to/config.yaml
```
