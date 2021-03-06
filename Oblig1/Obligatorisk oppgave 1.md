Gruppe: Yaguel van der Meij, Xohan Otero Barbosa, Vegard Marvik, Vegard
Trydal, Andreas Wöllert

                             Obligatorisk oppgave 1. IS-105
                                  Gruppe: Oldtimers
                          Xohan Otero Barbosa, Xohano@uia.no
                             Vegard Marvik, Vegarm@uia.no
                          Yaguel van der Meij, Yaguev@uia.no
                             Vegard Trydal, Vegart13@uia.no
                            Andreas Wöllert, Andrew17@uia.no

i.  https://imgur.com/1G9koAi

​a. Binære tall til desimaltall og motsatt: Konvertering fra binære til
desimaltall og omvendt kan gjøres enkelt ved bruk av en tabell. 1 i
binære tall betyr at verdien skal legges til, mens 0 betyr at verdien
ikke skal. Hvis vi blir tildelt et binært tall som for eksempel 1101 vet
vi ved å begynne fra høyre siden av tabellen at desimaltallet inneholder
1, 4 og 8 som til sammen blir 13.

Ved konvertering fra desimaltall til binære tall må vi fragmentere det
desimale tallet opp og se hva det er bygget opp av. Vi ser på tabellen
at for å bygge opp desimaltallet 13 trenger vi den største summen
tilgjengelig som er 13 eller lavere, i dette tilfellet 8. Vi har også
plass til 4, men hvis vi legger til 2 overskrider vi totalverdien.
Derfor hopper vi over 2 og legger til 1 istedet som summerer i 13.
https://imgur.com/1mQkPTV

Binære tall til heksadesimaltall og motsatt: Den enkleste måten å
konvertere mellom heksadesimale tall og binære tall er å se på de
forskjellige tallene og deres tilsvarende verdi i det andre
tallsystemet. I tredje forelesning av faget fikk vi en liten tabell
tildelt som gav oss verdiene av D, 3 og A. Hvis vi skal konvertere for
eksempel 0xA3D til binære tall blir det 011000111001.
https://i.imgur.com/tQgkCc0.png

​b. Desimaltall til heksadesimaltall og motsatt:

For å konvertere til et heksadesimalt tall fra desimaltall deler vi
tallet på 16, om vi får rest skriver vi det net, ellers skriver vi 0.
Etter det starter vi fra toppen igjen til svaret på heltall divisjonen
er 0. For eksempel tallet 435 https://i.imgur.com/dlu63K6.png Det
ferdige heksadesimale tallet blir 1b3.

Heksadesimal tall til desimaltall:

For å regne bruker vi en tabell hvor vi "oversetter" hver tall i det ved
å gange det med 16 opphøyd i posisjonen til tallet i rekka. La oss ta
for eksempel tallet BEEF https://i.imgur.com/R9HBRms.png Summen av
verdiene er altsa 15 + 224 + 358 + 45056 = 4887910 = BEEF

​ii. c. Grafisk fremstilling av benchmarks: https://imgur.com/a/D7GMR

Hva vi kan lese fra benchmark testingen er: \* Økning i mengde tall som
sorteres øker mengden nanosekund brukt delt på mengden operasjoner. \*
Økning i mengde tall som sorteres senker mengden ganger den looper
gjennom for sorteringsprosessen. \* Ns/op og loops ser ut til å ha en
direkte korrelasjon ettersom økning i en av de fører til senkning i den
andre. \* BSort tester dårligst på ns/op i samtlige gjennomganger, men
også best på loops. \* QSort tester dårligst på mengden ganger den
looper gjennom seg selv for sortering.
