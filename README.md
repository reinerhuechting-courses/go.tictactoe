# Tic Tac Toe

Hier wird das Spiel Tic Tac Toe in einer einfachen Variante auf der Konsole implementiert.

## Aufbau

Das eigentliche Spiel steht in der Datei `game.go`. Hier gibt es eine `main()`-Funktion, die das Spiel startet und steuert.
In den Dateien `boardlogic.go` und `gamelogic.go` finden sich Funktionen, mit denen der Zustand des Spielfelds und des Spiels
abgefragt werden können. In `boardio.go` und `gameio.go` gibt es Funktionen, die der Benutzerinteraktion dienen.
Diese Funktionen werden in der `main()`-funktion verwendet, um daraus das Spiel zu bauen.

Die Funktionen sind jeweils in eine Variante für das Spiel und das Spielfeld aufgeteilt, da sich die Spielfeld-Funktionen
nicht nur für dieses Spiel eignen, sondern für verschiedene Spiele wiederverwendbar sind.
