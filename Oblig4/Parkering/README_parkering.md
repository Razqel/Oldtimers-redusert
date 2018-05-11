# Oldtimers-redusert

 **Gruppemedlemmer:**
 
 * [Xohan Otero Barbosa](https://github.com/galirousa)
 * [Vegard Marvik](https://github.com/VMarvik)
 * [Yaguel van der Meij](https://github.com/Yaguel)
 * [Vegard Trydal](https://github.com/vegart13)
 * [Andreas Wöllert](https://github.com/Razqel)
 
 ----
 
 ## Systemspesifikasjoner

En nettside for å sjekke parkeringsmuligheter ved Stavanger by. Informasjon brukt på siden er hentet fra https://open.stavanger.kommune.no/dataset/stavanger-parkering/resource/d1bdc6eb-9b49-4f24-89c2-ab9f5ce2acce Stavanger Kommune, som viser antall plasser igjen på 9 parkeringshus i tillegg til lokasjonen til husene, navn, og når informasjonen ble oppdatert. På siden finner man også et kart med lokasjonen til parkeringshusene vist med røde markører.


Det blir sendt en request til APIen når den blir åpnet, og hvert andre minutt etter det. Informasjonen den får tilbake er det siste oppdateringen som ble gjort i parkeringshusene, det vil si sist gang APIen selv ble oppdatert.


## Systemarkitektur
        
Konsolidering Serveren skal hente data hvert andre minutt fra Stavanger parkerings api og skrive med tekst hvor mange ledige plasser det er igjen og vise på kart gjennom googles kart api hvor de ulike parkeringsplassene er. Dette skal da gjøre det lett for bruker å kunne se og velge hvor han skal parkere. 
Nedenfor en illustrasjon som viser logikken, både ved henting og sending av data. 

![Systemarkitektur](https://github.com/Razqel/Oldtimers-redusert/blob/Oblig4/Parkering/Bilde/Structurestv.png)