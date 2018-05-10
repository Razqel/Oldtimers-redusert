# Oldtimers-redusert

 **Gruppemedlemmer:**
 
 * [Xohan Otero Barbosa](https://github.com/galirousa)
 * [Vegard Marvik](https://github.com/VMarvik)
 * [Yaguel van der Meij](https://github.com/Yaguel)
 * [Vegard Trydal](https://github.com/vegart13)
 * [Andreas Wöllert](https://github.com/Razqel)
 
 ----
 
 ## Systemspesifikasjoner
 
En nettside for utviklere som skal kunne finne relevant data om spesifikke problemer (issues) på github, dette blir funnet gjennom: https://api.github.com/search/issues api, webapplikasjonen skal finne de 30 mest relevante issues i henhold til hvor spesifikk man har vært med eget søk. Dette skal da gjøre det lett for brukeren å finne relevant info om hva man har problemer med. 
I tillegg skal utviklere som ønsker å hjelpe andre med sine problemer lett kunne finne fram et spesifikt tema som er relevant med de sin erfaring, eller en repository som de er kjent med som f.eks. golang/go.
Brukeren kan søke på et tema og for å gjør søkkravene enda spesifikt bruker man "keywords" som man kan finne på: https://developer.github.com/v3/search/#search-issues. 
e.g. repo:googlecloudplatform/golang-samples (for å spesifisere repository), is:open/closed (for å spesifisere om problemet er løst eller ikke).
Etter man har fylt inn søkekravene, får man opp de 30 mest aktuelle problemer med saksnummer som kan trykkes på og linkes rett til problemet på github.com, dette følget av tittel til selve problemet, så brukernavn med link, og informasjon om hvor mange timer og hvor mange dager siden problemet har blitt opprettet. Websiden oppdateres hver 10 minutter.

## Systemarktiktur 

TCP Server settes opp på localhost:8080/1, her får man informasjon om hvordan man skal bruke 
søkefeltet til å finne på de mest relaterte problemer. Etter at brukeren har fylt på søkekravene 
og kommer til neste stige, får man opp en liste av de 30 mest relevante problemer som blir 
hentet fra Github’s issue API. 
Siden som blir spurt etter informasjon er
https://api.github.com/search/issues?q= {query}{&page,per_page,sort,order} <-- det er her 
kravene fra søkefeltet fyller inn resten som trenges til få fram resultatene.
Nedenfor en illustrasjon som viser logikken, både ved henting og sending av data. 

![Systemarkitektur](https://github.com/Razqel/Oldtimers-redusert/blob/master/Oblig4/Github/image/Structure.png)




