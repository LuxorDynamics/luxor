package generator

const Nouns = `[
  {
    "noun": "Quergang",
    "article": "der"
  },
  {
    "noun": "Querkopf",
    "article": "der"
  },
  {
    "noun": "Querpass",
    "article": "der"
  },
  {
    "noun": "Querschnitt",
    "article": "der"
  },
  {
    "noun": "Querschuss",
    "article": "der"
  },
  {
    "noun": "Querstreifen",
    "article": "der"
  },
  {
    "noun": "Querstrich",
    "article": "der"
  },
  {
    "noun": "Quertreiber",
    "article": "der"
  },
  {
    "noun": "Querulant",
    "article": "der"
  },
  {
    "noun": "Querverweis",
    "article": "der"
  },
  {
    "noun": "Quesenbandwurm",
    "article": "der"
  },
  {
    "noun": "Quetschhahn",
    "article": "der"
  },
  {
    "noun": "Quetzal",
    "article": "der"
  },
  {
    "noun": "Quietismus",
    "article": "der"
  },
  {
    "noun": "Quietschton",
    "article": "der"
  },
  {
    "noun": "Quintal",
    "article": "der"
  },
  {
    "noun": "Quintaner",
    "article": "der"
  },
  {
    "noun": "Quirl",
    "article": "der"
  },
  {
    "noun": "Quisling",
    "article": "der"
  },
  {
    "noun": "Quittungsblock",
    "article": "der"
  },
  {
    "noun": "Quizmaster",
    "article": "der"
  },
  {
    "noun": "Quotient",
    "article": "der"
  },
  {
    "noun": "Jammerlappen",
    "article": "der"
  },
  {
    "noun": "Jangtsekiang",
    "article": "der"
  },
  {
    "noun": "Janitschar",
    "article": "der"
  },
  {
    "noun": "Januar",
    "article": "der"
  },
  {
    "noun": "Japaner",
    "article": "der"
  },
  {
    "noun": "Jargon",
    "article": "der"
  },
  {
    "noun": "Jasager",
    "article": "der"
  },
  {
    "noun": "Jasmin",
    "article": "der"
  },
  {
    "noun": "Jaspis",
    "article": "der"
  },
  {
    "noun": "Jatagan",
    "article": "der"
  },
  {
    "noun": "Jauchzer",
    "article": "der"
  },
  {
    "noun": "Javaner",
    "article": "der"
  },
  {
    "noun": "Jazz",
    "article": "der"
  },
  {
    "noun": "Jazzer",
    "article": "der"
  },
  {
    "noun": "Jazzmusiker",
    "article": "der"
  },
  {
    "noun": "Jeansanzug",
    "article": "der"
  },
  {
    "noun": "Jeansstoff",
    "article": "der"
  },
  {
    "noun": "Jeep",
    "article": "der"
  },
  {
    "noun": "Jehova",
    "article": "der"
  },
  {
    "noun": "Jemen",
    "article": "der"
  },
  {
    "noun": "Jemenit",
    "article": "der"
  },
  {
    "noun": "Jen",
    "article": "der"
  },
  {
    "noun": "Jesuit",
    "article": "der"
  },
  {
    "noun": "Jesuitenorden",
    "article": "der"
  },
  {
    "noun": "Jesus",
    "article": "der"
  },
  {
    "noun": "Jet",
    "article": "der"
  },
  {
    "noun": "Jeton",
    "article": "der"
  },
  {
    "noun": "Jetstream",
    "article": "der"
  },
  {
    "noun": "Jetztmensch",
    "article": "der"
  },
  {
    "noun": "Jigger",
    "article": "der"
  },
  {
    "noun": "Jitterbug",
    "article": "der"
  },
  {
    "noun": "Jive",
    "article": "der"
  },
  {
    "noun": "Job",
    "article": "der"
  },
  {
    "noun": "Jockei",
    "article": "der"
  },
  {
    "noun": "Jodler",
    "article": "der"
  },
  {
    "noun": "Jogalehrer",
    "article": "der"
  },
  {
    "noun": "Jogger",
    "article": "der"
  },
  {
    "noun": "Jogginganzug",
    "article": "der"
  },
  {
    "noun": "Joggingschuh",
    "article": "der"
  },
  {
    "noun": "Johannisbeersaft",
    "article": "der"
  },
  {
    "noun": "Johannisbrotbaum",
    "article": "der"
  },
  {
    "noun": "Johannistag",
    "article": "der"
  },
  {
    "noun": "Joint",
    "article": "der"
  },
  {
    "noun": "Joker",
    "article": "der"
  },
  {
    "noun": "Jongleur",
    "article": "der"
  },
  {
    "noun": "Jordan",
    "article": "der"
  },
  {
    "noun": "Jordanier",
    "article": "der"
  },
  {
    "noun": "Journalismus",
    "article": "der"
  },
  {
    "noun": "Journalist",
    "article": "der"
  },
  {
    "noun": "Joystick",
    "article": "der"
  },
  {
    "noun": "Jubel",
    "article": "der"
  },
  {
    "noun": "Jubilar",
    "article": "der"
  },
  {
    "noun": "Juckreiz",
    "article": "der"
  },
  {
    "noun": "Jugendclub",
    "article": "der"
  },
  {
    "noun": "Jugenderzieher",
    "article": "der"
  },
  {
    "noun": "Jugendfreund",
    "article": "der"
  },
  {
    "noun": "Jugendrichter",
    "article": "der"
  },
  {
    "noun": "Jugendschutz",
    "article": "der"
  },
  {
    "noun": "Jugendstil",
    "article": "der"
  },
  {
    "noun": "Jugendstreich",
    "article": "der"
  },
  {
    "noun": "Jugendtraum",
    "article": "der"
  },
  {
    "noun": "Jugendverband",
    "article": "der"
  },
  {
    "noun": "Jugoslawe",
    "article": "der"
  },
  {
    "noun": "Juli",
    "article": "der"
  },
  {
    "noun": "Jumbo",
    "article": "der"
  },
  {
    "noun": "Jumper",
    "article": "der"
  },
  {
    "noun": "Jungbrunnen",
    "article": "der"
  },
  {
    "noun": "Jungenstreich",
    "article": "der"
  },
  {
    "noun": "Jungfernflug",
    "article": "der"
  },
  {
    "noun": "Jungfernkranich",
    "article": "der"
  },
  {
    "noun": "Junggeselle",
    "article": "der"
  },
  {
    "noun": "Jungsozialist",
    "article": "der"
  },
  {
    "noun": "Jungunternehmer",
    "article": "der"
  },
  {
    "noun": "Jungvogel",
    "article": "der"
  },
  {
    "noun": "Juni",
    "article": "der"
  },
  {
    "noun": "Junior",
    "article": "der"
  },
  {
    "noun": "Juniorchef",
    "article": "der"
  },
  {
    "noun": "Juniorpartner",
    "article": "der"
  },
  {
    "noun": "Junker",
    "article": "der"
  },
  {
    "noun": "Junkie",
    "article": "der"
  },
  {
    "noun": "Jupiter",
    "article": "der"
  },
  {
    "noun": "Jurastudent",
    "article": "der"
  },
  {
    "noun": "Jurist",
    "article": "der"
  },
  {
    "noun": "Juso",
    "article": "der"
  },
  {
    "noun": "Justitiar",
    "article": "der"
  },
  {
    "noun": "Justizirrtum",
    "article": "der"
  },
  {
    "noun": "Justizminister",
    "article": "der"
  },
  {
    "noun": "Justizpalast",
    "article": "der"
  },
  {
    "noun": "Justizrat",
    "article": "der"
  },
  {
    "noun": "Jutesack",
    "article": "der"
  },
  {
    "noun": "Juwelier",
    "article": "der"
  },
  {
    "noun": "Juwelierladen",
    "article": "der"
  },
  {
    "noun": "Jux",
    "article": "der"
  },
  {
    "noun": "Charismatiker",
    "article": "der"
  },
  {
    "noun": "Charme",
    "article": "der"
  },
  {
    "noun": "Charmeur",
    "article": "der"
  },
  {
    "noun": "Charter",
    "article": "der"
  },
  {
    "noun": "Charterflug",
    "article": "der"
  },
  {
    "noun": "Charterverkehr",
    "article": "der"
  },
  {
    "noun": "Chauffeur",
    "article": "der"
  },
  {
    "noun": "Chausseebaum",
    "article": "der"
  },
  {
    "noun": "Chauvi",
    "article": "der"
  },
  {
    "noun": "Chauvinismus",
    "article": "der"
  },
  {
    "noun": "Chauvinist",
    "article": "der"
  },
  {
    "noun": "Check",
    "article": "der"
  },
  {
    "noun": "Checkpoint",
    "article": "der"
  },
  {
    "noun": "Cheerleader",
    "article": "der"
  },
  {
    "noun": "Cheeseburger",
    "article": "der"
  },
  {
    "noun": "Chef",
    "article": "der"
  },
  {
    "noun": "Chefarzt",
    "article": "der"
  },
  {
    "noun": "Chefingenieur",
    "article": "der"
  },
  {
    "noun": "Chefkoch",
    "article": "der"
  },
  {
    "noun": "Chefredakteur",
    "article": "der"
  },
  {
    "noun": "Chemieingenieur",
    "article": "der"
  },
  {
    "noun": "Chemielaborant",
    "article": "der"
  },
  {
    "noun": "Chemikant",
    "article": "der"
  },
  {
    "noun": "Chemiker",
    "article": "der"
  },
  {
    "noun": "Chemismus",
    "article": "der"
  },
  {
    "noun": "Chemorezeptor",
    "article": "der"
  },
  {
    "noun": "Chemotechniker",
    "article": "der"
  },
  {
    "noun": "Cherub",
    "article": "der"
  },
  {
    "noun": "Chiasmus",
    "article": "der"
  },
  {
    "noun": "Chiffon",
    "article": "der"
  },
  {
    "noun": "Chiffreur",
    "article": "der"
  },
  {
    "noun": "Chilene",
    "article": "der"
  },
  {
    "noun": "Chilesalpeter",
    "article": "der"
  },
  {
    "noun": "Chili",
    "article": "der"
  },
  {
    "noun": "Chinakohl",
    "article": "der"
  },
  {
    "noun": "Chinese",
    "article": "der"
  },
  {
    "noun": "Chinook",
    "article": "der"
  },
  {
    "noun": "Chintz",
    "article": "der"
  },
  {
    "noun": "Chip",
    "article": "der"
  },
  {
    "noun": "Chiromant",
    "article": "der"
  },
  {
    "noun": "Chiropraktiker",
    "article": "der"
  },
  {
    "noun": "Chirurg",
    "article": "der"
  },
  {
    "noun": "Chiton",
    "article": "der"
  },
  {
    "noun": "Chlorkalk",
    "article": "der"
  },
  {
    "noun": "Chlorwasserstoff",
    "article": "der"
  },
  {
    "noun": "Choke",
    "article": "der"
  },
  {
    "noun": "Choleriker",
    "article": "der"
  },
  {
    "noun": "Cholesterinspiegel",
    "article": "der"
  },
  {
    "noun": "Chor",
    "article": "der"
  },
  {
    "noun": "Choral",
    "article": "der"
  },
  {
    "noun": "Choreograf",
    "article": "der"
  },
  {
    "noun": "Chorgesang",
    "article": "der"
  },
  {
    "noun": "Chorherr",
    "article": "der"
  },
  {
    "noun": "Chorknabe",
    "article": "der"
  },
  {
    "noun": "Chorleiter",
    "article": "der"
  },
  {
    "noun": "Chorregent",
    "article": "der"
  },
  {
    "noun": "Chromit",
    "article": "der"
  },
  {
    "noun": "Chromnickelstahl",
    "article": "der"
  },
  {
    "noun": "Chronist",
    "article": "der"
  },
  {
    "noun": "Chronograf",
    "article": "der"
  },
  {
    "noun": "Chronologe",
    "article": "der"
  },
  {
    "noun": "Chrysoberyll",
    "article": "der"
  },
  {
    "noun": "Chrysolith",
    "article": "der"
  },
  {
    "noun": "Chrysopras",
    "article": "der"
  },
  {
    "noun": "Chrysotil",
    "article": "der"
  },
  {
    "noun": "Chylus",
    "article": "der"
  },
  {
    "noun": "Cineast",
    "article": "der"
  },
  {
    "noun": "Cinquecentist",
    "article": "der"
  },
  {
    "noun": "Circus",
    "article": "der"
  },
  {
    "noun": "Circusartist",
    "article": "der"
  },
  {
    "noun": "Circusclown",
    "article": "der"
  },
  {
    "noun": "Circusdirektor",
    "article": "der"
  },
  {
    "noun": "Circusreiter",
    "article": "der"
  },
  {
    "noun": "Circuswagen",
    "article": "der"
  },
  {
    "noun": "Clan",
    "article": "der"
  },
  {
    "noun": "Clinch",
    "article": "der"
  },
  {
    "noun": "Clown",
    "article": "der"
  },
  {
    "noun": "Clownfisch",
    "article": "der"
  },
  {
    "noun": "Club",
    "article": "der"
  },
  {
    "noun": "Clubsessel",
    "article": "der"
  },
  {
    "noun": "Cluster",
    "article": "der"
  },
  {
    "noun": "Coach",
    "article": "der"
  },
  {
    "noun": "Cockerspaniel",
    "article": "der"
  },
  {
    "noun": "Cocktail",
    "article": "der"
  },
  {
    "noun": "Cocktailhappen",
    "article": "der"
  },
  {
    "noun": "Code",
    "article": "der"
  },
  {
    "noun": "Coelestin",
    "article": "der"
  },
  {
    "noun": "Comic",
    "article": "der"
  },
  {
    "noun": "Comicstrip",
    "article": "der"
  },
  {
    "noun": "Compiler",
    "article": "der"
  },
  {
    "noun": "Computer",
    "article": "der"
  },
  {
    "noun": "Computerausdruck",
    "article": "der"
  },
  {
    "noun": "Computerfachmann",
    "article": "der"
  },
  {
    "noun": "Computersatz",
    "article": "der"
  },
  {
    "noun": "Computerspezialist",
    "article": "der"
  },
  {
    "noun": "Computertomograf",
    "article": "der"
  },
  {
    "noun": "Confiseur",
    "article": "der"
  },
  {
    "noun": "Container",
    "article": "der"
  },
  {
    "noun": "Containerbahnhof",
    "article": "der"
  },
  {
    "noun": "Containerhafen",
    "article": "der"
  },
  {
    "noun": "Containerlastzug",
    "article": "der"
  },
  {
    "noun": "Containerverkehr",
    "article": "der"
  },
  {
    "noun": "Contest",
    "article": "der"
  },
  {
    "noun": "Controller",
    "article": "der"
  },
  {
    "noun": "Converter",
    "article": "der"
  },
  {
    "noun": "Copilot",
    "article": "der"
  },
  {
    "noun": "Coprozessor",
    "article": "der"
  },
  {
    "noun": "Cord",
    "article": "der"
  },
  {
    "noun": "Cordanzug",
    "article": "der"
  },
  {
    "noun": "Cordsamt",
    "article": "der"
  },
  {
    "noun": "Coroner",
    "article": "der"
  },
  {
    "noun": "Corpsgeist",
    "article": "der"
  },
  {
    "noun": "Corso",
    "article": "der"
  },
  {
    "noun": "Couchtisch",
    "article": "der"
  },
  {
    "noun": "Coup",
    "article": "der"
  },
  {
    "noun": "Coupon",
    "article": "der"
  },
  {
    "noun": "Cousin",
    "article": "der"
  },
  {
    "noun": "Cowboyhut",
    "article": "der"
  },
  {
    "noun": "Coyote",
    "article": "der"
  },
  {
    "noun": "Crash",
    "article": "der"
  },
  {
    "noun": "Crashkurs",
    "article": "der"
  },
  {
    "noun": "Crashtest",
    "article": "der"
  },
  {
    "noun": "Cristobalit",
    "article": "der"
  },
  {
    "noun": "Croupier",
    "article": "der"
  },
  {
    "noun": "Csardas",
    "article": "der"
  },
  {
    "noun": "Cunnilingus",
    "article": "der"
  },
  {
    "noun": "Cursor",
    "article": "der"
  },
  {
    "noun": "Cutter",
    "article": "der"
  },
  {
    "noun": "Cyberspace",
    "article": "der"
  },
  {
    "noun": "Obstanbau",
    "article": "der"
  },
  {
    "noun": "Obstbau",
    "article": "der"
  },
  {
    "noun": "Obstbaum",
    "article": "der"
  },
  {
    "noun": "Obstbrand",
    "article": "der"
  },
  {
    "noun": "Obstessig",
    "article": "der"
  },
  {
    "noun": "Obstgarten",
    "article": "der"
  },
  {
    "noun": "Obstkern",
    "article": "der"
  },
  {
    "noun": "Obstkuchen",
    "article": "der"
  },
  {
    "noun": "Obstler",
    "article": "der"
  },
  {
    "noun": "Obstsaft",
    "article": "der"
  },
  {
    "noun": "Obstsalat",
    "article": "der"
  },
  {
    "noun": "Obstschnaps",
    "article": "der"
  },
  {
    "noun": "Obsttag",
    "article": "der"
  },
  {
    "noun": "Obstwein",
    "article": "der"
  },
  {
    "noun": "Obus",
    "article": "der"
  },
  {
    "noun": "Ochs",
    "article": "der"
  },
  {
    "noun": "Ochsenfrosch",
    "article": "der"
  },
  {
    "noun": "Ochsenschwanz",
    "article": "der"
  },
  {
    "noun": "Ochsenziemer",
    "article": "der"
  },
  {
    "noun": "Odem",
    "article": "der"
  },
  {
    "noun": "Odor",
    "article": "der"
  },
  {
    "noun": "Ofen",
    "article": "der"
  },
  {
    "noun": "Ofenbauer",
    "article": "der"
  },
  {
    "noun": "Ofenhaken",
    "article": "der"
  },
  {
    "noun": "Ofenhocker",
    "article": "der"
  },
  {
    "noun": "Ofenschirm",
    "article": "der"
  },
  {
    "noun": "Ofensetzer",
    "article": "der"
  },
  {
    "noun": "Offenbarungseid",
    "article": "der"
  },
  {
    "noun": "Offensivkrieg",
    "article": "der"
  },
  {
    "noun": "Offiziant",
    "article": "der"
  },
  {
    "noun": "Offizier",
    "article": "der"
  },
  {
    "noun": "Offiziersrang",
    "article": "der"
  },
  {
    "noun": "Offroader",
    "article": "der"
  },
  {
    "noun": "Offsetdruck",
    "article": "der"
  },
  {
    "noun": "Oheim",
    "article": "der"
  },
  {
    "noun": "Ohnmachtsanfall",
    "article": "der"
  },
  {
    "noun": "Ohrenarzt",
    "article": "der"
  },
  {
    "noun": "Ohrenfluss",
    "article": "der"
  },
  {
    "noun": "Ohrenschmerz",
    "article": "der"
  },
  {
    "noun": "Ohrensessel",
    "article": "der"
  },
  {
    "noun": "Ohrenspiegel",
    "article": "der"
  },
  {
    "noun": "Ohrenzeuge",
    "article": "der"
  },
  {
    "noun": "Ohrring",
    "article": "der"
  },
  {
    "noun": "Ohrstecker",
    "article": "der"
  },
  {
    "noun": "Ohrwurm",
    "article": "der"
  },
  {
    "noun": "Okklusiv",
    "article": "der"
  },
  {
    "noun": "Okkultismus",
    "article": "der"
  },
  {
    "noun": "Okkultist",
    "article": "der"
  },
  {
    "noun": "Okkupant",
    "article": "der"
  },
  {
    "noun": "Oktant",
    "article": "der"
  },
  {
    "noun": "Oktavband",
    "article": "der"
  },
  {
    "noun": "Oktober",
    "article": "der"
  },
  {
    "noun": "Okzident",
    "article": "der"
  },
  {
    "noun": "Oldtimer",
    "article": "der"
  },
  {
    "noun": "Oleander",
    "article": "der"
  },
  {
    "noun": "Oleaster",
    "article": "der"
  },
  {
    "noun": "Oligarch",
    "article": "der"
  },
  {
    "noun": "Olivenbaum",
    "article": "der"
  },
  {
    "noun": "Olivenhain",
    "article": "der"
  },
  {
    "noun": "Olm",
    "article": "der"
  },
  {
    "noun": "Olymp",
    "article": "der"
  },
  {
    "noun": "Olympier",
    "article": "der"
  },
  {
    "noun": "Ombrograf",
    "article": "der"
  },
  {
    "noun": "Ombudsmann",
    "article": "der"
  },
  {
    "noun": "Omnibus",
    "article": "der"
  },
  {
    "noun": "Omnibusbahnhof",
    "article": "der"
  },
  {
    "noun": "Onkel",
    "article": "der"
  },
  {
    "noun": "Onkologe",
    "article": "der"
  },
  {
    "noun": "Onyx",
    "article": "der"
  },
  {
    "noun": "Oolith",
    "article": "der"
  },
  {
    "noun": "Opa",
    "article": "der"
  },
  {
    "noun": "Opal",
    "article": "der"
  },
  {
    "noun": "Operand",
    "article": "der"
  },
  {
    "noun": "Operateur",
    "article": "der"
  },
  {
    "noun": "Operationsbericht",
    "article": "der"
  },
  {
    "noun": "Operationsraum",
    "article": "der"
  },
  {
    "noun": "Operationssaal",
    "article": "der"
  },
  {
    "noun": "Operationstisch",
    "article": "der"
  },
  {
    "noun": "Operator",
    "article": "der"
  },
  {
    "noun": "Opernball",
    "article": "der"
  },
  {
    "noun": "Operngucker",
    "article": "der"
  },
  {
    "noun": "Opferaltar",
    "article": "der"
  },
  {
    "noun": "Opferschutz",
    "article": "der"
  },
  {
    "noun": "Opferstock",
    "article": "der"
  },
  {
    "noun": "Ophthalmologe",
    "article": "der"
  },
  {
    "noun": "Opisthotonus",
    "article": "der"
  },
  {
    "noun": "Opiumraucher",
    "article": "der"
  },
  {
    "noun": "Opiumschmuggel",
    "article": "der"
  },
  {
    "noun": "Opponent",
    "article": "der"
  },
  {
    "noun": "Opportunismus",
    "article": "der"
  },
  {
    "noun": "Opportunist",
    "article": "der"
  },
  {
    "noun": "Oppositionsgeist",
    "article": "der"
  },
  {
    "noun": "Optativ",
    "article": "der"
  },
  {
    "noun": "Optiker",
    "article": "der"
  },
  {
    "noun": "Optimismus",
    "article": "der"
  },
  {
    "noun": "Optimist",
    "article": "der"
  },
  {
    "noun": "Optionshandel",
    "article": "der"
  },
  {
    "noun": "Optionsmarkt",
    "article": "der"
  },
  {
    "noun": "Optionsschein",
    "article": "der"
  },
  {
    "noun": "Orakelspruch",
    "article": "der"
  },
  {
    "noun": "Oralverkehr",
    "article": "der"
  },
  {
    "noun": "Orangenbaum",
    "article": "der"
  },
  {
    "noun": "Orangensaft",
    "article": "der"
  },
  {
    "noun": "Orbit",
    "article": "der"
  },
  {
    "noun": "Orchestergraben",
    "article": "der"
  },
  {
    "noun": "Orchesterleiter",
    "article": "der"
  },
  {
    "noun": "Orden",
    "article": "der"
  },
  {
    "noun": "Ordensmann",
    "article": "der"
  },
  {
    "noun": "Orderscheck",
    "article": "der"
  },
  {
    "noun": "Ordinarius",
    "article": "der"
  },
  {
    "noun": "Ordner",
    "article": "der"
  },
  {
    "noun": "Ordnungsfimmel",
    "article": "der"
  },
  {
    "noun": "Ordnungsruf",
    "article": "der"
  },
  {
    "noun": "Oregano",
    "article": "der"
  },
  {
    "noun": "Organdy",
    "article": "der"
  },
  {
    "noun": "Organhandel",
    "article": "der"
  },
  {
    "noun": "Organiker",
    "article": "der"
  },
  {
    "noun": "Organisationsgrad",
    "article": "der"
  },
  {
    "noun": "Organisationsplan",
    "article": "der"
  },
  {
    "noun": "Organisator",
    "article": "der"
  },
  {
    "noun": "Organismus",
    "article": "der"
  },
  {
    "noun": "Organist",
    "article": "der"
  },
  {
    "noun": "Organspender",
    "article": "der"
  },
  {
    "noun": "Organza",
    "article": "der"
  },
  {
    "noun": "Orgelbauer",
    "article": "der"
  },
  {
    "noun": "Orient",
    "article": "der"
  },
  {
    "noun": "Orientale",
    "article": "der"
  },
  {
    "noun": "Orientalist",
    "article": "der"
  },
  {
    "noun": "Orientierungslauf",
    "article": "der"
  },
  {
    "noun": "Orientierungspunkt",
    "article": "der"
  },
  {
    "noun": "Orientierungssinn",
    "article": "der"
  },
  {
    "noun": "Orientteppich",
    "article": "der"
  },
  {
    "noun": "Originalbeleg",
    "article": "der"
  },
  {
    "noun": "Originaltext",
    "article": "der"
  },
  {
    "noun": "Originalton",
    "article": "der"
  },
  {
    "noun": "Orion",
    "article": "der"
  },
  {
    "noun": "Orkan",
    "article": "der"
  },
  {
    "noun": "Orkus",
    "article": "der"
  },
  {
    "noun": "Ornat",
    "article": "der"
  },
  {
    "noun": "Ornithologe",
    "article": "der"
  },
  {
    "noun": "Ortgang",
    "article": "der"
  },
  {
    "noun": "Ortolan",
    "article": "der"
  },
  {
    "noun": "Ortseingang",
    "article": "der"
  },
  {
    "noun": "Ortsname",
    "article": "der"
  },
  {
    "noun": "Ortssinn",
    "article": "der"
  },
  {
    "noun": "Ortstarif",
    "article": "der"
  },
  {
    "noun": "Ortsteil",
    "article": "der"
  },
  {
    "noun": "Ortstein",
    "article": "der"
  },
  {
    "noun": "Ortsvektor",
    "article": "der"
  },
  {
    "noun": "Ortsverkehr",
    "article": "der"
  },
  {
    "noun": "Ortsvorsteher",
    "article": "der"
  },
  {
    "noun": "Ortswechsel",
    "article": "der"
  },
  {
    "noun": "Ortszuschlag",
    "article": "der"
  },
  {
    "noun": "Oscar",
    "article": "der"
  },
  {
    "noun": "Osmane",
    "article": "der"
  },
  {
    "noun": "Ostblock",
    "article": "der"
  },
  {
    "noun": "Osterbrauch",
    "article": "der"
  },
  {
    "noun": "Osterhase",
    "article": "der"
  },
  {
    "noun": "Ostermontag",
    "article": "der"
  },
  {
    "noun": "Ostersonntag",
    "article": "der"
  },
  {
    "noun": "Ostwind",
    "article": "der"
  },
  {
    "noun": "Oszillator",
    "article": "der"
  },
  {
    "noun": "Oszillograf",
    "article": "der"
  },
  {
    "noun": "Otiater",
    "article": "der"
  },
  {
    "noun": "Otologe",
    "article": "der"
  },
  {
    "noun": "Ottokraftstoff",
    "article": "der"
  },
  {
    "noun": "Ottomane",
    "article": "der"
  },
  {
    "noun": "Ottomotor",
    "article": "der"
  },
  {
    "noun": "Overall",
    "article": "der"
  },
  {
    "noun": "Overheadprojektor",
    "article": "der"
  },
  {
    "noun": "Ovulationshemmer",
    "article": "der"
  },
  {
    "noun": "Oxer",
    "article": "der"
  },
  {
    "noun": "Ozean",
    "article": "der"
  },
  {
    "noun": "Ozeandampfer",
    "article": "der"
  },
  {
    "noun": "Ozeanograf",
    "article": "der"
  },
  {
    "noun": "Ozeanriese",
    "article": "der"
  },
  {
    "noun": "Ozelot",
    "article": "der"
  },
  {
    "noun": "Ozokerit",
    "article": "der"
  },
  {
    "noun": "Ozongehalt",
    "article": "der"
  },
  {
    "noun": "Ozonsmog",
    "article": "der"
  },
  {
    "noun": "Immoralist",
    "article": "der"
  },
  {
    "noun": "Imperativ",
    "article": "der"
  },
  {
    "noun": "Imperativsatz",
    "article": "der"
  },
  {
    "noun": "Imperator",
    "article": "der"
  },
  {
    "noun": "Imperialismus",
    "article": "der"
  },
  {
    "noun": "Impetus",
    "article": "der"
  },
  {
    "noun": "Impfarzt",
    "article": "der"
  },
  {
    "noun": "Impfausweis",
    "article": "der"
  },
  {
    "noun": "Impfschein",
    "article": "der"
  },
  {
    "noun": "Impfschutz",
    "article": "der"
  },
  {
    "noun": "Impfstoff",
    "article": "der"
  },
  {
    "noun": "Impfzwang",
    "article": "der"
  },
  {
    "noun": "Import",
    "article": "der"
  },
  {
    "noun": "Importanteil",
    "article": "der"
  },
  {
    "noun": "Importeur",
    "article": "der"
  },
  {
    "noun": "Importhandel",
    "article": "der"
  },
  {
    "noun": "Importzoll",
    "article": "der"
  },
  {
    "noun": "Impressionismus",
    "article": "der"
  },
  {
    "noun": "Impressionist",
    "article": "der"
  },
  {
    "noun": "Improvisator",
    "article": "der"
  },
  {
    "noun": "Impuls",
    "article": "der"
  },
  {
    "noun": "Impulsgeber",
    "article": "der"
  },
  {
    "noun": "Impulsgenerator",
    "article": "der"
  },
  {
    "noun": "Impulssatz",
    "article": "der"
  },
  {
    "noun": "Inbegriff",
    "article": "der"
  },
  {
    "noun": "Inch",
    "article": "der"
  },
  {
    "noun": "Inder",
    "article": "der"
  },
  {
    "noun": "Indeterminismus",
    "article": "der"
  },
  {
    "noun": "Index",
    "article": "der"
  },
  {
    "noun": "Indexlohn",
    "article": "der"
  },
  {
    "noun": "Indianer",
    "article": "der"
  },
  {
    "noun": "Indianerstamm",
    "article": "der"
  },
  {
    "noun": "Indifferentismus",
    "article": "der"
  },
  {
    "noun": "Indigolith",
    "article": "der"
  },
  {
    "noun": "Indikativ",
    "article": "der"
  },
  {
    "noun": "Indikator",
    "article": "der"
  },
  {
    "noun": "Indio",
    "article": "der"
  },
  {
    "noun": "Individualbereich",
    "article": "der"
  },
  {
    "noun": "Individualismus",
    "article": "der"
  },
  {
    "noun": "Individualist",
    "article": "der"
  },
  {
    "noun": "Individualverkehr",
    "article": "der"
  },
  {
    "noun": "Indologe",
    "article": "der"
  },
  {
    "noun": "Indonesier",
    "article": "der"
  },
  {
    "noun": "Indossant",
    "article": "der"
  },
  {
    "noun": "Indossatar",
    "article": "der"
  },
  {
    "noun": "Induktionsbeweis",
    "article": "der"
  },
  {
    "noun": "Induktionsmotor",
    "article": "der"
  },
  {
    "noun": "Induktionsofen",
    "article": "der"
  },
  {
    "noun": "Induktionsstrom",
    "article": "der"
  },
  {
    "noun": "Industrialismus",
    "article": "der"
  },
  {
    "noun": "Industriearbeiter",
    "article": "der"
  },
  {
    "noun": "Industriebau",
    "article": "der"
  },
  {
    "noun": "Industrieberater",
    "article": "der"
  },
  {
    "noun": "Industriebetrieb",
    "article": "der"
  },
  {
    "noun": "Industriediamant",
    "article": "der"
  },
  {
    "noun": "Industriekaufmann",
    "article": "der"
  },
  {
    "noun": "Industriekonzern",
    "article": "der"
  },
  {
    "noun": "Industriemagnat",
    "article": "der"
  },
  {
    "noun": "Industriemeister",
    "article": "der"
  },
  {
    "noun": "Industrieofen",
    "article": "der"
  },
  {
    "noun": "Industriepark",
    "article": "der"
  },
  {
    "noun": "Industrieroboter",
    "article": "der"
  },
  {
    "noun": "Industriestaat",
    "article": "der"
  },
  {
    "noun": "Industriezweig",
    "article": "der"
  },
  {
    "noun": "Infanterist",
    "article": "der"
  },
  {
    "noun": "Infantilismus",
    "article": "der"
  },
  {
    "noun": "Infarkt",
    "article": "der"
  },
  {
    "noun": "Infekt",
    "article": "der"
  },
  {
    "noun": "Infektionsherd",
    "article": "der"
  },
  {
    "noun": "Infiltrant",
    "article": "der"
  },
  {
    "noun": "Infinitiv",
    "article": "der"
  },
  {
    "noun": "Infinitivsatz",
    "article": "der"
  },
  {
    "noun": "Inflationsgewinn",
    "article": "der"
  },
  {
    "noun": "Informant",
    "article": "der"
  },
  {
    "noun": "Informatiker",
    "article": "der"
  },
  {
    "noun": "Informationsaustausch",
    "article": "der"
  },
  {
    "noun": "Informationsbedarf",
    "article": "der"
  },
  {
    "noun": "Informationsbesuch",
    "article": "der"
  },
  {
    "noun": "Informationsdienst",
    "article": "der"
  },
  {
    "noun": "Informationsfluss",
    "article": "der"
  },
  {
    "noun": "Informationsgehalt",
    "article": "der"
  },
  {
    "noun": "Informationsstand",
    "article": "der"
  },
  {
    "noun": "Informationswert",
    "article": "der"
  },
  {
    "noun": "Informator",
    "article": "der"
  },
  {
    "noun": "Infrarotfilm",
    "article": "der"
  },
  {
    "noun": "Infrarotstrahl",
    "article": "der"
  },
  {
    "noun": "Infrarotstrahler",
    "article": "der"
  },
  {
    "noun": "Infraschall",
    "article": "der"
  },
  {
    "noun": "Ingenieur",
    "article": "der"
  },
  {
    "noun": "Ingenieurbau",
    "article": "der"
  },
  {
    "noun": "Ingwer",
    "article": "der"
  },
  {
    "noun": "Inhaber",
    "article": "der"
  },
  {
    "noun": "Inhalt",
    "article": "der"
  },
  {
    "noun": "Inhibitor",
    "article": "der"
  },
  {
    "noun": "Initiationsritus",
    "article": "der"
  },
  {
    "noun": "Initiator",
    "article": "der"
  },
  {
    "noun": "Injektiv",
    "article": "der"
  },
  {
    "noun": "Injektivlaut",
    "article": "der"
  },
  {
    "noun": "Injektor",
    "article": "der"
  },
  {
    "noun": "Inka",
    "article": "der"
  },
  {
    "noun": "Inklusivpreis",
    "article": "der"
  },
  {
    "noun": "Inkreis",
    "article": "der"
  },
  {
    "noun": "Inkubator",
    "article": "der"
  },
  {
    "noun": "Inlaid",
    "article": "der"
  },
  {
    "noun": "Innenarchitekt",
    "article": "der"
  },
  {
    "noun": "Innendienst",
    "article": "der"
  },
  {
    "noun": "Innendruck",
    "article": "der"
  },
  {
    "noun": "Innendurchmesser",
    "article": "der"
  },
  {
    "noun": "Innenhof",
    "article": "der"
  },
  {
    "noun": "Innenminister",
    "article": "der"
  },
  {
    "noun": "Innenraum",
    "article": "der"
  },
  {
    "noun": "Innenrist",
    "article": "der"
  },
  {
    "noun": "Innenski",
    "article": "der"
  },
  {
    "noun": "Innenspiegel",
    "article": "der"
  },
  {
    "noun": "Innenwinkel",
    "article": "der"
  },
  {
    "noun": "Innovationsspross",
    "article": "der"
  },
  {
    "noun": "Inquisitor",
    "article": "der"
  },
  {
    "noun": "Insasse",
    "article": "der"
  },
  {
    "noun": "Insektenfresser",
    "article": "der"
  },
  {
    "noun": "Inselbewohner",
    "article": "der"
  },
  {
    "noun": "Inselstaat",
    "article": "der"
  },
  {
    "noun": "Inserent",
    "article": "der"
  },
  {
    "noun": "Insider",
    "article": "der"
  },
  {
    "noun": "Inspektionsgang",
    "article": "der"
  },
  {
    "noun": "Inspektor",
    "article": "der"
  },
  {
    "noun": "Inspirator",
    "article": "der"
  },
  {
    "noun": "Inspizient",
    "article": "der"
  },
  {
    "noun": "Installateur",
    "article": "der"
  },
  {
    "noun": "Instantkaffee",
    "article": "der"
  },
  {
    "noun": "Instanzenweg",
    "article": "der"
  },
  {
    "noun": "Instanzenzug",
    "article": "der"
  },
  {
    "noun": "Instinkt",
    "article": "der"
  },
  {
    "noun": "Institutionalismus",
    "article": "der"
  },
  {
    "noun": "Institutsleiter",
    "article": "der"
  },
  {
    "noun": "Instruktor",
    "article": "der"
  },
  {
    "noun": "Instrumentenbau",
    "article": "der"
  },
  {
    "noun": "Instrumentenkasten",
    "article": "der"
  },
  {
    "noun": "Instrumentenkoffer",
    "article": "der"
  },
  {
    "noun": "Instrumentenschrank",
    "article": "der"
  },
  {
    "noun": "Insulaner",
    "article": "der"
  },
  {
    "noun": "Insult",
    "article": "der"
  },
  {
    "noun": "Insurgent",
    "article": "der"
  },
  {
    "noun": "Integrand",
    "article": "der"
  },
  {
    "noun": "Integrationsprozess",
    "article": "der"
  },
  {
    "noun": "Integrator",
    "article": "der"
  },
  {
    "noun": "Intellekt",
    "article": "der"
  },
  {
    "noun": "Intellektualismus",
    "article": "der"
  },
  {
    "noun": "Intelligenzquotient",
    "article": "der"
  },
  {
    "noun": "Intelligenztest",
    "article": "der"
  },
  {
    "noun": "Intendant",
    "article": "der"
  },
  {
    "noun": "Intensivanbau",
    "article": "der"
  },
  {
    "noun": "Intensivkurs",
    "article": "der"
  },
  {
    "noun": "Intentionstremor",
    "article": "der"
  },
  {
    "noun": "Interessenausgleich",
    "article": "der"
  },
  {
    "noun": "Interessengegensatz",
    "article": "der"
  },
  {
    "noun": "Interessenkonflikt",
    "article": "der"
  },
  {
    "noun": "Interessent",
    "article": "der"
  },
  {
    "noun": "Interessenverband",
    "article": "der"
  },
  {
    "noun": "Interessenvertreter",
    "article": "der"
  },
  {
    "noun": "Interimsbescheid",
    "article": "der"
  },
  {
    "noun": "Interimsschein",
    "article": "der"
  },
  {
    "noun": "Internationalismus",
    "article": "der"
  },
  {
    "noun": "Internetanbieter",
    "article": "der"
  },
  {
    "noun": "Internetanschluss",
    "article": "der"
  },
  {
    "noun": "Internetauftritt",
    "article": "der"
  },
  {
    "noun": "Internetzugang",
    "article": "der"
  },
  {
    "noun": "Internist",
    "article": "der"
  },
  {
    "noun": "Interpellant",
    "article": "der"
  },
  {
    "noun": "Interpolator",
    "article": "der"
  },
  {
    "noun": "Interpret",
    "article": "der"
  },
  {
    "noun": "Interpreter",
    "article": "der"
  },
  {
    "noun": "Interrogativsatz",
    "article": "der"
  },
  {
    "noun": "Interventionismus",
    "article": "der"
  },
  {
    "noun": "Interviewer",
    "article": "der"
  },
  {
    "noun": "Intrigant",
    "article": "der"
  },
  {
    "noun": "Inuk",
    "article": "der"
  },
  {
    "noun": "Invertebrat",
    "article": "der"
  },
  {
    "noun": "Inverter",
    "article": "der"
  },
  {
    "noun": "Invertzucker",
    "article": "der"
  },
  {
    "noun": "Investitionshaushalt",
    "article": "der"
  },
  {
    "noun": "Investitionskredit",
    "article": "der"
  },
  {
    "noun": "Investiturstreit",
    "article": "der"
  },
  {
    "noun": "Investmentfonds",
    "article": "der"
  },
  {
    "noun": "Investor",
    "article": "der"
  },
  {
    "noun": "Inzest",
    "article": "der"
  },
  {
    "noun": "Inzidenzwinkel",
    "article": "der"
  },
  {
    "noun": "Inzisiv",
    "article": "der"
  },
  {
    "noun": "Ionenaustausch",
    "article": "der"
  },
  {
    "noun": "Ionenaustauscher",
    "article": "der"
  },
  {
    "noun": "Ionenbeschleuniger",
    "article": "der"
  },
  {
    "noun": "Ionenstrahl",
    "article": "der"
  },
  {
    "noun": "Ionisator",
    "article": "der"
  },
  {
    "noun": "Irak",
    "article": "der"
  },
  {
    "noun": "Iraker",
    "article": "der"
  },
  {
    "noun": "Iran",
    "article": "der"
  },
  {
    "noun": "Iraner",
    "article": "der"
  },
  {
    "noun": "Ire",
    "article": "der"
  },
  {
    "noun": "Irokesenschnitt",
    "article": "der"
  },
  {
    "noun": "Ironiker",
    "article": "der"
  },
  {
    "noun": "Irreal",
    "article": "der"
  },
  {
    "noun": "Irredentismus",
    "article": "der"
  },
  {
    "noun": "Irredentist",
    "article": "der"
  },
  {
    "noun": "Irrenarzt",
    "article": "der"
  },
  {
    "noun": "Irrgang",
    "article": "der"
  },
  {
    "noun": "Irrgarten",
    "article": "der"
  },
  {
    "noun": "Irrglaube",
    "article": "der"
  },
  {
    "noun": "Irrsinn",
    "article": "der"
  },
  {
    "noun": "Irrtum",
    "article": "der"
  },
  {
    "noun": "Irrweg",
    "article": "der"
  },
  {
    "noun": "Irrwisch",
    "article": "der"
  },
  {
    "noun": "Irrwitz",
    "article": "der"
  },
  {
    "noun": "Ischiasnerv",
    "article": "der"
  },
  {
    "noun": "Islam",
    "article": "der"
  },
  {
    "noun": "Isolationismus",
    "article": "der"
  },
  {
    "noun": "Isolationist",
    "article": "der"
  },
  {
    "noun": "Isolator",
    "article": "der"
  },
  {
    "noun": "Isolierer",
    "article": "der"
  },
  {
    "noun": "Isolierlack",
    "article": "der"
  },
  {
    "noun": "Isolierstoff",
    "article": "der"
  },
  {
    "noun": "Isomorphismus",
    "article": "der"
  },
  {
    "noun": "Israelit",
    "article": "der"
  },
  {
    "noun": "Isthmus",
    "article": "der"
  },
  {
    "noun": "Istrier",
    "article": "der"
  },
  {
    "noun": "Itaker",
    "article": "der"
  },
  {
    "noun": "Italiener",
    "article": "der"
  },
  {
    "noun": "Ubiquist",
    "article": "der"
  },
  {
    "noun": "Uferbau",
    "article": "der"
  },
  {
    "noun": "Ufologe",
    "article": "der"
  },
  {
    "noun": "Ugander",
    "article": "der"
  },
  {
    "noun": "Uhrkasten",
    "article": "der"
  },
  {
    "noun": "Uhrmacher",
    "article": "der"
  },
  {
    "noun": "Uhrzeiger",
    "article": "der"
  },
  {
    "noun": "Uhrzeigersinn",
    "article": "der"
  },
  {
    "noun": "Uhu",
    "article": "der"
  },
  {
    "noun": "Ukas",
    "article": "der"
  },
  {
    "noun": "Ukelei",
    "article": "der"
  },
  {
    "noun": "Ukrainer",
    "article": "der"
  },
  {
    "noun": "Ulan",
    "article": "der"
  },
  {
    "noun": "Ulema",
    "article": "der"
  },
  {
    "noun": "Ulk",
    "article": "der"
  },
  {
    "noun": "Ultimo",
    "article": "der"
  },
  {
    "noun": "Ultra",
    "article": "der"
  },
  {
    "noun": "Ultrakurzwellensender",
    "article": "der"
  },
  {
    "noun": "Ultramontanismus",
    "article": "der"
  },
  {
    "noun": "Ultraschall",
    "article": "der"
  },
  {
    "noun": "Umbau",
    "article": "der"
  },
  {
    "noun": "Umber",
    "article": "der"
  },
  {
    "noun": "Umbruch",
    "article": "der"
  },
  {
    "noun": "Umdruck",
    "article": "der"
  },
  {
    "noun": "Umfall",
    "article": "der"
  },
  {
    "noun": "Umfang",
    "article": "der"
  },
  {
    "noun": "Umfangswinkel",
    "article": "der"
  },
  {
    "noun": "Umformer",
    "article": "der"
  },
  {
    "noun": "Umgang",
    "article": "der"
  },
  {
    "noun": "Umgangston",
    "article": "der"
  },
  {
    "noun": "Umhang",
    "article": "der"
  },
  {
    "noun": "Umkehrfilm",
    "article": "der"
  },
  {
    "noun": "Umkehrschluss",
    "article": "der"
  },
  {
    "noun": "Umkleideraum",
    "article": "der"
  },
  {
    "noun": "Umkreis",
    "article": "der"
  },
  {
    "noun": "Umladebahnhof",
    "article": "der"
  },
  {
    "noun": "Umlauf",
    "article": "der"
  },
  {
    "noun": "Umlaut",
    "article": "der"
  },
  {
    "noun": "Umlegekragen",
    "article": "der"
  },
  {
    "noun": "Umluftbackofen",
    "article": "der"
  },
  {
    "noun": "Umraum",
    "article": "der"
  },
  {
    "noun": "Umrechnungskurs",
    "article": "der"
  },
  {
    "noun": "Umrichter",
    "article": "der"
  },
  {
    "noun": "Umriss",
    "article": "der"
  },
  {
    "noun": "Umsatz",
    "article": "der"
  },
  {
    "noun": "Umsatzanstieg",
    "article": "der"
  },
  {
    "noun": "Umschalter",
    "article": "der"
  },
  {
    "noun": "Umschalthebel",
    "article": "der"
  },
  {
    "noun": "Umschlag",
    "article": "der"
  },
  {
    "noun": "Umschlaghafen",
    "article": "der"
  },
  {
    "noun": "Umschlagplatz",
    "article": "der"
  },
  {
    "noun": "Umschlagtitel",
    "article": "der"
  },
  {
    "noun": "Umschweif",
    "article": "der"
  },
  {
    "noun": "Umschwung",
    "article": "der"
  },
  {
    "noun": "Umsetzer",
    "article": "der"
  },
  {
    "noun": "Umsiedler",
    "article": "der"
  },
  {
    "noun": "Umspanner",
    "article": "der"
  },
  {
    "noun": "Umstand",
    "article": "der"
  },
  {
    "noun": "Umstandssatz",
    "article": "der"
  },
  {
    "noun": "Umstellbahnhof",
    "article": "der"
  },
  {
    "noun": "Umstellungsprozess",
    "article": "der"
  },
  {
    "noun": "Umsturz",
    "article": "der"
  },
  {
    "noun": "Umsturzversuch",
    "article": "der"
  },
  {
    "noun": "Umtausch",
    "article": "der"
  },
  {
    "noun": "Umtrieb",
    "article": "der"
  },
  {
    "noun": "Umtrunk",
    "article": "der"
  },
  {
    "noun": "Umwandler",
    "article": "der"
  },
  {
    "noun": "Umwandlungsprozess",
    "article": "der"
  },
  {
    "noun": "Umweg",
    "article": "der"
  },
  {
    "noun": "Umwelteinfluss",
    "article": "der"
  },
  {
    "noun": "Umweltfaktor",
    "article": "der"
  },
  {
    "noun": "Umweltminister",
    "article": "der"
  },
  {
    "noun": "Umweltschaden",
    "article": "der"
  },
  {
    "noun": "Umweltschutz",
    "article": "der"
  },
  {
    "noun": "Umwohner",
    "article": "der"
  },
  {
    "noun": "Umzug",
    "article": "der"
  },
  {
    "noun": "Umzugstag",
    "article": "der"
  },
  {
    "noun": "Undank",
    "article": "der"
  },
  {
    "noun": "Undulator",
    "article": "der"
  },
  {
    "noun": "Unfall",
    "article": "der"
  },
  {
    "noun": "Unfallbericht",
    "article": "der"
  },
  {
    "noun": "Unfallgegner",
    "article": "der"
  },
  {
    "noun": "Unfallhergang",
    "article": "der"
  },
  {
    "noun": "Unfallort",
    "article": "der"
  },
  {
    "noun": "Unfallschutz",
    "article": "der"
  },
  {
    "noun": "Unfallwagen",
    "article": "der"
  },
  {
    "noun": "Unfallzeuge",
    "article": "der"
  },
  {
    "noun": "Unflat",
    "article": "der"
  },
  {
    "noun": "Unfriede",
    "article": "der"
  },
  {
    "noun": "Unfug",
    "article": "der"
  },
  {
    "noun": "Ungehorsam",
    "article": "der"
  },
  {
    "noun": "Ungeist",
    "article": "der"
  },
  {
    "noun": "Unglaube",
    "article": "der"
  },
  {
    "noun": "Unglimpf",
    "article": "der"
  },
  {
    "noun": "Unheilstifter",
    "article": "der"
  },
  {
    "noun": "Unhold",
    "article": "der"
  },
  {
    "noun": "Uniformrock",
    "article": "der"
  },
  {
    "noun": "Unionist",
    "article": "der"
  },
  {
    "noun": "Universalerbe",
    "article": "der"
  },
  {
    "noun": "Universalismus",
    "article": "der"
  },
  {
    "noun": "Universalmotor",
    "article": "der"
  },
  {
    "noun": "Unkenruf",
    "article": "der"
  },
  {
    "noun": "Unmensch",
    "article": "der"
  },
  {
    "noun": "Unmut",
    "article": "der"
  },
  {
    "noun": "Unpaarhufer",
    "article": "der"
  },
  {
    "noun": "Unrat",
    "article": "der"
  },
  {
    "noun": "Unruheherd",
    "article": "der"
  },
  {
    "noun": "Unruhestifter",
    "article": "der"
  },
  {
    "noun": "Unruhstifter",
    "article": "der"
  },
  {
    "noun": "Unschuldsengel",
    "article": "der"
  },
  {
    "noun": "Unsegen",
    "article": "der"
  },
  {
    "noun": "Unsicherheitsfaktor",
    "article": "der"
  },
  {
    "noun": "Unsinn",
    "article": "der"
  },
  {
    "noun": "Unstern",
    "article": "der"
  },
  {
    "noun": "Unter",
    "article": "der"
  },
  {
    "noun": "Unterabschnitt",
    "article": "der"
  },
  {
    "noun": "Unterarm",
    "article": "der"
  },
  {
    "noun": "Unterbau",
    "article": "der"
  },
  {
    "noun": "Unterbauch",
    "article": "der"
  },
  {
    "noun": "Unterbegriff",
    "article": "der"
  },
  {
    "noun": "Unterboden",
    "article": "der"
  },
  {
    "noun": "Unterbodenschutz",
    "article": "der"
  },
  {
    "noun": "Unterbrecher",
    "article": "der"
  },
  {
    "noun": "Unterbruch",
    "article": "der"
  },
  {
    "noun": "Unterdruck",
    "article": "der"
  },
  {
    "noun": "Unterfahrschutz",
    "article": "der"
  },
  {
    "noun": "Untergang",
    "article": "der"
  },
  {
    "noun": "Untergrund",
    "article": "der"
  },
  {
    "noun": "Unterhalt",
    "article": "der"
  },
  {
    "noun": "Unterhalter",
    "article": "der"
  },
  {
    "noun": "Unterhaltsanspruch",
    "article": "der"
  },
  {
    "noun": "Unterhaltungsfilm",
    "article": "der"
  },
  {
    "noun": "Unterhaltungswert",
    "article": "der"
  },
  {
    "noun": "Unterkiefer",
    "article": "der"
  },
  {
    "noun": "Unterkieferknochen",
    "article": "der"
  },
  {
    "noun": "Unterlauf",
    "article": "der"
  },
  {
    "noun": "Unterlegkeil",
    "article": "der"
  },
  {
    "noun": "Unterlegring",
    "article": "der"
  },
  {
    "noun": "Unterleib",
    "article": "der"
  },
  {
    "noun": "Untermieter",
    "article": "der"
  },
  {
    "noun": "Unternehmensberater",
    "article": "der"
  },
  {
    "noun": "Unternehmensleiter",
    "article": "der"
  },
  {
    "noun": "Unternehmenszusammenschluss",
    "article": "der"
  },
  {
    "noun": "Unternehmer",
    "article": "der"
  },
  {
    "noun": "Unternehmergeist",
    "article": "der"
  },
  {
    "noun": "Unternehmergewinn",
    "article": "der"
  },
  {
    "noun": "Unternehmerverband",
    "article": "der"
  },
  {
    "noun": "Unternehmungsgeist",
    "article": "der"
  },
  {
    "noun": "Unteroffizier",
    "article": "der"
  },
  {
    "noun": "Unteroffiziersrang",
    "article": "der"
  },
  {
    "noun": "Unterpunkt",
    "article": "der"
  },
  {
    "noun": "Unterricht",
    "article": "der"
  },
  {
    "noun": "Unterrichtsgegenstand",
    "article": "der"
  },
  {
    "noun": "Unterrichtsraum",
    "article": "der"
  },
  {
    "noun": "Unterrichtsstoff",
    "article": "der"
  },
  {
    "noun": "Unterrock",
    "article": "der"
  },
  {
    "noun": "Untersatz",
    "article": "der"
  },
  {
    "noun": "Unterschenkel",
    "article": "der"
  },
  {
    "noun": "Unterschied",
    "article": "der"
  },
  {
    "noun": "Unterschiedsbetrag",
    "article": "der"
  },
  {
    "noun": "Unterschlag",
    "article": "der"
  },
  {
    "noun": "Unterschlupf",
    "article": "der"
  },
  {
    "noun": "Unterschnitt",
    "article": "der"
  },
  {
    "noun": "Untersetzer",
    "article": "der"
  },
  {
    "noun": "Unterstand",
    "article": "der"
  },
  {
    "noun": "Unterstrich",
    "article": "der"
  },
  {
    "noun": "Untersuchungsausschuss",
    "article": "der"
  },
  {
    "noun": "Untersuchungsbericht",
    "article": "der"
  },
  {
    "noun": "Untersuchungsrichter",
    "article": "der"
  },
  {
    "noun": "Untertagebau",
    "article": "der"
  },
  {
    "noun": "Untertanengeist",
    "article": "der"
  },
  {
    "noun": "Unterteller",
    "article": "der"
  },
  {
    "noun": "Untertitel",
    "article": "der"
  },
  {
    "noun": "Unterton",
    "article": "der"
  },
  {
    "noun": "Unterwasserforscher",
    "article": "der"
  },
  {
    "noun": "Unterzeichner",
    "article": "der"
  },
  {
    "noun": "Unverstand",
    "article": "der"
  },
  {
    "noun": "Unwille",
    "article": "der"
  },
  {
    "noun": "Urahn",
    "article": "der"
  },
  {
    "noun": "Ural",
    "article": "der"
  },
  {
    "noun": "Uranbrenner",
    "article": "der"
  },
  {
    "noun": "Uranfang",
    "article": "der"
  },
  {
    "noun": "Uranglimmer",
    "article": "der"
  },
  {
    "noun": "Uranus",
    "article": "der"
  },
  {
    "noun": "Urbewohner",
    "article": "der"
  },
  {
    "noun": "Ureinwohner",
    "article": "der"
  },
  {
    "noun": "Urenkel",
    "article": "der"
  },
  {
    "noun": "Ureter",
    "article": "der"
  },
  {
    "noun": "Urgrund",
    "article": "der"
  },
  {
    "noun": "Urheber",
    "article": "der"
  },
  {
    "noun": "Urheberschutz",
    "article": "der"
  },
  {
    "noun": "Urin",
    "article": "der"
  },
  {
    "noun": "Urinstatus",
    "article": "der"
  },
  {
    "noun": "Urinstinkt",
    "article": "der"
  },
  {
    "noun": "Urknall",
    "article": "der"
  },
  {
    "noun": "Urkundenbeweis",
    "article": "der"
  },
  {
    "noun": "Urlaub",
    "article": "der"
  },
  {
    "noun": "Urlauber",
    "article": "der"
  },
  {
    "noun": "Urlaubsanspruch",
    "article": "der"
  },
  {
    "noun": "Urlaubsort",
    "article": "der"
  },
  {
    "noun": "Urlaubsplan",
    "article": "der"
  },
  {
    "noun": "Urlaubsschein",
    "article": "der"
  },
  {
    "noun": "Urlaubstag",
    "article": "der"
  },
  {
    "noun": "Urmensch",
    "article": "der"
  },
  {
    "noun": "Urnenfriedhof",
    "article": "der"
  },
  {
    "noun": "Urnengang",
    "article": "der"
  },
  {
    "noun": "Urologe",
    "article": "der"
  },
  {
    "noun": "Uropa",
    "article": "der"
  },
  {
    "noun": "Urquell",
    "article": "der"
  },
  {
    "noun": "Urschrei",
    "article": "der"
  },
  {
    "noun": "Ursprung",
    "article": "der"
  },
  {
    "noun": "Urstoff",
    "article": "der"
  },
  {
    "noun": "Urteilsspruch",
    "article": "der"
  },
  {
    "noun": "Urteilsvollzug",
    "article": "der"
  },
  {
    "noun": "Urtext",
    "article": "der"
  },
  {
    "noun": "Urtyp",
    "article": "der"
  },
  {
    "noun": "Urtypus",
    "article": "der"
  },
  {
    "noun": "Ururahn",
    "article": "der"
  },
  {
    "noun": "Ururenkel",
    "article": "der"
  },
  {
    "noun": "Urvater",
    "article": "der"
  },
  {
    "noun": "Urvogel",
    "article": "der"
  },
  {
    "noun": "Urwald",
    "article": "der"
  },
  {
    "noun": "Urzustand",
    "article": "der"
  },
  {
    "noun": "Usbeke",
    "article": "der"
  },
  {
    "noun": "Usurpator",
    "article": "der"
  },
  {
    "noun": "Usus",
    "article": "der"
  },
  {
    "noun": "Uterus",
    "article": "der"
  },
  {
    "noun": "Utilitarismus",
    "article": "der"
  },
  {
    "noun": "Utopismus",
    "article": "der"
  },
  {
    "noun": "Utopist",
    "article": "der"
  },
  {
    "noun": "Nachruf",
    "article": "der"
  },
  {
    "noun": "Nachruhm",
    "article": "der"
  },
  {
    "noun": "Nachsatz",
    "article": "der"
  },
  {
    "noun": "Nachschlag",
    "article": "der"
  },
  {
    "noun": "Nachschub",
    "article": "der"
  },
  {
    "noun": "Nachschuss",
    "article": "der"
  },
  {
    "noun": "Nachsichtwechsel",
    "article": "der"
  },
  {
    "noun": "Nachsommer",
    "article": "der"
  },
  {
    "noun": "Nachspann",
    "article": "der"
  },
  {
    "noun": "Nachtangriff",
    "article": "der"
  },
  {
    "noun": "Nachtbus",
    "article": "der"
  },
  {
    "noun": "Nachtclub",
    "article": "der"
  },
  {
    "noun": "Nachtdienst",
    "article": "der"
  },
  {
    "noun": "Nachteil",
    "article": "der"
  },
  {
    "noun": "Nachtfalter",
    "article": "der"
  },
  {
    "noun": "Nachtflug",
    "article": "der"
  },
  {
    "noun": "Nachtfrost",
    "article": "der"
  },
  {
    "noun": "Nachthimmel",
    "article": "der"
  },
  {
    "noun": "Nachtisch",
    "article": "der"
  },
  {
    "noun": "Nachtkasten",
    "article": "der"
  },
  {
    "noun": "Nachtmahr",
    "article": "der"
  },
  {
    "noun": "Nachtportier",
    "article": "der"
  },
  {
    "noun": "Nachtrag",
    "article": "der"
  },
  {
    "noun": "Nachtragshaushalt",
    "article": "der"
  },
  {
    "noun": "Nachtraubvogel",
    "article": "der"
  },
  {
    "noun": "Nachtschatten",
    "article": "der"
  },
  {
    "noun": "Nachtschlaf",
    "article": "der"
  },
  {
    "noun": "Nachtstrom",
    "article": "der"
  },
  {
    "noun": "Nachtstuhl",
    "article": "der"
  },
  {
    "noun": "Nachttarif",
    "article": "der"
  },
  {
    "noun": "Nachttisch",
    "article": "der"
  },
  {
    "noun": "Nachttopf",
    "article": "der"
  },
  {
    "noun": "Nachttresor",
    "article": "der"
  },
  {
    "noun": "Nachtvogel",
    "article": "der"
  },
  {
    "noun": "Nachtwandler",
    "article": "der"
  },
  {
    "noun": "Nachtzug",
    "article": "der"
  },
  {
    "noun": "Nachtzuschlag",
    "article": "der"
  },
  {
    "noun": "Nachweis",
    "article": "der"
  },
  {
    "noun": "Nachwinter",
    "article": "der"
  },
  {
    "noun": "Nachwuchs",
    "article": "der"
  },
  {
    "noun": "Nachwuchsmangel",
    "article": "der"
  },
  {
    "noun": "Nachwuchsspieler",
    "article": "der"
  },
  {
    "noun": "Nachzug",
    "article": "der"
  },
  {
    "noun": "Nackedei",
    "article": "der"
  },
  {
    "noun": "Nacken",
    "article": "der"
  },
  {
    "noun": "Nackenschutz",
    "article": "der"
  },
  {
    "noun": "Nackenwirbel",
    "article": "der"
  },
  {
    "noun": "Nadelbaum",
    "article": "der"
  },
  {
    "noun": "Nadeldrucker",
    "article": "der"
  },
  {
    "noun": "Nadelkopf",
    "article": "der"
  },
  {
    "noun": "Nadelstich",
    "article": "der"
  },
  {
    "noun": "Nadelstreifen",
    "article": "der"
  },
  {
    "noun": "Nadelstreifenanzug",
    "article": "der"
  },
  {
    "noun": "Nadelwald",
    "article": "der"
  },
  {
    "noun": "Nadir",
    "article": "der"
  },
  {
    "noun": "Nagel",
    "article": "der"
  },
  {
    "noun": "Nagelbohrer",
    "article": "der"
  },
  {
    "noun": "Nagelhautentferner",
    "article": "der"
  },
  {
    "noun": "Nagelkopf",
    "article": "der"
  },
  {
    "noun": "Nagellack",
    "article": "der"
  },
  {
    "noun": "Nagellackentferner",
    "article": "der"
  },
  {
    "noun": "Nagelschmied",
    "article": "der"
  },
  {
    "noun": "Nager",
    "article": "der"
  },
  {
    "noun": "Nagezahn",
    "article": "der"
  },
  {
    "noun": "Nahbereich",
    "article": "der"
  },
  {
    "noun": "Nahkampf",
    "article": "der"
  },
  {
    "noun": "Nahrungsstoff",
    "article": "der"
  },
  {
    "noun": "Nahrungstrieb",
    "article": "der"
  },
  {
    "noun": "Nahschuss",
    "article": "der"
  },
  {
    "noun": "Nahverkehr",
    "article": "der"
  },
  {
    "noun": "Nahverkehrszug",
    "article": "der"
  },
  {
    "noun": "Naivling",
    "article": "der"
  },
  {
    "noun": "Name",
    "article": "der"
  },
  {
    "noun": "Namensbruder",
    "article": "der"
  },
  {
    "noun": "Namensgeber",
    "article": "der"
  },
  {
    "noun": "Namenspatron",
    "article": "der"
  },
  {
    "noun": "Namenstag",
    "article": "der"
  },
  {
    "noun": "Namensvetter",
    "article": "der"
  },
  {
    "noun": "Namenszug",
    "article": "der"
  },
  {
    "noun": "Nandu",
    "article": "der"
  },
  {
    "noun": "Nanismus",
    "article": "der"
  },
  {
    "noun": "Napf",
    "article": "der"
  },
  {
    "noun": "Napfkuchen",
    "article": "der"
  },
  {
    "noun": "Narkosearzt",
    "article": "der"
  },
  {
    "noun": "Narkosefacharzt",
    "article": "der"
  },
  {
    "noun": "Narkotismus",
    "article": "der"
  },
  {
    "noun": "Narr",
    "article": "der"
  },
  {
    "noun": "Narrenstreich",
    "article": "der"
  },
  {
    "noun": "Narwal",
    "article": "der"
  },
  {
    "noun": "Narziss",
    "article": "der"
  },
  {
    "noun": "Narzissmus",
    "article": "der"
  },
  {
    "noun": "Nasalvokal",
    "article": "der"
  },
  {
    "noun": "Nascher",
    "article": "der"
  },
  {
    "noun": "Nasenbeinbruch",
    "article": "der"
  },
  {
    "noun": "Nasenpolyp",
    "article": "der"
  },
  {
    "noun": "Nasenring",
    "article": "der"
  },
  {
    "noun": "Nasenschleim",
    "article": "der"
  },
  {
    "noun": "Naseweis",
    "article": "der"
  },
  {
    "noun": "Nashornvogel",
    "article": "der"
  },
  {
    "noun": "Nationalcharakter",
    "article": "der"
  },
  {
    "noun": "Nationalfeiertag",
    "article": "der"
  },
  {
    "noun": "Nationalheld",
    "article": "der"
  },
  {
    "noun": "Nationalismus",
    "article": "der"
  },
  {
    "noun": "Nationalist",
    "article": "der"
  },
  {
    "noun": "Nationalpark",
    "article": "der"
  },
  {
    "noun": "Nationalpreis",
    "article": "der"
  },
  {
    "noun": "Nationalrat",
    "article": "der"
  },
  {
    "noun": "Nationalsozialismus",
    "article": "der"
  },
  {
    "noun": "Nationalsozialist",
    "article": "der"
  },
  {
    "noun": "Nationalspieler",
    "article": "der"
  },
  {
    "noun": "Nationalsport",
    "article": "der"
  },
  {
    "noun": "Nationalstaat",
    "article": "der"
  },
  {
    "noun": "Nationalstolz",
    "article": "der"
  },
  {
    "noun": "Nationaltanz",
    "article": "der"
  },
  {
    "noun": "Nationaltrainer",
    "article": "der"
  },
  {
    "noun": "Nativismus",
    "article": "der"
  },
  {
    "noun": "Nativist",
    "article": "der"
  },
  {
    "noun": "Natriumdampf",
    "article": "der"
  },
  {
    "noun": "Naturalbezug",
    "article": "der"
  },
  {
    "noun": "Naturalismus",
    "article": "der"
  },
  {
    "noun": "Naturalist",
    "article": "der"
  },
  {
    "noun": "Naturallohn",
    "article": "der"
  },
  {
    "noun": "Naturfilm",
    "article": "der"
  },
  {
    "noun": "Naturforscher",
    "article": "der"
  },
  {
    "noun": "Naturfreund",
    "article": "der"
  },
  {
    "noun": "Naturgummi",
    "article": "der"
  },
  {
    "noun": "Naturismus",
    "article": "der"
  },
  {
    "noun": "Naturist",
    "article": "der"
  },
  {
    "noun": "Naturkautschuk",
    "article": "der"
  },
  {
    "noun": "Naturlehrpfad",
    "article": "der"
  },
  {
    "noun": "Naturpark",
    "article": "der"
  },
  {
    "noun": "Naturschatz",
    "article": "der"
  },
  {
    "noun": "Naturschutz",
    "article": "der"
  },
  {
    "noun": "Naturschutzpark",
    "article": "der"
  },
  {
    "noun": "Natursekt",
    "article": "der"
  },
  {
    "noun": "Naturstein",
    "article": "der"
  },
  {
    "noun": "Naturstoff",
    "article": "der"
  },
  {
    "noun": "Naturton",
    "article": "der"
  },
  {
    "noun": "Naturtrieb",
    "article": "der"
  },
  {
    "noun": "Naturvorgang",
    "article": "der"
  },
  {
    "noun": "Naturwein",
    "article": "der"
  },
  {
    "noun": "Naturwissenschaftler",
    "article": "der"
  },
  {
    "noun": "Naturzustand",
    "article": "der"
  },
  {
    "noun": "Nauen",
    "article": "der"
  },
  {
    "noun": "Navigationsoffizier",
    "article": "der"
  },
  {
    "noun": "Navigationsraum",
    "article": "der"
  },
  {
    "noun": "Navigationssatellit",
    "article": "der"
  },
  {
    "noun": "Navigator",
    "article": "der"
  },
  {
    "noun": "Nazi",
    "article": "der"
  },
  {
    "noun": "Nazismus",
    "article": "der"
  },
  {
    "noun": "Nazist",
    "article": "der"
  },
  {
    "noun": "Neandertaler",
    "article": "der"
  },
  {
    "noun": "Nebel",
    "article": "der"
  },
  {
    "noun": "Nebelscheinwerfer",
    "article": "der"
  },
  {
    "noun": "Nebenanschluss",
    "article": "der"
  },
  {
    "noun": "Nebenausgang",
    "article": "der"
  },
  {
    "noun": "Nebenberuf",
    "article": "der"
  },
  {
    "noun": "Nebenbuhler",
    "article": "der"
  },
  {
    "noun": "Nebeneffekt",
    "article": "der"
  },
  {
    "noun": "Nebeneingang",
    "article": "der"
  },
  {
    "noun": "Nebenerwerb",
    "article": "der"
  },
  {
    "noun": "Nebenfluss",
    "article": "der"
  },
  {
    "noun": "Nebengedanke",
    "article": "der"
  },
  {
    "noun": "Nebengeschmack",
    "article": "der"
  },
  {
    "noun": "Nebenjob",
    "article": "der"
  },
  {
    "noun": "Nebenmensch",
    "article": "der"
  },
  {
    "noun": "Nebenpunkt",
    "article": "der"
  },
  {
    "noun": "Nebenraum",
    "article": "der"
  },
  {
    "noun": "Nebensatz",
    "article": "der"
  },
  {
    "noun": "Nebentisch",
    "article": "der"
  },
  {
    "noun": "Nebenton",
    "article": "der"
  },
  {
    "noun": "Nebenumstand",
    "article": "der"
  },
  {
    "noun": "Nebenverdienst",
    "article": "der"
  },
  {
    "noun": "Nebenweg",
    "article": "der"
  },
  {
    "noun": "Nebenwinkel",
    "article": "der"
  },
  {
    "noun": "Nebenzweck",
    "article": "der"
  },
  {
    "noun": "Nebenzweig",
    "article": "der"
  },
  {
    "noun": "Neck",
    "article": "der"
  },
  {
    "noun": "Neerstrom",
    "article": "der"
  },
  {
    "noun": "Neffe",
    "article": "der"
  },
  {
    "noun": "Negativfilm",
    "article": "der"
  },
  {
    "noun": "Negativzins",
    "article": "der"
  },
  {
    "noun": "Negator",
    "article": "der"
  },
  {
    "noun": "Nehmer",
    "article": "der"
  },
  {
    "noun": "Neid",
    "article": "der"
  },
  {
    "noun": "Neider",
    "article": "der"
  },
  {
    "noun": "Neigungsmesser",
    "article": "der"
  },
  {
    "noun": "Neigungswinkel",
    "article": "der"
  },
  {
    "noun": "Neinsager",
    "article": "der"
  },
  {
    "noun": "Nekrolog",
    "article": "der"
  },
  {
    "noun": "Nektar",
    "article": "der"
  },
  {
    "noun": "Nelkenpfeffer",
    "article": "der"
  },
  {
    "noun": "Nennbetrag",
    "article": "der"
  },
  {
    "noun": "Nenner",
    "article": "der"
  },
  {
    "noun": "Nennformsatz",
    "article": "der"
  },
  {
    "noun": "Nennonkel",
    "article": "der"
  },
  {
    "noun": "Nennungsschluss",
    "article": "der"
  },
  {
    "noun": "Nennwert",
    "article": "der"
  },
  {
    "noun": "Neofaschismus",
    "article": "der"
  },
  {
    "noun": "Neofaschist",
    "article": "der"
  },
  {
    "noun": "Neoimpressionismus",
    "article": "der"
  },
  {
    "noun": "Neoklassizismus",
    "article": "der"
  },
  {
    "noun": "Neokolonialismus",
    "article": "der"
  },
  {
    "noun": "Neoliberalismus",
    "article": "der"
  },
  {
    "noun": "Neologismus",
    "article": "der"
  },
  {
    "noun": "Neomarxismus",
    "article": "der"
  },
  {
    "noun": "Neonazi",
    "article": "der"
  },
  {
    "noun": "Neonazismus",
    "article": "der"
  },
  {
    "noun": "Neonazist",
    "article": "der"
  },
  {
    "noun": "Neophyt",
    "article": "der"
  },
  {
    "noun": "Neoprenanzug",
    "article": "der"
  },
  {
    "noun": "Nepaler",
    "article": "der"
  },
  {
    "noun": "Nepalese",
    "article": "der"
  },
  {
    "noun": "Nephrolith",
    "article": "der"
  },
  {
    "noun": "Nepote",
    "article": "der"
  },
  {
    "noun": "Nepotismus",
    "article": "der"
  },
  {
    "noun": "Nepp",
    "article": "der"
  },
  {
    "noun": "Nepper",
    "article": "der"
  },
  {
    "noun": "Neptun",
    "article": "der"
  },
  {
    "noun": "Nerfling",
    "article": "der"
  },
  {
    "noun": "Nerv",
    "article": "der"
  },
  {
    "noun": "Nervenarzt",
    "article": "der"
  },
  {
    "noun": "Nervenkitzel",
    "article": "der"
  },
  {
    "noun": "Nervenknoten",
    "article": "der"
  },
  {
    "noun": "Nervenkrieg",
    "article": "der"
  },
  {
    "noun": "Nervenschmerz",
    "article": "der"
  },
  {
    "noun": "Nervenschock",
    "article": "der"
  },
  {
    "noun": "Nervenzusammenbruch",
    "article": "der"
  },
  {
    "noun": "Nerz",
    "article": "der"
  },
  {
    "noun": "Nerzmantel",
    "article": "der"
  },
  {
    "noun": "Nessel",
    "article": "der"
  },
  {
    "noun": "Nesselausschlag",
    "article": "der"
  },
  {
    "noun": "Nesselfaden",
    "article": "der"
  },
  {
    "noun": "Nestbau",
    "article": "der"
  },
  {
    "noun": "Nesthocker",
    "article": "der"
  },
  {
    "noun": "Nestling",
    "article": "der"
  },
  {
    "noun": "Nestor",
    "article": "der"
  },
  {
    "noun": "Nettobetrag",
    "article": "der"
  },
  {
    "noun": "Nettoertrag",
    "article": "der"
  },
  {
    "noun": "Nettogewinn",
    "article": "der"
  },
  {
    "noun": "Nettolohn",
    "article": "der"
  },
  {
    "noun": "Nettopreis",
    "article": "der"
  },
  {
    "noun": "Nettoverdienst",
    "article": "der"
  },
  {
    "noun": "Nettowert",
    "article": "der"
  },
  {
    "noun": "Nettozahler",
    "article": "der"
  },
  {
    "noun": "Netzanschluss",
    "article": "der"
  },
  {
    "noun": "Netzball",
    "article": "der"
  },
  {
    "noun": "Netzbetreiber",
    "article": "der"
  },
  {
    "noun": "Netzbetrieb",
    "article": "der"
  },
  {
    "noun": "Netzgleichrichter",
    "article": "der"
  },
  {
    "noun": "Netzmagen",
    "article": "der"
  },
  {
    "noun": "Netzplan",
    "article": "der"
  },
  {
    "noun": "Netzstecker",
    "article": "der"
  },
  {
    "noun": "Netzstrumpf",
    "article": "der"
  },
  {
    "noun": "Neuanfang",
    "article": "der"
  },
  {
    "noun": "Neubau",
    "article": "der"
  },
  {
    "noun": "Neubeginn",
    "article": "der"
  },
  {
    "noun": "Neubruch",
    "article": "der"
  },
  {
    "noun": "Neudruck",
    "article": "der"
  },
  {
    "noun": "Neuerer",
    "article": "der"
  },
  {
    "noun": "Neugrad",
    "article": "der"
  },
  {
    "noun": "Neukauf",
    "article": "der"
  },
  {
    "noun": "Neukunde",
    "article": "der"
  },
  {
    "noun": "Neuling",
    "article": "der"
  },
  {
    "noun": "Neumond",
    "article": "der"
  },
  {
    "noun": "Neuner",
    "article": "der"
  },
  {
    "noun": "Neuplatonismus",
    "article": "der"
  },
  {
    "noun": "Neupreis",
    "article": "der"
  },
  {
    "noun": "Neurastheniker",
    "article": "der"
  },
  {
    "noun": "Neurochirurg",
    "article": "der"
  },
  {
    "noun": "Neurologe",
    "article": "der"
  },
  {
    "noun": "Neuromantiker",
    "article": "der"
  },
  {
    "noun": "Neurotiker",
    "article": "der"
  },
  {
    "noun": "Neurotransmitter",
    "article": "der"
  },
  {
    "noun": "Neuschnee",
    "article": "der"
  },
  {
    "noun": "Neustart",
    "article": "der"
  },
  {
    "noun": "Neutralismus",
    "article": "der"
  },
  {
    "noun": "Neutralist",
    "article": "der"
  },
  {
    "noun": "Neuwagen",
    "article": "der"
  },
  {
    "noun": "Neuwert",
    "article": "der"
  },
  {
    "noun": "Neuzustand",
    "article": "der"
  },
  {
    "noun": "Newcomer",
    "article": "der"
  },
  {
    "noun": "Nicaraguaner",
    "article": "der"
  },
  {
    "noun": "Nichtangriffspakt",
    "article": "der"
  },
  {
    "noun": "Nichterfolg",
    "article": "der"
  },
  {
    "noun": "Nichtfachmann",
    "article": "der"
  },
  {
    "noun": "Nichtgebrauch",
    "article": "der"
  },
  {
    "noun": "Nichtleiter",
    "article": "der"
  },
  {
    "noun": "Nichtraucher",
    "article": "der"
  },
  {
    "noun": "Nichtschwimmer",
    "article": "der"
  },
  {
    "noun": "Nichtsnutz",
    "article": "der"
  },
  {
    "noun": "Nichtstuer",
    "article": "der"
  },
  {
    "noun": "Nichtswisser",
    "article": "der"
  },
  {
    "noun": "Nichttrinker",
    "article": "der"
  },
  {
    "noun": "Nicker",
    "article": "der"
  },
  {
    "noun": "Niederdruck",
    "article": "der"
  },
  {
    "noun": "Niederdruckreifen",
    "article": "der"
  },
  {
    "noun": "Niedergang",
    "article": "der"
  },
  {
    "noun": "Niederrhein",
    "article": "der"
  },
  {
    "noun": "Niederschlag",
    "article": "der"
  },
  {
    "noun": "Niederschlagsmesser",
    "article": "der"
  },
  {
    "noun": "Niedriglohn",
    "article": "der"
  },
  {
    "noun": "Nierenbaum",
    "article": "der"
  },
  {
    "noun": "Nierenbraten",
    "article": "der"
  },
  {
    "noun": "Nierenstein",
    "article": "der"
  },
  {
    "noun": "Nieselregen",
    "article": "der"
  },
  {
    "noun": "Nieser",
    "article": "der"
  },
  {
    "noun": "Niethammer",
    "article": "der"
  },
  {
    "noun": "Nietkopf",
    "article": "der"
  },
  {
    "noun": "Nietnagel",
    "article": "der"
  },
  {
    "noun": "Niger",
    "article": "der"
  },
  {
    "noun": "Nigerianer",
    "article": "der"
  },
  {
    "noun": "Nightclub",
    "article": "der"
  },
  {
    "noun": "Nigrer",
    "article": "der"
  },
  {
    "noun": "Nihilismus",
    "article": "der"
  },
  {
    "noun": "Nihilist",
    "article": "der"
  },
  {
    "noun": "Nikolaus",
    "article": "der"
  },
  {
    "noun": "Nikolaustag",
    "article": "der"
  },
  {
    "noun": "Nil",
    "article": "der"
  },
  {
    "noun": "Nimbus",
    "article": "der"
  },
  {
    "noun": "Nimmersatt",
    "article": "der"
  },
  {
    "noun": "Nippel",
    "article": "der"
  },
  {
    "noun": "Nistkasten",
    "article": "der"
  },
  {
    "noun": "Nistplatz",
    "article": "der"
  },
  {
    "noun": "Niveauunterschied",
    "article": "der"
  },
  {
    "noun": "Nix",
    "article": "der"
  },
  {
    "noun": "Nobelpreis",
    "article": "der"
  },
  {
    "noun": "Nockenschalter",
    "article": "der"
  },
  {
    "noun": "Nomade",
    "article": "der"
  },
  {
    "noun": "Nomadenstamm",
    "article": "der"
  },
  {
    "noun": "Nominalbetrag",
    "article": "der"
  },
  {
    "noun": "Nominallohn",
    "article": "der"
  },
  {
    "noun": "Nominalwert",
    "article": "der"
  },
  {
    "noun": "Nominalzins",
    "article": "der"
  },
  {
    "noun": "Nominativ",
    "article": "der"
  },
  {
    "noun": "Nonius",
    "article": "der"
  },
  {
    "noun": "Nonkonformismus",
    "article": "der"
  },
  {
    "noun": "Nonkonformist",
    "article": "der"
  },
  {
    "noun": "Nonnenorden",
    "article": "der"
  },
  {
    "noun": "Nonsens",
    "article": "der"
  },
  {
    "noun": "Nordatlantikpakt",
    "article": "der"
  },
  {
    "noun": "Nordist",
    "article": "der"
  },
  {
    "noun": "Nordosten",
    "article": "der"
  },
  {
    "noun": "Nordostwind",
    "article": "der"
  },
  {
    "noun": "Nordpol",
    "article": "der"
  },
  {
    "noun": "Nordrand",
    "article": "der"
  },
  {
    "noun": "Nordstern",
    "article": "der"
  },
  {
    "noun": "Nordwesten",
    "article": "der"
  },
  {
    "noun": "Nordwestwind",
    "article": "der"
  },
  {
    "noun": "Nordwind",
    "article": "der"
  },
  {
    "noun": "Normaldruck",
    "article": "der"
  },
  {
    "noun": "Normalfall",
    "article": "der"
  },
  {
    "noun": "Normalfilm",
    "article": "der"
  },
  {
    "noun": "Normalverbraucher",
    "article": "der"
  },
  {
    "noun": "Normalverdiener",
    "article": "der"
  },
  {
    "noun": "Normalwert",
    "article": "der"
  },
  {
    "noun": "Normalzustand",
    "article": "der"
  },
  {
    "noun": "Normanne",
    "article": "der"
  },
  {
    "noun": "Normdruck",
    "article": "der"
  },
  {
    "noun": "Normenausschuss",
    "article": "der"
  },
  {
    "noun": "Notanker",
    "article": "der"
  },
  {
    "noun": "Notar",
    "article": "der"
  },
  {
    "noun": "Notarzt",
    "article": "der"
  },
  {
    "noun": "Notarztwagen",
    "article": "der"
  },
  {
    "noun": "Notausgang",
    "article": "der"
  },
  {
    "noun": "Notausstieg",
    "article": "der"
  },
  {
    "noun": "Notbehelf",
    "article": "der"
  },
  {
    "noun": "Notbetrieb",
    "article": "der"
  },
  {
    "noun": "Notdienst",
    "article": "der"
  },
  {
    "noun": "Notenaustausch",
    "article": "der"
  },
  {
    "noun": "Notendruck",
    "article": "der"
  },
  {
    "noun": "Notendurchschnitt",
    "article": "der"
  },
  {
    "noun": "Notfall",
    "article": "der"
  },
  {
    "noun": "Notfalldienst",
    "article": "der"
  },
  {
    "noun": "Notgroschen",
    "article": "der"
  },
  {
    "noun": "Nothafen",
    "article": "der"
  },
  {
    "noun": "Notizblock",
    "article": "der"
  },
  {
    "noun": "Notizzettel",
    "article": "der"
  },
  {
    "noun": "Notruf",
    "article": "der"
  },
  {
    "noun": "Notschalter",
    "article": "der"
  },
  {
    "noun": "Notsitz",
    "article": "der"
  },
  {
    "noun": "Notstand",
    "article": "der"
  },
  {
    "noun": "Notverkauf",
    "article": "der"
  },
  {
    "noun": "Novellist",
    "article": "der"
  },
  {
    "noun": "November",
    "article": "der"
  },
  {
    "noun": "Nuckel",
    "article": "der"
  },
  {
    "noun": "Nudelsalat",
    "article": "der"
  },
  {
    "noun": "Nudismus",
    "article": "der"
  },
  {
    "noun": "Nudist",
    "article": "der"
  },
  {
    "noun": "Nuklearsprengkopf",
    "article": "der"
  },
  {
    "noun": "Nukleus",
    "article": "der"
  },
  {
    "noun": "Nullmeridian",
    "article": "der"
  },
  {
    "noun": "Nullpunkt",
    "article": "der"
  },
  {
    "noun": "Nulltarif",
    "article": "der"
  },
  {
    "noun": "Nullwert",
    "article": "der"
  },
  {
    "noun": "Numerus",
    "article": "der"
  },
  {
    "noun": "Numismatiker",
    "article": "der"
  },
  {
    "noun": "Nuntius",
    "article": "der"
  },
  {
    "noun": "Nussbaum",
    "article": "der"
  },
  {
    "noun": "Nusskern",
    "article": "der"
  },
  {
    "noun": "Nussknacker",
    "article": "der"
  },
  {
    "noun": "Nuthobel",
    "article": "der"
  },
  {
    "noun": "Nutria",
    "article": "der"
  },
  {
    "noun": "Nutzbau",
    "article": "der"
  },
  {
    "noun": "Nutzeffekt",
    "article": "der"
  },
  {
    "noun": "Nutzer",
    "article": "der"
  },
  {
    "noun": "Nutzgarten",
    "article": "der"
  },
  {
    "noun": "Nutzungsgrad",
    "article": "der"
  },
  {
    "noun": "Nutzungsvertrag",
    "article": "der"
  },
  {
    "noun": "Nutzwert",
    "article": "der"
  },
  {
    "noun": "Nylonstrumpf",
    "article": "der"
  },
  {
    "noun": "Nymphensittich",
    "article": "der"
  },
  {
    "noun": "Nystagmus",
    "article": "der"
  },
  {
    "noun": "Zahnkranz",
    "article": "der"
  },
  {
    "noun": "Zahnlaut",
    "article": "der"
  },
  {
    "noun": "Zahnmediziner",
    "article": "der"
  },
  {
    "noun": "Zahnriemen",
    "article": "der"
  },
  {
    "noun": "Zahnschmelz",
    "article": "der"
  },
  {
    "noun": "Zahnschmerz",
    "article": "der"
  },
  {
    "noun": "Zahnspiegel",
    "article": "der"
  },
  {
    "noun": "Zahnstein",
    "article": "der"
  },
  {
    "noun": "Zahnstocher",
    "article": "der"
  },
  {
    "noun": "Zahnstumpf",
    "article": "der"
  },
  {
    "noun": "Zahntechniker",
    "article": "der"
  },
  {
    "noun": "Zahnverlust",
    "article": "der"
  },
  {
    "noun": "Zahnwal",
    "article": "der"
  },
  {
    "noun": "Zahnwechsel",
    "article": "der"
  },
  {
    "noun": "Zahnzement",
    "article": "der"
  },
  {
    "noun": "Zain",
    "article": "der"
  },
  {
    "noun": "Zander",
    "article": "der"
  },
  {
    "noun": "Zangengriff",
    "article": "der"
  },
  {
    "noun": "Zank",
    "article": "der"
  },
  {
    "noun": "Zankapfel",
    "article": "der"
  },
  {
    "noun": "Zankteufel",
    "article": "der"
  },
  {
    "noun": "Zapf",
    "article": "der"
  },
  {
    "noun": "Zapfenstreich",
    "article": "der"
  },
  {
    "noun": "Zapfenzieher",
    "article": "der"
  },
  {
    "noun": "Zapfer",
    "article": "der"
  },
  {
    "noun": "Zapfhahn",
    "article": "der"
  },
  {
    "noun": "Zaponlack",
    "article": "der"
  },
  {
    "noun": "Zar",
    "article": "der"
  },
  {
    "noun": "Zaster",
    "article": "der"
  },
  {
    "noun": "Zauber",
    "article": "der"
  },
  {
    "noun": "Zauberer",
    "article": "der"
  },
  {
    "noun": "Zauberspiegel",
    "article": "der"
  },
  {
    "noun": "Zauberspruch",
    "article": "der"
  },
  {
    "noun": "Zauberstab",
    "article": "der"
  },
  {
    "noun": "Zaubertrank",
    "article": "der"
  },
  {
    "noun": "Zaubertrick",
    "article": "der"
  },
  {
    "noun": "Zauberwald",
    "article": "der"
  },
  {
    "noun": "Zauderer",
    "article": "der"
  },
  {
    "noun": "Zaum",
    "article": "der"
  },
  {
    "noun": "Zaun",
    "article": "der"
  },
  {
    "noun": "Zaungast",
    "article": "der"
  },
  {
    "noun": "Zaunpfahl",
    "article": "der"
  },
  {
    "noun": "Zaunpfosten",
    "article": "der"
  },
  {
    "noun": "Zaunspfahl",
    "article": "der"
  },
  {
    "noun": "Zebrastreifen",
    "article": "der"
  },
  {
    "noun": "Zechbruder",
    "article": "der"
  },
  {
    "noun": "Zecher",
    "article": "der"
  },
  {
    "noun": "Zechkumpan",
    "article": "der"
  },
  {
    "noun": "Zechpreller",
    "article": "der"
  },
  {
    "noun": "Zeckenbiss",
    "article": "der"
  },
  {
    "noun": "Zedent",
    "article": "der"
  },
  {
    "noun": "Zeh",
    "article": "der"
  },
  {
    "noun": "Zehennagel",
    "article": "der"
  },
  {
    "noun": "Zehner",
    "article": "der"
  },
  {
    "noun": "Zehnjahresplan",
    "article": "der"
  },
  {
    "noun": "Zehnkampf",
    "article": "der"
  },
  {
    "noun": "Zehnmarkschein",
    "article": "der"
  },
  {
    "noun": "Zeichenblock",
    "article": "der"
  },
  {
    "noun": "Zeichencode",
    "article": "der"
  },
  {
    "noun": "Zeichenfilm",
    "article": "der"
  },
  {
    "noun": "Zeichengenerator",
    "article": "der"
  },
  {
    "noun": "Zeichenkarton",
    "article": "der"
  },
  {
    "noun": "Zeichenlehrer",
    "article": "der"
  },
  {
    "noun": "Zeichenleser",
    "article": "der"
  },
  {
    "noun": "Zeichensaal",
    "article": "der"
  },
  {
    "noun": "Zeichensatz",
    "article": "der"
  },
  {
    "noun": "Zeichenstift",
    "article": "der"
  },
  {
    "noun": "Zeichentisch",
    "article": "der"
  },
  {
    "noun": "Zeichentrickfilm",
    "article": "der"
  },
  {
    "noun": "Zeichenunterricht",
    "article": "der"
  },
  {
    "noun": "Zeichenwinkel",
    "article": "der"
  },
  {
    "noun": "Zeichner",
    "article": "der"
  },
  {
    "noun": "Zeidler",
    "article": "der"
  },
  {
    "noun": "Zeigefinger",
    "article": "der"
  },
  {
    "noun": "Zeiger",
    "article": "der"
  },
  {
    "noun": "Zeigestock",
    "article": "der"
  },
  {
    "noun": "Zeilenabstand",
    "article": "der"
  },
  {
    "noun": "Zeilenanfang",
    "article": "der"
  },
  {
    "noun": "Zeilendrucker",
    "article": "der"
  },
  {
    "noun": "Zeilendurchschuss",
    "article": "der"
  },
  {
    "noun": "Zeilenhonorar",
    "article": "der"
  },
  {
    "noun": "Zeilensprung",
    "article": "der"
  },
  {
    "noun": "Zeilenumbruch",
    "article": "der"
  },
  {
    "noun": "Zeisig",
    "article": "der"
  },
  {
    "noun": "Zeising",
    "article": "der"
  },
  {
    "noun": "Zeitablauf",
    "article": "der"
  },
  {
    "noun": "Zeitabschnitt",
    "article": "der"
  },
  {
    "noun": "Zeitabstand",
    "article": "der"
  },
  {
    "noun": "Zeitarbeiter",
    "article": "der"
  },
  {
    "noun": "Zeitaufwand",
    "article": "der"
  },
  {
    "noun": "Zeitausgleich",
    "article": "der"
  },
  {
    "noun": "Zeitbedarf",
    "article": "der"
  },
  {
    "noun": "Zeitbegriff",
    "article": "der"
  },
  {
    "noun": "Zeitdruck",
    "article": "der"
  },
  {
    "noun": "Zeitgeist",
    "article": "der"
  },
  {
    "noun": "Zeitgenosse",
    "article": "der"
  },
  {
    "noun": "Zeitgeschmack",
    "article": "der"
  },
  {
    "noun": "Zeitgewinn",
    "article": "der"
  },
  {
    "noun": "Zeithorizont",
    "article": "der"
  },
  {
    "noun": "Zeitlauf",
    "article": "der"
  },
  {
    "noun": "Zeitlohn",
    "article": "der"
  },
  {
    "noun": "Zeitmangel",
    "article": "der"
  },
  {
    "noun": "Zeitmesser",
    "article": "der"
  },
  {
    "noun": "Zeitnehmer",
    "article": "der"
  },
  {
    "noun": "Zeitplan",
    "article": "der"
  },
  {
    "noun": "Zeitpunkt",
    "article": "der"
  },
  {
    "noun": "Zeitraffer",
    "article": "der"
  },
  {
    "noun": "Zeitrahmen",
    "article": "der"
  },
  {
    "noun": "Zeitraum",
    "article": "der"
  },
  {
    "noun": "Zeitroman",
    "article": "der"
  },
  {
    "noun": "Zeitschalter",
    "article": "der"
  },
  {
    "noun": "Zeitschriftenverlag",
    "article": "der"
  },
  {
    "noun": "Zeitsinn",
    "article": "der"
  },
  {
    "noun": "Zeitsoldat",
    "article": "der"
  },
  {
    "noun": "Zeitungsartikel",
    "article": "der"
  },
  {
    "noun": "Zeitungsausschnitt",
    "article": "der"
  },
  {
    "noun": "Zeitungsbericht",
    "article": "der"
  },
  {
    "noun": "Zeitungsjunge",
    "article": "der"
  },
  {
    "noun": "Zeitungskiosk",
    "article": "der"
  },
  {
    "noun": "Zeitungskorrespondent",
    "article": "der"
  },
  {
    "noun": "Zeitungsleser",
    "article": "der"
  },
  {
    "noun": "Zeitungsredakteur",
    "article": "der"
  },
  {
    "noun": "Zeitungsstand",
    "article": "der"
  },
  {
    "noun": "Zeitungsstil",
    "article": "der"
  },
  {
    "noun": "Zeitungsverlag",
    "article": "der"
  },
  {
    "noun": "Zeitungsverleger",
    "article": "der"
  },
  {
    "noun": "Zeitunterschied",
    "article": "der"
  },
  {
    "noun": "Zeitverlauf",
    "article": "der"
  },
  {
    "noun": "Zeitverlust",
    "article": "der"
  },
  {
    "noun": "Zeitvertreib",
    "article": "der"
  },
  {
    "noun": "Zeitwert",
    "article": "der"
  },
  {
    "noun": "Zeitzeuge",
    "article": "der"
  },
  {
    "noun": "Zelebrant",
    "article": "der"
  },
  {
    "noun": "Zellengenosse",
    "article": "der"
  },
  {
    "noun": "Zellkern",
    "article": "der"
  },
  {
    "noun": "Zellstoff",
    "article": "der"
  },
  {
    "noun": "Zelot",
    "article": "der"
  },
  {
    "noun": "Zelotismus",
    "article": "der"
  },
  {
    "noun": "Zelter",
    "article": "der"
  },
  {
    "noun": "Zeltpflock",
    "article": "der"
  },
  {
    "noun": "Zeltplatz",
    "article": "der"
  },
  {
    "noun": "Zeltstock",
    "article": "der"
  },
  {
    "noun": "Zement",
    "article": "der"
  },
  {
    "noun": "Zementit",
    "article": "der"
  },
  {
    "noun": "Zementsack",
    "article": "der"
  },
  {
    "noun": "Zementsilo",
    "article": "der"
  },
  {
    "noun": "Zenit",
    "article": "der"
  },
  {
    "noun": "Zensor",
    "article": "der"
  },
  {
    "noun": "Zensus",
    "article": "der"
  },
  {
    "noun": "Zentaur",
    "article": "der"
  },
  {
    "noun": "Zentenar",
    "article": "der"
  },
  {
    "noun": "Zentner",
    "article": "der"
  },
  {
    "noun": "Zentralbankrat",
    "article": "der"
  },
  {
    "noun": "Zentralismus",
    "article": "der"
  },
  {
    "noun": "Zentralkatalog",
    "article": "der"
  },
  {
    "noun": "Zentralspeicher",
    "article": "der"
  },
  {
    "noun": "Zentralwert",
    "article": "der"
  },
  {
    "noun": "Zentriwinkel",
    "article": "der"
  },
  {
    "noun": "Zenturio",
    "article": "der"
  },
  {
    "noun": "Zeolith",
    "article": "der"
  },
  {
    "noun": "Zephir",
    "article": "der"
  },
  {
    "noun": "Zeppelin",
    "article": "der"
  },
  {
    "noun": "Zeremonienmeister",
    "article": "der"
  },
  {
    "noun": "Zerfall",
    "article": "der"
  },
  {
    "noun": "Zerfallsprozess",
    "article": "der"
  },
  {
    "noun": "Zerhacker",
    "article": "der"
  },
  {
    "noun": "Zerkleinerer",
    "article": "der"
  },
  {
    "noun": "Zerobond",
    "article": "der"
  },
  {
    "noun": "Zerrspiegel",
    "article": "der"
  },
  {
    "noun": "Zersetzungsprozess",
    "article": "der"
  },
  {
    "noun": "Zessionar",
    "article": "der"
  },
  {
    "noun": "Zettel",
    "article": "der"
  },
  {
    "noun": "Zettelkasten",
    "article": "der"
  },
  {
    "noun": "Zeugdruck",
    "article": "der"
  },
  {
    "noun": "Zeugenstand",
    "article": "der"
  },
  {
    "noun": "Zibet",
    "article": "der"
  },
  {
    "noun": "Zichorienkaffee",
    "article": "der"
  },
  {
    "noun": "Zickzack",
    "article": "der"
  },
  {
    "noun": "Zickzackkurs",
    "article": "der"
  },
  {
    "noun": "Zider",
    "article": "der"
  },
  {
    "noun": "Ziegel",
    "article": "der"
  },
  {
    "noun": "Ziegelbrenner",
    "article": "der"
  },
  {
    "noun": "Ziegelofen",
    "article": "der"
  },
  {
    "noun": "Ziegelstein",
    "article": "der"
  },
  {
    "noun": "Ziegenbart",
    "article": "der"
  },
  {
    "noun": "Ziegenbock",
    "article": "der"
  },
  {
    "noun": "Ziegenhirt",
    "article": "der"
  },
  {
    "noun": "Ziegenmelker",
    "article": "der"
  },
  {
    "noun": "Ziegler",
    "article": "der"
  },
  {
    "noun": "Ziehbrunnen",
    "article": "der"
  },
  {
    "noun": "Zieher",
    "article": "der"
  },
  {
    "noun": "Zielbahnhof",
    "article": "der"
  },
  {
    "noun": "Zielhafen",
    "article": "der"
  },
  {
    "noun": "Zielkonflikt",
    "article": "der"
  },
  {
    "noun": "Zielkorridor",
    "article": "der"
  },
  {
    "noun": "Zielort",
    "article": "der"
  },
  {
    "noun": "Zielpunkt",
    "article": "der"
  },
  {
    "noun": "Zielrichter",
    "article": "der"
  },
  {
    "noun": "Ziemer",
    "article": "der"
  },
  {
    "noun": "Zierat",
    "article": "der"
  },
  {
    "noun": "Zierfisch",
    "article": "der"
  },
  {
    "noun": "Ziergarten",
    "article": "der"
  },
  {
    "noun": "Zierstab",
    "article": "der"
  },
  {
    "noun": "Zierstich",
    "article": "der"
  },
  {
    "noun": "Zierstrauch",
    "article": "der"
  },
  {
    "noun": "Zierstreifen",
    "article": "der"
  },
  {
    "noun": "Zigarettenautomat",
    "article": "der"
  },
  {
    "noun": "Zigarettenrauch",
    "article": "der"
  },
  {
    "noun": "Zigarettenraucher",
    "article": "der"
  },
  {
    "noun": "Zigarettenstummel",
    "article": "der"
  },
  {
    "noun": "Zigarettentabak",
    "article": "der"
  },
  {
    "noun": "Zigarrenabschneider",
    "article": "der"
  },
  {
    "noun": "Zigarrenstummel",
    "article": "der"
  },
  {
    "noun": "Zigeuner",
    "article": "der"
  },
  {
    "noun": "Zimmerarrest",
    "article": "der"
  },
  {
    "noun": "Zimmerbrand",
    "article": "der"
  },
  {
    "noun": "Zimmerer",
    "article": "der"
  },
  {
    "noun": "Zimmergeselle",
    "article": "der"
  },
  {
    "noun": "Zimmerherr",
    "article": "der"
  },
  {
    "noun": "Zimmerkellner",
    "article": "der"
  },
  {
    "noun": "Zimmermann",
    "article": "der"
  },
  {
    "noun": "Zimmermeister",
    "article": "der"
  },
  {
    "noun": "Zimt",
    "article": "der"
  },
  {
    "noun": "Zimtapfel",
    "article": "der"
  },
  {
    "noun": "Zimtbaum",
    "article": "der"
  },
  {
    "noun": "Zinkdruck",
    "article": "der"
  },
  {
    "noun": "Zinkenist",
    "article": "der"
  },
  {
    "noun": "Zinker",
    "article": "der"
  },
  {
    "noun": "Zinksarg",
    "article": "der"
  },
  {
    "noun": "Zinkspat",
    "article": "der"
  },
  {
    "noun": "Zinnkrug",
    "article": "der"
  },
  {
    "noun": "Zinnober",
    "article": "der"
  },
  {
    "noun": "Zinnsoldat",
    "article": "der"
  },
  {
    "noun": "Zinnstein",
    "article": "der"
  },
  {
    "noun": "Zinnteller",
    "article": "der"
  },
  {
    "noun": "Zinnwaldit",
    "article": "der"
  },
  {
    "noun": "Zins",
    "article": "der"
  },
  {
    "noun": "Zinsabschlag",
    "article": "der"
  },
  {
    "noun": "Zinsabschnitt",
    "article": "der"
  },
  {
    "noun": "Zinsertrag",
    "article": "der"
  },
  {
    "noun": "Zinseszins",
    "article": "der"
  },
  {
    "noun": "Zinssatz",
    "article": "der"
  },
  {
    "noun": "Zinstermin",
    "article": "der"
  },
  {
    "noun": "Zinswucher",
    "article": "der"
  },
  {
    "noun": "Zionismus",
    "article": "der"
  },
  {
    "noun": "Zionist",
    "article": "der"
  },
  {
    "noun": "Zipfel",
    "article": "der"
  },
  {
    "noun": "Zirkel",
    "article": "der"
  },
  {
    "noun": "Zirkelbeweis",
    "article": "der"
  },
  {
    "noun": "Zirkelkasten",
    "article": "der"
  },
  {
    "noun": "Zirkelschluss",
    "article": "der"
  },
  {
    "noun": "Zirkon",
    "article": "der"
  },
  {
    "noun": "Zirkumflex",
    "article": "der"
  },
  {
    "noun": "Zirrostratus",
    "article": "der"
  },
  {
    "noun": "Zischlaut",
    "article": "der"
  },
  {
    "noun": "Ziseleur",
    "article": "der"
  },
  {
    "noun": "Ziselierer",
    "article": "der"
  },
  {
    "noun": "Zisterzienser",
    "article": "der"
  },
  {
    "noun": "Zisterzienserorden",
    "article": "der"
  },
  {
    "noun": "Zitronenbaum",
    "article": "der"
  },
  {
    "noun": "Zitronenfalter",
    "article": "der"
  },
  {
    "noun": "Zitronensaft",
    "article": "der"
  },
  {
    "noun": "Zitteraal",
    "article": "der"
  },
  {
    "noun": "Zitterrochen",
    "article": "der"
  },
  {
    "noun": "Zitwer",
    "article": "der"
  },
  {
    "noun": "Zitz",
    "article": "der"
  },
  {
    "noun": "Zivilanzug",
    "article": "der"
  },
  {
    "noun": "Zivildienst",
    "article": "der"
  },
  {
    "noun": "Zivilist",
    "article": "der"
  },
  {
    "noun": "Zivilprozess",
    "article": "der"
  },
  {
    "noun": "Zivilschutz",
    "article": "der"
  },
  {
    "noun": "Zivilstand",
    "article": "der"
  },
  {
    "noun": "Zivilstandsamt",
    "article": "der"
  },
  {
    "noun": "Zloty",
    "article": "der"
  },
  {
    "noun": "Zobel",
    "article": "der"
  },
  {
    "noun": "Zobelpelz",
    "article": "der"
  },
  {
    "noun": "Zocker",
    "article": "der"
  },
  {
    "noun": "Zodiakus",
    "article": "der"
  },
  {
    "noun": "Zoff",
    "article": "der"
  },
  {
    "noun": "Zoll",
    "article": "der"
  },
  {
    "noun": "Zollausschluss",
    "article": "der"
  },
  {
    "noun": "Zolleinnehmer",
    "article": "der"
  },
  {
    "noun": "Zollhund",
    "article": "der"
  },
  {
    "noun": "Zollkrieg",
    "article": "der"
  },
  {
    "noun": "Zollstock",
    "article": "der"
  },
  {
    "noun": "Zolltarif",
    "article": "der"
  },
  {
    "noun": "Zollvertrag",
    "article": "der"
  },
  {
    "noun": "Zombie",
    "article": "der"
  },
  {
    "noun": "Zonentarif",
    "article": "der"
  },
  {
    "noun": "Zoo",
    "article": "der"
  },
  {
    "noun": "Zoolith",
    "article": "der"
  },
  {
    "noun": "Zoologe",
    "article": "der"
  },
  {
    "noun": "Zoophage",
    "article": "der"
  },
  {
    "noun": "Zoophyt",
    "article": "der"
  },
  {
    "noun": "Zopf",
    "article": "der"
  },
  {
    "noun": "Zopfstil",
    "article": "der"
  },
  {
    "noun": "Zorn",
    "article": "der"
  },
  {
    "noun": "Zornausbruch",
    "article": "der"
  },
  {
    "noun": "Zuber",
    "article": "der"
  },
  {
    "noun": "Zubereiter",
    "article": "der"
  },
  {
    "noun": "Zubiss",
    "article": "der"
  },
  {
    "noun": "Zubringer",
    "article": "der"
  },
  {
    "noun": "Zubringerbus",
    "article": "der"
  },
  {
    "noun": "Zubringerdienst",
    "article": "der"
  },
  {
    "noun": "Zubringerverkehr",
    "article": "der"
  },
  {
    "noun": "Zuchtbulle",
    "article": "der"
  },
  {
    "noun": "Zuchteber",
    "article": "der"
  },
  {
    "noun": "Zuchthengst",
    "article": "der"
  },
  {
    "noun": "Zuchtmeister",
    "article": "der"
  },
  {
    "noun": "Zuchtstier",
    "article": "der"
  },
  {
    "noun": "Zuck",
    "article": "der"
  },
  {
    "noun": "Zucker",
    "article": "der"
  },
  {
    "noun": "Zuckergehalt",
    "article": "der"
  },
  {
    "noun": "Zuckerguss",
    "article": "der"
  },
  {
    "noun": "Zuckerhut",
    "article": "der"
  },
  {
    "noun": "Zuckerkandis",
    "article": "der"
  },
  {
    "noun": "Zuckerkuchen",
    "article": "der"
  },
  {
    "noun": "Zuckerstreuer",
    "article": "der"
  },
  {
    "noun": "Zudrang",
    "article": "der"
  },
  {
    "noun": "Zufahrtsweg",
    "article": "der"
  },
  {
    "noun": "Zufall",
    "article": "der"
  },
  {
    "noun": "Zufallsfund",
    "article": "der"
  },
  {
    "noun": "Zufallsgenerator",
    "article": "der"
  },
  {
    "noun": "Zufallstreffer",
    "article": "der"
  },
  {
    "noun": "Zufluchtsort",
    "article": "der"
  },
  {
    "noun": "Zufluss",
    "article": "der"
  },
  {
    "noun": "Zugang",
    "article": "der"
  },
  {
    "noun": "Zugangsweg",
    "article": "der"
  },
  {
    "noun": "Zugbegleiter",
    "article": "der"
  },
  {
    "noun": "Zugewinn",
    "article": "der"
  },
  {
    "noun": "Zugochse",
    "article": "der"
  },
  {
    "noun": "Zugriff",
    "article": "der"
  },
  {
    "noun": "Zugsverkehr",
    "article": "der"
  },
  {
    "noun": "Zugverkehr",
    "article": "der"
  },
  {
    "noun": "Zugvogel",
    "article": "der"
  },
  {
    "noun": "Zugwind",
    "article": "der"
  },
  {
    "noun": "Zugzwang",
    "article": "der"
  },
  {
    "noun": "Zukauf",
    "article": "der"
  },
  {
    "noun": "Zukunftsforscher",
    "article": "der"
  },
  {
    "noun": "Zukunftsglaube",
    "article": "der"
  },
  {
    "noun": "Zukunftsplan",
    "article": "der"
  },
  {
    "noun": "Zukunftsroman",
    "article": "der"
  },
  {
    "noun": "Zukunftsstaat",
    "article": "der"
  },
  {
    "noun": "Zukunftstraum",
    "article": "der"
  },
  {
    "noun": "Zulassungsschein",
    "article": "der"
  },
  {
    "noun": "Zulauf",
    "article": "der"
  },
  {
    "noun": "Zulieferant",
    "article": "der"
  },
  {
    "noun": "Zulieferbetrieb",
    "article": "der"
  },
  {
    "noun": "Zulieferer",
    "article": "der"
  },
  {
    "noun": "Zuname",
    "article": "der"
  },
  {
    "noun": "Zunder",
    "article": "der"
  },
  {
    "noun": "Zunftbrief",
    "article": "der"
  },
  {
    "noun": "Zunftgenosse",
    "article": "der"
  },
  {
    "noun": "Zunftmeister",
    "article": "der"
  },
  {
    "noun": "Zungenbelag",
    "article": "der"
  },
  {
    "noun": "Zungenbrecher",
    "article": "der"
  },
  {
    "noun": "Zungenkrampf",
    "article": "der"
  },
  {
    "noun": "Zungenkuss",
    "article": "der"
  },
  {
    "noun": "Zungenlaut",
    "article": "der"
  },
  {
    "noun": "Zungenschlag",
    "article": "der"
  },
  {
    "noun": "Zurichter",
    "article": "der"
  },
  {
    "noun": "Zuruf",
    "article": "der"
  },
  {
    "noun": "Zusammenbau",
    "article": "der"
  },
  {
    "noun": "Zusammenbruch",
    "article": "der"
  },
  {
    "noun": "Zusammenfall",
    "article": "der"
  },
  {
    "noun": "Zusammenfluss",
    "article": "der"
  },
  {
    "noun": "Zusammenhalt",
    "article": "der"
  },
  {
    "noun": "Zusammenhang",
    "article": "der"
  },
  {
    "noun": "Zusammenklang",
    "article": "der"
  },
  {
    "noun": "Zusammenprall",
    "article": "der"
  },
  {
    "noun": "Zusammenschluss",
    "article": "der"
  },
  {
    "noun": "Zusammensturz",
    "article": "der"
  },
  {
    "noun": "Zusatz",
    "article": "der"
  },
  {
    "noun": "Zusatzantrag",
    "article": "der"
  },
  {
    "noun": "Zusatzstoff",
    "article": "der"
  },
  {
    "noun": "Zuschauer",
    "article": "der"
  },
  {
    "noun": "Zuschauerraum",
    "article": "der"
  },
  {
    "noun": "Zuschlag",
    "article": "der"
  },
  {
    "noun": "Zuschlagssatz",
    "article": "der"
  },
  {
    "noun": "Zuschlagsstoff",
    "article": "der"
  },
  {
    "noun": "Zuschlagstoff",
    "article": "der"
  },
  {
    "noun": "Zuschneider",
    "article": "der"
  },
  {
    "noun": "Zuschnitt",
    "article": "der"
  },
  {
    "noun": "Zuschuss",
    "article": "der"
  },
  {
    "noun": "Zuschussbetrieb",
    "article": "der"
  },
  {
    "noun": "Zuschussbogen",
    "article": "der"
  },
  {
    "noun": "Zuseher",
    "article": "der"
  },
  {
    "noun": "Zuspruch",
    "article": "der"
  },
  {
    "noun": "Zustand",
    "article": "der"
  },
  {
    "noun": "Zusteller",
    "article": "der"
  },
  {
    "noun": "Zustellversuch",
    "article": "der"
  },
  {
    "noun": "Zustrom",
    "article": "der"
  },
  {
    "noun": "Zutritt",
    "article": "der"
  },
  {
    "noun": "Zutrunk",
    "article": "der"
  },
  {
    "noun": "Zuverdienst",
    "article": "der"
  },
  {
    "noun": "Zuwachs",
    "article": "der"
  },
  {
    "noun": "Zuwanderer",
    "article": "der"
  },
  {
    "noun": "Zuzug",
    "article": "der"
  },
  {
    "noun": "Zwang",
    "article": "der"
  },
  {
    "noun": "Zwangsarbeiter",
    "article": "der"
  },
  {
    "noun": "Zwangskurs",
    "article": "der"
  },
  {
    "noun": "Zwangsneurotiker",
    "article": "der"
  },
  {
    "noun": "Zwangsvergleich",
    "article": "der"
  },
  {
    "noun": "Zweck",
    "article": "der"
  },
  {
    "noun": "Zweckaufwand",
    "article": "der"
  },
  {
    "noun": "Zweckbau",
    "article": "der"
  },
  {
    "noun": "Zweckoptimismus",
    "article": "der"
  },
  {
    "noun": "Zweckpessimismus",
    "article": "der"
  },
  {
    "noun": "Zwecksatz",
    "article": "der"
  },
  {
    "noun": "Zweckverband",
    "article": "der"
  },
  {
    "noun": "Zweiachser",
    "article": "der"
  },
  {
    "noun": "Zweibeiner",
    "article": "der"
  },
  {
    "noun": "Zweidecker",
    "article": "der"
  },
  {
    "noun": "Zweier",
    "article": "der"
  },
  {
    "noun": "Zweieranschluss",
    "article": "der"
  },
  {
    "noun": "Zweierbob",
    "article": "der"
  },
  {
    "noun": "Zweifarbendruck",
    "article": "der"
  },
  {
    "noun": "Zweifel",
    "article": "der"
  },
  {
    "noun": "Zweifelsfall",
    "article": "der"
  },
  {
    "noun": "Zweifler",
    "article": "der"
  },
  {
    "noun": "Zweig",
    "article": "der"
  },
  {
    "noun": "Zweigbetrieb",
    "article": "der"
  },
  {
    "noun": "Zweihufer",
    "article": "der"
  },
  {
    "noun": "Zweijahresplan",
    "article": "der"
  },
  {
    "noun": "Zweikampf",
    "article": "der"
  },
  {
    "noun": "Zweikomponentenkleber",
    "article": "der"
  },
  {
    "noun": "Zweimaster",
    "article": "der"
  },
  {
    "noun": "Zweiphasenstrom",
    "article": "der"
  },
  {
    "noun": "Zweiradfahrer",
    "article": "der"
  },
  {
    "noun": "Zweireiher",
    "article": "der"
  },
  {
    "noun": "Zweiruderer",
    "article": "der"
  },
  {
    "noun": "Zweisitzer",
    "article": "der"
  },
  {
    "noun": "Zweispitz",
    "article": "der"
  },
  {
    "noun": "Zweisternegeneral",
    "article": "der"
  },
  {
    "noun": "Zweitakter",
    "article": "der"
  },
  {
    "noun": "Zweitaktmotor",
    "article": "der"
  },
  {
    "noun": "Zweitbesitzer",
    "article": "der"
  },
  {
    "noun": "Zweitdruck",
    "article": "der"
  },
  {
    "noun": "Zweiteiler",
    "article": "der"
  },
  {
    "noun": "Zweitschlag",
    "article": "der"
  },
  {
    "noun": "Zweitwagen",
    "article": "der"
  },
  {
    "noun": "Zweitwohnsitz",
    "article": "der"
  },
  {
    "noun": "Zweivierteltakt",
    "article": "der"
  },
  {
    "noun": "Zweizeiler",
    "article": "der"
  },
  {
    "noun": "Zweizylindermotor",
    "article": "der"
  },
  {
    "noun": "Zwerg",
    "article": "der"
  },
  {
    "noun": "Zwergbaum",
    "article": "der"
  },
  {
    "noun": "Zwergdackel",
    "article": "der"
  },
  {
    "noun": "Zwergpudel",
    "article": "der"
  },
  {
    "noun": "Zwergstaat",
    "article": "der"
  },
  {
    "noun": "Zwergtaucher",
    "article": "der"
  },
  {
    "noun": "Zwergwuchs",
    "article": "der"
  },
  {
    "noun": "Zwetschenbaum",
    "article": "der"
  },
  {
    "noun": "Zwetschenkern",
    "article": "der"
  },
  {
    "noun": "Zwetschgenkuchen",
    "article": "der"
  },
  {
    "noun": "Zwickel",
    "article": "der"
  },
  {
    "noun": "Zwicker",
    "article": "der"
  },
  {
    "noun": "Zwieback",
    "article": "der"
  },
  {
    "noun": "Zwiebelkuchen",
    "article": "der"
  },
  {
    "noun": "Zwiebelmarmor",
    "article": "der"
  },
  {
    "noun": "Zwiebelring",
    "article": "der"
  },
  {
    "noun": "Zwiebelturm",
    "article": "der"
  },
  {
    "noun": "Zwiegesang",
    "article": "der"
  },
  {
    "noun": "Zwielaut",
    "article": "der"
  },
  {
    "noun": "Zwiespalt",
    "article": "der"
  },
  {
    "noun": "Zwilch",
    "article": "der"
  },
  {
    "noun": "Zwilling",
    "article": "der"
  },
  {
    "noun": "Zwillingsbruder",
    "article": "der"
  },
  {
    "noun": "Zwillingsreifen",
    "article": "der"
  },
  {
    "noun": "Zwinger",
    "article": "der"
  },
  {
    "noun": "Zwingherr",
    "article": "der"
  },
  {
    "noun": "Zwirn",
    "article": "der"
  },
  {
    "noun": "Zwirnsfaden",
    "article": "der"
  },
  {
    "noun": "Zwischenakt",
    "article": "der"
  },
  {
    "noun": "Zwischenaufenthalt",
    "article": "der"
  },
  {
    "noun": "Zwischenbericht",
    "article": "der"
  },
  {
    "noun": "Zwischenbescheid",
    "article": "der"
  },
  {
    "noun": "Zwischenboden",
    "article": "der"
  },
  {
    "noun": "Zwischenfall",
    "article": "der"
  },
  {
    "noun": "Zwischenfruchtbau",
    "article": "der"
  },
  {
    "noun": "Zwischenhalt",
    "article": "der"
  },
  {
    "noun": "Zwischenhandel",
    "article": "der"
  },
  {
    "noun": "Zwischenkredit",
    "article": "der"
  },
  {
    "noun": "Zwischenlauf",
    "article": "der"
  },
  {
    "noun": "Zwischenraum",
    "article": "der"
  },
  {
    "noun": "Zwischenring",
    "article": "der"
  },
  {
    "noun": "Zwischenruf",
    "article": "der"
  },
  {
    "noun": "Zwischenrufer",
    "article": "der"
  },
  {
    "noun": "Zwischensatz",
    "article": "der"
  },
  {
    "noun": "Zwischenschein",
    "article": "der"
  },
  {
    "noun": "Zwischenschritt",
    "article": "der"
  },
  {
    "noun": "Zwischenspeicher",
    "article": "der"
  },
  {
    "noun": "Zwischenspurt",
    "article": "der"
  },
  {
    "noun": "Zwischenstock",
    "article": "der"
  },
  {
    "noun": "Zwischenstopp",
    "article": "der"
  },
  {
    "noun": "Zwischentext",
    "article": "der"
  },
  {
    "noun": "Zwischentitel",
    "article": "der"
  },
  {
    "noun": "Zwischenton",
    "article": "der"
  },
  {
    "noun": "Zwischenwirt",
    "article": "der"
  },
  {
    "noun": "Zwist",
    "article": "der"
  },
  {
    "noun": "Zwitter",
    "article": "der"
  },
  {
    "noun": "Zyklop",
    "article": "der"
  },
  {
    "noun": "Zyklus",
    "article": "der"
  },
  {
    "noun": "Zylinder",
    "article": "der"
  },
  {
    "noun": "Zylinderblock",
    "article": "der"
  },
  {
    "noun": "Zylinderhut",
    "article": "der"
  },
  {
    "noun": "Zylinderkopf",
    "article": "der"
  },
  {
    "noun": "Zyniker",
    "article": "der"
  },
  {
    "noun": "Zynismus",
    "article": "der"
  },
  {
    "noun": "Zyprer",
    "article": "der"
  },
  {
    "noun": "Zypriot",
    "article": "der"
  },
  {
    "noun": "Dan",
    "article": "der"
  },
  {
    "noun": "Danaer",
    "article": "der"
  },
  {
    "noun": "Dandy",
    "article": "der"
  },
  {
    "noun": "Dankbrief",
    "article": "der"
  },
  {
    "noun": "Dankesbrief",
    "article": "der"
  },
  {
    "noun": "Dankgottesdienst",
    "article": "der"
  },
  {
    "noun": "Darlehensgeber",
    "article": "der"
  },
  {
    "noun": "Darlehensnehmer",
    "article": "der"
  },
  {
    "noun": "Darlehensvertrag",
    "article": "der"
  },
  {
    "noun": "Darlehenszins",
    "article": "der"
  },
  {
    "noun": "Darleiher",
    "article": "der"
  },
  {
    "noun": "Darling",
    "article": "der"
  },
  {
    "noun": "Darm",
    "article": "der"
  },
  {
    "noun": "Darmkrebs",
    "article": "der"
  },
  {
    "noun": "Darmverschluss",
    "article": "der"
  },
  {
    "noun": "Darrofen",
    "article": "der"
  },
  {
    "noun": "Darwinismus",
    "article": "der"
  },
  {
    "noun": "Darwinist",
    "article": "der"
  },
  {
    "noun": "Daseinskampf",
    "article": "der"
  },
  {
    "noun": "Datenaustausch",
    "article": "der"
  },
  {
    "noun": "Datenbestand",
    "article": "der"
  },
  {
    "noun": "Datensatz",
    "article": "der"
  },
  {
    "noun": "Datenschutz",
    "article": "der"
  },
  {
    "noun": "Datenspeicher",
    "article": "der"
  },
  {
    "noun": "Datentypist",
    "article": "der"
  },
  {
    "noun": "Dativ",
    "article": "der"
  },
  {
    "noun": "Datolith",
    "article": "der"
  },
  {
    "noun": "Datowechsel",
    "article": "der"
  },
  {
    "noun": "Dattelkern",
    "article": "der"
  },
  {
    "noun": "Datumsstempel",
    "article": "der"
  },
  {
    "noun": "Datumstempel",
    "article": "der"
  },
  {
    "noun": "Dauerauftrag",
    "article": "der"
  },
  {
    "noun": "Dauerbetrieb",
    "article": "der"
  },
  {
    "noun": "Dauerbrandofen",
    "article": "der"
  },
  {
    "noun": "Dauerbrenner",
    "article": "der"
  },
  {
    "noun": "Dauerfrost",
    "article": "der"
  },
  {
    "noun": "Dauergast",
    "article": "der"
  },
  {
    "noun": "Dauerkunde",
    "article": "der"
  },
  {
    "noun": "Dauerlauf",
    "article": "der"
  },
  {
    "noun": "Dauermagnet",
    "article": "der"
  },
  {
    "noun": "Dauerposten",
    "article": "der"
  },
  {
    "noun": "Dauerregen",
    "article": "der"
  },
  {
    "noun": "Dauerschaden",
    "article": "der"
  },
  {
    "noun": "Dauerseller",
    "article": "der"
  },
  {
    "noun": "Dauerstrom",
    "article": "der"
  },
  {
    "noun": "Dauertest",
    "article": "der"
  },
  {
    "noun": "Dauerton",
    "article": "der"
  },
  {
    "noun": "Dauerwald",
    "article": "der"
  },
  {
    "noun": "Dauerzustand",
    "article": "der"
  },
  {
    "noun": "Daumen",
    "article": "der"
  },
  {
    "noun": "Daumenabdruck",
    "article": "der"
  },
  {
    "noun": "Daumennagel",
    "article": "der"
  },
  {
    "noun": "Davidstern",
    "article": "der"
  },
  {
    "noun": "Davit",
    "article": "der"
  },
  {
    "noun": "DAX",
    "article": "der"
  },
  {
    "noun": "Deal",
    "article": "der"
  },
  {
    "noun": "Dealer",
    "article": "der"
  },
  {
    "noun": "Debetsaldo",
    "article": "der"
  },
  {
    "noun": "Debitor",
    "article": "der"
  },
  {
    "noun": "Debugger",
    "article": "der"
  },
  {
    "noun": "Dechan",
    "article": "der"
  },
  {
    "noun": "Deckanstrich",
    "article": "der"
  },
  {
    "noun": "Deckel",
    "article": "der"
  },
  {
    "noun": "Deckelkorb",
    "article": "der"
  },
  {
    "noun": "Deckenbalken",
    "article": "der"
  },
  {
    "noun": "Deckenfluter",
    "article": "der"
  },
  {
    "noun": "Deckenputz",
    "article": "der"
  },
  {
    "noun": "Deckenventilator",
    "article": "der"
  },
  {
    "noun": "Deckhengst",
    "article": "der"
  },
  {
    "noun": "Deckmantel",
    "article": "der"
  },
  {
    "noun": "Deckname",
    "article": "der"
  },
  {
    "noun": "Deckpassagier",
    "article": "der"
  },
  {
    "noun": "Deckungsgraben",
    "article": "der"
  },
  {
    "noun": "Deckungskauf",
    "article": "der"
  },
  {
    "noun": "Decoder",
    "article": "der"
  },
  {
    "noun": "Decubitus",
    "article": "der"
  },
  {
    "noun": "Defekt",
    "article": "der"
  },
  {
    "noun": "Defensivkrieg",
    "article": "der"
  },
  {
    "noun": "Defiliermarsch",
    "article": "der"
  },
  {
    "noun": "Definitionsbereich",
    "article": "der"
  },
  {
    "noun": "Defizient",
    "article": "der"
  },
  {
    "noun": "Deflektor",
    "article": "der"
  },
  {
    "noun": "Defraudant",
    "article": "der"
  },
  {
    "noun": "Defroster",
    "article": "der"
  },
  {
    "noun": "Degen",
    "article": "der"
  },
  {
    "noun": "Degengriff",
    "article": "der"
  },
  {
    "noun": "Degout",
    "article": "der"
  },
  {
    "noun": "Dehnungsmesser",
    "article": "der"
  },
  {
    "noun": "Deich",
    "article": "der"
  },
  {
    "noun": "Deichbau",
    "article": "der"
  },
  {
    "noun": "Deichbruch",
    "article": "der"
  },
  {
    "noun": "Deichselbruch",
    "article": "der"
  },
  {
    "noun": "Deismus",
    "article": "der"
  },
  {
    "noun": "Deist",
    "article": "der"
  },
  {
    "noun": "Dekapode",
    "article": "der"
  },
  {
    "noun": "Deklamator",
    "article": "der"
  },
  {
    "noun": "Dekorateur",
    "article": "der"
  },
  {
    "noun": "Dekorationsstoff",
    "article": "der"
  },
  {
    "noun": "Delegat",
    "article": "der"
  },
  {
    "noun": "Delfin",
    "article": "der"
  },
  {
    "noun": "Delinquent",
    "article": "der"
  },
  {
    "noun": "Delkrederefonds",
    "article": "der"
  },
  {
    "noun": "Deltamuskel",
    "article": "der"
  },
  {
    "noun": "Demagoge",
    "article": "der"
  },
  {
    "noun": "Demant",
    "article": "der"
  },
  {
    "noun": "Demiurg",
    "article": "der"
  },
  {
    "noun": "Demodulator",
    "article": "der"
  },
  {
    "noun": "Demograf",
    "article": "der"
  },
  {
    "noun": "Demokrat",
    "article": "der"
  },
  {
    "noun": "Demokratismus",
    "article": "der"
  },
  {
    "noun": "Demonstrant",
    "article": "der"
  },
  {
    "noun": "Demoskop",
    "article": "der"
  },
  {
    "noun": "Dendrit",
    "article": "der"
  },
  {
    "noun": "Dendrologe",
    "article": "der"
  },
  {
    "noun": "Denker",
    "article": "der"
  },
  {
    "noun": "Denkfehler",
    "article": "der"
  },
  {
    "noun": "Denkmalschutz",
    "article": "der"
  },
  {
    "noun": "Denkprozess",
    "article": "der"
  },
  {
    "noun": "Denkspruch",
    "article": "der"
  },
  {
    "noun": "Denkzettel",
    "article": "der"
  },
  {
    "noun": "Dental",
    "article": "der"
  },
  {
    "noun": "Dentist",
    "article": "der"
  },
  {
    "noun": "Denunziant",
    "article": "der"
  },
  {
    "noun": "Deoroller",
    "article": "der"
  },
  {
    "noun": "Deponent",
    "article": "der"
  },
  {
    "noun": "Depositar",
    "article": "der"
  },
  {
    "noun": "Depotinhaber",
    "article": "der"
  },
  {
    "noun": "Depotwechsel",
    "article": "der"
  },
  {
    "noun": "Depotwert",
    "article": "der"
  },
  {
    "noun": "Depp",
    "article": "der"
  },
  {
    "noun": "Derivationswinkel",
    "article": "der"
  },
  {
    "noun": "Dermatologe",
    "article": "der"
  },
  {
    "noun": "Derrickkran",
    "article": "der"
  },
  {
    "noun": "Derwisch",
    "article": "der"
  },
  {
    "noun": "Desensibilisator",
    "article": "der"
  },
  {
    "noun": "Deserteur",
    "article": "der"
  },
  {
    "noun": "Designer",
    "article": "der"
  },
  {
    "noun": "Desintegrator",
    "article": "der"
  },
  {
    "noun": "Deskriptor",
    "article": "der"
  },
  {
    "noun": "Despot",
    "article": "der"
  },
  {
    "noun": "Despotismus",
    "article": "der"
  },
  {
    "noun": "Dessertteller",
    "article": "der"
  },
  {
    "noun": "Dessertwein",
    "article": "der"
  },
  {
    "noun": "Destillierapparat",
    "article": "der"
  },
  {
    "noun": "Destillierkolben",
    "article": "der"
  },
  {
    "noun": "Destinatar",
    "article": "der"
  },
  {
    "noun": "Deszendent",
    "article": "der"
  },
  {
    "noun": "Detailhandel",
    "article": "der"
  },
  {
    "noun": "Detaillist",
    "article": "der"
  },
  {
    "noun": "Detektiv",
    "article": "der"
  },
  {
    "noun": "Detektor",
    "article": "der"
  },
  {
    "noun": "Determinismus",
    "article": "der"
  },
  {
    "noun": "Determinist",
    "article": "der"
  },
  {
    "noun": "Detonator",
    "article": "der"
  },
  {
    "noun": "Deuter",
    "article": "der"
  },
  {
    "noun": "Devisenhandel",
    "article": "der"
  },
  {
    "noun": "Devisenkurs",
    "article": "der"
  },
  {
    "noun": "Devisenmarkt",
    "article": "der"
  },
  {
    "noun": "Devisenschmuggel",
    "article": "der"
  },
  {
    "noun": "Dezember",
    "article": "der"
  },
  {
    "noun": "Dezimalbruch",
    "article": "der"
  },
  {
    "noun": "Dezimalwert",
    "article": "der"
  },
  {
    "noun": "Diabas",
    "article": "der"
  },
  {
    "noun": "Diabetes",
    "article": "der"
  },
  {
    "noun": "Diadoche",
    "article": "der"
  },
  {
    "noun": "Diagonalreifen",
    "article": "der"
  },
  {
    "noun": "Diakon",
    "article": "der"
  },
  {
    "noun": "Dialekt",
    "article": "der"
  },
  {
    "noun": "Dialektforscher",
    "article": "der"
  },
  {
    "noun": "Dialektiker",
    "article": "der"
  },
  {
    "noun": "Dialog",
    "article": "der"
  },
  {
    "noun": "Dialogist",
    "article": "der"
  },
  {
    "noun": "Dialysator",
    "article": "der"
  },
  {
    "noun": "Diamagnetismus",
    "article": "der"
  },
  {
    "noun": "Diamant",
    "article": "der"
  },
  {
    "noun": "Diamantbohrer",
    "article": "der"
  },
  {
    "noun": "Diamantglanz",
    "article": "der"
  },
  {
    "noun": "Diamantring",
    "article": "der"
  },
  {
    "noun": "Diamantstaub",
    "article": "der"
  },
  {
    "noun": "Diameter",
    "article": "der"
  },
  {
    "noun": "Diapason",
    "article": "der"
  },
  {
    "noun": "Diapir",
    "article": "der"
  },
  {
    "noun": "Diaprojektor",
    "article": "der"
  },
  {
    "noun": "Diaspor",
    "article": "der"
  },
  {
    "noun": "Diatomeenschlamm",
    "article": "der"
  },
  {
    "noun": "Diatomit",
    "article": "der"
  },
  {
    "noun": "Diavortrag",
    "article": "der"
  },
  {
    "noun": "Dichroismus",
    "article": "der"
  },
  {
    "noun": "Dichroit",
    "article": "der"
  },
  {
    "noun": "Dichtemesser",
    "article": "der"
  },
  {
    "noun": "Dichter",
    "article": "der"
  },
  {
    "noun": "Dichterling",
    "article": "der"
  },
  {
    "noun": "Dichtungsring",
    "article": "der"
  },
  {
    "noun": "Dickbauch",
    "article": "der"
  },
  {
    "noun": "Dickdarm",
    "article": "der"
  },
  {
    "noun": "Dickdarmkrebs",
    "article": "der"
  },
  {
    "noun": "Dickkopf",
    "article": "der"
  },
  {
    "noun": "Didaktiker",
    "article": "der"
  },
  {
    "noun": "Dieb",
    "article": "der"
  },
  {
    "noun": "Diebstahl",
    "article": "der"
  },
  {
    "noun": "Diener",
    "article": "der"
  },
  {
    "noun": "Dienst",
    "article": "der"
  },
  {
    "noun": "Dienstag",
    "article": "der"
  },
  {
    "noun": "Dienstagabend",
    "article": "der"
  },
  {
    "noun": "Dienstantritt",
    "article": "der"
  },
  {
    "noun": "Dienstanzug",
    "article": "der"
  },
  {
    "noun": "Dienstauftrag",
    "article": "der"
  },
  {
    "noun": "Dienstausweis",
    "article": "der"
  },
  {
    "noun": "Dienstbefehl",
    "article": "der"
  },
  {
    "noun": "Dienstbereich",
    "article": "der"
  },
  {
    "noun": "Dienstbetrieb",
    "article": "der"
  },
  {
    "noun": "Dienstbote",
    "article": "der"
  },
  {
    "noun": "Diensteid",
    "article": "der"
  },
  {
    "noun": "Diensteifer",
    "article": "der"
  },
  {
    "noun": "Dienstgeber",
    "article": "der"
  },
  {
    "noun": "Dienstgebrauch",
    "article": "der"
  },
  {
    "noun": "Dienstgrad",
    "article": "der"
  },
  {
    "noun": "Dienstherr",
    "article": "der"
  },
  {
    "noun": "Diensthund",
    "article": "der"
  },
  {
    "noun": "Dienstleister",
    "article": "der"
  },
  {
    "noun": "Dienstleistungsberuf",
    "article": "der"
  },
  {
    "noun": "Dienstleistungsbetrieb",
    "article": "der"
  },
  {
    "noun": "Dienstleistungssektor",
    "article": "der"
  },
  {
    "noun": "Dienstleistungsverkehr",
    "article": "der"
  },
  {
    "noun": "Dienstmann",
    "article": "der"
  },
  {
    "noun": "Dienstnehmer",
    "article": "der"
  },
  {
    "noun": "Dienstplan",
    "article": "der"
  },
  {
    "noun": "Dienstschluss",
    "article": "der"
  },
  {
    "noun": "Dienststempel",
    "article": "der"
  },
  {
    "noun": "Dienstvertrag",
    "article": "der"
  },
  {
    "noun": "Dienstwagen",
    "article": "der"
  },
  {
    "noun": "Dienstweg",
    "article": "der"
  },
  {
    "noun": "Diesel",
    "article": "der"
  },
  {
    "noun": "Dieselkraftstoff",
    "article": "der"
  },
  {
    "noun": "Dieselmotor",
    "article": "der"
  },
  {
    "noun": "Dietrich",
    "article": "der"
  },
  {
    "noun": "Differentialquotient",
    "article": "der"
  },
  {
    "noun": "Differentialschutz",
    "article": "der"
  },
  {
    "noun": "Differentialzoll",
    "article": "der"
  },
  {
    "noun": "Differentiator",
    "article": "der"
  },
  {
    "noun": "Differenzbetrag",
    "article": "der"
  },
  {
    "noun": "Differenzenquotient",
    "article": "der"
  },
  {
    "noun": "Diffusor",
    "article": "der"
  },
  {
    "noun": "Digitaldrucker",
    "article": "der"
  },
  {
    "noun": "Digitalrechner",
    "article": "der"
  },
  {
    "noun": "Digraf",
    "article": "der"
  },
  {
    "noun": "Diktator",
    "article": "der"
  },
  {
    "noun": "Dildo",
    "article": "der"
  },
  {
    "noun": "Dilettant",
    "article": "der"
  },
  {
    "noun": "Dilettantismus",
    "article": "der"
  },
  {
    "noun": "Dill",
    "article": "der"
  },
  {
    "noun": "Dimeter",
    "article": "der"
  },
  {
    "noun": "Dimmer",
    "article": "der"
  },
  {
    "noun": "Dimorphismus",
    "article": "der"
  },
  {
    "noun": "Dinar",
    "article": "der"
  },
  {
    "noun": "Dingo",
    "article": "der"
  },
  {
    "noun": "Dinkel",
    "article": "der"
  },
  {
    "noun": "Dinosaurier",
    "article": "der"
  },
  {
    "noun": "Dinosaurus",
    "article": "der"
  },
  {
    "noun": "Diopsid",
    "article": "der"
  },
  {
    "noun": "Dioptas",
    "article": "der"
  },
  {
    "noun": "Diorit",
    "article": "der"
  },
  {
    "noun": "Dip",
    "article": "der"
  },
  {
    "noun": "Diphthong",
    "article": "der"
  },
  {
    "noun": "Diplomand",
    "article": "der"
  },
  {
    "noun": "Diplomat",
    "article": "der"
  },
  {
    "noun": "Diplomatenkoffer",
    "article": "der"
  },
  {
    "noun": "Dipol",
    "article": "der"
  },
  {
    "noun": "Direktbezug",
    "article": "der"
  },
  {
    "noun": "Direktflug",
    "article": "der"
  },
  {
    "noun": "Direktimport",
    "article": "der"
  },
  {
    "noun": "Direktionsassistent",
    "article": "der"
  },
  {
    "noun": "Direktor",
    "article": "der"
  },
  {
    "noun": "Direktverkauf",
    "article": "der"
  },
  {
    "noun": "Dirigent",
    "article": "der"
  },
  {
    "noun": "Dirigismus",
    "article": "der"
  },
  {
    "noun": "Dirndlrock",
    "article": "der"
  },
  {
    "noun": "Discounter",
    "article": "der"
  },
  {
    "noun": "Discountladen",
    "article": "der"
  },
  {
    "noun": "Diseur",
    "article": "der"
  },
  {
    "noun": "Diskant",
    "article": "der"
  },
  {
    "noun": "Diskjockey",
    "article": "der"
  },
  {
    "noun": "Diskont",
    "article": "der"
  },
  {
    "noun": "Diskontsatz",
    "article": "der"
  },
  {
    "noun": "Diskriminator",
    "article": "der"
  },
  {
    "noun": "Diskurs",
    "article": "der"
  },
  {
    "noun": "Diskus",
    "article": "der"
  },
  {
    "noun": "Diskussionsbeitrag",
    "article": "der"
  },
  {
    "noun": "Diskussionsgegenstand",
    "article": "der"
  },
  {
    "noun": "Diskussionsleiter",
    "article": "der"
  },
  {
    "noun": "Diskussionsredner",
    "article": "der"
  },
  {
    "noun": "Diskussionsteilnehmer",
    "article": "der"
  },
  {
    "noun": "Diskuswerfer",
    "article": "der"
  },
  {
    "noun": "Dispatcher",
    "article": "der"
  },
  {
    "noun": "Dispens",
    "article": "der"
  },
  {
    "noun": "Disponent",
    "article": "der"
  },
  {
    "noun": "Dispositionskredit",
    "article": "der"
  },
  {
    "noun": "Disput",
    "article": "der"
  },
  {
    "noun": "Disputant",
    "article": "der"
  },
  {
    "noun": "Dissertant",
    "article": "der"
  },
  {
    "noun": "Dissident",
    "article": "der"
  },
  {
    "noun": "Distanzwechsel",
    "article": "der"
  },
  {
    "noun": "Distelfink",
    "article": "der"
  },
  {
    "noun": "Disthen",
    "article": "der"
  },
  {
    "noun": "Distributor",
    "article": "der"
  },
  {
    "noun": "Distrikt",
    "article": "der"
  },
  {
    "noun": "Dithyrambus",
    "article": "der"
  },
  {
    "noun": "Diversant",
    "article": "der"
  },
  {
    "noun": "Dividend",
    "article": "der"
  },
  {
    "noun": "Dividendenertrag",
    "article": "der"
  },
  {
    "noun": "Dividendenschein",
    "article": "der"
  },
  {
    "noun": "Divisionskommandeur",
    "article": "der"
  },
  {
    "noun": "Divisor",
    "article": "der"
  },
  {
    "noun": "Diwan",
    "article": "der"
  },
  {
    "noun": "Dixie",
    "article": "der"
  },
  {
    "noun": "Dobermann",
    "article": "der"
  },
  {
    "noun": "Dobermannpinscher",
    "article": "der"
  },
  {
    "noun": "Docht",
    "article": "der"
  },
  {
    "noun": "Dockarbeiter",
    "article": "der"
  },
  {
    "noun": "Docker",
    "article": "der"
  },
  {
    "noun": "Doge",
    "article": "der"
  },
  {
    "noun": "Dogger",
    "article": "der"
  },
  {
    "noun": "Dogmatiker",
    "article": "der"
  },
  {
    "noun": "Dogmatismus",
    "article": "der"
  },
  {
    "noun": "Doktor",
    "article": "der"
  },
  {
    "noun": "Doktorand",
    "article": "der"
  },
  {
    "noun": "Doktorhut",
    "article": "der"
  },
  {
    "noun": "Doktoringenieur",
    "article": "der"
  },
  {
    "noun": "Doktorvater",
    "article": "der"
  },
  {
    "noun": "Dokumentar",
    "article": "der"
  },
  {
    "noun": "Dokumentarbericht",
    "article": "der"
  },
  {
    "noun": "Dokumentarfilm",
    "article": "der"
  },
  {
    "noun": "Dolch",
    "article": "der"
  },
  {
    "noun": "Dolchstich",
    "article": "der"
  },
  {
    "noun": "Dolerit",
    "article": "der"
  },
  {
    "noun": "Dollar",
    "article": "der"
  },
  {
    "noun": "Dollarkurs",
    "article": "der"
  },
  {
    "noun": "Dolmen",
    "article": "der"
  },
  {
    "noun": "Dolmetsch",
    "article": "der"
  },
  {
    "noun": "Dolmetscher",
    "article": "der"
  },
  {
    "noun": "Dolomit",
    "article": "der"
  },
  {
    "noun": "Dom",
    "article": "der"
  },
  {
    "noun": "Domherr",
    "article": "der"
  },
  {
    "noun": "Dominikaner",
    "article": "der"
  },
  {
    "noun": "Dominoeffekt",
    "article": "der"
  },
  {
    "noun": "Dominostein",
    "article": "der"
  },
  {
    "noun": "Domizilwechsel",
    "article": "der"
  },
  {
    "noun": "Dompteur",
    "article": "der"
  },
  {
    "noun": "Don",
    "article": "der"
  },
  {
    "noun": "Donator",
    "article": "der"
  },
  {
    "noun": "Donner",
    "article": "der"
  },
  {
    "noun": "Donnerbalken",
    "article": "der"
  },
  {
    "noun": "Donnerschlag",
    "article": "der"
  },
  {
    "noun": "Donnerstag",
    "article": "der"
  },
  {
    "noun": "Doofmann",
    "article": "der"
  },
  {
    "noun": "Dopingverdacht",
    "article": "der"
  },
  {
    "noun": "Doppelagent",
    "article": "der"
  },
  {
    "noun": "Doppelboden",
    "article": "der"
  },
  {
    "noun": "Doppelbruch",
    "article": "der"
  },
  {
    "noun": "Doppelbuchstabe",
    "article": "der"
  },
  {
    "noun": "Doppeldecker",
    "article": "der"
  },
  {
    "noun": "Doppelfehler",
    "article": "der"
  },
  {
    "noun": "Doppelklick",
    "article": "der"
  },
  {
    "noun": "Doppelkopf",
    "article": "der"
  },
  {
    "noun": "Doppelname",
    "article": "der"
  },
  {
    "noun": "Doppelpass",
    "article": "der"
  },
  {
    "noun": "Doppelpunkt",
    "article": "der"
  },
  {
    "noun": "Doppelsieg",
    "article": "der"
  },
  {
    "noun": "Doppelsinn",
    "article": "der"
  },
  {
    "noun": "Doppelstecker",
    "article": "der"
  },
  {
    "noun": "Doppelstern",
    "article": "der"
  },
  {
    "noun": "Doppelverdiener",
    "article": "der"
  },
  {
    "noun": "Doppelverdienst",
    "article": "der"
  },
  {
    "noun": "Doppelvokal",
    "article": "der"
  },
  {
    "noun": "Doppelzentner",
    "article": "der"
  },
  {
    "noun": "Dopplereffekt",
    "article": "der"
  },
  {
    "noun": "Dorfbewohner",
    "article": "der"
  },
  {
    "noun": "Dorfplatz",
    "article": "der"
  },
  {
    "noun": "Dorfschulze",
    "article": "der"
  },
  {
    "noun": "Dorftrottel",
    "article": "der"
  },
  {
    "noun": "Dorn",
    "article": "der"
  },
  {
    "noun": "Dornbusch",
    "article": "der"
  },
  {
    "noun": "Dornenstrauch",
    "article": "der"
  },
  {
    "noun": "Dornhai",
    "article": "der"
  },
  {
    "noun": "Dorsch",
    "article": "der"
  },
  {
    "noun": "Dottersack",
    "article": "der"
  },
  {
    "noun": "Download",
    "article": "der"
  },
  {
    "noun": "Doyen",
    "article": "der"
  },
  {
    "noun": "Dozent",
    "article": "der"
  },
  {
    "noun": "Drache",
    "article": "der"
  },
  {
    "noun": "Drachen",
    "article": "der"
  },
  {
    "noun": "Drachenflieger",
    "article": "der"
  },
  {
    "noun": "Dragoman",
    "article": "der"
  },
  {
    "noun": "Dragoner",
    "article": "der"
  },
  {
    "noun": "Dragster",
    "article": "der"
  },
  {
    "noun": "Draht",
    "article": "der"
  },
  {
    "noun": "Drahtbesen",
    "article": "der"
  },
  {
    "noun": "Drahtesel",
    "article": "der"
  },
  {
    "noun": "Drahtfunk",
    "article": "der"
  },
  {
    "noun": "Drahtkorb",
    "article": "der"
  },
  {
    "noun": "Drahtseilakt",
    "article": "der"
  },
  {
    "noun": "Drahtstift",
    "article": "der"
  },
  {
    "noun": "Drahtzaun",
    "article": "der"
  },
  {
    "noun": "Drahtzieher",
    "article": "der"
  },
  {
    "noun": "Drain",
    "article": "der"
  },
  {
    "noun": "Drall",
    "article": "der"
  },
  {
    "noun": "Dramatiker",
    "article": "der"
  },
  {
    "noun": "Dramaturg",
    "article": "der"
  },
  {
    "noun": "Drang",
    "article": "der"
  },
  {
    "noun": "Drechsler",
    "article": "der"
  },
  {
    "noun": "Dreck",
    "article": "der"
  },
  {
    "noun": "Dreckfink",
    "article": "der"
  },
  {
    "noun": "Drecksack",
    "article": "der"
  },
  {
    "noun": "Dreckskerl",
    "article": "der"
  },
  {
    "noun": "Dreckspatz",
    "article": "der"
  },
  {
    "noun": "Dreh",
    "article": "der"
  },
  {
    "noun": "Drehbleistift",
    "article": "der"
  },
  {
    "noun": "Drehbuchautor",
    "article": "der"
  },
  {
    "noun": "Dreher",
    "article": "der"
  },
  {
    "noun": "Drehimpuls",
    "article": "der"
  },
  {
    "noun": "Drehknopf",
    "article": "der"
  },
  {
    "noun": "Drehkolbenmotor",
    "article": "der"
  },
  {
    "noun": "Drehkondensator",
    "article": "der"
  },
  {
    "noun": "Drehkran",
    "article": "der"
  },
  {
    "noun": "Drehorgelspieler",
    "article": "der"
  },
  {
    "noun": "Drehort",
    "article": "der"
  },
  {
    "noun": "Drehpunkt",
    "article": "der"
  },
  {
    "noun": "Drehschalter",
    "article": "der"
  },
  {
    "noun": "Drehsinn",
    "article": "der"
  },
  {
    "noun": "Drehstahl",
    "article": "der"
  },
  {
    "noun": "Drehstrom",
    "article": "der"
  },
  {
    "noun": "Drehstuhl",
    "article": "der"
  },
  {
    "noun": "Drehzahlmesser",
    "article": "der"
  },
  {
    "noun": "Dreier",
    "article": "der"
  },
  {
    "noun": "Dreifarbendruck",
    "article": "der"
  },
  {
    "noun": "Dreikanter",
    "article": "der"
  },
  {
    "noun": "Dreiklang",
    "article": "der"
  },
  {
    "noun": "Dreilaut",
    "article": "der"
  },
  {
    "noun": "Dreimaster",
    "article": "der"
  },
  {
    "noun": "Dreiphasenstrom",
    "article": "der"
  },
  {
    "noun": "Dreipunktgurt",
    "article": "der"
  },
  {
    "noun": "Dreisatz",
    "article": "der"
  },
  {
    "noun": "Dreispitz",
    "article": "der"
  },
  {
    "noun": "Dreisprung",
    "article": "der"
  },
  {
    "noun": "Dreitagebart",
    "article": "der"
  },
  {
    "noun": "Dreivierteltakt",
    "article": "der"
  },
  {
    "noun": "Dreizack",
    "article": "der"
  },
  {
    "noun": "Drell",
    "article": "der"
  },
  {
    "noun": "Drempel",
    "article": "der"
  },
  {
    "noun": "Drescher",
    "article": "der"
  },
  {
    "noun": "Dreschflegel",
    "article": "der"
  },
  {
    "noun": "Dress",
    "article": "der"
  },
  {
    "noun": "Dresseur",
    "article": "der"
  },
  {
    "noun": "Dressman",
    "article": "der"
  },
  {
    "noun": "Drill",
    "article": "der"
  },
  {
    "noun": "Drillbohrer",
    "article": "der"
  },
  {
    "noun": "Drillich",
    "article": "der"
  },
  {
    "noun": "Drilling",
    "article": "der"
  },
  {
    "noun": "Dringlichkeitsantrag",
    "article": "der"
  },
  {
    "noun": "Drittschuldner",
    "article": "der"
  },
  {
    "noun": "Drogenfahnder",
    "article": "der"
  },
  {
    "noun": "Drogenkonsum",
    "article": "der"
  },
  {
    "noun": "Drogenkonsument",
    "article": "der"
  },
  {
    "noun": "Drogenkurier",
    "article": "der"
  },
  {
    "noun": "Drogeriemarkt",
    "article": "der"
  },
  {
    "noun": "Drogist",
    "article": "der"
  },
  {
    "noun": "Drohbrief",
    "article": "der"
  },
  {
    "noun": "Droschkenkutscher",
    "article": "der"
  },
  {
    "noun": "Druck",
    "article": "der"
  },
  {
    "noun": "Druckabfall",
    "article": "der"
  },
  {
    "noun": "Druckanstieg",
    "article": "der"
  },
  {
    "noun": "Druckausgleich",
    "article": "der"
  },
  {
    "noun": "Druckbleistift",
    "article": "der"
  },
  {
    "noun": "Druckbogen",
    "article": "der"
  },
  {
    "noun": "Druckbuchstabe",
    "article": "der"
  },
  {
    "noun": "Drucker",
    "article": "der"
  },
  {
    "noun": "Druckertreiber",
    "article": "der"
  },
  {
    "noun": "Druckfehler",
    "article": "der"
  },
  {
    "noun": "Druckfehlerteufel",
    "article": "der"
  },
  {
    "noun": "Druckknopf",
    "article": "der"
  },
  {
    "noun": "Drucklufthammer",
    "article": "der"
  },
  {
    "noun": "Druckmesser",
    "article": "der"
  },
  {
    "noun": "Druckort",
    "article": "der"
  },
  {
    "noun": "Druckpunkt",
    "article": "der"
  },
  {
    "noun": "Druckraum",
    "article": "der"
  },
  {
    "noun": "Druckregler",
    "article": "der"
  },
  {
    "noun": "Druckstoff",
    "article": "der"
  },
  {
    "noun": "Drucktopf",
    "article": "der"
  },
  {
    "noun": "Druckunterschied",
    "article": "der"
  },
  {
    "noun": "Druckverband",
    "article": "der"
  },
  {
    "noun": "Druckwasserreaktor",
    "article": "der"
  },
  {
    "noun": "Druide",
    "article": "der"
  },
  {
    "noun": "Drusch",
    "article": "der"
  },
  {
    "noun": "Dschihad",
    "article": "der"
  },
  {
    "noun": "Dual",
    "article": "der"
  },
  {
    "noun": "Dualismus",
    "article": "der"
  },
  {
    "noun": "Dualist",
    "article": "der"
  },
  {
    "noun": "Ductus",
    "article": "der"
  },
  {
    "noun": "Dudelsack",
    "article": "der"
  },
  {
    "noun": "Dudelsackpfeifer",
    "article": "der"
  },
  {
    "noun": "Dudelsackspieler",
    "article": "der"
  },
  {
    "noun": "Duellant",
    "article": "der"
  },
  {
    "noun": "Duft",
    "article": "der"
  },
  {
    "noun": "Dufthauch",
    "article": "der"
  },
  {
    "noun": "Duftstoff",
    "article": "der"
  },
  {
    "noun": "Dukaten",
    "article": "der"
  },
  {
    "noun": "Duktus",
    "article": "der"
  },
  {
    "noun": "Dulder",
    "article": "der"
  },
  {
    "noun": "Dummerjan",
    "article": "der"
  },
  {
    "noun": "Dummkopf",
    "article": "der"
  },
  {
    "noun": "Dummy",
    "article": "der"
  },
  {
    "noun": "Dumper",
    "article": "der"
  },
  {
    "noun": "Dumpingpreis",
    "article": "der"
  },
  {
    "noun": "Dung",
    "article": "der"
  },
  {
    "noun": "Dunghaufen",
    "article": "der"
  },
  {
    "noun": "Dungkarren",
    "article": "der"
  },
  {
    "noun": "Dungstreuer",
    "article": "der"
  },
  {
    "noun": "Dunkelmann",
    "article": "der"
  },
  {
    "noun": "Dunkelstrom",
    "article": "der"
  },
  {
    "noun": "Dunst",
    "article": "der"
  },
  {
    "noun": "Dunstschleier",
    "article": "der"
  },
  {
    "noun": "Duplexbetrieb",
    "article": "der"
  },
  {
    "noun": "Durchblick",
    "article": "der"
  },
  {
    "noun": "Durchbruch",
    "article": "der"
  },
  {
    "noun": "Durchfall",
    "article": "der"
  },
  {
    "noun": "Durchfaller",
    "article": "der"
  },
  {
    "noun": "Durchflug",
    "article": "der"
  },
  {
    "noun": "Durchfluss",
    "article": "der"
  },
  {
    "noun": "Durchflussmesser",
    "article": "der"
  },
  {
    "noun": "Durchfuhrhandel",
    "article": "der"
  },
  {
    "noun": "Durchgang",
    "article": "der"
  },
  {
    "noun": "Durchgangsbahnhof",
    "article": "der"
  },
  {
    "noun": "Durchgangsverkehr",
    "article": "der"
  },
  {
    "noun": "Durchhang",
    "article": "der"
  },
  {
    "noun": "Durchhau",
    "article": "der"
  },
  {
    "noun": "Durchhieb",
    "article": "der"
  },
  {
    "noun": "Durchlass",
    "article": "der"
  },
  {
    "noun": "Durchlauf",
    "article": "der"
  },
  {
    "noun": "Durchlauferhitzer",
    "article": "der"
  },
  {
    "noun": "Durchmarsch",
    "article": "der"
  },
  {
    "noun": "Durchmesser",
    "article": "der"
  },
  {
    "noun": "Durchsatz",
    "article": "der"
  },
  {
    "noun": "Durchschlag",
    "article": "der"
  },
  {
    "noun": "Durchschlupf",
    "article": "der"
  },
  {
    "noun": "Durchschnitt",
    "article": "der"
  },
  {
    "noun": "Durchschnittsertrag",
    "article": "der"
  },
  {
    "noun": "Durchschnittslohn",
    "article": "der"
  },
  {
    "noun": "Durchschnittsmensch",
    "article": "der"
  },
  {
    "noun": "Durchschnittsverdienst",
    "article": "der"
  },
  {
    "noun": "Durchschnittswert",
    "article": "der"
  },
  {
    "noun": "Durchschuss",
    "article": "der"
  },
  {
    "noun": "Durchstich",
    "article": "der"
  },
  {
    "noun": "Durchsuchungsbefehl",
    "article": "der"
  },
  {
    "noun": "Durchzug",
    "article": "der"
  },
  {
    "noun": "Duroplast",
    "article": "der"
  },
  {
    "noun": "Durst",
    "article": "der"
  },
  {
    "noun": "Duschkopf",
    "article": "der"
  },
  {
    "noun": "Duschraum",
    "article": "der"
  },
  {
    "noun": "Duschvorhang",
    "article": "der"
  },
  {
    "noun": "Dusel",
    "article": "der"
  },
  {
    "noun": "Dussel",
    "article": "der"
  },
  {
    "noun": "Dutzendmensch",
    "article": "der"
  },
  {
    "noun": "Dynamo",
    "article": "der"
  },
  {
    "noun": "Dyskrasit",
    "article": "der"
  },
  {
    "noun": "Lagerhalter",
    "article": "der"
  },
  {
    "noun": "Lagerist",
    "article": "der"
  },
  {
    "noun": "Lagerplatz",
    "article": "der"
  },
  {
    "noun": "Lagerraum",
    "article": "der"
  },
  {
    "noun": "Lagerschein",
    "article": "der"
  },
  {
    "noun": "Lagerverwalter",
    "article": "der"
  },
  {
    "noun": "Lagophthalmus",
    "article": "der"
  },
  {
    "noun": "Lahn",
    "article": "der"
  },
  {
    "noun": "Laib",
    "article": "der"
  },
  {
    "noun": "Laich",
    "article": "der"
  },
  {
    "noun": "Laichplatz",
    "article": "der"
  },
  {
    "noun": "Laie",
    "article": "der"
  },
  {
    "noun": "Laienbruder",
    "article": "der"
  },
  {
    "noun": "Laienrichter",
    "article": "der"
  },
  {
    "noun": "Laienstand",
    "article": "der"
  },
  {
    "noun": "Laizismus",
    "article": "der"
  },
  {
    "noun": "Lakai",
    "article": "der"
  },
  {
    "noun": "Lakkolith",
    "article": "der"
  },
  {
    "noun": "Lakonismus",
    "article": "der"
  },
  {
    "noun": "Lammbraten",
    "article": "der"
  },
  {
    "noun": "Lampendocht",
    "article": "der"
  },
  {
    "noun": "Lampenschirm",
    "article": "der"
  },
  {
    "noun": "Lamprophyr",
    "article": "der"
  },
  {
    "noun": "Lancier",
    "article": "der"
  },
  {
    "noun": "Langfinger",
    "article": "der"
  },
  {
    "noun": "Langlauf",
    "article": "der"
  },
  {
    "noun": "Langlochziegel",
    "article": "der"
  },
  {
    "noun": "Langweiler",
    "article": "der"
  },
  {
    "noun": "Langwellensender",
    "article": "der"
  },
  {
    "noun": "Lanzenfarn",
    "article": "der"
  },
  {
    "noun": "Lanzenstich",
    "article": "der"
  },
  {
    "noun": "Laote",
    "article": "der"
  },
  {
    "noun": "Lapislazuli",
    "article": "der"
  },
  {
    "noun": "Lappe",
    "article": "der"
  },
  {
    "noun": "Lappen",
    "article": "der"
  },
  {
    "noun": "Lappentaucher",
    "article": "der"
  },
  {
    "noun": "Lapsus",
    "article": "der"
  },
  {
    "noun": "Laser",
    "article": "der"
  },
  {
    "noun": "Laserdrucker",
    "article": "der"
  },
  {
    "noun": "Laserstrahl",
    "article": "der"
  },
  {
    "noun": "Lastarm",
    "article": "der"
  },
  {
    "noun": "Lastenaufzug",
    "article": "der"
  },
  {
    "noun": "Lastenausgleich",
    "article": "der"
  },
  {
    "noun": "Lastkahn",
    "article": "der"
  },
  {
    "noun": "Lastkraftwagen",
    "article": "der"
  },
  {
    "noun": "Lastverteiler",
    "article": "der"
  },
  {
    "noun": "Lastwagen",
    "article": "der"
  },
  {
    "noun": "Lastwagenfahrer",
    "article": "der"
  },
  {
    "noun": "Lastzug",
    "article": "der"
  },
  {
    "noun": "Lasurit",
    "article": "der"
  },
  {
    "noun": "Lasurlack",
    "article": "der"
  },
  {
    "noun": "Lateralplan",
    "article": "der"
  },
  {
    "noun": "Laterit",
    "article": "der"
  },
  {
    "noun": "Lateritboden",
    "article": "der"
  },
  {
    "noun": "Laternenpfahl",
    "article": "der"
  },
  {
    "noun": "Latex",
    "article": "der"
  },
  {
    "noun": "Latinismus",
    "article": "der"
  },
  {
    "noun": "Latsch",
    "article": "der"
  },
  {
    "noun": "Lattenrost",
    "article": "der"
  },
  {
    "noun": "Lattenzaun",
    "article": "der"
  },
  {
    "noun": "Latz",
    "article": "der"
  },
  {
    "noun": "Laubbaum",
    "article": "der"
  },
  {
    "noun": "Laubengang",
    "article": "der"
  },
  {
    "noun": "Laubfall",
    "article": "der"
  },
  {
    "noun": "Laubfrosch",
    "article": "der"
  },
  {
    "noun": "Laubwald",
    "article": "der"
  },
  {
    "noun": "Lauch",
    "article": "der"
  },
  {
    "noun": "Lauf",
    "article": "der"
  },
  {
    "noun": "Laufbursche",
    "article": "der"
  },
  {
    "noun": "Laufgang",
    "article": "der"
  },
  {
    "noun": "Laufhund",
    "article": "der"
  },
  {
    "noun": "Laufjunge",
    "article": "der"
  },
  {
    "noun": "Laufkran",
    "article": "der"
  },
  {
    "noun": "Laufschritt",
    "article": "der"
  },
  {
    "noun": "Laufschuh",
    "article": "der"
  },
  {
    "noun": "Laufstall",
    "article": "der"
  },
  {
    "noun": "Laufsteg",
    "article": "der"
  },
  {
    "noun": "Laufstuhl",
    "article": "der"
  },
  {
    "noun": "Laufvogel",
    "article": "der"
  },
  {
    "noun": "Laufzettel",
    "article": "der"
  },
  {
    "noun": "Laureat",
    "article": "der"
  },
  {
    "noun": "Lausbub",
    "article": "der"
  },
  {
    "noun": "Lausbube",
    "article": "der"
  },
  {
    "noun": "Lauschangriff",
    "article": "der"
  },
  {
    "noun": "Lauscher",
    "article": "der"
  },
  {
    "noun": "Lausebengel",
    "article": "der"
  },
  {
    "noun": "Lausekerl",
    "article": "der"
  },
  {
    "noun": "Lauser",
    "article": "der"
  },
  {
    "noun": "Lausitzer",
    "article": "der"
  },
  {
    "noun": "Laut",
    "article": "der"
  },
  {
    "noun": "Lautsprecher",
    "article": "der"
  },
  {
    "noun": "Lautstand",
    "article": "der"
  },
  {
    "noun": "Lavastrom",
    "article": "der"
  },
  {
    "noun": "Layouter",
    "article": "der"
  },
  {
    "noun": "Lazarettzug",
    "article": "der"
  },
  {
    "noun": "Leader",
    "article": "der"
  },
  {
    "noun": "Lebemann",
    "article": "der"
  },
  {
    "noun": "Lebensabend",
    "article": "der"
  },
  {
    "noun": "Lebensabriss",
    "article": "der"
  },
  {
    "noun": "Lebensabschnitt",
    "article": "der"
  },
  {
    "noun": "Lebensabschnittspartner",
    "article": "der"
  },
  {
    "noun": "Lebensbaum",
    "article": "der"
  },
  {
    "noun": "Lebensbereich",
    "article": "der"
  },
  {
    "noun": "Lebensgang",
    "article": "der"
  },
  {
    "noun": "Lebensgenuss",
    "article": "der"
  },
  {
    "noun": "Lebenshaltungsindex",
    "article": "der"
  },
  {
    "noun": "Lebenshunger",
    "article": "der"
  },
  {
    "noun": "Lebensinhalt",
    "article": "der"
  },
  {
    "noun": "Lebenskampf",
    "article": "der"
  },
  {
    "noun": "Lebenslauf",
    "article": "der"
  },
  {
    "noun": "Lebensmittelchemiker",
    "article": "der"
  },
  {
    "noun": "Lebensmittelhandel",
    "article": "der"
  },
  {
    "noun": "Lebensmittelpreis",
    "article": "der"
  },
  {
    "noun": "Lebensmut",
    "article": "der"
  },
  {
    "noun": "Lebensnerv",
    "article": "der"
  },
  {
    "noun": "Lebenspartner",
    "article": "der"
  },
  {
    "noun": "Lebensraum",
    "article": "der"
  },
  {
    "noun": "Lebensretter",
    "article": "der"
  },
  {
    "noun": "Lebensrhythmus",
    "article": "der"
  },
  {
    "noun": "Lebenssaft",
    "article": "der"
  },
  {
    "noun": "Lebensstandard",
    "article": "der"
  },
  {
    "noun": "Lebensstil",
    "article": "der"
  },
  {
    "noun": "Lebensunterhalt",
    "article": "der"
  },
  {
    "noun": "Lebenswandel",
    "article": "der"
  },
  {
    "noun": "Lebensweg",
    "article": "der"
  },
  {
    "noun": "Lebenswille",
    "article": "der"
  },
  {
    "noun": "Lebenszweck",
    "article": "der"
  },
  {
    "noun": "Leberfleck",
    "article": "der"
  },
  {
    "noun": "Leberkrebs",
    "article": "der"
  },
  {
    "noun": "Leberlappen",
    "article": "der"
  },
  {
    "noun": "Leberschaden",
    "article": "der"
  },
  {
    "noun": "Lebertran",
    "article": "der"
  },
  {
    "noun": "Leberwurstbaum",
    "article": "der"
  },
  {
    "noun": "Lebkuchen",
    "article": "der"
  },
  {
    "noun": "Lebtag",
    "article": "der"
  },
  {
    "noun": "Leckerbissen",
    "article": "der"
  },
  {
    "noun": "Ledereinband",
    "article": "der"
  },
  {
    "noun": "Lederer",
    "article": "der"
  },
  {
    "noun": "Lederhandschuh",
    "article": "der"
  },
  {
    "noun": "Lederkoffer",
    "article": "der"
  },
  {
    "noun": "Ledernacken",
    "article": "der"
  },
  {
    "noun": "Lederriemen",
    "article": "der"
  },
  {
    "noun": "Lederschuh",
    "article": "der"
  },
  {
    "noun": "Lederstiefel",
    "article": "der"
  },
  {
    "noun": "Leerdarm",
    "article": "der"
  },
  {
    "noun": "Leerlauf",
    "article": "der"
  },
  {
    "noun": "Leerverkauf",
    "article": "der"
  },
  {
    "noun": "Leerzug",
    "article": "der"
  },
  {
    "noun": "Legastheniker",
    "article": "der"
  },
  {
    "noun": "Legationsrat",
    "article": "der"
  },
  {
    "noun": "Leger",
    "article": "der"
  },
  {
    "noun": "Legionar",
    "article": "der"
  },
  {
    "noun": "Legionssoldat",
    "article": "der"
  },
  {
    "noun": "Leguan",
    "article": "der"
  },
  {
    "noun": "Lehm",
    "article": "der"
  },
  {
    "noun": "Lehmbau",
    "article": "der"
  },
  {
    "noun": "Lehmboden",
    "article": "der"
  },
  {
    "noun": "Lehmziegel",
    "article": "der"
  },
  {
    "noun": "Lehnsessel",
    "article": "der"
  },
  {
    "noun": "Lehnsherr",
    "article": "der"
  },
  {
    "noun": "Lehnsmann",
    "article": "der"
  },
  {
    "noun": "Lehnstuhl",
    "article": "der"
  },
  {
    "noun": "Lehrauftrag",
    "article": "der"
  },
  {
    "noun": "Lehrbehelf",
    "article": "der"
  },
  {
    "noun": "Lehrberuf",
    "article": "der"
  },
  {
    "noun": "Lehrbetrieb",
    "article": "der"
  },
  {
    "noun": "Lehrbogen",
    "article": "der"
  },
  {
    "noun": "Lehrbrief",
    "article": "der"
  },
  {
    "noun": "Lehrbub",
    "article": "der"
  },
  {
    "noun": "Lehrdorn",
    "article": "der"
  },
  {
    "noun": "Lehrer",
    "article": "der"
  },
  {
    "noun": "Lehrfilm",
    "article": "der"
  },
  {
    "noun": "Lehrgang",
    "article": "der"
  },
  {
    "noun": "Lehrgegenstand",
    "article": "der"
  },
  {
    "noun": "Lehrinhalt",
    "article": "der"
  },
  {
    "noun": "Lehrjunge",
    "article": "der"
  },
  {
    "noun": "Lehrling",
    "article": "der"
  },
  {
    "noun": "Lehrmeister",
    "article": "der"
  },
  {
    "noun": "Lehrpfad",
    "article": "der"
  },
  {
    "noun": "Lehrplan",
    "article": "der"
  },
  {
    "noun": "Lehrsaal",
    "article": "der"
  },
  {
    "noun": "Lehrsatz",
    "article": "der"
  },
  {
    "noun": "Lehrstoff",
    "article": "der"
  },
  {
    "noun": "Lehrstuhl",
    "article": "der"
  },
  {
    "noun": "Lehrstuhlinhaber",
    "article": "der"
  },
  {
    "noun": "Lehrvertrag",
    "article": "der"
  },
  {
    "noun": "Leib",
    "article": "der"
  },
  {
    "noun": "Leibarzt",
    "article": "der"
  },
  {
    "noun": "Leibkoch",
    "article": "der"
  },
  {
    "noun": "Leibschmerz",
    "article": "der"
  },
  {
    "noun": "Leichdorn",
    "article": "der"
  },
  {
    "noun": "Leichenbeschauer",
    "article": "der"
  },
  {
    "noun": "Leichenbestatter",
    "article": "der"
  },
  {
    "noun": "Leichenfledderer",
    "article": "der"
  },
  {
    "noun": "Leichengeruch",
    "article": "der"
  },
  {
    "noun": "Leichenschmaus",
    "article": "der"
  },
  {
    "noun": "Leichenwagen",
    "article": "der"
  },
  {
    "noun": "Leichenzug",
    "article": "der"
  },
  {
    "noun": "Leichnam",
    "article": "der"
  },
  {
    "noun": "Leichtathlet",
    "article": "der"
  },
  {
    "noun": "Leichtbau",
    "article": "der"
  },
  {
    "noun": "Leichtbeton",
    "article": "der"
  },
  {
    "noun": "Leichter",
    "article": "der"
  },
  {
    "noun": "Leichtlohn",
    "article": "der"
  },
  {
    "noun": "Leichtmatrose",
    "article": "der"
  },
  {
    "noun": "Leichtsinn",
    "article": "der"
  },
  {
    "noun": "Leichtsinnsfehler",
    "article": "der"
  },
  {
    "noun": "Leichtwasserreaktor",
    "article": "der"
  },
  {
    "noun": "Leidensdruck",
    "article": "der"
  },
  {
    "noun": "Leidensgenosse",
    "article": "der"
  },
  {
    "noun": "Leierkasten",
    "article": "der"
  },
  {
    "noun": "Leierkastenmann",
    "article": "der"
  },
  {
    "noun": "Leiharbeiter",
    "article": "der"
  },
  {
    "noun": "Leiharbeitnehmer",
    "article": "der"
  },
  {
    "noun": "Leihgeber",
    "article": "der"
  },
  {
    "noun": "Leihverkehr",
    "article": "der"
  },
  {
    "noun": "Leihwagen",
    "article": "der"
  },
  {
    "noun": "Leim",
    "article": "der"
  },
  {
    "noun": "Lein",
    "article": "der"
  },
  {
    "noun": "Leinacker",
    "article": "der"
  },
  {
    "noun": "Leineneinband",
    "article": "der"
  },
  {
    "noun": "Leinpfad",
    "article": "der"
  },
  {
    "noun": "Leinsame",
    "article": "der"
  },
  {
    "noun": "Leinweber",
    "article": "der"
  },
  {
    "noun": "Leisetreter",
    "article": "der"
  },
  {
    "noun": "Leist",
    "article": "der"
  },
  {
    "noun": "Leistenbruch",
    "article": "der"
  },
  {
    "noun": "Leistungsabfall",
    "article": "der"
  },
  {
    "noun": "Leistungsanreiz",
    "article": "der"
  },
  {
    "noun": "Leistungsanspruch",
    "article": "der"
  },
  {
    "noun": "Leistungsanstieg",
    "article": "der"
  },
  {
    "noun": "Leistungsdruck",
    "article": "der"
  },
  {
    "noun": "Leistungsgrad",
    "article": "der"
  },
  {
    "noun": "Leistungskurs",
    "article": "der"
  },
  {
    "noun": "Leistungslohn",
    "article": "der"
  },
  {
    "noun": "Leistungsmesser",
    "article": "der"
  },
  {
    "noun": "Leistungsnachweis",
    "article": "der"
  },
  {
    "noun": "Leistungsschutz",
    "article": "der"
  },
  {
    "noun": "Leistungssport",
    "article": "der"
  },
  {
    "noun": "Leistungsstand",
    "article": "der"
  },
  {
    "noun": "Leistungstest",
    "article": "der"
  },
  {
    "noun": "Leistungsumfang",
    "article": "der"
  },
  {
    "noun": "Leistungsvergleich",
    "article": "der"
  },
  {
    "noun": "Leistungszwang",
    "article": "der"
  },
  {
    "noun": "Leitartikel",
    "article": "der"
  },
  {
    "noun": "Leitartikelschreiber",
    "article": "der"
  },
  {
    "noun": "Leitartikler",
    "article": "der"
  },
  {
    "noun": "Leiter",
    "article": "der"
  },
  {
    "noun": "Leiterwagen",
    "article": "der"
  },
  {
    "noun": "Leitfaden",
    "article": "der"
  },
  {
    "noun": "Leitgedanke",
    "article": "der"
  },
  {
    "noun": "Leithammel",
    "article": "der"
  },
  {
    "noun": "Leitsatz",
    "article": "der"
  },
  {
    "noun": "Leitspruch",
    "article": "der"
  },
  {
    "noun": "Leitstand",
    "article": "der"
  },
  {
    "noun": "Leitstern",
    "article": "der"
  },
  {
    "noun": "Leitstrahl",
    "article": "der"
  },
  {
    "noun": "Leitungsdraht",
    "article": "der"
  },
  {
    "noun": "Leitungshahn",
    "article": "der"
  },
  {
    "noun": "Leitungsmast",
    "article": "der"
  },
  {
    "noun": "Leitungswiderstand",
    "article": "der"
  },
  {
    "noun": "Leitweg",
    "article": "der"
  },
  {
    "noun": "Leitwert",
    "article": "der"
  },
  {
    "noun": "Leitwolf",
    "article": "der"
  },
  {
    "noun": "Leitzins",
    "article": "der"
  },
  {
    "noun": "Lektor",
    "article": "der"
  },
  {
    "noun": "Lemming",
    "article": "der"
  },
  {
    "noun": "Lendenbraten",
    "article": "der"
  },
  {
    "noun": "Lendenschurz",
    "article": "der"
  },
  {
    "noun": "Lendenwirbel",
    "article": "der"
  },
  {
    "noun": "Leng",
    "article": "der"
  },
  {
    "noun": "Leninismus",
    "article": "der"
  },
  {
    "noun": "Leninist",
    "article": "der"
  },
  {
    "noun": "Lenker",
    "article": "der"
  },
  {
    "noun": "Lenkungsausschuss",
    "article": "der"
  },
  {
    "noun": "Lenz",
    "article": "der"
  },
  {
    "noun": "Leopard",
    "article": "der"
  },
  {
    "noun": "Lepidolith",
    "article": "der"
  },
  {
    "noun": "Lepidomelan",
    "article": "der"
  },
  {
    "noun": "Lepidopterologe",
    "article": "der"
  },
  {
    "noun": "Leptit",
    "article": "der"
  },
  {
    "noun": "Lerner",
    "article": "der"
  },
  {
    "noun": "Lernprozess",
    "article": "der"
  },
  {
    "noun": "Lernstoff",
    "article": "der"
  },
  {
    "noun": "Leser",
    "article": "der"
  },
  {
    "noun": "Leserbrief",
    "article": "der"
  },
  {
    "noun": "Leserkreis",
    "article": "der"
  },
  {
    "noun": "Leserwunsch",
    "article": "der"
  },
  {
    "noun": "Lesesaal",
    "article": "der"
  },
  {
    "noun": "Lesestoff",
    "article": "der"
  },
  {
    "noun": "Lette",
    "article": "der"
  },
  {
    "noun": "Letten",
    "article": "der"
  },
  {
    "noun": "Lettersetdruck",
    "article": "der"
  },
  {
    "noun": "Lettner",
    "article": "der"
  },
  {
    "noun": "Letztverbraucher",
    "article": "der"
  },
  {
    "noun": "Leuchter",
    "article": "der"
  },
  {
    "noun": "Leuchtschirm",
    "article": "der"
  },
  {
    "noun": "Leuchtstoff",
    "article": "der"
  },
  {
    "noun": "Leuchtturm",
    "article": "der"
  },
  {
    "noun": "Leuchtzeiger",
    "article": "der"
  },
  {
    "noun": "Leugner",
    "article": "der"
  },
  {
    "noun": "Leukozyt",
    "article": "der"
  },
  {
    "noun": "Leumund",
    "article": "der"
  },
  {
    "noun": "Leuteschinder",
    "article": "der"
  },
  {
    "noun": "Leutnant",
    "article": "der"
  },
  {
    "noun": "Levantiner",
    "article": "der"
  },
  {
    "noun": "Level",
    "article": "der"
  },
  {
    "noun": "Leviatan",
    "article": "der"
  },
  {
    "noun": "Levit",
    "article": "der"
  },
  {
    "noun": "Lexikograf",
    "article": "der"
  },
  {
    "noun": "Libanese",
    "article": "der"
  },
  {
    "noun": "Libanon",
    "article": "der"
  },
  {
    "noun": "Liberalismus",
    "article": "der"
  },
  {
    "noun": "Liberalist",
    "article": "der"
  },
  {
    "noun": "Liberianer",
    "article": "der"
  },
  {
    "noun": "Liberier",
    "article": "der"
  },
  {
    "noun": "Libero",
    "article": "der"
  },
  {
    "noun": "Libyer",
    "article": "der"
  },
  {
    "noun": "Lichtbildervortrag",
    "article": "der"
  },
  {
    "noun": "Lichtblick",
    "article": "der"
  },
  {
    "noun": "Lichtbogen",
    "article": "der"
  },
  {
    "noun": "Lichtdruck",
    "article": "der"
  },
  {
    "noun": "Lichteffekt",
    "article": "der"
  },
  {
    "noun": "Lichteinfall",
    "article": "der"
  },
  {
    "noun": "Lichterbaum",
    "article": "der"
  },
  {
    "noun": "Lichtfleck",
    "article": "der"
  },
  {
    "noun": "Lichtgriffel",
    "article": "der"
  },
  {
    "noun": "Lichthalter",
    "article": "der"
  },
  {
    "noun": "Lichthof",
    "article": "der"
  },
  {
    "noun": "Lichtkegel",
    "article": "der"
  },
  {
    "noun": "Lichtmast",
    "article": "der"
  },
  {
    "noun": "Lichtmesser",
    "article": "der"
  },
  {
    "noun": "Lichtpunkt",
    "article": "der"
  },
  {
    "noun": "Lichtreiz",
    "article": "der"
  },
  {
    "noun": "Lichtsatz",
    "article": "der"
  },
  {
    "noun": "Lichtschacht",
    "article": "der"
  },
  {
    "noun": "Lichtschalter",
    "article": "der"
  },
  {
    "noun": "Lichtschein",
    "article": "der"
  },
  {
    "noun": "Lichtschimmer",
    "article": "der"
  },
  {
    "noun": "Lichtschutzfaktor",
    "article": "der"
  },
  {
    "noun": "Lichtspalt",
    "article": "der"
  },
  {
    "noun": "Lichtstift",
    "article": "der"
  },
  {
    "noun": "Lichtstrahl",
    "article": "der"
  },
  {
    "noun": "Lichtstreifen",
    "article": "der"
  },
  {
    "noun": "Lichtstrom",
    "article": "der"
  },
  {
    "noun": "Lichtwert",
    "article": "der"
  },
  {
    "noun": "Lidrand",
    "article": "der"
  },
  {
    "noun": "Lidschatten",
    "article": "der"
  },
  {
    "noun": "Liebesakt",
    "article": "der"
  },
  {
    "noun": "Liebesapfel",
    "article": "der"
  },
  {
    "noun": "Liebesbote",
    "article": "der"
  },
  {
    "noun": "Liebesbrief",
    "article": "der"
  },
  {
    "noun": "Liebesdienst",
    "article": "der"
  },
  {
    "noun": "Liebesfilm",
    "article": "der"
  },
  {
    "noun": "Liebesgott",
    "article": "der"
  },
  {
    "noun": "Liebeskummer",
    "article": "der"
  },
  {
    "noun": "Liebesroman",
    "article": "der"
  },
  {
    "noun": "Liebestrank",
    "article": "der"
  },
  {
    "noun": "Liebhaber",
    "article": "der"
  },
  {
    "noun": "Liebhaberwert",
    "article": "der"
  },
  {
    "noun": "Liebling",
    "article": "der"
  },
  {
    "noun": "Lieblingslehrer",
    "article": "der"
  },
  {
    "noun": "Lieblingsplatz",
    "article": "der"
  },
  {
    "noun": "Liebreiz",
    "article": "der"
  },
  {
    "noun": "Liederabend",
    "article": "der"
  },
  {
    "noun": "Liederjan",
    "article": "der"
  },
  {
    "noun": "Liedermacher",
    "article": "der"
  },
  {
    "noun": "Liedtext",
    "article": "der"
  },
  {
    "noun": "Lieferant",
    "article": "der"
  },
  {
    "noun": "Lieferauftrag",
    "article": "der"
  },
  {
    "noun": "Lieferer",
    "article": "der"
  },
  {
    "noun": "Lieferschein",
    "article": "der"
  },
  {
    "noun": "Liefertag",
    "article": "der"
  },
  {
    "noun": "Liefertermin",
    "article": "der"
  },
  {
    "noun": "Lieferungsort",
    "article": "der"
  },
  {
    "noun": "Lieferungsvertrag",
    "article": "der"
  },
  {
    "noun": "Liefervertrag",
    "article": "der"
  },
  {
    "noun": "Lieferwagen",
    "article": "der"
  },
  {
    "noun": "Liegeplatz",
    "article": "der"
  },
  {
    "noun": "Liegesitz",
    "article": "der"
  },
  {
    "noun": "Liegestuhl",
    "article": "der"
  },
  {
    "noun": "Liegewagen",
    "article": "der"
  },
  {
    "noun": "Ligand",
    "article": "der"
  },
  {
    "noun": "Lignit",
    "article": "der"
  },
  {
    "noun": "Liguster",
    "article": "der"
  },
  {
    "noun": "Liktor",
    "article": "der"
  },
  {
    "noun": "Lilak",
    "article": "der"
  },
  {
    "noun": "Liliputaner",
    "article": "der"
  },
  {
    "noun": "Limbo",
    "article": "der"
  },
  {
    "noun": "Limbus",
    "article": "der"
  },
  {
    "noun": "Limerick",
    "article": "der"
  },
  {
    "noun": "Limes",
    "article": "der"
  },
  {
    "noun": "Limettensaft",
    "article": "der"
  },
  {
    "noun": "Limnologe",
    "article": "der"
  },
  {
    "noun": "Lindenbaum",
    "article": "der"
  },
  {
    "noun": "Lindenhonig",
    "article": "der"
  },
  {
    "noun": "Lindwurm",
    "article": "der"
  },
  {
    "noun": "Linearbeschleuniger",
    "article": "der"
  },
  {
    "noun": "Linearmotor",
    "article": "der"
  },
  {
    "noun": "Lingual",
    "article": "der"
  },
  {
    "noun": "Linguallaut",
    "article": "der"
  },
  {
    "noun": "Linguist",
    "article": "der"
  },
  {
    "noun": "Linienbus",
    "article": "der"
  },
  {
    "noun": "Linienflug",
    "article": "der"
  },
  {
    "noun": "Linienrichter",
    "article": "der"
  },
  {
    "noun": "Linienschreiber",
    "article": "der"
  },
  {
    "noun": "Linienverkehr",
    "article": "der"
  },
  {
    "noun": "Link",
    "article": "der"
  },
  {
    "noun": "Linker",
    "article": "der"
  },
  {
    "noun": "Linksabbieger",
    "article": "der"
  },
  {
    "noun": "Linksdrall",
    "article": "der"
  },
  {
    "noun": "Linkslenker",
    "article": "der"
  },
  {
    "noun": "Linksradikalismus",
    "article": "der"
  },
  {
    "noun": "Linksruck",
    "article": "der"
  },
  {
    "noun": "Linksverkehr",
    "article": "der"
  },
  {
    "noun": "Linolschnitt",
    "article": "der"
  },
  {
    "noun": "Lipizzaner",
    "article": "der"
  },
  {
    "noun": "Lippenlaut",
    "article": "der"
  },
  {
    "noun": "Lippenstift",
    "article": "der"
  },
  {
    "noun": "Lippfisch",
    "article": "der"
  },
  {
    "noun": "Liquidationswert",
    "article": "der"
  },
  {
    "noun": "Liquidator",
    "article": "der"
  },
  {
    "noun": "Liquor",
    "article": "der"
  },
  {
    "noun": "Listenauszug",
    "article": "der"
  },
  {
    "noun": "Listenpreis",
    "article": "der"
  },
  {
    "noun": "Litauer",
    "article": "der"
  },
  {
    "noun": "Literat",
    "article": "der"
  },
  {
    "noun": "Literaturhistoriker",
    "article": "der"
  },
  {
    "noun": "Literaturkritiker",
    "article": "der"
  },
  {
    "noun": "Literaturpreis",
    "article": "der"
  },
  {
    "noun": "Lithograf",
    "article": "der"
  },
  {
    "noun": "Lizenzentzug",
    "article": "der"
  },
  {
    "noun": "Lizenzgeber",
    "article": "der"
  },
  {
    "noun": "Lizenzinhaber",
    "article": "der"
  },
  {
    "noun": "Lizenznehmer",
    "article": "der"
  },
  {
    "noun": "Lizenzvertrag",
    "article": "der"
  },
  {
    "noun": "Lizitant",
    "article": "der"
  },
  {
    "noun": "LKW",
    "article": "der"
  },
  {
    "noun": "Lobbyismus",
    "article": "der"
  },
  {
    "noun": "Lobbyist",
    "article": "der"
  },
  {
    "noun": "Lobgesang",
    "article": "der"
  },
  {
    "noun": "Lobredner",
    "article": "der"
  },
  {
    "noun": "Lochbeitel",
    "article": "der"
  },
  {
    "noun": "Locher",
    "article": "der"
  },
  {
    "noun": "Lochstreifen",
    "article": "der"
  },
  {
    "noun": "Lochziegel",
    "article": "der"
  },
  {
    "noun": "Lockenkopf",
    "article": "der"
  },
  {
    "noun": "Lockenstab",
    "article": "der"
  },
  {
    "noun": "Lockenwickel",
    "article": "der"
  },
  {
    "noun": "Lockenwickler",
    "article": "der"
  },
  {
    "noun": "Lockruf",
    "article": "der"
  },
  {
    "noun": "Lockspitzel",
    "article": "der"
  },
  {
    "noun": "Lockvogel",
    "article": "der"
  },
  {
    "noun": "Loden",
    "article": "der"
  },
  {
    "noun": "Logarithmand",
    "article": "der"
  },
  {
    "noun": "Logarithmus",
    "article": "der"
  },
  {
    "noun": "Logenbruder",
    "article": "der"
  },
  {
    "noun": "Logenplatz",
    "article": "der"
  },
  {
    "noun": "Logger",
    "article": "der"
  },
  {
    "noun": "Logiker",
    "article": "der"
  },
  {
    "noun": "Logistiker",
    "article": "der"
  },
  {
    "noun": "Lohn",
    "article": "der"
  },
  {
    "noun": "Lohnabbau",
    "article": "der"
  },
  {
    "noun": "Lohnabschluss",
    "article": "der"
  },
  {
    "noun": "Lohnabzug",
    "article": "der"
  },
  {
    "noun": "Lohnarbeiter",
    "article": "der"
  },
  {
    "noun": "Lohnausfall",
    "article": "der"
  },
  {
    "noun": "Lohnausgleich",
    "article": "der"
  },
  {
    "noun": "Lohnbuchhalter",
    "article": "der"
  },
  {
    "noun": "Lohnkostenzuschuss",
    "article": "der"
  },
  {
    "noun": "Lohnsteuerjahresausgleich",
    "article": "der"
  },
  {
    "noun": "Lohnstopp",
    "article": "der"
  },
  {
    "noun": "Lohnstreifen",
    "article": "der"
  },
  {
    "noun": "Lohntag",
    "article": "der"
  },
  {
    "noun": "Lohntarif",
    "article": "der"
  },
  {
    "noun": "Lohntarifvertrag",
    "article": "der"
  },
  {
    "noun": "Lohnverzicht",
    "article": "der"
  },
  {
    "noun": "Lohnzettel",
    "article": "der"
  },
  {
    "noun": "Lokalreporter",
    "article": "der"
  },
  {
    "noun": "Lokalsatz",
    "article": "der"
  },
  {
    "noun": "Lokalverkehr",
    "article": "der"
  },
  {
    "noun": "Lokalzug",
    "article": "der"
  },
  {
    "noun": "Lokativ",
    "article": "der"
  },
  {
    "noun": "Lolch",
    "article": "der"
  },
  {
    "noun": "Lombardkredit",
    "article": "der"
  },
  {
    "noun": "Lombardsatz",
    "article": "der"
  },
  {
    "noun": "Longdrink",
    "article": "der"
  },
  {
    "noun": "Look",
    "article": "der"
  },
  {
    "noun": "Lorbeer",
    "article": "der"
  },
  {
    "noun": "Lorbeerbaum",
    "article": "der"
  },
  {
    "noun": "Lorbeerkranz",
    "article": "der"
  },
  {
    "noun": "Lord",
    "article": "der"
  },
  {
    "noun": "Loskauf",
    "article": "der"
  },
  {
    "noun": "Lotos",
    "article": "der"
  },
  {
    "noun": "Lotse",
    "article": "der"
  },
  {
    "noun": "Lotsendienst",
    "article": "der"
  },
  {
    "noun": "Lotsenzwang",
    "article": "der"
  },
  {
    "noun": "Lottogewinn",
    "article": "der"
  },
  {
    "noun": "Lotus",
    "article": "der"
  },
  {
    "noun": "Loyalist",
    "article": "der"
  },
  {
    "noun": "Luchs",
    "article": "der"
  },
  {
    "noun": "Luddismus",
    "article": "der"
  },
  {
    "noun": "Luddit",
    "article": "der"
  },
  {
    "noun": "Luffaschwamm",
    "article": "der"
  },
  {
    "noun": "Luftakrobat",
    "article": "der"
  },
  {
    "noun": "Luftangriff",
    "article": "der"
  },
  {
    "noun": "Luftaustausch",
    "article": "der"
  },
  {
    "noun": "Luftballon",
    "article": "der"
  },
  {
    "noun": "Luftbefeuchter",
    "article": "der"
  },
  {
    "noun": "Luftdruck",
    "article": "der"
  },
  {
    "noun": "Luftdruckmesser",
    "article": "der"
  },
  {
    "noun": "Luftfahrer",
    "article": "der"
  },
  {
    "noun": "Luftgeist",
    "article": "der"
  },
  {
    "noun": "Lufthafen",
    "article": "der"
  },
  {
    "noun": "Lufthauch",
    "article": "der"
  },
  {
    "noun": "Luftkalk",
    "article": "der"
  },
  {
    "noun": "Luftkampf",
    "article": "der"
  },
  {
    "noun": "Luftkorridor",
    "article": "der"
  },
  {
    "noun": "Luftkrieg",
    "article": "der"
  },
  {
    "noun": "Luftkurort",
    "article": "der"
  },
  {
    "noun": "Luftmangel",
    "article": "der"
  },
  {
    "noun": "Luftpirat",
    "article": "der"
  },
  {
    "noun": "Luftpostbrief",
    "article": "der"
  },
  {
    "noun": "Luftpostleichtbrief",
    "article": "der"
  },
  {
    "noun": "Luftraum",
    "article": "der"
  },
  {
    "noun": "Luftreifen",
    "article": "der"
  },
  {
    "noun": "Luftsack",
    "article": "der"
  },
  {
    "noun": "Luftsauerstoff",
    "article": "der"
  },
  {
    "noun": "Luftschacht",
    "article": "der"
  },
  {
    "noun": "Luftschadstoff",
    "article": "der"
  },
  {
    "noun": "Luftschall",
    "article": "der"
  },
  {
    "noun": "Luftschiffer",
    "article": "der"
  },
  {
    "noun": "Luftschlitz",
    "article": "der"
  },
  {
    "noun": "Luftschutz",
    "article": "der"
  },
  {
    "noun": "Luftschutzbunker",
    "article": "der"
  },
  {
    "noun": "Luftschutzkeller",
    "article": "der"
  },
  {
    "noun": "Luftschutzraum",
    "article": "der"
  },
  {
    "noun": "Luftsport",
    "article": "der"
  },
  {
    "noun": "Luftsprung",
    "article": "der"
  },
  {
    "noun": "Luftstickstoff",
    "article": "der"
  },
  {
    "noun": "Luftstrom",
    "article": "der"
  },
  {
    "noun": "Lufttorpedo",
    "article": "der"
  },
  {
    "noun": "Luftverkehr",
    "article": "der"
  },
  {
    "noun": "Luftwechsel",
    "article": "der"
  },
  {
    "noun": "Luftweg",
    "article": "der"
  },
  {
    "noun": "Luftwiderstand",
    "article": "der"
  },
  {
    "noun": "Luftwirbel",
    "article": "der"
  },
  {
    "noun": "Luftzug",
    "article": "der"
  },
  {
    "noun": "Lug",
    "article": "der"
  },
  {
    "noun": "Lumb",
    "article": "der"
  },
  {
    "noun": "Lumbalwirbel",
    "article": "der"
  },
  {
    "noun": "Lummel",
    "article": "der"
  },
  {
    "noun": "Lump",
    "article": "der"
  },
  {
    "noun": "Lumpenhund",
    "article": "der"
  },
  {
    "noun": "Lumpensammler",
    "article": "der"
  },
  {
    "noun": "Lunch",
    "article": "der"
  },
  {
    "noun": "Lungenkrebs",
    "article": "der"
  },
  {
    "noun": "Lunker",
    "article": "der"
  },
  {
    "noun": "Lupus",
    "article": "der"
  },
  {
    "noun": "Lurch",
    "article": "der"
  },
  {
    "noun": "Lustgarten",
    "article": "der"
  },
  {
    "noun": "Lustknabe",
    "article": "der"
  },
  {
    "noun": "Lustmolch",
    "article": "der"
  },
  {
    "noun": "Lutheraner",
    "article": "der"
  },
  {
    "noun": "Lutscher",
    "article": "der"
  },
  {
    "noun": "Luxus",
    "article": "der"
  },
  {
    "noun": "Luxusartikel",
    "article": "der"
  },
  {
    "noun": "Luxusdampfer",
    "article": "der"
  },
  {
    "noun": "Luxuswagen",
    "article": "der"
  },
  {
    "noun": "Lymphknoten",
    "article": "der"
  },
  {
    "noun": "Lymphozyt",
    "article": "der"
  },
  {
    "noun": "Lyriker",
    "article": "der"
  },
  {
    "noun": "Lyrismus",
    "article": "der"
  },
  {
    "noun": "Velar",
    "article": "der"
  },
  {
    "noun": "Velarlaut",
    "article": "der"
  },
  {
    "noun": "Venezianer",
    "article": "der"
  },
  {
    "noun": "Venezolaner",
    "article": "der"
  },
  {
    "noun": "Venezueler",
    "article": "der"
  },
  {
    "noun": "Ventilator",
    "article": "der"
  },
  {
    "noun": "Ventrikel",
    "article": "der"
  },
  {
    "noun": "Venus",
    "article": "der"
  },
  {
    "noun": "Veranlasser",
    "article": "der"
  },
  {
    "noun": "Veranstalter",
    "article": "der"
  },
  {
    "noun": "Veranstaltungskalender",
    "article": "der"
  },
  {
    "noun": "Veranstaltungsraum",
    "article": "der"
  },
  {
    "noun": "Verantwortungsbereich",
    "article": "der"
  },
  {
    "noun": "Verbalerotiker",
    "article": "der"
  },
  {
    "noun": "Verbalismus",
    "article": "der"
  },
  {
    "noun": "Verband",
    "article": "der"
  },
  {
    "noun": "Verbandkasten",
    "article": "der"
  },
  {
    "noun": "Verbandsflug",
    "article": "der"
  },
  {
    "noun": "Verbandskasten",
    "article": "der"
  },
  {
    "noun": "Verbandsmull",
    "article": "der"
  },
  {
    "noun": "Verbandstoff",
    "article": "der"
  },
  {
    "noun": "Verbesserer",
    "article": "der"
  },
  {
    "noun": "Verbesserungsvorschlag",
    "article": "der"
  },
  {
    "noun": "Verbinder",
    "article": "der"
  },
  {
    "noun": "Verbindungsgang",
    "article": "der"
  },
  {
    "noun": "Verbindungsmann",
    "article": "der"
  },
  {
    "noun": "Verbindungsoffizier",
    "article": "der"
  },
  {
    "noun": "Verbindungspunkt",
    "article": "der"
  },
  {
    "noun": "Verbindungsstecker",
    "article": "der"
  },
  {
    "noun": "Verbindungsstrich",
    "article": "der"
  },
  {
    "noun": "Verbindungsweg",
    "article": "der"
  },
  {
    "noun": "Verbleib",
    "article": "der"
  },
  {
    "noun": "Verbrauch",
    "article": "der"
  },
  {
    "noun": "Verbraucher",
    "article": "der"
  },
  {
    "noun": "Verbraucherkredit",
    "article": "der"
  },
  {
    "noun": "Verbrauchermarkt",
    "article": "der"
  },
  {
    "noun": "Verbraucherpreis",
    "article": "der"
  },
  {
    "noun": "Verbraucherschutz",
    "article": "der"
  },
  {
    "noun": "Verbraucherverband",
    "article": "der"
  },
  {
    "noun": "Verbrauchswert",
    "article": "der"
  },
  {
    "noun": "Verbrecher",
    "article": "der"
  },
  {
    "noun": "Verbreiter",
    "article": "der"
  },
  {
    "noun": "Verbrennungsmotor",
    "article": "der"
  },
  {
    "noun": "Verbrennungsvorgang",
    "article": "der"
  },
  {
    "noun": "Verbund",
    "article": "der"
  },
  {
    "noun": "Verbundguss",
    "article": "der"
  },
  {
    "noun": "Verbundpflasterstein",
    "article": "der"
  },
  {
    "noun": "Verbundstoff",
    "article": "der"
  },
  {
    "noun": "Verbundwerkstoff",
    "article": "der"
  },
  {
    "noun": "Verdacht",
    "article": "der"
  },
  {
    "noun": "Verdachtsgrund",
    "article": "der"
  },
  {
    "noun": "Verdampfer",
    "article": "der"
  },
  {
    "noun": "Verdauungsapparat",
    "article": "der"
  },
  {
    "noun": "Verdauungskanal",
    "article": "der"
  },
  {
    "noun": "Verdauungstrakt",
    "article": "der"
  },
  {
    "noun": "Verderb",
    "article": "der"
  },
  {
    "noun": "Verderber",
    "article": "der"
  },
  {
    "noun": "Verdichter",
    "article": "der"
  },
  {
    "noun": "Verdiener",
    "article": "der"
  },
  {
    "noun": "Verdienstausfall",
    "article": "der"
  },
  {
    "noun": "Verdienstorden",
    "article": "der"
  },
  {
    "noun": "Verdruss",
    "article": "der"
  },
  {
    "noun": "Verdunstungsmesser",
    "article": "der"
  },
  {
    "noun": "Verehrer",
    "article": "der"
  },
  {
    "noun": "Verein",
    "article": "der"
  },
  {
    "noun": "Vereinsbeitrag",
    "article": "der"
  },
  {
    "noun": "Verfahrensingenieur",
    "article": "der"
  },
  {
    "noun": "Verfahrenstechniker",
    "article": "der"
  },
  {
    "noun": "Verfall",
    "article": "der"
  },
  {
    "noun": "Verfallstag",
    "article": "der"
  },
  {
    "noun": "Verfasser",
    "article": "der"
  },
  {
    "noun": "Verfassungsbruch",
    "article": "der"
  },
  {
    "noun": "Verfassungsschutz",
    "article": "der"
  },
  {
    "noun": "Verfassungsstaat",
    "article": "der"
  },
  {
    "noun": "Verfechter",
    "article": "der"
  },
  {
    "noun": "Verfertiger",
    "article": "der"
  },
  {
    "noun": "Verfolger",
    "article": "der"
  },
  {
    "noun": "Verfolgungswahn",
    "article": "der"
  },
  {
    "noun": "Verfremdungseffekt",
    "article": "der"
  },
  {
    "noun": "Vergeltungsakt",
    "article": "der"
  },
  {
    "noun": "Vergeltungsschlag",
    "article": "der"
  },
  {
    "noun": "Vergeuder",
    "article": "der"
  },
  {
    "noun": "Vergleich",
    "article": "der"
  },
  {
    "noun": "Vergleichspunkt",
    "article": "der"
  },
  {
    "noun": "Vergleichssatz",
    "article": "der"
  },
  {
    "noun": "Vergleichswert",
    "article": "der"
  },
  {
    "noun": "Vergolder",
    "article": "der"
  },
  {
    "noun": "Verhaftungsbefehl",
    "article": "der"
  },
  {
    "noun": "Verhaltensforscher",
    "article": "der"
  },
  {
    "noun": "Verhaltenskodex",
    "article": "der"
  },
  {
    "noun": "Verhandlungsort",
    "article": "der"
  },
  {
    "noun": "Verhandlungstag",
    "article": "der"
  },
  {
    "noun": "Verkauf",
    "article": "der"
  },
  {
    "noun": "Verkaufsartikel",
    "article": "der"
  },
  {
    "noun": "Verkaufsautomat",
    "article": "der"
  },
  {
    "noun": "Verkaufsleiter",
    "article": "der"
  },
  {
    "noun": "Verkaufspreis",
    "article": "der"
  },
  {
    "noun": "Verkaufsraum",
    "article": "der"
  },
  {
    "noun": "Verkaufsschlager",
    "article": "der"
  },
  {
    "noun": "Verkaufsstand",
    "article": "der"
  },
  {
    "noun": "Verkaufstisch",
    "article": "der"
  },
  {
    "noun": "Verkaufswert",
    "article": "der"
  },
  {
    "noun": "Verkehr",
    "article": "der"
  },
  {
    "noun": "Verkehrsbetrieb",
    "article": "der"
  },
  {
    "noun": "Verkehrsfluss",
    "article": "der"
  },
  {
    "noun": "Verkehrsfunk",
    "article": "der"
  },
  {
    "noun": "Verkehrsknotenpunkt",
    "article": "der"
  },
  {
    "noun": "Verkehrskollaps",
    "article": "der"
  },
  {
    "noun": "Verkehrsminister",
    "article": "der"
  },
  {
    "noun": "Verkehrspolizist",
    "article": "der"
  },
  {
    "noun": "Verkehrsrowdy",
    "article": "der"
  },
  {
    "noun": "Verkehrsstau",
    "article": "der"
  },
  {
    "noun": "Verkehrsstrom",
    "article": "der"
  },
  {
    "noun": "Verkehrsteilnehmer",
    "article": "der"
  },
  {
    "noun": "Verkehrsunfall",
    "article": "der"
  },
  {
    "noun": "Verkehrsunterricht",
    "article": "der"
  },
  {
    "noun": "Verkehrsverbund",
    "article": "der"
  },
  {
    "noun": "Verkehrsverein",
    "article": "der"
  },
  {
    "noun": "Verkehrsweg",
    "article": "der"
  },
  {
    "noun": "Verkehrswert",
    "article": "der"
  },
  {
    "noun": "Verlad",
    "article": "der"
  },
  {
    "noun": "Verladekran",
    "article": "der"
  },
  {
    "noun": "Verladeplatz",
    "article": "der"
  },
  {
    "noun": "Verlader",
    "article": "der"
  },
  {
    "noun": "Verlag",
    "article": "der"
  },
  {
    "noun": "Verlagsbuchhandel",
    "article": "der"
  },
  {
    "noun": "Verlagskatalog",
    "article": "der"
  },
  {
    "noun": "Verlauf",
    "article": "der"
  },
  {
    "noun": "Verleger",
    "article": "der"
  },
  {
    "noun": "Verleih",
    "article": "der"
  },
  {
    "noun": "Verleiher",
    "article": "der"
  },
  {
    "noun": "Verletzer",
    "article": "der"
  },
  {
    "noun": "Verleumder",
    "article": "der"
  },
  {
    "noun": "Verlierer",
    "article": "der"
  },
  {
    "noun": "Verlobungsring",
    "article": "der"
  },
  {
    "noun": "Verlust",
    "article": "der"
  },
  {
    "noun": "Verlustausgleich",
    "article": "der"
  },
  {
    "noun": "Vermerk",
    "article": "der"
  },
  {
    "noun": "Vermesser",
    "article": "der"
  },
  {
    "noun": "Vermessungsingenieur",
    "article": "der"
  },
  {
    "noun": "Vermieter",
    "article": "der"
  },
  {
    "noun": "Vermittler",
    "article": "der"
  },
  {
    "noun": "Vermittlungsausschuss",
    "article": "der"
  },
  {
    "noun": "Vermittlungsversuch",
    "article": "der"
  },
  {
    "noun": "Vernichtungskrieg",
    "article": "der"
  },
  {
    "noun": "Vernunftschluss",
    "article": "der"
  },
  {
    "noun": "Verputz",
    "article": "der"
  },
  {
    "noun": "Verrat",
    "article": "der"
  },
  {
    "noun": "Verrechnungspreis",
    "article": "der"
  },
  {
    "noun": "Verrechnungsscheck",
    "article": "der"
  },
  {
    "noun": "Verriss",
    "article": "der"
  },
  {
    "noun": "Verruf",
    "article": "der"
  },
  {
    "noun": "Vers",
    "article": "der"
  },
  {
    "noun": "Versager",
    "article": "der"
  },
  {
    "noun": "Versalbuchstabe",
    "article": "der"
  },
  {
    "noun": "Versammlungsort",
    "article": "der"
  },
  {
    "noun": "Versammlungsraum",
    "article": "der"
  },
  {
    "noun": "Versammlungssaal",
    "article": "der"
  },
  {
    "noun": "Versand",
    "article": "der"
  },
  {
    "noun": "Versandhandel",
    "article": "der"
  },
  {
    "noun": "Versatz",
    "article": "der"
  },
  {
    "noun": "Versbau",
    "article": "der"
  },
  {
    "noun": "Verschiebebahnhof",
    "article": "der"
  },
  {
    "noun": "Verschiffungshafen",
    "article": "der"
  },
  {
    "noun": "Verschlag",
    "article": "der"
  },
  {
    "noun": "Verschluss",
    "article": "der"
  },
  {
    "noun": "Verschlussdeckel",
    "article": "der"
  },
  {
    "noun": "Verschlusslaut",
    "article": "der"
  },
  {
    "noun": "Verschmutzer",
    "article": "der"
  },
  {
    "noun": "Verschnitt",
    "article": "der"
  },
  {
    "noun": "Verschwender",
    "article": "der"
  },
  {
    "noun": "Versemacher",
    "article": "der"
  },
  {
    "noun": "Versicherer",
    "article": "der"
  },
  {
    "noun": "Versicherungsanspruch",
    "article": "der"
  },
  {
    "noun": "Versicherungsbeitrag",
    "article": "der"
  },
  {
    "noun": "Versicherungsbetrug",
    "article": "der"
  },
  {
    "noun": "Versicherungsfall",
    "article": "der"
  },
  {
    "noun": "Versicherungskaufmann",
    "article": "der"
  },
  {
    "noun": "Versicherungsmakler",
    "article": "der"
  },
  {
    "noun": "Versicherungsmathematiker",
    "article": "der"
  },
  {
    "noun": "Versicherungsnehmer",
    "article": "der"
  },
  {
    "noun": "Versicherungsschein",
    "article": "der"
  },
  {
    "noun": "Versicherungsschutz",
    "article": "der"
  },
  {
    "noun": "Versicherungsvertrag",
    "article": "der"
  },
  {
    "noun": "Versicherungsvertreter",
    "article": "der"
  },
  {
    "noun": "Versicherungswert",
    "article": "der"
  },
  {
    "noun": "Versorger",
    "article": "der"
  },
  {
    "noun": "Versorgungsanspruch",
    "article": "der"
  },
  {
    "noun": "Versorgungsausgleich",
    "article": "der"
  },
  {
    "noun": "Versorgungsbetrieb",
    "article": "der"
  },
  {
    "noun": "Versorgungsengpass",
    "article": "der"
  },
  {
    "noun": "Versprecher",
    "article": "der"
  },
  {
    "noun": "Verstand",
    "article": "der"
  },
  {
    "noun": "Verstandesmensch",
    "article": "der"
  },
  {
    "noun": "Versteigerer",
    "article": "der"
  },
  {
    "noun": "Versteller",
    "article": "der"
  },
  {
    "noun": "Versuch",
    "article": "der"
  },
  {
    "noun": "Versucher",
    "article": "der"
  },
  {
    "noun": "Versuchsballon",
    "article": "der"
  },
  {
    "noun": "Versuchsstand",
    "article": "der"
  },
  {
    "noun": "Vertebrat",
    "article": "der"
  },
  {
    "noun": "Verteidiger",
    "article": "der"
  },
  {
    "noun": "Verteidigungskrieg",
    "article": "der"
  },
  {
    "noun": "Verteidigungsminister",
    "article": "der"
  },
  {
    "noun": "Verteidigungspakt",
    "article": "der"
  },
  {
    "noun": "Verteidigungszustand",
    "article": "der"
  },
  {
    "noun": "Verteiler",
    "article": "der"
  },
  {
    "noun": "Verteilerkasten",
    "article": "der"
  },
  {
    "noun": "Verteilungskampf",
    "article": "der"
  },
  {
    "noun": "Vertrag",
    "article": "der"
  },
  {
    "noun": "Vertragsabschluss",
    "article": "der"
  },
  {
    "noun": "Vertragsbruch",
    "article": "der"
  },
  {
    "noun": "Vertragsentwurf",
    "article": "der"
  },
  {
    "noun": "Vertragspartner",
    "article": "der"
  },
  {
    "noun": "Vertragspunkt",
    "article": "der"
  },
  {
    "noun": "Vertragsschluss",
    "article": "der"
  },
  {
    "noun": "Vertragsstaat",
    "article": "der"
  },
  {
    "noun": "Vertragstext",
    "article": "der"
  },
  {
    "noun": "Vertrauensarzt",
    "article": "der"
  },
  {
    "noun": "Vertrauensbeweis",
    "article": "der"
  },
  {
    "noun": "Vertrauensbruch",
    "article": "der"
  },
  {
    "noun": "Vertrauensmann",
    "article": "der"
  },
  {
    "noun": "Vertrauensvorschuss",
    "article": "der"
  },
  {
    "noun": "Vertreiber",
    "article": "der"
  },
  {
    "noun": "Vertreter",
    "article": "der"
  },
  {
    "noun": "Vertrieb",
    "article": "der"
  },
  {
    "noun": "Vertriebsleiter",
    "article": "der"
  },
  {
    "noun": "Vertriebsweg",
    "article": "der"
  },
  {
    "noun": "Veruntreuer",
    "article": "der"
  },
  {
    "noun": "Verursacher",
    "article": "der"
  },
  {
    "noun": "Vervielfacher",
    "article": "der"
  },
  {
    "noun": "Verwahrer",
    "article": "der"
  },
  {
    "noun": "Verwalter",
    "article": "der"
  },
  {
    "noun": "Verwaltungsakt",
    "article": "der"
  },
  {
    "noun": "Verwaltungsapparat",
    "article": "der"
  },
  {
    "noun": "Verwaltungsaufwand",
    "article": "der"
  },
  {
    "noun": "Verwaltungsausschuss",
    "article": "der"
  },
  {
    "noun": "Verwaltungsbezirk",
    "article": "der"
  },
  {
    "noun": "Verwaltungsdienst",
    "article": "der"
  },
  {
    "noun": "Verwaltungsgerichtshof",
    "article": "der"
  },
  {
    "noun": "Verwaltungsrat",
    "article": "der"
  },
  {
    "noun": "Verwandtschaftsgrad",
    "article": "der"
  },
  {
    "noun": "Verweigerer",
    "article": "der"
  },
  {
    "noun": "Verweis",
    "article": "der"
  },
  {
    "noun": "Verwendungszweck",
    "article": "der"
  },
  {
    "noun": "Verweser",
    "article": "der"
  },
  {
    "noun": "Verwundetentransport",
    "article": "der"
  },
  {
    "noun": "Verzehr",
    "article": "der"
  },
  {
    "noun": "Verzehrer",
    "article": "der"
  },
  {
    "noun": "Verzerrer",
    "article": "der"
  },
  {
    "noun": "Verzicht",
    "article": "der"
  },
  {
    "noun": "Verzug",
    "article": "der"
  },
  {
    "noun": "Verzugszins",
    "article": "der"
  },
  {
    "noun": "Veteran",
    "article": "der"
  },
  {
    "noun": "Vetter",
    "article": "der"
  },
  {
    "noun": "Vexierspiegel",
    "article": "der"
  },
  {
    "noun": "Vibrator",
    "article": "der"
  },
  {
    "noun": "Vicomte",
    "article": "der"
  },
  {
    "noun": "Videoclip",
    "article": "der"
  },
  {
    "noun": "Videorecorder",
    "article": "der"
  },
  {
    "noun": "Videotext",
    "article": "der"
  },
  {
    "noun": "Viehbestand",
    "article": "der"
  },
  {
    "noun": "Viehhandel",
    "article": "der"
  },
  {
    "noun": "Viehhirt",
    "article": "der"
  },
  {
    "noun": "Viehmarkt",
    "article": "der"
  },
  {
    "noun": "Viehstall",
    "article": "der"
  },
  {
    "noun": "Viehtreiber",
    "article": "der"
  },
  {
    "noun": "Viehwagen",
    "article": "der"
  },
  {
    "noun": "Vielschreiber",
    "article": "der"
  },
  {
    "noun": "Vielstoffmotor",
    "article": "der"
  },
  {
    "noun": "Vielwisser",
    "article": "der"
  },
  {
    "noun": "Vielzeller",
    "article": "der"
  },
  {
    "noun": "Vierbeiner",
    "article": "der"
  },
  {
    "noun": "Vierer",
    "article": "der"
  },
  {
    "noun": "Viererbob",
    "article": "der"
  },
  {
    "noun": "Viererzug",
    "article": "der"
  },
  {
    "noun": "Vierfarbendruck",
    "article": "der"
  },
  {
    "noun": "Vierling",
    "article": "der"
  },
  {
    "noun": "Vierpass",
    "article": "der"
  },
  {
    "noun": "Vierpol",
    "article": "der"
  },
  {
    "noun": "Vierradantrieb",
    "article": "der"
  },
  {
    "noun": "Viersitzer",
    "article": "der"
  },
  {
    "noun": "Viertakter",
    "article": "der"
  },
  {
    "noun": "Viertaktmotor",
    "article": "der"
  },
  {
    "noun": "Viertelfinalist",
    "article": "der"
  },
  {
    "noun": "Viertelkreis",
    "article": "der"
  },
  {
    "noun": "Viertelstab",
    "article": "der"
  },
  {
    "noun": "Viervierteltakt",
    "article": "der"
  },
  {
    "noun": "Vierzeiler",
    "article": "der"
  },
  {
    "noun": "Vierzylindermotor",
    "article": "der"
  },
  {
    "noun": "Vietnamese",
    "article": "der"
  },
  {
    "noun": "Vikar",
    "article": "der"
  },
  {
    "noun": "Violinist",
    "article": "der"
  },
  {
    "noun": "Virilismus",
    "article": "der"
  },
  {
    "noun": "Virologe",
    "article": "der"
  },
  {
    "noun": "Virtuose",
    "article": "der"
  },
  {
    "noun": "Visagist",
    "article": "der"
  },
  {
    "noun": "Vitaminmangel",
    "article": "der"
  },
  {
    "noun": "Vizeadmiral",
    "article": "der"
  },
  {
    "noun": "Vizekanzler",
    "article": "der"
  },
  {
    "noun": "Vizekonsul",
    "article": "der"
  },
  {
    "noun": "Vliesstoff",
    "article": "der"
  },
  {
    "noun": "Vocoder",
    "article": "der"
  },
  {
    "noun": "Vogel",
    "article": "der"
  },
  {
    "noun": "Vogelbeerbaum",
    "article": "der"
  },
  {
    "noun": "Vogeldreck",
    "article": "der"
  },
  {
    "noun": "Vogeldunst",
    "article": "der"
  },
  {
    "noun": "Vogelfang",
    "article": "der"
  },
  {
    "noun": "Vogelflug",
    "article": "der"
  },
  {
    "noun": "Vogelgesang",
    "article": "der"
  },
  {
    "noun": "Vogelherd",
    "article": "der"
  },
  {
    "noun": "Vogelkundler",
    "article": "der"
  },
  {
    "noun": "Vogelmist",
    "article": "der"
  },
  {
    "noun": "Vogelruf",
    "article": "der"
  },
  {
    "noun": "Vogelschlag",
    "article": "der"
  },
  {
    "noun": "Vogelschutz",
    "article": "der"
  },
  {
    "noun": "Vogelschwarm",
    "article": "der"
  },
  {
    "noun": "Vogelsteller",
    "article": "der"
  },
  {
    "noun": "Vogelzug",
    "article": "der"
  },
  {
    "noun": "Vogt",
    "article": "der"
  },
  {
    "noun": "Voile",
    "article": "der"
  },
  {
    "noun": "Vokal",
    "article": "der"
  },
  {
    "noun": "Vokalismus",
    "article": "der"
  },
  {
    "noun": "Vokalist",
    "article": "der"
  },
  {
    "noun": "Vokativ",
    "article": "der"
  },
  {
    "noun": "Volksarmist",
    "article": "der"
  },
  {
    "noun": "Volksaufstand",
    "article": "der"
  },
  {
    "noun": "Volksbrauch",
    "article": "der"
  },
  {
    "noun": "Volkscharakter",
    "article": "der"
  },
  {
    "noun": "Volksdichter",
    "article": "der"
  },
  {
    "noun": "Volksentscheid",
    "article": "der"
  },
  {
    "noun": "Volksfeind",
    "article": "der"
  },
  {
    "noun": "Volksgenosse",
    "article": "der"
  },
  {
    "noun": "Volksglaube",
    "article": "der"
  },
  {
    "noun": "Volksheld",
    "article": "der"
  },
  {
    "noun": "Volkskundler",
    "article": "der"
  },
  {
    "noun": "Volkslauf",
    "article": "der"
  },
  {
    "noun": "Volksmund",
    "article": "der"
  },
  {
    "noun": "Volkspolizist",
    "article": "der"
  },
  {
    "noun": "Volksredner",
    "article": "der"
  },
  {
    "noun": "Volksschullehrer",
    "article": "der"
  },
  {
    "noun": "Volkssport",
    "article": "der"
  },
  {
    "noun": "Volksstamm",
    "article": "der"
  },
  {
    "noun": "Volkstanz",
    "article": "der"
  },
  {
    "noun": "Volkstrauertag",
    "article": "der"
  },
  {
    "noun": "Volksvertreter",
    "article": "der"
  },
  {
    "noun": "Volkswirt",
    "article": "der"
  },
  {
    "noun": "Volkswirtschaftler",
    "article": "der"
  },
  {
    "noun": "Volkszorn",
    "article": "der"
  },
  {
    "noun": "Vollautomat",
    "article": "der"
  },
  {
    "noun": "Vollbart",
    "article": "der"
  },
  {
    "noun": "Vollbesitz",
    "article": "der"
  },
  {
    "noun": "Volldampf",
    "article": "der"
  },
  {
    "noun": "Vollender",
    "article": "der"
  },
  {
    "noun": "Volleyball",
    "article": "der"
  },
  {
    "noun": "Volleyballspieler",
    "article": "der"
  },
  {
    "noun": "Vollgummireifen",
    "article": "der"
  },
  {
    "noun": "Vollidiot",
    "article": "der"
  },
  {
    "noun": "Volljurist",
    "article": "der"
  },
  {
    "noun": "Vollkasko",
    "article": "der"
  },
  {
    "noun": "Vollmachtgeber",
    "article": "der"
  },
  {
    "noun": "Vollmatrose",
    "article": "der"
  },
  {
    "noun": "Vollmond",
    "article": "der"
  },
  {
    "noun": "Vollname",
    "article": "der"
  },
  {
    "noun": "Vollrausch",
    "article": "der"
  },
  {
    "noun": "Vollreifen",
    "article": "der"
  },
  {
    "noun": "Vollstrecker",
    "article": "der"
  },
  {
    "noun": "Vollstreckungsbefehl",
    "article": "der"
  },
  {
    "noun": "Vollstreckungstitel",
    "article": "der"
  },
  {
    "noun": "Volltreffer",
    "article": "der"
  },
  {
    "noun": "Vollzug",
    "article": "der"
  },
  {
    "noun": "Volvulus",
    "article": "der"
  },
  {
    "noun": "Voodoo",
    "article": "der"
  },
  {
    "noun": "Vorabdruck",
    "article": "der"
  },
  {
    "noun": "Vorabend",
    "article": "der"
  },
  {
    "noun": "Voranschlag",
    "article": "der"
  },
  {
    "noun": "Vorarbeiter",
    "article": "der"
  },
  {
    "noun": "Vorbau",
    "article": "der"
  },
  {
    "noun": "Vorbedacht",
    "article": "der"
  },
  {
    "noun": "Vorbehalt",
    "article": "der"
  },
  {
    "noun": "Vorbeimarsch",
    "article": "der"
  },
  {
    "noun": "Vorbereiter",
    "article": "der"
  },
  {
    "noun": "Vorbereitungsdienst",
    "article": "der"
  },
  {
    "noun": "Vorbereitungskurs",
    "article": "der"
  },
  {
    "noun": "Vorbericht",
    "article": "der"
  },
  {
    "noun": "Vorbescheid",
    "article": "der"
  },
  {
    "noun": "Vorbesitzer",
    "article": "der"
  },
  {
    "noun": "Vorblick",
    "article": "der"
  },
  {
    "noun": "Vorbote",
    "article": "der"
  },
  {
    "noun": "Vordenker",
    "article": "der"
  },
  {
    "noun": "Vorderarm",
    "article": "der"
  },
  {
    "noun": "Vorderdarm",
    "article": "der"
  },
  {
    "noun": "Vordergaumen",
    "article": "der"
  },
  {
    "noun": "Vordergrund",
    "article": "der"
  },
  {
    "noun": "Vorderlader",
    "article": "der"
  },
  {
    "noun": "Vorderlauf",
    "article": "der"
  },
  {
    "noun": "Vordermann",
    "article": "der"
  },
  {
    "noun": "Vorderradantrieb",
    "article": "der"
  },
  {
    "noun": "Vordersatz",
    "article": "der"
  },
  {
    "noun": "Vorderschinken",
    "article": "der"
  },
  {
    "noun": "Vordersitz",
    "article": "der"
  },
  {
    "noun": "Vordersteven",
    "article": "der"
  },
  {
    "noun": "Vorderzahn",
    "article": "der"
  },
  {
    "noun": "Vordruck",
    "article": "der"
  },
  {
    "noun": "Vorentscheidungslauf",
    "article": "der"
  },
  {
    "noun": "Vorentwurf",
    "article": "der"
  },
  {
    "noun": "Vorfahr",
    "article": "der"
  },
  {
    "noun": "Vorfall",
    "article": "der"
  },
  {
    "noun": "Vorfilm",
    "article": "der"
  },
  {
    "noun": "Vorfluter",
    "article": "der"
  },
  {
    "noun": "Vorgang",
    "article": "der"
  },
  {
    "noun": "Vorgarten",
    "article": "der"
  },
  {
    "noun": "Vorgeschmack",
    "article": "der"
  },
  {
    "noun": "Vorgriff",
    "article": "der"
  },
  {
    "noun": "Vorhafen",
    "article": "der"
  },
  {
    "noun": "Vorhalt",
    "article": "der"
  },
  {
    "noun": "Vorhang",
    "article": "der"
  },
  {
    "noun": "Vorhof",
    "article": "der"
  },
  {
    "noun": "Vorkauf",
    "article": "der"
  },
  {
    "noun": "Vorkoster",
    "article": "der"
  },
  {
    "noun": "Vorlauf",
    "article": "der"
  },
  {
    "noun": "Vorleger",
    "article": "der"
  },
  {
    "noun": "Vorleser",
    "article": "der"
  },
  {
    "noun": "Vorlesewettbewerb",
    "article": "der"
  },
  {
    "noun": "Vormagen",
    "article": "der"
  },
  {
    "noun": "Vormann",
    "article": "der"
  },
  {
    "noun": "Vormarsch",
    "article": "der"
  },
  {
    "noun": "Vormast",
    "article": "der"
  },
  {
    "noun": "Vormittag",
    "article": "der"
  },
  {
    "noun": "Vormonat",
    "article": "der"
  },
  {
    "noun": "Vormund",
    "article": "der"
  },
  {
    "noun": "Vorname",
    "article": "der"
  },
  {
    "noun": "Vorort",
    "article": "der"
  },
  {
    "noun": "Vorortsverkehr",
    "article": "der"
  },
  {
    "noun": "Vorortszug",
    "article": "der"
  },
  {
    "noun": "Vorortzug",
    "article": "der"
  },
  {
    "noun": "Vorplatz",
    "article": "der"
  },
  {
    "noun": "Vorposten",
    "article": "der"
  },
  {
    "noun": "Vorrang",
    "article": "der"
  },
  {
    "noun": "Vorrat",
    "article": "der"
  },
  {
    "noun": "Vorratskeller",
    "article": "der"
  },
  {
    "noun": "Vorratsraum",
    "article": "der"
  },
  {
    "noun": "Vorraum",
    "article": "der"
  },
  {
    "noun": "Vorredner",
    "article": "der"
  },
  {
    "noun": "Vorreiter",
    "article": "der"
  },
  {
    "noun": "Vorruhestand",
    "article": "der"
  },
  {
    "noun": "Vorsaal",
    "article": "der"
  },
  {
    "noun": "Vorsager",
    "article": "der"
  },
  {
    "noun": "Vorsatz",
    "article": "der"
  },
  {
    "noun": "Vorschaltwiderstand",
    "article": "der"
  },
  {
    "noun": "Vorschlag",
    "article": "der"
  },
  {
    "noun": "Vorschlaghammer",
    "article": "der"
  },
  {
    "noun": "Vorschoter",
    "article": "der"
  },
  {
    "noun": "Vorschub",
    "article": "der"
  },
  {
    "noun": "Vorschuss",
    "article": "der"
  },
  {
    "noun": "Vorschusslorbeer",
    "article": "der"
  },
  {
    "noun": "Vorsitz",
    "article": "der"
  },
  {
    "noun": "Vorsitzer",
    "article": "der"
  },
  {
    "noun": "Vorsommer",
    "article": "der"
  },
  {
    "noun": "Vorspann",
    "article": "der"
  },
  {
    "noun": "Vorspruch",
    "article": "der"
  },
  {
    "noun": "Vorsprung",
    "article": "der"
  },
  {
    "noun": "Vorstand",
    "article": "der"
  },
  {
    "noun": "Vorstandstisch",
    "article": "der"
  },
  {
    "noun": "Vorstecker",
    "article": "der"
  },
  {
    "noun": "Vorsteher",
    "article": "der"
  },
  {
    "noun": "Vorstehhund",
    "article": "der"
  },
  {
    "noun": "Vorstellungsbeginn",
    "article": "der"
  },
  {
    "noun": "Vorstopper",
    "article": "der"
  },
  {
    "noun": "Vortag",
    "article": "der"
  },
  {
    "noun": "Vorteil",
    "article": "der"
  },
  {
    "noun": "Vortrab",
    "article": "der"
  },
  {
    "noun": "Vortrag",
    "article": "der"
  },
  {
    "noun": "Vortragsabend",
    "article": "der"
  },
  {
    "noun": "Vortragsraum",
    "article": "der"
  },
  {
    "noun": "Vortragssaal",
    "article": "der"
  },
  {
    "noun": "Vortrieb",
    "article": "der"
  },
  {
    "noun": "Vortritt",
    "article": "der"
  },
  {
    "noun": "Vortrupp",
    "article": "der"
  },
  {
    "noun": "Vorturner",
    "article": "der"
  },
  {
    "noun": "Vorvater",
    "article": "der"
  },
  {
    "noun": "Vorverkauf",
    "article": "der"
  },
  {
    "noun": "Vorversuch",
    "article": "der"
  },
  {
    "noun": "Vorvertrag",
    "article": "der"
  },
  {
    "noun": "Vorwand",
    "article": "der"
  },
  {
    "noun": "Vorweis",
    "article": "der"
  },
  {
    "noun": "Vorwitz",
    "article": "der"
  },
  {
    "noun": "Vorwurf",
    "article": "der"
  },
  {
    "noun": "Vorzeitmensch",
    "article": "der"
  },
  {
    "noun": "Vorzug",
    "article": "der"
  },
  {
    "noun": "Vorzugspreis",
    "article": "der"
  },
  {
    "noun": "Votant",
    "article": "der"
  },
  {
    "noun": "Voyeur",
    "article": "der"
  },
  {
    "noun": "Vulkan",
    "article": "der"
  },
  {
    "noun": "Vulkanausbruch",
    "article": "der"
  },
  {
    "noun": "Vulkaniseur",
    "article": "der"
  },
  {
    "noun": "Vulkanismus",
    "article": "der"
  },
  {
    "noun": "Vulkanologe",
    "article": "der"
  },
  {
    "noun": "Wagenfond",
    "article": "der"
  },
  {
    "noun": "Wagenheber",
    "article": "der"
  },
  {
    "noun": "Wagenlenker",
    "article": "der"
  },
  {
    "noun": "Wagenpark",
    "article": "der"
  },
  {
    "noun": "Wagenschlag",
    "article": "der"
  },
  {
    "noun": "Wagentyp",
    "article": "der"
  },
  {
    "noun": "Waggon",
    "article": "der"
  },
  {
    "noun": "Wahlaufruf",
    "article": "der"
  },
  {
    "noun": "Wahlausgang",
    "article": "der"
  },
  {
    "noun": "Wahlausschuss",
    "article": "der"
  },
  {
    "noun": "Wahlbetrug",
    "article": "der"
  },
  {
    "noun": "Wahlbezirk",
    "article": "der"
  },
  {
    "noun": "Wahlgang",
    "article": "der"
  },
  {
    "noun": "Wahlhelfer",
    "article": "der"
  },
  {
    "noun": "Wahlkampf",
    "article": "der"
  },
  {
    "noun": "Wahlkreis",
    "article": "der"
  },
  {
    "noun": "Wahlleiter",
    "article": "der"
  },
  {
    "noun": "Wahlmann",
    "article": "der"
  },
  {
    "noun": "Wahlredner",
    "article": "der"
  },
  {
    "noun": "Wahlsieg",
    "article": "der"
  },
  {
    "noun": "Wahlspruch",
    "article": "der"
  },
  {
    "noun": "Wahlvorschlag",
    "article": "der"
  },
  {
    "noun": "Wahlzettel",
    "article": "der"
  },
  {
    "noun": "Wahn",
    "article": "der"
  },
  {
    "noun": "Wahnsinn",
    "article": "der"
  },
  {
    "noun": "Wahnwitz",
    "article": "der"
  },
  {
    "noun": "Wahrheitsbeweis",
    "article": "der"
  },
  {
    "noun": "Wahrheitsgehalt",
    "article": "der"
  },
  {
    "noun": "Wahrheitssinn",
    "article": "der"
  },
  {
    "noun": "Wahrsager",
    "article": "der"
  },
  {
    "noun": "Wahrscheinlichkeitsgrad",
    "article": "der"
  },
  {
    "noun": "Wahrspruch",
    "article": "der"
  },
  {
    "noun": "Waid",
    "article": "der"
  },
  {
    "noun": "Waidmann",
    "article": "der"
  },
  {
    "noun": "Waidmesser",
    "article": "der"
  },
  {
    "noun": "Wal",
    "article": "der"
  },
  {
    "noun": "Wald",
    "article": "der"
  },
  {
    "noun": "Waldarbeiter",
    "article": "der"
  },
  {
    "noun": "Waldbau",
    "article": "der"
  },
  {
    "noun": "Waldbestand",
    "article": "der"
  },
  {
    "noun": "Waldboden",
    "article": "der"
  },
  {
    "noun": "Waldbrand",
    "article": "der"
  },
  {
    "noun": "Waldkauz",
    "article": "der"
  },
  {
    "noun": "Waldmeister",
    "article": "der"
  },
  {
    "noun": "Waldorfsalat",
    "article": "der"
  },
  {
    "noun": "Waldrand",
    "article": "der"
  },
  {
    "noun": "Waldreichtum",
    "article": "der"
  },
  {
    "noun": "Waldsaum",
    "article": "der"
  },
  {
    "noun": "Waldschaden",
    "article": "der"
  },
  {
    "noun": "Waldschrat",
    "article": "der"
  },
  {
    "noun": "Waldstorch",
    "article": "der"
  },
  {
    "noun": "Waldvogel",
    "article": "der"
  },
  {
    "noun": "Waldweg",
    "article": "der"
  },
  {
    "noun": "Walfang",
    "article": "der"
  },
  {
    "noun": "Walfisch",
    "article": "der"
  },
  {
    "noun": "Walker",
    "article": "der"
  },
  {
    "noun": "Wall",
    "article": "der"
  },
  {
    "noun": "Wallach",
    "article": "der"
  },
  {
    "noun": "Wallfahrer",
    "article": "der"
  },
  {
    "noun": "Wallfahrtsort",
    "article": "der"
  },
  {
    "noun": "Wallgraben",
    "article": "der"
  },
  {
    "noun": "Wallone",
    "article": "der"
  },
  {
    "noun": "Walm",
    "article": "der"
  },
  {
    "noun": "Walnussbaum",
    "article": "der"
  },
  {
    "noun": "Walzdraht",
    "article": "der"
  },
  {
    "noun": "Walzstahl",
    "article": "der"
  },
  {
    "noun": "Wandalismus",
    "article": "der"
  },
  {
    "noun": "Wandarm",
    "article": "der"
  },
  {
    "noun": "Wandbehang",
    "article": "der"
  },
  {
    "noun": "Wandel",
    "article": "der"
  },
  {
    "noun": "Wandelgang",
    "article": "der"
  },
  {
    "noun": "Wandelstern",
    "article": "der"
  },
  {
    "noun": "Wanderarbeiter",
    "article": "der"
  },
  {
    "noun": "Wanderbursche",
    "article": "der"
  },
  {
    "noun": "Wandercircus",
    "article": "der"
  },
  {
    "noun": "Wanderer",
    "article": "der"
  },
  {
    "noun": "Wanderfalke",
    "article": "der"
  },
  {
    "noun": "Wanderfisch",
    "article": "der"
  },
  {
    "noun": "Wandergeselle",
    "article": "der"
  },
  {
    "noun": "Wanderpokal",
    "article": "der"
  },
  {
    "noun": "Wanderprediger",
    "article": "der"
  },
  {
    "noun": "Wanderpreis",
    "article": "der"
  },
  {
    "noun": "Wanderschuh",
    "article": "der"
  },
  {
    "noun": "Wandersmann",
    "article": "der"
  },
  {
    "noun": "Wanderstab",
    "article": "der"
  },
  {
    "noun": "Wanderstock",
    "article": "der"
  },
  {
    "noun": "Wandertrieb",
    "article": "der"
  },
  {
    "noun": "Wanderverein",
    "article": "der"
  },
  {
    "noun": "Wandervogel",
    "article": "der"
  },
  {
    "noun": "Wanderweg",
    "article": "der"
  },
  {
    "noun": "Wandkalender",
    "article": "der"
  },
  {
    "noun": "Wandler",
    "article": "der"
  },
  {
    "noun": "Wandleuchter",
    "article": "der"
  },
  {
    "noun": "Wandlungsprozess",
    "article": "der"
  },
  {
    "noun": "Wandpfeiler",
    "article": "der"
  },
  {
    "noun": "Wandschirm",
    "article": "der"
  },
  {
    "noun": "Wandschrank",
    "article": "der"
  },
  {
    "noun": "Wandspiegel",
    "article": "der"
  },
  {
    "noun": "Wandteller",
    "article": "der"
  },
  {
    "noun": "Wandteppich",
    "article": "der"
  },
  {
    "noun": "Wankelmotor",
    "article": "der"
  },
  {
    "noun": "Wankelmut",
    "article": "der"
  },
  {
    "noun": "Wanst",
    "article": "der"
  },
  {
    "noun": "Wapiti",
    "article": "der"
  },
  {
    "noun": "Wappenspruch",
    "article": "der"
  },
  {
    "noun": "Waran",
    "article": "der"
  },
  {
    "noun": "Warenabsatz",
    "article": "der"
  },
  {
    "noun": "Warenaufzug",
    "article": "der"
  },
  {
    "noun": "Warenausgang",
    "article": "der"
  },
  {
    "noun": "Warenaustausch",
    "article": "der"
  },
  {
    "noun": "Warenautomat",
    "article": "der"
  },
  {
    "noun": "Warenbegleitschein",
    "article": "der"
  },
  {
    "noun": "Warenbestand",
    "article": "der"
  },
  {
    "noun": "Wareneingang",
    "article": "der"
  },
  {
    "noun": "Warenexport",
    "article": "der"
  },
  {
    "noun": "Warenhandel",
    "article": "der"
  },
  {
    "noun": "Warenimport",
    "article": "der"
  },
  {
    "noun": "Warenkorb",
    "article": "der"
  },
  {
    "noun": "Warenkredit",
    "article": "der"
  },
  {
    "noun": "Warenposten",
    "article": "der"
  },
  {
    "noun": "Warentest",
    "article": "der"
  },
  {
    "noun": "Warentransport",
    "article": "der"
  },
  {
    "noun": "Warenumsatz",
    "article": "der"
  },
  {
    "noun": "Warenumschlag",
    "article": "der"
  },
  {
    "noun": "Warenverkehr",
    "article": "der"
  },
  {
    "noun": "Warenwert",
    "article": "der"
  },
  {
    "noun": "Warenzeichenschutz",
    "article": "der"
  },
  {
    "noun": "Warenzoll",
    "article": "der"
  },
  {
    "noun": "Warmstart",
    "article": "der"
  },
  {
    "noun": "Warmwasserbereiter",
    "article": "der"
  },
  {
    "noun": "Warmwasserspeicher",
    "article": "der"
  },
  {
    "noun": "Warnblinker",
    "article": "der"
  },
  {
    "noun": "Warner",
    "article": "der"
  },
  {
    "noun": "Warnruf",
    "article": "der"
  },
  {
    "noun": "Warnschuss",
    "article": "der"
  },
  {
    "noun": "Warnstreik",
    "article": "der"
  },
  {
    "noun": "Warnton",
    "article": "der"
  },
  {
    "noun": "Warrant",
    "article": "der"
  },
  {
    "noun": "Wart",
    "article": "der"
  },
  {
    "noun": "Warteraum",
    "article": "der"
  },
  {
    "noun": "Wartesaal",
    "article": "der"
  },
  {
    "noun": "Wartungsdienst",
    "article": "der"
  },
  {
    "noun": "Wartungsvertrag",
    "article": "der"
  },
  {
    "noun": "Warzenhof",
    "article": "der"
  },
  {
    "noun": "Waschautomat",
    "article": "der"
  },
  {
    "noun": "Waschbottich",
    "article": "der"
  },
  {
    "noun": "Waschkorb",
    "article": "der"
  },
  {
    "noun": "Waschlappen",
    "article": "der"
  },
  {
    "noun": "Waschraum",
    "article": "der"
  },
  {
    "noun": "Waschsalon",
    "article": "der"
  },
  {
    "noun": "Waschtag",
    "article": "der"
  },
  {
    "noun": "Waschtisch",
    "article": "der"
  },
  {
    "noun": "Waschtrog",
    "article": "der"
  },
  {
    "noun": "Waschvollautomat",
    "article": "der"
  },
  {
    "noun": "Waschvorgang",
    "article": "der"
  },
  {
    "noun": "Waschzettel",
    "article": "der"
  },
  {
    "noun": "Wasen",
    "article": "der"
  },
  {
    "noun": "Wasserabfluss",
    "article": "der"
  },
  {
    "noun": "Wasseranschluss",
    "article": "der"
  },
  {
    "noun": "Wasserarm",
    "article": "der"
  },
  {
    "noun": "Wasserball",
    "article": "der"
  },
  {
    "noun": "Wasserbau",
    "article": "der"
  },
  {
    "noun": "Wasserbedarf",
    "article": "der"
  },
  {
    "noun": "Wasserbock",
    "article": "der"
  },
  {
    "noun": "Wasserbruch",
    "article": "der"
  },
  {
    "noun": "Wasserdampf",
    "article": "der"
  },
  {
    "noun": "Wasserdruck",
    "article": "der"
  },
  {
    "noun": "Wassereimer",
    "article": "der"
  },
  {
    "noun": "Wassereinbruch",
    "article": "der"
  },
  {
    "noun": "Wasserfall",
    "article": "der"
  },
  {
    "noun": "Wasserfleck",
    "article": "der"
  },
  {
    "noun": "Wasserfloh",
    "article": "der"
  },
  {
    "noun": "Wasserflughafen",
    "article": "der"
  },
  {
    "noun": "Wasserfrosch",
    "article": "der"
  },
  {
    "noun": "Wassergehalt",
    "article": "der"
  },
  {
    "noun": "Wassergraben",
    "article": "der"
  },
  {
    "noun": "Wasserhahn",
    "article": "der"
  },
  {
    "noun": "Wasserhaushalt",
    "article": "der"
  },
  {
    "noun": "Wasserkasten",
    "article": "der"
  },
  {
    "noun": "Wasserkelch",
    "article": "der"
  },
  {
    "noun": "Wasserkessel",
    "article": "der"
  },
  {
    "noun": "Wasserkopf",
    "article": "der"
  },
  {
    "noun": "Wasserkreislauf",
    "article": "der"
  },
  {
    "noun": "Wasserkrug",
    "article": "der"
  },
  {
    "noun": "Wasserlauf",
    "article": "der"
  },
  {
    "noun": "Wassermangel",
    "article": "der"
  },
  {
    "noun": "Wassermann",
    "article": "der"
  },
  {
    "noun": "Wassermesser",
    "article": "der"
  },
  {
    "noun": "Wasserpreis",
    "article": "der"
  },
  {
    "noun": "Wasserreichtum",
    "article": "der"
  },
  {
    "noun": "Wasserschaden",
    "article": "der"
  },
  {
    "noun": "Wasserschi",
    "article": "der"
  },
  {
    "noun": "Wasserschlag",
    "article": "der"
  },
  {
    "noun": "Wasserschlauch",
    "article": "der"
  },
  {
    "noun": "Wasserschwall",
    "article": "der"
  },
  {
    "noun": "Wasserski",
    "article": "der"
  },
  {
    "noun": "Wasserspeier",
    "article": "der"
  },
  {
    "noun": "Wasserspiegel",
    "article": "der"
  },
  {
    "noun": "Wassersport",
    "article": "der"
  },
  {
    "noun": "Wasserstand",
    "article": "der"
  },
  {
    "noun": "Wasserstandsanzeiger",
    "article": "der"
  },
  {
    "noun": "Wasserstein",
    "article": "der"
  },
  {
    "noun": "Wasserstoff",
    "article": "der"
  },
  {
    "noun": "Wasserstrahl",
    "article": "der"
  },
  {
    "noun": "Wassertank",
    "article": "der"
  },
  {
    "noun": "Wassertrog",
    "article": "der"
  },
  {
    "noun": "Wassertropfen",
    "article": "der"
  },
  {
    "noun": "Wasserturm",
    "article": "der"
  },
  {
    "noun": "Wasserverbrauch",
    "article": "der"
  },
  {
    "noun": "Wasserverlust",
    "article": "der"
  },
  {
    "noun": "Wasservogel",
    "article": "der"
  },
  {
    "noun": "Wasservorrat",
    "article": "der"
  },
  {
    "noun": "Wasserweg",
    "article": "der"
  },
  {
    "noun": "Wasserwerfer",
    "article": "der"
  },
  {
    "noun": "Wattebausch",
    "article": "der"
  },
  {
    "noun": "Wattstrom",
    "article": "der"
  },
  {
    "noun": "Wattwurm",
    "article": "der"
  },
  {
    "noun": "Watvogel",
    "article": "der"
  },
  {
    "noun": "Weber",
    "article": "der"
  },
  {
    "noun": "Weberkamm",
    "article": "der"
  },
  {
    "noun": "Weberknecht",
    "article": "der"
  },
  {
    "noun": "Webervogel",
    "article": "der"
  },
  {
    "noun": "Webmaster",
    "article": "der"
  },
  {
    "noun": "Webpelz",
    "article": "der"
  },
  {
    "noun": "Webstuhl",
    "article": "der"
  },
  {
    "noun": "Wechsel",
    "article": "der"
  },
  {
    "noun": "Wechselbrief",
    "article": "der"
  },
  {
    "noun": "Wechselgesang",
    "article": "der"
  },
  {
    "noun": "Wechselkredit",
    "article": "der"
  },
  {
    "noun": "Wechselkurs",
    "article": "der"
  },
  {
    "noun": "Wechselnehmer",
    "article": "der"
  },
  {
    "noun": "Wechselprotest",
    "article": "der"
  },
  {
    "noun": "Wechselprozess",
    "article": "der"
  },
  {
    "noun": "Wechselrahmen",
    "article": "der"
  },
  {
    "noun": "Wechselregress",
    "article": "der"
  },
  {
    "noun": "Wechselrichter",
    "article": "der"
  },
  {
    "noun": "Wechselschalter",
    "article": "der"
  },
  {
    "noun": "Wechselschritt",
    "article": "der"
  },
  {
    "noun": "Wechselschuldner",
    "article": "der"
  },
  {
    "noun": "Wechselstrom",
    "article": "der"
  },
  {
    "noun": "Wechselstromkreis",
    "article": "der"
  },
  {
    "noun": "Wechselstromwiderstand",
    "article": "der"
  },
  {
    "noun": "Wechselverkehr",
    "article": "der"
  },
  {
    "noun": "Wechsler",
    "article": "der"
  },
  {
    "noun": "Weckdienst",
    "article": "der"
  },
  {
    "noun": "Wecker",
    "article": "der"
  },
  {
    "noun": "Weckruf",
    "article": "der"
  },
  {
    "noun": "Wedel",
    "article": "der"
  },
  {
    "noun": "Weg",
    "article": "der"
  },
  {
    "noun": "Wegbereiter",
    "article": "der"
  },
  {
    "noun": "Wegebau",
    "article": "der"
  },
  {
    "noun": "Wegelagerer",
    "article": "der"
  },
  {
    "noun": "Wegemesser",
    "article": "der"
  },
  {
    "noun": "Wegerich",
    "article": "der"
  },
  {
    "noun": "Wegfall",
    "article": "der"
  },
  {
    "noun": "Weggang",
    "article": "der"
  },
  {
    "noun": "Wegmesser",
    "article": "der"
  },
  {
    "noun": "Wegrand",
    "article": "der"
  },
  {
    "noun": "Wegweiser",
    "article": "der"
  },
  {
    "noun": "Wegwerfartikel",
    "article": "der"
  },
  {
    "noun": "Wegzug",
    "article": "der"
  },
  {
    "noun": "Wehrbereich",
    "article": "der"
  },
  {
    "noun": "Wehrdienst",
    "article": "der"
  },
  {
    "noun": "Wehrdienstverweigerer",
    "article": "der"
  },
  {
    "noun": "Wehrersatzdienst",
    "article": "der"
  },
  {
    "noun": "Wehretat",
    "article": "der"
  },
  {
    "noun": "Wehrgang",
    "article": "der"
  },
  {
    "noun": "Wehrpass",
    "article": "der"
  },
  {
    "noun": "Wehrsold",
    "article": "der"
  },
  {
    "noun": "Wehrstand",
    "article": "der"
  },
  {
    "noun": "Wehrturm",
    "article": "der"
  },
  {
    "noun": "Weibel",
    "article": "der"
  },
  {
    "noun": "Weiberfeind",
    "article": "der"
  },
  {
    "noun": "Weiberheld",
    "article": "der"
  },
  {
    "noun": "Weiberknecht",
    "article": "der"
  },
  {
    "noun": "Weichensteller",
    "article": "der"
  },
  {
    "noun": "Weichling",
    "article": "der"
  },
  {
    "noun": "Weichmacher",
    "article": "der"
  },
  {
    "noun": "Weichselbaum",
    "article": "der"
  },
  {
    "noun": "Weidenbaum",
    "article": "der"
  },
  {
    "noun": "Weidenkorb",
    "article": "der"
  },
  {
    "noun": "Weideplatz",
    "article": "der"
  },
  {
    "noun": "Weiderich",
    "article": "der"
  },
  {
    "noun": "Weidling",
    "article": "der"
  },
  {
    "noun": "Weihbischof",
    "article": "der"
  },
  {
    "noun": "Weiher",
    "article": "der"
  },
  {
    "noun": "Weihkessel",
    "article": "der"
  },
  {
    "noun": "Weihnachtsabend",
    "article": "der"
  },
  {
    "noun": "Weihnachtsbaum",
    "article": "der"
  },
  {
    "noun": "Weihnachtsbaumschmuck",
    "article": "der"
  },
  {
    "noun": "Weihnachtseinkauf",
    "article": "der"
  },
  {
    "noun": "Weihnachtskarpfen",
    "article": "der"
  },
  {
    "noun": "Weihnachtsmann",
    "article": "der"
  },
  {
    "noun": "Weihnachtsmarkt",
    "article": "der"
  },
  {
    "noun": "Weihnachtsstern",
    "article": "der"
  },
  {
    "noun": "Weihnachtsstollen",
    "article": "der"
  },
  {
    "noun": "Weihnachtsteller",
    "article": "der"
  },
  {
    "noun": "Weihrauch",
    "article": "der"
  },
  {
    "noun": "Weiler",
    "article": "der"
  },
  {
    "noun": "Wein",
    "article": "der"
  },
  {
    "noun": "Weinanbau",
    "article": "der"
  },
  {
    "noun": "Weinbau",
    "article": "der"
  },
  {
    "noun": "Weinbauer",
    "article": "der"
  },
  {
    "noun": "Weinberg",
    "article": "der"
  },
  {
    "noun": "Weinbrand",
    "article": "der"
  },
  {
    "noun": "Weinbrandverschnitt",
    "article": "der"
  },
  {
    "noun": "Weinessig",
    "article": "der"
  },
  {
    "noun": "Weingarten",
    "article": "der"
  },
  {
    "noun": "Weingeist",
    "article": "der"
  },
  {
    "noun": "Weingott",
    "article": "der"
  },
  {
    "noun": "Weinheber",
    "article": "der"
  },
  {
    "noun": "Weinkeller",
    "article": "der"
  },
  {
    "noun": "Weinkenner",
    "article": "der"
  },
  {
    "noun": "Weinkrampf",
    "article": "der"
  },
  {
    "noun": "Weinliebhaber",
    "article": "der"
  },
  {
    "noun": "Weinschlauch",
    "article": "der"
  },
  {
    "noun": "Weinstein",
    "article": "der"
  },
  {
    "noun": "Weinstock",
    "article": "der"
  },
  {
    "noun": "Weintrinker",
    "article": "der"
  },
  {
    "noun": "Weisheitszahn",
    "article": "der"
  },
  {
    "noun": "Weissager",
    "article": "der"
  },
  {
    "noun": "Weitblick",
    "article": "der"
  },
  {
    "noun": "Weiterbestand",
    "article": "der"
  },
  {
    "noun": "Weiterflug",
    "article": "der"
  },
  {
    "noun": "Weitergang",
    "article": "der"
  },
  {
    "noun": "Weiterverkauf",
    "article": "der"
  },
  {
    "noun": "Weitschuss",
    "article": "der"
  },
  {
    "noun": "Weitsprung",
    "article": "der"
  },
  {
    "noun": "Weitstrahler",
    "article": "der"
  },
  {
    "noun": "Weizen",
    "article": "der"
  },
  {
    "noun": "Weizenkeim",
    "article": "der"
  },
  {
    "noun": "Wellbaum",
    "article": "der"
  },
  {
    "noun": "Wellenbereich",
    "article": "der"
  },
  {
    "noun": "Wellenberg",
    "article": "der"
  },
  {
    "noun": "Wellenbrecher",
    "article": "der"
  },
  {
    "noun": "Wellengang",
    "article": "der"
  },
  {
    "noun": "Wellenkamm",
    "article": "der"
  },
  {
    "noun": "Wellenleiter",
    "article": "der"
  },
  {
    "noun": "Wellenreiter",
    "article": "der"
  },
  {
    "noun": "Wellenschlag",
    "article": "der"
  },
  {
    "noun": "Wellensittich",
    "article": "der"
  },
  {
    "noun": "Wellenstrahl",
    "article": "der"
  },
  {
    "noun": "Wellenzug",
    "article": "der"
  },
  {
    "noun": "Welpe",
    "article": "der"
  },
  {
    "noun": "Wels",
    "article": "der"
  },
  {
    "noun": "Weltcup",
    "article": "der"
  },
  {
    "noun": "Weltenbummler",
    "article": "der"
  },
  {
    "noun": "Welterfolg",
    "article": "der"
  },
  {
    "noun": "Weltergewichtler",
    "article": "der"
  },
  {
    "noun": "Weltfriede",
    "article": "der"
  },
  {
    "noun": "Weltgewerkschaftsbund",
    "article": "der"
  },
  {
    "noun": "Welthandel",
    "article": "der"
  },
  {
    "noun": "Weltkrieg",
    "article": "der"
  },
  {
    "noun": "Weltlauf",
    "article": "der"
  },
  {
    "noun": "Weltmann",
    "article": "der"
  },
  {
    "noun": "Weltmarkt",
    "article": "der"
  },
  {
    "noun": "Weltmeister",
    "article": "der"
  },
  {
    "noun": "Weltpokal",
    "article": "der"
  },
  {
    "noun": "Weltraum",
    "article": "der"
  },
  {
    "noun": "Weltraumfahrer",
    "article": "der"
  },
  {
    "noun": "Weltrekord",
    "article": "der"
  },
  {
    "noun": "Weltrekordler",
    "article": "der"
  },
  {
    "noun": "Weltruf",
    "article": "der"
  },
  {
    "noun": "Weltruhm",
    "article": "der"
  },
  {
    "noun": "Weltschmerz",
    "article": "der"
  },
  {
    "noun": "Weltsicherheitsrat",
    "article": "der"
  },
  {
    "noun": "Weltspartag",
    "article": "der"
  },
  {
    "noun": "Weltteil",
    "article": "der"
  },
  {
    "noun": "Weltumsegler",
    "article": "der"
  },
  {
    "noun": "Weltverbesserer",
    "article": "der"
  },
  {
    "noun": "Weltwirtschaftsgipfel",
    "article": "der"
  },
  {
    "noun": "Wemfall",
    "article": "der"
  },
  {
    "noun": "Wendehals",
    "article": "der"
  },
  {
    "noun": "Wendekreis",
    "article": "der"
  },
  {
    "noun": "Wendeplatz",
    "article": "der"
  },
  {
    "noun": "Wendepunkt",
    "article": "der"
  },
  {
    "noun": "Wenfall",
    "article": "der"
  },
  {
    "noun": "Werbeanteil",
    "article": "der"
  },
  {
    "noun": "Werbeberater",
    "article": "der"
  },
  {
    "noun": "Werbebrief",
    "article": "der"
  },
  {
    "noun": "Werbeetat",
    "article": "der"
  },
  {
    "noun": "Werbefachmann",
    "article": "der"
  },
  {
    "noun": "Werbefeldzug",
    "article": "der"
  },
  {
    "noun": "Werbefilm",
    "article": "der"
  },
  {
    "noun": "Werbefunk",
    "article": "der"
  },
  {
    "noun": "Werbegag",
    "article": "der"
  },
  {
    "noun": "Werbegrafiker",
    "article": "der"
  },
  {
    "noun": "Werbeleiter",
    "article": "der"
  },
  {
    "noun": "Werbepreis",
    "article": "der"
  },
  {
    "noun": "Werbeprospekt",
    "article": "der"
  },
  {
    "noun": "Werber",
    "article": "der"
  },
  {
    "noun": "Werbeslogan",
    "article": "der"
  },
  {
    "noun": "Werbespot",
    "article": "der"
  },
  {
    "noun": "Werbespruch",
    "article": "der"
  },
  {
    "noun": "Werbetext",
    "article": "der"
  },
  {
    "noun": "Werbetexter",
    "article": "der"
  },
  {
    "noun": "Werbevertreter",
    "article": "der"
  },
  {
    "noun": "Werdegang",
    "article": "der"
  },
  {
    "noun": "Werfall",
    "article": "der"
  },
  {
    "noun": "Werfer",
    "article": "der"
  },
  {
    "noun": "Werkarzt",
    "article": "der"
  },
  {
    "noun": "Werklehrer",
    "article": "der"
  },
  {
    "noun": "Werkleiter",
    "article": "der"
  },
  {
    "noun": "Werklieferungsvertrag",
    "article": "der"
  },
  {
    "noun": "Werkmeister",
    "article": "der"
  },
  {
    "noun": "Werkraum",
    "article": "der"
  },
  {
    "noun": "Werksarzt",
    "article": "der"
  },
  {
    "noun": "Werkschutz",
    "article": "der"
  },
  {
    "noun": "Werksleiter",
    "article": "der"
  },
  {
    "noun": "Werkstoff",
    "article": "der"
  },
  {
    "noun": "Werkstudent",
    "article": "der"
  },
  {
    "noun": "Werktag",
    "article": "der"
  },
  {
    "noun": "Werktisch",
    "article": "der"
  },
  {
    "noun": "Werkvertrag",
    "article": "der"
  },
  {
    "noun": "Werkzeugkasten",
    "article": "der"
  },
  {
    "noun": "Werkzeugmacher",
    "article": "der"
  },
  {
    "noun": "Werkzeugschrank",
    "article": "der"
  },
  {
    "noun": "Werkzeugstahl",
    "article": "der"
  },
  {
    "noun": "Wermut",
    "article": "der"
  },
  {
    "noun": "Wermutbruder",
    "article": "der"
  },
  {
    "noun": "Wermutstropfen",
    "article": "der"
  },
  {
    "noun": "Wermutwein",
    "article": "der"
  },
  {
    "noun": "Wertbrief",
    "article": "der"
  },
  {
    "noun": "Wertewandel",
    "article": "der"
  },
  {
    "noun": "Wertgegenstand",
    "article": "der"
  },
  {
    "noun": "Wertmesser",
    "article": "der"
  },
  {
    "noun": "Wertstoff",
    "article": "der"
  },
  {
    "noun": "Wertungslauf",
    "article": "der"
  },
  {
    "noun": "Wertungspunkt",
    "article": "der"
  },
  {
    "noun": "Wertverlust",
    "article": "der"
  },
  {
    "noun": "Wertzoll",
    "article": "der"
  },
  {
    "noun": "Wertzuwachs",
    "article": "der"
  },
  {
    "noun": "Werwolf",
    "article": "der"
  },
  {
    "noun": "Wesenszug",
    "article": "der"
  },
  {
    "noun": "Wesfall",
    "article": "der"
  },
  {
    "noun": "Wesir",
    "article": "der"
  },
  {
    "noun": "Wespenstich",
    "article": "der"
  },
  {
    "noun": "Westen",
    "article": "der"
  },
  {
    "noun": "Western",
    "article": "der"
  },
  {
    "noun": "Westfale",
    "article": "der"
  },
  {
    "noun": "Westover",
    "article": "der"
  },
  {
    "noun": "Westwind",
    "article": "der"
  },
  {
    "noun": "Wettbewerb",
    "article": "der"
  },
  {
    "noun": "Wettbewerber",
    "article": "der"
  },
  {
    "noun": "Wetteifer",
    "article": "der"
  },
  {
    "noun": "Wetteiferer",
    "article": "der"
  },
  {
    "noun": "Wetterbericht",
    "article": "der"
  },
  {
    "noun": "Wetterdienst",
    "article": "der"
  },
  {
    "noun": "Wetterfrosch",
    "article": "der"
  },
  {
    "noun": "Wetterhahn",
    "article": "der"
  },
  {
    "noun": "Wetterprophet",
    "article": "der"
  },
  {
    "noun": "Wettersatellit",
    "article": "der"
  },
  {
    "noun": "Wetterschacht",
    "article": "der"
  },
  {
    "noun": "Wetterschaden",
    "article": "der"
  },
  {
    "noun": "Wetterschutz",
    "article": "der"
  },
  {
    "noun": "Wettersturz",
    "article": "der"
  },
  {
    "noun": "Wetterumschwung",
    "article": "der"
  },
  {
    "noun": "Wetterwechsel",
    "article": "der"
  },
  {
    "noun": "Wettkampf",
    "article": "der"
  },
  {
    "noun": "Wettlauf",
    "article": "der"
  },
  {
    "noun": "Wettstreit",
    "article": "der"
  },
  {
    "noun": "Wetzstahl",
    "article": "der"
  },
  {
    "noun": "Wetzstein",
    "article": "der"
  },
  {
    "noun": "Whirlpool",
    "article": "der"
  },
  {
    "noun": "Whiskey",
    "article": "der"
  },
  {
    "noun": "Whisky",
    "article": "der"
  },
  {
    "noun": "Wichser",
    "article": "der"
  },
  {
    "noun": "Wicht",
    "article": "der"
  },
  {
    "noun": "Wichtel",
    "article": "der"
  },
  {
    "noun": "Wichtigtuer",
    "article": "der"
  },
  {
    "noun": "Wickel",
    "article": "der"
  },
  {
    "noun": "Wickelraum",
    "article": "der"
  },
  {
    "noun": "Wickelrock",
    "article": "der"
  },
  {
    "noun": "Wickeltisch",
    "article": "der"
  },
  {
    "noun": "Wickler",
    "article": "der"
  },
  {
    "noun": "Widder",
    "article": "der"
  },
  {
    "noun": "Widerdruck",
    "article": "der"
  },
  {
    "noun": "Widerhaken",
    "article": "der"
  },
  {
    "noun": "Widerhall",
    "article": "der"
  },
  {
    "noun": "Widerhalt",
    "article": "der"
  },
  {
    "noun": "Widerling",
    "article": "der"
  },
  {
    "noun": "Widerpart",
    "article": "der"
  },
  {
    "noun": "Widerrist",
    "article": "der"
  },
  {
    "noun": "Widerruf",
    "article": "der"
  },
  {
    "noun": "Widersacher",
    "article": "der"
  },
  {
    "noun": "Widerschein",
    "article": "der"
  },
  {
    "noun": "Widersinn",
    "article": "der"
  },
  {
    "noun": "Widerspruch",
    "article": "der"
  },
  {
    "noun": "Widerspruchsgeist",
    "article": "der"
  },
  {
    "noun": "Widerstand",
    "article": "der"
  },
  {
    "noun": "Widerstandsmesser",
    "article": "der"
  },
  {
    "noun": "Widerstrahl",
    "article": "der"
  },
  {
    "noun": "Widerstreit",
    "article": "der"
  },
  {
    "noun": "Widerwille",
    "article": "der"
  },
  {
    "noun": "Wiedehopf",
    "article": "der"
  },
  {
    "noun": "Wiederabdruck",
    "article": "der"
  },
  {
    "noun": "Wiederaufbau",
    "article": "der"
  },
  {
    "noun": "Wiederbeginn",
    "article": "der"
  },
  {
    "noun": "Wiederbelebungsversuch",
    "article": "der"
  },
  {
    "noun": "Wiederbeschaffungswert",
    "article": "der"
  },
  {
    "noun": "Wiederdruck",
    "article": "der"
  },
  {
    "noun": "Wiedereinstieg",
    "article": "der"
  },
  {
    "noun": "Wiedereintritt",
    "article": "der"
  },
  {
    "noun": "Wiederholungsfall",
    "article": "der"
  },
  {
    "noun": "Wiederkauf",
    "article": "der"
  },
  {
    "noun": "Wiederverkauf",
    "article": "der"
  },
  {
    "noun": "Wiederverkaufswert",
    "article": "der"
  },
  {
    "noun": "Wiegendruck",
    "article": "der"
  },
  {
    "noun": "Wiesbaum",
    "article": "der"
  },
  {
    "noun": "Wiesenchampignon",
    "article": "der"
  },
  {
    "noun": "Wiesengrund",
    "article": "der"
  },
  {
    "noun": "Wigwam",
    "article": "der"
  },
  {
    "noun": "Wikinger",
    "article": "der"
  },
  {
    "noun": "Wildbach",
    "article": "der"
  },
  {
    "noun": "Wildbraten",
    "article": "der"
  },
  {
    "noun": "Wilddieb",
    "article": "der"
  },
  {
    "noun": "Wilderer",
    "article": "der"
  },
  {
    "noun": "Wildesel",
    "article": "der"
  },
  {
    "noun": "Wildfang",
    "article": "der"
  },
  {
    "noun": "Wildgeschmack",
    "article": "der"
  },
  {
    "noun": "Wildheger",
    "article": "der"
  },
  {
    "noun": "Wildhund",
    "article": "der"
  },
  {
    "noun": "Wildling",
    "article": "der"
  },
  {
    "noun": "Wildpark",
    "article": "der"
  },
  {
    "noun": "Wildreis",
    "article": "der"
  },
  {
    "noun": "Wildstand",
    "article": "der"
  },
  {
    "noun": "Wildtyp",
    "article": "der"
  },
  {
    "noun": "Wildverbiss",
    "article": "der"
  },
  {
    "noun": "Wildwechsel",
    "article": "der"
  },
  {
    "noun": "Wildwestfilm",
    "article": "der"
  },
  {
    "noun": "Wildwuchs",
    "article": "der"
  },
  {
    "noun": "Wille",
    "article": "der"
  },
  {
    "noun": "Willensakt",
    "article": "der"
  },
  {
    "noun": "Willkomm",
    "article": "der"
  },
  {
    "noun": "Willkommenstrunk",
    "article": "der"
  },
  {
    "noun": "Wimpel",
    "article": "der"
  },
  {
    "noun": "Wind",
    "article": "der"
  },
  {
    "noun": "Windabweiser",
    "article": "der"
  },
  {
    "noun": "Windbeutel",
    "article": "der"
  },
  {
    "noun": "Winderhitzer",
    "article": "der"
  },
  {
    "noun": "Windfang",
    "article": "der"
  },
  {
    "noun": "Windhafer",
    "article": "der"
  },
  {
    "noun": "Windhund",
    "article": "der"
  },
  {
    "noun": "Windjammer",
    "article": "der"
  },
  {
    "noun": "Windkanal",
    "article": "der"
  },
  {
    "noun": "Windmesser",
    "article": "der"
  },
  {
    "noun": "Windpark",
    "article": "der"
  },
  {
    "noun": "Windsack",
    "article": "der"
  },
  {
    "noun": "Windschatten",
    "article": "der"
  },
  {
    "noun": "Windschirm",
    "article": "der"
  },
  {
    "noun": "Windschutz",
    "article": "der"
  },
  {
    "noun": "Windzug",
    "article": "der"
  },
  {
    "noun": "Wink",
    "article": "der"
  },
  {
    "noun": "Winkel",
    "article": "der"
  },
  {
    "noun": "Winkeladvokat",
    "article": "der"
  },
  {
    "noun": "Winkelgrad",
    "article": "der"
  },
  {
    "noun": "Winkelhaken",
    "article": "der"
  },
  {
    "noun": "Winkelmesser",
    "article": "der"
  },
  {
    "noun": "Winkelstahl",
    "article": "der"
  },
  {
    "noun": "Winkelzug",
    "article": "der"
  },
  {
    "noun": "Winker",
    "article": "der"
  },
  {
    "noun": "Winter",
    "article": "der"
  },
  {
    "noun": "Winterdienst",
    "article": "der"
  },
  {
    "noun": "Wintereinbruch",
    "article": "der"
  },
  {
    "noun": "Winterfahrplan",
    "article": "der"
  },
  {
    "noun": "Wintergarten",
    "article": "der"
  },
  {
    "noun": "Winterhafen",
    "article": "der"
  },
  {
    "noun": "Winterkurort",
    "article": "der"
  },
  {
    "noun": "Winterling",
    "article": "der"
  },
  {
    "noun": "Wintermantel",
    "article": "der"
  },
  {
    "noun": "Wintermonat",
    "article": "der"
  },
  {
    "noun": "Winterquartier",
    "article": "der"
  },
  {
    "noun": "Winterreifen",
    "article": "der"
  },
  {
    "noun": "Wintersanfang",
    "article": "der"
  },
  {
    "noun": "Winterschlaf",
    "article": "der"
  },
  {
    "noun": "Winterschlussverkauf",
    "article": "der"
  },
  {
    "noun": "Wintersport",
    "article": "der"
  },
  {
    "noun": "Wintervorrat",
    "article": "der"
  },
  {
    "noun": "Winzer",
    "article": "der"
  },
  {
    "noun": "Winzling",
    "article": "der"
  },
  {
    "noun": "Wipfel",
    "article": "der"
  },
  {
    "noun": "Wirbel",
    "article": "der"
  },
  {
    "noun": "Wirbelknochen",
    "article": "der"
  },
  {
    "noun": "Wirbelstrom",
    "article": "der"
  },
  {
    "noun": "Wirbelsturm",
    "article": "der"
  },
  {
    "noun": "Wirbelwind",
    "article": "der"
  },
  {
    "noun": "Wirklichkeitsmensch",
    "article": "der"
  },
  {
    "noun": "Wirklichkeitssinn",
    "article": "der"
  },
  {
    "noun": "Wirkstoff",
    "article": "der"
  },
  {
    "noun": "Wirkungsbereich",
    "article": "der"
  },
  {
    "noun": "Wirkungsgrad",
    "article": "der"
  },
  {
    "noun": "Wirkungskreis",
    "article": "der"
  },
  {
    "noun": "Wirkungsquerschnitt",
    "article": "der"
  },
  {
    "noun": "Wirrkopf",
    "article": "der"
  },
  {
    "noun": "Wirrwarr",
    "article": "der"
  },
  {
    "noun": "Wirsing",
    "article": "der"
  },
  {
    "noun": "Wirsingkohl",
    "article": "der"
  },
  {
    "noun": "Wirt",
    "article": "der"
  },
  {
    "noun": "Wirtel",
    "article": "der"
  },
  {
    "noun": "Wirtschaftler",
    "article": "der"
  },
  {
    "noun": "Wirtschaftsaufschwung",
    "article": "der"
  },
  {
    "noun": "Wirtschaftsausschuss",
    "article": "der"
  },
  {
    "noun": "Wirtschaftsberater",
    "article": "der"
  },
  {
    "noun": "Wirtschaftshof",
    "article": "der"
  },
  {
    "noun": "Wirtschaftsingenieur",
    "article": "der"
  },
  {
    "noun": "Wirtschaftskreislauf",
    "article": "der"
  },
  {
    "noun": "Wirtschaftskrieg",
    "article": "der"
  },
  {
    "noun": "Wirtschaftsliberalismus",
    "article": "der"
  },
  {
    "noun": "Wirtschaftsminister",
    "article": "der"
  },
  {
    "noun": "Wirtschaftsplan",
    "article": "der"
  },
  {
    "noun": "Wirtschaftspolitiker",
    "article": "der"
  },
  {
    "noun": "Wirtschaftsrat",
    "article": "der"
  },
  {
    "noun": "Wirtschaftsraum",
    "article": "der"
  },
  {
    "noun": "Wirtschaftssektor",
    "article": "der"
  },
  {
    "noun": "Wirtschaftsstandort",
    "article": "der"
  },
  {
    "noun": "Wirtschaftsteil",
    "article": "der"
  },
  {
    "noun": "Wirtschaftsverband",
    "article": "der"
  },
  {
    "noun": "Wirtschaftswissenschaftler",
    "article": "der"
  },
  {
    "noun": "Wirtschaftszweig",
    "article": "der"
  },
  {
    "noun": "Wisch",
    "article": "der"
  },
  {
    "noun": "Wischer",
    "article": "der"
  },
  {
    "noun": "Wischlappen",
    "article": "der"
  },
  {
    "noun": "Wisent",
    "article": "der"
  },
  {
    "noun": "Wissenschafter",
    "article": "der"
  },
  {
    "noun": "Wissenschaftler",
    "article": "der"
  },
  {
    "noun": "Wissensdrang",
    "article": "der"
  },
  {
    "noun": "Wissensdurst",
    "article": "der"
  },
  {
    "noun": "Wissensschatz",
    "article": "der"
  },
  {
    "noun": "Wissensstand",
    "article": "der"
  },
  {
    "noun": "Wissenstransfer",
    "article": "der"
  },
  {
    "noun": "Wissenszweig",
    "article": "der"
  },
  {
    "noun": "Witterungseinfluss",
    "article": "der"
  },
  {
    "noun": "Witterungsumschlag",
    "article": "der"
  },
  {
    "noun": "Witwenstand",
    "article": "der"
  },
  {
    "noun": "Witwer",
    "article": "der"
  },
  {
    "noun": "Witz",
    "article": "der"
  },
  {
    "noun": "Witzbold",
    "article": "der"
  },
  {
    "noun": "Witzling",
    "article": "der"
  },
  {
    "noun": "Wobbelsender",
    "article": "der"
  },
  {
    "noun": "Wochenendkurs",
    "article": "der"
  },
  {
    "noun": "Wochenfluss",
    "article": "der"
  },
  {
    "noun": "Wochenlohn",
    "article": "der"
  },
  {
    "noun": "Wochenmarkt",
    "article": "der"
  },
  {
    "noun": "Wochenspielplan",
    "article": "der"
  },
  {
    "noun": "Wochentag",
    "article": "der"
  },
  {
    "noun": "Wodka",
    "article": "der"
  },
  {
    "noun": "Wodu",
    "article": "der"
  },
  {
    "noun": "Wohlfahrtsstaat",
    "article": "der"
  },
  {
    "noun": "Wohlgeruch",
    "article": "der"
  },
  {
    "noun": "Wohlgeschmack",
    "article": "der"
  },
  {
    "noun": "Wohlklang",
    "article": "der"
  },
  {
    "noun": "Wohllaut",
    "article": "der"
  },
  {
    "noun": "Wohlstand",
    "article": "der"
  },
  {
    "noun": "Wohnbau",
    "article": "der"
  },
  {
    "noun": "Wohnbereich",
    "article": "der"
  },
  {
    "noun": "Wohnbezirk",
    "article": "der"
  },
  {
    "noun": "Wohnblock",
    "article": "der"
  },
  {
    "noun": "Wohnkomfort",
    "article": "der"
  },
  {
    "noun": "Wohnort",
    "article": "der"
  },
  {
    "noun": "Wohnraum",
    "article": "der"
  },
  {
    "noun": "Wohnsitz",
    "article": "der"
  },
  {
    "noun": "Wohnungsbau",
    "article": "der"
  },
  {
    "noun": "Wohnungsbrand",
    "article": "der"
  },
  {
    "noun": "Wohnungsinhaber",
    "article": "der"
  },
  {
    "noun": "Wohnungsmangel",
    "article": "der"
  },
  {
    "noun": "Wohnungsmarkt",
    "article": "der"
  },
  {
    "noun": "Wohnungstausch",
    "article": "der"
  },
  {
    "noun": "Wohnungswechsel",
    "article": "der"
  },
  {
    "noun": "Wohnwagen",
    "article": "der"
  },
  {
    "noun": "Wok",
    "article": "der"
  },
  {
    "noun": "Wolf",
    "article": "der"
  },
  {
    "noun": "Wolframfaden",
    "article": "der"
  },
  {
    "noun": "Wolfshund",
    "article": "der"
  },
  {
    "noun": "Wolfshunger",
    "article": "der"
  },
  {
    "noun": "Wolfsmensch",
    "article": "der"
  },
  {
    "noun": "Wolfspelz",
    "article": "der"
  },
  {
    "noun": "Wolfsrachen",
    "article": "der"
  },
  {
    "noun": "Wolkenbruch",
    "article": "der"
  },
  {
    "noun": "Wolkenfetzen",
    "article": "der"
  },
  {
    "noun": "Wolkenhimmel",
    "article": "der"
  },
  {
    "noun": "Wolkenkratzer",
    "article": "der"
  },
  {
    "noun": "Wolkenschieber",
    "article": "der"
  },
  {
    "noun": "Wollhandel",
    "article": "der"
  },
  {
    "noun": "Wollkamm",
    "article": "der"
  },
  {
    "noun": "Wollmantel",
    "article": "der"
  },
  {
    "noun": "Wollschal",
    "article": "der"
  },
  {
    "noun": "Wollstoff",
    "article": "der"
  },
  {
    "noun": "Wollstrumpf",
    "article": "der"
  },
  {
    "noun": "Wolpertinger",
    "article": "der"
  },
  {
    "noun": "Wombat",
    "article": "der"
  },
  {
    "noun": "Wonnemonat",
    "article": "der"
  },
  {
    "noun": "Wonneproppen",
    "article": "der"
  },
  {
    "noun": "Worb",
    "article": "der"
  },
  {
    "noun": "Workshop",
    "article": "der"
  },
  {
    "noun": "Wortakzent",
    "article": "der"
  },
  {
    "noun": "Wortbruch",
    "article": "der"
  },
  {
    "noun": "Wortklauber",
    "article": "der"
  },
  {
    "noun": "Wortlaut",
    "article": "der"
  },
  {
    "noun": "Wortreichtum",
    "article": "der"
  },
  {
    "noun": "Wortschatz",
    "article": "der"
  },
  {
    "noun": "Wortschwall",
    "article": "der"
  },
  {
    "noun": "Wortsinn",
    "article": "der"
  },
  {
    "noun": "Wortstamm",
    "article": "der"
  },
  {
    "noun": "Wortstreit",
    "article": "der"
  },
  {
    "noun": "Wortverdreher",
    "article": "der"
  },
  {
    "noun": "Wortwechsel",
    "article": "der"
  },
  {
    "noun": "Wrasen",
    "article": "der"
  },
  {
    "noun": "Wucher",
    "article": "der"
  },
  {
    "noun": "Wucherer",
    "article": "der"
  },
  {
    "noun": "Wucherpreis",
    "article": "der"
  },
  {
    "noun": "Wucherzins",
    "article": "der"
  },
  {
    "noun": "Wuchs",
    "article": "der"
  },
  {
    "noun": "Wuchsstoff",
    "article": "der"
  },
  {
    "noun": "Wuchtbaum",
    "article": "der"
  },
  {
    "noun": "Wundarzt",
    "article": "der"
  },
  {
    "noun": "Wundbrand",
    "article": "der"
  },
  {
    "noun": "Wunderdoktor",
    "article": "der"
  },
  {
    "noun": "Wunderglaube",
    "article": "der"
  },
  {
    "noun": "Wunderheiler",
    "article": "der"
  },
  {
    "noun": "Wunderknabe",
    "article": "der"
  },
  {
    "noun": "Wundstarrkrampf",
    "article": "der"
  },
  {
    "noun": "Wundverband",
    "article": "der"
  },
  {
    "noun": "Wunsch",
    "article": "der"
  },
  {
    "noun": "Wunschsatz",
    "article": "der"
  },
  {
    "noun": "Wunschtraum",
    "article": "der"
  },
  {
    "noun": "Wunschzettel",
    "article": "der"
  },
  {
    "noun": "Wurf",
    "article": "der"
  },
  {
    "noun": "Wurfkreis",
    "article": "der"
  },
  {
    "noun": "Wurfpfeil",
    "article": "der"
  },
  {
    "noun": "Wurfring",
    "article": "der"
  },
  {
    "noun": "Wurmbefall",
    "article": "der"
  },
  {
    "noun": "Wurmfortsatz",
    "article": "der"
  },
  {
    "noun": "Wurstsalat",
    "article": "der"
  },
  {
    "noun": "Wurstzipfel",
    "article": "der"
  },
  {
    "noun": "Wurzelballen",
    "article": "der"
  },
  {
    "noun": "Wurzelexponent",
    "article": "der"
  },
  {
    "noun": "Wurzelstock",
    "article": "der"
  },
  {
    "noun": "Wuschelkopf",
    "article": "der"
  },
  {
    "noun": "Wust",
    "article": "der"
  },
  {
    "noun": "Wutanfall",
    "article": "der"
  },
  {
    "noun": "Wutausbruch",
    "article": "der"
  },
  {
    "noun": "Radius",
    "article": "der"
  },
  {
    "noun": "Radiusvektor",
    "article": "der"
  },
  {
    "noun": "Radkasten",
    "article": "der"
  },
  {
    "noun": "Radkranz",
    "article": "der"
  },
  {
    "noun": "Radler",
    "article": "der"
  },
  {
    "noun": "Radschuh",
    "article": "der"
  },
  {
    "noun": "Radsport",
    "article": "der"
  },
  {
    "noun": "Radsportler",
    "article": "der"
  },
  {
    "noun": "Radstand",
    "article": "der"
  },
  {
    "noun": "Radsturz",
    "article": "der"
  },
  {
    "noun": "Radwechsel",
    "article": "der"
  },
  {
    "noun": "Radweg",
    "article": "der"
  },
  {
    "noun": "Raffke",
    "article": "der"
  },
  {
    "noun": "Raglan",
    "article": "der"
  },
  {
    "noun": "Raglanmantel",
    "article": "der"
  },
  {
    "noun": "Ragtime",
    "article": "der"
  },
  {
    "noun": "Rahm",
    "article": "der"
  },
  {
    "noun": "Rahmapfel",
    "article": "der"
  },
  {
    "noun": "Rahmenbeschluss",
    "article": "der"
  },
  {
    "noun": "Rahmenplan",
    "article": "der"
  },
  {
    "noun": "Rahmentarif",
    "article": "der"
  },
  {
    "noun": "Rahmentarifvertrag",
    "article": "der"
  },
  {
    "noun": "Rain",
    "article": "der"
  },
  {
    "noun": "Rainfarn",
    "article": "der"
  },
  {
    "noun": "Rajah",
    "article": "der"
  },
  {
    "noun": "Raketenantrieb",
    "article": "der"
  },
  {
    "noun": "Raketenmotor",
    "article": "der"
  },
  {
    "noun": "Raketenstart",
    "article": "der"
  },
  {
    "noun": "Raketentreibstoff",
    "article": "der"
  },
  {
    "noun": "Raketenwerfer",
    "article": "der"
  },
  {
    "noun": "Ramadan",
    "article": "der"
  },
  {
    "noun": "Ramaneffekt",
    "article": "der"
  },
  {
    "noun": "Rammbock",
    "article": "der"
  },
  {
    "noun": "Rammklotz",
    "article": "der"
  },
  {
    "noun": "Rammler",
    "article": "der"
  },
  {
    "noun": "Ramsch",
    "article": "der"
  },
  {
    "noun": "Ramschladen",
    "article": "der"
  },
  {
    "noun": "Rancher",
    "article": "der"
  },
  {
    "noun": "Rand",
    "article": "der"
  },
  {
    "noun": "Randal",
    "article": "der"
  },
  {
    "noun": "Randalierer",
    "article": "der"
  },
  {
    "noun": "Randausgleich",
    "article": "der"
  },
  {
    "noun": "Randbezirk",
    "article": "der"
  },
  {
    "noun": "Randstein",
    "article": "der"
  },
  {
    "noun": "Randsteller",
    "article": "der"
  },
  {
    "noun": "Randstreifen",
    "article": "der"
  },
  {
    "noun": "Ranft",
    "article": "der"
  },
  {
    "noun": "Rang",
    "article": "der"
  },
  {
    "noun": "Rangierbahnhof",
    "article": "der"
  },
  {
    "noun": "Rangierer",
    "article": "der"
  },
  {
    "noun": "Rangiermeister",
    "article": "der"
  },
  {
    "noun": "Rank",
    "article": "der"
  },
  {
    "noun": "Rap",
    "article": "der"
  },
  {
    "noun": "Rappel",
    "article": "der"
  },
  {
    "noun": "Rappelkopf",
    "article": "der"
  },
  {
    "noun": "Rappenspalter",
    "article": "der"
  },
  {
    "noun": "Rapport",
    "article": "der"
  },
  {
    "noun": "Raps",
    "article": "der"
  },
  {
    "noun": "Rapsacker",
    "article": "der"
  },
  {
    "noun": "Rapssame",
    "article": "der"
  },
  {
    "noun": "Rasenplatz",
    "article": "der"
  },
  {
    "noun": "Rasensprenger",
    "article": "der"
  },
  {
    "noun": "Rasenstreifen",
    "article": "der"
  },
  {
    "noun": "Raser",
    "article": "der"
  },
  {
    "noun": "Rasierapparat",
    "article": "der"
  },
  {
    "noun": "Rasierer",
    "article": "der"
  },
  {
    "noun": "Rasierpinsel",
    "article": "der"
  },
  {
    "noun": "Rasierschaum",
    "article": "der"
  },
  {
    "noun": "Rasierspiegel",
    "article": "der"
  },
  {
    "noun": "Rassehund",
    "article": "der"
  },
  {
    "noun": "Rassenfanatiker",
    "article": "der"
  },
  {
    "noun": "Rassenhass",
    "article": "der"
  },
  {
    "noun": "Rassenkampf",
    "article": "der"
  },
  {
    "noun": "Rassenkonflikt",
    "article": "der"
  },
  {
    "noun": "Rassenkrawall",
    "article": "der"
  },
  {
    "noun": "Rassismus",
    "article": "der"
  },
  {
    "noun": "Rassist",
    "article": "der"
  },
  {
    "noun": "Rasterpunkt",
    "article": "der"
  },
  {
    "noun": "Rasthof",
    "article": "der"
  },
  {
    "noun": "Rastplatz",
    "article": "der"
  },
  {
    "noun": "Rat",
    "article": "der"
  },
  {
    "noun": "Ratenbetrag",
    "article": "der"
  },
  {
    "noun": "Ratenkauf",
    "article": "der"
  },
  {
    "noun": "Ratenkredit",
    "article": "der"
  },
  {
    "noun": "Ratenzahlungskredit",
    "article": "der"
  },
  {
    "noun": "Rater",
    "article": "der"
  },
  {
    "noun": "Ratgeber",
    "article": "der"
  },
  {
    "noun": "Ratiodetektor",
    "article": "der"
  },
  {
    "noun": "Rationalismus",
    "article": "der"
  },
  {
    "noun": "Rationalist",
    "article": "der"
  },
  {
    "noun": "Ratsbeschluss",
    "article": "der"
  },
  {
    "noun": "Ratschlag",
    "article": "der"
  },
  {
    "noun": "Ratschluss",
    "article": "der"
  },
  {
    "noun": "Ratsherr",
    "article": "der"
  },
  {
    "noun": "Ratskeller",
    "article": "der"
  },
  {
    "noun": "Rattler",
    "article": "der"
  },
  {
    "noun": "Ratz",
    "article": "der"
  },
  {
    "noun": "Raub",
    "article": "der"
  },
  {
    "noun": "Raubauz",
    "article": "der"
  },
  {
    "noun": "Raubbau",
    "article": "der"
  },
  {
    "noun": "Raubdruck",
    "article": "der"
  },
  {
    "noun": "Raubfisch",
    "article": "der"
  },
  {
    "noun": "Raubritter",
    "article": "der"
  },
  {
    "noun": "Raubvogel",
    "article": "der"
  },
  {
    "noun": "Raubzug",
    "article": "der"
  },
  {
    "noun": "Rauch",
    "article": "der"
  },
  {
    "noun": "Rauchabzug",
    "article": "der"
  },
  {
    "noun": "Raucher",
    "article": "der"
  },
  {
    "noun": "Raucherhusten",
    "article": "der"
  },
  {
    "noun": "Rauchfang",
    "article": "der"
  },
  {
    "noun": "Rauchfangkehrer",
    "article": "der"
  },
  {
    "noun": "Rauchmelder",
    "article": "der"
  },
  {
    "noun": "Rauchpilz",
    "article": "der"
  },
  {
    "noun": "Rauchschwaden",
    "article": "der"
  },
  {
    "noun": "Rauchverzehrer",
    "article": "der"
  },
  {
    "noun": "Rauchvorhang",
    "article": "der"
  },
  {
    "noun": "Raufbold",
    "article": "der"
  },
  {
    "noun": "Raufer",
    "article": "der"
  },
  {
    "noun": "Raufhandel",
    "article": "der"
  },
  {
    "noun": "Raufrost",
    "article": "der"
  },
  {
    "noun": "Rauhputz",
    "article": "der"
  },
  {
    "noun": "Rauhreif",
    "article": "der"
  },
  {
    "noun": "Raumanzug",
    "article": "der"
  },
  {
    "noun": "Raumausstatter",
    "article": "der"
  },
  {
    "noun": "Raumfahrer",
    "article": "der"
  },
  {
    "noun": "Raumflug",
    "article": "der"
  },
  {
    "noun": "Raumgewinn",
    "article": "der"
  },
  {
    "noun": "Rauminhalt",
    "article": "der"
  },
  {
    "noun": "Raumklang",
    "article": "der"
  },
  {
    "noun": "Raummangel",
    "article": "der"
  },
  {
    "noun": "Raumpfleger",
    "article": "der"
  },
  {
    "noun": "Raumteiler",
    "article": "der"
  },
  {
    "noun": "Raumton",
    "article": "der"
  },
  {
    "noun": "Raumwinkel",
    "article": "der"
  },
  {
    "noun": "Raunzer",
    "article": "der"
  },
  {
    "noun": "Raupenbagger",
    "article": "der"
  },
  {
    "noun": "Raupenschlepper",
    "article": "der"
  },
  {
    "noun": "Rausch",
    "article": "der"
  },
  {
    "noun": "Rauschgifthandel",
    "article": "der"
  },
  {
    "noun": "Rauschgoldengel",
    "article": "der"
  },
  {
    "noun": "Rausschmiss",
    "article": "der"
  },
  {
    "noun": "Raviolo",
    "article": "der"
  },
  {
    "noun": "Reaktant",
    "article": "der"
  },
  {
    "noun": "Reaktionsmotor",
    "article": "der"
  },
  {
    "noun": "Reaktor",
    "article": "der"
  },
  {
    "noun": "Reaktorkern",
    "article": "der"
  },
  {
    "noun": "Realgar",
    "article": "der"
  },
  {
    "noun": "Realisationswert",
    "article": "der"
  },
  {
    "noun": "Realisator",
    "article": "der"
  },
  {
    "noun": "Realismus",
    "article": "der"
  },
  {
    "noun": "Realist",
    "article": "der"
  },
  {
    "noun": "Realkontrakt",
    "article": "der"
  },
  {
    "noun": "Realkredit",
    "article": "der"
  },
  {
    "noun": "Reallohn",
    "article": "der"
  },
  {
    "noun": "Realpolitiker",
    "article": "der"
  },
  {
    "noun": "Realwert",
    "article": "der"
  },
  {
    "noun": "Realzins",
    "article": "der"
  },
  {
    "noun": "Rebbach",
    "article": "der"
  },
  {
    "noun": "Rebell",
    "article": "der"
  },
  {
    "noun": "Rebensaft",
    "article": "der"
  },
  {
    "noun": "Rebound",
    "article": "der"
  },
  {
    "noun": "Rebstock",
    "article": "der"
  },
  {
    "noun": "Rechenautomat",
    "article": "der"
  },
  {
    "noun": "Rechenfehler",
    "article": "der"
  },
  {
    "noun": "Rechenschaftsbericht",
    "article": "der"
  },
  {
    "noun": "Rechenschieber",
    "article": "der"
  },
  {
    "noun": "Rechenstab",
    "article": "der"
  },
  {
    "noun": "Rechenunterricht",
    "article": "der"
  },
  {
    "noun": "Rechercheur",
    "article": "der"
  },
  {
    "noun": "Rechner",
    "article": "der"
  },
  {
    "noun": "Rechnungsbetrag",
    "article": "der"
  },
  {
    "noun": "Rechnungshof",
    "article": "der"
  },
  {
    "noun": "Rechnungsposten",
    "article": "der"
  },
  {
    "noun": "Rechtfertigungsversuch",
    "article": "der"
  },
  {
    "noun": "Rechtsabbieger",
    "article": "der"
  },
  {
    "noun": "Rechtsanspruch",
    "article": "der"
  },
  {
    "noun": "Rechtsanwalt",
    "article": "der"
  },
  {
    "noun": "Rechtsanwaltsgehilfe",
    "article": "der"
  },
  {
    "noun": "Rechtsausleger",
    "article": "der"
  },
  {
    "noun": "Rechtsausschuss",
    "article": "der"
  },
  {
    "noun": "Rechtsbegriff",
    "article": "der"
  },
  {
    "noun": "Rechtsbehelf",
    "article": "der"
  },
  {
    "noun": "Rechtsbeistand",
    "article": "der"
  },
  {
    "noun": "Rechtsberater",
    "article": "der"
  },
  {
    "noun": "Rechtsbrauch",
    "article": "der"
  },
  {
    "noun": "Rechtsbrecher",
    "article": "der"
  },
  {
    "noun": "Rechtsbruch",
    "article": "der"
  },
  {
    "noun": "Rechtschreibfehler",
    "article": "der"
  },
  {
    "noun": "Rechtschreibungsfehler",
    "article": "der"
  },
  {
    "noun": "Rechtsdrall",
    "article": "der"
  },
  {
    "noun": "Rechtsfall",
    "article": "der"
  },
  {
    "noun": "Rechtsgang",
    "article": "der"
  },
  {
    "noun": "Rechtsgrund",
    "article": "der"
  },
  {
    "noun": "Rechtsgrundsatz",
    "article": "der"
  },
  {
    "noun": "Rechtshandel",
    "article": "der"
  },
  {
    "noun": "Rechtsirrtum",
    "article": "der"
  },
  {
    "noun": "Rechtslenker",
    "article": "der"
  },
  {
    "noun": "Rechtsnachfolger",
    "article": "der"
  },
  {
    "noun": "Rechtspfleger",
    "article": "der"
  },
  {
    "noun": "Rechtsradikalismus",
    "article": "der"
  },
  {
    "noun": "Rechtsruck",
    "article": "der"
  },
  {
    "noun": "Rechtsschutz",
    "article": "der"
  },
  {
    "noun": "Rechtsspruch",
    "article": "der"
  },
  {
    "noun": "Rechtsstaat",
    "article": "der"
  },
  {
    "noun": "Rechtsstandpunkt",
    "article": "der"
  },
  {
    "noun": "Rechtsstreit",
    "article": "der"
  },
  {
    "noun": "Rechtstitel",
    "article": "der"
  },
  {
    "noun": "Rechtsverkehr",
    "article": "der"
  },
  {
    "noun": "Rechtsvertreter",
    "article": "der"
  },
  {
    "noun": "Rechtsweg",
    "article": "der"
  },
  {
    "noun": "Recorder",
    "article": "der"
  },
  {
    "noun": "Redakteur",
    "article": "der"
  },
  {
    "noun": "Redaktionsschluss",
    "article": "der"
  },
  {
    "noun": "Redaktor",
    "article": "der"
  },
  {
    "noun": "Redeschwall",
    "article": "der"
  },
  {
    "noun": "Rediskont",
    "article": "der"
  },
  {
    "noun": "Redner",
    "article": "der"
  },
  {
    "noun": "Reduktionismus",
    "article": "der"
  },
  {
    "noun": "Reduktor",
    "article": "der"
  },
  {
    "noun": "Reeder",
    "article": "der"
  },
  {
    "noun": "Reeper",
    "article": "der"
  },
  {
    "noun": "Reexport",
    "article": "der"
  },
  {
    "noun": "Referendar",
    "article": "der"
  },
  {
    "noun": "Referent",
    "article": "der"
  },
  {
    "noun": "Reflektant",
    "article": "der"
  },
  {
    "noun": "Reflektor",
    "article": "der"
  },
  {
    "noun": "Reflex",
    "article": "der"
  },
  {
    "noun": "Reflexionsgrad",
    "article": "der"
  },
  {
    "noun": "Reflexionsnebel",
    "article": "der"
  },
  {
    "noun": "Reflexionswinkel",
    "article": "der"
  },
  {
    "noun": "Reflux",
    "article": "der"
  },
  {
    "noun": "Reformator",
    "article": "der"
  },
  {
    "noun": "Reformer",
    "article": "der"
  },
  {
    "noun": "Reformismus",
    "article": "der"
  },
  {
    "noun": "Reformplan",
    "article": "der"
  },
  {
    "noun": "Refrain",
    "article": "der"
  },
  {
    "noun": "Refraktor",
    "article": "der"
  },
  {
    "noun": "Refrigerator",
    "article": "der"
  },
  {
    "noun": "Regelfall",
    "article": "der"
  },
  {
    "noun": "Regelkreis",
    "article": "der"
  },
  {
    "noun": "Regelsatz",
    "article": "der"
  },
  {
    "noun": "Regenbogen",
    "article": "der"
  },
  {
    "noun": "Regenerator",
    "article": "der"
  },
  {
    "noun": "Regenfall",
    "article": "der"
  },
  {
    "noun": "Regenguss",
    "article": "der"
  },
  {
    "noun": "Regenmacher",
    "article": "der"
  },
  {
    "noun": "Regenmantel",
    "article": "der"
  },
  {
    "noun": "Regenmesser",
    "article": "der"
  },
  {
    "noun": "Regenpfeifer",
    "article": "der"
  },
  {
    "noun": "Regenschatten",
    "article": "der"
  },
  {
    "noun": "Regenschauer",
    "article": "der"
  },
  {
    "noun": "Regenschirm",
    "article": "der"
  },
  {
    "noun": "Regenschreiber",
    "article": "der"
  },
  {
    "noun": "Regenschutz",
    "article": "der"
  },
  {
    "noun": "Regent",
    "article": "der"
  },
  {
    "noun": "Regentag",
    "article": "der"
  },
  {
    "noun": "Regentropfen",
    "article": "der"
  },
  {
    "noun": "Regenumhang",
    "article": "der"
  },
  {
    "noun": "Regenwald",
    "article": "der"
  },
  {
    "noun": "Regenwurm",
    "article": "der"
  },
  {
    "noun": "Reggae",
    "article": "der"
  },
  {
    "noun": "Regieassistent",
    "article": "der"
  },
  {
    "noun": "Regiebetrieb",
    "article": "der"
  },
  {
    "noun": "Regiefehler",
    "article": "der"
  },
  {
    "noun": "Regierungsantritt",
    "article": "der"
  },
  {
    "noun": "Regierungsausschuss",
    "article": "der"
  },
  {
    "noun": "Regierungsbezirk",
    "article": "der"
  },
  {
    "noun": "Regierungschef",
    "article": "der"
  },
  {
    "noun": "Regierungsrat",
    "article": "der"
  },
  {
    "noun": "Regierungssitz",
    "article": "der"
  },
  {
    "noun": "Regierungssprecher",
    "article": "der"
  },
  {
    "noun": "Regierungsvertreter",
    "article": "der"
  },
  {
    "noun": "Regierungswechsel",
    "article": "der"
  },
  {
    "noun": "Regimekritiker",
    "article": "der"
  },
  {
    "noun": "Regimentsarzt",
    "article": "der"
  },
  {
    "noun": "Regionalismus",
    "article": "der"
  },
  {
    "noun": "Regisseur",
    "article": "der"
  },
  {
    "noun": "Registrator",
    "article": "der"
  },
  {
    "noun": "Registrierapparat",
    "article": "der"
  },
  {
    "noun": "Regler",
    "article": "der"
  },
  {
    "noun": "Regress",
    "article": "der"
  },
  {
    "noun": "Regressand",
    "article": "der"
  },
  {
    "noun": "Regressanspruch",
    "article": "der"
  },
  {
    "noun": "Regressor",
    "article": "der"
  },
  {
    "noun": "Regulator",
    "article": "der"
  },
  {
    "noun": "Rehbock",
    "article": "der"
  },
  {
    "noun": "Rehbraten",
    "article": "der"
  },
  {
    "noun": "Rehposten",
    "article": "der"
  },
  {
    "noun": "Reibekuchen",
    "article": "der"
  },
  {
    "noun": "Reibelaut",
    "article": "der"
  },
  {
    "noun": "Reibungskoeffizient",
    "article": "der"
  },
  {
    "noun": "Reibungsverlust",
    "article": "der"
  },
  {
    "noun": "Reibungswiderstand",
    "article": "der"
  },
  {
    "noun": "Reichsapfel",
    "article": "der"
  },
  {
    "noun": "Reichskanzler",
    "article": "der"
  },
  {
    "noun": "Reichstag",
    "article": "der"
  },
  {
    "noun": "Reichsverweser",
    "article": "der"
  },
  {
    "noun": "Reichtum",
    "article": "der"
  },
  {
    "noun": "Reif",
    "article": "der"
  },
  {
    "noun": "Reifendruck",
    "article": "der"
  },
  {
    "noun": "Reifenplatzer",
    "article": "der"
  },
  {
    "noun": "Reifenschaden",
    "article": "der"
  },
  {
    "noun": "Reifenwechsel",
    "article": "der"
  },
  {
    "noun": "Reifeprozess",
    "article": "der"
  },
  {
    "noun": "Reigen",
    "article": "der"
  },
  {
    "noun": "Reihenmotor",
    "article": "der"
  },
  {
    "noun": "Reiher",
    "article": "der"
  },
  {
    "noun": "Reim",
    "article": "der"
  },
  {
    "noun": "Reimport",
    "article": "der"
  },
  {
    "noun": "Reinbestand",
    "article": "der"
  },
  {
    "noun": "Reinbetrag",
    "article": "der"
  },
  {
    "noun": "Reinertrag",
    "article": "der"
  },
  {
    "noun": "Reinfall",
    "article": "der"
  },
  {
    "noun": "Reinfarkt",
    "article": "der"
  },
  {
    "noun": "Reingewinn",
    "article": "der"
  },
  {
    "noun": "Reinheitsgrad",
    "article": "der"
  },
  {
    "noun": "Reiniger",
    "article": "der"
  },
  {
    "noun": "Reinraum",
    "article": "der"
  },
  {
    "noun": "Reinstoff",
    "article": "der"
  },
  {
    "noun": "Reisbesen",
    "article": "der"
  },
  {
    "noun": "Reisbrei",
    "article": "der"
  },
  {
    "noun": "Reisebedarf",
    "article": "der"
  },
  {
    "noun": "Reisebegleiter",
    "article": "der"
  },
  {
    "noun": "Reisebericht",
    "article": "der"
  },
  {
    "noun": "Reisebus",
    "article": "der"
  },
  {
    "noun": "Reisekoffer",
    "article": "der"
  },
  {
    "noun": "Reisekorb",
    "article": "der"
  },
  {
    "noun": "Reiseleiter",
    "article": "der"
  },
  {
    "noun": "Reisemarschall",
    "article": "der"
  },
  {
    "noun": "Reisepass",
    "article": "der"
  },
  {
    "noun": "Reiseplan",
    "article": "der"
  },
  {
    "noun": "Reiseprospekt",
    "article": "der"
  },
  {
    "noun": "Reiseproviant",
    "article": "der"
  },
  {
    "noun": "Reisescheck",
    "article": "der"
  },
  {
    "noun": "Reiseveranstalter",
    "article": "der"
  },
  {
    "noun": "Reiseverkehr",
    "article": "der"
  },
  {
    "noun": "Reiseverkehrskaufmann",
    "article": "der"
  },
  {
    "noun": "Reisevertreter",
    "article": "der"
  },
  {
    "noun": "Reisewecker",
    "article": "der"
  },
  {
    "noun": "Reiseweg",
    "article": "der"
  },
  {
    "noun": "Reisezug",
    "article": "der"
  },
  {
    "noun": "Reisezugwagen",
    "article": "der"
  },
  {
    "noun": "Reisigbesen",
    "article": "der"
  },
  {
    "noun": "Reitanzug",
    "article": "der"
  },
  {
    "noun": "Reitdress",
    "article": "der"
  },
  {
    "noun": "Reiter",
    "article": "der"
  },
  {
    "noun": "Reiterangriff",
    "article": "der"
  },
  {
    "noun": "Reitknecht",
    "article": "der"
  },
  {
    "noun": "Reitsattel",
    "article": "der"
  },
  {
    "noun": "Reitsport",
    "article": "der"
  },
  {
    "noun": "Reitstall",
    "article": "der"
  },
  {
    "noun": "Reitstiefel",
    "article": "der"
  },
  {
    "noun": "Reitstock",
    "article": "der"
  },
  {
    "noun": "Reitunterricht",
    "article": "der"
  },
  {
    "noun": "Reitweg",
    "article": "der"
  },
  {
    "noun": "Reiz",
    "article": "der"
  },
  {
    "noun": "Reizhusten",
    "article": "der"
  },
  {
    "noun": "Reizker",
    "article": "der"
  },
  {
    "noun": "Reizstrom",
    "article": "der"
  },
  {
    "noun": "Reklamant",
    "article": "der"
  },
  {
    "noun": "Reklamefeldzug",
    "article": "der"
  },
  {
    "noun": "Reklamerummel",
    "article": "der"
  },
  {
    "noun": "Reklamezettel",
    "article": "der"
  },
  {
    "noun": "Rekord",
    "article": "der"
  },
  {
    "noun": "Rekordbesuch",
    "article": "der"
  },
  {
    "noun": "Rekordhalter",
    "article": "der"
  },
  {
    "noun": "Rekordinhaber",
    "article": "der"
  },
  {
    "noun": "Rekordler",
    "article": "der"
  },
  {
    "noun": "Rekordversuch",
    "article": "der"
  },
  {
    "noun": "Rekrut",
    "article": "der"
  },
  {
    "noun": "Rektawechsel",
    "article": "der"
  },
  {
    "noun": "Rektor",
    "article": "der"
  },
  {
    "noun": "Rekuperator",
    "article": "der"
  },
  {
    "noun": "Rekurs",
    "article": "der"
  },
  {
    "noun": "Relaissatellit",
    "article": "der"
  },
  {
    "noun": "Relationsbegriff",
    "article": "der"
  },
  {
    "noun": "Relativismus",
    "article": "der"
  },
  {
    "noun": "Relativsatz",
    "article": "der"
  },
  {
    "noun": "Reliefdruck",
    "article": "der"
  },
  {
    "noun": "Reliefglobus",
    "article": "der"
  },
  {
    "noun": "Religionskrieg",
    "article": "der"
  },
  {
    "noun": "Religionslehrer",
    "article": "der"
  },
  {
    "noun": "Religionsstifter",
    "article": "der"
  },
  {
    "noun": "Religionsstreit",
    "article": "der"
  },
  {
    "noun": "Religionsunterricht",
    "article": "der"
  },
  {
    "noun": "Religionswechsel",
    "article": "der"
  },
  {
    "noun": "Reliquienschrein",
    "article": "der"
  },
  {
    "noun": "Rembours",
    "article": "der"
  },
  {
    "noun": "Remigrant",
    "article": "der"
  },
  {
    "noun": "Remittent",
    "article": "der"
  },
  {
    "noun": "Rempler",
    "article": "der"
  },
  {
    "noun": "Rempter",
    "article": "der"
  },
  {
    "noun": "Renegat",
    "article": "der"
  },
  {
    "noun": "Renner",
    "article": "der"
  },
  {
    "noun": "Rennfahrer",
    "article": "der"
  },
  {
    "noun": "Rennreifen",
    "article": "der"
  },
  {
    "noun": "Rennsattel",
    "article": "der"
  },
  {
    "noun": "Rennschlitten",
    "article": "der"
  },
  {
    "noun": "Rennsport",
    "article": "der"
  },
  {
    "noun": "Rennstall",
    "article": "der"
  },
  {
    "noun": "Renntag",
    "article": "der"
  },
  {
    "noun": "Rennwagen",
    "article": "der"
  },
  {
    "noun": "Renommist",
    "article": "der"
  },
  {
    "noun": "Rentenanspruch",
    "article": "der"
  },
  {
    "noun": "Rentenberater",
    "article": "der"
  },
  {
    "noun": "Rentenmarkt",
    "article": "der"
  },
  {
    "noun": "Rentner",
    "article": "der"
  },
  {
    "noun": "Reorganisator",
    "article": "der"
  },
  {
    "noun": "Rep",
    "article": "der"
  },
  {
    "noun": "Reparateur",
    "article": "der"
  },
  {
    "noun": "Repeater",
    "article": "der"
  },
  {
    "noun": "Repetent",
    "article": "der"
  },
  {
    "noun": "Repetitor",
    "article": "der"
  },
  {
    "noun": "Report",
    "article": "der"
  },
  {
    "noun": "Reporter",
    "article": "der"
  },
  {
    "noun": "Reprint",
    "article": "der"
  },
  {
    "noun": "Reproduktionsfaktor",
    "article": "der"
  },
  {
    "noun": "Republikaner",
    "article": "der"
  },
  {
    "noun": "Repulsionsmotor",
    "article": "der"
  },
  {
    "noun": "Requisiteur",
    "article": "der"
  },
  {
    "noun": "Reservefonds",
    "article": "der"
  },
  {
    "noun": "Reserveoffizier",
    "article": "der"
  },
  {
    "noun": "Reservereifen",
    "article": "der"
  },
  {
    "noun": "Reservespieler",
    "article": "der"
  },
  {
    "noun": "Reservetank",
    "article": "der"
  },
  {
    "noun": "Reserveteil",
    "article": "der"
  },
  {
    "noun": "Reservist",
    "article": "der"
  },
  {
    "noun": "Resonanzboden",
    "article": "der"
  },
  {
    "noun": "Resonanzkasten",
    "article": "der"
  },
  {
    "noun": "Resonanzraum",
    "article": "der"
  },
  {
    "noun": "Resonator",
    "article": "der"
  },
  {
    "noun": "Respekt",
    "article": "der"
  },
  {
    "noun": "Respirationsapparat",
    "article": "der"
  },
  {
    "noun": "Respirator",
    "article": "der"
  },
  {
    "noun": "Ressortleiter",
    "article": "der"
  },
  {
    "noun": "Rest",
    "article": "der"
  },
  {
    "noun": "Restalkohol",
    "article": "der"
  },
  {
    "noun": "Restant",
    "article": "der"
  },
  {
    "noun": "Restaurantbesitzer",
    "article": "der"
  },
  {
    "noun": "Restaurateur",
    "article": "der"
  },
  {
    "noun": "Restaurator",
    "article": "der"
  },
  {
    "noun": "Restbestand",
    "article": "der"
  },
  {
    "noun": "Restbetrag",
    "article": "der"
  },
  {
    "noun": "Resteverkauf",
    "article": "der"
  },
  {
    "noun": "Restmagnetismus",
    "article": "der"
  },
  {
    "noun": "Restposten",
    "article": "der"
  },
  {
    "noun": "Reststoff",
    "article": "der"
  },
  {
    "noun": "Resturlaub",
    "article": "der"
  },
  {
    "noun": "Restwert",
    "article": "der"
  },
  {
    "noun": "Restzucker",
    "article": "der"
  },
  {
    "noun": "Retentionsraum",
    "article": "der"
  },
  {
    "noun": "Retortengrafit",
    "article": "der"
  },
  {
    "noun": "Retsina",
    "article": "der"
  },
  {
    "noun": "Retter",
    "article": "der"
  },
  {
    "noun": "Rettich",
    "article": "der"
  },
  {
    "noun": "Rettichsalat",
    "article": "der"
  },
  {
    "noun": "Rettungsanker",
    "article": "der"
  },
  {
    "noun": "Rettungsassistent",
    "article": "der"
  },
  {
    "noun": "Rettungsdienst",
    "article": "der"
  },
  {
    "noun": "Rettungshubschrauber",
    "article": "der"
  },
  {
    "noun": "Rettungsring",
    "article": "der"
  },
  {
    "noun": "Rettungsschwimmer",
    "article": "der"
  },
  {
    "noun": "Rettungsversuch",
    "article": "der"
  },
  {
    "noun": "Rettungswagen",
    "article": "der"
  },
  {
    "noun": "Retuscheur",
    "article": "der"
  },
  {
    "noun": "Revanchekrieg",
    "article": "der"
  },
  {
    "noun": "Revanchismus",
    "article": "der"
  },
  {
    "noun": "Revanchist",
    "article": "der"
  },
  {
    "noun": "Revisionismus",
    "article": "der"
  },
  {
    "noun": "Revisionist",
    "article": "der"
  },
  {
    "noun": "Revisor",
    "article": "der"
  },
  {
    "noun": "Revoluzzer",
    "article": "der"
  },
  {
    "noun": "Revolver",
    "article": "der"
  },
  {
    "noun": "Revolverheld",
    "article": "der"
  },
  {
    "noun": "Revolverkopf",
    "article": "der"
  },
  {
    "noun": "Revuefilm",
    "article": "der"
  },
  {
    "noun": "Rezensent",
    "article": "der"
  },
  {
    "noun": "Rezeptor",
    "article": "der"
  },
  {
    "noun": "Rezess",
    "article": "der"
  },
  {
    "noun": "Rezipient",
    "article": "der"
  },
  {
    "noun": "Rezitator",
    "article": "der"
  },
  {
    "noun": "Rhabarberkuchen",
    "article": "der"
  },
  {
    "noun": "Rhapsode",
    "article": "der"
  },
  {
    "noun": "Rhein",
    "article": "der"
  },
  {
    "noun": "Rheinfall",
    "article": "der"
  },
  {
    "noun": "Rheinkiesel",
    "article": "der"
  },
  {
    "noun": "Rheinwein",
    "article": "der"
  },
  {
    "noun": "Rheostat",
    "article": "der"
  },
  {
    "noun": "Rhesusfaktor",
    "article": "der"
  },
  {
    "noun": "Rhetor",
    "article": "der"
  },
  {
    "noun": "Rhetoriker",
    "article": "der"
  },
  {
    "noun": "Rheumatiker",
    "article": "der"
  },
  {
    "noun": "Rheumatismus",
    "article": "der"
  },
  {
    "noun": "Rheumatologe",
    "article": "der"
  },
  {
    "noun": "Rhinologe",
    "article": "der"
  },
  {
    "noun": "Rhombus",
    "article": "der"
  },
  {
    "noun": "Rhonchus",
    "article": "der"
  },
  {
    "noun": "Rhythmus",
    "article": "der"
  },
  {
    "noun": "Rial",
    "article": "der"
  },
  {
    "noun": "RIAS",
    "article": "der"
  },
  {
    "noun": "Richter",
    "article": "der"
  },
  {
    "noun": "Richterspruch",
    "article": "der"
  },
  {
    "noun": "Richterstand",
    "article": "der"
  },
  {
    "noun": "Richterstuhl",
    "article": "der"
  },
  {
    "noun": "Richtfunk",
    "article": "der"
  },
  {
    "noun": "Richtkreis",
    "article": "der"
  },
  {
    "noun": "Richtmeister",
    "article": "der"
  },
  {
    "noun": "Richtplatz",
    "article": "der"
  },
  {
    "noun": "Richtpreis",
    "article": "der"
  },
  {
    "noun": "Richtpunkt",
    "article": "der"
  },
  {
    "noun": "Richtsatz",
    "article": "der"
  },
  {
    "noun": "Richtstrahler",
    "article": "der"
  },
  {
    "noun": "Richtungsanzeiger",
    "article": "der"
  },
  {
    "noun": "Richtungspfeil",
    "article": "der"
  },
  {
    "noun": "Richtungsverkehr",
    "article": "der"
  },
  {
    "noun": "Richtungswechsel",
    "article": "der"
  },
  {
    "noun": "Richtwert",
    "article": "der"
  },
  {
    "noun": "Riecher",
    "article": "der"
  },
  {
    "noun": "Riechkolben",
    "article": "der"
  },
  {
    "noun": "Riechnerv",
    "article": "der"
  },
  {
    "noun": "Riechstoff",
    "article": "der"
  },
  {
    "noun": "Riedbock",
    "article": "der"
  },
  {
    "noun": "Riegel",
    "article": "der"
  },
  {
    "noun": "Riemen",
    "article": "der"
  },
  {
    "noun": "Riemenantrieb",
    "article": "der"
  },
  {
    "noun": "Riementrieb",
    "article": "der"
  },
  {
    "noun": "Riemer",
    "article": "der"
  },
  {
    "noun": "Riesenbau",
    "article": "der"
  },
  {
    "noun": "Riesenerfolg",
    "article": "der"
  },
  {
    "noun": "Riesenschritt",
    "article": "der"
  },
  {
    "noun": "Riesenslalom",
    "article": "der"
  },
  {
    "noun": "Riesenstern",
    "article": "der"
  },
  {
    "noun": "Riesentorlauf",
    "article": "der"
  },
  {
    "noun": "Riesenwuchs",
    "article": "der"
  },
  {
    "noun": "Riesling",
    "article": "der"
  },
  {
    "noun": "Rigolpflug",
    "article": "der"
  },
  {
    "noun": "Rigorismus",
    "article": "der"
  },
  {
    "noun": "Rigorosum",
    "article": "der"
  },
  {
    "noun": "Rinderbraten",
    "article": "der"
  },
  {
    "noun": "Rindertalg",
    "article": "der"
  },
  {
    "noun": "Rinderwahn",
    "article": "der"
  },
  {
    "noun": "Rinderwahnsinn",
    "article": "der"
  },
  {
    "noun": "Rindsbraten",
    "article": "der"
  },
  {
    "noun": "Ringel",
    "article": "der"
  },
  {
    "noun": "Ringelschwanz",
    "article": "der"
  },
  {
    "noun": "Ringelwurm",
    "article": "der"
  },
  {
    "noun": "Ringer",
    "article": "der"
  },
  {
    "noun": "Ringfinger",
    "article": "der"
  },
  {
    "noun": "Ringkampf",
    "article": "der"
  },
  {
    "noun": "Ringofen",
    "article": "der"
  },
  {
    "noun": "Ringtausch",
    "article": "der"
  },
  {
    "noun": "Ringwall",
    "article": "der"
  },
  {
    "noun": "Ringwechsel",
    "article": "der"
  },
  {
    "noun": "Rinnstein",
    "article": "der"
  },
  {
    "noun": "Rippenbogen",
    "article": "der"
  },
  {
    "noun": "Rippenbruch",
    "article": "der"
  },
  {
    "noun": "Rippenknochen",
    "article": "der"
  },
  {
    "noun": "Rippenknorpel",
    "article": "der"
  },
  {
    "noun": "Rips",
    "article": "der"
  },
  {
    "noun": "Risikofaktor",
    "article": "der"
  },
  {
    "noun": "Risikopatient",
    "article": "der"
  },
  {
    "noun": "Riss",
    "article": "der"
  },
  {
    "noun": "Rist",
    "article": "der"
  },
  {
    "noun": "Ritt",
    "article": "der"
  },
  {
    "noun": "Rittberger",
    "article": "der"
  },
  {
    "noun": "Ritter",
    "article": "der"
  },
  {
    "noun": "Ritterdienst",
    "article": "der"
  },
  {
    "noun": "Ritterfalter",
    "article": "der"
  },
  {
    "noun": "Ritterling",
    "article": "der"
  },
  {
    "noun": "Ritterorden",
    "article": "der"
  },
  {
    "noun": "Ritterschlag",
    "article": "der"
  },
  {
    "noun": "Rittersmann",
    "article": "der"
  },
  {
    "noun": "Rittersporn",
    "article": "der"
  },
  {
    "noun": "Ritterstand",
    "article": "der"
  },
  {
    "noun": "Ritterstern",
    "article": "der"
  },
  {
    "noun": "Rittmeister",
    "article": "der"
  },
  {
    "noun": "Ritualismus",
    "article": "der"
  },
  {
    "noun": "Ritualist",
    "article": "der"
  },
  {
    "noun": "Ritus",
    "article": "der"
  },
  {
    "noun": "Ritz",
    "article": "der"
  },
  {
    "noun": "Rivale",
    "article": "der"
  },
  {
    "noun": "Rizinus",
    "article": "der"
  },
  {
    "noun": "Roadie",
    "article": "der"
  },
  {
    "noun": "Robbenfang",
    "article": "der"
  },
  {
    "noun": "Robbenschlag",
    "article": "der"
  },
  {
    "noun": "Robber",
    "article": "der"
  },
  {
    "noun": "Roboter",
    "article": "der"
  },
  {
    "noun": "Rochen",
    "article": "der"
  },
  {
    "noun": "Rockaufschlag",
    "article": "der"
  },
  {
    "noun": "Rockbund",
    "article": "der"
  },
  {
    "noun": "Rocker",
    "article": "der"
  },
  {
    "noun": "Rocksaum",
    "article": "der"
  },
  {
    "noun": "Rockstar",
    "article": "der"
  },
  {
    "noun": "Rockzipfel",
    "article": "der"
  },
  {
    "noun": "Rodelschlitten",
    "article": "der"
  },
  {
    "noun": "Rodelsport",
    "article": "der"
  },
  {
    "noun": "Rodler",
    "article": "der"
  },
  {
    "noun": "Rogen",
    "article": "der"
  },
  {
    "noun": "Rogenstein",
    "article": "der"
  },
  {
    "noun": "Roggen",
    "article": "der"
  },
  {
    "noun": "Rohbau",
    "article": "der"
  },
  {
    "noun": "Rohdiamant",
    "article": "der"
  },
  {
    "noun": "Rohentwurf",
    "article": "der"
  },
  {
    "noun": "Rohertrag",
    "article": "der"
  },
  {
    "noun": "Rohgewinn",
    "article": "der"
  },
  {
    "noun": "Rohkaffee",
    "article": "der"
  },
  {
    "noun": "Rohling",
    "article": "der"
  },
  {
    "noun": "Rohrabschneider",
    "article": "der"
  },
  {
    "noun": "Rohrbruch",
    "article": "der"
  },
  {
    "noun": "Rohrkrepierer",
    "article": "der"
  },
  {
    "noun": "Rohrrahmen",
    "article": "der"
  },
  {
    "noun": "Rohrspatz",
    "article": "der"
  },
  {
    "noun": "Rohrstock",
    "article": "der"
  },
  {
    "noun": "Rohrzucker",
    "article": "der"
  },
  {
    "noun": "Rohstahl",
    "article": "der"
  },
  {
    "noun": "Rohstoff",
    "article": "der"
  },
  {
    "noun": "Rohstoffmangel",
    "article": "der"
  },
  {
    "noun": "Rohstoffpreis",
    "article": "der"
  },
  {
    "noun": "Rohumsatz",
    "article": "der"
  },
  {
    "noun": "Rohzucker",
    "article": "der"
  },
  {
    "noun": "Rohzustand",
    "article": "der"
  },
  {
    "noun": "Rolladenschrank",
    "article": "der"
  },
  {
    "noun": "Rollbalken",
    "article": "der"
  },
  {
    "noun": "Rollball",
    "article": "der"
  },
  {
    "noun": "Rollentausch",
    "article": "der"
  },
  {
    "noun": "Rollentext",
    "article": "der"
  },
  {
    "noun": "Roller",
    "article": "der"
  },
  {
    "noun": "Rollfilm",
    "article": "der"
  },
  {
    "noun": "Rolli",
    "article": "der"
  },
  {
    "noun": "Rollkragen",
    "article": "der"
  },
  {
    "noun": "Rollkragenpullover",
    "article": "der"
  },
  {
    "noun": "Rollmops",
    "article": "der"
  },
  {
    "noun": "Rollschinken",
    "article": "der"
  },
  {
    "noun": "Rollschrank",
    "article": "der"
  },
  {
    "noun": "Rollschuh",
    "article": "der"
  },
  {
    "noun": "Rollstuhl",
    "article": "der"
  },
  {
    "noun": "Rollstuhlfahrer",
    "article": "der"
  },
  {
    "noun": "Rollwagen",
    "article": "der"
  },
  {
    "noun": "Roman",
    "article": "der"
  },
  {
    "noun": "Romanautor",
    "article": "der"
  },
  {
    "noun": "Romancier",
    "article": "der"
  },
  {
    "noun": "Romanheld",
    "article": "der"
  },
  {
    "noun": "Romanist",
    "article": "der"
  },
  {
    "noun": "Romanleser",
    "article": "der"
  },
  {
    "noun": "Romanschreiber",
    "article": "der"
  },
  {
    "noun": "Romanschriftsteller",
    "article": "der"
  },
  {
    "noun": "Romantiker",
    "article": "der"
  },
  {
    "noun": "Romantizismus",
    "article": "der"
  },
  {
    "noun": "Roquefort",
    "article": "der"
  },
  {
    "noun": "Rosenduft",
    "article": "der"
  },
  {
    "noun": "Rosengarten",
    "article": "der"
  },
  {
    "noun": "Rosenkohl",
    "article": "der"
  },
  {
    "noun": "Rosenkranz",
    "article": "der"
  },
  {
    "noun": "Rosenmontag",
    "article": "der"
  },
  {
    "noun": "Rosenstock",
    "article": "der"
  },
  {
    "noun": "Rosenstrauch",
    "article": "der"
  },
  {
    "noun": "Rosmarin",
    "article": "der"
  },
  {
    "noun": "Rossapfel",
    "article": "der"
  },
  {
    "noun": "Rosselenker",
    "article": "der"
  },
  {
    "noun": "Rosskamm",
    "article": "der"
  },
  {
    "noun": "Rost",
    "article": "der"
  },
  {
    "noun": "Rostansatz",
    "article": "der"
  },
  {
    "noun": "Rostbefall",
    "article": "der"
  },
  {
    "noun": "Rostbraten",
    "article": "der"
  },
  {
    "noun": "Rostfleck",
    "article": "der"
  },
  {
    "noun": "Rostpilz",
    "article": "der"
  },
  {
    "noun": "Rostschutz",
    "article": "der"
  },
  {
    "noun": "Rostschutzanstrich",
    "article": "der"
  },
  {
    "noun": "Rostumwandler",
    "article": "der"
  },
  {
    "noun": "Rotang",
    "article": "der"
  },
  {
    "noun": "Rotationsdruck",
    "article": "der"
  },
  {
    "noun": "Rotationskolbenmotor",
    "article": "der"
  },
  {
    "noun": "Rotator",
    "article": "der"
  },
  {
    "noun": "Rotbarsch",
    "article": "der"
  },
  {
    "noun": "Rotdorn",
    "article": "der"
  },
  {
    "noun": "Rotfuchs",
    "article": "der"
  },
  {
    "noun": "Rotguss",
    "article": "der"
  },
  {
    "noun": "Rothirsch",
    "article": "der"
  },
  {
    "noun": "Rotklee",
    "article": "der"
  },
  {
    "noun": "Rotkohl",
    "article": "der"
  },
  {
    "noun": "Rotkopf",
    "article": "der"
  },
  {
    "noun": "Rotlauf",
    "article": "der"
  },
  {
    "noun": "Rotlichtbezirk",
    "article": "der"
  },
  {
    "noun": "Rotmilan",
    "article": "der"
  },
  {
    "noun": "Rotor",
    "article": "der"
  },
  {
    "noun": "Rotschimmel",
    "article": "der"
  },
  {
    "noun": "Rotschopf",
    "article": "der"
  },
  {
    "noun": "Rotstift",
    "article": "der"
  },
  {
    "noun": "Rottweiler",
    "article": "der"
  },
  {
    "noun": "Rotwein",
    "article": "der"
  },
  {
    "noun": "Rotweinfleck",
    "article": "der"
  },
  {
    "noun": "Rotz",
    "article": "der"
  },
  {
    "noun": "Router",
    "article": "der"
  },
  {
    "noun": "Routinier",
    "article": "der"
  },
  {
    "noun": "Rowdy",
    "article": "der"
  },
  {
    "noun": "Royalist",
    "article": "der"
  },
  {
    "noun": "Rubel",
    "article": "der"
  },
  {
    "noun": "Rubin",
    "article": "der"
  },
  {
    "noun": "Ruch",
    "article": "der"
  },
  {
    "noun": "Ruck",
    "article": "der"
  },
  {
    "noun": "Rucksack",
    "article": "der"
  },
  {
    "noun": "Rucksacktourist",
    "article": "der"
  },
  {
    "noun": "Ruderclub",
    "article": "der"
  },
  {
    "noun": "Ruderer",
    "article": "der"
  },
  {
    "noun": "Ruderlagenanzeiger",
    "article": "der"
  },
  {
    "noun": "Ruderschaft",
    "article": "der"
  },
  {
    "noun": "Ruderschlag",
    "article": "der"
  },
  {
    "noun": "Rudersport",
    "article": "der"
  },
  {
    "noun": "Ruderverein",
    "article": "der"
  },
  {
    "noun": "Ruf",
    "article": "der"
  },
  {
    "noun": "Rufer",
    "article": "der"
  },
  {
    "noun": "Ruffall",
    "article": "der"
  },
  {
    "noun": "Rufname",
    "article": "der"
  },
  {
    "noun": "Ruheort",
    "article": "der"
  },
  {
    "noun": "Ruheplatz",
    "article": "der"
  },
  {
    "noun": "Ruheposten",
    "article": "der"
  },
  {
    "noun": "Ruhepunkt",
    "article": "der"
  },
  {
    "noun": "Ruheraum",
    "article": "der"
  },
  {
    "noun": "Ruheschmerz",
    "article": "der"
  },
  {
    "noun": "Ruhesitz",
    "article": "der"
  },
  {
    "noun": "Ruhestand",
    "article": "der"
  },
  {
    "noun": "Ruhestifter",
    "article": "der"
  },
  {
    "noun": "Ruhestrom",
    "article": "der"
  },
  {
    "noun": "Ruhetag",
    "article": "der"
  },
  {
    "noun": "Ruhezustand",
    "article": "der"
  },
  {
    "noun": "Ruhm",
    "article": "der"
  },
  {
    "noun": "Ruhmestag",
    "article": "der"
  },
  {
    "noun": "Ruhrpott",
    "article": "der"
  },
  {
    "noun": "Ruin",
    "article": "der"
  },
  {
    "noun": "Rum",
    "article": "der"
  },
  {
    "noun": "Rummel",
    "article": "der"
  },
  {
    "noun": "Rummelplatz",
    "article": "der"
  },
  {
    "noun": "Rumor",
    "article": "der"
  },
  {
    "noun": "Rumpf",
    "article": "der"
  },
  {
    "noun": "Rumtopf",
    "article": "der"
  },
  {
    "noun": "Run",
    "article": "der"
  },
  {
    "noun": "Rundbau",
    "article": "der"
  },
  {
    "noun": "Rundblick",
    "article": "der"
  },
  {
    "noun": "Rundbogen",
    "article": "der"
  },
  {
    "noun": "Rundbrief",
    "article": "der"
  },
  {
    "noun": "Runderlass",
    "article": "der"
  },
  {
    "noun": "Rundflug",
    "article": "der"
  },
  {
    "noun": "Rundfunk",
    "article": "der"
  },
  {
    "noun": "Rundfunkempfang",
    "article": "der"
  },
  {
    "noun": "Rundfunkmechaniker",
    "article": "der"
  },
  {
    "noun": "Rundfunksender",
    "article": "der"
  },
  {
    "noun": "Rundfunksprecher",
    "article": "der"
  },
  {
    "noun": "Rundfunktechniker",
    "article": "der"
  },
  {
    "noun": "Rundfunkteilnehmer",
    "article": "der"
  },
  {
    "noun": "Rundgang",
    "article": "der"
  },
  {
    "noun": "Rundgesang",
    "article": "der"
  },
  {
    "noun": "Rundkurs",
    "article": "der"
  },
  {
    "noun": "Rundlauf",
    "article": "der"
  },
  {
    "noun": "Rundruf",
    "article": "der"
  },
  {
    "noun": "Rundspruch",
    "article": "der"
  },
  {
    "noun": "Rundstab",
    "article": "der"
  },
  {
    "noun": "Runenstein",
    "article": "der"
  },
  {
    "noun": "Rungenwagen",
    "article": "der"
  },
  {
    "noun": "Russe",
    "article": "der"
  },
  {
    "noun": "Russenkittel",
    "article": "der"
  },
  {
    "noun": "Rutil",
    "article": "der"
  },
  {
    "noun": "Rutsch",
    "article": "der"
  },
  {
    "noun": "Haft",
    "article": "der"
  },
  {
    "noun": "Haftbefehl",
    "article": "der"
  },
  {
    "noun": "Haftreifen",
    "article": "der"
  },
  {
    "noun": "Hafturlaub",
    "article": "der"
  },
  {
    "noun": "Hag",
    "article": "der"
  },
  {
    "noun": "Hagebuttentee",
    "article": "der"
  },
  {
    "noun": "Hagedorn",
    "article": "der"
  },
  {
    "noun": "Hagel",
    "article": "der"
  },
  {
    "noun": "Hagelschaden",
    "article": "der"
  },
  {
    "noun": "Hagelschauer",
    "article": "der"
  },
  {
    "noun": "Hagelsturm",
    "article": "der"
  },
  {
    "noun": "Hahn",
    "article": "der"
  },
  {
    "noun": "Hahnenkamm",
    "article": "der"
  },
  {
    "noun": "Hahnenkampf",
    "article": "der"
  },
  {
    "noun": "Hahnenschrei",
    "article": "der"
  },
  {
    "noun": "Hahnrei",
    "article": "der"
  },
  {
    "noun": "Hai",
    "article": "der"
  },
  {
    "noun": "Haifisch",
    "article": "der"
  },
  {
    "noun": "Hain",
    "article": "der"
  },
  {
    "noun": "Haitianer",
    "article": "der"
  },
  {
    "noun": "Hakenwurm",
    "article": "der"
  },
  {
    "noun": "Hakim",
    "article": "der"
  },
  {
    "noun": "Halbaffe",
    "article": "der"
  },
  {
    "noun": "Halbautomat",
    "article": "der"
  },
  {
    "noun": "Halbbogen",
    "article": "der"
  },
  {
    "noun": "Halbbruder",
    "article": "der"
  },
  {
    "noun": "Halbedelstein",
    "article": "der"
  },
  {
    "noun": "Halbgott",
    "article": "der"
  },
  {
    "noun": "Halbkreis",
    "article": "der"
  },
  {
    "noun": "Halblederband",
    "article": "der"
  },
  {
    "noun": "Halbleinenband",
    "article": "der"
  },
  {
    "noun": "Halbleiter",
    "article": "der"
  },
  {
    "noun": "Halbmesser",
    "article": "der"
  },
  {
    "noun": "Halbmond",
    "article": "der"
  },
  {
    "noun": "Halbschatten",
    "article": "der"
  },
  {
    "noun": "Halbschlaf",
    "article": "der"
  },
  {
    "noun": "Halbschuh",
    "article": "der"
  },
  {
    "noun": "Halbstiefel",
    "article": "der"
  },
  {
    "noun": "Halbton",
    "article": "der"
  },
  {
    "noun": "Halbvokal",
    "article": "der"
  },
  {
    "noun": "Halbzeitpfiff",
    "article": "der"
  },
  {
    "noun": "Halfterriemen",
    "article": "der"
  },
  {
    "noun": "Halit",
    "article": "der"
  },
  {
    "noun": "Hall",
    "article": "der"
  },
  {
    "noun": "Halleffekt",
    "article": "der"
  },
  {
    "noun": "Hallodri",
    "article": "der"
  },
  {
    "noun": "Halm",
    "article": "der"
  },
  {
    "noun": "Halo",
    "article": "der"
  },
  {
    "noun": "Halogenscheinwerfer",
    "article": "der"
  },
  {
    "noun": "Halophyt",
    "article": "der"
  },
  {
    "noun": "Hals",
    "article": "der"
  },
  {
    "noun": "Halsabschneider",
    "article": "der"
  },
  {
    "noun": "Halsausschnitt",
    "article": "der"
  },
  {
    "noun": "Halskragen",
    "article": "der"
  },
  {
    "noun": "Halsschmerz",
    "article": "der"
  },
  {
    "noun": "Halsschmuck",
    "article": "der"
  },
  {
    "noun": "Halswirbel",
    "article": "der"
  },
  {
    "noun": "Halswirbelbruch",
    "article": "der"
  },
  {
    "noun": "Halt",
    "article": "der"
  },
  {
    "noun": "Haltegriff",
    "article": "der"
  },
  {
    "noun": "Haltegurt",
    "article": "der"
  },
  {
    "noun": "Halteplatz",
    "article": "der"
  },
  {
    "noun": "Haltepunkt",
    "article": "der"
  },
  {
    "noun": "Halter",
    "article": "der"
  },
  {
    "noun": "Halteriemen",
    "article": "der"
  },
  {
    "noun": "Halunke",
    "article": "der"
  },
  {
    "noun": "Hammel",
    "article": "der"
  },
  {
    "noun": "Hammelbraten",
    "article": "der"
  },
  {
    "noun": "Hammer",
    "article": "der"
  },
  {
    "noun": "Hammerfisch",
    "article": "der"
  },
  {
    "noun": "Hammerhai",
    "article": "der"
  },
  {
    "noun": "Hammerkopf",
    "article": "der"
  },
  {
    "noun": "Hammerschlag",
    "article": "der"
  },
  {
    "noun": "Hammerschmied",
    "article": "der"
  },
  {
    "noun": "Hampelmann",
    "article": "der"
  },
  {
    "noun": "Hamster",
    "article": "der"
  },
  {
    "noun": "Hamsterer",
    "article": "der"
  },
  {
    "noun": "Handapparat",
    "article": "der"
  },
  {
    "noun": "Handarbeiter",
    "article": "der"
  },
  {
    "noun": "Handatlas",
    "article": "der"
  },
  {
    "noun": "Handball",
    "article": "der"
  },
  {
    "noun": "Handballer",
    "article": "der"
  },
  {
    "noun": "Handballspieler",
    "article": "der"
  },
  {
    "noun": "Handbetrieb",
    "article": "der"
  },
  {
    "noun": "Handbohrer",
    "article": "der"
  },
  {
    "noun": "Handdruck",
    "article": "der"
  },
  {
    "noun": "Handel",
    "article": "der"
  },
  {
    "noun": "Handelsagent",
    "article": "der"
  },
  {
    "noun": "Handelsartikel",
    "article": "der"
  },
  {
    "noun": "Handelsbetrieb",
    "article": "der"
  },
  {
    "noun": "Handelsbrauch",
    "article": "der"
  },
  {
    "noun": "Handelsbrief",
    "article": "der"
  },
  {
    "noun": "Handelsdampfer",
    "article": "der"
  },
  {
    "noun": "Handelsgeist",
    "article": "der"
  },
  {
    "noun": "Handelshafen",
    "article": "der"
  },
  {
    "noun": "Handelskauf",
    "article": "der"
  },
  {
    "noun": "Handelskrieg",
    "article": "der"
  },
  {
    "noun": "Handelsmakler",
    "article": "der"
  },
  {
    "noun": "Handelsmann",
    "article": "der"
  },
  {
    "noun": "Handelsname",
    "article": "der"
  },
  {
    "noun": "Handelspartner",
    "article": "der"
  },
  {
    "noun": "Handelsplatz",
    "article": "der"
  },
  {
    "noun": "Handelsrichter",
    "article": "der"
  },
  {
    "noun": "Handelsstand",
    "article": "der"
  },
  {
    "noun": "Handelsverkehr",
    "article": "der"
  },
  {
    "noun": "Handelsvertrag",
    "article": "der"
  },
  {
    "noun": "Handelsvertreter",
    "article": "der"
  },
  {
    "noun": "Handelsweg",
    "article": "der"
  },
  {
    "noun": "Handelswert",
    "article": "der"
  },
  {
    "noun": "Handelszweig",
    "article": "der"
  },
  {
    "noun": "Handfeger",
    "article": "der"
  },
  {
    "noun": "Handgalopp",
    "article": "der"
  },
  {
    "noun": "Handgebrauch",
    "article": "der"
  },
  {
    "noun": "Handgriff",
    "article": "der"
  },
  {
    "noun": "Handhebel",
    "article": "der"
  },
  {
    "noun": "Handkarren",
    "article": "der"
  },
  {
    "noun": "Handkoffer",
    "article": "der"
  },
  {
    "noun": "Handkuss",
    "article": "der"
  },
  {
    "noun": "Handlanger",
    "article": "der"
  },
  {
    "noun": "Handlauf",
    "article": "der"
  },
  {
    "noun": "Handleser",
    "article": "der"
  },
  {
    "noun": "Handlungsbedarf",
    "article": "der"
  },
  {
    "noun": "Handlungsgehilfe",
    "article": "der"
  },
  {
    "noun": "Handlungsspielraum",
    "article": "der"
  },
  {
    "noun": "Handschlag",
    "article": "der"
  },
  {
    "noun": "Handschriftendeuter",
    "article": "der"
  },
  {
    "noun": "Handschuh",
    "article": "der"
  },
  {
    "noun": "Handschutz",
    "article": "der"
  },
  {
    "noun": "Handspiegel",
    "article": "der"
  },
  {
    "noun": "Handstand",
    "article": "der"
  },
  {
    "noun": "Handstreich",
    "article": "der"
  },
  {
    "noun": "Handtuchhalter",
    "article": "der"
  },
  {
    "noun": "Handwagen",
    "article": "der"
  },
  {
    "noun": "Handwerker",
    "article": "der"
  },
  {
    "noun": "Handwerkerstand",
    "article": "der"
  },
  {
    "noun": "Handwerksberuf",
    "article": "der"
  },
  {
    "noun": "Handwerksbetrieb",
    "article": "der"
  },
  {
    "noun": "Handwerksmann",
    "article": "der"
  },
  {
    "noun": "Handwerksmeister",
    "article": "der"
  },
  {
    "noun": "Handwurzelknochen",
    "article": "der"
  },
  {
    "noun": "Handzettel",
    "article": "der"
  },
  {
    "noun": "Hanf",
    "article": "der"
  },
  {
    "noun": "Hanfanbau",
    "article": "der"
  },
  {
    "noun": "Hanfsame",
    "article": "der"
  },
  {
    "noun": "Hanfstengel",
    "article": "der"
  },
  {
    "noun": "Hang",
    "article": "der"
  },
  {
    "noun": "Hangar",
    "article": "der"
  },
  {
    "noun": "Hanger",
    "article": "der"
  },
  {
    "noun": "Hanseat",
    "article": "der"
  },
  {
    "noun": "Hansnarr",
    "article": "der"
  },
  {
    "noun": "Hardcopy",
    "article": "der"
  },
  {
    "noun": "Hardliner",
    "article": "der"
  },
  {
    "noun": "Harem",
    "article": "der"
  },
  {
    "noun": "Harfenist",
    "article": "der"
  },
  {
    "noun": "Harfenklang",
    "article": "der"
  },
  {
    "noun": "Harfenspieler",
    "article": "der"
  },
  {
    "noun": "Harlekin",
    "article": "der"
  },
  {
    "noun": "Harm",
    "article": "der"
  },
  {
    "noun": "Harn",
    "article": "der"
  },
  {
    "noun": "Harndrang",
    "article": "der"
  },
  {
    "noun": "Harnisch",
    "article": "der"
  },
  {
    "noun": "Harnleiter",
    "article": "der"
  },
  {
    "noun": "Harnstoff",
    "article": "der"
  },
  {
    "noun": "Harnweg",
    "article": "der"
  },
  {
    "noun": "Harpunier",
    "article": "der"
  },
  {
    "noun": "Harsch",
    "article": "der"
  },
  {
    "noun": "Harschschnee",
    "article": "der"
  },
  {
    "noun": "Hartbeton",
    "article": "der"
  },
  {
    "noun": "Hartbrandziegel",
    "article": "der"
  },
  {
    "noun": "Hartguss",
    "article": "der"
  },
  {
    "noun": "Hartriegel",
    "article": "der"
  },
  {
    "noun": "Hartweizen",
    "article": "der"
  },
  {
    "noun": "Hasardeur",
    "article": "der"
  },
  {
    "noun": "Hascher",
    "article": "der"
  },
  {
    "noun": "Haschischraucher",
    "article": "der"
  },
  {
    "noun": "Hase",
    "article": "der"
  },
  {
    "noun": "Hasel",
    "article": "der"
  },
  {
    "noun": "Haselbusch",
    "article": "der"
  },
  {
    "noun": "Haselstrauch",
    "article": "der"
  },
  {
    "noun": "Hasenbraten",
    "article": "der"
  },
  {
    "noun": "Haspel",
    "article": "der"
  },
  {
    "noun": "Hass",
    "article": "der"
  },
  {
    "noun": "Hasser",
    "article": "der"
  },
  {
    "noun": "Hattrick",
    "article": "der"
  },
  {
    "noun": "Haubentaucher",
    "article": "der"
  },
  {
    "noun": "Hauch",
    "article": "der"
  },
  {
    "noun": "Hauchlaut",
    "article": "der"
  },
  {
    "noun": "Haudegen",
    "article": "der"
  },
  {
    "noun": "Hauer",
    "article": "der"
  },
  {
    "noun": "Haufen",
    "article": "der"
  },
  {
    "noun": "Hauklotz",
    "article": "der"
  },
  {
    "noun": "Hauptabnehmer",
    "article": "der"
  },
  {
    "noun": "Hauptaltar",
    "article": "der"
  },
  {
    "noun": "Hauptausgang",
    "article": "der"
  },
  {
    "noun": "Hauptbahnhof",
    "article": "der"
  },
  {
    "noun": "Hauptberuf",
    "article": "der"
  },
  {
    "noun": "Hauptbestandteil",
    "article": "der"
  },
  {
    "noun": "Hauptbuchhalter",
    "article": "der"
  },
  {
    "noun": "Hauptdarsteller",
    "article": "der"
  },
  {
    "noun": "Haupteingang",
    "article": "der"
  },
  {
    "noun": "Haupterbe",
    "article": "der"
  },
  {
    "noun": "Hauptfaktor",
    "article": "der"
  },
  {
    "noun": "Hauptfeldwebel",
    "article": "der"
  },
  {
    "noun": "Hauptfilm",
    "article": "der"
  },
  {
    "noun": "Hauptgang",
    "article": "der"
  },
  {
    "noun": "Hauptgedanke",
    "article": "der"
  },
  {
    "noun": "Hauptgewinn",
    "article": "der"
  },
  {
    "noun": "Hauptgewinner",
    "article": "der"
  },
  {
    "noun": "Hauptgrund",
    "article": "der"
  },
  {
    "noun": "Haupthahn",
    "article": "der"
  },
  {
    "noun": "Hauptinhalt",
    "article": "der"
  },
  {
    "noun": "Hauptkatalog",
    "article": "der"
  },
  {
    "noun": "Hauptlehrer",
    "article": "der"
  },
  {
    "noun": "Hauptmangel",
    "article": "der"
  },
  {
    "noun": "Hauptmann",
    "article": "der"
  },
  {
    "noun": "Hauptmieter",
    "article": "der"
  },
  {
    "noun": "Hauptnenner",
    "article": "der"
  },
  {
    "noun": "Hauptpreis",
    "article": "der"
  },
  {
    "noun": "Hauptpunkt",
    "article": "der"
  },
  {
    "noun": "Hauptquartier",
    "article": "der"
  },
  {
    "noun": "Hauptredner",
    "article": "der"
  },
  {
    "noun": "Hauptsatz",
    "article": "der"
  },
  {
    "noun": "Hauptschriftleiter",
    "article": "der"
  },
  {
    "noun": "Hauptschulabschluss",
    "article": "der"
  },
  {
    "noun": "Hauptschuldner",
    "article": "der"
  },
  {
    "noun": "Hauptsitz",
    "article": "der"
  },
  {
    "noun": "Hauptspeicher",
    "article": "der"
  },
  {
    "noun": "Hauptstrang",
    "article": "der"
  },
  {
    "noun": "Hauptteil",
    "article": "der"
  },
  {
    "noun": "Haupttitel",
    "article": "der"
  },
  {
    "noun": "Haupttreffer",
    "article": "der"
  },
  {
    "noun": "Hauptunterschied",
    "article": "der"
  },
  {
    "noun": "Hauptverdiener",
    "article": "der"
  },
  {
    "noun": "Hauptverkehr",
    "article": "der"
  },
  {
    "noun": "Hauptwert",
    "article": "der"
  },
  {
    "noun": "Hauptwohnsitz",
    "article": "der"
  },
  {
    "noun": "Hauptzug",
    "article": "der"
  },
  {
    "noun": "Hausaltar",
    "article": "der"
  },
  {
    "noun": "Hausanschluss",
    "article": "der"
  },
  {
    "noun": "Hausanzug",
    "article": "der"
  },
  {
    "noun": "Hausarrest",
    "article": "der"
  },
  {
    "noun": "Hausarzt",
    "article": "der"
  },
  {
    "noun": "Hausbau",
    "article": "der"
  },
  {
    "noun": "Hausbedarf",
    "article": "der"
  },
  {
    "noun": "Hausbesetzer",
    "article": "der"
  },
  {
    "noun": "Hausbesitzer",
    "article": "der"
  },
  {
    "noun": "Hausbesuch",
    "article": "der"
  },
  {
    "noun": "Hausbewohner",
    "article": "der"
  },
  {
    "noun": "Hausbock",
    "article": "der"
  },
  {
    "noun": "Hausbrand",
    "article": "der"
  },
  {
    "noun": "Hausdetektiv",
    "article": "der"
  },
  {
    "noun": "Hausdiener",
    "article": "der"
  },
  {
    "noun": "Hausdrachen",
    "article": "der"
  },
  {
    "noun": "Hauseingang",
    "article": "der"
  },
  {
    "noun": "Hausflur",
    "article": "der"
  },
  {
    "noun": "Hausfreund",
    "article": "der"
  },
  {
    "noun": "Hausfriede",
    "article": "der"
  },
  {
    "noun": "Hausfriedensbruch",
    "article": "der"
  },
  {
    "noun": "Hausgarten",
    "article": "der"
  },
  {
    "noun": "Hausgast",
    "article": "der"
  },
  {
    "noun": "Hausgenosse",
    "article": "der"
  },
  {
    "noun": "Hausgott",
    "article": "der"
  },
  {
    "noun": "Haushalt",
    "article": "der"
  },
  {
    "noun": "Haushaltsartikel",
    "article": "der"
  },
  {
    "noun": "Haushaltsausgleich",
    "article": "der"
  },
  {
    "noun": "Haushaltsausschuss",
    "article": "der"
  },
  {
    "noun": "Haushaltsfehlbetrag",
    "article": "der"
  },
  {
    "noun": "Haushaltsplan",
    "article": "der"
  },
  {
    "noun": "Haushaltsposten",
    "article": "der"
  },
  {
    "noun": "Haushaltstag",
    "article": "der"
  },
  {
    "noun": "Haushaltsvorstand",
    "article": "der"
  },
  {
    "noun": "Haushaltungsvorstand",
    "article": "der"
  },
  {
    "noun": "Hausherr",
    "article": "der"
  },
  {
    "noun": "Haushofmeister",
    "article": "der"
  },
  {
    "noun": "Hausierer",
    "article": "der"
  },
  {
    "noun": "Hausjurist",
    "article": "der"
  },
  {
    "noun": "Hausknecht",
    "article": "der"
  },
  {
    "noun": "Hauslehrer",
    "article": "der"
  },
  {
    "noun": "Hausmann",
    "article": "der"
  },
  {
    "noun": "Hausmannit",
    "article": "der"
  },
  {
    "noun": "Hausmeister",
    "article": "der"
  },
  {
    "noun": "Hausplatz",
    "article": "der"
  },
  {
    "noun": "Hausputz",
    "article": "der"
  },
  {
    "noun": "Hausrat",
    "article": "der"
  },
  {
    "noun": "Hausrock",
    "article": "der"
  },
  {
    "noun": "Hausschuh",
    "article": "der"
  },
  {
    "noun": "Hausschwamm",
    "article": "der"
  },
  {
    "noun": "Hausstand",
    "article": "der"
  },
  {
    "noun": "Haussuchungsbefehl",
    "article": "der"
  },
  {
    "noun": "Hausvater",
    "article": "der"
  },
  {
    "noun": "Hausverstand",
    "article": "der"
  },
  {
    "noun": "Hausverwalter",
    "article": "der"
  },
  {
    "noun": "Hauswart",
    "article": "der"
  },
  {
    "noun": "Hauswirt",
    "article": "der"
  },
  {
    "noun": "Hauszins",
    "article": "der"
  },
  {
    "noun": "Hautarzt",
    "article": "der"
  },
  {
    "noun": "Hautausschlag",
    "article": "der"
  },
  {
    "noun": "Hautfetzen",
    "article": "der"
  },
  {
    "noun": "Hautkratzer",
    "article": "der"
  },
  {
    "noun": "Hautkrebs",
    "article": "der"
  },
  {
    "noun": "Hautlappen",
    "article": "der"
  },
  {
    "noun": "Hautpilz",
    "article": "der"
  },
  {
    "noun": "Hauttyp",
    "article": "der"
  },
  {
    "noun": "Hautwolf",
    "article": "der"
  },
  {
    "noun": "Hawaiianer",
    "article": "der"
  },
  {
    "noun": "Hebel",
    "article": "der"
  },
  {
    "noun": "Hebelarm",
    "article": "der"
  },
  {
    "noun": "Hebelgriff",
    "article": "der"
  },
  {
    "noun": "Heber",
    "article": "der"
  },
  {
    "noun": "Hebesatz",
    "article": "der"
  },
  {
    "noun": "Hecht",
    "article": "der"
  },
  {
    "noun": "Hechtsprung",
    "article": "der"
  },
  {
    "noun": "Heckantrieb",
    "article": "der"
  },
  {
    "noun": "Heckmotor",
    "article": "der"
  },
  {
    "noun": "Heckspoiler",
    "article": "der"
  },
  {
    "noun": "Hederich",
    "article": "der"
  },
  {
    "noun": "Hedonismus",
    "article": "der"
  },
  {
    "noun": "Hedonist",
    "article": "der"
  },
  {
    "noun": "Heeresbericht",
    "article": "der"
  },
  {
    "noun": "Heeresdienst",
    "article": "der"
  },
  {
    "noun": "Heeresflieger",
    "article": "der"
  },
  {
    "noun": "Heereszug",
    "article": "der"
  },
  {
    "noun": "Hefekuchen",
    "article": "der"
  },
  {
    "noun": "Hefepilz",
    "article": "der"
  },
  {
    "noun": "Hefeteig",
    "article": "der"
  },
  {
    "noun": "Hefter",
    "article": "der"
  },
  {
    "noun": "Heftfaden",
    "article": "der"
  },
  {
    "noun": "Heftstich",
    "article": "der"
  },
  {
    "noun": "Heger",
    "article": "der"
  },
  {
    "noun": "Hehler",
    "article": "der"
  },
  {
    "noun": "Heide",
    "article": "der"
  },
  {
    "noun": "Heidenrespekt",
    "article": "der"
  },
  {
    "noun": "Heidenspektakel",
    "article": "der"
  },
  {
    "noun": "Heilbrunnen",
    "article": "der"
  },
  {
    "noun": "Heilbutt",
    "article": "der"
  },
  {
    "noun": "Heiler",
    "article": "der"
  },
  {
    "noun": "Heilerfolg",
    "article": "der"
  },
  {
    "noun": "Heiligabend",
    "article": "der"
  },
  {
    "noun": "Heiligenschein",
    "article": "der"
  },
  {
    "noun": "Heiligenschrein",
    "article": "der"
  },
  {
    "noun": "Heilpraktiker",
    "article": "der"
  },
  {
    "noun": "Heilungsprozess",
    "article": "der"
  },
  {
    "noun": "Heimarbeiter",
    "article": "der"
  },
  {
    "noun": "Heimatabend",
    "article": "der"
  },
  {
    "noun": "Heimatdichter",
    "article": "der"
  },
  {
    "noun": "Heimatfilm",
    "article": "der"
  },
  {
    "noun": "Heimathafen",
    "article": "der"
  },
  {
    "noun": "Heimatort",
    "article": "der"
  },
  {
    "noun": "Heimatstaat",
    "article": "der"
  },
  {
    "noun": "Heimcomputer",
    "article": "der"
  },
  {
    "noun": "Heimgang",
    "article": "der"
  },
  {
    "noun": "Heimkehrer",
    "article": "der"
  },
  {
    "noun": "Heimleiter",
    "article": "der"
  },
  {
    "noun": "Heimlichtuer",
    "article": "der"
  },
  {
    "noun": "Heimsieg",
    "article": "der"
  },
  {
    "noun": "Heimtrainer",
    "article": "der"
  },
  {
    "noun": "Heimweg",
    "article": "der"
  },
  {
    "noun": "Heimwerker",
    "article": "der"
  },
  {
    "noun": "Heini",
    "article": "der"
  },
  {
    "noun": "Heiratsantrag",
    "article": "der"
  },
  {
    "noun": "Heiratskandidat",
    "article": "der"
  },
  {
    "noun": "Heiratsschwindler",
    "article": "der"
  },
  {
    "noun": "Heiratsvermittler",
    "article": "der"
  },
  {
    "noun": "Heizapparat",
    "article": "der"
  },
  {
    "noun": "Heizer",
    "article": "der"
  },
  {
    "noun": "Heizkessel",
    "article": "der"
  },
  {
    "noun": "Heizleiter",
    "article": "der"
  },
  {
    "noun": "Heizofen",
    "article": "der"
  },
  {
    "noun": "Heizraum",
    "article": "der"
  },
  {
    "noun": "Heizstrahler",
    "article": "der"
  },
  {
    "noun": "Heizstrom",
    "article": "der"
  },
  {
    "noun": "Heizungskeller",
    "article": "der"
  },
  {
    "noun": "Heizungsmonteur",
    "article": "der"
  },
  {
    "noun": "Heizwert",
    "article": "der"
  },
  {
    "noun": "Hektarertrag",
    "article": "der"
  },
  {
    "noun": "Hektograf",
    "article": "der"
  },
  {
    "noun": "Held",
    "article": "der"
  },
  {
    "noun": "Heldenfriedhof",
    "article": "der"
  },
  {
    "noun": "Heldenmut",
    "article": "der"
  },
  {
    "noun": "Helfer",
    "article": "der"
  },
  {
    "noun": "Helfershelfer",
    "article": "der"
  },
  {
    "noun": "Helikopter",
    "article": "der"
  },
  {
    "noun": "Heliodor",
    "article": "der"
  },
  {
    "noun": "Heliograf",
    "article": "der"
  },
  {
    "noun": "Heliport",
    "article": "der"
  },
  {
    "noun": "Hellebardier",
    "article": "der"
  },
  {
    "noun": "Hellebardist",
    "article": "der"
  },
  {
    "noun": "Hellene",
    "article": "der"
  },
  {
    "noun": "Hellenismus",
    "article": "der"
  },
  {
    "noun": "Hellenist",
    "article": "der"
  },
  {
    "noun": "Helligkeitsregler",
    "article": "der"
  },
  {
    "noun": "Hellseher",
    "article": "der"
  },
  {
    "noun": "Helm",
    "article": "der"
  },
  {
    "noun": "Helophyt",
    "article": "der"
  },
  {
    "noun": "Helot",
    "article": "der"
  },
  {
    "noun": "Helvetier",
    "article": "der"
  },
  {
    "noun": "Helvetismus",
    "article": "der"
  },
  {
    "noun": "Hemdenstoff",
    "article": "der"
  },
  {
    "noun": "Hemdknopf",
    "article": "der"
  },
  {
    "noun": "Hemdkragen",
    "article": "der"
  },
  {
    "noun": "Hemdzipfel",
    "article": "der"
  },
  {
    "noun": "Hemmer",
    "article": "der"
  },
  {
    "noun": "Hemmklotz",
    "article": "der"
  },
  {
    "noun": "Hemmschuh",
    "article": "der"
  },
  {
    "noun": "Hemmstoff",
    "article": "der"
  },
  {
    "noun": "Hengst",
    "article": "der"
  },
  {
    "noun": "Henkel",
    "article": "der"
  },
  {
    "noun": "Henker",
    "article": "der"
  },
  {
    "noun": "Henkersknecht",
    "article": "der"
  },
  {
    "noun": "Heraldiker",
    "article": "der"
  },
  {
    "noun": "Herausforderer",
    "article": "der"
  },
  {
    "noun": "Herausgeber",
    "article": "der"
  },
  {
    "noun": "Herbst",
    "article": "der"
  },
  {
    "noun": "Herbstanfang",
    "article": "der"
  },
  {
    "noun": "Herbstling",
    "article": "der"
  },
  {
    "noun": "Herbstmonat",
    "article": "der"
  },
  {
    "noun": "Herbstmond",
    "article": "der"
  },
  {
    "noun": "Herbstregen",
    "article": "der"
  },
  {
    "noun": "Herbststurm",
    "article": "der"
  },
  {
    "noun": "Herd",
    "article": "der"
  },
  {
    "noun": "Herdentrieb",
    "article": "der"
  },
  {
    "noun": "Hergang",
    "article": "der"
  },
  {
    "noun": "Hering",
    "article": "der"
  },
  {
    "noun": "Heringsfang",
    "article": "der"
  },
  {
    "noun": "Heringssalat",
    "article": "der"
  },
  {
    "noun": "Herkules",
    "article": "der"
  },
  {
    "noun": "Herkunftsort",
    "article": "der"
  },
  {
    "noun": "Hermaphrodit",
    "article": "der"
  },
  {
    "noun": "Heroinismus",
    "article": "der"
  },
  {
    "noun": "Heroismus",
    "article": "der"
  },
  {
    "noun": "Herold",
    "article": "der"
  },
  {
    "noun": "Heronsball",
    "article": "der"
  },
  {
    "noun": "Heros",
    "article": "der"
  },
  {
    "noun": "Herpes",
    "article": "der"
  },
  {
    "noun": "Herr",
    "article": "der"
  },
  {
    "noun": "Herrenabend",
    "article": "der"
  },
  {
    "noun": "Herrenanzug",
    "article": "der"
  },
  {
    "noun": "Herrenausstatter",
    "article": "der"
  },
  {
    "noun": "Herrenbesuch",
    "article": "der"
  },
  {
    "noun": "Herrendienst",
    "article": "der"
  },
  {
    "noun": "Herrenpilz",
    "article": "der"
  },
  {
    "noun": "Herrensitz",
    "article": "der"
  },
  {
    "noun": "Herrgott",
    "article": "der"
  },
  {
    "noun": "Herrschaftsbereich",
    "article": "der"
  },
  {
    "noun": "Herrscher",
    "article": "der"
  },
  {
    "noun": "Hersteller",
    "article": "der"
  },
  {
    "noun": "Herstellungsfehler",
    "article": "der"
  },
  {
    "noun": "Herstellungspreis",
    "article": "der"
  },
  {
    "noun": "Herstellungsprozess",
    "article": "der"
  },
  {
    "noun": "Herumtreiber",
    "article": "der"
  },
  {
    "noun": "Herzanfall",
    "article": "der"
  },
  {
    "noun": "Herzautomatismus",
    "article": "der"
  },
  {
    "noun": "Herzbeutel",
    "article": "der"
  },
  {
    "noun": "Herzbube",
    "article": "der"
  },
  {
    "noun": "Herzchirurg",
    "article": "der"
  },
  {
    "noun": "Herzensbrecher",
    "article": "der"
  },
  {
    "noun": "Herzensbruder",
    "article": "der"
  },
  {
    "noun": "Herzenserguss",
    "article": "der"
  },
  {
    "noun": "Herzensgrund",
    "article": "der"
  },
  {
    "noun": "Herzenstrost",
    "article": "der"
  },
  {
    "noun": "Herzenswunsch",
    "article": "der"
  },
  {
    "noun": "Herzfehler",
    "article": "der"
  },
  {
    "noun": "Herzinfarkt",
    "article": "der"
  },
  {
    "noun": "Herzkatheter",
    "article": "der"
  },
  {
    "noun": "Herzklappenfehler",
    "article": "der"
  },
  {
    "noun": "Herzmuskel",
    "article": "der"
  },
  {
    "noun": "Herzog",
    "article": "der"
  },
  {
    "noun": "Herzpatient",
    "article": "der"
  },
  {
    "noun": "Herzrhythmus",
    "article": "der"
  },
  {
    "noun": "Herzschlag",
    "article": "der"
  },
  {
    "noun": "Herzschmerz",
    "article": "der"
  },
  {
    "noun": "Herzschrittmacher",
    "article": "der"
  },
  {
    "noun": "Herzspezialist",
    "article": "der"
  },
  {
    "noun": "Herzton",
    "article": "der"
  },
  {
    "noun": "Hesse",
    "article": "der"
  },
  {
    "noun": "Hetzer",
    "article": "der"
  },
  {
    "noun": "Hetzhund",
    "article": "der"
  },
  {
    "noun": "Heuboden",
    "article": "der"
  },
  {
    "noun": "Heuchler",
    "article": "der"
  },
  {
    "noun": "Heuer",
    "article": "der"
  },
  {
    "noun": "Heuervertrag",
    "article": "der"
  },
  {
    "noun": "Heuhaufen",
    "article": "der"
  },
  {
    "noun": "Heuler",
    "article": "der"
  },
  {
    "noun": "Heulton",
    "article": "der"
  },
  {
    "noun": "Heumonat",
    "article": "der"
  },
  {
    "noun": "Heuschnupfen",
    "article": "der"
  },
  {
    "noun": "Heuschober",
    "article": "der"
  },
  {
    "noun": "Heuwagen",
    "article": "der"
  },
  {
    "noun": "Heuwender",
    "article": "der"
  },
  {
    "noun": "Hexameter",
    "article": "der"
  },
  {
    "noun": "Hexapode",
    "article": "der"
  },
  {
    "noun": "Hexenmeister",
    "article": "der"
  },
  {
    "noun": "Hexenprozess",
    "article": "der"
  },
  {
    "noun": "Hexenring",
    "article": "der"
  },
  {
    "noun": "Hexensabbat",
    "article": "der"
  },
  {
    "noun": "Hexenschuss",
    "article": "der"
  },
  {
    "noun": "Hexentanz",
    "article": "der"
  },
  {
    "noun": "Hexentanzplatz",
    "article": "der"
  },
  {
    "noun": "Hexer",
    "article": "der"
  },
  {
    "noun": "Hiatus",
    "article": "der"
  },
  {
    "noun": "Hibiskus",
    "article": "der"
  },
  {
    "noun": "Hickhack",
    "article": "der"
  },
  {
    "noun": "Hieb",
    "article": "der"
  },
  {
    "noun": "Hieromant",
    "article": "der"
  },
  {
    "noun": "Hilferuf",
    "article": "der"
  },
  {
    "noun": "Hilfsarbeiter",
    "article": "der"
  },
  {
    "noun": "Hilfsassistent",
    "article": "der"
  },
  {
    "noun": "Hilfsdienst",
    "article": "der"
  },
  {
    "noun": "Hilfslehrer",
    "article": "der"
  },
  {
    "noun": "Hilfsmotor",
    "article": "der"
  },
  {
    "noun": "Hilfsstoff",
    "article": "der"
  },
  {
    "noun": "Himalaja",
    "article": "der"
  },
  {
    "noun": "Himbeergeist",
    "article": "der"
  },
  {
    "noun": "Himbeersaft",
    "article": "der"
  },
  {
    "noun": "Himbeerstrauch",
    "article": "der"
  },
  {
    "noun": "Himmel",
    "article": "der"
  },
  {
    "noun": "Himmelfahrtstag",
    "article": "der"
  },
  {
    "noun": "Himmelsbogen",
    "article": "der"
  },
  {
    "noun": "Himmelspol",
    "article": "der"
  },
  {
    "noun": "Himmelsschreiber",
    "article": "der"
  },
  {
    "noun": "Himmelsstrich",
    "article": "der"
  },
  {
    "noun": "Hinauswurf",
    "article": "der"
  },
  {
    "noun": "Hindernislauf",
    "article": "der"
  },
  {
    "noun": "Hinderungsgrund",
    "article": "der"
  },
  {
    "noun": "Hindu",
    "article": "der"
  },
  {
    "noun": "Hinduismus",
    "article": "der"
  },
  {
    "noun": "Hinflug",
    "article": "der"
  },
  {
    "noun": "Hingang",
    "article": "der"
  },
  {
    "noun": "Hinkelstein",
    "article": "der"
  },
  {
    "noun": "Hinschied",
    "article": "der"
  },
  {
    "noun": "Hinterachsantrieb",
    "article": "der"
  },
  {
    "noun": "Hinterausgang",
    "article": "der"
  },
  {
    "noun": "Hinterdarm",
    "article": "der"
  },
  {
    "noun": "Hintergaumenlaut",
    "article": "der"
  },
  {
    "noun": "Hintergedanke",
    "article": "der"
  },
  {
    "noun": "Hintergrund",
    "article": "der"
  },
  {
    "noun": "Hinterhalt",
    "article": "der"
  },
  {
    "noun": "Hinterhof",
    "article": "der"
  },
  {
    "noun": "Hinterkipper",
    "article": "der"
  },
  {
    "noun": "Hinterkopf",
    "article": "der"
  },
  {
    "noun": "Hinterlauf",
    "article": "der"
  },
  {
    "noun": "Hinterleger",
    "article": "der"
  },
  {
    "noun": "Hinterlegungsschein",
    "article": "der"
  },
  {
    "noun": "Hintermann",
    "article": "der"
  },
  {
    "noun": "Hinterradantrieb",
    "article": "der"
  },
  {
    "noun": "Hintersatz",
    "article": "der"
  },
  {
    "noun": "Hintersinn",
    "article": "der"
  },
  {
    "noun": "Hintertreppenroman",
    "article": "der"
  },
  {
    "noun": "Hinweg",
    "article": "der"
  },
  {
    "noun": "Hinweis",
    "article": "der"
  },
  {
    "noun": "Hippie",
    "article": "der"
  },
  {
    "noun": "Hippocampus",
    "article": "der"
  },
  {
    "noun": "Hippogryph",
    "article": "der"
  },
  {
    "noun": "Hippokampus",
    "article": "der"
  },
  {
    "noun": "Hippokratiker",
    "article": "der"
  },
  {
    "noun": "Hippopotamus",
    "article": "der"
  },
  {
    "noun": "Hirnschaden",
    "article": "der"
  },
  {
    "noun": "Hirnschlag",
    "article": "der"
  },
  {
    "noun": "Hirnstamm",
    "article": "der"
  },
  {
    "noun": "Hirntumor",
    "article": "der"
  },
  {
    "noun": "Hirsch",
    "article": "der"
  },
  {
    "noun": "Hirsebrei",
    "article": "der"
  },
  {
    "noun": "Hirt",
    "article": "der"
  },
  {
    "noun": "Hirtenbrief",
    "article": "der"
  },
  {
    "noun": "Hirtenknabe",
    "article": "der"
  },
  {
    "noun": "Hirtenstab",
    "article": "der"
  },
  {
    "noun": "Hisbollah",
    "article": "der"
  },
  {
    "noun": "Historiker",
    "article": "der"
  },
  {
    "noun": "Historiograf",
    "article": "der"
  },
  {
    "noun": "Historismus",
    "article": "der"
  },
  {
    "noun": "Hit",
    "article": "der"
  },
  {
    "noun": "Hitzegrad",
    "article": "der"
  },
  {
    "noun": "Hitzeschild",
    "article": "der"
  },
  {
    "noun": "Hitzkopf",
    "article": "der"
  },
  {
    "noun": "Hitzschlag",
    "article": "der"
  },
  {
    "noun": "Hobbock",
    "article": "der"
  },
  {
    "noun": "Hobbyraum",
    "article": "der"
  },
  {
    "noun": "Hobel",
    "article": "der"
  },
  {
    "noun": "Hobelspan",
    "article": "der"
  },
  {
    "noun": "Hochadel",
    "article": "der"
  },
  {
    "noun": "Hochaltar",
    "article": "der"
  },
  {
    "noun": "Hochbau",
    "article": "der"
  },
  {
    "noun": "Hochbauingenieur",
    "article": "der"
  },
  {
    "noun": "Hochbetrieb",
    "article": "der"
  },
  {
    "noun": "Hochdruck",
    "article": "der"
  },
  {
    "noun": "Hochfrequenzgleichrichter",
    "article": "der"
  },
  {
    "noun": "Hochfrequenzstrom",
    "article": "der"
  },
  {
    "noun": "Hochgenuss",
    "article": "der"
  },
  {
    "noun": "Hochgeschwindigkeitsverkehr",
    "article": "der"
  },
  {
    "noun": "Hochgeschwindigkeitszug",
    "article": "der"
  },
  {
    "noun": "Hochglanz",
    "article": "der"
  },
  {
    "noun": "Hochleistungssport",
    "article": "der"
  },
  {
    "noun": "Hochmut",
    "article": "der"
  },
  {
    "noun": "Hochnebel",
    "article": "der"
  },
  {
    "noun": "Hochofen",
    "article": "der"
  },
  {
    "noun": "Hochschrank",
    "article": "der"
  },
  {
    "noun": "Hochschulabschluss",
    "article": "der"
  },
  {
    "noun": "Hochschulabsolvent",
    "article": "der"
  },
  {
    "noun": "Hochschullehrer",
    "article": "der"
  },
  {
    "noun": "Hochseilartist",
    "article": "der"
  },
  {
    "noun": "Hochsicherheitstrakt",
    "article": "der"
  },
  {
    "noun": "Hochsinn",
    "article": "der"
  },
  {
    "noun": "Hochsitz",
    "article": "der"
  },
  {
    "noun": "Hochsommer",
    "article": "der"
  },
  {
    "noun": "Hochspannungsmast",
    "article": "der"
  },
  {
    "noun": "Hochsprung",
    "article": "der"
  },
  {
    "noun": "Hochstapler",
    "article": "der"
  },
  {
    "noun": "Hochstart",
    "article": "der"
  },
  {
    "noun": "Hochtemperaturreaktor",
    "article": "der"
  },
  {
    "noun": "Hochton",
    "article": "der"
  },
  {
    "noun": "Hochtourist",
    "article": "der"
  },
  {
    "noun": "Hochverrat",
    "article": "der"
  },
  {
    "noun": "Hochwald",
    "article": "der"
  },
  {
    "noun": "Hochzeitsflug",
    "article": "der"
  },
  {
    "noun": "Hochzeitsgast",
    "article": "der"
  },
  {
    "noun": "Hochzeitstag",
    "article": "der"
  },
  {
    "noun": "Hochzeitszug",
    "article": "der"
  },
  {
    "noun": "Hockeyball",
    "article": "der"
  },
  {
    "noun": "Hockeyspieler",
    "article": "der"
  },
  {
    "noun": "Hoden",
    "article": "der"
  },
  {
    "noun": "Hodenbruch",
    "article": "der"
  },
  {
    "noun": "Hodenhochstand",
    "article": "der"
  },
  {
    "noun": "Hodenkrebs",
    "article": "der"
  },
  {
    "noun": "Hodensack",
    "article": "der"
  },
  {
    "noun": "Hof",
    "article": "der"
  },
  {
    "noun": "Hofdichter",
    "article": "der"
  },
  {
    "noun": "Hofdienst",
    "article": "der"
  },
  {
    "noun": "Hoferbe",
    "article": "der"
  },
  {
    "noun": "Hoffnungsfunke",
    "article": "der"
  },
  {
    "noun": "Hoffnungslauf",
    "article": "der"
  },
  {
    "noun": "Hoffnungsschimmer",
    "article": "der"
  },
  {
    "noun": "Hofhund",
    "article": "der"
  },
  {
    "noun": "Hofmeister",
    "article": "der"
  },
  {
    "noun": "Hofnarr",
    "article": "der"
  },
  {
    "noun": "Hofpoet",
    "article": "der"
  },
  {
    "noun": "Hofrat",
    "article": "der"
  },
  {
    "noun": "Hofraum",
    "article": "der"
  },
  {
    "noun": "Hofschauspieler",
    "article": "der"
  },
  {
    "noun": "Hofstaat",
    "article": "der"
  },
  {
    "noun": "Hohepriester",
    "article": "der"
  },
  {
    "noun": "Hohlball",
    "article": "der"
  },
  {
    "noun": "Hohlblockstein",
    "article": "der"
  },
  {
    "noun": "Hohlkopf",
    "article": "der"
  },
  {
    "noun": "Hohlleiter",
    "article": "der"
  },
  {
    "noun": "Hohlraum",
    "article": "der"
  },
  {
    "noun": "Hohlsaum",
    "article": "der"
  },
  {
    "noun": "Hohlschliff",
    "article": "der"
  },
  {
    "noun": "Hohlspiegel",
    "article": "der"
  },
  {
    "noun": "Hohlweg",
    "article": "der"
  },
  {
    "noun": "Hohlziegel",
    "article": "der"
  },
  {
    "noun": "Hohn",
    "article": "der"
  },
  {
    "noun": "Hokuspokus",
    "article": "der"
  },
  {
    "noun": "Holismus",
    "article": "der"
  },
  {
    "noun": "Holk",
    "article": "der"
  },
  {
    "noun": "Holm",
    "article": "der"
  },
  {
    "noun": "Holmgang",
    "article": "der"
  },
  {
    "noun": "Holocaust",
    "article": "der"
  },
  {
    "noun": "Holotypus",
    "article": "der"
  },
  {
    "noun": "Holunder",
    "article": "der"
  },
  {
    "noun": "Holunderbusch",
    "article": "der"
  },
  {
    "noun": "Holunderstrauch",
    "article": "der"
  },
  {
    "noun": "Holundertee",
    "article": "der"
  },
  {
    "noun": "Holzabfuhrweg",
    "article": "der"
  },
  {
    "noun": "Holzapfel",
    "article": "der"
  },
  {
    "noun": "Holzarbeiter",
    "article": "der"
  },
  {
    "noun": "Holzbildhauer",
    "article": "der"
  },
  {
    "noun": "Holzblock",
    "article": "der"
  },
  {
    "noun": "Holzbock",
    "article": "der"
  },
  {
    "noun": "Holzboden",
    "article": "der"
  },
  {
    "noun": "Holzbohrer",
    "article": "der"
  },
  {
    "noun": "Holzdiebstahl",
    "article": "der"
  },
  {
    "noun": "Holzeinschlag",
    "article": "der"
  },
  {
    "noun": "Holzer",
    "article": "der"
  },
  {
    "noun": "Holzessig",
    "article": "der"
  },
  {
    "noun": "Holzhacker",
    "article": "der"
  },
  {
    "noun": "Holzhammer",
    "article": "der"
  },
  {
    "noun": "Holzhandel",
    "article": "der"
  },
  {
    "noun": "Holzhauer",
    "article": "der"
  },
  {
    "noun": "Holzhaufen",
    "article": "der"
  },
  {
    "noun": "Holzkasten",
    "article": "der"
  },
  {
    "noun": "Holzkeil",
    "article": "der"
  },
  {
    "noun": "Holzklotz",
    "article": "der"
  },
  {
    "noun": "Holzknecht",
    "article": "der"
  },
  {
    "noun": "Holzkopf",
    "article": "der"
  },
  {
    "noun": "Holzlack",
    "article": "der"
  },
  {
    "noun": "Holzleim",
    "article": "der"
  },
  {
    "noun": "Holzmarkt",
    "article": "der"
  },
  {
    "noun": "Holznagel",
    "article": "der"
  },
  {
    "noun": "Holzofen",
    "article": "der"
  },
  {
    "noun": "Holzopal",
    "article": "der"
  },
  {
    "noun": "Holzpantoffel",
    "article": "der"
  },
  {
    "noun": "Holzpflock",
    "article": "der"
  },
  {
    "noun": "Holzplatz",
    "article": "der"
  },
  {
    "noun": "Holzrahmen",
    "article": "der"
  },
  {
    "noun": "Holzschimmel",
    "article": "der"
  },
  {
    "noun": "Holzschneider",
    "article": "der"
  },
  {
    "noun": "Holzschnitt",
    "article": "der"
  },
  {
    "noun": "Holzschnitzer",
    "article": "der"
  },
  {
    "noun": "Holzschuh",
    "article": "der"
  },
  {
    "noun": "Holzschuppen",
    "article": "der"
  },
  {
    "noun": "Holzschutz",
    "article": "der"
  },
  {
    "noun": "Holzspan",
    "article": "der"
  },
  {
    "noun": "Holzsplitter",
    "article": "der"
  },
  {
    "noun": "Holzstoff",
    "article": "der"
  },
  {
    "noun": "Holzstuhl",
    "article": "der"
  },
  {
    "noun": "Holzteer",
    "article": "der"
  },
  {
    "noun": "Holztisch",
    "article": "der"
  },
  {
    "noun": "Holzverschlag",
    "article": "der"
  },
  {
    "noun": "Holzweg",
    "article": "der"
  },
  {
    "noun": "Holzwurm",
    "article": "der"
  },
  {
    "noun": "Hometrainer",
    "article": "der"
  },
  {
    "noun": "Hominid",
    "article": "der"
  },
  {
    "noun": "Homomorphismus",
    "article": "der"
  },
  {
    "noun": "Honduraner",
    "article": "der"
  },
  {
    "noun": "Honig",
    "article": "der"
  },
  {
    "noun": "Honigfresser",
    "article": "der"
  },
  {
    "noun": "Honigkuchen",
    "article": "der"
  },
  {
    "noun": "Honigmond",
    "article": "der"
  },
  {
    "noun": "Honigseim",
    "article": "der"
  },
  {
    "noun": "Honigtau",
    "article": "der"
  },
  {
    "noun": "Honigwein",
    "article": "der"
  },
  {
    "noun": "Honorant",
    "article": "der"
  },
  {
    "noun": "Honoratior",
    "article": "der"
  },
  {
    "noun": "Hooligan",
    "article": "der"
  },
  {
    "noun": "Hooliganismus",
    "article": "der"
  },
  {
    "noun": "Hopfenanbau",
    "article": "der"
  },
  {
    "noun": "Horcher",
    "article": "der"
  },
  {
    "noun": "Horchposten",
    "article": "der"
  },
  {
    "noun": "Horizont",
    "article": "der"
  },
  {
    "noun": "Horizontalkonzern",
    "article": "der"
  },
  {
    "noun": "Hormonhaushalt",
    "article": "der"
  },
  {
    "noun": "Hormonspiegel",
    "article": "der"
  },
  {
    "noun": "Hornfels",
    "article": "der"
  },
  {
    "noun": "Hornhai",
    "article": "der"
  },
  {
    "noun": "Hornhecht",
    "article": "der"
  },
  {
    "noun": "Hornissenstich",
    "article": "der"
  },
  {
    "noun": "Hornist",
    "article": "der"
  },
  {
    "noun": "Hornochs",
    "article": "der"
  },
  {
    "noun": "Hornung",
    "article": "der"
  },
  {
    "noun": "Horror",
    "article": "der"
  },
  {
    "noun": "Horrorfilm",
    "article": "der"
  },
  {
    "noun": "Horst",
    "article": "der"
  },
  {
    "noun": "Hort",
    "article": "der"
  },
  {
    "noun": "Hosenanzug",
    "article": "der"
  },
  {
    "noun": "Hosenaufschlag",
    "article": "der"
  },
  {
    "noun": "Hosenbandorden",
    "article": "der"
  },
  {
    "noun": "Hosenboden",
    "article": "der"
  },
  {
    "noun": "Hosenbund",
    "article": "der"
  },
  {
    "noun": "Hosenmatz",
    "article": "der"
  },
  {
    "noun": "Hosenschlitz",
    "article": "der"
  },
  {
    "noun": "Hosenstall",
    "article": "der"
  },
  {
    "noun": "Hospitalismus",
    "article": "der"
  },
  {
    "noun": "Hospitant",
    "article": "der"
  },
  {
    "noun": "Host",
    "article": "der"
  },
  {
    "noun": "Hostrechner",
    "article": "der"
  },
  {
    "noun": "Hotelbesitzer",
    "article": "der"
  },
  {
    "noun": "Hoteldirektor",
    "article": "der"
  },
  {
    "noun": "Hotelgast",
    "article": "der"
  },
  {
    "noun": "Hotelier",
    "article": "der"
  },
  {
    "noun": "Hotelmanager",
    "article": "der"
  },
  {
    "noun": "Hotelnachweis",
    "article": "der"
  },
  {
    "noun": "Hotspot",
    "article": "der"
  },
  {
    "noun": "Hottentotte",
    "article": "der"
  },
  {
    "noun": "Hub",
    "article": "der"
  },
  {
    "noun": "Hubraum",
    "article": "der"
  },
  {
    "noun": "Hubschrauber",
    "article": "der"
  },
  {
    "noun": "Hubstapler",
    "article": "der"
  },
  {
    "noun": "Huchen",
    "article": "der"
  },
  {
    "noun": "Huckepackverkehr",
    "article": "der"
  },
  {
    "noun": "Hudel",
    "article": "der"
  },
  {
    "noun": "Huf",
    "article": "der"
  },
  {
    "noun": "Hufbeschlag",
    "article": "der"
  },
  {
    "noun": "Hufeisenmagnet",
    "article": "der"
  },
  {
    "noun": "Huflattich",
    "article": "der"
  },
  {
    "noun": "Hufnagel",
    "article": "der"
  },
  {
    "noun": "Hufschlag",
    "article": "der"
  },
  {
    "noun": "Hufschmied",
    "article": "der"
  },
  {
    "noun": "Hugenotte",
    "article": "der"
  },
  {
    "noun": "Hughestelegraf",
    "article": "der"
  },
  {
    "noun": "Humanismus",
    "article": "der"
  },
  {
    "noun": "Humanist",
    "article": "der"
  },
  {
    "noun": "Humanmediziner",
    "article": "der"
  },
  {
    "noun": "Humboldtstrom",
    "article": "der"
  },
  {
    "noun": "Humbug",
    "article": "der"
  },
  {
    "noun": "Hummer",
    "article": "der"
  },
  {
    "noun": "Humor",
    "article": "der"
  },
  {
    "noun": "Humorist",
    "article": "der"
  },
  {
    "noun": "Humpen",
    "article": "der"
  },
  {
    "noun": "Humus",
    "article": "der"
  },
  {
    "noun": "Humusboden",
    "article": "der"
  },
  {
    "noun": "Hund",
    "article": "der"
  },
  {
    "noun": "Hundebandwurm",
    "article": "der"
  },
  {
    "noun": "Hundebesitzer",
    "article": "der"
  },
  {
    "noun": "Hundehalter",
    "article": "der"
  },
  {
    "noun": "Hundekot",
    "article": "der"
  },
  {
    "noun": "Hundekuchen",
    "article": "der"
  },
  {
    "noun": "Hundeliebhaber",
    "article": "der"
  },
  {
    "noun": "Hundenarr",
    "article": "der"
  },
  {
    "noun": "Hunderter",
    "article": "der"
  },
  {
    "noun": "Hundertsatz",
    "article": "der"
  },
  {
    "noun": "Hundesalon",
    "article": "der"
  },
  {
    "noun": "Hundeschlitten",
    "article": "der"
  },
  {
    "noun": "Hundesohn",
    "article": "der"
  },
  {
    "noun": "Hundezwinger",
    "article": "der"
  },
  {
    "noun": "Hundsfott",
    "article": "der"
  },
  {
    "noun": "Hundszahn",
    "article": "der"
  },
  {
    "noun": "Hunger",
    "article": "der"
  },
  {
    "noun": "Hungerleider",
    "article": "der"
  },
  {
    "noun": "Hungerlohn",
    "article": "der"
  },
  {
    "noun": "Hungerstreik",
    "article": "der"
  },
  {
    "noun": "Hunne",
    "article": "der"
  },
  {
    "noun": "Hunt",
    "article": "der"
  },
  {
    "noun": "Hunter",
    "article": "der"
  },
  {
    "noun": "Hurone",
    "article": "der"
  },
  {
    "noun": "Hurrapatriot",
    "article": "der"
  },
  {
    "noun": "Hurrapatriotismus",
    "article": "der"
  },
  {
    "noun": "Hurraruf",
    "article": "der"
  },
  {
    "noun": "Hurrikan",
    "article": "der"
  },
  {
    "noun": "Husar",
    "article": "der"
  },
  {
    "noun": "Hustenanfall",
    "article": "der"
  },
  {
    "noun": "Hustenreiz",
    "article": "der"
  },
  {
    "noun": "Hustensaft",
    "article": "der"
  },
  {
    "noun": "Hutmacher",
    "article": "der"
  },
  {
    "noun": "Hutzucker",
    "article": "der"
  },
  {
    "noun": "Hyaloklastit",
    "article": "der"
  },
  {
    "noun": "Hyalophan",
    "article": "der"
  },
  {
    "noun": "Hybride",
    "article": "der"
  },
  {
    "noun": "Hybridrechner",
    "article": "der"
  },
  {
    "noun": "Hydrant",
    "article": "der"
  },
  {
    "noun": "Hydrargillit",
    "article": "der"
  },
  {
    "noun": "Hydrophan",
    "article": "der"
  },
  {
    "noun": "Hydroplan",
    "article": "der"
  },
  {
    "noun": "Hygieniker",
    "article": "der"
  },
  {
    "noun": "Hygrograf",
    "article": "der"
  },
  {
    "noun": "Hygrostat",
    "article": "der"
  },
  {
    "noun": "Hylozoismus",
    "article": "der"
  },
  {
    "noun": "Hype",
    "article": "der"
  },
  {
    "noun": "Hyperlink",
    "article": "der"
  },
  {
    "noun": "Hyperschall",
    "article": "der"
  },
  {
    "noun": "Hypersthen",
    "article": "der"
  },
  {
    "noun": "Hypertext",
    "article": "der"
  },
  {
    "noun": "Hypnotiseur",
    "article": "der"
  },
  {
    "noun": "Hypnotismus",
    "article": "der"
  },
  {
    "noun": "Hypochonder",
    "article": "der"
  },
  {
    "noun": "Hypokrit",
    "article": "der"
  },
  {
    "noun": "Hypothekar",
    "article": "der"
  },
  {
    "noun": "Hypothekarkredit",
    "article": "der"
  },
  {
    "noun": "Hypothekenbrief",
    "article": "der"
  },
  {
    "noun": "Hypothekenpfandbrief",
    "article": "der"
  },
  {
    "noun": "Hypothekenschuldner",
    "article": "der"
  },
  {
    "noun": "Hypothekenzins",
    "article": "der"
  },
  {
    "noun": "Hysteriker",
    "article": "der"
  },
  {
    "noun": "Magnetiseur",
    "article": "der"
  },
  {
    "noun": "Magnetismus",
    "article": "der"
  },
  {
    "noun": "Magnetit",
    "article": "der"
  },
  {
    "noun": "Magnetkern",
    "article": "der"
  },
  {
    "noun": "Magnetkompass",
    "article": "der"
  },
  {
    "noun": "Magnetkopf",
    "article": "der"
  },
  {
    "noun": "Magnetpol",
    "article": "der"
  },
  {
    "noun": "Magnetstreifen",
    "article": "der"
  },
  {
    "noun": "Mahagonibaum",
    "article": "der"
  },
  {
    "noun": "Maharadscha",
    "article": "der"
  },
  {
    "noun": "Mahlsand",
    "article": "der"
  },
  {
    "noun": "Mahlstein",
    "article": "der"
  },
  {
    "noun": "Mahlzahn",
    "article": "der"
  },
  {
    "noun": "Mahnbescheid",
    "article": "der"
  },
  {
    "noun": "Mahnbrief",
    "article": "der"
  },
  {
    "noun": "Mahner",
    "article": "der"
  },
  {
    "noun": "Mahr",
    "article": "der"
  },
  {
    "noun": "Mai",
    "article": "der"
  },
  {
    "noun": "Maibaum",
    "article": "der"
  },
  {
    "noun": "Maifisch",
    "article": "der"
  },
  {
    "noun": "Main",
    "article": "der"
  },
  {
    "noun": "Mais",
    "article": "der"
  },
  {
    "noun": "Maisbrei",
    "article": "der"
  },
  {
    "noun": "Maiskolben",
    "article": "der"
  },
  {
    "noun": "Maitrieb",
    "article": "der"
  },
  {
    "noun": "Major",
    "article": "der"
  },
  {
    "noun": "Majoran",
    "article": "der"
  },
  {
    "noun": "Makak",
    "article": "der"
  },
  {
    "noun": "Makedonier",
    "article": "der"
  },
  {
    "noun": "Makel",
    "article": "der"
  },
  {
    "noun": "Makler",
    "article": "der"
  },
  {
    "noun": "Makrobefehl",
    "article": "der"
  },
  {
    "noun": "Makrokosmos",
    "article": "der"
  },
  {
    "noun": "Malachit",
    "article": "der"
  },
  {
    "noun": "Malaie",
    "article": "der"
  },
  {
    "noun": "Malawier",
    "article": "der"
  },
  {
    "noun": "Malaysier",
    "article": "der"
  },
  {
    "noun": "Maler",
    "article": "der"
  },
  {
    "noun": "Malier",
    "article": "der"
  },
  {
    "noun": "Malkasten",
    "article": "der"
  },
  {
    "noun": "Malm",
    "article": "der"
  },
  {
    "noun": "Malstift",
    "article": "der"
  },
  {
    "noun": "Maltese",
    "article": "der"
  },
  {
    "noun": "Malteserorden",
    "article": "der"
  },
  {
    "noun": "Malzextrakt",
    "article": "der"
  },
  {
    "noun": "Mammon",
    "article": "der"
  },
  {
    "noun": "Mammonismus",
    "article": "der"
  },
  {
    "noun": "Mammutbaum",
    "article": "der"
  },
  {
    "noun": "Mammutfilm",
    "article": "der"
  },
  {
    "noun": "Manager",
    "article": "der"
  },
  {
    "noun": "Mandant",
    "article": "der"
  },
  {
    "noun": "Mandarinenbaum",
    "article": "der"
  },
  {
    "noun": "Mandatar",
    "article": "der"
  },
  {
    "noun": "Mandatsverlust",
    "article": "der"
  },
  {
    "noun": "Mandelbaum",
    "article": "der"
  },
  {
    "noun": "Mandrill",
    "article": "der"
  },
  {
    "noun": "Manganit",
    "article": "der"
  },
  {
    "noun": "Manganspat",
    "article": "der"
  },
  {
    "noun": "Mangel",
    "article": "der"
  },
  {
    "noun": "Mangobaum",
    "article": "der"
  },
  {
    "noun": "Mangold",
    "article": "der"
  },
  {
    "noun": "Mangrovenbaum",
    "article": "der"
  },
  {
    "noun": "Manierismus",
    "article": "der"
  },
  {
    "noun": "Manierist",
    "article": "der"
  },
  {
    "noun": "Manilahanf",
    "article": "der"
  },
  {
    "noun": "Manipulant",
    "article": "der"
  },
  {
    "noun": "Manipulator",
    "article": "der"
  },
  {
    "noun": "Mann",
    "article": "der"
  },
  {
    "noun": "Mannesstamm",
    "article": "der"
  },
  {
    "noun": "Mannmonat",
    "article": "der"
  },
  {
    "noun": "Mannschaftsgeist",
    "article": "der"
  },
  {
    "noun": "Mannschaftskamerad",
    "article": "der"
  },
  {
    "noun": "Mannschaftskampf",
    "article": "der"
  },
  {
    "noun": "Mannschaftssport",
    "article": "der"
  },
  {
    "noun": "Mannschaftswagen",
    "article": "der"
  },
  {
    "noun": "Manostat",
    "article": "der"
  },
  {
    "noun": "Manschettenknopf",
    "article": "der"
  },
  {
    "noun": "Mantel",
    "article": "der"
  },
  {
    "noun": "Manteltarif",
    "article": "der"
  },
  {
    "noun": "Manteltarifvertrag",
    "article": "der"
  },
  {
    "noun": "Maoismus",
    "article": "der"
  },
  {
    "noun": "Maoist",
    "article": "der"
  },
  {
    "noun": "Marabu",
    "article": "der"
  },
  {
    "noun": "Maracujasaft",
    "article": "der"
  },
  {
    "noun": "Maraschino",
    "article": "der"
  },
  {
    "noun": "Marasmus",
    "article": "der"
  },
  {
    "noun": "Marathon",
    "article": "der"
  },
  {
    "noun": "Marathonlauf",
    "article": "der"
  },
  {
    "noun": "Marder",
    "article": "der"
  },
  {
    "noun": "Marderhund",
    "article": "der"
  },
  {
    "noun": "Marderpelz",
    "article": "der"
  },
  {
    "noun": "Marginalismus",
    "article": "der"
  },
  {
    "noun": "Marie",
    "article": "der"
  },
  {
    "noun": "Marineinfanterist",
    "article": "der"
  },
  {
    "noun": "Marineoffizier",
    "article": "der"
  },
  {
    "noun": "Marinesoldat",
    "article": "der"
  },
  {
    "noun": "Marinismus",
    "article": "der"
  },
  {
    "noun": "Markenartikel",
    "article": "der"
  },
  {
    "noun": "Markenname",
    "article": "der"
  },
  {
    "noun": "Markenschutz",
    "article": "der"
  },
  {
    "noun": "Marker",
    "article": "der"
  },
  {
    "noun": "Marketender",
    "article": "der"
  },
  {
    "noun": "Marketingspezialist",
    "article": "der"
  },
  {
    "noun": "Markierstift",
    "article": "der"
  },
  {
    "noun": "Markknochen",
    "article": "der"
  },
  {
    "noun": "Markstein",
    "article": "der"
  },
  {
    "noun": "Markt",
    "article": "der"
  },
  {
    "noun": "Marktanteil",
    "article": "der"
  },
  {
    "noun": "Marktbericht",
    "article": "der"
  },
  {
    "noun": "Marktflecken",
    "article": "der"
  },
  {
    "noun": "Marktforscher",
    "article": "der"
  },
  {
    "noun": "Marktplatz",
    "article": "der"
  },
  {
    "noun": "Marktpreis",
    "article": "der"
  },
  {
    "noun": "Marktschreier",
    "article": "der"
  },
  {
    "noun": "Marktstand",
    "article": "der"
  },
  {
    "noun": "Markttag",
    "article": "der"
  },
  {
    "noun": "Markttrend",
    "article": "der"
  },
  {
    "noun": "Marktwert",
    "article": "der"
  },
  {
    "noun": "Marlspieker",
    "article": "der"
  },
  {
    "noun": "Marmor",
    "article": "der"
  },
  {
    "noun": "Marmorblock",
    "article": "der"
  },
  {
    "noun": "Marmorbruch",
    "article": "der"
  },
  {
    "noun": "Marmorgips",
    "article": "der"
  },
  {
    "noun": "Marmorkuchen",
    "article": "der"
  },
  {
    "noun": "Marodeur",
    "article": "der"
  },
  {
    "noun": "Marokkaner",
    "article": "der"
  },
  {
    "noun": "Marquis",
    "article": "der"
  },
  {
    "noun": "Mars",
    "article": "der"
  },
  {
    "noun": "Marsbewohner",
    "article": "der"
  },
  {
    "noun": "Marsch",
    "article": "der"
  },
  {
    "noun": "Marschall",
    "article": "der"
  },
  {
    "noun": "Marschallstab",
    "article": "der"
  },
  {
    "noun": "Marschbefehl",
    "article": "der"
  },
  {
    "noun": "Marschkompass",
    "article": "der"
  },
  {
    "noun": "Marsmensch",
    "article": "der"
  },
  {
    "noun": "Marstall",
    "article": "der"
  },
  {
    "noun": "Marterpfahl",
    "article": "der"
  },
  {
    "noun": "Martinstag",
    "article": "der"
  },
  {
    "noun": "Martyrer",
    "article": "der"
  },
  {
    "noun": "Marxismus",
    "article": "der"
  },
  {
    "noun": "Marxist",
    "article": "der"
  },
  {
    "noun": "Maschendraht",
    "article": "der"
  },
  {
    "noun": "Maschendrahtzaun",
    "article": "der"
  },
  {
    "noun": "Maschinenantrieb",
    "article": "der"
  },
  {
    "noun": "Maschinenarbeiter",
    "article": "der"
  },
  {
    "noun": "Maschinenbau",
    "article": "der"
  },
  {
    "noun": "Maschinenbauer",
    "article": "der"
  },
  {
    "noun": "Maschinenbauingenieur",
    "article": "der"
  },
  {
    "noun": "Maschinenbaumechaniker",
    "article": "der"
  },
  {
    "noun": "Maschinencode",
    "article": "der"
  },
  {
    "noun": "Maschinendefekt",
    "article": "der"
  },
  {
    "noun": "Maschineneinsatz",
    "article": "der"
  },
  {
    "noun": "Maschinenmeister",
    "article": "der"
  },
  {
    "noun": "Maschinenpark",
    "article": "der"
  },
  {
    "noun": "Maschinenraum",
    "article": "der"
  },
  {
    "noun": "Maschinensatz",
    "article": "der"
  },
  {
    "noun": "Maschinenschaden",
    "article": "der"
  },
  {
    "noun": "Maschinenschlosser",
    "article": "der"
  },
  {
    "noun": "Maschinenschreiber",
    "article": "der"
  },
  {
    "noun": "Maschinentelegraf",
    "article": "der"
  },
  {
    "noun": "Maschinist",
    "article": "der"
  },
  {
    "noun": "Maskaron",
    "article": "der"
  },
  {
    "noun": "Maskenball",
    "article": "der"
  },
  {
    "noun": "Maskenbildner",
    "article": "der"
  },
  {
    "noun": "Maskenverleih",
    "article": "der"
  },
  {
    "noun": "Masochismus",
    "article": "der"
  },
  {
    "noun": "Masochist",
    "article": "der"
  },
  {
    "noun": "Massagesalon",
    "article": "der"
  },
  {
    "noun": "Massenandrang",
    "article": "der"
  },
  {
    "noun": "Massenartikel",
    "article": "der"
  },
  {
    "noun": "Massenbedarf",
    "article": "der"
  },
  {
    "noun": "Massendefekt",
    "article": "der"
  },
  {
    "noun": "Massenmittelpunkt",
    "article": "der"
  },
  {
    "noun": "Massenspeicher",
    "article": "der"
  },
  {
    "noun": "Massenspektrograf",
    "article": "der"
  },
  {
    "noun": "Massensport",
    "article": "der"
  },
  {
    "noun": "Massenstreik",
    "article": "der"
  },
  {
    "noun": "Massentourismus",
    "article": "der"
  },
  {
    "noun": "Massenverbrauch",
    "article": "der"
  },
  {
    "noun": "Masseur",
    "article": "der"
  },
  {
    "noun": "Massivbau",
    "article": "der"
  },
  {
    "noun": "Mast",
    "article": "der"
  },
  {
    "noun": "Mastdarm",
    "article": "der"
  },
  {
    "noun": "Mastdarmspiegel",
    "article": "der"
  },
  {
    "noun": "Master",
    "article": "der"
  },
  {
    "noun": "Mastik",
    "article": "der"
  },
  {
    "noun": "Mastix",
    "article": "der"
  },
  {
    "noun": "Mastkorb",
    "article": "der"
  },
  {
    "noun": "Mastochse",
    "article": "der"
  },
  {
    "noun": "Masut",
    "article": "der"
  },
  {
    "noun": "Matador",
    "article": "der"
  },
  {
    "noun": "Materialbedarf",
    "article": "der"
  },
  {
    "noun": "Materialfehler",
    "article": "der"
  },
  {
    "noun": "Materialfluss",
    "article": "der"
  },
  {
    "noun": "Materialismus",
    "article": "der"
  },
  {
    "noun": "Materialist",
    "article": "der"
  },
  {
    "noun": "Materialmangel",
    "article": "der"
  },
  {
    "noun": "Materialverbrauch",
    "article": "der"
  },
  {
    "noun": "Mathematiker",
    "article": "der"
  },
  {
    "noun": "Mathematiklehrer",
    "article": "der"
  },
  {
    "noun": "Mathematikunterricht",
    "article": "der"
  },
  {
    "noun": "Matjeshering",
    "article": "der"
  },
  {
    "noun": "Matrixdrucker",
    "article": "der"
  },
  {
    "noun": "Matrose",
    "article": "der"
  },
  {
    "noun": "Matrosenanzug",
    "article": "der"
  },
  {
    "noun": "Matrosenkragen",
    "article": "der"
  },
  {
    "noun": "Mattlack",
    "article": "der"
  },
  {
    "noun": "Maturand",
    "article": "der"
  },
  {
    "noun": "Maturant",
    "article": "der"
  },
  {
    "noun": "Mauerfall",
    "article": "der"
  },
  {
    "noun": "Mauerhaken",
    "article": "der"
  },
  {
    "noun": "Mauersegler",
    "article": "der"
  },
  {
    "noun": "Mauerspecht",
    "article": "der"
  },
  {
    "noun": "Mauerverband",
    "article": "der"
  },
  {
    "noun": "Mauervorsprung",
    "article": "der"
  },
  {
    "noun": "Mauerziegel",
    "article": "der"
  },
  {
    "noun": "Maulbeerbaum",
    "article": "der"
  },
  {
    "noun": "Maulbeerspinner",
    "article": "der"
  },
  {
    "noun": "Maulesel",
    "article": "der"
  },
  {
    "noun": "Maulheld",
    "article": "der"
  },
  {
    "noun": "Maulkorb",
    "article": "der"
  },
  {
    "noun": "Maulkorberlass",
    "article": "der"
  },
  {
    "noun": "Maulwurf",
    "article": "der"
  },
  {
    "noun": "Maurer",
    "article": "der"
  },
  {
    "noun": "Maurerhammer",
    "article": "der"
  },
  {
    "noun": "Maurermeister",
    "article": "der"
  },
  {
    "noun": "Maurerpolier",
    "article": "der"
  },
  {
    "noun": "Mauritier",
    "article": "der"
  },
  {
    "noun": "Mausklick",
    "article": "der"
  },
  {
    "noun": "Maximalpreis",
    "article": "der"
  },
  {
    "noun": "Maximalwert",
    "article": "der"
  },
  {
    "noun": "Mazedonier",
    "article": "der"
  },
  {
    "noun": "Mechaniker",
    "article": "der"
  },
  {
    "noun": "Mechanismus",
    "article": "der"
  },
  {
    "noun": "Meckerer",
    "article": "der"
  },
  {
    "noun": "Medaillengewinner",
    "article": "der"
  },
  {
    "noun": "Medailleur",
    "article": "der"
  },
  {
    "noun": "Median",
    "article": "der"
  },
  {
    "noun": "Medianwert",
    "article": "der"
  },
  {
    "noun": "Mediator",
    "article": "der"
  },
  {
    "noun": "Medienfachmann",
    "article": "der"
  },
  {
    "noun": "Medienkonzern",
    "article": "der"
  },
  {
    "noun": "Medienrummel",
    "article": "der"
  },
  {
    "noun": "Medienverbund",
    "article": "der"
  },
  {
    "noun": "Medikamentenschrank",
    "article": "der"
  },
  {
    "noun": "Medikus",
    "article": "der"
  },
  {
    "noun": "Medizinalrat",
    "article": "der"
  },
  {
    "noun": "Medizinball",
    "article": "der"
  },
  {
    "noun": "Mediziner",
    "article": "der"
  },
  {
    "noun": "Medizinmann",
    "article": "der"
  },
  {
    "noun": "Medizinschrank",
    "article": "der"
  },
  {
    "noun": "Medizinstudent",
    "article": "der"
  },
  {
    "noun": "Meeraal",
    "article": "der"
  },
  {
    "noun": "Meerbusen",
    "article": "der"
  },
  {
    "noun": "Meeresarm",
    "article": "der"
  },
  {
    "noun": "Meeresbiologe",
    "article": "der"
  },
  {
    "noun": "Meeresblick",
    "article": "der"
  },
  {
    "noun": "Meeresboden",
    "article": "der"
  },
  {
    "noun": "Meeresgrund",
    "article": "der"
  },
  {
    "noun": "Meeresspiegel",
    "article": "der"
  },
  {
    "noun": "Meeresstrand",
    "article": "der"
  },
  {
    "noun": "Meergott",
    "article": "der"
  },
  {
    "noun": "Meerrettich",
    "article": "der"
  },
  {
    "noun": "Meerschaum",
    "article": "der"
  },
  {
    "noun": "Mehlbrei",
    "article": "der"
  },
  {
    "noun": "Mehltau",
    "article": "der"
  },
  {
    "noun": "Mehlwurm",
    "article": "der"
  },
  {
    "noun": "Mehraufwand",
    "article": "der"
  },
  {
    "noun": "Mehrbedarf",
    "article": "der"
  },
  {
    "noun": "Mehrbetrag",
    "article": "der"
  },
  {
    "noun": "Mehrertrag",
    "article": "der"
  },
  {
    "noun": "Mehrfarbendruck",
    "article": "der"
  },
  {
    "noun": "Mehrheitsbeschaffer",
    "article": "der"
  },
  {
    "noun": "Mehrheitsbeschluss",
    "article": "der"
  },
  {
    "noun": "Mehrkampf",
    "article": "der"
  },
  {
    "noun": "Mehrphasenstrom",
    "article": "der"
  },
  {
    "noun": "Mehrpreis",
    "article": "der"
  },
  {
    "noun": "Mehrprogrammbetrieb",
    "article": "der"
  },
  {
    "noun": "Mehrverbrauch",
    "article": "der"
  },
  {
    "noun": "Mehrwert",
    "article": "der"
  },
  {
    "noun": "Meier",
    "article": "der"
  },
  {
    "noun": "Meierhof",
    "article": "der"
  },
  {
    "noun": "Meilenstein",
    "article": "der"
  },
  {
    "noun": "Meiler",
    "article": "der"
  },
  {
    "noun": "Meineid",
    "article": "der"
  },
  {
    "noun": "Meinungsaustausch",
    "article": "der"
  },
  {
    "noun": "Meinungsbildner",
    "article": "der"
  },
  {
    "noun": "Meinungsforscher",
    "article": "der"
  },
  {
    "noun": "Meinungsmacher",
    "article": "der"
  },
  {
    "noun": "Meinungsstreit",
    "article": "der"
  },
  {
    "noun": "Meister",
    "article": "der"
  },
  {
    "noun": "Meisterdieb",
    "article": "der"
  },
  {
    "noun": "Meistergesang",
    "article": "der"
  },
  {
    "noun": "Meistersang",
    "article": "der"
  },
  {
    "noun": "Meisterschaftler",
    "article": "der"
  },
  {
    "noun": "Meisterschuss",
    "article": "der"
  },
  {
    "noun": "Meistersinger",
    "article": "der"
  },
  {
    "noun": "Mekong",
    "article": "der"
  },
  {
    "noun": "Melancholiker",
    "article": "der"
  },
  {
    "noun": "Melanesier",
    "article": "der"
  },
  {
    "noun": "Melanit",
    "article": "der"
  },
  {
    "noun": "Meldebogen",
    "article": "der"
  },
  {
    "noun": "Meldefahrer",
    "article": "der"
  },
  {
    "noun": "Melder",
    "article": "der"
  },
  {
    "noun": "Meldereiter",
    "article": "der"
  },
  {
    "noun": "Meldezettel",
    "article": "der"
  },
  {
    "noun": "Melkeimer",
    "article": "der"
  },
  {
    "noun": "Melker",
    "article": "der"
  },
  {
    "noun": "Melkstand",
    "article": "der"
  },
  {
    "noun": "Meltau",
    "article": "der"
  },
  {
    "noun": "Mengenbegriff",
    "article": "der"
  },
  {
    "noun": "Mengenrabatt",
    "article": "der"
  },
  {
    "noun": "Menhir",
    "article": "der"
  },
  {
    "noun": "Meniskus",
    "article": "der"
  },
  {
    "noun": "Meniskusriss",
    "article": "der"
  },
  {
    "noun": "Mensch",
    "article": "der"
  },
  {
    "noun": "Menschenaffe",
    "article": "der"
  },
  {
    "noun": "Menschenauflauf",
    "article": "der"
  },
  {
    "noun": "Menschenfeind",
    "article": "der"
  },
  {
    "noun": "Menschenfresser",
    "article": "der"
  },
  {
    "noun": "Menschenfreund",
    "article": "der"
  },
  {
    "noun": "Menschengeist",
    "article": "der"
  },
  {
    "noun": "Menschenhandel",
    "article": "der"
  },
  {
    "noun": "Menschenhass",
    "article": "der"
  },
  {
    "noun": "Menschenkenner",
    "article": "der"
  },
  {
    "noun": "Menschenraub",
    "article": "der"
  },
  {
    "noun": "Menschenschlag",
    "article": "der"
  },
  {
    "noun": "Menschenverstand",
    "article": "der"
  },
  {
    "noun": "Menschenwitz",
    "article": "der"
  },
  {
    "noun": "Mentor",
    "article": "der"
  },
  {
    "noun": "Mergel",
    "article": "der"
  },
  {
    "noun": "Meridian",
    "article": "der"
  },
  {
    "noun": "Meridiankreis",
    "article": "der"
  },
  {
    "noun": "Merkantilismus",
    "article": "der"
  },
  {
    "noun": "Merkantilist",
    "article": "der"
  },
  {
    "noun": "Merker",
    "article": "der"
  },
  {
    "noun": "Merksatz",
    "article": "der"
  },
  {
    "noun": "Merkurialismus",
    "article": "der"
  },
  {
    "noun": "Merkvers",
    "article": "der"
  },
  {
    "noun": "Merkzettel",
    "article": "der"
  },
  {
    "noun": "Merlin",
    "article": "der"
  },
  {
    "noun": "Mesner",
    "article": "der"
  },
  {
    "noun": "Messbecher",
    "article": "der"
  },
  {
    "noun": "Messbereich",
    "article": "der"
  },
  {
    "noun": "Messdiener",
    "article": "der"
  },
  {
    "noun": "Messeausweis",
    "article": "der"
  },
  {
    "noun": "Messebesucher",
    "article": "der"
  },
  {
    "noun": "Messekatalog",
    "article": "der"
  },
  {
    "noun": "Messergriff",
    "article": "der"
  },
  {
    "noun": "Messerheld",
    "article": "der"
  },
  {
    "noun": "Messerschmied",
    "article": "der"
  },
  {
    "noun": "Messerstecher",
    "article": "der"
  },
  {
    "noun": "Messerstich",
    "article": "der"
  },
  {
    "noun": "Messerwerfer",
    "article": "der"
  },
  {
    "noun": "Messestand",
    "article": "der"
  },
  {
    "noun": "Messianismus",
    "article": "der"
  },
  {
    "noun": "Messias",
    "article": "der"
  },
  {
    "noun": "Messingbeschlag",
    "article": "der"
  },
  {
    "noun": "Messingdraht",
    "article": "der"
  },
  {
    "noun": "Messtisch",
    "article": "der"
  },
  {
    "noun": "Messwandler",
    "article": "der"
  },
  {
    "noun": "Messwert",
    "article": "der"
  },
  {
    "noun": "Messwiderstand",
    "article": "der"
  },
  {
    "noun": "Messzylinder",
    "article": "der"
  },
  {
    "noun": "Met",
    "article": "der"
  },
  {
    "noun": "Metabolit",
    "article": "der"
  },
  {
    "noun": "Metallarbeiter",
    "article": "der"
  },
  {
    "noun": "Metalldampf",
    "article": "der"
  },
  {
    "noun": "Metalldetektor",
    "article": "der"
  },
  {
    "noun": "Metallfaden",
    "article": "der"
  },
  {
    "noun": "Metallkleber",
    "article": "der"
  },
  {
    "noun": "Metallschutz",
    "article": "der"
  },
  {
    "noun": "Metallspan",
    "article": "der"
  },
  {
    "noun": "Metallstab",
    "article": "der"
  },
  {
    "noun": "Metallstift",
    "article": "der"
  },
  {
    "noun": "Metallurg",
    "article": "der"
  },
  {
    "noun": "Metallurge",
    "article": "der"
  },
  {
    "noun": "Metamorphismus",
    "article": "der"
  },
  {
    "noun": "Metaphysiker",
    "article": "der"
  },
  {
    "noun": "Metaplasmus",
    "article": "der"
  },
  {
    "noun": "Meteorismus",
    "article": "der"
  },
  {
    "noun": "Meteorit",
    "article": "der"
  },
  {
    "noun": "Meteorologe",
    "article": "der"
  },
  {
    "noun": "Meteorstein",
    "article": "der"
  },
  {
    "noun": "Meterstab",
    "article": "der"
  },
  {
    "noun": "Methodiker",
    "article": "der"
  },
  {
    "noun": "Methodismus",
    "article": "der"
  },
  {
    "noun": "Methodist",
    "article": "der"
  },
  {
    "noun": "Methusalem",
    "article": "der"
  },
  {
    "noun": "Methylalkohol",
    "article": "der"
  },
  {
    "noun": "Metropolit",
    "article": "der"
  },
  {
    "noun": "Metzger",
    "article": "der"
  },
  {
    "noun": "Meuterer",
    "article": "der"
  },
  {
    "noun": "Mexikaner",
    "article": "der"
  },
  {
    "noun": "Mezzosopran",
    "article": "der"
  },
  {
    "noun": "Mief",
    "article": "der"
  },
  {
    "noun": "Miesepeter",
    "article": "der"
  },
  {
    "noun": "Miesling",
    "article": "der"
  },
  {
    "noun": "Miesmacher",
    "article": "der"
  },
  {
    "noun": "Mieter",
    "article": "der"
  },
  {
    "noun": "Mieterschutz",
    "article": "der"
  },
  {
    "noun": "Mietertrag",
    "article": "der"
  },
  {
    "noun": "Mietkauf",
    "article": "der"
  },
  {
    "noun": "Mietling",
    "article": "der"
  },
  {
    "noun": "Mietpreis",
    "article": "der"
  },
  {
    "noun": "Mietverlust",
    "article": "der"
  },
  {
    "noun": "Mietvertrag",
    "article": "der"
  },
  {
    "noun": "Mietwagen",
    "article": "der"
  },
  {
    "noun": "Mietzins",
    "article": "der"
  },
  {
    "noun": "Mietzuschuss",
    "article": "der"
  },
  {
    "noun": "Migrant",
    "article": "der"
  },
  {
    "noun": "Mikrobiologe",
    "article": "der"
  },
  {
    "noun": "Mikrochip",
    "article": "der"
  },
  {
    "noun": "Mikrocomputer",
    "article": "der"
  },
  {
    "noun": "Mikrofilm",
    "article": "der"
  },
  {
    "noun": "Mikrofongalgen",
    "article": "der"
  },
  {
    "noun": "Mikrokosmos",
    "article": "der"
  },
  {
    "noun": "Mikrolith",
    "article": "der"
  },
  {
    "noun": "Mikronesier",
    "article": "der"
  },
  {
    "noun": "Mikroorganismus",
    "article": "der"
  },
  {
    "noun": "Mikroprozessor",
    "article": "der"
  },
  {
    "noun": "Mikrowellenherd",
    "article": "der"
  },
  {
    "noun": "Mikrozensus",
    "article": "der"
  },
  {
    "noun": "Milchbrei",
    "article": "der"
  },
  {
    "noun": "Milchkaffee",
    "article": "der"
  },
  {
    "noun": "Milchmann",
    "article": "der"
  },
  {
    "noun": "Milchner",
    "article": "der"
  },
  {
    "noun": "Milchreis",
    "article": "der"
  },
  {
    "noun": "Milchsaft",
    "article": "der"
  },
  {
    "noun": "Milchtopf",
    "article": "der"
  },
  {
    "noun": "Milchzahn",
    "article": "der"
  },
  {
    "noun": "Milchzucker",
    "article": "der"
  },
  {
    "noun": "Milderungsgrund",
    "article": "der"
  },
  {
    "noun": "Militarismus",
    "article": "der"
  },
  {
    "noun": "Militarist",
    "article": "der"
  },
  {
    "noun": "Millionenauftrag",
    "article": "der"
  },
  {
    "noun": "Millionenschaden",
    "article": "der"
  },
  {
    "noun": "Milzbrand",
    "article": "der"
  },
  {
    "noun": "Milzriss",
    "article": "der"
  },
  {
    "noun": "Mime",
    "article": "der"
  },
  {
    "noun": "Minderbedarf",
    "article": "der"
  },
  {
    "noun": "Minderbetrag",
    "article": "der"
  },
  {
    "noun": "Minderertrag",
    "article": "der"
  },
  {
    "noun": "Minderheitenschutz",
    "article": "der"
  },
  {
    "noun": "Minderwertigkeitskomplex",
    "article": "der"
  },
  {
    "noun": "Mindestabstand",
    "article": "der"
  },
  {
    "noun": "Mindestbeitrag",
    "article": "der"
  },
  {
    "noun": "Mindestbetrag",
    "article": "der"
  },
  {
    "noun": "Mindestlohn",
    "article": "der"
  },
  {
    "noun": "Mindestpreis",
    "article": "der"
  },
  {
    "noun": "Mindestumtausch",
    "article": "der"
  },
  {
    "noun": "Minenleger",
    "article": "der"
  },
  {
    "noun": "Minensucher",
    "article": "der"
  },
  {
    "noun": "Minenwerfer",
    "article": "der"
  },
  {
    "noun": "Mineraloge",
    "article": "der"
  },
  {
    "noun": "Mineralstoff",
    "article": "der"
  },
  {
    "noun": "Minimalbetrag",
    "article": "der"
  },
  {
    "noun": "Minimalist",
    "article": "der"
  },
  {
    "noun": "Minimalwert",
    "article": "der"
  },
  {
    "noun": "Minirock",
    "article": "der"
  },
  {
    "noun": "Minispion",
    "article": "der"
  },
  {
    "noun": "Minister",
    "article": "der"
  },
  {
    "noun": "Ministerialrat",
    "article": "der"
  },
  {
    "noun": "Ministerrat",
    "article": "der"
  },
  {
    "noun": "Ministrant",
    "article": "der"
  },
  {
    "noun": "Minnesang",
    "article": "der"
  },
  {
    "noun": "Minnesinger",
    "article": "der"
  },
  {
    "noun": "Minor",
    "article": "der"
  },
  {
    "noun": "Minotaur",
    "article": "der"
  },
  {
    "noun": "Minuend",
    "article": "der"
  },
  {
    "noun": "Minusbetrag",
    "article": "der"
  },
  {
    "noun": "Minuspol",
    "article": "der"
  },
  {
    "noun": "Minuspunkt",
    "article": "der"
  },
  {
    "noun": "Minutenzeiger",
    "article": "der"
  },
  {
    "noun": "Mirabellenbaum",
    "article": "der"
  },
  {
    "noun": "Misanthrop",
    "article": "der"
  },
  {
    "noun": "Mischer",
    "article": "der"
  },
  {
    "noun": "Mischkonzern",
    "article": "der"
  },
  {
    "noun": "Mischkristall",
    "article": "der"
  },
  {
    "noun": "Mischling",
    "article": "der"
  },
  {
    "noun": "Mischmasch",
    "article": "der"
  },
  {
    "noun": "Mischwald",
    "article": "der"
  },
  {
    "noun": "Misogam",
    "article": "der"
  },
  {
    "noun": "Misogyn",
    "article": "der"
  },
  {
    "noun": "Misserfolg",
    "article": "der"
  },
  {
    "noun": "Missgriff",
    "article": "der"
  },
  {
    "noun": "Missionar",
    "article": "der"
  },
  {
    "noun": "Missionschef",
    "article": "der"
  },
  {
    "noun": "Missklang",
    "article": "der"
  },
  {
    "noun": "Missmut",
    "article": "der"
  },
  {
    "noun": "Misston",
    "article": "der"
  },
  {
    "noun": "Misstrauensantrag",
    "article": "der"
  },
  {
    "noun": "Misswuchs",
    "article": "der"
  },
  {
    "noun": "Mist",
    "article": "der"
  },
  {
    "noun": "Mistelzweig",
    "article": "der"
  },
  {
    "noun": "Mister",
    "article": "der"
  },
  {
    "noun": "Mistfink",
    "article": "der"
  },
  {
    "noun": "Misthaufen",
    "article": "der"
  },
  {
    "noun": "Mistkerl",
    "article": "der"
  },
  {
    "noun": "Mistral",
    "article": "der"
  },
  {
    "noun": "Mitarbeiter",
    "article": "der"
  },
  {
    "noun": "Mitarbeiterstab",
    "article": "der"
  },
  {
    "noun": "Mitautor",
    "article": "der"
  },
  {
    "noun": "Mitbesitz",
    "article": "der"
  },
  {
    "noun": "Mitbesitzer",
    "article": "der"
  },
  {
    "noun": "Mitbewerber",
    "article": "der"
  },
  {
    "noun": "Mitbewohner",
    "article": "der"
  },
  {
    "noun": "Miterbe",
    "article": "der"
  },
  {
    "noun": "Mitesser",
    "article": "der"
  },
  {
    "noun": "Mitfahrer",
    "article": "der"
  },
  {
    "noun": "Mitgliedsausweis",
    "article": "der"
  },
  {
    "noun": "Mitgliedsbeitrag",
    "article": "der"
  },
  {
    "noun": "Mitgliedsstaat",
    "article": "der"
  },
  {
    "noun": "Mitgliedstaat",
    "article": "der"
  },
  {
    "noun": "Mitherausgeber",
    "article": "der"
  },
  {
    "noun": "Mitinhaber",
    "article": "der"
  },
  {
    "noun": "Mitlaut",
    "article": "der"
  },
  {
    "noun": "Mitmensch",
    "article": "der"
  },
  {
    "noun": "Mitnahmeeffekt",
    "article": "der"
  },
  {
    "noun": "Mitnahmepreis",
    "article": "der"
  },
  {
    "noun": "Mitschnitt",
    "article": "der"
  },
  {
    "noun": "Mitspieler",
    "article": "der"
  },
  {
    "noun": "Mitstreiter",
    "article": "der"
  },
  {
    "noun": "Mittag",
    "article": "der"
  },
  {
    "noun": "Mittagskreis",
    "article": "der"
  },
  {
    "noun": "Mittagsschlaf",
    "article": "der"
  },
  {
    "noun": "Mittagstisch",
    "article": "der"
  },
  {
    "noun": "Mittelbau",
    "article": "der"
  },
  {
    "noun": "Mittelbetrieb",
    "article": "der"
  },
  {
    "noun": "Mittelfeldspieler",
    "article": "der"
  },
  {
    "noun": "Mittelfinger",
    "article": "der"
  },
  {
    "noun": "Mittelgang",
    "article": "der"
  },
  {
    "noun": "Mittelgewichtler",
    "article": "der"
  },
  {
    "noun": "Mittelklassewagen",
    "article": "der"
  },
  {
    "noun": "Mittelkreis",
    "article": "der"
  },
  {
    "noun": "Mittelmeerraum",
    "article": "der"
  },
  {
    "noun": "Mittelmotor",
    "article": "der"
  },
  {
    "noun": "Mittelpunkt",
    "article": "der"
  },
  {
    "noun": "Mittelscheitel",
    "article": "der"
  },
  {
    "noun": "Mittelsmann",
    "article": "der"
  },
  {
    "noun": "Mittelstand",
    "article": "der"
  },
  {
    "noun": "Mittelstreckenlauf",
    "article": "der"
  },
  {
    "noun": "Mittelstreifen",
    "article": "der"
  },
  {
    "noun": "Mittelweg",
    "article": "der"
  },
  {
    "noun": "Mittelwert",
    "article": "der"
  },
  {
    "noun": "Mittler",
    "article": "der"
  },
  {
    "noun": "Mittsommer",
    "article": "der"
  },
  {
    "noun": "Mittwinter",
    "article": "der"
  },
  {
    "noun": "Mittwoch",
    "article": "der"
  },
  {
    "noun": "Mittwochabend",
    "article": "der"
  },
  {
    "noun": "Mitunterzeichner",
    "article": "der"
  },
  {
    "noun": "Mitverfasser",
    "article": "der"
  },
  {
    "noun": "Mitwisser",
    "article": "der"
  },
  {
    "noun": "Mix",
    "article": "der"
  },
  {
    "noun": "Mixbecher",
    "article": "der"
  },
  {
    "noun": "Mixer",
    "article": "der"
  },
  {
    "noun": "Mob",
    "article": "der"
  },
  {
    "noun": "Mobilfunk",
    "article": "der"
  },
  {
    "noun": "Mod",
    "article": "der"
  },
  {
    "noun": "Modalsatz",
    "article": "der"
  },
  {
    "noun": "Modder",
    "article": "der"
  },
  {
    "noun": "Modeartikel",
    "article": "der"
  },
  {
    "noun": "Modedesigner",
    "article": "der"
  },
  {
    "noun": "Modefan",
    "article": "der"
  },
  {
    "noun": "Modefotograf",
    "article": "der"
  },
  {
    "noun": "Modellbauer",
    "article": "der"
  },
  {
    "noun": "Modelleur",
    "article": "der"
  },
  {
    "noun": "Modellfall",
    "article": "der"
  },
  {
    "noun": "Modellflug",
    "article": "der"
  },
  {
    "noun": "Modellierer",
    "article": "der"
  },
  {
    "noun": "Modellierton",
    "article": "der"
  },
  {
    "noun": "Modellversuch",
    "article": "der"
  },
  {
    "noun": "Modemacher",
    "article": "der"
  },
  {
    "noun": "Modenarr",
    "article": "der"
  },
  {
    "noun": "Moder",
    "article": "der"
  },
  {
    "noun": "Moderator",
    "article": "der"
  },
  {
    "noun": "Modergeruch",
    "article": "der"
  },
  {
    "noun": "Modernismus",
    "article": "der"
  },
  {
    "noun": "Modeschmuck",
    "article": "der"
  },
  {
    "noun": "Modezar",
    "article": "der"
  },
  {
    "noun": "Modezeichner",
    "article": "der"
  },
  {
    "noun": "Modulator",
    "article": "der"
  },
  {
    "noun": "Modus",
    "article": "der"
  },
  {
    "noun": "Mogler",
    "article": "der"
  },
  {
    "noun": "Mogul",
    "article": "der"
  },
  {
    "noun": "Mohair",
    "article": "der"
  },
  {
    "noun": "Mohn",
    "article": "der"
  },
  {
    "noun": "Mohnkuchen",
    "article": "der"
  },
  {
    "noun": "Mohnsaft",
    "article": "der"
  },
  {
    "noun": "Mohnsame",
    "article": "der"
  },
  {
    "noun": "Mohr",
    "article": "der"
  },
  {
    "noun": "Mohrenkopf",
    "article": "der"
  },
  {
    "noun": "Mokassin",
    "article": "der"
  },
  {
    "noun": "Mokka",
    "article": "der"
  },
  {
    "noun": "Molch",
    "article": "der"
  },
  {
    "noun": "Moldawier",
    "article": "der"
  },
  {
    "noun": "Mollakkord",
    "article": "der"
  },
  {
    "noun": "Molldreiklang",
    "article": "der"
  },
  {
    "noun": "Moloch",
    "article": "der"
  },
  {
    "noun": "Monarch",
    "article": "der"
  },
  {
    "noun": "Monarchismus",
    "article": "der"
  },
  {
    "noun": "Monarchist",
    "article": "der"
  },
  {
    "noun": "Monat",
    "article": "der"
  },
  {
    "noun": "Monatsanfang",
    "article": "der"
  },
  {
    "noun": "Monatsbeginn",
    "article": "der"
  },
  {
    "noun": "Monatsbeitrag",
    "article": "der"
  },
  {
    "noun": "Monatslohn",
    "article": "der"
  },
  {
    "noun": "Monatsname",
    "article": "der"
  },
  {
    "noun": "Mond",
    "article": "der"
  },
  {
    "noun": "Mondaufgang",
    "article": "der"
  },
  {
    "noun": "Mondkrater",
    "article": "der"
  },
  {
    "noun": "Mondschein",
    "article": "der"
  },
  {
    "noun": "Mondscheintarif",
    "article": "der"
  },
  {
    "noun": "Mondstein",
    "article": "der"
  },
  {
    "noun": "Monduntergang",
    "article": "der"
  },
  {
    "noun": "Mondwechsel",
    "article": "der"
  },
  {
    "noun": "Monegasse",
    "article": "der"
  },
  {
    "noun": "Monetarismus",
    "article": "der"
  },
  {
    "noun": "Mongole",
    "article": "der"
  },
  {
    "noun": "Mongolenfleck",
    "article": "der"
  },
  {
    "noun": "Mongolismus",
    "article": "der"
  },
  {
    "noun": "Monismus",
    "article": "der"
  },
  {
    "noun": "Monitor",
    "article": "der"
  },
  {
    "noun": "Monochromator",
    "article": "der"
  },
  {
    "noun": "Monogenismus",
    "article": "der"
  },
  {
    "noun": "Monolog",
    "article": "der"
  },
  {
    "noun": "Monomane",
    "article": "der"
  },
  {
    "noun": "Monoplan",
    "article": "der"
  },
  {
    "noun": "Monopolismus",
    "article": "der"
  },
  {
    "noun": "Monopolist",
    "article": "der"
  },
  {
    "noun": "Monopolkapitalismus",
    "article": "der"
  },
  {
    "noun": "Monopolpreis",
    "article": "der"
  },
  {
    "noun": "Monotheismus",
    "article": "der"
  },
  {
    "noun": "Monotheist",
    "article": "der"
  },
  {
    "noun": "Monsterbau",
    "article": "der"
  },
  {
    "noun": "Monsterfilm",
    "article": "der"
  },
  {
    "noun": "Monsun",
    "article": "der"
  },
  {
    "noun": "Montag",
    "article": "der"
  },
  {
    "noun": "Montagebau",
    "article": "der"
  },
  {
    "noun": "Montenegriner",
    "article": "der"
  },
  {
    "noun": "Monteur",
    "article": "der"
  },
  {
    "noun": "Montierer",
    "article": "der"
  },
  {
    "noun": "Mooskrepp",
    "article": "der"
  },
  {
    "noun": "Mop",
    "article": "der"
  },
  {
    "noun": "Mops",
    "article": "der"
  },
  {
    "noun": "Moralapostel",
    "article": "der"
  },
  {
    "noun": "Moralismus",
    "article": "der"
  },
  {
    "noun": "Moralist",
    "article": "der"
  },
  {
    "noun": "Moralkodex",
    "article": "der"
  },
  {
    "noun": "Moralphilosoph",
    "article": "der"
  },
  {
    "noun": "Morast",
    "article": "der"
  },
  {
    "noun": "Morgenmantel",
    "article": "der"
  },
  {
    "noun": "Morgenmuffel",
    "article": "der"
  },
  {
    "noun": "Morgennebel",
    "article": "der"
  },
  {
    "noun": "Morgenrock",
    "article": "der"
  },
  {
    "noun": "Morgenspaziergang",
    "article": "der"
  },
  {
    "noun": "Morgenstern",
    "article": "der"
  },
  {
    "noun": "Morgentau",
    "article": "der"
  },
  {
    "noun": "Mormone",
    "article": "der"
  },
  {
    "noun": "Morphinismus",
    "article": "der"
  },
  {
    "noun": "Morphinist",
    "article": "der"
  },
  {
    "noun": "Morseapparat",
    "article": "der"
  },
  {
    "noun": "Mosambikaner",
    "article": "der"
  },
  {
    "noun": "Moschus",
    "article": "der"
  },
  {
    "noun": "Moschusgeruch",
    "article": "der"
  },
  {
    "noun": "Moschusochse",
    "article": "der"
  },
  {
    "noun": "Mosel",
    "article": "der"
  },
  {
    "noun": "Moselwein",
    "article": "der"
  },
  {
    "noun": "Moskito",
    "article": "der"
  },
  {
    "noun": "Moskowiter",
    "article": "der"
  },
  {
    "noun": "Most",
    "article": "der"
  },
  {
    "noun": "Mostrich",
    "article": "der"
  },
  {
    "noun": "Motivationsschub",
    "article": "der"
  },
  {
    "noun": "Motivwagen",
    "article": "der"
  },
  {
    "noun": "Motor",
    "article": "der"
  },
  {
    "noun": "Motorblock",
    "article": "der"
  },
  {
    "noun": "Motordefekt",
    "article": "der"
  },
  {
    "noun": "Motorenbau",
    "article": "der"
  },
  {
    "noun": "Motorgenerator",
    "article": "der"
  },
  {
    "noun": "Motorpflug",
    "article": "der"
  },
  {
    "noun": "Motorradfahrer",
    "article": "der"
  },
  {
    "noun": "Motorradhelm",
    "article": "der"
  },
  {
    "noun": "Motorradsport",
    "article": "der"
  },
  {
    "noun": "Motorraum",
    "article": "der"
  },
  {
    "noun": "Motorroller",
    "article": "der"
  },
  {
    "noun": "Motorschaden",
    "article": "der"
  },
  {
    "noun": "Motorschlepper",
    "article": "der"
  },
  {
    "noun": "Motorschlitten",
    "article": "der"
  },
  {
    "noun": "Motorsegler",
    "article": "der"
  },
  {
    "noun": "Motorsport",
    "article": "der"
  },
  {
    "noun": "Motorwagen",
    "article": "der"
  },
  {
    "noun": "Mozzarella",
    "article": "der"
  },
  {
    "noun": "Muckefuck",
    "article": "der"
  },
  {
    "noun": "Muff",
    "article": "der"
  },
  {
    "noun": "Muffelofen",
    "article": "der"
  },
  {
    "noun": "Muffin",
    "article": "der"
  },
  {
    "noun": "Mulatte",
    "article": "der"
  },
  {
    "noun": "Mulch",
    "article": "der"
  },
  {
    "noun": "Mullah",
    "article": "der"
  },
  {
    "noun": "Mulm",
    "article": "der"
  },
  {
    "noun": "Multilateralismus",
    "article": "der"
  },
  {
    "noun": "Multiplier",
    "article": "der"
  },
  {
    "noun": "Multiplikand",
    "article": "der"
  },
  {
    "noun": "Multiplikator",
    "article": "der"
  },
  {
    "noun": "Multipol",
    "article": "der"
  },
  {
    "noun": "Multivibrator",
    "article": "der"
  },
  {
    "noun": "Mumm",
    "article": "der"
  },
  {
    "noun": "Mummenschanz",
    "article": "der"
  },
  {
    "noun": "Mumpitz",
    "article": "der"
  },
  {
    "noun": "Mumps",
    "article": "der"
  },
  {
    "noun": "Mund",
    "article": "der"
  },
  {
    "noun": "Mundartdichter",
    "article": "der"
  },
  {
    "noun": "Mundartsprecher",
    "article": "der"
  },
  {
    "noun": "Mundgeruch",
    "article": "der"
  },
  {
    "noun": "Mundraub",
    "article": "der"
  },
  {
    "noun": "Mundschenk",
    "article": "der"
  },
  {
    "noun": "Mundschutz",
    "article": "der"
  },
  {
    "noun": "Mundverkehr",
    "article": "der"
  },
  {
    "noun": "Mundvorrat",
    "article": "der"
  },
  {
    "noun": "Mundwinkel",
    "article": "der"
  },
  {
    "noun": "Mungo",
    "article": "der"
  },
  {
    "noun": "Munitionsbunker",
    "article": "der"
  },
  {
    "noun": "Muntermacher",
    "article": "der"
  },
  {
    "noun": "Murks",
    "article": "der"
  },
  {
    "noun": "Muschelkrebs",
    "article": "der"
  },
  {
    "noun": "Muschik",
    "article": "der"
  },
  {
    "noun": "Museumsaufseher",
    "article": "der"
  },
  {
    "noun": "Museumsdirektor",
    "article": "der"
  },
  {
    "noun": "Musikant",
    "article": "der"
  },
  {
    "noun": "Musikantenknochen",
    "article": "der"
  },
  {
    "noun": "Musikautomat",
    "article": "der"
  },
  {
    "noun": "Musiker",
    "article": "der"
  },
  {
    "noun": "Musikfreund",
    "article": "der"
  },
  {
    "noun": "Musikgeschmack",
    "article": "der"
  },
  {
    "noun": "Musikliebhaber",
    "article": "der"
  },
  {
    "noun": "Musikologe",
    "article": "der"
  },
  {
    "noun": "Musikpreis",
    "article": "der"
  },
  {
    "noun": "Musikschrank",
    "article": "der"
  },
  {
    "noun": "Musikverlag",
    "article": "der"
  },
  {
    "noun": "Musikwissenschaftler",
    "article": "der"
  },
  {
    "noun": "Muskat",
    "article": "der"
  },
  {
    "noun": "Muskel",
    "article": "der"
  },
  {
    "noun": "Muskelfaserriss",
    "article": "der"
  },
  {
    "noun": "Muskelkater",
    "article": "der"
  },
  {
    "noun": "Muskelkrampf",
    "article": "der"
  },
  {
    "noun": "Muskelmagen",
    "article": "der"
  },
  {
    "noun": "Muskelmann",
    "article": "der"
  },
  {
    "noun": "Muskelprotz",
    "article": "der"
  },
  {
    "noun": "Muskelriss",
    "article": "der"
  },
  {
    "noun": "Muskelschmerz",
    "article": "der"
  },
  {
    "noun": "Muskelschwund",
    "article": "der"
  },
  {
    "noun": "Musketier",
    "article": "der"
  },
  {
    "noun": "Musselin",
    "article": "der"
  },
  {
    "noun": "Mustang",
    "article": "der"
  },
  {
    "noun": "Musterbetrieb",
    "article": "der"
  },
  {
    "noun": "Musterbrief",
    "article": "der"
  },
  {
    "noun": "Musterhaushalt",
    "article": "der"
  },
  {
    "noun": "Musterknabe",
    "article": "der"
  },
  {
    "noun": "Musterkoffer",
    "article": "der"
  },
  {
    "noun": "Musterprozess",
    "article": "der"
  },
  {
    "noun": "Musterschutz",
    "article": "der"
  },
  {
    "noun": "Mustervertrag",
    "article": "der"
  },
  {
    "noun": "Musterzeichner",
    "article": "der"
  },
  {
    "noun": "Mut",
    "article": "der"
  },
  {
    "noun": "Mutismus",
    "article": "der"
  },
  {
    "noun": "Mutterboden",
    "article": "der"
  },
  {
    "noun": "Mutterinstinkt",
    "article": "der"
  },
  {
    "noun": "Mutterkomplex",
    "article": "der"
  },
  {
    "noun": "Mutterkonzern",
    "article": "der"
  },
  {
    "noun": "Mutterkuchen",
    "article": "der"
  },
  {
    "noun": "Mutterleib",
    "article": "der"
  },
  {
    "noun": "Muttermund",
    "article": "der"
  },
  {
    "noun": "Mutterpass",
    "article": "der"
  },
  {
    "noun": "Mutterschaftsurlaub",
    "article": "der"
  },
  {
    "noun": "Mutterschutz",
    "article": "der"
  },
  {
    "noun": "Muttersprachler",
    "article": "der"
  },
  {
    "noun": "Muttertag",
    "article": "der"
  },
  {
    "noun": "Mutwille",
    "article": "der"
  },
  {
    "noun": "Myograf",
    "article": "der"
  },
  {
    "noun": "Myrtenkranz",
    "article": "der"
  },
  {
    "noun": "Myrtenzweig",
    "article": "der"
  },
  {
    "noun": "Mystiker",
    "article": "der"
  },
  {
    "noun": "Mystizismus",
    "article": "der"
  },
  {
    "noun": "Mythos",
    "article": "der"
  },
  {
    "noun": "Egotismus",
    "article": "der"
  },
  {
    "noun": "Egotist",
    "article": "der"
  },
  {
    "noun": "Egotrip",
    "article": "der"
  },
  {
    "noun": "Egozentriker",
    "article": "der"
  },
  {
    "noun": "Eheberater",
    "article": "der"
  },
  {
    "noun": "Ehebrecher",
    "article": "der"
  },
  {
    "noun": "Ehebruch",
    "article": "der"
  },
  {
    "noun": "Ehebund",
    "article": "der"
  },
  {
    "noun": "Ehefeind",
    "article": "der"
  },
  {
    "noun": "Ehegatte",
    "article": "der"
  },
  {
    "noun": "Ehegemahl",
    "article": "der"
  },
  {
    "noun": "Ehegenosse",
    "article": "der"
  },
  {
    "noun": "Ehekontrakt",
    "article": "der"
  },
  {
    "noun": "Ehekrach",
    "article": "der"
  },
  {
    "noun": "Ehemann",
    "article": "der"
  },
  {
    "noun": "Ehepartner",
    "article": "der"
  },
  {
    "noun": "Ehering",
    "article": "der"
  },
  {
    "noun": "Ehestand",
    "article": "der"
  },
  {
    "noun": "Ehestifter",
    "article": "der"
  },
  {
    "noun": "Ehestreit",
    "article": "der"
  },
  {
    "noun": "Ehevermittler",
    "article": "der"
  },
  {
    "noun": "Ehevertrag",
    "article": "der"
  },
  {
    "noun": "Ehrabschneider",
    "article": "der"
  },
  {
    "noun": "Ehrendoktor",
    "article": "der"
  },
  {
    "noun": "Ehreneintritt",
    "article": "der"
  },
  {
    "noun": "Ehrengast",
    "article": "der"
  },
  {
    "noun": "Ehrenkodex",
    "article": "der"
  },
  {
    "noun": "Ehrenkonsul",
    "article": "der"
  },
  {
    "noun": "Ehrenmann",
    "article": "der"
  },
  {
    "noun": "Ehrenname",
    "article": "der"
  },
  {
    "noun": "Ehrenplatz",
    "article": "der"
  },
  {
    "noun": "Ehrentag",
    "article": "der"
  },
  {
    "noun": "Ehrentanz",
    "article": "der"
  },
  {
    "noun": "Ehrentitel",
    "article": "der"
  },
  {
    "noun": "Ehrentreffer",
    "article": "der"
  },
  {
    "noun": "Ehrgeiz",
    "article": "der"
  },
  {
    "noun": "Ehrgeizling",
    "article": "der"
  },
  {
    "noun": "Ehrverlust",
    "article": "der"
  },
  {
    "noun": "Eibisch",
    "article": "der"
  },
  {
    "noun": "Eibischtee",
    "article": "der"
  },
  {
    "noun": "Eichbaum",
    "article": "der"
  },
  {
    "noun": "Eichenbaum",
    "article": "der"
  },
  {
    "noun": "Eichenwald",
    "article": "der"
  },
  {
    "noun": "Eicher",
    "article": "der"
  },
  {
    "noun": "Eichmeister",
    "article": "der"
  },
  {
    "noun": "Eichschein",
    "article": "der"
  },
  {
    "noun": "Eichstempel",
    "article": "der"
  },
  {
    "noun": "Eichstrich",
    "article": "der"
  },
  {
    "noun": "Eichwert",
    "article": "der"
  },
  {
    "noun": "Eid",
    "article": "der"
  },
  {
    "noun": "Eidam",
    "article": "der"
  },
  {
    "noun": "Eidbruch",
    "article": "der"
  },
  {
    "noun": "Eideshelfer",
    "article": "der"
  },
  {
    "noun": "Eidgenosse",
    "article": "der"
  },
  {
    "noun": "Eierbecher",
    "article": "der"
  },
  {
    "noun": "Eierkocher",
    "article": "der"
  },
  {
    "noun": "Eierkopf",
    "article": "der"
  },
  {
    "noun": "Eierkuchen",
    "article": "der"
  },
  {
    "noun": "Eierschaum",
    "article": "der"
  },
  {
    "noun": "Eierschwamm",
    "article": "der"
  },
  {
    "noun": "Eierstock",
    "article": "der"
  },
  {
    "noun": "Eiertanz",
    "article": "der"
  },
  {
    "noun": "Eifer",
    "article": "der"
  },
  {
    "noun": "Eiferer",
    "article": "der"
  },
  {
    "noun": "Eiffelturm",
    "article": "der"
  },
  {
    "noun": "Eigenanteil",
    "article": "der"
  },
  {
    "noun": "Eigenantrieb",
    "article": "der"
  },
  {
    "noun": "Eigenbedarf",
    "article": "der"
  },
  {
    "noun": "Eigenbesitz",
    "article": "der"
  },
  {
    "noun": "Eigenbesitzer",
    "article": "der"
  },
  {
    "noun": "Eigenbetrieb",
    "article": "der"
  },
  {
    "noun": "Eigenname",
    "article": "der"
  },
  {
    "noun": "Eigennutz",
    "article": "der"
  },
  {
    "noun": "Eigensinn",
    "article": "der"
  },
  {
    "noun": "Eigentumserwerb",
    "article": "der"
  },
  {
    "noun": "Eigentumsverlust",
    "article": "der"
  },
  {
    "noun": "Eigentumsvorbehalt",
    "article": "der"
  },
  {
    "noun": "Eigenverbrauch",
    "article": "der"
  },
  {
    "noun": "Eigenwechsel",
    "article": "der"
  },
  {
    "noun": "Eigenwert",
    "article": "der"
  },
  {
    "noun": "Eigenwille",
    "article": "der"
  },
  {
    "noun": "Eigner",
    "article": "der"
  },
  {
    "noun": "Eignungstest",
    "article": "der"
  },
  {
    "noun": "Eilauftrag",
    "article": "der"
  },
  {
    "noun": "Eilbote",
    "article": "der"
  },
  {
    "noun": "Eilbrief",
    "article": "der"
  },
  {
    "noun": "Eileiter",
    "article": "der"
  },
  {
    "noun": "Eilmarsch",
    "article": "der"
  },
  {
    "noun": "Eilschritt",
    "article": "der"
  },
  {
    "noun": "Eilzug",
    "article": "der"
  },
  {
    "noun": "Eimer",
    "article": "der"
  },
  {
    "noun": "Eimerkettenbagger",
    "article": "der"
  },
  {
    "noun": "Einakter",
    "article": "der"
  },
  {
    "noun": "Einband",
    "article": "der"
  },
  {
    "noun": "Einbau",
    "article": "der"
  },
  {
    "noun": "Einbaumotor",
    "article": "der"
  },
  {
    "noun": "Einbauschrank",
    "article": "der"
  },
  {
    "noun": "Einblick",
    "article": "der"
  },
  {
    "noun": "Einbrecher",
    "article": "der"
  },
  {
    "noun": "Einbruch",
    "article": "der"
  },
  {
    "noun": "Einbruchdiebstahl",
    "article": "der"
  },
  {
    "noun": "Einbruchsdiebstahl",
    "article": "der"
  },
  {
    "noun": "Eindecker",
    "article": "der"
  },
  {
    "noun": "Eindringling",
    "article": "der"
  },
  {
    "noun": "Eindruck",
    "article": "der"
  },
  {
    "noun": "Einer",
    "article": "der"
  },
  {
    "noun": "Einfall",
    "article": "der"
  },
  {
    "noun": "Einfallsreichtum",
    "article": "der"
  },
  {
    "noun": "Einfallswinkel",
    "article": "der"
  },
  {
    "noun": "Einfaltspinsel",
    "article": "der"
  },
  {
    "noun": "Einfang",
    "article": "der"
  },
  {
    "noun": "Einfangprozess",
    "article": "der"
  },
  {
    "noun": "Einflug",
    "article": "der"
  },
  {
    "noun": "Einfluss",
    "article": "der"
  },
  {
    "noun": "Einflussbereich",
    "article": "der"
  },
  {
    "noun": "Einfuhrhafen",
    "article": "der"
  },
  {
    "noun": "Einfuhrhandel",
    "article": "der"
  },
  {
    "noun": "Einfuhrstopp",
    "article": "der"
  },
  {
    "noun": "Einfuhrvertrag",
    "article": "der"
  },
  {
    "noun": "Einfuhrzoll",
    "article": "der"
  },
  {
    "noun": "Eingabefehler",
    "article": "der"
  },
  {
    "noun": "Eingang",
    "article": "der"
  },
  {
    "noun": "Eingangssteuersatz",
    "article": "der"
  },
  {
    "noun": "Eingangstag",
    "article": "der"
  },
  {
    "noun": "Eingangstest",
    "article": "der"
  },
  {
    "noun": "Eingangsvermerk",
    "article": "der"
  },
  {
    "noun": "Eingangszoll",
    "article": "der"
  },
  {
    "noun": "Eingeweidebruch",
    "article": "der"
  },
  {
    "noun": "Eingeweidewurm",
    "article": "der"
  },
  {
    "noun": "Eingriff",
    "article": "der"
  },
  {
    "noun": "Einguss",
    "article": "der"
  },
  {
    "noun": "Einhandmischer",
    "article": "der"
  },
  {
    "noun": "Einhebelmischer",
    "article": "der"
  },
  {
    "noun": "Einheitspreis",
    "article": "der"
  },
  {
    "noun": "Einheitsstaat",
    "article": "der"
  },
  {
    "noun": "Einheitstarif",
    "article": "der"
  },
  {
    "noun": "Einheitswert",
    "article": "der"
  },
  {
    "noun": "Einhelfer",
    "article": "der"
  },
  {
    "noun": "Einhufer",
    "article": "der"
  },
  {
    "noun": "Einigungsversuch",
    "article": "der"
  },
  {
    "noun": "Einigungsvertrag",
    "article": "der"
  },
  {
    "noun": "Einkauf",
    "article": "der"
  },
  {
    "noun": "Einkaufsbeutel",
    "article": "der"
  },
  {
    "noun": "Einkaufsbummel",
    "article": "der"
  },
  {
    "noun": "Einkaufskorb",
    "article": "der"
  },
  {
    "noun": "Einkaufspreis",
    "article": "der"
  },
  {
    "noun": "Einkaufswagen",
    "article": "der"
  },
  {
    "noun": "Einklang",
    "article": "der"
  },
  {
    "noun": "Einkommensausfall",
    "article": "der"
  },
  {
    "noun": "Einkommenszuwachs",
    "article": "der"
  },
  {
    "noun": "Einkristall",
    "article": "der"
  },
  {
    "noun": "Einlass",
    "article": "der"
  },
  {
    "noun": "Einlauf",
    "article": "der"
  },
  {
    "noun": "Einlegeboden",
    "article": "der"
  },
  {
    "noun": "Einleger",
    "article": "der"
  },
  {
    "noun": "Einlieferer",
    "article": "der"
  },
  {
    "noun": "Einlieferungsschein",
    "article": "der"
  },
  {
    "noun": "Einmachtopf",
    "article": "der"
  },
  {
    "noun": "Einmarsch",
    "article": "der"
  },
  {
    "noun": "Einmaster",
    "article": "der"
  },
  {
    "noun": "Einnehmer",
    "article": "der"
  },
  {
    "noun": "Einpauker",
    "article": "der"
  },
  {
    "noun": "Einpersonenhaushalt",
    "article": "der"
  },
  {
    "noun": "Einphasenstrom",
    "article": "der"
  },
  {
    "noun": "Einphasenwechselstrom",
    "article": "der"
  },
  {
    "noun": "Einreichungstermin",
    "article": "der"
  },
  {
    "noun": "Einrichter",
    "article": "der"
  },
  {
    "noun": "Einrichtungsgegenstand",
    "article": "der"
  },
  {
    "noun": "Einriss",
    "article": "der"
  },
  {
    "noun": "Einsatz",
    "article": "der"
  },
  {
    "noun": "Einsatzbefehl",
    "article": "der"
  },
  {
    "noun": "Einsatzleiter",
    "article": "der"
  },
  {
    "noun": "Einsatzplan",
    "article": "der"
  },
  {
    "noun": "Einsatzring",
    "article": "der"
  },
  {
    "noun": "Einsatzwagen",
    "article": "der"
  },
  {
    "noun": "Einschalthebel",
    "article": "der"
  },
  {
    "noun": "Einschlag",
    "article": "der"
  },
  {
    "noun": "Einschluss",
    "article": "der"
  },
  {
    "noun": "Einschnitt",
    "article": "der"
  },
  {
    "noun": "Einschreibebrief",
    "article": "der"
  },
  {
    "noun": "Einschub",
    "article": "der"
  },
  {
    "noun": "Einschuss",
    "article": "der"
  },
  {
    "noun": "Einsender",
    "article": "der"
  },
  {
    "noun": "Einsendeschluss",
    "article": "der"
  },
  {
    "noun": "Einser",
    "article": "der"
  },
  {
    "noun": "Einsiedler",
    "article": "der"
  },
  {
    "noun": "Einsiedlerkrebs",
    "article": "der"
  },
  {
    "noun": "Einsilber",
    "article": "der"
  },
  {
    "noun": "Einspritzmotor",
    "article": "der"
  },
  {
    "noun": "Einspruch",
    "article": "der"
  },
  {
    "noun": "Einstand",
    "article": "der"
  },
  {
    "noun": "Einstandspreis",
    "article": "der"
  },
  {
    "noun": "Einsteiger",
    "article": "der"
  },
  {
    "noun": "Einsteller",
    "article": "der"
  },
  {
    "noun": "Einstellknopf",
    "article": "der"
  },
  {
    "noun": "Einstellraum",
    "article": "der"
  },
  {
    "noun": "Einstellungsstopp",
    "article": "der"
  },
  {
    "noun": "Einstellungstermin",
    "article": "der"
  },
  {
    "noun": "Einstellungstest",
    "article": "der"
  },
  {
    "noun": "Einstich",
    "article": "der"
  },
  {
    "noun": "Einstieg",
    "article": "der"
  },
  {
    "noun": "Einsturz",
    "article": "der"
  },
  {
    "noun": "Eintausch",
    "article": "der"
  },
  {
    "noun": "Eintopf",
    "article": "der"
  },
  {
    "noun": "Eintrag",
    "article": "der"
  },
  {
    "noun": "Eintragungsvermerk",
    "article": "der"
  },
  {
    "noun": "Eintreiber",
    "article": "der"
  },
  {
    "noun": "Eintritt",
    "article": "der"
  },
  {
    "noun": "Eintrittspreis",
    "article": "der"
  },
  {
    "noun": "Einwand",
    "article": "der"
  },
  {
    "noun": "Einwanderer",
    "article": "der"
  },
  {
    "noun": "Einwanderungsstrom",
    "article": "der"
  },
  {
    "noun": "Einweghahn",
    "article": "der"
  },
  {
    "noun": "Einweiser",
    "article": "der"
  },
  {
    "noun": "Einwirkungsbereich",
    "article": "der"
  },
  {
    "noun": "Einwohner",
    "article": "der"
  },
  {
    "noun": "Einwurf",
    "article": "der"
  },
  {
    "noun": "Einzahler",
    "article": "der"
  },
  {
    "noun": "Einzahlungsbeleg",
    "article": "der"
  },
  {
    "noun": "Einzahlungsschalter",
    "article": "der"
  },
  {
    "noun": "Einzahlungsschein",
    "article": "der"
  },
  {
    "noun": "Einzelantrieb",
    "article": "der"
  },
  {
    "noun": "Einzelblatteinzug",
    "article": "der"
  },
  {
    "noun": "Einzelfall",
    "article": "der"
  },
  {
    "noun": "Einzelhandel",
    "article": "der"
  },
  {
    "noun": "Einzelhandelskaufmann",
    "article": "der"
  },
  {
    "noun": "Einzelhandelspreis",
    "article": "der"
  },
  {
    "noun": "Einzelkampf",
    "article": "der"
  },
  {
    "noun": "Einzelkaufmann",
    "article": "der"
  },
  {
    "noun": "Einzelmensch",
    "article": "der"
  },
  {
    "noun": "Einzelposten",
    "article": "der"
  },
  {
    "noun": "Einzelpreis",
    "article": "der"
  },
  {
    "noun": "Einzelrichter",
    "article": "der"
  },
  {
    "noun": "Einzelstaat",
    "article": "der"
  },
  {
    "noun": "Einzelunternehmer",
    "article": "der"
  },
  {
    "noun": "Einzelunterricht",
    "article": "der"
  },
  {
    "noun": "Einzelverkauf",
    "article": "der"
  },
  {
    "noun": "Einzelvertrag",
    "article": "der"
  },
  {
    "noun": "Einzelwert",
    "article": "der"
  },
  {
    "noun": "Einzug",
    "article": "der"
  },
  {
    "noun": "Einzylindermotor",
    "article": "der"
  },
  {
    "noun": "Eisbecher",
    "article": "der"
  },
  {
    "noun": "Eisberg",
    "article": "der"
  },
  {
    "noun": "Eisbeutel",
    "article": "der"
  },
  {
    "noun": "Eisblock",
    "article": "der"
  },
  {
    "noun": "Eisbrecher",
    "article": "der"
  },
  {
    "noun": "Eischnee",
    "article": "der"
  },
  {
    "noun": "Eisenbahnarbeiter",
    "article": "der"
  },
  {
    "noun": "Eisenbahndamm",
    "article": "der"
  },
  {
    "noun": "Eisenbahner",
    "article": "der"
  },
  {
    "noun": "Eisenbahnfahrplan",
    "article": "der"
  },
  {
    "noun": "Eisenbahnknotenpunkt",
    "article": "der"
  },
  {
    "noun": "Eisenbahnschaffner",
    "article": "der"
  },
  {
    "noun": "Eisenbahntarif",
    "article": "der"
  },
  {
    "noun": "Eisenbahntransport",
    "article": "der"
  },
  {
    "noun": "Eisenbahnverkehr",
    "article": "der"
  },
  {
    "noun": "Eisenbahnwagen",
    "article": "der"
  },
  {
    "noun": "Eisenbahnwaggon",
    "article": "der"
  },
  {
    "noun": "Eisenbahnzug",
    "article": "der"
  },
  {
    "noun": "Eisenbeschlag",
    "article": "der"
  },
  {
    "noun": "Eisenbeton",
    "article": "der"
  },
  {
    "noun": "Eisendraht",
    "article": "der"
  },
  {
    "noun": "Eisenfresser",
    "article": "der"
  },
  {
    "noun": "Eisengehalt",
    "article": "der"
  },
  {
    "noun": "Eisenguss",
    "article": "der"
  },
  {
    "noun": "Eisenhaken",
    "article": "der"
  },
  {
    "noun": "Eisenhammer",
    "article": "der"
  },
  {
    "noun": "Eisenhut",
    "article": "der"
  },
  {
    "noun": "Eisenkern",
    "article": "der"
  },
  {
    "noun": "Eisenkies",
    "article": "der"
  },
  {
    "noun": "Eisenkiesel",
    "article": "der"
  },
  {
    "noun": "Eisenkitt",
    "article": "der"
  },
  {
    "noun": "Eisenkoks",
    "article": "der"
  },
  {
    "noun": "Eisenmangel",
    "article": "der"
  },
  {
    "noun": "Eisenring",
    "article": "der"
  },
  {
    "noun": "Eisenspat",
    "article": "der"
  },
  {
    "noun": "Eisenstaub",
    "article": "der"
  },
  {
    "noun": "Eishockeyspieler",
    "article": "der"
  },
  {
    "noun": "Eiskaffee",
    "article": "der"
  },
  {
    "noun": "Eiskasten",
    "article": "der"
  },
  {
    "noun": "Eiskeller",
    "article": "der"
  },
  {
    "noun": "Eiskristall",
    "article": "der"
  },
  {
    "noun": "Eiskunstlauf",
    "article": "der"
  },
  {
    "noun": "Eislauf",
    "article": "der"
  },
  {
    "noun": "Eispickel",
    "article": "der"
  },
  {
    "noun": "Eisprung",
    "article": "der"
  },
  {
    "noun": "Eispunkt",
    "article": "der"
  },
  {
    "noun": "Eissalon",
    "article": "der"
  },
  {
    "noun": "Eisschnellauf",
    "article": "der"
  },
  {
    "noun": "Eisschrank",
    "article": "der"
  },
  {
    "noun": "Eisstausee",
    "article": "der"
  },
  {
    "noun": "Eisstock",
    "article": "der"
  },
  {
    "noun": "Eissturmvogel",
    "article": "der"
  },
  {
    "noun": "Eistaucher",
    "article": "der"
  },
  {
    "noun": "Eistee",
    "article": "der"
  },
  {
    "noun": "Eisvogel",
    "article": "der"
  },
  {
    "noun": "Eiszapfen",
    "article": "der"
  },
  {
    "noun": "Eiter",
    "article": "der"
  },
  {
    "noun": "Eiterherd",
    "article": "der"
  },
  {
    "noun": "Eiterpfropf",
    "article": "der"
  },
  {
    "noun": "Eiterpickel",
    "article": "der"
  },
  {
    "noun": "Ejektor",
    "article": "der"
  },
  {
    "noun": "Eklat",
    "article": "der"
  },
  {
    "noun": "Eklektiker",
    "article": "der"
  },
  {
    "noun": "Eklektizismus",
    "article": "der"
  },
  {
    "noun": "Eklogit",
    "article": "der"
  },
  {
    "noun": "Elan",
    "article": "der"
  },
  {
    "noun": "Elativ",
    "article": "der"
  },
  {
    "noun": "Elbeseitenkanal",
    "article": "der"
  },
  {
    "noun": "Elch",
    "article": "der"
  },
  {
    "noun": "Elchbulle",
    "article": "der"
  },
  {
    "noun": "Elefant",
    "article": "der"
  },
  {
    "noun": "Elefantenbulle",
    "article": "der"
  },
  {
    "noun": "Elektriker",
    "article": "der"
  },
  {
    "noun": "Elektroartikel",
    "article": "der"
  },
  {
    "noun": "Elektrobus",
    "article": "der"
  },
  {
    "noun": "Elektroenzephalograf",
    "article": "der"
  },
  {
    "noun": "Elektroherd",
    "article": "der"
  },
  {
    "noun": "Elektroingenieur",
    "article": "der"
  },
  {
    "noun": "Elektroinstallateur",
    "article": "der"
  },
  {
    "noun": "Elektrokardiograf",
    "article": "der"
  },
  {
    "noun": "Elektrokauter",
    "article": "der"
  },
  {
    "noun": "Elektrokocher",
    "article": "der"
  },
  {
    "noun": "Elektrokonzern",
    "article": "der"
  },
  {
    "noun": "Elektrokrampf",
    "article": "der"
  },
  {
    "noun": "Elektrolyseur",
    "article": "der"
  },
  {
    "noun": "Elektrolyt",
    "article": "der"
  },
  {
    "noun": "Elektromagnet",
    "article": "der"
  },
  {
    "noun": "Elektromagnetismus",
    "article": "der"
  },
  {
    "noun": "Elektromaschinenbauer",
    "article": "der"
  },
  {
    "noun": "Elektromechaniker",
    "article": "der"
  },
  {
    "noun": "Elektromonteur",
    "article": "der"
  },
  {
    "noun": "Elektromotor",
    "article": "der"
  },
  {
    "noun": "Elektronenradius",
    "article": "der"
  },
  {
    "noun": "Elektronenrechner",
    "article": "der"
  },
  {
    "noun": "Elektronenspin",
    "article": "der"
  },
  {
    "noun": "Elektronenstrahl",
    "article": "der"
  },
  {
    "noun": "Elektronenstrom",
    "article": "der"
  },
  {
    "noun": "Elektroniker",
    "article": "der"
  },
  {
    "noun": "Elektronikschrott",
    "article": "der"
  },
  {
    "noun": "Elektroofen",
    "article": "der"
  },
  {
    "noun": "Elektrorasierer",
    "article": "der"
  },
  {
    "noun": "Elektroschock",
    "article": "der"
  },
  {
    "noun": "Elektrosmog",
    "article": "der"
  },
  {
    "noun": "Elektrotechniker",
    "article": "der"
  },
  {
    "noun": "Elektrotonus",
    "article": "der"
  },
  {
    "noun": "Elektrozaun",
    "article": "der"
  },
  {
    "noun": "Elementarbegriff",
    "article": "der"
  },
  {
    "noun": "Elementarmagnet",
    "article": "der"
  },
  {
    "noun": "Elementarschaden",
    "article": "der"
  },
  {
    "noun": "Elementarunterricht",
    "article": "der"
  },
  {
    "noun": "Elevator",
    "article": "der"
  },
  {
    "noun": "Eleve",
    "article": "der"
  },
  {
    "noun": "Elfenbeinturm",
    "article": "der"
  },
  {
    "noun": "Elfenreigen",
    "article": "der"
  },
  {
    "noun": "Elfer",
    "article": "der"
  },
  {
    "noun": "Elfmeter",
    "article": "der"
  },
  {
    "noun": "Elfmeterpunkt",
    "article": "der"
  },
  {
    "noun": "Elitismus",
    "article": "der"
  },
  {
    "noun": "Ellbogen",
    "article": "der"
  },
  {
    "noun": "Ellbogenmensch",
    "article": "der"
  },
  {
    "noun": "Ellenbogen",
    "article": "der"
  },
  {
    "noun": "Elternabend",
    "article": "der"
  },
  {
    "noun": "Elternbeirat",
    "article": "der"
  },
  {
    "noun": "Elternteil",
    "article": "der"
  },
  {
    "noun": "Eluvialhorizont",
    "article": "der"
  },
  {
    "noun": "Emaildraht",
    "article": "der"
  },
  {
    "noun": "Emaillierofen",
    "article": "der"
  },
  {
    "noun": "Embolus",
    "article": "der"
  },
  {
    "noun": "Embryo",
    "article": "der"
  },
  {
    "noun": "Emigrant",
    "article": "der"
  },
  {
    "noun": "Emir",
    "article": "der"
  },
  {
    "noun": "Emissionskurs",
    "article": "der"
  },
  {
    "noun": "Emissionsschutz",
    "article": "der"
  },
  {
    "noun": "Emissionswert",
    "article": "der"
  },
  {
    "noun": "Emittent",
    "article": "der"
  },
  {
    "noun": "Emitter",
    "article": "der"
  },
  {
    "noun": "Empfang",
    "article": "der"
  },
  {
    "noun": "Empfangschef",
    "article": "der"
  },
  {
    "noun": "Empfangsraum",
    "article": "der"
  },
  {
    "noun": "Empfangsschalter",
    "article": "der"
  },
  {
    "noun": "Empfangsschein",
    "article": "der"
  },
  {
    "noun": "Empfangstag",
    "article": "der"
  },
  {
    "noun": "Empfehlungsbrief",
    "article": "der"
  },
  {
    "noun": "Empfindungsnerv",
    "article": "der"
  },
  {
    "noun": "Empiriker",
    "article": "der"
  },
  {
    "noun": "Empirismus",
    "article": "der"
  },
  {
    "noun": "Emu",
    "article": "der"
  },
  {
    "noun": "Emulator",
    "article": "der"
  },
  {
    "noun": "Emulgator",
    "article": "der"
  },
  {
    "noun": "Endbescheid",
    "article": "der"
  },
  {
    "noun": "Endbetrag",
    "article": "der"
  },
  {
    "noun": "Enddarm",
    "article": "der"
  },
  {
    "noun": "Endeffekt",
    "article": "der"
  },
  {
    "noun": "Endemismus",
    "article": "der"
  },
  {
    "noun": "Endkampf",
    "article": "der"
  },
  {
    "noun": "Endlauf",
    "article": "der"
  },
  {
    "noun": "Endokrinologe",
    "article": "der"
  },
  {
    "noun": "Endomorphismus",
    "article": "der"
  },
  {
    "noun": "Endophyt",
    "article": "der"
  },
  {
    "noun": "Endpunkt",
    "article": "der"
  },
  {
    "noun": "Endsieg",
    "article": "der"
  },
  {
    "noun": "Endspurt",
    "article": "der"
  },
  {
    "noun": "Endstand",
    "article": "der"
  },
  {
    "noun": "Endverbraucher",
    "article": "der"
  },
  {
    "noun": "Endverbraucherpreis",
    "article": "der"
  },
  {
    "noun": "Endzustand",
    "article": "der"
  },
  {
    "noun": "Endzweck",
    "article": "der"
  },
  {
    "noun": "Energieaufwand",
    "article": "der"
  },
  {
    "noun": "Energiebedarf",
    "article": "der"
  },
  {
    "noun": "Energiebereich",
    "article": "der"
  },
  {
    "noun": "Energiegehalt",
    "article": "der"
  },
  {
    "noun": "Energiehaushalt",
    "article": "der"
  },
  {
    "noun": "Energiekonzern",
    "article": "der"
  },
  {
    "noun": "Energielieferant",
    "article": "der"
  },
  {
    "noun": "Energiesatz",
    "article": "der"
  },
  {
    "noun": "Energieverbrauch",
    "article": "der"
  },
  {
    "noun": "Energieverlust",
    "article": "der"
  },
  {
    "noun": "Energievorrat",
    "article": "der"
  },
  {
    "noun": "Engel",
    "article": "der"
  },
  {
    "noun": "Engelmacher",
    "article": "der"
  },
  {
    "noun": "Engerling",
    "article": "der"
  },
  {
    "noun": "Englischkurs",
    "article": "der"
  },
  {
    "noun": "Englischlehrer",
    "article": "der"
  },
  {
    "noun": "Engpass",
    "article": "der"
  },
  {
    "noun": "Engroshandel",
    "article": "der"
  },
  {
    "noun": "Enkel",
    "article": "der"
  },
  {
    "noun": "Enkelsohn",
    "article": "der"
  },
  {
    "noun": "Enstatit",
    "article": "der"
  },
  {
    "noun": "Entdecker",
    "article": "der"
  },
  {
    "noun": "Entenbraten",
    "article": "der"
  },
  {
    "noun": "Entenschnabel",
    "article": "der"
  },
  {
    "noun": "Entenwal",
    "article": "der"
  },
  {
    "noun": "Enterhaken",
    "article": "der"
  },
  {
    "noun": "Enterich",
    "article": "der"
  },
  {
    "noun": "Entertainer",
    "article": "der"
  },
  {
    "noun": "Entfall",
    "article": "der"
  },
  {
    "noun": "Entfernungsmesser",
    "article": "der"
  },
  {
    "noun": "Entfeuchter",
    "article": "der"
  },
  {
    "noun": "Entfroster",
    "article": "der"
  },
  {
    "noun": "Enthusiasmus",
    "article": "der"
  },
  {
    "noun": "Enthusiast",
    "article": "der"
  },
  {
    "noun": "Entkerner",
    "article": "der"
  },
  {
    "noun": "Entlader",
    "article": "der"
  },
  {
    "noun": "Entlassungsgrund",
    "article": "der"
  },
  {
    "noun": "Entlassungsschein",
    "article": "der"
  },
  {
    "noun": "Entlastungszeuge",
    "article": "der"
  },
  {
    "noun": "Entlastungszug",
    "article": "der"
  },
  {
    "noun": "Entleiher",
    "article": "der"
  },
  {
    "noun": "Entomologe",
    "article": "der"
  },
  {
    "noun": "Entsafter",
    "article": "der"
  },
  {
    "noun": "Entscheid",
    "article": "der"
  },
  {
    "noun": "Entscheidungsbedarf",
    "article": "der"
  },
  {
    "noun": "Entscheidungskampf",
    "article": "der"
  },
  {
    "noun": "Entscheidungslauf",
    "article": "der"
  },
  {
    "noun": "Entschluss",
    "article": "der"
  },
  {
    "noun": "Entschuldigungsbrief",
    "article": "der"
  },
  {
    "noun": "Entschuldigungsgrund",
    "article": "der"
  },
  {
    "noun": "Entschuldigungszettel",
    "article": "der"
  },
  {
    "noun": "Entsetzensschrei",
    "article": "der"
  },
  {
    "noun": "Entstauber",
    "article": "der"
  },
  {
    "noun": "Entstehungsort",
    "article": "der"
  },
  {
    "noun": "Entstehungsprozess",
    "article": "der"
  },
  {
    "noun": "Entwickler",
    "article": "der"
  },
  {
    "noun": "Entwicklungsgang",
    "article": "der"
  },
  {
    "noun": "Entwicklungsgrad",
    "article": "der"
  },
  {
    "noun": "Entwicklungshelfer",
    "article": "der"
  },
  {
    "noun": "Entwicklungsingenieur",
    "article": "der"
  },
  {
    "noun": "Entwicklungsplan",
    "article": "der"
  },
  {
    "noun": "Entwicklungsprozess",
    "article": "der"
  },
  {
    "noun": "Entwicklungsroman",
    "article": "der"
  },
  {
    "noun": "Entwicklungsstand",
    "article": "der"
  },
  {
    "noun": "Entwicklungszustand",
    "article": "der"
  },
  {
    "noun": "Entwurf",
    "article": "der"
  },
  {
    "noun": "Entzerrer",
    "article": "der"
  },
  {
    "noun": "Entzifferer",
    "article": "der"
  },
  {
    "noun": "Entzug",
    "article": "der"
  },
  {
    "noun": "Enzephalograf",
    "article": "der"
  },
  {
    "noun": "Enzian",
    "article": "der"
  },
  {
    "noun": "Epidot",
    "article": "der"
  },
  {
    "noun": "Epigone",
    "article": "der"
  },
  {
    "noun": "Epiker",
    "article": "der"
  },
  {
    "noun": "Epikureer",
    "article": "der"
  },
  {
    "noun": "Epileptiker",
    "article": "der"
  },
  {
    "noun": "Epilog",
    "article": "der"
  },
  {
    "noun": "Epiphyt",
    "article": "der"
  },
  {
    "noun": "Epirot",
    "article": "der"
  },
  {
    "noun": "Episkopalismus",
    "article": "der"
  },
  {
    "noun": "Erbadel",
    "article": "der"
  },
  {
    "noun": "Erbanspruch",
    "article": "der"
  },
  {
    "noun": "Erbanteil",
    "article": "der"
  },
  {
    "noun": "Erbauer",
    "article": "der"
  },
  {
    "noun": "Erbbauzins",
    "article": "der"
  },
  {
    "noun": "Erbfaktor",
    "article": "der"
  },
  {
    "noun": "Erbfehler",
    "article": "der"
  },
  {
    "noun": "Erbfeind",
    "article": "der"
  },
  {
    "noun": "Erbfolgekrieg",
    "article": "der"
  },
  {
    "noun": "Erbfolger",
    "article": "der"
  },
  {
    "noun": "Erblasser",
    "article": "der"
  },
  {
    "noun": "Erbprinz",
    "article": "der"
  },
  {
    "noun": "Erbschaden",
    "article": "der"
  },
  {
    "noun": "Erbschaftsanspruch",
    "article": "der"
  },
  {
    "noun": "Erbschein",
    "article": "der"
  },
  {
    "noun": "Erbschleicher",
    "article": "der"
  },
  {
    "noun": "Erbsenbrei",
    "article": "der"
  },
  {
    "noun": "Erbsenstein",
    "article": "der"
  },
  {
    "noun": "Erbverzicht",
    "article": "der"
  },
  {
    "noun": "Erdapfel",
    "article": "der"
  },
  {
    "noun": "Erdaushub",
    "article": "der"
  },
  {
    "noun": "Erdball",
    "article": "der"
  },
  {
    "noun": "Erdbebenanzeiger",
    "article": "der"
  },
  {
    "noun": "Erdbebenherd",
    "article": "der"
  },
  {
    "noun": "Erdbebenmesser",
    "article": "der"
  },
  {
    "noun": "Erdbeerbaum",
    "article": "der"
  },
  {
    "noun": "Erdbeersekt",
    "article": "der"
  },
  {
    "noun": "Erdbewohner",
    "article": "der"
  },
  {
    "noun": "Erdboden",
    "article": "der"
  },
  {
    "noun": "Erdbohrer",
    "article": "der"
  },
  {
    "noun": "Erdenbewohner",
    "article": "der"
  },
  {
    "noun": "Erdfall",
    "article": "der"
  },
  {
    "noun": "Erdgeist",
    "article": "der"
  },
  {
    "noun": "Erdgeruch",
    "article": "der"
  },
  {
    "noun": "Erdkern",
    "article": "der"
  },
  {
    "noun": "Erdklumpen",
    "article": "der"
  },
  {
    "noun": "Erdkundler",
    "article": "der"
  },
  {
    "noun": "Erdling",
    "article": "der"
  },
  {
    "noun": "Erdmagnetismus",
    "article": "der"
  },
  {
    "noun": "Erdmittelpunkt",
    "article": "der"
  },
  {
    "noun": "Erdrutsch",
    "article": "der"
  },
  {
    "noun": "Erdrutschsieg",
    "article": "der"
  },
  {
    "noun": "Erdschatten",
    "article": "der"
  },
  {
    "noun": "Erdschluss",
    "article": "der"
  },
  {
    "noun": "Erdstrom",
    "article": "der"
  },
  {
    "noun": "Erdteil",
    "article": "der"
  },
  {
    "noun": "Erdumfang",
    "article": "der"
  },
  {
    "noun": "Erdumlauf",
    "article": "der"
  },
  {
    "noun": "Erdwall",
    "article": "der"
  },
  {
    "noun": "Eremit",
    "article": "der"
  },
  {
    "noun": "Erfahrungsaustausch",
    "article": "der"
  },
  {
    "noun": "Erfahrungsbereich",
    "article": "der"
  },
  {
    "noun": "Erfahrungsbericht",
    "article": "der"
  },
  {
    "noun": "Erfahrungswert",
    "article": "der"
  },
  {
    "noun": "Erfinder",
    "article": "der"
  },
  {
    "noun": "Erfindergeist",
    "article": "der"
  },
  {
    "noun": "Erfinderschutz",
    "article": "der"
  },
  {
    "noun": "Erfindungsschutz",
    "article": "der"
  },
  {
    "noun": "Erfolg",
    "article": "der"
  },
  {
    "noun": "Erfolgsautor",
    "article": "der"
  },
  {
    "noun": "Erfolgshonorar",
    "article": "der"
  },
  {
    "noun": "Erfolgsnachweis",
    "article": "der"
  },
  {
    "noun": "Erfolgsroman",
    "article": "der"
  },
  {
    "noun": "Erforscher",
    "article": "der"
  },
  {
    "noun": "Erfrischungsraum",
    "article": "der"
  },
  {
    "noun": "Ergograf",
    "article": "der"
  },
  {
    "noun": "Ergotismus",
    "article": "der"
  },
  {
    "noun": "Erguss",
    "article": "der"
  },
  {
    "noun": "Erhalt",
    "article": "der"
  },
  {
    "noun": "Erhalter",
    "article": "der"
  },
  {
    "noun": "Erhitzer",
    "article": "der"
  },
  {
    "noun": "Erholungsaufenthalt",
    "article": "der"
  },
  {
    "noun": "Erholungsraum",
    "article": "der"
  },
  {
    "noun": "Erholungsurlaub",
    "article": "der"
  },
  {
    "noun": "Erholungswert",
    "article": "der"
  },
  {
    "noun": "Erinnerungswert",
    "article": "der"
  },
  {
    "noun": "Erisapfel",
    "article": "der"
  },
  {
    "noun": "Erkenntnisprozess",
    "article": "der"
  },
  {
    "noun": "Erker",
    "article": "der"
  },
  {
    "noun": "Erlass",
    "article": "der"
  },
  {
    "noun": "Erlaubnisschein",
    "article": "der"
  },
  {
    "noun": "Erlebensfall",
    "article": "der"
  },
  {
    "noun": "Erlenmeyerkolben",
    "article": "der"
  },
  {
    "noun": "Erlenzeisig",
    "article": "der"
  },
  {
    "noun": "Ermessensspielraum",
    "article": "der"
  },
  {
    "noun": "Ermittler",
    "article": "der"
  },
  {
    "noun": "Ermittlungsausschuss",
    "article": "der"
  },
  {
    "noun": "Ermittlungsrichter",
    "article": "der"
  },
  {
    "noun": "Erneuerer",
    "article": "der"
  },
  {
    "noun": "Erneuerungsschein",
    "article": "der"
  },
  {
    "noun": "Ernstfall",
    "article": "der"
  },
  {
    "noun": "Erntearbeiter",
    "article": "der"
  },
  {
    "noun": "Ernteausfall",
    "article": "der"
  },
  {
    "noun": "Ernteeinsatz",
    "article": "der"
  },
  {
    "noun": "Ernteertrag",
    "article": "der"
  },
  {
    "noun": "Erntehelfer",
    "article": "der"
  },
  {
    "noun": "Erntekranz",
    "article": "der"
  },
  {
    "noun": "Erntemonat",
    "article": "der"
  },
  {
    "noun": "Erntesegen",
    "article": "der"
  },
  {
    "noun": "Eroberer",
    "article": "der"
  },
  {
    "noun": "Eroberungsfeldzug",
    "article": "der"
  },
  {
    "noun": "Eroberungskrieg",
    "article": "der"
  },
  {
    "noun": "Eros",
    "article": "der"
  },
  {
    "noun": "Erotiker",
    "article": "der"
  },
  {
    "noun": "Erpel",
    "article": "der"
  },
  {
    "noun": "Erpresser",
    "article": "der"
  },
  {
    "noun": "Erpressungsversuch",
    "article": "der"
  },
  {
    "noun": "Erreger",
    "article": "der"
  },
  {
    "noun": "Erretter",
    "article": "der"
  },
  {
    "noun": "Ersatz",
    "article": "der"
  },
  {
    "noun": "Ersatzanspruch",
    "article": "der"
  },
  {
    "noun": "Ersatzdienst",
    "article": "der"
  },
  {
    "noun": "Ersatzfahrer",
    "article": "der"
  },
  {
    "noun": "Ersatzmann",
    "article": "der"
  },
  {
    "noun": "Ersatzreifen",
    "article": "der"
  },
  {
    "noun": "Ersatzspieler",
    "article": "der"
  },
  {
    "noun": "Ersatzstoff",
    "article": "der"
  },
  {
    "noun": "Ersatzwagen",
    "article": "der"
  },
  {
    "noun": "Erschaffer",
    "article": "der"
  },
  {
    "noun": "Erscheinungsort",
    "article": "der"
  },
  {
    "noun": "Erstattungsantrag",
    "article": "der"
  },
  {
    "noun": "Erstdruck",
    "article": "der"
  },
  {
    "noun": "Ersteher",
    "article": "der"
  },
  {
    "noun": "Ersteller",
    "article": "der"
  },
  {
    "noun": "Ersterwerber",
    "article": "der"
  },
  {
    "noun": "Erstflug",
    "article": "der"
  },
  {
    "noun": "Ersthelfer",
    "article": "der"
  },
  {
    "noun": "Erstling",
    "article": "der"
  },
  {
    "noun": "Ersttagsbrief",
    "article": "der"
  },
  {
    "noun": "Ertrag",
    "article": "der"
  },
  {
    "noun": "Ertragsausfall",
    "article": "der"
  },
  {
    "noun": "Ertragswert",
    "article": "der"
  },
  {
    "noun": "Erwartungshorizont",
    "article": "der"
  },
  {
    "noun": "Erweis",
    "article": "der"
  },
  {
    "noun": "Erweiterungsbau",
    "article": "der"
  },
  {
    "noun": "Erwerb",
    "article": "der"
  },
  {
    "noun": "Erwerber",
    "article": "der"
  },
  {
    "noun": "Erwerbssinn",
    "article": "der"
  },
  {
    "noun": "Erwerbstrieb",
    "article": "der"
  },
  {
    "noun": "Erythrozyt",
    "article": "der"
  },
  {
    "noun": "Erzbergbau",
    "article": "der"
  },
  {
    "noun": "Erzbischof",
    "article": "der"
  },
  {
    "noun": "Erzdiakon",
    "article": "der"
  },
  {
    "noun": "Erzengel",
    "article": "der"
  },
  {
    "noun": "Erzeuger",
    "article": "der"
  },
  {
    "noun": "Erzeugerpreis",
    "article": "der"
  },
  {
    "noun": "Erzfeind",
    "article": "der"
  },
  {
    "noun": "Erzgang",
    "article": "der"
  },
  {
    "noun": "Erzgauner",
    "article": "der"
  },
  {
    "noun": "Erzherzog",
    "article": "der"
  },
  {
    "noun": "Erzieher",
    "article": "der"
  },
  {
    "noun": "Erziehungsurlaub",
    "article": "der"
  },
  {
    "noun": "Erzkanzler",
    "article": "der"
  },
  {
    "noun": "Erzpriester",
    "article": "der"
  },
  {
    "noun": "Erzrivale",
    "article": "der"
  },
  {
    "noun": "Erzvater",
    "article": "der"
  },
  {
    "noun": "Esel",
    "article": "der"
  },
  {
    "noun": "Eselhengst",
    "article": "der"
  },
  {
    "noun": "Eselsschrei",
    "article": "der"
  },
  {
    "noun": "Eskimo",
    "article": "der"
  },
  {
    "noun": "Esoteriker",
    "article": "der"
  },
  {
    "noun": "Espresso",
    "article": "der"
  },
  {
    "noun": "Esprit",
    "article": "der"
  },
  {
    "noun": "Essayist",
    "article": "der"
  },
  {
    "noun": "Esser",
    "article": "der"
  },
  {
    "noun": "Essig",
    "article": "der"
  },
  {
    "noun": "Essigstich",
    "article": "der"
  },
  {
    "noun": "Essteller",
    "article": "der"
  },
  {
    "noun": "Esstisch",
    "article": "der"
  },
  {
    "noun": "Este",
    "article": "der"
  },
  {
    "noun": "Ester",
    "article": "der"
  },
  {
    "noun": "Estragon",
    "article": "der"
  },
  {
    "noun": "Estrich",
    "article": "der"
  },
  {
    "noun": "Estrichleger",
    "article": "der"
  },
  {
    "noun": "Etalon",
    "article": "der"
  },
  {
    "noun": "Etat",
    "article": "der"
  },
  {
    "noun": "Etatentwurf",
    "article": "der"
  },
  {
    "noun": "Ethiker",
    "article": "der"
  },
  {
    "noun": "Ethologe",
    "article": "der"
  },
  {
    "noun": "Etikettenschwindel",
    "article": "der"
  },
  {
    "noun": "Etrusker",
    "article": "der"
  },
  {
    "noun": "Etymologe",
    "article": "der"
  },
  {
    "noun": "Eukalyptus",
    "article": "der"
  },
  {
    "noun": "Eulenschrei",
    "article": "der"
  },
  {
    "noun": "Eulenspiegel",
    "article": "der"
  },
  {
    "noun": "Eunuch",
    "article": "der"
  },
  {
    "noun": "Euphemismus",
    "article": "der"
  },
  {
    "noun": "Euphrat",
    "article": "der"
  },
  {
    "noun": "Eurasier",
    "article": "der"
  },
  {
    "noun": "Euro",
    "article": "der"
  },
  {
    "noun": "Eurocent",
    "article": "der"
  },
  {
    "noun": "Eurocheque",
    "article": "der"
  },
  {
    "noun": "Eurodollar",
    "article": "der"
  },
  {
    "noun": "Euroskeptiker",
    "article": "der"
  },
  {
    "noun": "Evangelist",
    "article": "der"
  },
  {
    "noun": "Evaporator",
    "article": "der"
  },
  {
    "noun": "Evolutionismus",
    "article": "der"
  },
  {
    "noun": "Evolutionist",
    "article": "der"
  },
  {
    "noun": "Ewenke",
    "article": "der"
  },
  {
    "noun": "Examinand",
    "article": "der"
  },
  {
    "noun": "Examinator",
    "article": "der"
  },
  {
    "noun": "Exekutivrat",
    "article": "der"
  },
  {
    "noun": "Exekutor",
    "article": "der"
  },
  {
    "noun": "Exerzierplatz",
    "article": "der"
  },
  {
    "noun": "Exfreund",
    "article": "der"
  },
  {
    "noun": "Exhaustor",
    "article": "der"
  },
  {
    "noun": "Exhibitionist",
    "article": "der"
  },
  {
    "noun": "Existenzkampf",
    "article": "der"
  },
  {
    "noun": "Exitus",
    "article": "der"
  },
  {
    "noun": "Exkaiser",
    "article": "der"
  },
  {
    "noun": "Exklusivbericht",
    "article": "der"
  },
  {
    "noun": "Exklusivvertrag",
    "article": "der"
  },
  {
    "noun": "Exkurs",
    "article": "der"
  },
  {
    "noun": "Exodus",
    "article": "der"
  },
  {
    "noun": "Exophthalmus",
    "article": "der"
  },
  {
    "noun": "Exorzismus",
    "article": "der"
  },
  {
    "noun": "Exorzist",
    "article": "der"
  },
  {
    "noun": "Exot",
    "article": "der"
  },
  {
    "noun": "Exotismus",
    "article": "der"
  },
  {
    "noun": "Expander",
    "article": "der"
  },
  {
    "noun": "Expedient",
    "article": "der"
  },
  {
    "noun": "Experimentator",
    "article": "der"
  },
  {
    "noun": "Experte",
    "article": "der"
  },
  {
    "noun": "Explosionsdruck",
    "article": "der"
  },
  {
    "noun": "Explosionsherd",
    "article": "der"
  },
  {
    "noun": "Explosionskrater",
    "article": "der"
  },
  {
    "noun": "Explosivlaut",
    "article": "der"
  },
  {
    "noun": "Exponent",
    "article": "der"
  },
  {
    "noun": "Export",
    "article": "der"
  },
  {
    "noun": "Exportanteil",
    "article": "der"
  },
  {
    "noun": "Exportartikel",
    "article": "der"
  },
  {
    "noun": "Exportauftrag",
    "article": "der"
  },
  {
    "noun": "Exporteur",
    "article": "der"
  },
  {
    "noun": "Exporthandel",
    "article": "der"
  },
  {
    "noun": "Exportkredit",
    "article": "der"
  },
  {
    "noun": "Exportleiter",
    "article": "der"
  },
  {
    "noun": "Exportmarkt",
    "article": "der"
  },
  {
    "noun": "Exportpreis",
    "article": "der"
  },
  {
    "noun": "Exportzoll",
    "article": "der"
  },
  {
    "noun": "Express",
    "article": "der"
  },
  {
    "noun": "Expressdienst",
    "article": "der"
  },
  {
    "noun": "Expressionismus",
    "article": "der"
  },
  {
    "noun": "Expressionist",
    "article": "der"
  },
  {
    "noun": "Expresszug",
    "article": "der"
  },
  {
    "noun": "Externspeicher",
    "article": "der"
  },
  {
    "noun": "Extinktionskoeffizient",
    "article": "der"
  },
  {
    "noun": "Extrakt",
    "article": "der"
  },
  {
    "noun": "Extraordinarius",
    "article": "der"
  },
  {
    "noun": "Extraprofit",
    "article": "der"
  },
  {
    "noun": "Extremfall",
    "article": "der"
  },
  {
    "noun": "Extremismus",
    "article": "der"
  },
  {
    "noun": "Extremist",
    "article": "der"
  },
  {
    "noun": "Extremsport",
    "article": "der"
  },
  {
    "noun": "Extremwert",
    "article": "der"
  },
  {
    "noun": "Extruder",
    "article": "der"
  },
  {
    "noun": "Exzenter",
    "article": "der"
  },
  {
    "noun": "Exzentriker",
    "article": "der"
  },
  {
    "noun": "Exzess",
    "article": "der"
  },
  {
    "noun": "Eyeliner",
    "article": "der"
  },
  {
    "noun": "Fadenwurm",
    "article": "der"
  },
  {
    "noun": "Fagottist",
    "article": "der"
  },
  {
    "noun": "Fahlsegler",
    "article": "der"
  },
  {
    "noun": "Fahneneid",
    "article": "der"
  },
  {
    "noun": "Fahnenjunker",
    "article": "der"
  },
  {
    "noun": "Fahnenmast",
    "article": "der"
  },
  {
    "noun": "Fahrausweis",
    "article": "der"
  },
  {
    "noun": "Fahrbahnwechsel",
    "article": "der"
  },
  {
    "noun": "Fahrbereich",
    "article": "der"
  },
  {
    "noun": "Fahrbetrieb",
    "article": "der"
  },
  {
    "noun": "Fahrdienst",
    "article": "der"
  },
  {
    "noun": "Fahrdienstleiter",
    "article": "der"
  },
  {
    "noun": "Fahrdraht",
    "article": "der"
  },
  {
    "noun": "Fahrer",
    "article": "der"
  },
  {
    "noun": "Fahrersitz",
    "article": "der"
  },
  {
    "noun": "Fahrgast",
    "article": "der"
  },
  {
    "noun": "Fahrgastraum",
    "article": "der"
  },
  {
    "noun": "Fahrkartenautomat",
    "article": "der"
  },
  {
    "noun": "Fahrkartenschalter",
    "article": "der"
  },
  {
    "noun": "Fahrkomfort",
    "article": "der"
  },
  {
    "noun": "Fahrkorb",
    "article": "der"
  },
  {
    "noun": "Fahrlehrer",
    "article": "der"
  },
  {
    "noun": "Fahrplan",
    "article": "der"
  },
  {
    "noun": "Fahrpreis",
    "article": "der"
  },
  {
    "noun": "Fahrradfahrer",
    "article": "der"
  },
  {
    "noun": "Fahrradhelm",
    "article": "der"
  },
  {
    "noun": "Fahrradkurier",
    "article": "der"
  },
  {
    "noun": "Fahrradreifen",
    "article": "der"
  },
  {
    "noun": "Fahrradsattel",
    "article": "der"
  },
  {
    "noun": "Fahrradsport",
    "article": "der"
  },
  {
    "noun": "Fahrradverleih",
    "article": "der"
  },
  {
    "noun": "Fahrradweg",
    "article": "der"
  },
  {
    "noun": "Fahrschein",
    "article": "der"
  },
  {
    "noun": "Fahrsteig",
    "article": "der"
  },
  {
    "noun": "Fahrsteiger",
    "article": "der"
  },
  {
    "noun": "Fahrstrahl",
    "article": "der"
  },
  {
    "noun": "Fahrstreifen",
    "article": "der"
  },
  {
    "noun": "Fahrstuhl",
    "article": "der"
  },
  {
    "noun": "Fahrstuhlschacht",
    "article": "der"
  },
  {
    "noun": "Fahrtenschreiber",
    "article": "der"
  },
  {
    "noun": "Fahrtest",
    "article": "der"
  },
  {
    "noun": "Fahrtrichtungsanzeiger",
    "article": "der"
  },
  {
    "noun": "Fahrtschreiber",
    "article": "der"
  },
  {
    "noun": "Fahrtwind",
    "article": "der"
  },
  {
    "noun": "Fahrweg",
    "article": "der"
  },
  {
    "noun": "Fahrzeugbau",
    "article": "der"
  },
  {
    "noun": "Fahrzeugbrief",
    "article": "der"
  },
  {
    "noun": "Fahrzeughalter",
    "article": "der"
  },
  {
    "noun": "Fahrzeuglenker",
    "article": "der"
  },
  {
    "noun": "Fahrzeugpark",
    "article": "der"
  },
  {
    "noun": "Fahrzeugrahmen",
    "article": "der"
  },
  {
    "noun": "Fahrzeugverkehr",
    "article": "der"
  },
  {
    "noun": "Fakir",
    "article": "der"
  },
  {
    "noun": "Faksimilestempel",
    "article": "der"
  },
  {
    "noun": "Faktis",
    "article": "der"
  },
  {
    "noun": "Faktor",
    "article": "der"
  },
  {
    "noun": "Fakturist",
    "article": "der"
  },
  {
    "noun": "Falangist",
    "article": "der"
  },
  {
    "noun": "Falke",
    "article": "der"
  },
  {
    "noun": "Falkner",
    "article": "der"
  },
  {
    "noun": "Fall",
    "article": "der"
  },
  {
    "noun": "Fallensteller",
    "article": "der"
  },
  {
    "noun": "Fallhammer",
    "article": "der"
  },
  {
    "noun": "Fallschirm",
    "article": "der"
  },
  {
    "noun": "Fallschirmabsprung",
    "article": "der"
  },
  {
    "noun": "Fallschirmspringer",
    "article": "der"
  },
  {
    "noun": "Fallstrick",
    "article": "der"
  },
  {
    "noun": "Fallwind",
    "article": "der"
  },
  {
    "noun": "Falscheid",
    "article": "der"
  },
  {
    "noun": "Falschspieler",
    "article": "der"
  },
  {
    "noun": "Falsettist",
    "article": "der"
  },
  {
    "noun": "Faltenbalg",
    "article": "der"
  },
  {
    "noun": "Faltenrock",
    "article": "der"
  },
  {
    "noun": "Falter",
    "article": "der"
  },
  {
    "noun": "Faltprospekt",
    "article": "der"
  },
  {
    "noun": "Faltstuhl",
    "article": "der"
  },
  {
    "noun": "Falz",
    "article": "der"
  },
  {
    "noun": "Falzziegel",
    "article": "der"
  },
  {
    "noun": "Familienbesitz",
    "article": "der"
  },
  {
    "noun": "Familienbetrieb",
    "article": "der"
  },
  {
    "noun": "Familienkreis",
    "article": "der"
  },
  {
    "noun": "Familienname",
    "article": "der"
  },
  {
    "noun": "Familienpass",
    "article": "der"
  },
  {
    "noun": "Familienrat",
    "article": "der"
  },
  {
    "noun": "Familienschmuck",
    "article": "der"
  },
  {
    "noun": "Familienstand",
    "article": "der"
  },
  {
    "noun": "Familienvater",
    "article": "der"
  },
  {
    "noun": "Familienvorstand",
    "article": "der"
  },
  {
    "noun": "Familienzuschlag",
    "article": "der"
  },
  {
    "noun": "Familienzuwachs",
    "article": "der"
  },
  {
    "noun": "Famulus",
    "article": "der"
  },
  {
    "noun": "Fan",
    "article": "der"
  },
  {
    "noun": "Fanatiker",
    "article": "der"
  },
  {
    "noun": "Fanatismus",
    "article": "der"
  },
  {
    "noun": "Fanclub",
    "article": "der"
  },
  {
    "noun": "Fang",
    "article": "der"
  },
  {
    "noun": "Fangarm",
    "article": "der"
  },
  {
    "noun": "Fangdamm",
    "article": "der"
  },
  {
    "noun": "Fangfehler",
    "article": "der"
  },
  {
    "noun": "Fanghaken",
    "article": "der"
  },
  {
    "noun": "Fanghandschuh",
    "article": "der"
  },
  {
    "noun": "Fangzahn",
    "article": "der"
  },
  {
    "noun": "Fant",
    "article": "der"
  },
  {
    "noun": "Fantasiepreis",
    "article": "der"
  },
  {
    "noun": "Fantast",
    "article": "der"
  },
  {
    "noun": "Fantasyroman",
    "article": "der"
  },
  {
    "noun": "Farbauszug",
    "article": "der"
  },
  {
    "noun": "Farbdrucker",
    "article": "der"
  },
  {
    "noun": "Farbendruck",
    "article": "der"
  },
  {
    "noun": "Farbentest",
    "article": "der"
  },
  {
    "noun": "Farbfernseher",
    "article": "der"
  },
  {
    "noun": "Farbfilm",
    "article": "der"
  },
  {
    "noun": "Farbkasten",
    "article": "der"
  },
  {
    "noun": "Farbkontrast",
    "article": "der"
  },
  {
    "noun": "Farbmonitor",
    "article": "der"
  },
  {
    "noun": "Farbstift",
    "article": "der"
  },
  {
    "noun": "Farbstoff",
    "article": "der"
  },
  {
    "noun": "Farbtest",
    "article": "der"
  },
  {
    "noun": "Farbton",
    "article": "der"
  },
  {
    "noun": "Farbtopf",
    "article": "der"
  },
  {
    "noun": "Farbtupfer",
    "article": "der"
  },
  {
    "noun": "Farbwechsel",
    "article": "der"
  },
  {
    "noun": "Farmer",
    "article": "der"
  },
  {
    "noun": "Farn",
    "article": "der"
  },
  {
    "noun": "Farnwedel",
    "article": "der"
  },
  {
    "noun": "Fasan",
    "article": "der"
  },
  {
    "noun": "Fasching",
    "article": "der"
  },
  {
    "noun": "Faschingsdienstag",
    "article": "der"
  },
  {
    "noun": "Faschingskrapfen",
    "article": "der"
  },
  {
    "noun": "Faschingsprinz",
    "article": "der"
  },
  {
    "noun": "Faschingsumzug",
    "article": "der"
  },
  {
    "noun": "Faschingszug",
    "article": "der"
  },
  {
    "noun": "Faschismus",
    "article": "der"
  },
  {
    "noun": "Faschist",
    "article": "der"
  },
  {
    "noun": "Fasel",
    "article": "der"
  },
  {
    "noun": "Faserschreiber",
    "article": "der"
  },
  {
    "noun": "Fassadenkletterer",
    "article": "der"
  },
  {
    "noun": "Fassbinder",
    "article": "der"
  },
  {
    "noun": "Fasshahn",
    "article": "der"
  },
  {
    "noun": "Fassreifen",
    "article": "der"
  },
  {
    "noun": "Fasswein",
    "article": "der"
  },
  {
    "noun": "Fastnachtsbrauch",
    "article": "der"
  },
  {
    "noun": "Fastnachtsdienstag",
    "article": "der"
  },
  {
    "noun": "Fasttag",
    "article": "der"
  },
  {
    "noun": "Faszikel",
    "article": "der"
  },
  {
    "noun": "Fatalismus",
    "article": "der"
  },
  {
    "noun": "Fatalist",
    "article": "der"
  },
  {
    "noun": "Fatzke",
    "article": "der"
  },
  {
    "noun": "Faulbaum",
    "article": "der"
  },
  {
    "noun": "Faulenzer",
    "article": "der"
  },
  {
    "noun": "Faulpelz",
    "article": "der"
  },
  {
    "noun": "Faulschlamm",
    "article": "der"
  },
  {
    "noun": "Faun",
    "article": "der"
  },
  {
    "noun": "Faustball",
    "article": "der"
  },
  {
    "noun": "Fausthandschuh",
    "article": "der"
  },
  {
    "noun": "Faustkampf",
    "article": "der"
  },
  {
    "noun": "Faustkeil",
    "article": "der"
  },
  {
    "noun": "Faustschlag",
    "article": "der"
  },
  {
    "noun": "Fauvismus",
    "article": "der"
  },
  {
    "noun": "Fauvist",
    "article": "der"
  },
  {
    "noun": "Fauxpas",
    "article": "der"
  },
  {
    "noun": "Favorit",
    "article": "der"
  },
  {
    "noun": "Faxenmacher",
    "article": "der"
  },
  {
    "noun": "Fazialis",
    "article": "der"
  },
  {
    "noun": "Februar",
    "article": "der"
  },
  {
    "noun": "Fechtboden",
    "article": "der"
  },
  {
    "noun": "Fechtbruder",
    "article": "der"
  },
  {
    "noun": "Fechtdegen",
    "article": "der"
  },
  {
    "noun": "Fechter",
    "article": "der"
  },
  {
    "noun": "Fechtmeister",
    "article": "der"
  },
  {
    "noun": "Federball",
    "article": "der"
  },
  {
    "noun": "Federbesen",
    "article": "der"
  },
  {
    "noun": "Federbusch",
    "article": "der"
  },
  {
    "noun": "Federfuchser",
    "article": "der"
  },
  {
    "noun": "Federhalter",
    "article": "der"
  },
  {
    "noun": "Federkiel",
    "article": "der"
  },
  {
    "noun": "Federkrieg",
    "article": "der"
  },
  {
    "noun": "Federstahl",
    "article": "der"
  },
  {
    "noun": "Federstrich",
    "article": "der"
  },
  {
    "noun": "Federzirkel",
    "article": "der"
  },
  {
    "noun": "Federzug",
    "article": "der"
  },
  {
    "noun": "Feeder",
    "article": "der"
  },
  {
    "noun": "Feger",
    "article": "der"
  },
  {
    "noun": "Fehlbestand",
    "article": "der"
  },
  {
    "noun": "Fehlbetrag",
    "article": "der"
  },
  {
    "noun": "Fehldruck",
    "article": "der"
  },
  {
    "noun": "Fehler",
    "article": "der"
  },
  {
    "noun": "Fehlgriff",
    "article": "der"
  },
  {
    "noun": "Fehlpass",
    "article": "der"
  },
  {
    "noun": "Fehlschlag",
    "article": "der"
  },
  {
    "noun": "Fehlschluss",
    "article": "der"
  },
  {
    "noun": "Fehlschuss",
    "article": "der"
  },
  {
    "noun": "Fehlstart",
    "article": "der"
  },
  {
    "noun": "Fehltritt",
    "article": "der"
  },
  {
    "noun": "Fehlversuch",
    "article": "der"
  },
  {
    "noun": "Fehlwurf",
    "article": "der"
  },
  {
    "noun": "Feierabend",
    "article": "der"
  },
  {
    "noun": "Feiertag",
    "article": "der"
  },
  {
    "noun": "Feigenbaum",
    "article": "der"
  },
  {
    "noun": "Feigling",
    "article": "der"
  },
  {
    "noun": "Feilenhauer",
    "article": "der"
  },
  {
    "noun": "Feilkloben",
    "article": "der"
  },
  {
    "noun": "Feilspan",
    "article": "der"
  },
  {
    "noun": "Feingehalt",
    "article": "der"
  },
  {
    "noun": "Feingehaltsstempel",
    "article": "der"
  },
  {
    "noun": "Feinguss",
    "article": "der"
  },
  {
    "noun": "Feinkostladen",
    "article": "der"
  },
  {
    "noun": "Feinmechaniker",
    "article": "der"
  },
  {
    "noun": "Feinschmecker",
    "article": "der"
  },
  {
    "noun": "Feinstaub",
    "article": "der"
  },
  {
    "noun": "Felchen",
    "article": "der"
  },
  {
    "noun": "Feldahorn",
    "article": "der"
  },
  {
    "noun": "Feldarbeiter",
    "article": "der"
  },
  {
    "noun": "Feldbau",
    "article": "der"
  },
  {
    "noun": "Felddienst",
    "article": "der"
  },
  {
    "noun": "Feldhandball",
    "article": "der"
  },
  {
    "noun": "Feldhase",
    "article": "der"
  },
  {
    "noun": "Feldherr",
    "article": "der"
  },
  {
    "noun": "Feldmarschall",
    "article": "der"
  },
  {
    "noun": "Feldmesser",
    "article": "der"
  },
  {
    "noun": "Feldrain",
    "article": "der"
  },
  {
    "noun": "Feldsalat",
    "article": "der"
  },
  {
    "noun": "Feldspat",
    "article": "der"
  },
  {
    "noun": "Feldsperling",
    "article": "der"
  },
  {
    "noun": "Feldstecher",
    "article": "der"
  },
  {
    "noun": "Feldstein",
    "article": "der"
  },
  {
    "noun": "Feldversuch",
    "article": "der"
  },
  {
    "noun": "Feldwebel",
    "article": "der"
  },
  {
    "noun": "Feldweg",
    "article": "der"
  },
  {
    "noun": "Feldweibel",
    "article": "der"
  },
  {
    "noun": "Feldzug",
    "article": "der"
  },
  {
    "noun": "Felgumschwung",
    "article": "der"
  },
  {
    "noun": "Fels",
    "article": "der"
  },
  {
    "noun": "Felsblock",
    "article": "der"
  },
  {
    "noun": "Felsbrocken",
    "article": "der"
  },
  {
    "noun": "Felsendom",
    "article": "der"
  },
  {
    "noun": "Felsgrat",
    "article": "der"
  },
  {
    "noun": "Felsgrund",
    "article": "der"
  },
  {
    "noun": "Felsit",
    "article": "der"
  },
  {
    "noun": "Felsspalt",
    "article": "der"
  },
  {
    "noun": "Felsvorsprung",
    "article": "der"
  },
  {
    "noun": "Feminismus",
    "article": "der"
  },
  {
    "noun": "Feminist",
    "article": "der"
  },
  {
    "noun": "Fench",
    "article": "der"
  },
  {
    "noun": "Fenchel",
    "article": "der"
  },
  {
    "noun": "Fencheltee",
    "article": "der"
  },
  {
    "noun": "Fender",
    "article": "der"
  },
  {
    "noun": "Fenek",
    "article": "der"
  },
  {
    "noun": "Fennich",
    "article": "der"
  },
  {
    "noun": "Fensterbriefumschlag",
    "article": "der"
  },
  {
    "noun": "Fenstergriff",
    "article": "der"
  },
  {
    "noun": "Fensterkitt",
    "article": "der"
  },
  {
    "noun": "Fensterladen",
    "article": "der"
  },
  {
    "noun": "Fensterplatz",
    "article": "der"
  },
  {
    "noun": "Fensterputzer",
    "article": "der"
  },
  {
    "noun": "Fensterrahmen",
    "article": "der"
  },
  {
    "noun": "Fensterstock",
    "article": "der"
  },
  {
    "noun": "Fenstersturz",
    "article": "der"
  },
  {
    "noun": "Ferienjob",
    "article": "der"
  },
  {
    "noun": "Ferienort",
    "article": "der"
  },
  {
    "noun": "Ferientag",
    "article": "der"
  },
  {
    "noun": "Fermenter",
    "article": "der"
  },
  {
    "noun": "Fernempfang",
    "article": "der"
  },
  {
    "noun": "Fernfahrer",
    "article": "der"
  },
  {
    "noun": "Fernkopierer",
    "article": "der"
  },
  {
    "noun": "Fernkurs",
    "article": "der"
  },
  {
    "noun": "Fernlehrgang",
    "article": "der"
  },
  {
    "noun": "Fernleihverkehr",
    "article": "der"
  },
  {
    "noun": "Fernmeldedienst",
    "article": "der"
  },
  {
    "noun": "Fernmeldesatellit",
    "article": "der"
  },
  {
    "noun": "Fernmeldeverkehr",
    "article": "der"
  },
  {
    "noun": "Fernost",
    "article": "der"
  },
  {
    "noun": "Fernruf",
    "article": "der"
  },
  {
    "noun": "Fernschreiber",
    "article": "der"
  },
  {
    "noun": "Fernsehapparat",
    "article": "der"
  },
  {
    "noun": "Fernseher",
    "article": "der"
  },
  {
    "noun": "Fernsehfilm",
    "article": "der"
  },
  {
    "noun": "Fernsehjournalist",
    "article": "der"
  },
  {
    "noun": "Fernsehkanal",
    "article": "der"
  },
  {
    "noun": "Fernsehsatellit",
    "article": "der"
  },
  {
    "noun": "Fernsehschirm",
    "article": "der"
  },
  {
    "noun": "Fernsehsender",
    "article": "der"
  },
  {
    "noun": "Fernsehspot",
    "article": "der"
  },
  {
    "noun": "Fernsehtechniker",
    "article": "der"
  },
  {
    "noun": "Fernsehteilnehmer",
    "article": "der"
  },
  {
    "noun": "Fernsehturm",
    "article": "der"
  },
  {
    "noun": "Fernsehzuschauer",
    "article": "der"
  },
  {
    "noun": "Fernsprechanschluss",
    "article": "der"
  },
  {
    "noun": "Fernsprechapparat",
    "article": "der"
  },
  {
    "noun": "Fernsprechauftragsdienst",
    "article": "der"
  },
  {
    "noun": "Fernsprechautomat",
    "article": "der"
  },
  {
    "noun": "Fernsprecher",
    "article": "der"
  },
  {
    "noun": "Fernsprechteilnehmer",
    "article": "der"
  },
  {
    "noun": "Ferntransport",
    "article": "der"
  },
  {
    "noun": "Fernunterricht",
    "article": "der"
  },
  {
    "noun": "Fernverkehr",
    "article": "der"
  },
  {
    "noun": "Ferrimagnetismus",
    "article": "der"
  },
  {
    "noun": "Ferrit",
    "article": "der"
  },
  {
    "noun": "Ferritkern",
    "article": "der"
  },
  {
    "noun": "Ferrograf",
    "article": "der"
  },
  {
    "noun": "Ferromagnetismus",
    "article": "der"
  },
  {
    "noun": "Fertigbau",
    "article": "der"
  },
  {
    "noun": "Fertigbeton",
    "article": "der"
  },
  {
    "noun": "Fertigungsablauf",
    "article": "der"
  },
  {
    "noun": "Fertigungsprozess",
    "article": "der"
  },
  {
    "noun": "Fesselballon",
    "article": "der"
  },
  {
    "noun": "Festabend",
    "article": "der"
  },
  {
    "noun": "Festakt",
    "article": "der"
  },
  {
    "noun": "Festbetrag",
    "article": "der"
  },
  {
    "noun": "Festordner",
    "article": "der"
  },
  {
    "noun": "Festplatz",
    "article": "der"
  },
  {
    "noun": "Festpreis",
    "article": "der"
  },
  {
    "noun": "Festpunkt",
    "article": "der"
  },
  {
    "noun": "Festredner",
    "article": "der"
  },
  {
    "noun": "Festsaal",
    "article": "der"
  },
  {
    "noun": "Feststoff",
    "article": "der"
  },
  {
    "noun": "Festtag",
    "article": "der"
  },
  {
    "noun": "Festumzug",
    "article": "der"
  },
  {
    "noun": "Festungsgraben",
    "article": "der"
  },
  {
    "noun": "Festwertspeicher",
    "article": "der"
  },
  {
    "noun": "Festzug",
    "article": "der"
  },
  {
    "noun": "Fetisch",
    "article": "der"
  },
  {
    "noun": "Fetischismus",
    "article": "der"
  },
  {
    "noun": "Fetischist",
    "article": "der"
  },
  {
    "noun": "Fettabscheider",
    "article": "der"
  },
  {
    "noun": "Fettbauch",
    "article": "der"
  },
  {
    "noun": "Fettdruck",
    "article": "der"
  },
  {
    "noun": "Fettfilm",
    "article": "der"
  },
  {
    "noun": "Fettfleck",
    "article": "der"
  },
  {
    "noun": "Fettgehalt",
    "article": "der"
  },
  {
    "noun": "Fettglanz",
    "article": "der"
  },
  {
    "noun": "Fettsack",
    "article": "der"
  },
  {
    "noun": "Fettwanst",
    "article": "der"
  },
  {
    "noun": "Fetus",
    "article": "der"
  },
  {
    "noun": "Feuchtigkeitsanzeiger",
    "article": "der"
  },
  {
    "noun": "Feuchtigkeitsgehalt",
    "article": "der"
  },
  {
    "noun": "Feuchtigkeitsgrad",
    "article": "der"
  },
  {
    "noun": "Feuchtigkeitsmesser",
    "article": "der"
  },
  {
    "noun": "Feuchtigkeitsregler",
    "article": "der"
  },
  {
    "noun": "Feuchtigkeitsschutz",
    "article": "der"
  },
  {
    "noun": "Feudalherr",
    "article": "der"
  },
  {
    "noun": "Feudalismus",
    "article": "der"
  },
  {
    "noun": "Feueralarm",
    "article": "der"
  },
  {
    "noun": "Feuerball",
    "article": "der"
  },
  {
    "noun": "Feuerbock",
    "article": "der"
  },
  {
    "noun": "Feuereifer",
    "article": "der"
  },
  {
    "noun": "Feuerkopf",
    "article": "der"
  },
  {
    "noun": "Feuermelder",
    "article": "der"
  },
  {
    "noun": "Feueropal",
    "article": "der"
  },
  {
    "noun": "Feuerraum",
    "article": "der"
  },
  {
    "noun": "Feuerrost",
    "article": "der"
  },
  {
    "noun": "Feuersalamander",
    "article": "der"
  },
  {
    "noun": "Feuerschaden",
    "article": "der"
  },
  {
    "noun": "Feuerschutz",
    "article": "der"
  },
  {
    "noun": "Feuerstein",
    "article": "der"
  },
  {
    "noun": "Feuersturm",
    "article": "der"
  },
  {
    "noun": "Feuerteufel",
    "article": "der"
  },
  {
    "noun": "Feuerton",
    "article": "der"
  },
  {
    "noun": "Feuerwehrhauptmann",
    "article": "der"
  },
  {
    "noun": "Feuerwehrmann",
    "article": "der"
  },
  {
    "noun": "Feuerwehrschlauch",
    "article": "der"
  },
  {
    "noun": "Feuerwerker",
    "article": "der"
  },
  {
    "noun": "Feuilletonist",
    "article": "der"
  },
  {
    "noun": "Feuilletonstil",
    "article": "der"
  },
  {
    "noun": "Feuilletonteil",
    "article": "der"
  },
  {
    "noun": "Fiaker",
    "article": "der"
  },
  {
    "noun": "Fichtenbaum",
    "article": "der"
  },
  {
    "noun": "Fichtenkreuzschnabel",
    "article": "der"
  },
  {
    "noun": "Fichtenwald",
    "article": "der"
  },
  {
    "noun": "Fichtenzapfen",
    "article": "der"
  },
  {
    "noun": "Fieberanfall",
    "article": "der"
  },
  {
    "noun": "Fiebermesser",
    "article": "der"
  },
  {
    "noun": "Fieberschauer",
    "article": "der"
  },
  {
    "noun": "Fiebertraum",
    "article": "der"
  },
  {
    "noun": "Fieberwahn",
    "article": "der"
  },
  {
    "noun": "Fiedelbogen",
    "article": "der"
  },
  {
    "noun": "Fieldistor",
    "article": "der"
  },
  {
    "noun": "Fiesling",
    "article": "der"
  },
  {
    "noun": "Figaro",
    "article": "der"
  },
  {
    "noun": "Figurant",
    "article": "der"
  },
  {
    "noun": "Filetbraten",
    "article": "der"
  },
  {
    "noun": "Filialbetrieb",
    "article": "der"
  },
  {
    "noun": "Filialleiter",
    "article": "der"
  },
  {
    "noun": "Filipino",
    "article": "der"
  },
  {
    "noun": "Film",
    "article": "der"
  },
  {
    "noun": "Filmdarsteller",
    "article": "der"
  },
  {
    "noun": "Filmemacher",
    "article": "der"
  },
  {
    "noun": "Filmpreis",
    "article": "der"
  },
  {
    "noun": "Filmproduzent",
    "article": "der"
  },
  {
    "noun": "Filmprojektor",
    "article": "der"
  },
  {
    "noun": "Filmregisseur",
    "article": "der"
  },
  {
    "noun": "Filmschauspieler",
    "article": "der"
  },
  {
    "noun": "Filmstar",
    "article": "der"
  },
  {
    "noun": "Filmstreifen",
    "article": "der"
  },
  {
    "noun": "Filmtransport",
    "article": "der"
  },
  {
    "noun": "Filmverleih",
    "article": "der"
  },
  {
    "noun": "Filterkaffee",
    "article": "der"
  },
  {
    "noun": "Filterkuchen",
    "article": "der"
  },
  {
    "noun": "Filz",
    "article": "der"
  },
  {
    "noun": "Filzhut",
    "article": "der"
  },
  {
    "noun": "Filzpantoffel",
    "article": "der"
  },
  {
    "noun": "Filzschreiber",
    "article": "der"
  },
  {
    "noun": "Filzstiefel",
    "article": "der"
  },
  {
    "noun": "Filzstift",
    "article": "der"
  },
  {
    "noun": "Fimmel",
    "article": "der"
  },
  {
    "noun": "Finalgegner",
    "article": "der"
  },
  {
    "noun": "Finalist",
    "article": "der"
  },
  {
    "noun": "Finalsatz",
    "article": "der"
  },
  {
    "noun": "Finanzausgleich",
    "article": "der"
  },
  {
    "noun": "Finanzausschuss",
    "article": "der"
  },
  {
    "noun": "Finanzbedarf",
    "article": "der"
  },
  {
    "noun": "Finanzberater",
    "article": "der"
  },
  {
    "noun": "Finanzbuchhalter",
    "article": "der"
  },
  {
    "noun": "Finanzdirektor",
    "article": "der"
  },
  {
    "noun": "Finanzexperte",
    "article": "der"
  },
  {
    "noun": "Finanzier",
    "article": "der"
  },
  {
    "noun": "Finanzierungsplan",
    "article": "der"
  },
  {
    "noun": "Finanzmagnat",
    "article": "der"
  },
  {
    "noun": "Finanzmakler",
    "article": "der"
  },
  {
    "noun": "Finanzmann",
    "article": "der"
  },
  {
    "noun": "Finanzmarkt",
    "article": "der"
  },
  {
    "noun": "Finanzminister",
    "article": "der"
  },
  {
    "noun": "Finanzplan",
    "article": "der"
  },
  {
    "noun": "Finanzplatz",
    "article": "der"
  },
  {
    "noun": "Finanzwechsel",
    "article": "der"
  },
  {
    "noun": "Finanzzoll",
    "article": "der"
  },
  {
    "noun": "Finder",
    "article": "der"
  },
  {
    "noun": "Finderlohn",
    "article": "der"
  },
  {
    "noun": "Findling",
    "article": "der"
  },
  {
    "noun": "Finger",
    "article": "der"
  },
  {
    "noun": "Fingerabdruck",
    "article": "der"
  },
  {
    "noun": "Fingerhut",
    "article": "der"
  },
  {
    "noun": "Fingerknochen",
    "article": "der"
  },
  {
    "noun": "Fingerling",
    "article": "der"
  },
  {
    "noun": "Fingernagel",
    "article": "der"
  },
  {
    "noun": "Fingerring",
    "article": "der"
  },
  {
    "noun": "Fingersatz",
    "article": "der"
  },
  {
    "noun": "Fingerzeig",
    "article": "der"
  },
  {
    "noun": "Fink",
    "article": "der"
  },
  {
    "noun": "Finkenschlag",
    "article": "der"
  },
  {
    "noun": "Finne",
    "article": "der"
  },
  {
    "noun": "Finnwal",
    "article": "der"
  },
  {
    "noun": "Fips",
    "article": "der"
  },
  {
    "noun": "Firlefanz",
    "article": "der"
  },
  {
    "noun": "Firmenchef",
    "article": "der"
  },
  {
    "noun": "Firmeninhaber",
    "article": "der"
  },
  {
    "noun": "Firmenmantel",
    "article": "der"
  },
  {
    "noun": "Firmenname",
    "article": "der"
  },
  {
    "noun": "Firmensitz",
    "article": "der"
  },
  {
    "noun": "Firmensprecher",
    "article": "der"
  },
  {
    "noun": "Firmenstempel",
    "article": "der"
  },
  {
    "noun": "Firmenwagen",
    "article": "der"
  },
  {
    "noun": "Firmenwert",
    "article": "der"
  },
  {
    "noun": "Firn",
    "article": "der"
  },
  {
    "noun": "Firnewein",
    "article": "der"
  },
  {
    "noun": "Firnis",
    "article": "der"
  },
  {
    "noun": "Firnschnee",
    "article": "der"
  },
  {
    "noun": "First",
    "article": "der"
  },
  {
    "noun": "Firstbalken",
    "article": "der"
  },
  {
    "noun": "Firstziegel",
    "article": "der"
  },
  {
    "noun": "Fisch",
    "article": "der"
  },
  {
    "noun": "Fischadler",
    "article": "der"
  },
  {
    "noun": "Fischbestand",
    "article": "der"
  },
  {
    "noun": "Fischdampfer",
    "article": "der"
  },
  {
    "noun": "Fischer",
    "article": "der"
  },
  {
    "noun": "Fischereihafen",
    "article": "der"
  },
  {
    "noun": "Fischerkahn",
    "article": "der"
  },
  {
    "noun": "Fischfang",
    "article": "der"
  },
  {
    "noun": "Fischgeruch",
    "article": "der"
  },
  {
    "noun": "Fischhandel",
    "article": "der"
  },
  {
    "noun": "Fischkutter",
    "article": "der"
  },
  {
    "noun": "Fischlaich",
    "article": "der"
  },
  {
    "noun": "Fischleim",
    "article": "der"
  },
  {
    "noun": "Fischmarkt",
    "article": "der"
  },
  {
    "noun": "Fischotter",
    "article": "der"
  },
  {
    "noun": "Fischpass",
    "article": "der"
  },
  {
    "noun": "Fischreichtum",
    "article": "der"
  },
  {
    "noun": "Fischreiher",
    "article": "der"
  },
  {
    "noun": "Fischrogen",
    "article": "der"
  },
  {
    "noun": "Fischsaurier",
    "article": "der"
  },
  {
    "noun": "Fischschwanz",
    "article": "der"
  },
  {
    "noun": "Fischschwarm",
    "article": "der"
  },
  {
    "noun": "Fischteich",
    "article": "der"
  },
  {
    "noun": "Fischtran",
    "article": "der"
  },
  {
    "noun": "Fischweg",
    "article": "der"
  },
  {
    "noun": "Fischweiher",
    "article": "der"
  },
  {
    "noun": "Fischzug",
    "article": "der"
  },
  {
    "noun": "Fiskus",
    "article": "der"
  },
  {
    "noun": "Fitis",
    "article": "der"
  },
  {
    "noun": "Fitnessraum",
    "article": "der"
  },
  {
    "noun": "Fittich",
    "article": "der"
  },
  {
    "noun": "Fixateur",
    "article": "der"
  },
  {
    "noun": "Fixer",
    "article": "der"
  },
  {
    "noun": "Fixpunkt",
    "article": "der"
  },
  {
    "noun": "Fixstern",
    "article": "der"
  },
  {
    "noun": "Fjell",
    "article": "der"
  },
  {
    "noun": "Fjord",
    "article": "der"
  },
  {
    "noun": "Flachbettscanner",
    "article": "der"
  },
  {
    "noun": "Flachbildschirm",
    "article": "der"
  },
  {
    "noun": "Flachdruck",
    "article": "der"
  },
  {
    "noun": "Flachkopf",
    "article": "der"
  },
  {
    "noun": "Flachmann",
    "article": "der"
  },
  {
    "noun": "Flachsbau",
    "article": "der"
  },
  {
    "noun": "Flachssame",
    "article": "der"
  },
  {
    "noun": "Flachstahl",
    "article": "der"
  },
  {
    "noun": "Flachwichser",
    "article": "der"
  },
  {
    "noun": "Flachziegel",
    "article": "der"
  },
  {
    "noun": "Fladen",
    "article": "der"
  },
  {
    "noun": "Flagellant",
    "article": "der"
  },
  {
    "noun": "Flaggenmast",
    "article": "der"
  },
  {
    "noun": "Flaggoffizier",
    "article": "der"
  },
  {
    "noun": "Flame",
    "article": "der"
  },
  {
    "noun": "Flamenco",
    "article": "der"
  },
  {
    "noun": "Flamingo",
    "article": "der"
  },
  {
    "noun": "Flammenbogen",
    "article": "der"
  },
  {
    "noun": "Flammenwerfer",
    "article": "der"
  },
  {
    "noun": "Flammeri",
    "article": "der"
  },
  {
    "noun": "Flammpunkt",
    "article": "der"
  },
  {
    "noun": "Flanell",
    "article": "der"
  },
  {
    "noun": "Flaneur",
    "article": "der"
  },
  {
    "noun": "Flankenangriff",
    "article": "der"
  },
  {
    "noun": "Flankenball",
    "article": "der"
  },
  {
    "noun": "Flankenlauf",
    "article": "der"
  },
  {
    "noun": "Flankenschutz",
    "article": "der"
  },
  {
    "noun": "Flansch",
    "article": "der"
  },
  {
    "noun": "Flaschenhals",
    "article": "der"
  },
  {
    "noun": "Flaschenwein",
    "article": "der"
  },
  {
    "noun": "Flaschenzug",
    "article": "der"
  },
  {
    "noun": "Flattergeist",
    "article": "der"
  },
  {
    "noun": "Flattersatz",
    "article": "der"
  },
  {
    "noun": "Flaum",
    "article": "der"
  },
  {
    "noun": "Flaumacher",
    "article": "der"
  },
  {
    "noun": "Flaumbart",
    "article": "der"
  },
  {
    "noun": "Flausch",
    "article": "der"
  },
  {
    "noun": "Flausenmacher",
    "article": "der"
  },
  {
    "noun": "Flechter",
    "article": "der"
  },
  {
    "noun": "Fleck",
    "article": "der"
  },
  {
    "noun": "Fleckentferner",
    "article": "der"
  },
  {
    "noun": "Flecktyphus",
    "article": "der"
  },
  {
    "noun": "Flederwisch",
    "article": "der"
  },
  {
    "noun": "Flegel",
    "article": "der"
  },
  {
    "noun": "Fleischbeschauer",
    "article": "der"
  },
  {
    "noun": "Fleischer",
    "article": "der"
  },
  {
    "noun": "Fleischerhaken",
    "article": "der"
  },
  {
    "noun": "Fleischesser",
    "article": "der"
  },
  {
    "noun": "Fleischextrakt",
    "article": "der"
  },
  {
    "noun": "Fleischfresser",
    "article": "der"
  },
  {
    "noun": "Fleischhacker",
    "article": "der"
  },
  {
    "noun": "Fleischhauer",
    "article": "der"
  },
  {
    "noun": "Fleischmarkt",
    "article": "der"
  },
  {
    "noun": "Fleischsaft",
    "article": "der"
  },
  {
    "noun": "Fleischsalat",
    "article": "der"
  },
  {
    "noun": "Fleischwolf",
    "article": "der"
  },
  {
    "noun": "Flexodruck",
    "article": "der"
  },
  {
    "noun": "Flickenteppich",
    "article": "der"
  },
  {
    "noun": "Flicker",
    "article": "der"
  },
  {
    "noun": "Flickschuster",
    "article": "der"
  },
  {
    "noun": "Flieder",
    "article": "der"
  },
  {
    "noun": "Fliegenpilz",
    "article": "der"
  },
  {
    "noun": "Fliegenschrank",
    "article": "der"
  },
  {
    "noun": "Fliegenschwarm",
    "article": "der"
  },
  {
    "noun": "Flieger",
    "article": "der"
  },
  {
    "noun": "Fliegeralarm",
    "article": "der"
  },
  {
    "noun": "Fliegerangriff",
    "article": "der"
  },
  {
    "noun": "Fliegerhorst",
    "article": "der"
  },
  {
    "noun": "Fliegeroffizier",
    "article": "der"
  },
  {
    "noun": "Fliesenbelag",
    "article": "der"
  },
  {
    "noun": "Fliesenleger",
    "article": "der"
  },
  {
    "noun": "Flimmer",
    "article": "der"
  },
  {
    "noun": "Flint",
    "article": "der"
  },
  {
    "noun": "Flintenschrot",
    "article": "der"
  },
  {
    "noun": "Flintenschuss",
    "article": "der"
  },
  {
    "noun": "Flipper",
    "article": "der"
  },
  {
    "noun": "Flirt",
    "article": "der"
  },
  {
    "noun": "Flirter",
    "article": "der"
  },
  {
    "noun": "Flitzer",
    "article": "der"
  },
  {
    "noun": "Floh",
    "article": "der"
  },
  {
    "noun": "Flohbiss",
    "article": "der"
  },
  {
    "noun": "Flohcircus",
    "article": "der"
  },
  {
    "noun": "Flohmarkt",
    "article": "der"
  },
  {
    "noun": "Flop",
    "article": "der"
  },
  {
    "noun": "Florettfechter",
    "article": "der"
  },
  {
    "noun": "Florist",
    "article": "der"
  },
  {
    "noun": "Flottillenadmiral",
    "article": "der"
  },
  {
    "noun": "Fluch",
    "article": "der"
  },
  {
    "noun": "Flucher",
    "article": "der"
  },
  {
    "noun": "Fluchthelfer",
    "article": "der"
  },
  {
    "noun": "Fluchtplan",
    "article": "der"
  },
  {
    "noun": "Fluchtpunkt",
    "article": "der"
  },
  {
    "noun": "Fluchtversuch",
    "article": "der"
  },
  {
    "noun": "Fluchtweg",
    "article": "der"
  },
  {
    "noun": "Flug",
    "article": "der"
  },
  {
    "noun": "Flugball",
    "article": "der"
  },
  {
    "noun": "Flugbegleiter",
    "article": "der"
  },
  {
    "noun": "Flugbetrieb",
    "article": "der"
  },
  {
    "noun": "Flugdienst",
    "article": "der"
  },
  {
    "noun": "Fluggast",
    "article": "der"
  },
  {
    "noun": "Flughafen",
    "article": "der"
  },
  {
    "noun": "Flughafer",
    "article": "der"
  },
  {
    "noun": "Flughund",
    "article": "der"
  },
  {
    "noun": "Fluglehrer",
    "article": "der"
  },
  {
    "noun": "Fluglotse",
    "article": "der"
  },
  {
    "noun": "Flugplan",
    "article": "der"
  },
  {
    "noun": "Flugplatz",
    "article": "der"
  },
  {
    "noun": "Flugpreis",
    "article": "der"
  },
  {
    "noun": "Flugsand",
    "article": "der"
  },
  {
    "noun": "Flugsaurier",
    "article": "der"
  },
  {
    "noun": "Flugschein",
    "article": "der"
  },
  {
    "noun": "Flugsimulator",
    "article": "der"
  },
  {
    "noun": "Flugsport",
    "article": "der"
  },
  {
    "noun": "Flugsteig",
    "article": "der"
  },
  {
    "noun": "Flugverkehr",
    "article": "der"
  },
  {
    "noun": "Flugweg",
    "article": "der"
  },
  {
    "noun": "Flugwetterdienst",
    "article": "der"
  },
  {
    "noun": "Flugzeugabsturz",
    "article": "der"
  },
  {
    "noun": "Flugzeugbau",
    "article": "der"
  },
  {
    "noun": "Flugzeugkonstrukteur",
    "article": "der"
  },
  {
    "noun": "Flugzeugpark",
    "article": "der"
  },
  {
    "noun": "Flugzeugrumpf",
    "article": "der"
  },
  {
    "noun": "Flugzeugtyp",
    "article": "der"
  },
  {
    "noun": "Flunkerer",
    "article": "der"
  },
  {
    "noun": "Fluoreszenzschirm",
    "article": "der"
  },
  {
    "noun": "Fluorit",
    "article": "der"
  },
  {
    "noun": "Fluortest",
    "article": "der"
  },
  {
    "noun": "Fluorwasserstoff",
    "article": "der"
  },
  {
    "noun": "Flur",
    "article": "der"
  },
  {
    "noun": "Flurname",
    "article": "der"
  },
  {
    "noun": "Flurschaden",
    "article": "der"
  },
  {
    "noun": "Fluss",
    "article": "der"
  },
  {
    "noun": "Flussarm",
    "article": "der"
  },
  {
    "noun": "Flussbarsch",
    "article": "der"
  },
  {
    "noun": "Flussfisch",
    "article": "der"
  },
  {
    "noun": "Flusshafen",
    "article": "der"
  },
  {
    "noun": "Flusskies",
    "article": "der"
  },
  {
    "noun": "Flusskrebs",
    "article": "der"
  },
  {
    "noun": "Flusslauf",
    "article": "der"
  },
  {
    "noun": "Flyer",
    "article": "der"
  },
  {
    "noun": "Fockmast",
    "article": "der"
  },
  {
    "noun": "Fokus",
    "article": "der"
  },
  {
    "noun": "Folgesatz",
    "article": "der"
  },
  {
    "noun": "Folgeschaden",
    "article": "der"
  },
  {
    "noun": "Foliant",
    "article": "der"
  },
  {
    "noun": "Folklorist",
    "article": "der"
  },
  {
    "noun": "Folksong",
    "article": "der"
  },
  {
    "noun": "Follikel",
    "article": "der"
  },
  {
    "noun": "Follikelsprung",
    "article": "der"
  },
  {
    "noun": "Folterer",
    "article": "der"
  },
  {
    "noun": "Folterknecht",
    "article": "der"
  },
  {
    "noun": "Fond",
    "article": "der"
  },
  {
    "noun": "Fonds",
    "article": "der"
  },
  {
    "noun": "Fonograf",
    "article": "der"
  },
  {
    "noun": "Fonotypist",
    "article": "der"
  },
  {
    "noun": "Font",
    "article": "der"
  },
  {
    "noun": "Football",
    "article": "der"
  },
  {
    "noun": "Fordismus",
    "article": "der"
  },
  {
    "noun": "Forint",
    "article": "der"
  },
  {
    "noun": "Formaldehyd",
    "article": "der"
  },
  {
    "noun": "Formalismus",
    "article": "der"
  },
  {
    "noun": "Formalist",
    "article": "der"
  },
  {
    "noun": "Formationsflug",
    "article": "der"
  },
  {
    "noun": "Former",
    "article": "der"
  },
  {
    "noun": "Formfehler",
    "article": "der"
  },
  {
    "noun": "Formgestalter",
    "article": "der"
  },
  {
    "noun": "Formsand",
    "article": "der"
  },
  {
    "noun": "Formschnitt",
    "article": "der"
  },
  {
    "noun": "Forscher",
    "article": "der"
  },
  {
    "noun": "Forschungsauftrag",
    "article": "der"
  },
  {
    "noun": "Forschungsbereich",
    "article": "der"
  },
  {
    "noun": "Forschungsbericht",
    "article": "der"
  },
  {
    "noun": "Forschungsgegenstand",
    "article": "der"
  },
  {
    "noun": "Forschungsplan",
    "article": "der"
  },
  {
    "noun": "Forst",
    "article": "der"
  },
  {
    "noun": "Forstbetrieb",
    "article": "der"
  },
  {
    "noun": "Forstmann",
    "article": "der"
  },
  {
    "noun": "Forstmeister",
    "article": "der"
  },
  {
    "noun": "Forstschaden",
    "article": "der"
  },
  {
    "noun": "Forstschutz",
    "article": "der"
  },
  {
    "noun": "Forstwirt",
    "article": "der"
  },
  {
    "noun": "Fortbestand",
    "article": "der"
  },
  {
    "noun": "Fortfall",
    "article": "der"
  },
  {
    "noun": "Fortgang",
    "article": "der"
  },
  {
    "noun": "Fortpflanzungstrieb",
    "article": "der"
  },
  {
    "noun": "Fortsatz",
    "article": "der"
  },
  {
    "noun": "Fortschritt",
    "article": "der"
  },
  {
    "noun": "Fortsetzungsroman",
    "article": "der"
  },
  {
    "noun": "Fotoapparat",
    "article": "der"
  },
  {
    "noun": "Fotoeffekt",
    "article": "der"
  },
  {
    "noun": "Fotograf",
    "article": "der"
  },
  {
    "noun": "Fotokopierautomat",
    "article": "der"
  },
  {
    "noun": "Fotokopierer",
    "article": "der"
  },
  {
    "noun": "Fotometer",
    "article": "der"
  },
  {
    "noun": "Fotosatz",
    "article": "der"
  },
  {
    "noun": "Fotowiderstand",
    "article": "der"
  },
  {
    "noun": "Foxterrier",
    "article": "der"
  },
  {
    "noun": "Foxtrott",
    "article": "der"
  },
  {
    "noun": "Frachtbrief",
    "article": "der"
  },
  {
    "noun": "Frachter",
    "article": "der"
  },
  {
    "noun": "Frachtkahn",
    "article": "der"
  },
  {
    "noun": "Frachtraum",
    "article": "der"
  },
  {
    "noun": "Frachtsatz",
    "article": "der"
  },
  {
    "noun": "Frachttarif",
    "article": "der"
  },
  {
    "noun": "Frachtverkehr",
    "article": "der"
  },
  {
    "noun": "Frachtvertrag",
    "article": "der"
  },
  {
    "noun": "Frack",
    "article": "der"
  },
  {
    "noun": "Fragebogen",
    "article": "der"
  },
  {
    "noun": "Fragenkatalog",
    "article": "der"
  },
  {
    "noun": "Fragenkomplex",
    "article": "der"
  },
  {
    "noun": "Frager",
    "article": "der"
  },
  {
    "noun": "Fragesatz",
    "article": "der"
  },
  {
    "noun": "Fragesteller",
    "article": "der"
  },
  {
    "noun": "Franc",
    "article": "der"
  },
  {
    "noun": "Franke",
    "article": "der"
  },
  {
    "noun": "Frankokanadier",
    "article": "der"
  },
  {
    "noun": "Franzbranntwein",
    "article": "der"
  },
  {
    "noun": "Franziskaner",
    "article": "der"
  },
  {
    "noun": "Franzose",
    "article": "der"
  },
  {
    "noun": "Fratz",
    "article": "der"
  },
  {
    "noun": "Frauenarzt",
    "article": "der"
  },
  {
    "noun": "Frauenberuf",
    "article": "der"
  },
  {
    "noun": "Frauenchor",
    "article": "der"
  },
  {
    "noun": "Frauenfeind",
    "article": "der"
  },
  {
    "noun": "Frauenhasser",
    "article": "der"
  },
  {
    "noun": "Frauenheld",
    "article": "der"
  },
  {
    "noun": "Frauenmantel",
    "article": "der"
  },
  {
    "noun": "Frauentag",
    "article": "der"
  },
  {
    "noun": "Frauenverein",
    "article": "der"
  },
  {
    "noun": "Freak",
    "article": "der"
  },
  {
    "noun": "Frechdachs",
    "article": "der"
  },
  {
    "noun": "Frechling",
    "article": "der"
  },
  {
    "noun": "Freibauer",
    "article": "der"
  },
  {
    "noun": "Freiberufler",
    "article": "der"
  },
  {
    "noun": "Freibetrag",
    "article": "der"
  },
  {
    "noun": "Freibeuter",
    "article": "der"
  },
  {
    "noun": "Freibord",
    "article": "der"
  },
  {
    "noun": "Freibrief",
    "article": "der"
  },
  {
    "noun": "Freidenker",
    "article": "der"
  },
  {
    "noun": "Freier",
    "article": "der"
  },
  {
    "noun": "Freiflug",
    "article": "der"
  },
  {
    "noun": "Freigang",
    "article": "der"
  },
  {
    "noun": "Freigeist",
    "article": "der"
  },
  {
    "noun": "Freihafen",
    "article": "der"
  },
  {
    "noun": "Freihandel",
    "article": "der"
  },
  {
    "noun": "Freiheitsbegriff",
    "article": "der"
  },
  {
    "noun": "Freiheitsdrang",
    "article": "der"
  },
  {
    "noun": "Freiheitsentzug",
    "article": "der"
  },
  {
    "noun": "Freiheitsgrad",
    "article": "der"
  },
  {
    "noun": "Freiheitskampf",
    "article": "der"
  },
  {
    "noun": "Freiheitskrieg",
    "article": "der"
  },
  {
    "noun": "Freiherr",
    "article": "der"
  },
  {
    "noun": "Freilauf",
    "article": "der"
  },
  {
    "noun": "Freimaurer",
    "article": "der"
  },
  {
    "noun": "Freimut",
    "article": "der"
  },
  {
    "noun": "Freiplatz",
    "article": "der"
  },
  {
    "noun": "Freiraum",
    "article": "der"
  },
  {
    "noun": "Freisass",
    "article": "der"
  },
  {
    "noun": "Freisinn",
    "article": "der"
  },
  {
    "noun": "Freispruch",
    "article": "der"
  },
  {
    "noun": "Freistaat",
    "article": "der"
  },
  {
    "noun": "Freitag",
    "article": "der"
  },
  {
    "noun": "Freitisch",
    "article": "der"
  },
  {
    "noun": "Freiumschlag",
    "article": "der"
  },
  {
    "noun": "Freiverkehr",
    "article": "der"
  },
  {
    "noun": "Freiwurf",
    "article": "der"
  },
  {
    "noun": "Freizeitpark",
    "article": "der"
  },
  {
    "noun": "Freizeitwert",
    "article": "der"
  },
  {
    "noun": "Fremdarbeiter",
    "article": "der"
  },
  {
    "noun": "Fremdbesitz",
    "article": "der"
  },
  {
    "noun": "Fremdenhass",
    "article": "der"
  },
  {
    "noun": "Fremdenverkehr",
    "article": "der"
  },
  {
    "noun": "Fremdling",
    "article": "der"
  },
  {
    "noun": "Fremdsprachenkorrespondent",
    "article": "der"
  },
  {
    "noun": "Fremdsprachenunterricht",
    "article": "der"
  },
  {
    "noun": "Fremdstoff",
    "article": "der"
  },
  {
    "noun": "Frequenzbereich",
    "article": "der"
  },
  {
    "noun": "Frequenzmesser",
    "article": "der"
  },
  {
    "noun": "Frequenzmodulator",
    "article": "der"
  },
  {
    "noun": "Freudenruf",
    "article": "der"
  },
  {
    "noun": "Freudenschrei",
    "article": "der"
  },
  {
    "noun": "Freudensprung",
    "article": "der"
  },
  {
    "noun": "Freudentag",
    "article": "der"
  },
  {
    "noun": "Freudianer",
    "article": "der"
  },
  {
    "noun": "Freund",
    "article": "der"
  },
  {
    "noun": "Freundeskreis",
    "article": "der"
  },
  {
    "noun": "Freundschaftsbeweis",
    "article": "der"
  },
  {
    "noun": "Freundschaftsdienst",
    "article": "der"
  },
  {
    "noun": "Freundschaftsring",
    "article": "der"
  },
  {
    "noun": "Freundschaftsvertrag",
    "article": "der"
  },
  {
    "noun": "Frevel",
    "article": "der"
  },
  {
    "noun": "Frevelmut",
    "article": "der"
  },
  {
    "noun": "Frevler",
    "article": "der"
  },
  {
    "noun": "Friede",
    "article": "der"
  },
  {
    "noun": "Friedensbote",
    "article": "der"
  },
  {
    "noun": "Friedensbruch",
    "article": "der"
  },
  {
    "noun": "Friedensnobelpreis",
    "article": "der"
  },
  {
    "noun": "Friedensrichter",
    "article": "der"
  },
  {
    "noun": "Friedensschluss",
    "article": "der"
  },
  {
    "noun": "Friedensstifter",
    "article": "der"
  },
  {
    "noun": "Friedensvertrag",
    "article": "der"
  },
  {
    "noun": "Friedenswille",
    "article": "der"
  },
  {
    "noun": "Friedhof",
    "article": "der"
  },
  {
    "noun": "Fries",
    "article": "der"
  },
  {
    "noun": "Friese",
    "article": "der"
  },
  {
    "noun": "Frigidaire",
    "article": "der"
  },
  {
    "noun": "Frikativ",
    "article": "der"
  },
  {
    "noun": "Frikativlaut",
    "article": "der"
  },
  {
    "noun": "Frischdampf",
    "article": "der"
  },
  {
    "noun": "Frischfisch",
    "article": "der"
  },
  {
    "noun": "Frischling",
    "article": "der"
  },
  {
    "noun": "Friseur",
    "article": "der"
  },
  {
    "noun": "Friseurladen",
    "article": "der"
  },
  {
    "noun": "Friseursalon",
    "article": "der"
  },
  {
    "noun": "Frisiermantel",
    "article": "der"
  },
  {
    "noun": "Fritter",
    "article": "der"
  },
  {
    "noun": "Frohmut",
    "article": "der"
  },
  {
    "noun": "Frohsinn",
    "article": "der"
  },
  {
    "noun": "Fronbauer",
    "article": "der"
  },
  {
    "noun": "Frondeur",
    "article": "der"
  },
  {
    "noun": "Frondienst",
    "article": "der"
  },
  {
    "noun": "Fronleichnam",
    "article": "der"
  },
  {
    "noun": "Frontalangriff",
    "article": "der"
  },
  {
    "noun": "Frontantrieb",
    "article": "der"
  },
  {
    "noun": "Frontlader",
    "article": "der"
  },
  {
    "noun": "Frontmotor",
    "article": "der"
  },
  {
    "noun": "Frontwechsel",
    "article": "der"
  },
  {
    "noun": "Frontzahn",
    "article": "der"
  },
  {
    "noun": "Frosch",
    "article": "der"
  },
  {
    "noun": "Froschlaich",
    "article": "der"
  },
  {
    "noun": "Froschschenkel",
    "article": "der"
  },
  {
    "noun": "Froschteich",
    "article": "der"
  },
  {
    "noun": "Frost",
    "article": "der"
  },
  {
    "noun": "Frostschaden",
    "article": "der"
  },
  {
    "noun": "Frostschutz",
    "article": "der"
  },
  {
    "noun": "Frotteestoff",
    "article": "der"
  },
  {
    "noun": "Fruchtertrag",
    "article": "der"
  },
  {
    "noun": "Fruchtknoten",
    "article": "der"
  },
  {
    "noun": "Fruchtkuchen",
    "article": "der"
  },
  {
    "noun": "Fruchtsaft",
    "article": "der"
  },
  {
    "noun": "Fruchtsalat",
    "article": "der"
  },
  {
    "noun": "Fruchtwechsel",
    "article": "der"
  },
  {
    "noun": "Fruchtwein",
    "article": "der"
  },
  {
    "noun": "Fruchtzucker",
    "article": "der"
  },
  {
    "noun": "Frust",
    "article": "der"
  },
  {
    "noun": "Fuchs",
    "article": "der"
  },
  {
    "noun": "Fuchsbalg",
    "article": "der"
  },
  {
    "noun": "Fuchsbandwurm",
    "article": "der"
  },
  {
    "noun": "Fuchsbau",
    "article": "der"
  },
  {
    "noun": "Fuchspelz",
    "article": "der"
  },
  {
    "noun": "Fuchsschwanz",
    "article": "der"
  },
  {
    "noun": "Fuhrlohn",
    "article": "der"
  },
  {
    "noun": "Fuhrmann",
    "article": "der"
  },
  {
    "noun": "Fuhrpark",
    "article": "der"
  },
  {
    "noun": "Fuhrunternehmer",
    "article": "der"
  },
  {
    "noun": "Fund",
    "article": "der"
  },
  {
    "noun": "Fundamentalismus",
    "article": "der"
  },
  {
    "noun": "Fundamentalist",
    "article": "der"
  },
  {
    "noun": "Fundgegenstand",
    "article": "der"
  },
  {
    "noun": "Fundi",
    "article": "der"
  },
  {
    "noun": "Fundort",
    "article": "der"
  },
  {
    "noun": "Fundus",
    "article": "der"
  },
  {
    "noun": "Funk",
    "article": "der"
  },
  {
    "noun": "Funkamateur",
    "article": "der"
  },
  {
    "noun": "Funkbetrieb",
    "article": "der"
  },
  {
    "noun": "Funkdienst",
    "article": "der"
  },
  {
    "noun": "Funke",
    "article": "der"
  },
  {
    "noun": "Funkenflug",
    "article": "der"
  },
  {
    "noun": "Funkeninduktor",
    "article": "der"
  },
  {
    "noun": "Funkenregen",
    "article": "der"
  },
  {
    "noun": "Funker",
    "article": "der"
  },
  {
    "noun": "Funkmechaniker",
    "article": "der"
  },
  {
    "noun": "Funkpeiler",
    "article": "der"
  },
  {
    "noun": "Funkschatten",
    "article": "der"
  },
  {
    "noun": "Funksprechverkehr",
    "article": "der"
  },
  {
    "noun": "Funkspruch",
    "article": "der"
  },
  {
    "noun": "Funkstreifenwagen",
    "article": "der"
  },
  {
    "noun": "Funkturm",
    "article": "der"
  },
  {
    "noun": "Funkverkehr",
    "article": "der"
  },
  {
    "noun": "Funkwagen",
    "article": "der"
  },
  {
    "noun": "Funkweg",
    "article": "der"
  },
  {
    "noun": "Furier",
    "article": "der"
  },
  {
    "noun": "Furz",
    "article": "der"
  },
  {
    "noun": "Fusel",
    "article": "der"
  },
  {
    "noun": "Fusionsreaktor",
    "article": "der"
  },
  {
    "noun": "Fusionsvertrag",
    "article": "der"
  },
  {
    "noun": "Fussel",
    "article": "der"
  },
  {
    "noun": "Futon",
    "article": "der"
  },
  {
    "noun": "Futterautomat",
    "article": "der"
  },
  {
    "noun": "Futternapf",
    "article": "der"
  },
  {
    "noun": "Futterneid",
    "article": "der"
  },
  {
    "noun": "Futtersack",
    "article": "der"
  },
  {
    "noun": "Futterstoff",
    "article": "der"
  },
  {
    "noun": "Futtertaft",
    "article": "der"
  },
  {
    "noun": "Futtertrog",
    "article": "der"
  },
  {
    "noun": "Futurismus",
    "article": "der"
  },
  {
    "noun": "Futurist",
    "article": "der"
  },
  {
    "noun": "Tagesordnungspunkt",
    "article": "der"
  },
  {
    "noun": "Tagesplan",
    "article": "der"
  },
  {
    "noun": "Tagespreis",
    "article": "der"
  },
  {
    "noun": "Tagesraum",
    "article": "der"
  },
  {
    "noun": "Tagesrhythmus",
    "article": "der"
  },
  {
    "noun": "Tagessatz",
    "article": "der"
  },
  {
    "noun": "Tagesumsatz",
    "article": "der"
  },
  {
    "noun": "Tageswert",
    "article": "der"
  },
  {
    "noun": "Tagfalter",
    "article": "der"
  },
  {
    "noun": "Tagtraum",
    "article": "der"
  },
  {
    "noun": "Tagungsort",
    "article": "der"
  },
  {
    "noun": "Taifun",
    "article": "der"
  },
  {
    "noun": "Taillenumfang",
    "article": "der"
  },
  {
    "noun": "Taiwanese",
    "article": "der"
  },
  {
    "noun": "Takt",
    "article": "der"
  },
  {
    "noun": "Taktfehler",
    "article": "der"
  },
  {
    "noun": "Taktiker",
    "article": "der"
  },
  {
    "noun": "Taktmesser",
    "article": "der"
  },
  {
    "noun": "Taktstrich",
    "article": "der"
  },
  {
    "noun": "Taktwechsel",
    "article": "der"
  },
  {
    "noun": "Talar",
    "article": "der"
  },
  {
    "noun": "Talboden",
    "article": "der"
  },
  {
    "noun": "Taler",
    "article": "der"
  },
  {
    "noun": "Talg",
    "article": "der"
  },
  {
    "noun": "Talisman",
    "article": "der"
  },
  {
    "noun": "Talk",
    "article": "der"
  },
  {
    "noun": "Talkessel",
    "article": "der"
  },
  {
    "noun": "Talkmaster",
    "article": "der"
  },
  {
    "noun": "Talkumpuder",
    "article": "der"
  },
  {
    "noun": "Talmiglanz",
    "article": "der"
  },
  {
    "noun": "Talmud",
    "article": "der"
  },
  {
    "noun": "Talmudist",
    "article": "der"
  },
  {
    "noun": "Talon",
    "article": "der"
  },
  {
    "noun": "Talweg",
    "article": "der"
  },
  {
    "noun": "Tambour",
    "article": "der"
  },
  {
    "noun": "Tampon",
    "article": "der"
  },
  {
    "noun": "Tanbur",
    "article": "der"
  },
  {
    "noun": "Tand",
    "article": "der"
  },
  {
    "noun": "Tandelmarkt",
    "article": "der"
  },
  {
    "noun": "Tandler",
    "article": "der"
  },
  {
    "noun": "Tang",
    "article": "der"
  },
  {
    "noun": "Tanga",
    "article": "der"
  },
  {
    "noun": "Tangaslip",
    "article": "der"
  },
  {
    "noun": "Tangens",
    "article": "der"
  },
  {
    "noun": "Tangenssatz",
    "article": "der"
  },
  {
    "noun": "Tangentenvektor",
    "article": "der"
  },
  {
    "noun": "Tango",
    "article": "der"
  },
  {
    "noun": "Tank",
    "article": "der"
  },
  {
    "noun": "Tanker",
    "article": "der"
  },
  {
    "noun": "Tankwagen",
    "article": "der"
  },
  {
    "noun": "Tankwart",
    "article": "der"
  },
  {
    "noun": "Tankzug",
    "article": "der"
  },
  {
    "noun": "Tannenbaum",
    "article": "der"
  },
  {
    "noun": "Tannenhonig",
    "article": "der"
  },
  {
    "noun": "Tannenpfeil",
    "article": "der"
  },
  {
    "noun": "Tannenwald",
    "article": "der"
  },
  {
    "noun": "Tannenzapfen",
    "article": "der"
  },
  {
    "noun": "Tannenzweig",
    "article": "der"
  },
  {
    "noun": "Tannicht",
    "article": "der"
  },
  {
    "noun": "Tansanier",
    "article": "der"
  },
  {
    "noun": "Tansanit",
    "article": "der"
  },
  {
    "noun": "Tanz",
    "article": "der"
  },
  {
    "noun": "Tanzabend",
    "article": "der"
  },
  {
    "noun": "Tanzboden",
    "article": "der"
  },
  {
    "noun": "Tanzkurs",
    "article": "der"
  },
  {
    "noun": "Tanzlehrer",
    "article": "der"
  },
  {
    "noun": "Tanzmeister",
    "article": "der"
  },
  {
    "noun": "Tanzpartner",
    "article": "der"
  },
  {
    "noun": "Tanzsaal",
    "article": "der"
  },
  {
    "noun": "Tanzschritt",
    "article": "der"
  },
  {
    "noun": "Tanzschuh",
    "article": "der"
  },
  {
    "noun": "Tanzsport",
    "article": "der"
  },
  {
    "noun": "Tanztee",
    "article": "der"
  },
  {
    "noun": "Tanzunterricht",
    "article": "der"
  },
  {
    "noun": "Tapetenkleister",
    "article": "der"
  },
  {
    "noun": "Tapezierer",
    "article": "der"
  },
  {
    "noun": "Tapfen",
    "article": "der"
  },
  {
    "noun": "Tapir",
    "article": "der"
  },
  {
    "noun": "Taps",
    "article": "der"
  },
  {
    "noun": "Tarif",
    "article": "der"
  },
  {
    "noun": "Tarifabschluss",
    "article": "der"
  },
  {
    "noun": "Tariflohn",
    "article": "der"
  },
  {
    "noun": "Tarifsatz",
    "article": "der"
  },
  {
    "noun": "Tarifvertrag",
    "article": "der"
  },
  {
    "noun": "Tarnanzug",
    "article": "der"
  },
  {
    "noun": "Tarnmantel",
    "article": "der"
  },
  {
    "noun": "Tarsus",
    "article": "der"
  },
  {
    "noun": "Tartan",
    "article": "der"
  },
  {
    "noun": "Tartaros",
    "article": "der"
  },
  {
    "noun": "Taschendieb",
    "article": "der"
  },
  {
    "noun": "Taschenfahrplan",
    "article": "der"
  },
  {
    "noun": "Tascheninhalt",
    "article": "der"
  },
  {
    "noun": "Taschenkalender",
    "article": "der"
  },
  {
    "noun": "Taschenkamm",
    "article": "der"
  },
  {
    "noun": "Taschenkrebs",
    "article": "der"
  },
  {
    "noun": "Taschenrechner",
    "article": "der"
  },
  {
    "noun": "Taschenschirm",
    "article": "der"
  },
  {
    "noun": "Taschenspiegel",
    "article": "der"
  },
  {
    "noun": "Taschenspieler",
    "article": "der"
  },
  {
    "noun": "Tastendruck",
    "article": "der"
  },
  {
    "noun": "Tastenfernsprecher",
    "article": "der"
  },
  {
    "noun": "Taster",
    "article": "der"
  },
  {
    "noun": "Tastsinn",
    "article": "der"
  },
  {
    "noun": "Tatablauf",
    "article": "der"
  },
  {
    "noun": "Tatbericht",
    "article": "der"
  },
  {
    "noun": "Tatbestand",
    "article": "der"
  },
  {
    "noun": "Tatendrang",
    "article": "der"
  },
  {
    "noun": "Tatort",
    "article": "der"
  },
  {
    "noun": "Tatsachenbericht",
    "article": "der"
  },
  {
    "noun": "Tattergreis",
    "article": "der"
  },
  {
    "noun": "Tatterich",
    "article": "der"
  },
  {
    "noun": "Tattersall",
    "article": "der"
  },
  {
    "noun": "Tatverdacht",
    "article": "der"
  },
  {
    "noun": "Tatzeuge",
    "article": "der"
  },
  {
    "noun": "Taubendreck",
    "article": "der"
  },
  {
    "noun": "Taubenschlag",
    "article": "der"
  },
  {
    "noun": "Tauber",
    "article": "der"
  },
  {
    "noun": "Taucher",
    "article": "der"
  },
  {
    "noun": "Taucheranzug",
    "article": "der"
  },
  {
    "noun": "Tauchsieder",
    "article": "der"
  },
  {
    "noun": "Tauchsport",
    "article": "der"
  },
  {
    "noun": "Taufname",
    "article": "der"
  },
  {
    "noun": "Taufpate",
    "article": "der"
  },
  {
    "noun": "Taufschein",
    "article": "der"
  },
  {
    "noun": "Taufstein",
    "article": "der"
  },
  {
    "noun": "Taufzeuge",
    "article": "der"
  },
  {
    "noun": "Taugenichts",
    "article": "der"
  },
  {
    "noun": "Taumel",
    "article": "der"
  },
  {
    "noun": "Taupunkt",
    "article": "der"
  },
  {
    "noun": "Tausch",
    "article": "der"
  },
  {
    "noun": "Tauscher",
    "article": "der"
  },
  {
    "noun": "Tauschhandel",
    "article": "der"
  },
  {
    "noun": "Tauschverkehr",
    "article": "der"
  },
  {
    "noun": "Tauschvertrag",
    "article": "der"
  },
  {
    "noun": "Tauschwert",
    "article": "der"
  },
  {
    "noun": "Tausender",
    "article": "der"
  },
  {
    "noun": "Tausendsasa",
    "article": "der"
  },
  {
    "noun": "Tautropfen",
    "article": "der"
  },
  {
    "noun": "Tauwind",
    "article": "der"
  },
  {
    "noun": "Taxator",
    "article": "der"
  },
  {
    "noun": "Taxichauffeur",
    "article": "der"
  },
  {
    "noun": "Taxifahrer",
    "article": "der"
  },
  {
    "noun": "Taxistand",
    "article": "der"
  },
  {
    "noun": "Taxkurs",
    "article": "der"
  },
  {
    "noun": "Taxus",
    "article": "der"
  },
  {
    "noun": "Taxwert",
    "article": "der"
  },
  {
    "noun": "Taylorismus",
    "article": "der"
  },
  {
    "noun": "Teakbaum",
    "article": "der"
  },
  {
    "noun": "Teamer",
    "article": "der"
  },
  {
    "noun": "Teamgeist",
    "article": "der"
  },
  {
    "noun": "Techniker",
    "article": "der"
  },
  {
    "noun": "Technokrat",
    "article": "der"
  },
  {
    "noun": "Technologe",
    "article": "der"
  },
  {
    "noun": "Technologiepark",
    "article": "der"
  },
  {
    "noun": "Technologietransfer",
    "article": "der"
  },
  {
    "noun": "Teckel",
    "article": "der"
  },
  {
    "noun": "Teddy",
    "article": "der"
  },
  {
    "noun": "Teebeutel",
    "article": "der"
  },
  {
    "noun": "Teekessel",
    "article": "der"
  },
  {
    "noun": "Teenager",
    "article": "der"
  },
  {
    "noun": "Teer",
    "article": "der"
  },
  {
    "noun": "Teerfarbstoff",
    "article": "der"
  },
  {
    "noun": "Teerstuhl",
    "article": "der"
  },
  {
    "noun": "Teestrauch",
    "article": "der"
  },
  {
    "noun": "Teewagen",
    "article": "der"
  },
  {
    "noun": "Tefsir",
    "article": "der"
  },
  {
    "noun": "Teich",
    "article": "der"
  },
  {
    "noun": "Teichmolch",
    "article": "der"
  },
  {
    "noun": "Teig",
    "article": "der"
  },
  {
    "noun": "Teilabschnitt",
    "article": "der"
  },
  {
    "noun": "Teilakzept",
    "article": "der"
  },
  {
    "noun": "Teilaspekt",
    "article": "der"
  },
  {
    "noun": "Teilbereich",
    "article": "der"
  },
  {
    "noun": "Teilbetrag",
    "article": "der"
  },
  {
    "noun": "Teilchenbeschleuniger",
    "article": "der"
  },
  {
    "noun": "Teiler",
    "article": "der"
  },
  {
    "noun": "Teilerfolg",
    "article": "der"
  },
  {
    "noun": "Teilhaber",
    "article": "der"
  },
  {
    "noun": "Teilnehmer",
    "article": "der"
  },
  {
    "noun": "Teilnehmerbetrieb",
    "article": "der"
  },
  {
    "noun": "Teilstrich",
    "article": "der"
  },
  {
    "noun": "Teilton",
    "article": "der"
  },
  {
    "noun": "Teilungsartikel",
    "article": "der"
  },
  {
    "noun": "Teilwert",
    "article": "der"
  },
  {
    "noun": "Teilzahlungskredit",
    "article": "der"
  },
  {
    "noun": "Teilzeitjob",
    "article": "der"
  },
  {
    "noun": "Teint",
    "article": "der"
  },
  {
    "noun": "Telebrief",
    "article": "der"
  },
  {
    "noun": "Telefonanruf",
    "article": "der"
  },
  {
    "noun": "Telefonanschluss",
    "article": "der"
  },
  {
    "noun": "Telefonapparat",
    "article": "der"
  },
  {
    "noun": "Telefondienst",
    "article": "der"
  },
  {
    "noun": "Telefondraht",
    "article": "der"
  },
  {
    "noun": "Telefonhandel",
    "article": "der"
  },
  {
    "noun": "Telefonist",
    "article": "der"
  },
  {
    "noun": "Telefonmast",
    "article": "der"
  },
  {
    "noun": "Telefonrundspruch",
    "article": "der"
  },
  {
    "noun": "Telefonverkauf",
    "article": "der"
  },
  {
    "noun": "Telefonverkehr",
    "article": "der"
  },
  {
    "noun": "Telegraf",
    "article": "der"
  },
  {
    "noun": "Telegrafenapparat",
    "article": "der"
  },
  {
    "noun": "Telegrafendraht",
    "article": "der"
  },
  {
    "noun": "Telegrafenmast",
    "article": "der"
  },
  {
    "noun": "Telegrafenverkehr",
    "article": "der"
  },
  {
    "noun": "Telegrafist",
    "article": "der"
  },
  {
    "noun": "Telegrammbote",
    "article": "der"
  },
  {
    "noun": "Telegrammstil",
    "article": "der"
  },
  {
    "noun": "Telekonverter",
    "article": "der"
  },
  {
    "noun": "Telekopierer",
    "article": "der"
  },
  {
    "noun": "Telemark",
    "article": "der"
  },
  {
    "noun": "Telemarkschwung",
    "article": "der"
  },
  {
    "noun": "Telepath",
    "article": "der"
  },
  {
    "noun": "Teleprompter",
    "article": "der"
  },
  {
    "noun": "Telespot",
    "article": "der"
  },
  {
    "noun": "Teletext",
    "article": "der"
  },
  {
    "noun": "Teller",
    "article": "der"
  },
  {
    "noun": "Tellerrand",
    "article": "der"
  },
  {
    "noun": "Tempel",
    "article": "der"
  },
  {
    "noun": "Tempelbau",
    "article": "der"
  },
  {
    "noun": "Tempelherr",
    "article": "der"
  },
  {
    "noun": "Tempelraub",
    "article": "der"
  },
  {
    "noun": "Tempelritter",
    "article": "der"
  },
  {
    "noun": "Tempeltanz",
    "article": "der"
  },
  {
    "noun": "Temperamentsausbruch",
    "article": "der"
  },
  {
    "noun": "Temperaturabfall",
    "article": "der"
  },
  {
    "noun": "Temperaturanstieg",
    "article": "der"
  },
  {
    "noun": "Temperaturausgleich",
    "article": "der"
  },
  {
    "noun": "Temperaturkoeffizient",
    "article": "der"
  },
  {
    "noun": "Temperaturmesser",
    "article": "der"
  },
  {
    "noun": "Temperaturregler",
    "article": "der"
  },
  {
    "noun": "Temperaturschreiber",
    "article": "der"
  },
  {
    "noun": "Temperatursturz",
    "article": "der"
  },
  {
    "noun": "Temperaturunterschied",
    "article": "der"
  },
  {
    "noun": "Temperaturwechsel",
    "article": "der"
  },
  {
    "noun": "Temperenzler",
    "article": "der"
  },
  {
    "noun": "Temperguss",
    "article": "der"
  },
  {
    "noun": "Templer",
    "article": "der"
  },
  {
    "noun": "Tempomat",
    "article": "der"
  },
  {
    "noun": "Temporalsatz",
    "article": "der"
  },
  {
    "noun": "Tempowechsel",
    "article": "der"
  },
  {
    "noun": "Tender",
    "article": "der"
  },
  {
    "noun": "Tenesmus",
    "article": "der"
  },
  {
    "noun": "Tennisarm",
    "article": "der"
  },
  {
    "noun": "Tennisball",
    "article": "der"
  },
  {
    "noun": "Tennisclub",
    "article": "der"
  },
  {
    "noun": "Tennisdress",
    "article": "der"
  },
  {
    "noun": "Tennisellenbogen",
    "article": "der"
  },
  {
    "noun": "Tennisplatz",
    "article": "der"
  },
  {
    "noun": "Tennisschuh",
    "article": "der"
  },
  {
    "noun": "Tennisspieler",
    "article": "der"
  },
  {
    "noun": "Tenno",
    "article": "der"
  },
  {
    "noun": "Tenor",
    "article": "der"
  },
  {
    "noun": "Tenorist",
    "article": "der"
  },
  {
    "noun": "Tensor",
    "article": "der"
  },
  {
    "noun": "Teppich",
    "article": "der"
  },
  {
    "noun": "Teppichboden",
    "article": "der"
  },
  {
    "noun": "Teppichklopfer",
    "article": "der"
  },
  {
    "noun": "Teppichschaum",
    "article": "der"
  },
  {
    "noun": "Tequila",
    "article": "der"
  },
  {
    "noun": "Term",
    "article": "der"
  },
  {
    "noun": "Termin",
    "article": "der"
  },
  {
    "noun": "Termindruck",
    "article": "der"
  },
  {
    "noun": "Terminismus",
    "article": "der"
  },
  {
    "noun": "Terminkalender",
    "article": "der"
  },
  {
    "noun": "Terminkurs",
    "article": "der"
  },
  {
    "noun": "Terminmarkt",
    "article": "der"
  },
  {
    "noun": "Terminplan",
    "article": "der"
  },
  {
    "noun": "Terminus",
    "article": "der"
  },
  {
    "noun": "Termitenstaat",
    "article": "der"
  },
  {
    "noun": "Terrazzo",
    "article": "der"
  },
  {
    "noun": "Terrier",
    "article": "der"
  },
  {
    "noun": "Territorialstaat",
    "article": "der"
  },
  {
    "noun": "Terror",
    "article": "der"
  },
  {
    "noun": "Terrorakt",
    "article": "der"
  },
  {
    "noun": "Terrorangriff",
    "article": "der"
  },
  {
    "noun": "Terroranschlag",
    "article": "der"
  },
  {
    "noun": "Terrorismus",
    "article": "der"
  },
  {
    "noun": "Terrorist",
    "article": "der"
  },
  {
    "noun": "Terzel",
    "article": "der"
  },
  {
    "noun": "Tesafilm",
    "article": "der"
  },
  {
    "noun": "Teslastrom",
    "article": "der"
  },
  {
    "noun": "Test",
    "article": "der"
  },
  {
    "noun": "Testamentsvollstrecker",
    "article": "der"
  },
  {
    "noun": "Testator",
    "article": "der"
  },
  {
    "noun": "Tester",
    "article": "der"
  },
  {
    "noun": "Testfahrer",
    "article": "der"
  },
  {
    "noun": "Testfall",
    "article": "der"
  },
  {
    "noun": "Testflug",
    "article": "der"
  },
  {
    "noun": "Testikel",
    "article": "der"
  },
  {
    "noun": "Testlauf",
    "article": "der"
  },
  {
    "noun": "Testpilot",
    "article": "der"
  },
  {
    "noun": "Teststreifen",
    "article": "der"
  },
  {
    "noun": "Tetanus",
    "article": "der"
  },
  {
    "noun": "Tetrachlorkohlenstoff",
    "article": "der"
  },
  {
    "noun": "Tetraedrit",
    "article": "der"
  },
  {
    "noun": "Tetrameter",
    "article": "der"
  },
  {
    "noun": "Teuerungszuschlag",
    "article": "der"
  },
  {
    "noun": "Teufel",
    "article": "der"
  },
  {
    "noun": "Teufelsbraten",
    "article": "der"
  },
  {
    "noun": "Teufelskerl",
    "article": "der"
  },
  {
    "noun": "Teufelskreis",
    "article": "der"
  },
  {
    "noun": "Teufelsrochen",
    "article": "der"
  },
  {
    "noun": "Teutone",
    "article": "der"
  },
  {
    "noun": "Texaner",
    "article": "der"
  },
  {
    "noun": "Text",
    "article": "der"
  },
  {
    "noun": "Textband",
    "article": "der"
  },
  {
    "noun": "Textbaustein",
    "article": "der"
  },
  {
    "noun": "Textdichter",
    "article": "der"
  },
  {
    "noun": "Texter",
    "article": "der"
  },
  {
    "noun": "Textilbetrieb",
    "article": "der"
  },
  {
    "noun": "Textilfabrikant",
    "article": "der"
  },
  {
    "noun": "Textiltechniker",
    "article": "der"
  },
  {
    "noun": "Textvergleich",
    "article": "der"
  },
  {
    "noun": "Thalamus",
    "article": "der"
  },
  {
    "noun": "Thanatologe",
    "article": "der"
  },
  {
    "noun": "Theaterabend",
    "article": "der"
  },
  {
    "noun": "Theaterbau",
    "article": "der"
  },
  {
    "noun": "Theaterbesuch",
    "article": "der"
  },
  {
    "noun": "Theaterbesucher",
    "article": "der"
  },
  {
    "noun": "Theaterdichter",
    "article": "der"
  },
  {
    "noun": "Theaterdirektor",
    "article": "der"
  },
  {
    "noun": "Theaterdonner",
    "article": "der"
  },
  {
    "noun": "Theatererfolg",
    "article": "der"
  },
  {
    "noun": "Theaterfriseur",
    "article": "der"
  },
  {
    "noun": "Theaterkritiker",
    "article": "der"
  },
  {
    "noun": "Theatermacher",
    "article": "der"
  },
  {
    "noun": "Theatermann",
    "article": "der"
  },
  {
    "noun": "Theaterregisseur",
    "article": "der"
  },
  {
    "noun": "Theatersaal",
    "article": "der"
  },
  {
    "noun": "Theatervorhang",
    "article": "der"
  },
  {
    "noun": "Theaterwissenschaftler",
    "article": "der"
  },
  {
    "noun": "Theaterzettel",
    "article": "der"
  },
  {
    "noun": "Theismus",
    "article": "der"
  },
  {
    "noun": "Theist",
    "article": "der"
  },
  {
    "noun": "Themenbereich",
    "article": "der"
  },
  {
    "noun": "Themenkatalog",
    "article": "der"
  },
  {
    "noun": "Themenkomplex",
    "article": "der"
  },
  {
    "noun": "Themenkreis",
    "article": "der"
  },
  {
    "noun": "Themenplan",
    "article": "der"
  },
  {
    "noun": "Theodolit",
    "article": "der"
  },
  {
    "noun": "Theokrat",
    "article": "der"
  },
  {
    "noun": "Theologe",
    "article": "der"
  },
  {
    "noun": "Theoretiker",
    "article": "der"
  },
  {
    "noun": "Theosoph",
    "article": "der"
  },
  {
    "noun": "Therapeut",
    "article": "der"
  },
  {
    "noun": "Thermistor",
    "article": "der"
  },
  {
    "noun": "Thermograf",
    "article": "der"
  },
  {
    "noun": "Thermokauter",
    "article": "der"
  },
  {
    "noun": "Thermophor",
    "article": "der"
  },
  {
    "noun": "Thermoplast",
    "article": "der"
  },
  {
    "noun": "Thermostat",
    "article": "der"
  },
  {
    "noun": "Thesaurus",
    "article": "der"
  },
  {
    "noun": "Thingplatz",
    "article": "der"
  },
  {
    "noun": "Thioplast",
    "article": "der"
  },
  {
    "noun": "Thorax",
    "article": "der"
  },
  {
    "noun": "Thraker",
    "article": "der"
  },
  {
    "noun": "Thriller",
    "article": "der"
  },
  {
    "noun": "Thrips",
    "article": "der"
  },
  {
    "noun": "Thrombozyt",
    "article": "der"
  },
  {
    "noun": "Thrombus",
    "article": "der"
  },
  {
    "noun": "Thron",
    "article": "der"
  },
  {
    "noun": "Thronerbe",
    "article": "der"
  },
  {
    "noun": "Thronfolger",
    "article": "der"
  },
  {
    "noun": "Thronhimmel",
    "article": "der"
  },
  {
    "noun": "Thronsaal",
    "article": "der"
  },
  {
    "noun": "Thronsessel",
    "article": "der"
  },
  {
    "noun": "Thronverzicht",
    "article": "der"
  },
  {
    "noun": "Thronwechsel",
    "article": "der"
  },
  {
    "noun": "Thunfisch",
    "article": "der"
  },
  {
    "noun": "Thymian",
    "article": "der"
  },
  {
    "noun": "Thymus",
    "article": "der"
  },
  {
    "noun": "Thyristor",
    "article": "der"
  },
  {
    "noun": "Tibetaner",
    "article": "der"
  },
  {
    "noun": "Tibeter",
    "article": "der"
  },
  {
    "noun": "Tic",
    "article": "der"
  },
  {
    "noun": "Tick",
    "article": "der"
  },
  {
    "noun": "Ticker",
    "article": "der"
  },
  {
    "noun": "Tidehafen",
    "article": "der"
  },
  {
    "noun": "Tiefangriff",
    "article": "der"
  },
  {
    "noun": "Tiefbau",
    "article": "der"
  },
  {
    "noun": "Tiefbauingenieur",
    "article": "der"
  },
  {
    "noun": "Tiefdecker",
    "article": "der"
  },
  {
    "noun": "Tiefdruck",
    "article": "der"
  },
  {
    "noun": "Tiefflieger",
    "article": "der"
  },
  {
    "noun": "Tiefflug",
    "article": "der"
  },
  {
    "noun": "Tiefgang",
    "article": "der"
  },
  {
    "noun": "Tieflader",
    "article": "der"
  },
  {
    "noun": "Tiefladewagen",
    "article": "der"
  },
  {
    "noun": "Tiefpass",
    "article": "der"
  },
  {
    "noun": "Tiefpunkt",
    "article": "der"
  },
  {
    "noun": "Tiefschlaf",
    "article": "der"
  },
  {
    "noun": "Tiefschlag",
    "article": "der"
  },
  {
    "noun": "Tiefseebergbau",
    "article": "der"
  },
  {
    "noun": "Tiefseefisch",
    "article": "der"
  },
  {
    "noun": "Tiefseeforscher",
    "article": "der"
  },
  {
    "noun": "Tiefseegraben",
    "article": "der"
  },
  {
    "noun": "Tiefseetaucher",
    "article": "der"
  },
  {
    "noun": "Tiefsinn",
    "article": "der"
  },
  {
    "noun": "Tiefstand",
    "article": "der"
  },
  {
    "noun": "Tiefstkurs",
    "article": "der"
  },
  {
    "noun": "Tiefstpreis",
    "article": "der"
  },
  {
    "noun": "Tiefstrahler",
    "article": "der"
  },
  {
    "noun": "Tiefststand",
    "article": "der"
  },
  {
    "noun": "Tiefstwert",
    "article": "der"
  },
  {
    "noun": "Tiefton",
    "article": "der"
  },
  {
    "noun": "Tiegel",
    "article": "der"
  },
  {
    "noun": "Tiegelofen",
    "article": "der"
  },
  {
    "noun": "Tierarzt",
    "article": "der"
  },
  {
    "noun": "Tierfilm",
    "article": "der"
  },
  {
    "noun": "Tierfreund",
    "article": "der"
  },
  {
    "noun": "Tiergarten",
    "article": "der"
  },
  {
    "noun": "Tierhalter",
    "article": "der"
  },
  {
    "noun": "Tierkreis",
    "article": "der"
  },
  {
    "noun": "Tierpark",
    "article": "der"
  },
  {
    "noun": "Tierpfleger",
    "article": "der"
  },
  {
    "noun": "Tierschutz",
    "article": "der"
  },
  {
    "noun": "Tierschutzverein",
    "article": "der"
  },
  {
    "noun": "Tierversuch",
    "article": "der"
  },
  {
    "noun": "Tiger",
    "article": "der"
  },
  {
    "noun": "Tigerhai",
    "article": "der"
  },
  {
    "noun": "Tigris",
    "article": "der"
  },
  {
    "noun": "Tilgungsfonds",
    "article": "der"
  },
  {
    "noun": "Timer",
    "article": "der"
  },
  {
    "noun": "Timpano",
    "article": "der"
  },
  {
    "noun": "Tinnef",
    "article": "der"
  },
  {
    "noun": "Tinnitus",
    "article": "der"
  },
  {
    "noun": "Tintenfisch",
    "article": "der"
  },
  {
    "noun": "Tintenfleck",
    "article": "der"
  },
  {
    "noun": "Tintengummi",
    "article": "der"
  },
  {
    "noun": "Tintenkiller",
    "article": "der"
  },
  {
    "noun": "Tintenklecks",
    "article": "der"
  },
  {
    "noun": "Tintenpilz",
    "article": "der"
  },
  {
    "noun": "Tintenstift",
    "article": "der"
  },
  {
    "noun": "Tintenstrahl",
    "article": "der"
  },
  {
    "noun": "Tintenstrahldrucker",
    "article": "der"
  },
  {
    "noun": "Tip",
    "article": "der"
  },
  {
    "noun": "Tippel",
    "article": "der"
  },
  {
    "noun": "Tippelbruder",
    "article": "der"
  },
  {
    "noun": "Tippfehler",
    "article": "der"
  },
  {
    "noun": "Tisch",
    "article": "der"
  },
  {
    "noun": "Tischbesen",
    "article": "der"
  },
  {
    "noun": "Tischcomputer",
    "article": "der"
  },
  {
    "noun": "Tischfernsprecher",
    "article": "der"
  },
  {
    "noun": "Tischgast",
    "article": "der"
  },
  {
    "noun": "Tischgrill",
    "article": "der"
  },
  {
    "noun": "Tischherr",
    "article": "der"
  },
  {
    "noun": "Tischler",
    "article": "der"
  },
  {
    "noun": "Tischnachbar",
    "article": "der"
  },
  {
    "noun": "Tischrand",
    "article": "der"
  },
  {
    "noun": "Tischrechner",
    "article": "der"
  },
  {
    "noun": "Tischtennisspieler",
    "article": "der"
  },
  {
    "noun": "Tischwein",
    "article": "der"
  },
  {
    "noun": "Titanit",
    "article": "der"
  },
  {
    "noun": "Titel",
    "article": "der"
  },
  {
    "noun": "Titelbogen",
    "article": "der"
  },
  {
    "noun": "Titelhalter",
    "article": "der"
  },
  {
    "noun": "Titelheld",
    "article": "der"
  },
  {
    "noun": "Titelkampf",
    "article": "der"
  },
  {
    "noun": "Titelschutz",
    "article": "der"
  },
  {
    "noun": "Titelsong",
    "article": "der"
  },
  {
    "noun": "Titelverteidiger",
    "article": "der"
  },
  {
    "noun": "Titer",
    "article": "der"
  },
  {
    "noun": "Titular",
    "article": "der"
  },
  {
    "noun": "Toast",
    "article": "der"
  },
  {
    "noun": "Toaster",
    "article": "der"
  },
  {
    "noun": "Toboggan",
    "article": "der"
  },
  {
    "noun": "Tobsuchtsanfall",
    "article": "der"
  },
  {
    "noun": "Tofu",
    "article": "der"
  },
  {
    "noun": "Togolese",
    "article": "der"
  },
  {
    "noun": "Toilettenartikel",
    "article": "der"
  },
  {
    "noun": "Toilettenraum",
    "article": "der"
  },
  {
    "noun": "Toilettentisch",
    "article": "der"
  },
  {
    "noun": "Toleranzbereich",
    "article": "der"
  },
  {
    "noun": "Tollkopf",
    "article": "der"
  },
  {
    "noun": "Tollpatsch",
    "article": "der"
  },
  {
    "noun": "Tomahawk",
    "article": "der"
  },
  {
    "noun": "Tomatensaft",
    "article": "der"
  },
  {
    "noun": "Tombak",
    "article": "der"
  },
  {
    "noun": "Tommy",
    "article": "der"
  },
  {
    "noun": "Ton",
    "article": "der"
  },
  {
    "noun": "Tonabnehmer",
    "article": "der"
  },
  {
    "noun": "Tonarm",
    "article": "der"
  },
  {
    "noun": "Tonausfall",
    "article": "der"
  },
  {
    "noun": "Tonboden",
    "article": "der"
  },
  {
    "noun": "Tondichter",
    "article": "der"
  },
  {
    "noun": "Tonfall",
    "article": "der"
  },
  {
    "noun": "Tonfilm",
    "article": "der"
  },
  {
    "noun": "Toningenieur",
    "article": "der"
  },
  {
    "noun": "Tonkopf",
    "article": "der"
  },
  {
    "noun": "Tonkrug",
    "article": "der"
  },
  {
    "noun": "Tonmeister",
    "article": "der"
  },
  {
    "noun": "Tonmischer",
    "article": "der"
  },
  {
    "noun": "Tonnengehalt",
    "article": "der"
  },
  {
    "noun": "Tonsatz",
    "article": "der"
  },
  {
    "noun": "Tonschiefer",
    "article": "der"
  },
  {
    "noun": "Tonschneider",
    "article": "der"
  },
  {
    "noun": "Tonsetzer",
    "article": "der"
  },
  {
    "noun": "Tontechniker",
    "article": "der"
  },
  {
    "noun": "Tontopf",
    "article": "der"
  },
  {
    "noun": "Tonumfang",
    "article": "der"
  },
  {
    "noun": "Tonus",
    "article": "der"
  },
  {
    "noun": "Tonwert",
    "article": "der"
  },
  {
    "noun": "Tonziegel",
    "article": "der"
  },
  {
    "noun": "Topas",
    "article": "der"
  },
  {
    "noun": "Topf",
    "article": "der"
  },
  {
    "noun": "Topfdeckel",
    "article": "der"
  },
  {
    "noun": "Topfhut",
    "article": "der"
  },
  {
    "noun": "Topfkratzer",
    "article": "der"
  },
  {
    "noun": "Topfkuchen",
    "article": "der"
  },
  {
    "noun": "Topflappen",
    "article": "der"
  },
  {
    "noun": "Topfreiniger",
    "article": "der"
  },
  {
    "noun": "Tophus",
    "article": "der"
  },
  {
    "noun": "Topograf",
    "article": "der"
  },
  {
    "noun": "Topp",
    "article": "der"
  },
  {
    "noun": "Torbogen",
    "article": "der"
  },
  {
    "noun": "Tordalk",
    "article": "der"
  },
  {
    "noun": "Toreador",
    "article": "der"
  },
  {
    "noun": "Torero",
    "article": "der"
  },
  {
    "noun": "Torf",
    "article": "der"
  },
  {
    "noun": "Torfballen",
    "article": "der"
  },
  {
    "noun": "Torfboden",
    "article": "der"
  },
  {
    "noun": "Torfmull",
    "article": "der"
  },
  {
    "noun": "Torfstecher",
    "article": "der"
  },
  {
    "noun": "Torfstich",
    "article": "der"
  },
  {
    "noun": "Torlauf",
    "article": "der"
  },
  {
    "noun": "Tormann",
    "article": "der"
  },
  {
    "noun": "Tornado",
    "article": "der"
  },
  {
    "noun": "Tornister",
    "article": "der"
  },
  {
    "noun": "Torpedo",
    "article": "der"
  },
  {
    "noun": "Torpfosten",
    "article": "der"
  },
  {
    "noun": "Torques",
    "article": "der"
  },
  {
    "noun": "Torraum",
    "article": "der"
  },
  {
    "noun": "Torschuss",
    "article": "der"
  },
  {
    "noun": "Torsionsbruch",
    "article": "der"
  },
  {
    "noun": "Torso",
    "article": "der"
  },
  {
    "noun": "Tortenboden",
    "article": "der"
  },
  {
    "noun": "Tortenguss",
    "article": "der"
  },
  {
    "noun": "Tortenheber",
    "article": "der"
  },
  {
    "noun": "Torus",
    "article": "der"
  },
  {
    "noun": "Torwart",
    "article": "der"
  },
  {
    "noun": "Torweg",
    "article": "der"
  },
  {
    "noun": "Tory",
    "article": "der"
  },
  {
    "noun": "Totalausfall",
    "article": "der"
  },
  {
    "noun": "Totalausverkauf",
    "article": "der"
  },
  {
    "noun": "Totaleindruck",
    "article": "der"
  },
  {
    "noun": "Totalisator",
    "article": "der"
  },
  {
    "noun": "Totalitarismus",
    "article": "der"
  },
  {
    "noun": "Totalschaden",
    "article": "der"
  },
  {
    "noun": "Totemismus",
    "article": "der"
  },
  {
    "noun": "Totempfahl",
    "article": "der"
  },
  {
    "noun": "Totenacker",
    "article": "der"
  },
  {
    "noun": "Totenfleck",
    "article": "der"
  },
  {
    "noun": "Totenkopf",
    "article": "der"
  },
  {
    "noun": "Totenkult",
    "article": "der"
  },
  {
    "noun": "Totenschein",
    "article": "der"
  },
  {
    "noun": "Totenschrein",
    "article": "der"
  },
  {
    "noun": "Totensonntag",
    "article": "der"
  },
  {
    "noun": "Totentanz",
    "article": "der"
  },
  {
    "noun": "Totentempel",
    "article": "der"
  },
  {
    "noun": "Totenvogel",
    "article": "der"
  },
  {
    "noun": "Totmannknopf",
    "article": "der"
  },
  {
    "noun": "Totogewinn",
    "article": "der"
  },
  {
    "noun": "Totoschein",
    "article": "der"
  },
  {
    "noun": "Totpunkt",
    "article": "der"
  },
  {
    "noun": "Totschlag",
    "article": "der"
  },
  {
    "noun": "Tourenwagen",
    "article": "der"
  },
  {
    "noun": "Tourismus",
    "article": "der"
  },
  {
    "noun": "Tourist",
    "article": "der"
  },
  {
    "noun": "Touristenrummel",
    "article": "der"
  },
  {
    "noun": "Touristenstrom",
    "article": "der"
  },
  {
    "noun": "Toxikologe",
    "article": "der"
  },
  {
    "noun": "Trab",
    "article": "der"
  },
  {
    "noun": "Trabant",
    "article": "der"
  },
  {
    "noun": "Traber",
    "article": "der"
  },
  {
    "noun": "Tracer",
    "article": "der"
  },
  {
    "noun": "Trachtenanzug",
    "article": "der"
  },
  {
    "noun": "Trachtenverein",
    "article": "der"
  },
  {
    "noun": "Trackball",
    "article": "der"
  },
  {
    "noun": "Traditionalismus",
    "article": "der"
  },
  {
    "noun": "Traditionalist",
    "article": "der"
  },
  {
    "noun": "Trafo",
    "article": "der"
  },
  {
    "noun": "Tragegriff",
    "article": "der"
  },
  {
    "noun": "Tragegurt",
    "article": "der"
  },
  {
    "noun": "Tragekomfort",
    "article": "der"
  },
  {
    "noun": "Traggriff",
    "article": "der"
  },
  {
    "noun": "Traghimmel",
    "article": "der"
  },
  {
    "noun": "Tragiker",
    "article": "der"
  },
  {
    "noun": "Tragkorb",
    "article": "der"
  },
  {
    "noun": "Tragriemen",
    "article": "der"
  },
  {
    "noun": "Tragschrauber",
    "article": "der"
  },
  {
    "noun": "Trailer",
    "article": "der"
  },
  {
    "noun": "Trainer",
    "article": "der"
  },
  {
    "noun": "Trainingsanzug",
    "article": "der"
  },
  {
    "noun": "Trainingsplan",
    "article": "der"
  },
  {
    "noun": "Traiteur",
    "article": "der"
  },
  {
    "noun": "Trakt",
    "article": "der"
  },
  {
    "noun": "Traktor",
    "article": "der"
  },
  {
    "noun": "Traktorist",
    "article": "der"
  },
  {
    "noun": "Tramp",
    "article": "der"
  },
  {
    "noun": "Trampelpfad",
    "article": "der"
  },
  {
    "noun": "Tramper",
    "article": "der"
  },
  {
    "noun": "Tran",
    "article": "der"
  },
  {
    "noun": "Trank",
    "article": "der"
  },
  {
    "noun": "Tranquilizer",
    "article": "der"
  },
  {
    "noun": "Transfer",
    "article": "der"
  },
  {
    "noun": "Transformator",
    "article": "der"
  },
  {
    "noun": "Transistor",
    "article": "der"
  },
  {
    "noun": "Transit",
    "article": "der"
  },
  {
    "noun": "Transithandel",
    "article": "der"
  },
  {
    "noun": "Transitverkehr",
    "article": "der"
  },
  {
    "noun": "Transitweg",
    "article": "der"
  },
  {
    "noun": "Transitzoll",
    "article": "der"
  },
  {
    "noun": "Translator",
    "article": "der"
  },
  {
    "noun": "Transmitter",
    "article": "der"
  },
  {
    "noun": "Transponder",
    "article": "der"
  },
  {
    "noun": "Transport",
    "article": "der"
  },
  {
    "noun": "Transportarbeiter",
    "article": "der"
  },
  {
    "noun": "Transportbeton",
    "article": "der"
  },
  {
    "noun": "Transporter",
    "article": "der"
  },
  {
    "noun": "Transporteur",
    "article": "der"
  },
  {
    "noun": "Transportunternehmer",
    "article": "der"
  },
  {
    "noun": "Transvestismus",
    "article": "der"
  },
  {
    "noun": "Transvestit",
    "article": "der"
  },
  {
    "noun": "Transvestitismus",
    "article": "der"
  },
  {
    "noun": "Trapezmuskel",
    "article": "der"
  },
  {
    "noun": "Trapper",
    "article": "der"
  },
  {
    "noun": "Trappist",
    "article": "der"
  },
  {
    "noun": "Trassant",
    "article": "der"
  },
  {
    "noun": "Trassat",
    "article": "der"
  },
  {
    "noun": "Tratsch",
    "article": "der"
  },
  {
    "noun": "Traubenholunder",
    "article": "der"
  },
  {
    "noun": "Traubenkamm",
    "article": "der"
  },
  {
    "noun": "Traubenmost",
    "article": "der"
  },
  {
    "noun": "Traubensaft",
    "article": "der"
  },
  {
    "noun": "Traubenzucker",
    "article": "der"
  },
  {
    "noun": "Trauerbrief",
    "article": "der"
  },
  {
    "noun": "Trauerfall",
    "article": "der"
  },
  {
    "noun": "Trauerflor",
    "article": "der"
  },
  {
    "noun": "Trauergottesdienst",
    "article": "der"
  },
  {
    "noun": "Trauermarsch",
    "article": "der"
  },
  {
    "noun": "Trauerrand",
    "article": "der"
  },
  {
    "noun": "Trauerschleier",
    "article": "der"
  },
  {
    "noun": "Trauerzug",
    "article": "der"
  },
  {
    "noun": "Traum",
    "article": "der"
  },
  {
    "noun": "Traumberuf",
    "article": "der"
  },
  {
    "noun": "Traumdeuter",
    "article": "der"
  },
  {
    "noun": "Traumjob",
    "article": "der"
  },
  {
    "noun": "Trauring",
    "article": "der"
  },
  {
    "noun": "Trauschein",
    "article": "der"
  },
  {
    "noun": "Trauzeuge",
    "article": "der"
  },
  {
    "noun": "Travellerscheck",
    "article": "der"
  },
  {
    "noun": "Trawler",
    "article": "der"
  },
  {
    "noun": "Treber",
    "article": "der"
  },
  {
    "noun": "Treck",
    "article": "der"
  },
  {
    "noun": "Trecker",
    "article": "der"
  },
  {
    "noun": "Treckerfahrer",
    "article": "der"
  },
  {
    "noun": "Treffer",
    "article": "der"
  },
  {
    "noun": "Treffpunkt",
    "article": "der"
  },
  {
    "noun": "Treibanker",
    "article": "der"
  },
  {
    "noun": "Treiber",
    "article": "der"
  },
  {
    "noun": "Treibhauseffekt",
    "article": "der"
  },
  {
    "noun": "Treibriemen",
    "article": "der"
  },
  {
    "noun": "Treibsand",
    "article": "der"
  },
  {
    "noun": "Treibschlag",
    "article": "der"
  },
  {
    "noun": "Treibstoff",
    "article": "der"
  },
  {
    "noun": "Treibstofftank",
    "article": "der"
  },
  {
    "noun": "Treibstoffverbrauch",
    "article": "der"
  },
  {
    "noun": "Treidelpfad",
    "article": "der"
  },
  {
    "noun": "Treidler",
    "article": "der"
  },
  {
    "noun": "Tremor",
    "article": "der"
  },
  {
    "noun": "Trenchcoat",
    "article": "der"
  },
  {
    "noun": "Trend",
    "article": "der"
  },
  {
    "noun": "Trenner",
    "article": "der"
  },
  {
    "noun": "Trennungsschmerz",
    "article": "der"
  },
  {
    "noun": "Trennungsstrich",
    "article": "der"
  },
  {
    "noun": "Trepan",
    "article": "der"
  },
  {
    "noun": "Treppenabsatz",
    "article": "der"
  },
  {
    "noun": "Treppenflur",
    "article": "der"
  },
  {
    "noun": "Treppenlauf",
    "article": "der"
  },
  {
    "noun": "Treppenschacht",
    "article": "der"
  },
  {
    "noun": "Tresen",
    "article": "der"
  },
  {
    "noun": "Tresor",
    "article": "der"
  },
  {
    "noun": "Tresorraum",
    "article": "der"
  },
  {
    "noun": "Trester",
    "article": "der"
  },
  {
    "noun": "Treter",
    "article": "der"
  },
  {
    "noun": "Tretroller",
    "article": "der"
  },
  {
    "noun": "Treubruch",
    "article": "der"
  },
  {
    "noun": "Treuebruch",
    "article": "der"
  },
  {
    "noun": "Treueeid",
    "article": "der"
  },
  {
    "noun": "Treueid",
    "article": "der"
  },
  {
    "noun": "Treuerabatt",
    "article": "der"
  },
  {
    "noun": "Treugeber",
    "article": "der"
  },
  {
    "noun": "Triathlet",
    "article": "der"
  },
  {
    "noun": "Tribalismus",
    "article": "der"
  },
  {
    "noun": "Tribrachys",
    "article": "der"
  },
  {
    "noun": "Tribun",
    "article": "der"
  },
  {
    "noun": "Tribut",
    "article": "der"
  },
  {
    "noun": "Trichter",
    "article": "der"
  },
  {
    "noun": "Trick",
    "article": "der"
  },
  {
    "noun": "Trickbetrug",
    "article": "der"
  },
  {
    "noun": "Trickfilm",
    "article": "der"
  },
  {
    "noun": "Trickser",
    "article": "der"
  },
  {
    "noun": "Trident",
    "article": "der"
  },
  {
    "noun": "Trieb",
    "article": "der"
  },
  {
    "noun": "Triebkopf",
    "article": "der"
  },
  {
    "noun": "Triebsand",
    "article": "der"
  },
  {
    "noun": "Triebverbrecher",
    "article": "der"
  },
  {
    "noun": "Triebwagen",
    "article": "der"
  },
  {
    "noun": "Triel",
    "article": "der"
  },
  {
    "noun": "Trieur",
    "article": "der"
  },
  {
    "noun": "Trigeminus",
    "article": "der"
  },
  {
    "noun": "Trigger",
    "article": "der"
  },
  {
    "noun": "Triller",
    "article": "der"
  },
  {
    "noun": "Trilobit",
    "article": "der"
  },
  {
    "noun": "Trimaran",
    "article": "der"
  },
  {
    "noun": "Trimeter",
    "article": "der"
  },
  {
    "noun": "Trimm",
    "article": "der"
  },
  {
    "noun": "Trimmer",
    "article": "der"
  },
  {
    "noun": "Trimmpfad",
    "article": "der"
  },
  {
    "noun": "Trimmtrab",
    "article": "der"
  },
  {
    "noun": "Trinkbecher",
    "article": "der"
  },
  {
    "noun": "Trinker",
    "article": "der"
  },
  {
    "noun": "Trinkhalm",
    "article": "der"
  },
  {
    "noun": "Trinkspruch",
    "article": "der"
  },
  {
    "noun": "Trip",
    "article": "der"
  },
  {
    "noun": "Tripelpunkt",
    "article": "der"
  },
  {
    "noun": "Trismus",
    "article": "der"
  },
  {
    "noun": "Tritagonist",
    "article": "der"
  },
  {
    "noun": "Tritt",
    "article": "der"
  },
  {
    "noun": "Trittbrettfahrer",
    "article": "der"
  },
  {
    "noun": "Triumph",
    "article": "der"
  },
  {
    "noun": "Triumphator",
    "article": "der"
  },
  {
    "noun": "Triumphbogen",
    "article": "der"
  },
  {
    "noun": "Triumphwagen",
    "article": "der"
  },
  {
    "noun": "Triumphzug",
    "article": "der"
  },
  {
    "noun": "Triumvir",
    "article": "der"
  },
  {
    "noun": "Trivialroman",
    "article": "der"
  },
  {
    "noun": "Trizeps",
    "article": "der"
  },
  {
    "noun": "Trockenapparat",
    "article": "der"
  },
  {
    "noun": "Trockenboden",
    "article": "der"
  },
  {
    "noun": "Trockenofen",
    "article": "der"
  },
  {
    "noun": "Trockenplatz",
    "article": "der"
  },
  {
    "noun": "Trockenrasierer",
    "article": "der"
  },
  {
    "noun": "Trockner",
    "article": "der"
  },
  {
    "noun": "Troer",
    "article": "der"
  },
  {
    "noun": "Trog",
    "article": "der"
  },
  {
    "noun": "Trojaner",
    "article": "der"
  },
  {
    "noun": "Trokar",
    "article": "der"
  },
  {
    "noun": "Troll",
    "article": "der"
  },
  {
    "noun": "Trolleybus",
    "article": "der"
  },
  {
    "noun": "Trollinger",
    "article": "der"
  },
  {
    "noun": "Trommelrevolver",
    "article": "der"
  },
  {
    "noun": "Trommelschlag",
    "article": "der"
  },
  {
    "noun": "Trommelschlegel",
    "article": "der"
  },
  {
    "noun": "Trommelstock",
    "article": "der"
  },
  {
    "noun": "Trommelwirbel",
    "article": "der"
  },
  {
    "noun": "Trommler",
    "article": "der"
  },
  {
    "noun": "Trompetenbaum",
    "article": "der"
  },
  {
    "noun": "Trompeter",
    "article": "der"
  },
  {
    "noun": "Trompetervogel",
    "article": "der"
  },
  {
    "noun": "Tropenanzug",
    "article": "der"
  },
  {
    "noun": "Tropenhelm",
    "article": "der"
  },
  {
    "noun": "Tropenkoller",
    "article": "der"
  },
  {
    "noun": "Tropf",
    "article": "der"
  },
  {
    "noun": "Tropfstein",
    "article": "der"
  },
  {
    "noun": "Tropismus",
    "article": "der"
  },
  {
    "noun": "Tropophyt",
    "article": "der"
  },
  {
    "noun": "Tropus",
    "article": "der"
  },
  {
    "noun": "Trost",
    "article": "der"
  },
  {
    "noun": "Trostpreis",
    "article": "der"
  },
  {
    "noun": "Trostspender",
    "article": "der"
  },
  {
    "noun": "Trostspruch",
    "article": "der"
  },
  {
    "noun": "Trott",
    "article": "der"
  },
  {
    "noun": "Trottel",
    "article": "der"
  },
  {
    "noun": "Trotz",
    "article": "der"
  },
  {
    "noun": "Trotzer",
    "article": "der"
  },
  {
    "noun": "Trotzkismus",
    "article": "der"
  },
  {
    "noun": "Trotzkist",
    "article": "der"
  },
  {
    "noun": "Trotzkopf",
    "article": "der"
  },
  {
    "noun": "Troubadour",
    "article": "der"
  },
  {
    "noun": "Trouble",
    "article": "der"
  },
  {
    "noun": "Trubel",
    "article": "der"
  },
  {
    "noun": "Trug",
    "article": "der"
  },
  {
    "noun": "Trugschluss",
    "article": "der"
  },
  {
    "noun": "Trumpf",
    "article": "der"
  },
  {
    "noun": "Trunk",
    "article": "der"
  },
  {
    "noun": "Trunkenbold",
    "article": "der"
  },
  {
    "noun": "Trupp",
    "article": "der"
  },
  {
    "noun": "Truppenabzug",
    "article": "der"
  },
  {
    "noun": "Truppenteil",
    "article": "der"
  },
  {
    "noun": "Truppentransporter",
    "article": "der"
  },
  {
    "noun": "Trust",
    "article": "der"
  },
  {
    "noun": "Truthahn",
    "article": "der"
  },
  {
    "noun": "Trutz",
    "article": "der"
  },
  {
    "noun": "Tschad",
    "article": "der"
  },
  {
    "noun": "Tschador",
    "article": "der"
  },
  {
    "noun": "Tschako",
    "article": "der"
  },
  {
    "noun": "Tscheche",
    "article": "der"
  },
  {
    "noun": "Tscherkesse",
    "article": "der"
  },
  {
    "noun": "Tscherokese",
    "article": "der"
  },
  {
    "noun": "Tschetschene",
    "article": "der"
  },
  {
    "noun": "Tschibuk",
    "article": "der"
  },
  {
    "noun": "Tsunami",
    "article": "der"
  },
  {
    "noun": "Tuberkel",
    "article": "der"
  },
  {
    "noun": "Tuberkelbazillus",
    "article": "der"
  },
  {
    "noun": "Tubus",
    "article": "der"
  },
  {
    "noun": "Tuchmacher",
    "article": "der"
  },
  {
    "noun": "Tuff",
    "article": "der"
  },
  {
    "noun": "Tuffstein",
    "article": "der"
  },
  {
    "noun": "Tugendbold",
    "article": "der"
  },
  {
    "noun": "Tukan",
    "article": "der"
  },
  {
    "noun": "Tulpenbaum",
    "article": "der"
  },
  {
    "noun": "Tummelplatz",
    "article": "der"
  },
  {
    "noun": "Tumor",
    "article": "der"
  },
  {
    "noun": "Tumormarker",
    "article": "der"
  },
  {
    "noun": "Tumult",
    "article": "der"
  },
  {
    "noun": "Tuner",
    "article": "der"
  },
  {
    "noun": "Tunesier",
    "article": "der"
  },
  {
    "noun": "Tunichtgut",
    "article": "der"
  },
  {
    "noun": "Tunnel",
    "article": "der"
  },
  {
    "noun": "Tunnelbau",
    "article": "der"
  },
  {
    "noun": "Tupfer",
    "article": "der"
  },
  {
    "noun": "Turban",
    "article": "der"
  },
  {
    "noun": "Turbogenerator",
    "article": "der"
  },
  {
    "noun": "Turbokompressor",
    "article": "der"
  },
  {
    "noun": "Turbolader",
    "article": "der"
  },
  {
    "noun": "Turbomotor",
    "article": "der"
  },
  {
    "noun": "Turf",
    "article": "der"
  },
  {
    "noun": "Turkologe",
    "article": "der"
  },
  {
    "noun": "Turm",
    "article": "der"
  },
  {
    "noun": "Turmalin",
    "article": "der"
  },
  {
    "noun": "Turmbau",
    "article": "der"
  },
  {
    "noun": "Turmdrehkran",
    "article": "der"
  },
  {
    "noun": "Turmfalke",
    "article": "der"
  },
  {
    "noun": "Turmknopf",
    "article": "der"
  },
  {
    "noun": "Turnanzug",
    "article": "der"
  },
  {
    "noun": "Turner",
    "article": "der"
  },
  {
    "noun": "Turnierreiter",
    "article": "der"
  },
  {
    "noun": "Turnierspieler",
    "article": "der"
  },
  {
    "noun": "Turniertanz",
    "article": "der"
  },
  {
    "noun": "Turnlehrer",
    "article": "der"
  },
  {
    "noun": "Turnplatz",
    "article": "der"
  },
  {
    "noun": "Turnsaal",
    "article": "der"
  },
  {
    "noun": "Turnschuh",
    "article": "der"
  },
  {
    "noun": "Turnunterricht",
    "article": "der"
  },
  {
    "noun": "Turnus",
    "article": "der"
  },
  {
    "noun": "Turnverein",
    "article": "der"
  },
  {
    "noun": "Tusch",
    "article": "der"
  },
  {
    "noun": "Tuschkasten",
    "article": "der"
  },
  {
    "noun": "Tutor",
    "article": "der"
  },
  {
    "noun": "Tweed",
    "article": "der"
  },
  {
    "noun": "Twist",
    "article": "der"
  },
  {
    "noun": "Typ",
    "article": "der"
  },
  {
    "noun": "Typensatz",
    "article": "der"
  },
  {
    "noun": "Typhus",
    "article": "der"
  },
  {
    "noun": "Typograf",
    "article": "der"
  },
  {
    "noun": "Typus",
    "article": "der"
  },
  {
    "noun": "Tyrann",
    "article": "der"
  },
  {
    "noun": "Ganzton",
    "article": "der"
  },
  {
    "noun": "Garant",
    "article": "der"
  },
  {
    "noun": "Garantieanspruch",
    "article": "der"
  },
  {
    "noun": "Garantielohn",
    "article": "der"
  },
  {
    "noun": "Garantieschein",
    "article": "der"
  },
  {
    "noun": "Garderobenschrank",
    "article": "der"
  },
  {
    "noun": "Garderobier",
    "article": "der"
  },
  {
    "noun": "Gardesoldat",
    "article": "der"
  },
  {
    "noun": "Gardist",
    "article": "der"
  },
  {
    "noun": "Garkoch",
    "article": "der"
  },
  {
    "noun": "Garnierit",
    "article": "der"
  },
  {
    "noun": "Garten",
    "article": "der"
  },
  {
    "noun": "Gartenarchitekt",
    "article": "der"
  },
  {
    "noun": "Gartenbau",
    "article": "der"
  },
  {
    "noun": "Gartenmeister",
    "article": "der"
  },
  {
    "noun": "Gartenrotschwanz",
    "article": "der"
  },
  {
    "noun": "Gartenschirm",
    "article": "der"
  },
  {
    "noun": "Gartenschlauch",
    "article": "der"
  },
  {
    "noun": "Gartenstuhl",
    "article": "der"
  },
  {
    "noun": "Gartenweg",
    "article": "der"
  },
  {
    "noun": "Gartenzaun",
    "article": "der"
  },
  {
    "noun": "Gartenzwerg",
    "article": "der"
  },
  {
    "noun": "Gasabzug",
    "article": "der"
  },
  {
    "noun": "Gasangriff",
    "article": "der"
  },
  {
    "noun": "Gasanschluss",
    "article": "der"
  },
  {
    "noun": "Gasaustritt",
    "article": "der"
  },
  {
    "noun": "Gasbeton",
    "article": "der"
  },
  {
    "noun": "Gasbrenner",
    "article": "der"
  },
  {
    "noun": "Gasdruck",
    "article": "der"
  },
  {
    "noun": "Gasgehalt",
    "article": "der"
  },
  {
    "noun": "Gasgenerator",
    "article": "der"
  },
  {
    "noun": "Gashahn",
    "article": "der"
  },
  {
    "noun": "Gashebel",
    "article": "der"
  },
  {
    "noun": "Gasherd",
    "article": "der"
  },
  {
    "noun": "Gaskessel",
    "article": "der"
  },
  {
    "noun": "Gaskocher",
    "article": "der"
  },
  {
    "noun": "Gasmotor",
    "article": "der"
  },
  {
    "noun": "Gasofen",
    "article": "der"
  },
  {
    "noun": "Gasometer",
    "article": "der"
  },
  {
    "noun": "Gasraum",
    "article": "der"
  },
  {
    "noun": "Gasschutz",
    "article": "der"
  },
  {
    "noun": "Gassenhauer",
    "article": "der"
  },
  {
    "noun": "Gassenjunge",
    "article": "der"
  },
  {
    "noun": "Gassenwitz",
    "article": "der"
  },
  {
    "noun": "Gasstrahl",
    "article": "der"
  },
  {
    "noun": "Gast",
    "article": "der"
  },
  {
    "noun": "Gastarbeiter",
    "article": "der"
  },
  {
    "noun": "Gastdozent",
    "article": "der"
  },
  {
    "noun": "Gastfreund",
    "article": "der"
  },
  {
    "noun": "Gastgeber",
    "article": "der"
  },
  {
    "noun": "Gasthof",
    "article": "der"
  },
  {
    "noun": "Gastprofessor",
    "article": "der"
  },
  {
    "noun": "Gastronom",
    "article": "der"
  },
  {
    "noun": "Gastropode",
    "article": "der"
  },
  {
    "noun": "Gastwirt",
    "article": "der"
  },
  {
    "noun": "Gasverbrauch",
    "article": "der"
  },
  {
    "noun": "Gaswascher",
    "article": "der"
  },
  {
    "noun": "Gattungsbegriff",
    "article": "der"
  },
  {
    "noun": "Gattungskauf",
    "article": "der"
  },
  {
    "noun": "Gattungsname",
    "article": "der"
  },
  {
    "noun": "Gauch",
    "article": "der"
  },
  {
    "noun": "Gaukler",
    "article": "der"
  },
  {
    "noun": "Gaul",
    "article": "der"
  },
  {
    "noun": "Gauleiter",
    "article": "der"
  },
  {
    "noun": "Gaumenbogen",
    "article": "der"
  },
  {
    "noun": "Gaumenkitzel",
    "article": "der"
  },
  {
    "noun": "Gaumenlaut",
    "article": "der"
  },
  {
    "noun": "Gauner",
    "article": "der"
  },
  {
    "noun": "Gaunerstreich",
    "article": "der"
  },
  {
    "noun": "Geber",
    "article": "der"
  },
  {
    "noun": "Gebetsmantel",
    "article": "der"
  },
  {
    "noun": "Gebetsteppich",
    "article": "der"
  },
  {
    "noun": "Gebieter",
    "article": "der"
  },
  {
    "noun": "Gebietsanspruch",
    "article": "der"
  },
  {
    "noun": "Gebirgler",
    "article": "der"
  },
  {
    "noun": "Gebirgsbach",
    "article": "der"
  },
  {
    "noun": "Gebirgsbewohner",
    "article": "der"
  },
  {
    "noun": "Gebirgsgrat",
    "article": "der"
  },
  {
    "noun": "Gebirgskamm",
    "article": "der"
  },
  {
    "noun": "Gebirgspass",
    "article": "der"
  },
  {
    "noun": "Gebirgsschlag",
    "article": "der"
  },
  {
    "noun": "Gebirgsstock",
    "article": "der"
  },
  {
    "noun": "Gebirgszug",
    "article": "der"
  },
  {
    "noun": "Gebissabdruck",
    "article": "der"
  },
  {
    "noun": "Gebrauch",
    "article": "der"
  },
  {
    "noun": "Gebrauchsartikel",
    "article": "der"
  },
  {
    "noun": "Gebrauchsgegenstand",
    "article": "der"
  },
  {
    "noun": "Gebrauchsmusterschutz",
    "article": "der"
  },
  {
    "noun": "Gebrauchswert",
    "article": "der"
  },
  {
    "noun": "Gebrauchtwagen",
    "article": "der"
  },
  {
    "noun": "Gebrauchtwagenhandel",
    "article": "der"
  },
  {
    "noun": "Gebrauchtwagenmarkt",
    "article": "der"
  },
  {
    "noun": "Gebrauchtwarenmarkt",
    "article": "der"
  },
  {
    "noun": "Geburtenzuwachs",
    "article": "der"
  },
  {
    "noun": "Geburtsfehler",
    "article": "der"
  },
  {
    "noun": "Geburtshelfer",
    "article": "der"
  },
  {
    "noun": "Geburtsjahrgang",
    "article": "der"
  },
  {
    "noun": "Geburtskanal",
    "article": "der"
  },
  {
    "noun": "Geburtsname",
    "article": "der"
  },
  {
    "noun": "Geburtsort",
    "article": "der"
  },
  {
    "noun": "Geburtsschein",
    "article": "der"
  },
  {
    "noun": "Geburtstag",
    "article": "der"
  },
  {
    "noun": "Geburtstagskuchen",
    "article": "der"
  },
  {
    "noun": "Geburtsvorgang",
    "article": "der"
  },
  {
    "noun": "Geck",
    "article": "der"
  },
  {
    "noun": "Gecko",
    "article": "der"
  },
  {
    "noun": "Gedanke",
    "article": "der"
  },
  {
    "noun": "Gedankenaustausch",
    "article": "der"
  },
  {
    "noun": "Gedankenblitz",
    "article": "der"
  },
  {
    "noun": "Gedankengang",
    "article": "der"
  },
  {
    "noun": "Gedankenkreis",
    "article": "der"
  },
  {
    "noun": "Gedankenleser",
    "article": "der"
  },
  {
    "noun": "Gedankenreichtum",
    "article": "der"
  },
  {
    "noun": "Gedankensplitter",
    "article": "der"
  },
  {
    "noun": "Gedankensprung",
    "article": "der"
  },
  {
    "noun": "Gedankenstrich",
    "article": "der"
  },
  {
    "noun": "Gedenkgottesdienst",
    "article": "der"
  },
  {
    "noun": "Gedenkstein",
    "article": "der"
  },
  {
    "noun": "Gedenktag",
    "article": "der"
  },
  {
    "noun": "Gedichtband",
    "article": "der"
  },
  {
    "noun": "Gefahrenbereich",
    "article": "der"
  },
  {
    "noun": "Gefahrenpunkt",
    "article": "der"
  },
  {
    "noun": "Gefahrguttransport",
    "article": "der"
  },
  {
    "noun": "Gefahrstoff",
    "article": "der"
  },
  {
    "noun": "Gefallenenfriedhof",
    "article": "der"
  },
  {
    "noun": "Gefangenenaufseher",
    "article": "der"
  },
  {
    "noun": "Gefangenentransport",
    "article": "der"
  },
  {
    "noun": "Gefechtskopf",
    "article": "der"
  },
  {
    "noun": "Gefechtsstand",
    "article": "der"
  },
  {
    "noun": "Gefechtsturm",
    "article": "der"
  },
  {
    "noun": "Gefolgsmann",
    "article": "der"
  },
  {
    "noun": "Gefrierpunkt",
    "article": "der"
  },
  {
    "noun": "Gefrierschrank",
    "article": "der"
  },
  {
    "noun": "Gefrierschutz",
    "article": "der"
  },
  {
    "noun": "Gegenangriff",
    "article": "der"
  },
  {
    "noun": "Gegenantrag",
    "article": "der"
  },
  {
    "noun": "Gegenbefehl",
    "article": "der"
  },
  {
    "noun": "Gegenbesuch",
    "article": "der"
  },
  {
    "noun": "Gegenbeweis",
    "article": "der"
  },
  {
    "noun": "Gegendienst",
    "article": "der"
  },
  {
    "noun": "Gegendruck",
    "article": "der"
  },
  {
    "noun": "Gegenentwurf",
    "article": "der"
  },
  {
    "noun": "Gegengrund",
    "article": "der"
  },
  {
    "noun": "Gegenkandidat",
    "article": "der"
  },
  {
    "noun": "Gegenkurs",
    "article": "der"
  },
  {
    "noun": "Gegenpart",
    "article": "der"
  },
  {
    "noun": "Gegenplan",
    "article": "der"
  },
  {
    "noun": "Gegenpol",
    "article": "der"
  },
  {
    "noun": "Gegensatz",
    "article": "der"
  },
  {
    "noun": "Gegenschlag",
    "article": "der"
  },
  {
    "noun": "Gegenspieler",
    "article": "der"
  },
  {
    "noun": "Gegensprechverkehr",
    "article": "der"
  },
  {
    "noun": "Gegenstand",
    "article": "der"
  },
  {
    "noun": "Gegenstandssatz",
    "article": "der"
  },
  {
    "noun": "Gegenstrom",
    "article": "der"
  },
  {
    "noun": "Gegenuhrzeigersinn",
    "article": "der"
  },
  {
    "noun": "Gegenverkehr",
    "article": "der"
  },
  {
    "noun": "Gegenvorschlag",
    "article": "der"
  },
  {
    "noun": "Gegenvorwurf",
    "article": "der"
  },
  {
    "noun": "Gegenwartsroman",
    "article": "der"
  },
  {
    "noun": "Gegenwert",
    "article": "der"
  },
  {
    "noun": "Gegenwind",
    "article": "der"
  },
  {
    "noun": "Gegenwohner",
    "article": "der"
  },
  {
    "noun": "Gegenzeuge",
    "article": "der"
  },
  {
    "noun": "Gegenzug",
    "article": "der"
  },
  {
    "noun": "Gegner",
    "article": "der"
  },
  {
    "noun": "Gehaltsabzug",
    "article": "der"
  },
  {
    "noun": "Gehaltsanspruch",
    "article": "der"
  },
  {
    "noun": "Gehaltstarif",
    "article": "der"
  },
  {
    "noun": "Gehaltsvorschuss",
    "article": "der"
  },
  {
    "noun": "Gehaltswunsch",
    "article": "der"
  },
  {
    "noun": "Geheimagent",
    "article": "der"
  },
  {
    "noun": "Geheimbericht",
    "article": "der"
  },
  {
    "noun": "Geheimbote",
    "article": "der"
  },
  {
    "noun": "Geheimbund",
    "article": "der"
  },
  {
    "noun": "Geheimcode",
    "article": "der"
  },
  {
    "noun": "Geheimdienst",
    "article": "der"
  },
  {
    "noun": "Geheimgang",
    "article": "der"
  },
  {
    "noun": "Geheimnisverrat",
    "article": "der"
  },
  {
    "noun": "Geheimrat",
    "article": "der"
  },
  {
    "noun": "Geheimtip",
    "article": "der"
  },
  {
    "noun": "Geheimtuer",
    "article": "der"
  },
  {
    "noun": "Gehilfe",
    "article": "der"
  },
  {
    "noun": "Gehirnnerv",
    "article": "der"
  },
  {
    "noun": "Gehirnschlag",
    "article": "der"
  },
  {
    "noun": "Gehirnschwund",
    "article": "der"
  },
  {
    "noun": "Gehirntumor",
    "article": "der"
  },
  {
    "noun": "Gehorsam",
    "article": "der"
  },
  {
    "noun": "Gehrock",
    "article": "der"
  },
  {
    "noun": "Gehsteig",
    "article": "der"
  },
  {
    "noun": "Gehstock",
    "article": "der"
  },
  {
    "noun": "Gehversuch",
    "article": "der"
  },
  {
    "noun": "Gehweg",
    "article": "der"
  },
  {
    "noun": "Geier",
    "article": "der"
  },
  {
    "noun": "Geifer",
    "article": "der"
  },
  {
    "noun": "Geigenbauer",
    "article": "der"
  },
  {
    "noun": "Geigenbogen",
    "article": "der"
  },
  {
    "noun": "Geigenkasten",
    "article": "der"
  },
  {
    "noun": "Geigenspieler",
    "article": "der"
  },
  {
    "noun": "Geiger",
    "article": "der"
  },
  {
    "noun": "Geiselnehmer",
    "article": "der"
  },
  {
    "noun": "Geiser",
    "article": "der"
  },
  {
    "noun": "Geist",
    "article": "der"
  },
  {
    "noun": "Geisterfahrer",
    "article": "der"
  },
  {
    "noun": "Geistesarbeiter",
    "article": "der"
  },
  {
    "noun": "Geistesblitz",
    "article": "der"
  },
  {
    "noun": "Geisteswissenschafter",
    "article": "der"
  },
  {
    "noun": "Geisteswissenschaftler",
    "article": "der"
  },
  {
    "noun": "Geisteszustand",
    "article": "der"
  },
  {
    "noun": "Geiz",
    "article": "der"
  },
  {
    "noun": "Geizhals",
    "article": "der"
  },
  {
    "noun": "Geizkragen",
    "article": "der"
  },
  {
    "noun": "Gelbguss",
    "article": "der"
  },
  {
    "noun": "Gelbschnabel",
    "article": "der"
  },
  {
    "noun": "Geldadel",
    "article": "der"
  },
  {
    "noun": "Geldanleger",
    "article": "der"
  },
  {
    "noun": "Geldausgabeautomat",
    "article": "der"
  },
  {
    "noun": "Geldautomat",
    "article": "der"
  },
  {
    "noun": "Geldbetrag",
    "article": "der"
  },
  {
    "noun": "Geldbeutel",
    "article": "der"
  },
  {
    "noun": "Gelderwerb",
    "article": "der"
  },
  {
    "noun": "Geldgeber",
    "article": "der"
  },
  {
    "noun": "Geldkreislauf",
    "article": "der"
  },
  {
    "noun": "Geldkurs",
    "article": "der"
  },
  {
    "noun": "Geldmangel",
    "article": "der"
  },
  {
    "noun": "Geldmarkt",
    "article": "der"
  },
  {
    "noun": "Geldpreis",
    "article": "der"
  },
  {
    "noun": "Geldsack",
    "article": "der"
  },
  {
    "noun": "Geldschein",
    "article": "der"
  },
  {
    "noun": "Geldschrank",
    "article": "der"
  },
  {
    "noun": "Geldschrankknacker",
    "article": "der"
  },
  {
    "noun": "Geldtransport",
    "article": "der"
  },
  {
    "noun": "Geldumlauf",
    "article": "der"
  },
  {
    "noun": "Geldverkehr",
    "article": "der"
  },
  {
    "noun": "Geldverleiher",
    "article": "der"
  },
  {
    "noun": "Geldverlust",
    "article": "der"
  },
  {
    "noun": "Geldwechsel",
    "article": "der"
  },
  {
    "noun": "Geldwert",
    "article": "der"
  },
  {
    "noun": "Gelegenheitsarbeiter",
    "article": "der"
  },
  {
    "noun": "Gelegenheitsdichter",
    "article": "der"
  },
  {
    "noun": "Gelegenheitsdieb",
    "article": "der"
  },
  {
    "noun": "Gelegenheitskauf",
    "article": "der"
  },
  {
    "noun": "Geleiter",
    "article": "der"
  },
  {
    "noun": "Geleitschutz",
    "article": "der"
  },
  {
    "noun": "Geleitzug",
    "article": "der"
  },
  {
    "noun": "Gelenkbus",
    "article": "der"
  },
  {
    "noun": "Gelenkknorpel",
    "article": "der"
  },
  {
    "noun": "Gelenkkopf",
    "article": "der"
  },
  {
    "noun": "Gelenkrheumatismus",
    "article": "der"
  },
  {
    "noun": "Gelenkschmerz",
    "article": "der"
  },
  {
    "noun": "Gelierzucker",
    "article": "der"
  },
  {
    "noun": "Geltungsbereich",
    "article": "der"
  },
  {
    "noun": "Geltungsdrang",
    "article": "der"
  },
  {
    "noun": "Gemahl",
    "article": "der"
  },
  {
    "noun": "Gemeindebezirk",
    "article": "der"
  },
  {
    "noun": "Gemeinderat",
    "article": "der"
  },
  {
    "noun": "Gemeindesaal",
    "article": "der"
  },
  {
    "noun": "Gemeindeverband",
    "article": "der"
  },
  {
    "noun": "Gemeindevorstand",
    "article": "der"
  },
  {
    "noun": "Gemeingebrauch",
    "article": "der"
  },
  {
    "noun": "Gemeinnutz",
    "article": "der"
  },
  {
    "noun": "Gemeinplatz",
    "article": "der"
  },
  {
    "noun": "Gemeinschaftsanschluss",
    "article": "der"
  },
  {
    "noun": "Gemeinschaftsgeist",
    "article": "der"
  },
  {
    "noun": "Gemeinschaftsraum",
    "article": "der"
  },
  {
    "noun": "Gemeinschaftssinn",
    "article": "der"
  },
  {
    "noun": "Gemeinschuldner",
    "article": "der"
  },
  {
    "noun": "Gemeinsinn",
    "article": "der"
  },
  {
    "noun": "Gemsbock",
    "article": "der"
  },
  {
    "noun": "Gendarm",
    "article": "der"
  },
  {
    "noun": "Gendefekt",
    "article": "der"
  },
  {
    "noun": "Genealoge",
    "article": "der"
  },
  {
    "noun": "Genehmigungsvorbehalt",
    "article": "der"
  },
  {
    "noun": "General",
    "article": "der"
  },
  {
    "noun": "Generaladmiral",
    "article": "der"
  },
  {
    "noun": "Generalagent",
    "article": "der"
  },
  {
    "noun": "Generalarzt",
    "article": "der"
  },
  {
    "noun": "Generalbass",
    "article": "der"
  },
  {
    "noun": "Generalbundesanwalt",
    "article": "der"
  },
  {
    "noun": "Generaldirektor",
    "article": "der"
  },
  {
    "noun": "Generalfeldmarschall",
    "article": "der"
  },
  {
    "noun": "Generalgouverneur",
    "article": "der"
  },
  {
    "noun": "Generalintendant",
    "article": "der"
  },
  {
    "noun": "Generalissimus",
    "article": "der"
  },
  {
    "noun": "Generalist",
    "article": "der"
  },
  {
    "noun": "Generalkonsul",
    "article": "der"
  },
  {
    "noun": "Generalleutnant",
    "article": "der"
  },
  {
    "noun": "Generalmajor",
    "article": "der"
  },
  {
    "noun": "Generalmusikdirektor",
    "article": "der"
  },
  {
    "noun": "Generalnenner",
    "article": "der"
  },
  {
    "noun": "Generaloberst",
    "article": "der"
  },
  {
    "noun": "Generalstaatsanwalt",
    "article": "der"
  },
  {
    "noun": "Generalstab",
    "article": "der"
  },
  {
    "noun": "Generalstreik",
    "article": "der"
  },
  {
    "noun": "Generalvertreter",
    "article": "der"
  },
  {
    "noun": "Generalvikar",
    "article": "der"
  },
  {
    "noun": "Generationenvertrag",
    "article": "der"
  },
  {
    "noun": "Generationskonflikt",
    "article": "der"
  },
  {
    "noun": "Generationsunterschied",
    "article": "der"
  },
  {
    "noun": "Generationswechsel",
    "article": "der"
  },
  {
    "noun": "Generator",
    "article": "der"
  },
  {
    "noun": "Genetiker",
    "article": "der"
  },
  {
    "noun": "Genickschuss",
    "article": "der"
  },
  {
    "noun": "Geniestreich",
    "article": "der"
  },
  {
    "noun": "Genitalbereich",
    "article": "der"
  },
  {
    "noun": "Genitiv",
    "article": "der"
  },
  {
    "noun": "Genius",
    "article": "der"
  },
  {
    "noun": "Genosse",
    "article": "der"
  },
  {
    "noun": "Genossenschaftler",
    "article": "der"
  },
  {
    "noun": "Genossenschaftsanteil",
    "article": "der"
  },
  {
    "noun": "Genossenschaftsverband",
    "article": "der"
  },
  {
    "noun": "Genotyp",
    "article": "der"
  },
  {
    "noun": "Gentleman",
    "article": "der"
  },
  {
    "noun": "Gentransfer",
    "article": "der"
  },
  {
    "noun": "Genuskauf",
    "article": "der"
  },
  {
    "noun": "Genuss",
    "article": "der"
  },
  {
    "noun": "Genussmensch",
    "article": "der"
  },
  {
    "noun": "Geograf",
    "article": "der"
  },
  {
    "noun": "Geologe",
    "article": "der"
  },
  {
    "noun": "Geometer",
    "article": "der"
  },
  {
    "noun": "Geophysiker",
    "article": "der"
  },
  {
    "noun": "Georgier",
    "article": "der"
  },
  {
    "noun": "Gepard",
    "article": "der"
  },
  {
    "noun": "Gerant",
    "article": "der"
  },
  {
    "noun": "Gerber",
    "article": "der"
  },
  {
    "noun": "Gerbstoff",
    "article": "der"
  },
  {
    "noun": "Gerechtigkeitssinn",
    "article": "der"
  },
  {
    "noun": "Gerfalke",
    "article": "der"
  },
  {
    "noun": "Geriater",
    "article": "der"
  },
  {
    "noun": "Gerichtsbescheid",
    "article": "der"
  },
  {
    "noun": "Gerichtsbeschluss",
    "article": "der"
  },
  {
    "noun": "Gerichtsdiener",
    "article": "der"
  },
  {
    "noun": "Gerichtsdolmetscher",
    "article": "der"
  },
  {
    "noun": "Gerichtsentscheid",
    "article": "der"
  },
  {
    "noun": "Gerichtshof",
    "article": "der"
  },
  {
    "noun": "Gerichtsmediziner",
    "article": "der"
  },
  {
    "noun": "Gerichtsort",
    "article": "der"
  },
  {
    "noun": "Gerichtssaal",
    "article": "der"
  },
  {
    "noun": "Gerichtsstand",
    "article": "der"
  },
  {
    "noun": "Gerichtstermin",
    "article": "der"
  },
  {
    "noun": "Gerichtsvollzieher",
    "article": "der"
  },
  {
    "noun": "Germer",
    "article": "der"
  },
  {
    "noun": "Gerneklug",
    "article": "der"
  },
  {
    "noun": "Gerontologe",
    "article": "der"
  },
  {
    "noun": "Gerstensaft",
    "article": "der"
  },
  {
    "noun": "Gerstenschrot",
    "article": "der"
  },
  {
    "noun": "Gerstenzucker",
    "article": "der"
  },
  {
    "noun": "Geruch",
    "article": "der"
  },
  {
    "noun": "Geruchsnerv",
    "article": "der"
  },
  {
    "noun": "Geruchssinn",
    "article": "der"
  },
  {
    "noun": "Geruchsstoff",
    "article": "der"
  },
  {
    "noun": "Geruchsverschluss",
    "article": "der"
  },
  {
    "noun": "Gesamtbetrag",
    "article": "der"
  },
  {
    "noun": "Gesamteindruck",
    "article": "der"
  },
  {
    "noun": "Gesamterbe",
    "article": "der"
  },
  {
    "noun": "Gesamtertrag",
    "article": "der"
  },
  {
    "noun": "Gesamtgewinn",
    "article": "der"
  },
  {
    "noun": "Gesamtkatalog",
    "article": "der"
  },
  {
    "noun": "Gesamtschaden",
    "article": "der"
  },
  {
    "noun": "Gesamtschuldner",
    "article": "der"
  },
  {
    "noun": "Gesamtsieg",
    "article": "der"
  },
  {
    "noun": "Gesamtumsatz",
    "article": "der"
  },
  {
    "noun": "Gesamtverband",
    "article": "der"
  },
  {
    "noun": "Gesamtwert",
    "article": "der"
  },
  {
    "noun": "Gesang",
    "article": "der"
  },
  {
    "noun": "Gesanglehrer",
    "article": "der"
  },
  {
    "noun": "Gesangsunterricht",
    "article": "der"
  },
  {
    "noun": "Gesangsverein",
    "article": "der"
  },
  {
    "noun": "Gesangverein",
    "article": "der"
  },
  {
    "noun": "Geschenkgutschein",
    "article": "der"
  },
  {
    "noun": "Geschenkkorb",
    "article": "der"
  },
  {
    "noun": "Geschichtsforscher",
    "article": "der"
  },
  {
    "noun": "Geschichtsschreiber",
    "article": "der"
  },
  {
    "noun": "Geschirraufzug",
    "article": "der"
  },
  {
    "noun": "Geschirrmacher",
    "article": "der"
  },
  {
    "noun": "Geschirrschrank",
    "article": "der"
  },
  {
    "noun": "Geschmack",
    "article": "der"
  },
  {
    "noun": "Geschmacksnerv",
    "article": "der"
  },
  {
    "noun": "Geschmackssinn",
    "article": "der"
  },
  {
    "noun": "Geschmacksstoff",
    "article": "der"
  },
  {
    "noun": "Geschosshagel",
    "article": "der"
  },
  {
    "noun": "Geschossmantel",
    "article": "der"
  },
  {
    "noun": "Geschwaderkommodore",
    "article": "der"
  },
  {
    "noun": "Geschwindigkeitsmesser",
    "article": "der"
  },
  {
    "noun": "Geselle",
    "article": "der"
  },
  {
    "noun": "Gesellenbrief",
    "article": "der"
  },
  {
    "noun": "Gesellschafter",
    "article": "der"
  },
  {
    "noun": "Gesellschaftsanteil",
    "article": "der"
  },
  {
    "noun": "Gesellschaftsanzug",
    "article": "der"
  },
  {
    "noun": "Gesellschaftsraum",
    "article": "der"
  },
  {
    "noun": "Gesellschaftsvertrag",
    "article": "der"
  },
  {
    "noun": "Gesetzentwurf",
    "article": "der"
  },
  {
    "noun": "Gesetzesartikel",
    "article": "der"
  },
  {
    "noun": "Gesetzesbrecher",
    "article": "der"
  },
  {
    "noun": "Gesetzesentwurf",
    "article": "der"
  },
  {
    "noun": "Gesetzgeber",
    "article": "der"
  },
  {
    "noun": "Gesichtsausdruck",
    "article": "der"
  },
  {
    "noun": "Gesichtserker",
    "article": "der"
  },
  {
    "noun": "Gesichtskrebs",
    "article": "der"
  },
  {
    "noun": "Gesichtskreis",
    "article": "der"
  },
  {
    "noun": "Gesichtsmuskel",
    "article": "der"
  },
  {
    "noun": "Gesichtsnerv",
    "article": "der"
  },
  {
    "noun": "Gesichtspuder",
    "article": "der"
  },
  {
    "noun": "Gesichtspunkt",
    "article": "der"
  },
  {
    "noun": "Gesichtssinn",
    "article": "der"
  },
  {
    "noun": "Gesichtsverlust",
    "article": "der"
  },
  {
    "noun": "Gesichtswinkel",
    "article": "der"
  },
  {
    "noun": "Gesichtszug",
    "article": "der"
  },
  {
    "noun": "Gesinnungsgenosse",
    "article": "der"
  },
  {
    "noun": "Gesinnungslump",
    "article": "der"
  },
  {
    "noun": "Gesinnungswandel",
    "article": "der"
  },
  {
    "noun": "Gesinnungswechsel",
    "article": "der"
  },
  {
    "noun": "Gespiele",
    "article": "der"
  },
  {
    "noun": "Gestalter",
    "article": "der"
  },
  {
    "noun": "Gestank",
    "article": "der"
  },
  {
    "noun": "Gestapokeller",
    "article": "der"
  },
  {
    "noun": "Gesteinsblock",
    "article": "der"
  },
  {
    "noun": "Gesteinsbohrer",
    "article": "der"
  },
  {
    "noun": "Gesteinsbrocken",
    "article": "der"
  },
  {
    "noun": "Gesuchsteller",
    "article": "der"
  },
  {
    "noun": "Gesundbrunnen",
    "article": "der"
  },
  {
    "noun": "Gesundheitsdienst",
    "article": "der"
  },
  {
    "noun": "Gesundheitsminister",
    "article": "der"
  },
  {
    "noun": "Gesundheitsschaden",
    "article": "der"
  },
  {
    "noun": "Gesundheitsschuh",
    "article": "der"
  },
  {
    "noun": "Gesundheitsschutz",
    "article": "der"
  },
  {
    "noun": "Gesundheitszustand",
    "article": "der"
  },
  {
    "noun": "Getreideanbau",
    "article": "der"
  },
  {
    "noun": "Getreidebau",
    "article": "der"
  },
  {
    "noun": "Getreideboden",
    "article": "der"
  },
  {
    "noun": "Getreidehalm",
    "article": "der"
  },
  {
    "noun": "Getreidehandel",
    "article": "der"
  },
  {
    "noun": "Getreidepreis",
    "article": "der"
  },
  {
    "noun": "Getreidespeicher",
    "article": "der"
  },
  {
    "noun": "Getriebeschaden",
    "article": "der"
  },
  {
    "noun": "Getter",
    "article": "der"
  },
  {
    "noun": "Gettoblaster",
    "article": "der"
  },
  {
    "noun": "Gewehrkolben",
    "article": "der"
  },
  {
    "noun": "Gewehrlauf",
    "article": "der"
  },
  {
    "noun": "Gewehrschuss",
    "article": "der"
  },
  {
    "noun": "Gewerbebetrieb",
    "article": "der"
  },
  {
    "noun": "Gewerbepark",
    "article": "der"
  },
  {
    "noun": "Gewerbeschein",
    "article": "der"
  },
  {
    "noun": "Gewerbeverein",
    "article": "der"
  },
  {
    "noun": "Gewerbezweig",
    "article": "der"
  },
  {
    "noun": "Gewerkschafter",
    "article": "der"
  },
  {
    "noun": "Gewerkschaftler",
    "article": "der"
  },
  {
    "noun": "Gewerkschaftsbeitrag",
    "article": "der"
  },
  {
    "noun": "Gewerkschaftsbund",
    "article": "der"
  },
  {
    "noun": "Gewerkschaftsvertreter",
    "article": "der"
  },
  {
    "noun": "Gewichtheber",
    "article": "der"
  },
  {
    "noun": "Gewichtsausgleich",
    "article": "der"
  },
  {
    "noun": "Gewichtsverlust",
    "article": "der"
  },
  {
    "noun": "Gewichtszoll",
    "article": "der"
  },
  {
    "noun": "Gewindebohrer",
    "article": "der"
  },
  {
    "noun": "Gewindeschneider",
    "article": "der"
  },
  {
    "noun": "Gewindestift",
    "article": "der"
  },
  {
    "noun": "Gewinn",
    "article": "der"
  },
  {
    "noun": "Gewinnanteil",
    "article": "der"
  },
  {
    "noun": "Gewinner",
    "article": "der"
  },
  {
    "noun": "Gewinnvortrag",
    "article": "der"
  },
  {
    "noun": "Gewissensbiss",
    "article": "der"
  },
  {
    "noun": "Gewissenskampf",
    "article": "der"
  },
  {
    "noun": "Gewissenskonflikt",
    "article": "der"
  },
  {
    "noun": "Gewissenszwang",
    "article": "der"
  },
  {
    "noun": "Gewitterregen",
    "article": "der"
  },
  {
    "noun": "Gewitterschauer",
    "article": "der"
  },
  {
    "noun": "Gewittersturm",
    "article": "der"
  },
  {
    "noun": "Gewohnheitsmensch",
    "article": "der"
  },
  {
    "noun": "Gewohnheitstrinker",
    "article": "der"
  },
  {
    "noun": "Gewohnheitsverbrecher",
    "article": "der"
  },
  {
    "noun": "Geysir",
    "article": "der"
  },
  {
    "noun": "Gezeitenwechsel",
    "article": "der"
  },
  {
    "noun": "Ghanaer",
    "article": "der"
  },
  {
    "noun": "Ghostwriter",
    "article": "der"
  },
  {
    "noun": "Gibbon",
    "article": "der"
  },
  {
    "noun": "Gichtknoten",
    "article": "der"
  },
  {
    "noun": "Giebel",
    "article": "der"
  },
  {
    "noun": "Giftbecher",
    "article": "der"
  },
  {
    "noun": "Gifthauch",
    "article": "der"
  },
  {
    "noun": "Giftmischer",
    "article": "der"
  },
  {
    "noun": "Giftpilz",
    "article": "der"
  },
  {
    "noun": "Giftstoff",
    "article": "der"
  },
  {
    "noun": "Giftzahn",
    "article": "der"
  },
  {
    "noun": "Giftzwerg",
    "article": "der"
  },
  {
    "noun": "Gig",
    "article": "der"
  },
  {
    "noun": "Gigant",
    "article": "der"
  },
  {
    "noun": "Gigantismus",
    "article": "der"
  },
  {
    "noun": "Gigolo",
    "article": "der"
  },
  {
    "noun": "Gimpel",
    "article": "der"
  },
  {
    "noun": "Gin",
    "article": "der"
  },
  {
    "noun": "Gingan",
    "article": "der"
  },
  {
    "noun": "Ginkgo",
    "article": "der"
  },
  {
    "noun": "Ginkgobaum",
    "article": "der"
  },
  {
    "noun": "Ginseng",
    "article": "der"
  },
  {
    "noun": "Ginster",
    "article": "der"
  },
  {
    "noun": "Gipfel",
    "article": "der"
  },
  {
    "noun": "Gipfelpunkt",
    "article": "der"
  },
  {
    "noun": "Gips",
    "article": "der"
  },
  {
    "noun": "Gipsabdruck",
    "article": "der"
  },
  {
    "noun": "Gipsabguss",
    "article": "der"
  },
  {
    "noun": "Gipser",
    "article": "der"
  },
  {
    "noun": "Gipskarton",
    "article": "der"
  },
  {
    "noun": "Gipsverband",
    "article": "der"
  },
  {
    "noun": "Girant",
    "article": "der"
  },
  {
    "noun": "Girat",
    "article": "der"
  },
  {
    "noun": "Giratar",
    "article": "der"
  },
  {
    "noun": "Girlitz",
    "article": "der"
  },
  {
    "noun": "Giroverkehr",
    "article": "der"
  },
  {
    "noun": "Gitarrenspieler",
    "article": "der"
  },
  {
    "noun": "Gitarrist",
    "article": "der"
  },
  {
    "noun": "Gitterbaustein",
    "article": "der"
  },
  {
    "noun": "Gittermast",
    "article": "der"
  },
  {
    "noun": "Gitterrost",
    "article": "der"
  },
  {
    "noun": "Gitterstab",
    "article": "der"
  },
  {
    "noun": "Gitterzaun",
    "article": "der"
  },
  {
    "noun": "Gladiator",
    "article": "der"
  },
  {
    "noun": "Glanz",
    "article": "der"
  },
  {
    "noun": "Glanzpunkt",
    "article": "der"
  },
  {
    "noun": "Glasballon",
    "article": "der"
  },
  {
    "noun": "Glasbaustein",
    "article": "der"
  },
  {
    "noun": "Glascontainer",
    "article": "der"
  },
  {
    "noun": "Glaser",
    "article": "der"
  },
  {
    "noun": "Glaserkitt",
    "article": "der"
  },
  {
    "noun": "Glasfaden",
    "article": "der"
  },
  {
    "noun": "Glasfluss",
    "article": "der"
  },
  {
    "noun": "Glaskasten",
    "article": "der"
  },
  {
    "noun": "Glaskolben",
    "article": "der"
  },
  {
    "noun": "Glasmaler",
    "article": "der"
  },
  {
    "noun": "Glaspalast",
    "article": "der"
  },
  {
    "noun": "Glasreiniger",
    "article": "der"
  },
  {
    "noun": "Glasschirm",
    "article": "der"
  },
  {
    "noun": "Glasschleifer",
    "article": "der"
  },
  {
    "noun": "Glasschliff",
    "article": "der"
  },
  {
    "noun": "Glasschrank",
    "article": "der"
  },
  {
    "noun": "Glassplitter",
    "article": "der"
  },
  {
    "noun": "Glassturz",
    "article": "der"
  },
  {
    "noun": "Glastisch",
    "article": "der"
  },
  {
    "noun": "Glasziegel",
    "article": "der"
  },
  {
    "noun": "Glatzkopf",
    "article": "der"
  },
  {
    "noun": "Glaube",
    "article": "der"
  },
  {
    "noun": "Glaubensartikel",
    "article": "der"
  },
  {
    "noun": "Glaubenseifer",
    "article": "der"
  },
  {
    "noun": "Glaubensgenosse",
    "article": "der"
  },
  {
    "noun": "Glaubenskampf",
    "article": "der"
  },
  {
    "noun": "Glaubenskrieg",
    "article": "der"
  },
  {
    "noun": "Glaubenssatz",
    "article": "der"
  },
  {
    "noun": "Glaubensstreit",
    "article": "der"
  },
  {
    "noun": "Glaubenswechsel",
    "article": "der"
  },
  {
    "noun": "Glaubenszweifel",
    "article": "der"
  },
  {
    "noun": "Glaukonit",
    "article": "der"
  },
  {
    "noun": "Glaukophan",
    "article": "der"
  },
  {
    "noun": "Glee",
    "article": "der"
  },
  {
    "noun": "Gleichgewichtssinn",
    "article": "der"
  },
  {
    "noun": "Gleichheitsgrundsatz",
    "article": "der"
  },
  {
    "noun": "Gleichklang",
    "article": "der"
  },
  {
    "noun": "Gleichlauf",
    "article": "der"
  },
  {
    "noun": "Gleichmacher",
    "article": "der"
  },
  {
    "noun": "Gleichmut",
    "article": "der"
  },
  {
    "noun": "Gleichrichter",
    "article": "der"
  },
  {
    "noun": "Gleichschritt",
    "article": "der"
  },
  {
    "noun": "Gleichsetzungsakkusativ",
    "article": "der"
  },
  {
    "noun": "Gleichsetzungsnominativ",
    "article": "der"
  },
  {
    "noun": "Gleichstand",
    "article": "der"
  },
  {
    "noun": "Gleichstrom",
    "article": "der"
  },
  {
    "noun": "Gleichstromgenerator",
    "article": "der"
  },
  {
    "noun": "Gleisanschluss",
    "article": "der"
  },
  {
    "noun": "Gleisbau",
    "article": "der"
  },
  {
    "noun": "Gleiter",
    "article": "der"
  },
  {
    "noun": "Gleitflug",
    "article": "der"
  },
  {
    "noun": "Gleitlaut",
    "article": "der"
  },
  {
    "noun": "Gleitschutz",
    "article": "der"
  },
  {
    "noun": "Gleitwinkel",
    "article": "der"
  },
  {
    "noun": "Gletscher",
    "article": "der"
  },
  {
    "noun": "Gletscherbach",
    "article": "der"
  },
  {
    "noun": "Gletscherschliff",
    "article": "der"
  },
  {
    "noun": "Gletschertisch",
    "article": "der"
  },
  {
    "noun": "Gletschertopf",
    "article": "der"
  },
  {
    "noun": "Gliedersatz",
    "article": "der"
  },
  {
    "noun": "Gliederzug",
    "article": "der"
  },
  {
    "noun": "Gliedsatz",
    "article": "der"
  },
  {
    "noun": "Gliedstaat",
    "article": "der"
  },
  {
    "noun": "Glimmer",
    "article": "der"
  },
  {
    "noun": "Glimmerschiefer",
    "article": "der"
  },
  {
    "noun": "Glimmstengel",
    "article": "der"
  },
  {
    "noun": "Glimmstrom",
    "article": "der"
  },
  {
    "noun": "Glitter",
    "article": "der"
  },
  {
    "noun": "Glitzer",
    "article": "der"
  },
  {
    "noun": "Globetrotter",
    "article": "der"
  },
  {
    "noun": "Globigerinenschlamm",
    "article": "der"
  },
  {
    "noun": "Globus",
    "article": "der"
  },
  {
    "noun": "Glockenguss",
    "article": "der"
  },
  {
    "noun": "Glockenrock",
    "article": "der"
  },
  {
    "noun": "Glockenschlag",
    "article": "der"
  },
  {
    "noun": "Glockenschwengel",
    "article": "der"
  },
  {
    "noun": "Glockenton",
    "article": "der"
  },
  {
    "noun": "Glockenturm",
    "article": "der"
  },
  {
    "noun": "Glossospasmus",
    "article": "der"
  },
  {
    "noun": "Glotzer",
    "article": "der"
  },
  {
    "noun": "Gnadenakt",
    "article": "der"
  },
  {
    "noun": "Gnadenbeweis",
    "article": "der"
  },
  {
    "noun": "Gnadenerlass",
    "article": "der"
  },
  {
    "noun": "Gnadenweg",
    "article": "der"
  },
  {
    "noun": "Gneis",
    "article": "der"
  },
  {
    "noun": "Gnom",
    "article": "der"
  },
  {
    "noun": "Gnomiker",
    "article": "der"
  },
  {
    "noun": "Gobelin",
    "article": "der"
  },
  {
    "noun": "Gockel",
    "article": "der"
  },
  {
    "noun": "Goi",
    "article": "der"
  },
  {
    "noun": "Goldafter",
    "article": "der"
  },
  {
    "noun": "Goldbarren",
    "article": "der"
  },
  {
    "noun": "Goldbarsch",
    "article": "der"
  },
  {
    "noun": "Goldbestand",
    "article": "der"
  },
  {
    "noun": "Goldesel",
    "article": "der"
  },
  {
    "noun": "Goldfasan",
    "article": "der"
  },
  {
    "noun": "Goldfisch",
    "article": "der"
  },
  {
    "noun": "Goldgehalt",
    "article": "der"
  },
  {
    "noun": "Goldgrund",
    "article": "der"
  },
  {
    "noun": "Goldhamster",
    "article": "der"
  },
  {
    "noun": "Goldhandel",
    "article": "der"
  },
  {
    "noun": "Goldjunge",
    "article": "der"
  },
  {
    "noun": "Goldklumpen",
    "article": "der"
  },
  {
    "noun": "Goldkurs",
    "article": "der"
  },
  {
    "noun": "Goldlack",
    "article": "der"
  },
  {
    "noun": "Goldmacher",
    "article": "der"
  },
  {
    "noun": "Goldmarkt",
    "article": "der"
  },
  {
    "noun": "Goldpreis",
    "article": "der"
  },
  {
    "noun": "Goldrahmen",
    "article": "der"
  },
  {
    "noun": "Goldrausch",
    "article": "der"
  },
  {
    "noun": "Goldregen",
    "article": "der"
  },
  {
    "noun": "Goldring",
    "article": "der"
  },
  {
    "noun": "Goldschatz",
    "article": "der"
  },
  {
    "noun": "Goldschmied",
    "article": "der"
  },
  {
    "noun": "Goldstandard",
    "article": "der"
  },
  {
    "noun": "Goldstaub",
    "article": "der"
  },
  {
    "noun": "Goldsucher",
    "article": "der"
  },
  {
    "noun": "Golfball",
    "article": "der"
  },
  {
    "noun": "Golfplatz",
    "article": "der"
  },
  {
    "noun": "Golfspieler",
    "article": "der"
  },
  {
    "noun": "Golfstaat",
    "article": "der"
  },
  {
    "noun": "Golfstrom",
    "article": "der"
  },
  {
    "noun": "Gondoliere",
    "article": "der"
  },
  {
    "noun": "Goodwill",
    "article": "der"
  },
  {
    "noun": "Gorgonzola",
    "article": "der"
  },
  {
    "noun": "Gorilla",
    "article": "der"
  },
  {
    "noun": "Gott",
    "article": "der"
  },
  {
    "noun": "Gottesacker",
    "article": "der"
  },
  {
    "noun": "Gottesbegriff",
    "article": "der"
  },
  {
    "noun": "Gottesdienst",
    "article": "der"
  },
  {
    "noun": "Gottesdienstbesucher",
    "article": "der"
  },
  {
    "noun": "Gottesglaube",
    "article": "der"
  },
  {
    "noun": "Gotteslohn",
    "article": "der"
  },
  {
    "noun": "Gottessohn",
    "article": "der"
  },
  {
    "noun": "Gottseibeiuns",
    "article": "der"
  },
  {
    "noun": "Gottvater",
    "article": "der"
  },
  {
    "noun": "Gourmand",
    "article": "der"
  },
  {
    "noun": "Gourmet",
    "article": "der"
  },
  {
    "noun": "Gouverneur",
    "article": "der"
  },
  {
    "noun": "Grabenbagger",
    "article": "der"
  },
  {
    "noun": "Grabenbruch",
    "article": "der"
  },
  {
    "noun": "Grabgesang",
    "article": "der"
  },
  {
    "noun": "Grabscher",
    "article": "der"
  },
  {
    "noun": "Grabschmuck",
    "article": "der"
  },
  {
    "noun": "Grabspruch",
    "article": "der"
  },
  {
    "noun": "Grabstein",
    "article": "der"
  },
  {
    "noun": "Grabstichel",
    "article": "der"
  },
  {
    "noun": "Gradient",
    "article": "der"
  },
  {
    "noun": "Gradmesser",
    "article": "der"
  },
  {
    "noun": "Gradunterschied",
    "article": "der"
  },
  {
    "noun": "Grafenstand",
    "article": "der"
  },
  {
    "noun": "Grafentitel",
    "article": "der"
  },
  {
    "noun": "Grafiker",
    "article": "der"
  },
  {
    "noun": "Grafit",
    "article": "der"
  },
  {
    "noun": "Grafologe",
    "article": "der"
  },
  {
    "noun": "Gral",
    "article": "der"
  },
  {
    "noun": "Gralsritter",
    "article": "der"
  },
  {
    "noun": "Grammatiker",
    "article": "der"
  },
  {
    "noun": "Granat",
    "article": "der"
  },
  {
    "noun": "Granatapfel",
    "article": "der"
  },
  {
    "noun": "Granatapfelbaum",
    "article": "der"
  },
  {
    "noun": "Granatsplitter",
    "article": "der"
  },
  {
    "noun": "Granattrichter",
    "article": "der"
  },
  {
    "noun": "Granatwerfer",
    "article": "der"
  },
  {
    "noun": "Grande",
    "article": "der"
  },
  {
    "noun": "Granit",
    "article": "der"
  },
  {
    "noun": "Granitblock",
    "article": "der"
  },
  {
    "noun": "Granitporphyr",
    "article": "der"
  },
  {
    "noun": "Granodiorit",
    "article": "der"
  },
  {
    "noun": "Grant",
    "article": "der"
  },
  {
    "noun": "Granulit",
    "article": "der"
  },
  {
    "noun": "Grapefruitsaft",
    "article": "der"
  },
  {
    "noun": "Graphitstab",
    "article": "der"
  },
  {
    "noun": "Grappa",
    "article": "der"
  },
  {
    "noun": "Grapscher",
    "article": "der"
  },
  {
    "noun": "Graptolith",
    "article": "der"
  },
  {
    "noun": "Grasboden",
    "article": "der"
  },
  {
    "noun": "Graser",
    "article": "der"
  },
  {
    "noun": "Grasfrosch",
    "article": "der"
  },
  {
    "noun": "Grashalm",
    "article": "der"
  },
  {
    "noun": "Grasplatz",
    "article": "der"
  },
  {
    "noun": "Grassame",
    "article": "der"
  },
  {
    "noun": "Grasschnitt",
    "article": "der"
  },
  {
    "noun": "Grasstreifen",
    "article": "der"
  },
  {
    "noun": "Grasteppich",
    "article": "der"
  },
  {
    "noun": "Graswuchs",
    "article": "der"
  },
  {
    "noun": "Grat",
    "article": "der"
  },
  {
    "noun": "Grathobel",
    "article": "der"
  },
  {
    "noun": "Gratulant",
    "article": "der"
  },
  {
    "noun": "Graubart",
    "article": "der"
  },
  {
    "noun": "Graufischer",
    "article": "der"
  },
  {
    "noun": "Grauguss",
    "article": "der"
  },
  {
    "noun": "Graukardinal",
    "article": "der"
  },
  {
    "noun": "Graupapagei",
    "article": "der"
  },
  {
    "noun": "Graupelregen",
    "article": "der"
  },
  {
    "noun": "Graupelschauer",
    "article": "der"
  },
  {
    "noun": "Graureiher",
    "article": "der"
  },
  {
    "noun": "Grauschimmel",
    "article": "der"
  },
  {
    "noun": "Grauschleier",
    "article": "der"
  },
  {
    "noun": "Grauspecht",
    "article": "der"
  },
  {
    "noun": "Grauton",
    "article": "der"
  },
  {
    "noun": "Grauwal",
    "article": "der"
  },
  {
    "noun": "Grauwert",
    "article": "der"
  },
  {
    "noun": "Graveur",
    "article": "der"
  },
  {
    "noun": "Gravis",
    "article": "der"
  },
  {
    "noun": "Greif",
    "article": "der"
  },
  {
    "noun": "Greifarm",
    "article": "der"
  },
  {
    "noun": "Greifer",
    "article": "der"
  },
  {
    "noun": "Greifreflex",
    "article": "der"
  },
  {
    "noun": "Greifvogel",
    "article": "der"
  },
  {
    "noun": "Greis",
    "article": "der"
  },
  {
    "noun": "Grenadier",
    "article": "der"
  },
  {
    "noun": "Grenzausgleich",
    "article": "der"
  },
  {
    "noun": "Grenzbahnhof",
    "article": "der"
  },
  {
    "noun": "Grenzbaum",
    "article": "der"
  },
  {
    "noun": "Grenzbereich",
    "article": "der"
  },
  {
    "noun": "Grenzbewohner",
    "article": "der"
  },
  {
    "noun": "Grenzbezirk",
    "article": "der"
  },
  {
    "noun": "Grenzfall",
    "article": "der"
  },
  {
    "noun": "Grenzfluss",
    "article": "der"
  },
  {
    "noun": "Grenzkonflikt",
    "article": "der"
  },
  {
    "noun": "Grenznutzen",
    "article": "der"
  },
  {
    "noun": "Grenzort",
    "article": "der"
  },
  {
    "noun": "Grenzposten",
    "article": "der"
  },
  {
    "noun": "Grenzpreis",
    "article": "der"
  },
  {
    "noun": "Grenzpunkt",
    "article": "der"
  },
  {
    "noun": "Grenzraum",
    "article": "der"
  },
  {
    "noun": "Grenzschutz",
    "article": "der"
  },
  {
    "noun": "Grenzstein",
    "article": "der"
  },
  {
    "noun": "Grenzverkehr",
    "article": "der"
  },
  {
    "noun": "Grenzwert",
    "article": "der"
  },
  {
    "noun": "Grenzzwischenfall",
    "article": "der"
  },
  {
    "noun": "Greuel",
    "article": "der"
  },
  {
    "noun": "Greyhound",
    "article": "der"
  },
  {
    "noun": "Grieche",
    "article": "der"
  },
  {
    "noun": "Griesgram",
    "article": "der"
  },
  {
    "noun": "Griff",
    "article": "der"
  },
  {
    "noun": "Griffel",
    "article": "der"
  },
  {
    "noun": "Grill",
    "article": "der"
  },
  {
    "noun": "Grillplatz",
    "article": "der"
  },
  {
    "noun": "Grillrost",
    "article": "der"
  },
  {
    "noun": "Grimassenschneider",
    "article": "der"
  },
  {
    "noun": "Grimm",
    "article": "der"
  },
  {
    "noun": "Grimmdarm",
    "article": "der"
  },
  {
    "noun": "Grind",
    "article": "der"
  },
  {
    "noun": "Gringo",
    "article": "der"
  },
  {
    "noun": "Grips",
    "article": "der"
  },
  {
    "noun": "Grison",
    "article": "der"
  },
  {
    "noun": "Grobian",
    "article": "der"
  },
  {
    "noun": "Grog",
    "article": "der"
  },
  {
    "noun": "Groll",
    "article": "der"
  },
  {
    "noun": "Groschen",
    "article": "der"
  },
  {
    "noun": "Groschenroman",
    "article": "der"
  },
  {
    "noun": "Grossular",
    "article": "der"
  },
  {
    "noun": "Grottenbau",
    "article": "der"
  },
  {
    "noun": "Grottenolm",
    "article": "der"
  },
  {
    "noun": "Grubber",
    "article": "der"
  },
  {
    "noun": "Grubenarbeiter",
    "article": "der"
  },
  {
    "noun": "Grubenbau",
    "article": "der"
  },
  {
    "noun": "Grubenbetrieb",
    "article": "der"
  },
  {
    "noun": "Grubenbrand",
    "article": "der"
  },
  {
    "noun": "Grufti",
    "article": "der"
  },
  {
    "noun": "Grund",
    "article": "der"
  },
  {
    "noun": "Grundablass",
    "article": "der"
  },
  {
    "noun": "Grundanstrich",
    "article": "der"
  },
  {
    "noun": "Grundbass",
    "article": "der"
  },
  {
    "noun": "Grundbau",
    "article": "der"
  },
  {
    "noun": "Grundbaustein",
    "article": "der"
  },
  {
    "noun": "Grundbedarf",
    "article": "der"
  },
  {
    "noun": "Grundbegriff",
    "article": "der"
  },
  {
    "noun": "Grundbesitz",
    "article": "der"
  },
  {
    "noun": "Grundbesitzer",
    "article": "der"
  },
  {
    "noun": "Grundbetrag",
    "article": "der"
  },
  {
    "noun": "Grundbucheintrag",
    "article": "der"
  },
  {
    "noun": "Grunderwerb",
    "article": "der"
  },
  {
    "noun": "Grundfreibetrag",
    "article": "der"
  },
  {
    "noun": "Grundgedanke",
    "article": "der"
  },
  {
    "noun": "Grundherr",
    "article": "der"
  },
  {
    "noun": "Grundkonsens",
    "article": "der"
  },
  {
    "noun": "Grundkurs",
    "article": "der"
  },
  {
    "noun": "Grundlohn",
    "article": "der"
  },
  {
    "noun": "Grundpfeiler",
    "article": "der"
  },
  {
    "noun": "Grundpreis",
    "article": "der"
  },
  {
    "noun": "Grundriss",
    "article": "der"
  },
  {
    "noun": "Grundsatz",
    "article": "der"
  },
  {
    "noun": "Grundstein",
    "article": "der"
  },
  {
    "noun": "Grundstock",
    "article": "der"
  },
  {
    "noun": "Grundstoff",
    "article": "der"
  },
  {
    "noun": "Grundtarif",
    "article": "der"
  },
  {
    "noun": "Grundtext",
    "article": "der"
  },
  {
    "noun": "Grundton",
    "article": "der"
  },
  {
    "noun": "Grundumsatz",
    "article": "der"
  },
  {
    "noun": "Grundwasserspiegel",
    "article": "der"
  },
  {
    "noun": "Grundwasserstand",
    "article": "der"
  },
  {
    "noun": "Grundwehrdienst",
    "article": "der"
  },
  {
    "noun": "Grundwert",
    "article": "der"
  },
  {
    "noun": "Grundwortschatz",
    "article": "der"
  },
  {
    "noun": "Grundzins",
    "article": "der"
  },
  {
    "noun": "Grundzug",
    "article": "der"
  },
  {
    "noun": "Grundzustand",
    "article": "der"
  },
  {
    "noun": "Grunge",
    "article": "der"
  },
  {
    "noun": "Grunzochse",
    "article": "der"
  },
  {
    "noun": "Gruppenakkord",
    "article": "der"
  },
  {
    "noun": "Gruppendruck",
    "article": "der"
  },
  {
    "noun": "Gruppenleiter",
    "article": "der"
  },
  {
    "noun": "Gruppenstart",
    "article": "der"
  },
  {
    "noun": "Gruppenzwang",
    "article": "der"
  },
  {
    "noun": "Gruselfilm",
    "article": "der"
  },
  {
    "noun": "Guajakbaum",
    "article": "der"
  },
  {
    "noun": "Guano",
    "article": "der"
  },
  {
    "noun": "Guatemalteke",
    "article": "der"
  },
  {
    "noun": "Gucker",
    "article": "der"
  },
  {
    "noun": "Guerilla",
    "article": "der"
  },
  {
    "noun": "Guerillakrieg",
    "article": "der"
  },
  {
    "noun": "Gugelhopf",
    "article": "der"
  },
  {
    "noun": "Gugelhupf",
    "article": "der"
  },
  {
    "noun": "Guineer",
    "article": "der"
  },
  {
    "noun": "Gulden",
    "article": "der"
  },
  {
    "noun": "Gummiball",
    "article": "der"
  },
  {
    "noun": "Gummibaum",
    "article": "der"
  },
  {
    "noun": "Gummibelag",
    "article": "der"
  },
  {
    "noun": "Gummihandschuh",
    "article": "der"
  },
  {
    "noun": "Gummilack",
    "article": "der"
  },
  {
    "noun": "Gummimantel",
    "article": "der"
  },
  {
    "noun": "Gummireifen",
    "article": "der"
  },
  {
    "noun": "Gummiring",
    "article": "der"
  },
  {
    "noun": "Gummischlauch",
    "article": "der"
  },
  {
    "noun": "Gummischuh",
    "article": "der"
  },
  {
    "noun": "Gummistiefel",
    "article": "der"
  },
  {
    "noun": "Gummistrumpf",
    "article": "der"
  },
  {
    "noun": "Gummizug",
    "article": "der"
  },
  {
    "noun": "Gundermann",
    "article": "der"
  },
  {
    "noun": "Gunstbeweis",
    "article": "der"
  },
  {
    "noun": "Guppy",
    "article": "der"
  },
  {
    "noun": "Gurkenhobel",
    "article": "der"
  },
  {
    "noun": "Gurkensalat",
    "article": "der"
  },
  {
    "noun": "Gurt",
    "article": "der"
  },
  {
    "noun": "Gurtmuffel",
    "article": "der"
  },
  {
    "noun": "Guru",
    "article": "der"
  },
  {
    "noun": "Guslar",
    "article": "der"
  },
  {
    "noun": "Guss",
    "article": "der"
  },
  {
    "noun": "Gussasphalt",
    "article": "der"
  },
  {
    "noun": "Gussbeton",
    "article": "der"
  },
  {
    "noun": "Gussregen",
    "article": "der"
  },
  {
    "noun": "Gussteil",
    "article": "der"
  },
  {
    "noun": "Gusto",
    "article": "der"
  },
  {
    "noun": "Gutachter",
    "article": "der"
  },
  {
    "noun": "Gutenachtkuss",
    "article": "der"
  },
  {
    "noun": "Gutmensch",
    "article": "der"
  },
  {
    "noun": "Gutpunkt",
    "article": "der"
  },
  {
    "noun": "Gutsbesitzer",
    "article": "der"
  },
  {
    "noun": "Gutschein",
    "article": "der"
  },
  {
    "noun": "Gutsherr",
    "article": "der"
  },
  {
    "noun": "Gutshof",
    "article": "der"
  },
  {
    "noun": "Guyaner",
    "article": "der"
  },
  {
    "noun": "Guyot",
    "article": "der"
  },
  {
    "noun": "Gymnasiast",
    "article": "der"
  },
  {
    "noun": "Gymnast",
    "article": "der"
  },
  {
    "noun": "Gymnastiker",
    "article": "der"
  },
  {
    "noun": "Panikkauf",
    "article": "der"
  },
  {
    "noun": "Panikmacher",
    "article": "der"
  },
  {
    "noun": "Pannendienst",
    "article": "der"
  },
  {
    "noun": "Pannenstreifen",
    "article": "der"
  },
  {
    "noun": "Panoramablick",
    "article": "der"
  },
  {
    "noun": "Panoramabus",
    "article": "der"
  },
  {
    "noun": "Pansen",
    "article": "der"
  },
  {
    "noun": "Panslawismus",
    "article": "der"
  },
  {
    "noun": "Panslawist",
    "article": "der"
  },
  {
    "noun": "Panter",
    "article": "der"
  },
  {
    "noun": "Pantheismus",
    "article": "der"
  },
  {
    "noun": "Pantoffel",
    "article": "der"
  },
  {
    "noun": "Pantoffelheld",
    "article": "der"
  },
  {
    "noun": "Pantophage",
    "article": "der"
  },
  {
    "noun": "Panzer",
    "article": "der"
  },
  {
    "noun": "Panzergrenadier",
    "article": "der"
  },
  {
    "noun": "Panzerkommandant",
    "article": "der"
  },
  {
    "noun": "Panzerkreuzer",
    "article": "der"
  },
  {
    "noun": "Panzerschrank",
    "article": "der"
  },
  {
    "noun": "Panzerturm",
    "article": "der"
  },
  {
    "noun": "Panzerwagen",
    "article": "der"
  },
  {
    "noun": "Papa",
    "article": "der"
  },
  {
    "noun": "Papagei",
    "article": "der"
  },
  {
    "noun": "Papageienfisch",
    "article": "der"
  },
  {
    "noun": "Papageitaucher",
    "article": "der"
  },
  {
    "noun": "Paparazzo",
    "article": "der"
  },
  {
    "noun": "Papi",
    "article": "der"
  },
  {
    "noun": "Papierbogen",
    "article": "der"
  },
  {
    "noun": "Papiereinzug",
    "article": "der"
  },
  {
    "noun": "Papierfetzen",
    "article": "der"
  },
  {
    "noun": "Papierflieger",
    "article": "der"
  },
  {
    "noun": "Papierkorb",
    "article": "der"
  },
  {
    "noun": "Papierkram",
    "article": "der"
  },
  {
    "noun": "Papierkrieg",
    "article": "der"
  },
  {
    "noun": "Papiersack",
    "article": "der"
  },
  {
    "noun": "Papierstapel",
    "article": "der"
  },
  {
    "noun": "Papierstau",
    "article": "der"
  },
  {
    "noun": "Papierstreifen",
    "article": "der"
  },
  {
    "noun": "Papiertiger",
    "article": "der"
  },
  {
    "noun": "Papierwolf",
    "article": "der"
  },
  {
    "noun": "Papismus",
    "article": "der"
  },
  {
    "noun": "Papist",
    "article": "der"
  },
  {
    "noun": "Pappband",
    "article": "der"
  },
  {
    "noun": "Pappdeckel",
    "article": "der"
  },
  {
    "noun": "Pappelbaum",
    "article": "der"
  },
  {
    "noun": "Pappendeckel",
    "article": "der"
  },
  {
    "noun": "Pappkarton",
    "article": "der"
  },
  {
    "noun": "Pappschnee",
    "article": "der"
  },
  {
    "noun": "Paprika",
    "article": "der"
  },
  {
    "noun": "Paps",
    "article": "der"
  },
  {
    "noun": "Papst",
    "article": "der"
  },
  {
    "noun": "Papua",
    "article": "der"
  },
  {
    "noun": "Papyrus",
    "article": "der"
  },
  {
    "noun": "Parabolspiegel",
    "article": "der"
  },
  {
    "noun": "Paradeiser",
    "article": "der"
  },
  {
    "noun": "Parademarsch",
    "article": "der"
  },
  {
    "noun": "Paradeschritt",
    "article": "der"
  },
  {
    "noun": "Paradiesapfel",
    "article": "der"
  },
  {
    "noun": "Paradiesvogel",
    "article": "der"
  },
  {
    "noun": "Paradigmenwechsel",
    "article": "der"
  },
  {
    "noun": "Paragraf",
    "article": "der"
  },
  {
    "noun": "Paragummi",
    "article": "der"
  },
  {
    "noun": "Parakautschuk",
    "article": "der"
  },
  {
    "noun": "Parallelismus",
    "article": "der"
  },
  {
    "noun": "Parallelkreis",
    "article": "der"
  },
  {
    "noun": "Paralytiker",
    "article": "der"
  },
  {
    "noun": "Paramagnetismus",
    "article": "der"
  },
  {
    "noun": "Parameter",
    "article": "der"
  },
  {
    "noun": "Paranoiker",
    "article": "der"
  },
  {
    "noun": "Parasit",
    "article": "der"
  },
  {
    "noun": "Parasitismus",
    "article": "der"
  },
  {
    "noun": "Paratyphus",
    "article": "der"
  },
  {
    "noun": "Parcours",
    "article": "der"
  },
  {
    "noun": "Parder",
    "article": "der"
  },
  {
    "noun": "Pariwert",
    "article": "der"
  },
  {
    "noun": "Park",
    "article": "der"
  },
  {
    "noun": "Parkausweis",
    "article": "der"
  },
  {
    "noun": "Parkettboden",
    "article": "der"
  },
  {
    "noun": "Parkettleger",
    "article": "der"
  },
  {
    "noun": "Parkinsonismus",
    "article": "der"
  },
  {
    "noun": "Parkplatz",
    "article": "der"
  },
  {
    "noun": "Parkschein",
    "article": "der"
  },
  {
    "noun": "Parkstreifen",
    "article": "der"
  },
  {
    "noun": "Parlamentarier",
    "article": "der"
  },
  {
    "noun": "Parlamentsausschuss",
    "article": "der"
  },
  {
    "noun": "Parmaschinken",
    "article": "der"
  },
  {
    "noun": "Parmesan",
    "article": "der"
  },
  {
    "noun": "Parodist",
    "article": "der"
  },
  {
    "noun": "Paroxysmus",
    "article": "der"
  },
  {
    "noun": "Part",
    "article": "der"
  },
  {
    "noun": "Parteibonze",
    "article": "der"
  },
  {
    "noun": "Parteichef",
    "article": "der"
  },
  {
    "noun": "Parteienverkehr",
    "article": "der"
  },
  {
    "noun": "Parteifreund",
    "article": "der"
  },
  {
    "noun": "Parteigenosse",
    "article": "der"
  },
  {
    "noun": "Parteikongress",
    "article": "der"
  },
  {
    "noun": "Parteisprecher",
    "article": "der"
  },
  {
    "noun": "Parteitag",
    "article": "der"
  },
  {
    "noun": "Partialbruch",
    "article": "der"
  },
  {
    "noun": "Partialton",
    "article": "der"
  },
  {
    "noun": "Partikularismus",
    "article": "der"
  },
  {
    "noun": "Partikularist",
    "article": "der"
  },
  {
    "noun": "Partisan",
    "article": "der"
  },
  {
    "noun": "Partisanenkrieg",
    "article": "der"
  },
  {
    "noun": "Partizipationsschein",
    "article": "der"
  },
  {
    "noun": "Partner",
    "article": "der"
  },
  {
    "noun": "Partnerschaftsvertrag",
    "article": "der"
  },
  {
    "noun": "Partnertausch",
    "article": "der"
  },
  {
    "noun": "Partyservice",
    "article": "der"
  },
  {
    "noun": "Pasch",
    "article": "der"
  },
  {
    "noun": "Pascher",
    "article": "der"
  },
  {
    "noun": "Pass",
    "article": "der"
  },
  {
    "noun": "Passagier",
    "article": "der"
  },
  {
    "noun": "Passagierdampfer",
    "article": "der"
  },
  {
    "noun": "Passagierverkehr",
    "article": "der"
  },
  {
    "noun": "Passant",
    "article": "der"
  },
  {
    "noun": "Passat",
    "article": "der"
  },
  {
    "noun": "Passatwind",
    "article": "der"
  },
  {
    "noun": "Passgang",
    "article": "der"
  },
  {
    "noun": "Passierschein",
    "article": "der"
  },
  {
    "noun": "Passivposten",
    "article": "der"
  },
  {
    "noun": "Passivsaldo",
    "article": "der"
  },
  {
    "noun": "Passivzins",
    "article": "der"
  },
  {
    "noun": "Passus",
    "article": "der"
  },
  {
    "noun": "Pastellmaler",
    "article": "der"
  },
  {
    "noun": "Pastellton",
    "article": "der"
  },
  {
    "noun": "Pastiche",
    "article": "der"
  },
  {
    "noun": "Pastinak",
    "article": "der"
  },
  {
    "noun": "Pastor",
    "article": "der"
  },
  {
    "noun": "Patenonkel",
    "article": "der"
  },
  {
    "noun": "Patensohn",
    "article": "der"
  },
  {
    "noun": "Patentanspruch",
    "article": "der"
  },
  {
    "noun": "Patentanwalt",
    "article": "der"
  },
  {
    "noun": "Patentinhaber",
    "article": "der"
  },
  {
    "noun": "Patentschutz",
    "article": "der"
  },
  {
    "noun": "Pater",
    "article": "der"
  },
  {
    "noun": "Paternosteraufzug",
    "article": "der"
  },
  {
    "noun": "Pathologe",
    "article": "der"
  },
  {
    "noun": "Patient",
    "article": "der"
  },
  {
    "noun": "Patriarch",
    "article": "der"
  },
  {
    "noun": "Patriot",
    "article": "der"
  },
  {
    "noun": "Patriotismus",
    "article": "der"
  },
  {
    "noun": "Patrizier",
    "article": "der"
  },
  {
    "noun": "Patzer",
    "article": "der"
  },
  {
    "noun": "Paukenschlag",
    "article": "der"
  },
  {
    "noun": "Pauker",
    "article": "der"
  },
  {
    "noun": "Paukist",
    "article": "der"
  },
  {
    "noun": "Pauperismus",
    "article": "der"
  },
  {
    "noun": "Pauschalbeitrag",
    "article": "der"
  },
  {
    "noun": "Pauschalbetrag",
    "article": "der"
  },
  {
    "noun": "Pauschalpreis",
    "article": "der"
  },
  {
    "noun": "Pauschalsatz",
    "article": "der"
  },
  {
    "noun": "Pauschalvertrag",
    "article": "der"
  },
  {
    "noun": "Pauschbetrag",
    "article": "der"
  },
  {
    "noun": "Pauschsatz",
    "article": "der"
  },
  {
    "noun": "Pausenhof",
    "article": "der"
  },
  {
    "noun": "Pausenraum",
    "article": "der"
  },
  {
    "noun": "Pavian",
    "article": "der"
  },
  {
    "noun": "Pavillon",
    "article": "der"
  },
  {
    "noun": "Pazifik",
    "article": "der"
  },
  {
    "noun": "Pazifismus",
    "article": "der"
  },
  {
    "noun": "Pazifist",
    "article": "der"
  },
  {
    "noun": "Pechvogel",
    "article": "der"
  },
  {
    "noun": "Pedell",
    "article": "der"
  },
  {
    "noun": "Peer",
    "article": "der"
  },
  {
    "noun": "Pegasus",
    "article": "der"
  },
  {
    "noun": "Pegel",
    "article": "der"
  },
  {
    "noun": "Pegelstand",
    "article": "der"
  },
  {
    "noun": "Peiler",
    "article": "der"
  },
  {
    "noun": "Peiniger",
    "article": "der"
  },
  {
    "noun": "Peitschenhieb",
    "article": "der"
  },
  {
    "noun": "Peitschenknall",
    "article": "der"
  },
  {
    "noun": "Peitschenschlag",
    "article": "der"
  },
  {
    "noun": "Pekinese",
    "article": "der"
  },
  {
    "noun": "Pelikan",
    "article": "der"
  },
  {
    "noun": "Pelz",
    "article": "der"
  },
  {
    "noun": "Pelzhandel",
    "article": "der"
  },
  {
    "noun": "Pelzkragen",
    "article": "der"
  },
  {
    "noun": "Pelzmantel",
    "article": "der"
  },
  {
    "noun": "Pemphigus",
    "article": "der"
  },
  {
    "noun": "Pendelbus",
    "article": "der"
  },
  {
    "noun": "Pendelverkehr",
    "article": "der"
  },
  {
    "noun": "Pendelzug",
    "article": "der"
  },
  {
    "noun": "Pendler",
    "article": "der"
  },
  {
    "noun": "Penni",
    "article": "der"
  },
  {
    "noun": "Penny",
    "article": "der"
  },
  {
    "noun": "Pensionist",
    "article": "der"
  },
  {
    "noun": "Pensionsanspruch",
    "article": "der"
  },
  {
    "noun": "Pensionsfonds",
    "article": "der"
  },
  {
    "noun": "Pentameter",
    "article": "der"
  },
  {
    "noun": "Pep",
    "article": "der"
  },
  {
    "noun": "Peperone",
    "article": "der"
  },
  {
    "noun": "Perfektionismus",
    "article": "der"
  },
  {
    "noun": "Perfektionist",
    "article": "der"
  },
  {
    "noun": "Perforator",
    "article": "der"
  },
  {
    "noun": "Peridot",
    "article": "der"
  },
  {
    "noun": "Perinatologe",
    "article": "der"
  },
  {
    "noun": "Periodenerfolg",
    "article": "der"
  },
  {
    "noun": "Perlenglanz",
    "article": "der"
  },
  {
    "noun": "Perlenschmuck",
    "article": "der"
  },
  {
    "noun": "Perlentaucher",
    "article": "der"
  },
  {
    "noun": "Perlitguss",
    "article": "der"
  },
  {
    "noun": "Perlmuttknopf",
    "article": "der"
  },
  {
    "noun": "Perlreis",
    "article": "der"
  },
  {
    "noun": "Perlstein",
    "article": "der"
  },
  {
    "noun": "Perlwein",
    "article": "der"
  },
  {
    "noun": "Permanentmagnet",
    "article": "der"
  },
  {
    "noun": "Peronismus",
    "article": "der"
  },
  {
    "noun": "Perron",
    "article": "der"
  },
  {
    "noun": "Perser",
    "article": "der"
  },
  {
    "noun": "Perserteppich",
    "article": "der"
  },
  {
    "noun": "Perseus",
    "article": "der"
  },
  {
    "noun": "Personalabbau",
    "article": "der"
  },
  {
    "noun": "Personalaufwand",
    "article": "der"
  },
  {
    "noun": "Personalausweis",
    "article": "der"
  },
  {
    "noun": "Personalbedarf",
    "article": "der"
  },
  {
    "noun": "Personalbestand",
    "article": "der"
  },
  {
    "noun": "Personalchef",
    "article": "der"
  },
  {
    "noun": "Personalcomputer",
    "article": "der"
  },
  {
    "noun": "Personalkredit",
    "article": "der"
  },
  {
    "noun": "Personalleiter",
    "article": "der"
  },
  {
    "noun": "Personalmangel",
    "article": "der"
  },
  {
    "noun": "Personalrat",
    "article": "der"
  },
  {
    "noun": "Personalreferent",
    "article": "der"
  },
  {
    "noun": "Personalstand",
    "article": "der"
  },
  {
    "noun": "Personalvertreter",
    "article": "der"
  },
  {
    "noun": "Personalwechsel",
    "article": "der"
  },
  {
    "noun": "Personalzuwachs",
    "article": "der"
  },
  {
    "noun": "Personenaufzug",
    "article": "der"
  },
  {
    "noun": "Personendampfer",
    "article": "der"
  },
  {
    "noun": "Personenkraftwagen",
    "article": "der"
  },
  {
    "noun": "Personenkreis",
    "article": "der"
  },
  {
    "noun": "Personenkult",
    "article": "der"
  },
  {
    "noun": "Personenname",
    "article": "der"
  },
  {
    "noun": "Personenschaden",
    "article": "der"
  },
  {
    "noun": "Personenschutz",
    "article": "der"
  },
  {
    "noun": "Personenstand",
    "article": "der"
  },
  {
    "noun": "Personentransport",
    "article": "der"
  },
  {
    "noun": "Personenverkehr",
    "article": "der"
  },
  {
    "noun": "Personenwagen",
    "article": "der"
  },
  {
    "noun": "Personenzug",
    "article": "der"
  },
  {
    "noun": "Peruaner",
    "article": "der"
  },
  {
    "noun": "Perustrom",
    "article": "der"
  },
  {
    "noun": "Pessimismus",
    "article": "der"
  },
  {
    "noun": "Pessimist",
    "article": "der"
  },
  {
    "noun": "Pestgeruch",
    "article": "der"
  },
  {
    "noun": "Pesthauch",
    "article": "der"
  },
  {
    "noun": "Petrodollar",
    "article": "der"
  },
  {
    "noun": "Petrologe",
    "article": "der"
  },
  {
    "noun": "Petticoat",
    "article": "der"
  },
  {
    "noun": "Petzer",
    "article": "der"
  },
  {
    "noun": "Peyotl",
    "article": "der"
  },
  {
    "noun": "Pfad",
    "article": "der"
  },
  {
    "noun": "Pfader",
    "article": "der"
  },
  {
    "noun": "Pfadfinder",
    "article": "der"
  },
  {
    "noun": "Pfahl",
    "article": "der"
  },
  {
    "noun": "Pfahlbau",
    "article": "der"
  },
  {
    "noun": "Pfandbrief",
    "article": "der"
  },
  {
    "noun": "Pfandleiher",
    "article": "der"
  },
  {
    "noun": "Pfandschein",
    "article": "der"
  },
  {
    "noun": "Pfandschuldner",
    "article": "der"
  },
  {
    "noun": "Pfannenstiel",
    "article": "der"
  },
  {
    "noun": "Pfannkuchen",
    "article": "der"
  },
  {
    "noun": "Pfarrbezirk",
    "article": "der"
  },
  {
    "noun": "Pfarrer",
    "article": "der"
  },
  {
    "noun": "Pfarrgemeinderat",
    "article": "der"
  },
  {
    "noun": "Pfarrhof",
    "article": "der"
  },
  {
    "noun": "Pfau",
    "article": "der"
  },
  {
    "noun": "Pfauenhahn",
    "article": "der"
  },
  {
    "noun": "Pfeffer",
    "article": "der"
  },
  {
    "noun": "Pfefferkuchen",
    "article": "der"
  },
  {
    "noun": "Pfefferminztee",
    "article": "der"
  },
  {
    "noun": "Pfefferstreuer",
    "article": "der"
  },
  {
    "noun": "Pfeifenkopf",
    "article": "der"
  },
  {
    "noun": "Pfeifenraucher",
    "article": "der"
  },
  {
    "noun": "Pfeifenstopfer",
    "article": "der"
  },
  {
    "noun": "Pfeifenstrauch",
    "article": "der"
  },
  {
    "noun": "Pfeifentabak",
    "article": "der"
  },
  {
    "noun": "Pfeifer",
    "article": "der"
  },
  {
    "noun": "Pfeil",
    "article": "der"
  },
  {
    "noun": "Pfeiler",
    "article": "der"
  },
  {
    "noun": "Pfeilhecht",
    "article": "der"
  },
  {
    "noun": "Pfennig",
    "article": "der"
  },
  {
    "noun": "Pfennigabsatz",
    "article": "der"
  },
  {
    "noun": "Pfennigbetrag",
    "article": "der"
  },
  {
    "noun": "Pfennigfuchser",
    "article": "der"
  },
  {
    "noun": "Pferch",
    "article": "der"
  },
  {
    "noun": "Pferdeapfel",
    "article": "der"
  },
  {
    "noun": "Pferdedieb",
    "article": "der"
  },
  {
    "noun": "Pferdehof",
    "article": "der"
  },
  {
    "noun": "Pferdehuf",
    "article": "der"
  },
  {
    "noun": "Pferdemarkt",
    "article": "der"
  },
  {
    "noun": "Pferdemetzger",
    "article": "der"
  },
  {
    "noun": "Pferdemist",
    "article": "der"
  },
  {
    "noun": "Pferdepfleger",
    "article": "der"
  },
  {
    "noun": "Pferderennsport",
    "article": "der"
  },
  {
    "noun": "Pferdeschlachter",
    "article": "der"
  },
  {
    "noun": "Pferdeschwanz",
    "article": "der"
  },
  {
    "noun": "Pferdesport",
    "article": "der"
  },
  {
    "noun": "Pferdestall",
    "article": "der"
  },
  {
    "noun": "Pferdewagen",
    "article": "der"
  },
  {
    "noun": "Pferdewirt",
    "article": "der"
  },
  {
    "noun": "Pfiff",
    "article": "der"
  },
  {
    "noun": "Pfifferling",
    "article": "der"
  },
  {
    "noun": "Pfiffikus",
    "article": "der"
  },
  {
    "noun": "Pfingstmontag",
    "article": "der"
  },
  {
    "noun": "Pfingstochse",
    "article": "der"
  },
  {
    "noun": "Pfingstsonntag",
    "article": "der"
  },
  {
    "noun": "Pfirsich",
    "article": "der"
  },
  {
    "noun": "Pfirsichbaum",
    "article": "der"
  },
  {
    "noun": "Pflanzenanbau",
    "article": "der"
  },
  {
    "noun": "Pflanzenbau",
    "article": "der"
  },
  {
    "noun": "Pflanzenextrakt",
    "article": "der"
  },
  {
    "noun": "Pflanzenfresser",
    "article": "der"
  },
  {
    "noun": "Pflanzenschutz",
    "article": "der"
  },
  {
    "noun": "Pflanzer",
    "article": "der"
  },
  {
    "noun": "Pflasterer",
    "article": "der"
  },
  {
    "noun": "Pflastermaler",
    "article": "der"
  },
  {
    "noun": "Pflasterstein",
    "article": "der"
  },
  {
    "noun": "Pflaumenbaum",
    "article": "der"
  },
  {
    "noun": "Pflaumenkern",
    "article": "der"
  },
  {
    "noun": "Pflaumenkuchen",
    "article": "der"
  },
  {
    "noun": "Pflegedienst",
    "article": "der"
  },
  {
    "noun": "Pflegefall",
    "article": "der"
  },
  {
    "noun": "Pfleger",
    "article": "der"
  },
  {
    "noun": "Pflegesohn",
    "article": "der"
  },
  {
    "noun": "Pflegevater",
    "article": "der"
  },
  {
    "noun": "Pflegling",
    "article": "der"
  },
  {
    "noun": "Pflichtbeitrag",
    "article": "der"
  },
  {
    "noun": "Pflichtbesuch",
    "article": "der"
  },
  {
    "noun": "Pflichteifer",
    "article": "der"
  },
  {
    "noun": "Pflichtgegenstand",
    "article": "der"
  },
  {
    "noun": "Pflichtteilsanspruch",
    "article": "der"
  },
  {
    "noun": "Pflichtverteidiger",
    "article": "der"
  },
  {
    "noun": "Pflock",
    "article": "der"
  },
  {
    "noun": "Pflug",
    "article": "der"
  },
  {
    "noun": "Pflugsterz",
    "article": "der"
  },
  {
    "noun": "Pfosten",
    "article": "der"
  },
  {
    "noun": "Pfostenschuss",
    "article": "der"
  },
  {
    "noun": "Pfriem",
    "article": "der"
  },
  {
    "noun": "Pfropf",
    "article": "der"
  },
  {
    "noun": "Pfuhl",
    "article": "der"
  },
  {
    "noun": "Pfusch",
    "article": "der"
  },
  {
    "noun": "Pfuscher",
    "article": "der"
  },
  {
    "noun": "Phallus",
    "article": "der"
  },
  {
    "noun": "Phantomschmerz",
    "article": "der"
  },
  {
    "noun": "Pharmakologe",
    "article": "der"
  },
  {
    "noun": "Pharmazeut",
    "article": "der"
  },
  {
    "noun": "Phasenmesser",
    "article": "der"
  },
  {
    "noun": "Phasensprung",
    "article": "der"
  },
  {
    "noun": "Phenoplast",
    "article": "der"
  },
  {
    "noun": "Philanthrop",
    "article": "der"
  },
  {
    "noun": "Philatelist",
    "article": "der"
  },
  {
    "noun": "Philharmoniker",
    "article": "der"
  },
  {
    "noun": "Philippiner",
    "article": "der"
  },
  {
    "noun": "Philister",
    "article": "der"
  },
  {
    "noun": "Philogyn",
    "article": "der"
  },
  {
    "noun": "Philologe",
    "article": "der"
  },
  {
    "noun": "Philosemit",
    "article": "der"
  },
  {
    "noun": "Philosoph",
    "article": "der"
  },
  {
    "noun": "Philosophaster",
    "article": "der"
  },
  {
    "noun": "Phlegmatiker",
    "article": "der"
  },
  {
    "noun": "Phlox",
    "article": "der"
  },
  {
    "noun": "Phonetiker",
    "article": "der"
  },
  {
    "noun": "Phosphorit",
    "article": "der"
  },
  {
    "noun": "Phrasendrescher",
    "article": "der"
  },
  {
    "noun": "Physiker",
    "article": "der"
  },
  {
    "noun": "Physikunterricht",
    "article": "der"
  },
  {
    "noun": "Physiognomiker",
    "article": "der"
  },
  {
    "noun": "Physiologe",
    "article": "der"
  },
  {
    "noun": "Physiotherapeut",
    "article": "der"
  },
  {
    "noun": "Pianist",
    "article": "der"
  },
  {
    "noun": "Pianospieler",
    "article": "der"
  },
  {
    "noun": "Pickel",
    "article": "der"
  },
  {
    "noun": "Picknickkorb",
    "article": "der"
  },
  {
    "noun": "Piepser",
    "article": "der"
  },
  {
    "noun": "Pier",
    "article": "der"
  },
  {
    "noun": "Pierrot",
    "article": "der"
  },
  {
    "noun": "Pigmentfleck",
    "article": "der"
  },
  {
    "noun": "Pigmentschwund",
    "article": "der"
  },
  {
    "noun": "Pilaster",
    "article": "der"
  },
  {
    "noun": "Pilau",
    "article": "der"
  },
  {
    "noun": "Pilchard",
    "article": "der"
  },
  {
    "noun": "Pilger",
    "article": "der"
  },
  {
    "noun": "Pilgerpfad",
    "article": "der"
  },
  {
    "noun": "Pilgerweg",
    "article": "der"
  },
  {
    "noun": "Pillenknick",
    "article": "der"
  },
  {
    "noun": "Pilot",
    "article": "der"
  },
  {
    "noun": "Pilotfilm",
    "article": "der"
  },
  {
    "noun": "Pilz",
    "article": "der"
  },
  {
    "noun": "Pilzbefall",
    "article": "der"
  },
  {
    "noun": "Pilzsammler",
    "article": "der"
  },
  {
    "noun": "Pimmel",
    "article": "der"
  },
  {
    "noun": "Pimpernell",
    "article": "der"
  },
  {
    "noun": "Pincheffekt",
    "article": "der"
  },
  {
    "noun": "Pinguin",
    "article": "der"
  },
  {
    "noun": "Pinienkern",
    "article": "der"
  },
  {
    "noun": "Pinienwald",
    "article": "der"
  },
  {
    "noun": "Pinienzapfen",
    "article": "der"
  },
  {
    "noun": "Pinscher",
    "article": "der"
  },
  {
    "noun": "Pinsel",
    "article": "der"
  },
  {
    "noun": "Pinselstrich",
    "article": "der"
  },
  {
    "noun": "Pinsler",
    "article": "der"
  },
  {
    "noun": "Pionier",
    "article": "der"
  },
  {
    "noun": "Pioniergeist",
    "article": "der"
  },
  {
    "noun": "Pipelinepionier",
    "article": "der"
  },
  {
    "noun": "Pipifax",
    "article": "der"
  },
  {
    "noun": "Pips",
    "article": "der"
  },
  {
    "noun": "Piranha",
    "article": "der"
  },
  {
    "noun": "Pirat",
    "article": "der"
  },
  {
    "noun": "Piratensender",
    "article": "der"
  },
  {
    "noun": "Pirol",
    "article": "der"
  },
  {
    "noun": "Pisspott",
    "article": "der"
  },
  {
    "noun": "Pistolengriff",
    "article": "der"
  },
  {
    "noun": "Pistolenlauf",
    "article": "der"
  },
  {
    "noun": "Pithekanthropus",
    "article": "der"
  },
  {
    "noun": "Piz",
    "article": "der"
  },
  {
    "noun": "Pizzateig",
    "article": "der"
  },
  {
    "noun": "PKW",
    "article": "der"
  },
  {
    "noun": "Plafond",
    "article": "der"
  },
  {
    "noun": "Plagegeist",
    "article": "der"
  },
  {
    "noun": "Plagiator",
    "article": "der"
  },
  {
    "noun": "Plakatkleber",
    "article": "der"
  },
  {
    "noun": "Planer",
    "article": "der"
  },
  {
    "noun": "Planet",
    "article": "der"
  },
  {
    "noun": "Planetoid",
    "article": "der"
  },
  {
    "noun": "Plantagenbesitzer",
    "article": "der"
  },
  {
    "noun": "Planungsaufwand",
    "article": "der"
  },
  {
    "noun": "Planungshorizont",
    "article": "der"
  },
  {
    "noun": "Planungsleiter",
    "article": "der"
  },
  {
    "noun": "Planungszeitraum",
    "article": "der"
  },
  {
    "noun": "Planwagen",
    "article": "der"
  },
  {
    "noun": "Plasmabrenner",
    "article": "der"
  },
  {
    "noun": "Plastifikator",
    "article": "der"
  },
  {
    "noun": "Plastikbecher",
    "article": "der"
  },
  {
    "noun": "Plastikbeutel",
    "article": "der"
  },
  {
    "noun": "Plastiker",
    "article": "der"
  },
  {
    "noun": "Plastiksack",
    "article": "der"
  },
  {
    "noun": "Plastiksprengstoff",
    "article": "der"
  },
  {
    "noun": "Platindraht",
    "article": "der"
  },
  {
    "noun": "Platoniker",
    "article": "der"
  },
  {
    "noun": "Platonismus",
    "article": "der"
  },
  {
    "noun": "Plattenbau",
    "article": "der"
  },
  {
    "noun": "Plattenleger",
    "article": "der"
  },
  {
    "noun": "Plattensee",
    "article": "der"
  },
  {
    "noun": "Plattenspeicher",
    "article": "der"
  },
  {
    "noun": "Plattenspieler",
    "article": "der"
  },
  {
    "noun": "Plattenteller",
    "article": "der"
  },
  {
    "noun": "Plattenwechsler",
    "article": "der"
  },
  {
    "noun": "Plattfisch",
    "article": "der"
  },
  {
    "noun": "Plattwurm",
    "article": "der"
  },
  {
    "noun": "Platz",
    "article": "der"
  },
  {
    "noun": "Platzanweiser",
    "article": "der"
  },
  {
    "noun": "Platzbedarf",
    "article": "der"
  },
  {
    "noun": "Platzhalter",
    "article": "der"
  },
  {
    "noun": "Platzhirsch",
    "article": "der"
  },
  {
    "noun": "Platzmangel",
    "article": "der"
  },
  {
    "noun": "Platzordner",
    "article": "der"
  },
  {
    "noun": "Platzregen",
    "article": "der"
  },
  {
    "noun": "Platzverweis",
    "article": "der"
  },
  {
    "noun": "Platzwart",
    "article": "der"
  },
  {
    "noun": "Platzwechsel",
    "article": "der"
  },
  {
    "noun": "Plauderer",
    "article": "der"
  },
  {
    "noun": "Plauderton",
    "article": "der"
  },
  {
    "noun": "Plausch",
    "article": "der"
  },
  {
    "noun": "Plebejer",
    "article": "der"
  },
  {
    "noun": "Pleitegeier",
    "article": "der"
  },
  {
    "noun": "Plenarsaal",
    "article": "der"
  },
  {
    "noun": "Pleonasmus",
    "article": "der"
  },
  {
    "noun": "Plesiosaurier",
    "article": "der"
  },
  {
    "noun": "Pleuel",
    "article": "der"
  },
  {
    "noun": "Plexus",
    "article": "der"
  },
  {
    "noun": "Plisseerock",
    "article": "der"
  },
  {
    "noun": "Plot",
    "article": "der"
  },
  {
    "noun": "Plotter",
    "article": "der"
  },
  {
    "noun": "Plumps",
    "article": "der"
  },
  {
    "noun": "Plunder",
    "article": "der"
  },
  {
    "noun": "Plural",
    "article": "der"
  },
  {
    "noun": "Pluralismus",
    "article": "der"
  },
  {
    "noun": "Pluspol",
    "article": "der"
  },
  {
    "noun": "Pluto",
    "article": "der"
  },
  {
    "noun": "Plutokrat",
    "article": "der"
  },
  {
    "noun": "Pneumokokkus",
    "article": "der"
  },
  {
    "noun": "Podcast",
    "article": "der"
  },
  {
    "noun": "Podex",
    "article": "der"
  },
  {
    "noun": "Podsol",
    "article": "der"
  },
  {
    "noun": "Poet",
    "article": "der"
  },
  {
    "noun": "Point",
    "article": "der"
  },
  {
    "noun": "Pointer",
    "article": "der"
  },
  {
    "noun": "Pointillismus",
    "article": "der"
  },
  {
    "noun": "Pointillist",
    "article": "der"
  },
  {
    "noun": "Pokal",
    "article": "der"
  },
  {
    "noun": "Pokalsieger",
    "article": "der"
  },
  {
    "noun": "Pol",
    "article": "der"
  },
  {
    "noun": "Polarforscher",
    "article": "der"
  },
  {
    "noun": "Polarfuchs",
    "article": "der"
  },
  {
    "noun": "Polarisationsstrom",
    "article": "der"
  },
  {
    "noun": "Polarisator",
    "article": "der"
  },
  {
    "noun": "Polarkreis",
    "article": "der"
  },
  {
    "noun": "Polarograf",
    "article": "der"
  },
  {
    "noun": "Polarstern",
    "article": "der"
  },
  {
    "noun": "Polder",
    "article": "der"
  },
  {
    "noun": "Polier",
    "article": "der"
  },
  {
    "noun": "Polierer",
    "article": "der"
  },
  {
    "noun": "Politikaster",
    "article": "der"
  },
  {
    "noun": "Politiker",
    "article": "der"
  },
  {
    "noun": "Politologe",
    "article": "der"
  },
  {
    "noun": "Polizeidirektor",
    "article": "der"
  },
  {
    "noun": "Polizeieinsatz",
    "article": "der"
  },
  {
    "noun": "Polizeifunk",
    "article": "der"
  },
  {
    "noun": "Polizeihund",
    "article": "der"
  },
  {
    "noun": "Polizeikommandant",
    "article": "der"
  },
  {
    "noun": "Polizeikommissar",
    "article": "der"
  },
  {
    "noun": "Polizeischutz",
    "article": "der"
  },
  {
    "noun": "Polizeispitzel",
    "article": "der"
  },
  {
    "noun": "Polizeistaat",
    "article": "der"
  },
  {
    "noun": "Polizist",
    "article": "der"
  },
  {
    "noun": "Pollen",
    "article": "der"
  },
  {
    "noun": "Pollenflug",
    "article": "der"
  },
  {
    "noun": "Pollenschlauch",
    "article": "der"
  },
  {
    "noun": "Poller",
    "article": "der"
  },
  {
    "noun": "Polschuh",
    "article": "der"
  },
  {
    "noun": "Polsucher",
    "article": "der"
  },
  {
    "noun": "Polterabend",
    "article": "der"
  },
  {
    "noun": "Poltergeist",
    "article": "der"
  },
  {
    "noun": "Polyester",
    "article": "der"
  },
  {
    "noun": "Polygamist",
    "article": "der"
  },
  {
    "noun": "Polygenismus",
    "article": "der"
  },
  {
    "noun": "Polygraf",
    "article": "der"
  },
  {
    "noun": "Polyhistor",
    "article": "der"
  },
  {
    "noun": "Polymorphismus",
    "article": "der"
  },
  {
    "noun": "Polynesier",
    "article": "der"
  },
  {
    "noun": "Polyp",
    "article": "der"
  },
  {
    "noun": "Polytheismus",
    "article": "der"
  },
  {
    "noun": "Polytheist",
    "article": "der"
  },
  {
    "noun": "Polyurethanschaum",
    "article": "der"
  },
  {
    "noun": "Pomologe",
    "article": "der"
  },
  {
    "noun": "Pomp",
    "article": "der"
  },
  {
    "noun": "Pompon",
    "article": "der"
  },
  {
    "noun": "Poncho",
    "article": "der"
  },
  {
    "noun": "Pontifex",
    "article": "der"
  },
  {
    "noun": "Ponton",
    "article": "der"
  },
  {
    "noun": "Pool",
    "article": "der"
  },
  {
    "noun": "Pop",
    "article": "der"
  },
  {
    "noun": "Popanz",
    "article": "der"
  },
  {
    "noun": "Popel",
    "article": "der"
  },
  {
    "noun": "Popelin",
    "article": "der"
  },
  {
    "noun": "Popo",
    "article": "der"
  },
  {
    "noun": "Populismus",
    "article": "der"
  },
  {
    "noun": "Populist",
    "article": "der"
  },
  {
    "noun": "Porno",
    "article": "der"
  },
  {
    "noun": "Pornodarsteller",
    "article": "der"
  },
  {
    "noun": "Pornofilm",
    "article": "der"
  },
  {
    "noun": "Pornograf",
    "article": "der"
  },
  {
    "noun": "Pornostar",
    "article": "der"
  },
  {
    "noun": "Porphyr",
    "article": "der"
  },
  {
    "noun": "Porree",
    "article": "der"
  },
  {
    "noun": "Porridge",
    "article": "der"
  },
  {
    "noun": "Port",
    "article": "der"
  },
  {
    "noun": "Portalkran",
    "article": "der"
  },
  {
    "noun": "Portier",
    "article": "der"
  },
  {
    "noun": "Portraitmaler",
    "article": "der"
  },
  {
    "noun": "Portugiese",
    "article": "der"
  },
  {
    "noun": "Portwein",
    "article": "der"
  },
  {
    "noun": "Posamenter",
    "article": "der"
  },
  {
    "noun": "Posaunenchor",
    "article": "der"
  },
  {
    "noun": "Posaunist",
    "article": "der"
  },
  {
    "noun": "Poseur",
    "article": "der"
  },
  {
    "noun": "Positionswinkel",
    "article": "der"
  },
  {
    "noun": "Positivismus",
    "article": "der"
  },
  {
    "noun": "Positivist",
    "article": "der"
  },
  {
    "noun": "Postauftrag",
    "article": "der"
  },
  {
    "noun": "Postausgang",
    "article": "der"
  },
  {
    "noun": "Postbezirk",
    "article": "der"
  },
  {
    "noun": "Postbote",
    "article": "der"
  },
  {
    "noun": "Postcheck",
    "article": "der"
  },
  {
    "noun": "Postdienst",
    "article": "der"
  },
  {
    "noun": "Posteingang",
    "article": "der"
  },
  {
    "noun": "Postgirodienst",
    "article": "der"
  },
  {
    "noun": "Postillion",
    "article": "der"
  },
  {
    "noun": "Postkasten",
    "article": "der"
  },
  {
    "noun": "Postmeister",
    "article": "der"
  },
  {
    "noun": "Postsack",
    "article": "der"
  },
  {
    "noun": "Postschalter",
    "article": "der"
  },
  {
    "noun": "Postscheck",
    "article": "der"
  },
  {
    "noun": "Postscheckverkehr",
    "article": "der"
  },
  {
    "noun": "Poststempel",
    "article": "der"
  },
  {
    "noun": "Postverkehr",
    "article": "der"
  },
  {
    "noun": "Postversand",
    "article": "der"
  },
  {
    "noun": "Postwagen",
    "article": "der"
  },
  {
    "noun": "Postweg",
    "article": "der"
  },
  {
    "noun": "Postzug",
    "article": "der"
  },
  {
    "noun": "Potentat",
    "article": "der"
  },
  {
    "noun": "Pott",
    "article": "der"
  },
  {
    "noun": "Pottfisch",
    "article": "der"
  },
  {
    "noun": "Pottwal",
    "article": "der"
  },
  {
    "noun": "Powidl",
    "article": "der"
  },
  {
    "noun": "Prachtjunge",
    "article": "der"
  },
  {
    "noun": "Prachtkerl",
    "article": "der"
  },
  {
    "noun": "Pragmatiker",
    "article": "der"
  },
  {
    "noun": "Pragmatismus",
    "article": "der"
  },
  {
    "noun": "Prahler",
    "article": "der"
  },
  {
    "noun": "Prahm",
    "article": "der"
  },
  {
    "noun": "Praktikant",
    "article": "der"
  },
  {
    "noun": "Praktiker",
    "article": "der"
  },
  {
    "noun": "Praktikumsplatz",
    "article": "der"
  },
  {
    "noun": "Prallhang",
    "article": "der"
  },
  {
    "noun": "Pralltopf",
    "article": "der"
  },
  {
    "noun": "Pranger",
    "article": "der"
  },
  {
    "noun": "Prasser",
    "article": "der"
  },
  {
    "noun": "Prater",
    "article": "der"
  },
  {
    "noun": "Praxisbezug",
    "article": "der"
  },
  {
    "noun": "Prediger",
    "article": "der"
  },
  {
    "noun": "Preis",
    "article": "der"
  },
  {
    "noun": "Preisabbau",
    "article": "der"
  },
  {
    "noun": "Preisabschlag",
    "article": "der"
  },
  {
    "noun": "Preisanstieg",
    "article": "der"
  },
  {
    "noun": "Preisaufschlag",
    "article": "der"
  },
  {
    "noun": "Preisauftrieb",
    "article": "der"
  },
  {
    "noun": "Preisbrecher",
    "article": "der"
  },
  {
    "noun": "Preisdruck",
    "article": "der"
  },
  {
    "noun": "Preisindex",
    "article": "der"
  },
  {
    "noun": "Preisnachlass",
    "article": "der"
  },
  {
    "noun": "Preisrichter",
    "article": "der"
  },
  {
    "noun": "Preisschlager",
    "article": "der"
  },
  {
    "noun": "Preisstopp",
    "article": "der"
  },
  {
    "noun": "Preissturz",
    "article": "der"
  },
  {
    "noun": "Preistreiber",
    "article": "der"
  },
  {
    "noun": "Preisunterschied",
    "article": "der"
  },
  {
    "noun": "Preisverfall",
    "article": "der"
  },
  {
    "noun": "Preisvergleich",
    "article": "der"
  },
  {
    "noun": "Preisvorteil",
    "article": "der"
  },
  {
    "noun": "Preiszuschlag",
    "article": "der"
  },
  {
    "noun": "Prellbock",
    "article": "der"
  },
  {
    "noun": "Preller",
    "article": "der"
  },
  {
    "noun": "Prellstein",
    "article": "der"
  },
  {
    "noun": "Premier",
    "article": "der"
  },
  {
    "noun": "Premierminister",
    "article": "der"
  },
  {
    "noun": "Presbyterianer",
    "article": "der"
  },
  {
    "noun": "Presseausweis",
    "article": "der"
  },
  {
    "noun": "Pressebericht",
    "article": "der"
  },
  {
    "noun": "Presseberichterstatter",
    "article": "der"
  },
  {
    "noun": "Pressechef",
    "article": "der"
  },
  {
    "noun": "Pressedienst",
    "article": "der"
  },
  {
    "noun": "Presser",
    "article": "der"
  },
  {
    "noun": "Pressereferent",
    "article": "der"
  },
  {
    "noun": "Pressesprecher",
    "article": "der"
  },
  {
    "noun": "Pressevertreter",
    "article": "der"
  },
  {
    "noun": "Presskopf",
    "article": "der"
  },
  {
    "noun": "Pressluftbohrer",
    "article": "der"
  },
  {
    "noun": "Presslufthammer",
    "article": "der"
  },
  {
    "noun": "Prestigegewinn",
    "article": "der"
  },
  {
    "noun": "Prestigeverlust",
    "article": "der"
  },
  {
    "noun": "Pretest",
    "article": "der"
  },
  {
    "noun": "Priapismus",
    "article": "der"
  },
  {
    "noun": "Priel",
    "article": "der"
  },
  {
    "noun": "Priem",
    "article": "der"
  },
  {
    "noun": "Priester",
    "article": "der"
  },
  {
    "noun": "Primararzt",
    "article": "der"
  },
  {
    "noun": "Primarius",
    "article": "der"
  },
  {
    "noun": "Primas",
    "article": "der"
  },
  {
    "noun": "Primawechsel",
    "article": "der"
  },
  {
    "noun": "Primer",
    "article": "der"
  },
  {
    "noun": "Primitivismus",
    "article": "der"
  },
  {
    "noun": "Primus",
    "article": "der"
  },
  {
    "noun": "Prinz",
    "article": "der"
  },
  {
    "noun": "Prinzgemahl",
    "article": "der"
  },
  {
    "noun": "Prinzipienreiter",
    "article": "der"
  },
  {
    "noun": "Prinzipienstreit",
    "article": "der"
  },
  {
    "noun": "Prior",
    "article": "der"
  },
  {
    "noun": "Pritschenwagen",
    "article": "der"
  },
  {
    "noun": "Privatbesitz",
    "article": "der"
  },
  {
    "noun": "Privatbetrieb",
    "article": "der"
  },
  {
    "noun": "Privatdetektiv",
    "article": "der"
  },
  {
    "noun": "Privatdozent",
    "article": "der"
  },
  {
    "noun": "Privatfunk",
    "article": "der"
  },
  {
    "noun": "Privathaushalt",
    "article": "der"
  },
  {
    "noun": "Privatier",
    "article": "der"
  },
  {
    "noun": "Privatinvestor",
    "article": "der"
  },
  {
    "noun": "Privatist",
    "article": "der"
  },
  {
    "noun": "Privatkunde",
    "article": "der"
  },
  {
    "noun": "Privatlehrer",
    "article": "der"
  },
  {
    "noun": "Privatmann",
    "article": "der"
  },
  {
    "noun": "Privatpatient",
    "article": "der"
  },
  {
    "noun": "Privatsektor",
    "article": "der"
  },
  {
    "noun": "Probabilismus",
    "article": "der"
  },
  {
    "noun": "Proband",
    "article": "der"
  },
  {
    "noun": "Probeabzug",
    "article": "der"
  },
  {
    "noun": "Probealarm",
    "article": "der"
  },
  {
    "noun": "Probeauftrag",
    "article": "der"
  },
  {
    "noun": "Probebetrieb",
    "article": "der"
  },
  {
    "noun": "Probedruck",
    "article": "der"
  },
  {
    "noun": "Probeflug",
    "article": "der"
  },
  {
    "noun": "Probelauf",
    "article": "der"
  },
  {
    "noun": "Probeversuch",
    "article": "der"
  },
  {
    "noun": "Probierer",
    "article": "der"
  },
  {
    "noun": "Probierstein",
    "article": "der"
  },
  {
    "noun": "Problembereich",
    "article": "der"
  },
  {
    "noun": "Problemfall",
    "article": "der"
  },
  {
    "noun": "Problemkreis",
    "article": "der"
  },
  {
    "noun": "Prodekan",
    "article": "der"
  },
  {
    "noun": "Produktenhandel",
    "article": "der"
  },
  {
    "noun": "Produktionsablauf",
    "article": "der"
  },
  {
    "noun": "Produktionsausfall",
    "article": "der"
  },
  {
    "noun": "Produktionsbetrieb",
    "article": "der"
  },
  {
    "noun": "Produktionsengpass",
    "article": "der"
  },
  {
    "noun": "Produktionsfaktor",
    "article": "der"
  },
  {
    "noun": "Produktionsfluss",
    "article": "der"
  },
  {
    "noun": "Produktionsleiter",
    "article": "der"
  },
  {
    "noun": "Produktionsplan",
    "article": "der"
  },
  {
    "noun": "Produktionspreis",
    "article": "der"
  },
  {
    "noun": "Produktionsprozess",
    "article": "der"
  },
  {
    "noun": "Produktionswert",
    "article": "der"
  },
  {
    "noun": "Produktionszweig",
    "article": "der"
  },
  {
    "noun": "Produktmanager",
    "article": "der"
  },
  {
    "noun": "Produzent",
    "article": "der"
  },
  {
    "noun": "Profanbau",
    "article": "der"
  },
  {
    "noun": "Professional",
    "article": "der"
  },
  {
    "noun": "Professionalismus",
    "article": "der"
  },
  {
    "noun": "Professionist",
    "article": "der"
  },
  {
    "noun": "Professor",
    "article": "der"
  },
  {
    "noun": "Profi",
    "article": "der"
  },
  {
    "noun": "Profiboxer",
    "article": "der"
  },
  {
    "noun": "Profilblock",
    "article": "der"
  },
  {
    "noun": "Profilreifen",
    "article": "der"
  },
  {
    "noun": "Profilstahl",
    "article": "der"
  },
  {
    "noun": "Profispieler",
    "article": "der"
  },
  {
    "noun": "Profit",
    "article": "der"
  },
  {
    "noun": "Profitmacher",
    "article": "der"
  },
  {
    "noun": "Profos",
    "article": "der"
  },
  {
    "noun": "Prognostiker",
    "article": "der"
  },
  {
    "noun": "Programmablauf",
    "article": "der"
  },
  {
    "noun": "Programmabsturz",
    "article": "der"
  },
  {
    "noun": "Programmcode",
    "article": "der"
  },
  {
    "noun": "Programmfehler",
    "article": "der"
  },
  {
    "noun": "Programmierer",
    "article": "der"
  },
  {
    "noun": "Programmierfehler",
    "article": "der"
  },
  {
    "noun": "Programmpunkt",
    "article": "der"
  },
  {
    "noun": "Programmstart",
    "article": "der"
  },
  {
    "noun": "Progress",
    "article": "der"
  },
  {
    "noun": "Prohibitionist",
    "article": "der"
  },
  {
    "noun": "Prohibitivzoll",
    "article": "der"
  },
  {
    "noun": "Projektant",
    "article": "der"
  },
  {
    "noun": "Projektbericht",
    "article": "der"
  },
  {
    "noun": "Projektingenieur",
    "article": "der"
  },
  {
    "noun": "Projektionsapparat",
    "article": "der"
  },
  {
    "noun": "Projektionsschirm",
    "article": "der"
  },
  {
    "noun": "Projektionsstrahl",
    "article": "der"
  },
  {
    "noun": "Projektleiter",
    "article": "der"
  },
  {
    "noun": "Projektmanager",
    "article": "der"
  },
  {
    "noun": "Projektor",
    "article": "der"
  },
  {
    "noun": "Prokonsul",
    "article": "der"
  },
  {
    "noun": "Proktologe",
    "article": "der"
  },
  {
    "noun": "Prokurator",
    "article": "der"
  },
  {
    "noun": "Prokurist",
    "article": "der"
  },
  {
    "noun": "Prolaps",
    "article": "der"
  },
  {
    "noun": "Prolet",
    "article": "der"
  },
  {
    "noun": "Proletarier",
    "article": "der"
  },
  {
    "noun": "Proll",
    "article": "der"
  },
  {
    "noun": "Prolog",
    "article": "der"
  },
  {
    "noun": "Prolongationswechsel",
    "article": "der"
  },
  {
    "noun": "Promenadenweg",
    "article": "der"
  },
  {
    "noun": "Promi",
    "article": "der"
  },
  {
    "noun": "Promotor",
    "article": "der"
  },
  {
    "noun": "Propagandaapparat",
    "article": "der"
  },
  {
    "noun": "Propagandafeldzug",
    "article": "der"
  },
  {
    "noun": "Propagandaminister",
    "article": "der"
  },
  {
    "noun": "Propagandist",
    "article": "der"
  },
  {
    "noun": "Propeller",
    "article": "der"
  },
  {
    "noun": "Propellerantrieb",
    "article": "der"
  },
  {
    "noun": "Prophet",
    "article": "der"
  },
  {
    "noun": "Proportionalsatz",
    "article": "der"
  },
  {
    "noun": "Propst",
    "article": "der"
  },
  {
    "noun": "Propusk",
    "article": "der"
  },
  {
    "noun": "Prorektor",
    "article": "der"
  },
  {
    "noun": "Prosaiker",
    "article": "der"
  },
  {
    "noun": "Prosecco",
    "article": "der"
  },
  {
    "noun": "Prosektor",
    "article": "der"
  },
  {
    "noun": "Prospekt",
    "article": "der"
  },
  {
    "noun": "Prospektor",
    "article": "der"
  },
  {
    "noun": "Prostatakrebs",
    "article": "der"
  },
  {
    "noun": "Protagonist",
    "article": "der"
  },
  {
    "noun": "Protektionismus",
    "article": "der"
  },
  {
    "noun": "Protektionist",
    "article": "der"
  },
  {
    "noun": "Protektor",
    "article": "der"
  },
  {
    "noun": "Protest",
    "article": "der"
  },
  {
    "noun": "Protestant",
    "article": "der"
  },
  {
    "noun": "Protestantismus",
    "article": "der"
  },
  {
    "noun": "Protestbrief",
    "article": "der"
  },
  {
    "noun": "Protestmarsch",
    "article": "der"
  },
  {
    "noun": "Protestruf",
    "article": "der"
  },
  {
    "noun": "Protestsong",
    "article": "der"
  },
  {
    "noun": "Proteststreik",
    "article": "der"
  },
  {
    "noun": "Proteststurm",
    "article": "der"
  },
  {
    "noun": "Protonenstrahl",
    "article": "der"
  },
  {
    "noun": "Prototyp",
    "article": "der"
  },
  {
    "noun": "Protz",
    "article": "der"
  },
  {
    "noun": "Proviant",
    "article": "der"
  },
  {
    "noun": "Provinzialismus",
    "article": "der"
  },
  {
    "noun": "Provinzler",
    "article": "der"
  },
  {
    "noun": "Provisionssatz",
    "article": "der"
  },
  {
    "noun": "Provisor",
    "article": "der"
  },
  {
    "noun": "Provokateur",
    "article": "der"
  },
  {
    "noun": "Prozentpunkt",
    "article": "der"
  },
  {
    "noun": "Prozentsatz",
    "article": "der"
  },
  {
    "noun": "Prozentwert",
    "article": "der"
  },
  {
    "noun": "Prozess",
    "article": "der"
  },
  {
    "noun": "Prozessablauf",
    "article": "der"
  },
  {
    "noun": "Prozessbericht",
    "article": "der"
  },
  {
    "noun": "Prozessdampf",
    "article": "der"
  },
  {
    "noun": "Prozessgegner",
    "article": "der"
  },
  {
    "noun": "Prozessor",
    "article": "der"
  },
  {
    "noun": "Prozessrechner",
    "article": "der"
  },
  {
    "noun": "Prozessvergleich",
    "article": "der"
  },
  {
    "noun": "Prunk",
    "article": "der"
  },
  {
    "noun": "Prunkbau",
    "article": "der"
  },
  {
    "noun": "Prunksaal",
    "article": "der"
  },
  {
    "noun": "Pruritus",
    "article": "der"
  },
  {
    "noun": "Psalm",
    "article": "der"
  },
  {
    "noun": "Psalmist",
    "article": "der"
  },
  {
    "noun": "Psalter",
    "article": "der"
  },
  {
    "noun": "Pseudokrupp",
    "article": "der"
  },
  {
    "noun": "Pseudotumor",
    "article": "der"
  },
  {
    "noun": "Psychiater",
    "article": "der"
  },
  {
    "noun": "Psychoanalytiker",
    "article": "der"
  },
  {
    "noun": "Psychologe",
    "article": "der"
  },
  {
    "noun": "Psychopath",
    "article": "der"
  },
  {
    "noun": "Psychotherapeut",
    "article": "der"
  },
  {
    "noun": "Ptyalismus",
    "article": "der"
  },
  {
    "noun": "Publikumserfolg",
    "article": "der"
  },
  {
    "noun": "Publikumsliebling",
    "article": "der"
  },
  {
    "noun": "Publikumsmagnet",
    "article": "der"
  },
  {
    "noun": "Publizist",
    "article": "der"
  },
  {
    "noun": "Puck",
    "article": "der"
  },
  {
    "noun": "Pudding",
    "article": "der"
  },
  {
    "noun": "Pudel",
    "article": "der"
  },
  {
    "noun": "Puderzucker",
    "article": "der"
  },
  {
    "noun": "Pueblo",
    "article": "der"
  },
  {
    "noun": "Puffer",
    "article": "der"
  },
  {
    "noun": "Pufferspeicher",
    "article": "der"
  },
  {
    "noun": "Pufferstaat",
    "article": "der"
  },
  {
    "noun": "Puffmais",
    "article": "der"
  },
  {
    "noun": "Puffreis",
    "article": "der"
  },
  {
    "noun": "Pulk",
    "article": "der"
  },
  {
    "noun": "Pulli",
    "article": "der"
  },
  {
    "noun": "Pullover",
    "article": "der"
  },
  {
    "noun": "Pullunder",
    "article": "der"
  },
  {
    "noun": "Pulp",
    "article": "der"
  },
  {
    "noun": "Puls",
    "article": "der"
  },
  {
    "noun": "Pulsar",
    "article": "der"
  },
  {
    "noun": "Pulsator",
    "article": "der"
  },
  {
    "noun": "Pulsgenerator",
    "article": "der"
  },
  {
    "noun": "Pulsschlag",
    "article": "der"
  },
  {
    "noun": "Pulverdampf",
    "article": "der"
  },
  {
    "noun": "Pulverisator",
    "article": "der"
  },
  {
    "noun": "Pulverkaffee",
    "article": "der"
  },
  {
    "noun": "Pulverschnee",
    "article": "der"
  },
  {
    "noun": "Puma",
    "article": "der"
  },
  {
    "noun": "Pummel",
    "article": "der"
  },
  {
    "noun": "Pumpenschwengel",
    "article": "der"
  },
  {
    "noun": "Pumpernickel",
    "article": "der"
  },
  {
    "noun": "Pumps",
    "article": "der"
  },
  {
    "noun": "Punk",
    "article": "der"
  },
  {
    "noun": "Punker",
    "article": "der"
  },
  {
    "noun": "Punkrock",
    "article": "der"
  },
  {
    "noun": "Punkt",
    "article": "der"
  },
  {
    "noun": "Punktrichter",
    "article": "der"
  },
  {
    "noun": "Punktroller",
    "article": "der"
  },
  {
    "noun": "Punktsieg",
    "article": "der"
  },
  {
    "noun": "Punsch",
    "article": "der"
  },
  {
    "noun": "Puppendoktor",
    "article": "der"
  },
  {
    "noun": "Puppenspieler",
    "article": "der"
  },
  {
    "noun": "Puppenwagen",
    "article": "der"
  },
  {
    "noun": "Pups",
    "article": "der"
  },
  {
    "noun": "Pupser",
    "article": "der"
  },
  {
    "noun": "Purismus",
    "article": "der"
  },
  {
    "noun": "Purist",
    "article": "der"
  },
  {
    "noun": "Puritaner",
    "article": "der"
  },
  {
    "noun": "Puritanismus",
    "article": "der"
  },
  {
    "noun": "Purpur",
    "article": "der"
  },
  {
    "noun": "Purzel",
    "article": "der"
  },
  {
    "noun": "Purzelbaum",
    "article": "der"
  },
  {
    "noun": "Puter",
    "article": "der"
  },
  {
    "noun": "Putsch",
    "article": "der"
  },
  {
    "noun": "Putschist",
    "article": "der"
  },
  {
    "noun": "Putschversuch",
    "article": "der"
  },
  {
    "noun": "Putter",
    "article": "der"
  },
  {
    "noun": "Putto",
    "article": "der"
  },
  {
    "noun": "Putz",
    "article": "der"
  },
  {
    "noun": "Putzer",
    "article": "der"
  },
  {
    "noun": "Putzfimmel",
    "article": "der"
  },
  {
    "noun": "Putzlappen",
    "article": "der"
  },
  {
    "noun": "Putztag",
    "article": "der"
  },
  {
    "noun": "Pylon",
    "article": "der"
  },
  {
    "noun": "Pylorus",
    "article": "der"
  },
  {
    "noun": "Pyramidenstumpf",
    "article": "der"
  },
  {
    "noun": "Pyrit",
    "article": "der"
  },
  {
    "noun": "Pyromane",
    "article": "der"
  },
  {
    "noun": "Pyrophor",
    "article": "der"
  },
  {
    "noun": "Pyrotechniker",
    "article": "der"
  },
  {
    "noun": "Pyrrhichius",
    "article": "der"
  },
  {
    "noun": "Pyrrhussieg",
    "article": "der"
  },
  {
    "noun": "Python",
    "article": "der"
  },
  {
    "noun": "Bahntransport",
    "article": "der"
  },
  {
    "noun": "Baikalsee",
    "article": "der"
  },
  {
    "noun": "Baissespekulant",
    "article": "der"
  },
  {
    "noun": "Baissier",
    "article": "der"
  },
  {
    "noun": "Bajazzo",
    "article": "der"
  },
  {
    "noun": "Bajonettverschluss",
    "article": "der"
  },
  {
    "noun": "Bakkalaureus",
    "article": "der"
  },
  {
    "noun": "Bakken",
    "article": "der"
  },
  {
    "noun": "Bakterienstamm",
    "article": "der"
  },
  {
    "noun": "Bakteriologe",
    "article": "der"
  },
  {
    "noun": "Balanceakt",
    "article": "der"
  },
  {
    "noun": "Balancier",
    "article": "der"
  },
  {
    "noun": "Baldachin",
    "article": "der"
  },
  {
    "noun": "Baldrian",
    "article": "der"
  },
  {
    "noun": "Balkan",
    "article": "der"
  },
  {
    "noun": "Balkanologe",
    "article": "der"
  },
  {
    "noun": "Balken",
    "article": "der"
  },
  {
    "noun": "Balkencode",
    "article": "der"
  },
  {
    "noun": "Balkon",
    "article": "der"
  },
  {
    "noun": "Ball",
    "article": "der"
  },
  {
    "noun": "Ballast",
    "article": "der"
  },
  {
    "noun": "Ballaststoff",
    "article": "der"
  },
  {
    "noun": "Ballermann",
    "article": "der"
  },
  {
    "noun": "Balljunge",
    "article": "der"
  },
  {
    "noun": "Ballon",
    "article": "der"
  },
  {
    "noun": "Ballonfahrer",
    "article": "der"
  },
  {
    "noun": "Ballonreifen",
    "article": "der"
  },
  {
    "noun": "Ballsaal",
    "article": "der"
  },
  {
    "noun": "Ballspieler",
    "article": "der"
  },
  {
    "noun": "Ballsport",
    "article": "der"
  },
  {
    "noun": "Ballungsraum",
    "article": "der"
  },
  {
    "noun": "Ballwechsel",
    "article": "der"
  },
  {
    "noun": "Balsam",
    "article": "der"
  },
  {
    "noun": "Balte",
    "article": "der"
  },
  {
    "noun": "Bambus",
    "article": "der"
  },
  {
    "noun": "Bammel",
    "article": "der"
  },
  {
    "noun": "Bananenstecker",
    "article": "der"
  },
  {
    "noun": "Banause",
    "article": "der"
  },
  {
    "noun": "Bandagist",
    "article": "der"
  },
  {
    "noun": "Bandgenerator",
    "article": "der"
  },
  {
    "noun": "Bandit",
    "article": "der"
  },
  {
    "noun": "Bandscheibenschaden",
    "article": "der"
  },
  {
    "noun": "Bandscheibenvorfall",
    "article": "der"
  },
  {
    "noun": "Bandstahl",
    "article": "der"
  },
  {
    "noun": "Bandwurm",
    "article": "der"
  },
  {
    "noun": "Bandwurmsatz",
    "article": "der"
  },
  {
    "noun": "Bankakzept",
    "article": "der"
  },
  {
    "noun": "Bankauftrag",
    "article": "der"
  },
  {
    "noun": "Bankausweis",
    "article": "der"
  },
  {
    "noun": "Bankautomat",
    "article": "der"
  },
  {
    "noun": "Bankdirektor",
    "article": "der"
  },
  {
    "noun": "Banker",
    "article": "der"
  },
  {
    "noun": "Bankfachmann",
    "article": "der"
  },
  {
    "noun": "Bankfeiertag",
    "article": "der"
  },
  {
    "noun": "Bankhalter",
    "article": "der"
  },
  {
    "noun": "Bankier",
    "article": "der"
  },
  {
    "noun": "Bankkaufmann",
    "article": "der"
  },
  {
    "noun": "Bankkrach",
    "article": "der"
  },
  {
    "noun": "Bankkredit",
    "article": "der"
  },
  {
    "noun": "Bankkunde",
    "article": "der"
  },
  {
    "noun": "Bankomat",
    "article": "der"
  },
  {
    "noun": "Bankraub",
    "article": "der"
  },
  {
    "noun": "Bankrotteur",
    "article": "der"
  },
  {
    "noun": "Banksafe",
    "article": "der"
  },
  {
    "noun": "Bankschalter",
    "article": "der"
  },
  {
    "noun": "Bankscheck",
    "article": "der"
  },
  {
    "noun": "Bankverkehr",
    "article": "der"
  },
  {
    "noun": "Bankzins",
    "article": "der"
  },
  {
    "noun": "Bann",
    "article": "der"
  },
  {
    "noun": "Bannbruch",
    "article": "der"
  },
  {
    "noun": "Bannfluch",
    "article": "der"
  },
  {
    "noun": "Baratt",
    "article": "der"
  },
  {
    "noun": "Barbar",
    "article": "der"
  },
  {
    "noun": "Barbarismus",
    "article": "der"
  },
  {
    "noun": "Barbestand",
    "article": "der"
  },
  {
    "noun": "Barbetrag",
    "article": "der"
  },
  {
    "noun": "Barbier",
    "article": "der"
  },
  {
    "noun": "Barchent",
    "article": "der"
  },
  {
    "noun": "Barde",
    "article": "der"
  },
  {
    "noun": "Barhocker",
    "article": "der"
  },
  {
    "noun": "Bariton",
    "article": "der"
  },
  {
    "noun": "Barkarole",
    "article": "der"
  },
  {
    "noun": "Barkauf",
    "article": "der"
  },
  {
    "noun": "Barkeeper",
    "article": "der"
  },
  {
    "noun": "Barmann",
    "article": "der"
  },
  {
    "noun": "Barmixer",
    "article": "der"
  },
  {
    "noun": "Barockstil",
    "article": "der"
  },
  {
    "noun": "Barograf",
    "article": "der"
  },
  {
    "noun": "Barometerstand",
    "article": "der"
  },
  {
    "noun": "Baron",
    "article": "der"
  },
  {
    "noun": "Barothermograf",
    "article": "der"
  },
  {
    "noun": "Barrakuda",
    "article": "der"
  },
  {
    "noun": "Barren",
    "article": "der"
  },
  {
    "noun": "Barscheck",
    "article": "der"
  },
  {
    "noun": "Bart",
    "article": "der"
  },
  {
    "noun": "Bartgeier",
    "article": "der"
  },
  {
    "noun": "Bartwuchs",
    "article": "der"
  },
  {
    "noun": "Baryt",
    "article": "der"
  },
  {
    "noun": "Basalt",
    "article": "der"
  },
  {
    "noun": "Basar",
    "article": "der"
  },
  {
    "noun": "Baseball",
    "article": "der"
  },
  {
    "noun": "Basilisk",
    "article": "der"
  },
  {
    "noun": "Basiskurs",
    "article": "der"
  },
  {
    "noun": "Basiswinkel",
    "article": "der"
  },
  {
    "noun": "Baske",
    "article": "der"
  },
  {
    "noun": "Basketball",
    "article": "der"
  },
  {
    "noun": "Bass",
    "article": "der"
  },
  {
    "noun": "Bassist",
    "article": "der"
  },
  {
    "noun": "Bast",
    "article": "der"
  },
  {
    "noun": "Bastard",
    "article": "der"
  },
  {
    "noun": "Bastler",
    "article": "der"
  },
  {
    "noun": "Bataillonskommandeur",
    "article": "der"
  },
  {
    "noun": "Batholith",
    "article": "der"
  },
  {
    "noun": "Batist",
    "article": "der"
  },
  {
    "noun": "Batteriestrom",
    "article": "der"
  },
  {
    "noun": "Batzen",
    "article": "der"
  },
  {
    "noun": "Bau",
    "article": "der"
  },
  {
    "noun": "Bauabschnitt",
    "article": "der"
  },
  {
    "noun": "Bauarbeiter",
    "article": "der"
  },
  {
    "noun": "Baubeginn",
    "article": "der"
  },
  {
    "noun": "Baublock",
    "article": "der"
  },
  {
    "noun": "Bauboom",
    "article": "der"
  },
  {
    "noun": "Bauch",
    "article": "der"
  },
  {
    "noun": "Bauchansatz",
    "article": "der"
  },
  {
    "noun": "Bauchgurt",
    "article": "der"
  },
  {
    "noun": "Bauchklatscher",
    "article": "der"
  },
  {
    "noun": "Bauchladen",
    "article": "der"
  },
  {
    "noun": "Bauchmuskel",
    "article": "der"
  },
  {
    "noun": "Bauchnabel",
    "article": "der"
  },
  {
    "noun": "Bauchredner",
    "article": "der"
  },
  {
    "noun": "Bauchschmerz",
    "article": "der"
  },
  {
    "noun": "Bauchschnitt",
    "article": "der"
  },
  {
    "noun": "Bauchschuss",
    "article": "der"
  },
  {
    "noun": "Bauchspeck",
    "article": "der"
  },
  {
    "noun": "Bauchtanz",
    "article": "der"
  },
  {
    "noun": "Bauchumfang",
    "article": "der"
  },
  {
    "noun": "Bauentwurf",
    "article": "der"
  },
  {
    "noun": "Bauernaufstand",
    "article": "der"
  },
  {
    "noun": "Bauernbursche",
    "article": "der"
  },
  {
    "noun": "Bauernhof",
    "article": "der"
  },
  {
    "noun": "Bauernjunge",
    "article": "der"
  },
  {
    "noun": "Bauernstand",
    "article": "der"
  },
  {
    "noun": "Bauerntanz",
    "article": "der"
  },
  {
    "noun": "Bauerntrampel",
    "article": "der"
  },
  {
    "noun": "Bauersmann",
    "article": "der"
  },
  {
    "noun": "Baufachmann",
    "article": "der"
  },
  {
    "noun": "Baugrund",
    "article": "der"
  },
  {
    "noun": "Bauherr",
    "article": "der"
  },
  {
    "noun": "Bauingenieur",
    "article": "der"
  },
  {
    "noun": "Baukasten",
    "article": "der"
  },
  {
    "noun": "Baukostenanschlag",
    "article": "der"
  },
  {
    "noun": "Baukostenzuschuss",
    "article": "der"
  },
  {
    "noun": "Baukran",
    "article": "der"
  },
  {
    "noun": "Baukredit",
    "article": "der"
  },
  {
    "noun": "Bauleiter",
    "article": "der"
  },
  {
    "noun": "Baum",
    "article": "der"
  },
  {
    "noun": "Baumarkt",
    "article": "der"
  },
  {
    "noun": "Baumeister",
    "article": "der"
  },
  {
    "noun": "Baumfarn",
    "article": "der"
  },
  {
    "noun": "Baumgarten",
    "article": "der"
  },
  {
    "noun": "Baumkuchen",
    "article": "der"
  },
  {
    "noun": "Baummarder",
    "article": "der"
  },
  {
    "noun": "Baumsarg",
    "article": "der"
  },
  {
    "noun": "Baumstamm",
    "article": "der"
  },
  {
    "noun": "Baumstrunk",
    "article": "der"
  },
  {
    "noun": "Baumstumpf",
    "article": "der"
  },
  {
    "noun": "Baumwipfel",
    "article": "der"
  },
  {
    "noun": "Baumwollpullover",
    "article": "der"
  },
  {
    "noun": "Baumwollsame",
    "article": "der"
  },
  {
    "noun": "Baumwollstoff",
    "article": "der"
  },
  {
    "noun": "Bauplan",
    "article": "der"
  },
  {
    "noun": "Bauplatz",
    "article": "der"
  },
  {
    "noun": "Baupreis",
    "article": "der"
  },
  {
    "noun": "Baurat",
    "article": "der"
  },
  {
    "noun": "Bausatz",
    "article": "der"
  },
  {
    "noun": "Bausch",
    "article": "der"
  },
  {
    "noun": "Bauschlosser",
    "article": "der"
  },
  {
    "noun": "Bauschreiner",
    "article": "der"
  },
  {
    "noun": "Bauschutt",
    "article": "der"
  },
  {
    "noun": "Bausektor",
    "article": "der"
  },
  {
    "noun": "Bausparer",
    "article": "der"
  },
  {
    "noun": "Bausparvertrag",
    "article": "der"
  },
  {
    "noun": "Baustahl",
    "article": "der"
  },
  {
    "noun": "Baustein",
    "article": "der"
  },
  {
    "noun": "Baustil",
    "article": "der"
  },
  {
    "noun": "Baustoff",
    "article": "der"
  },
  {
    "noun": "Baustopp",
    "article": "der"
  },
  {
    "noun": "Bautischler",
    "article": "der"
  },
  {
    "noun": "Bauunternehmer",
    "article": "der"
  },
  {
    "noun": "Bauvertrag",
    "article": "der"
  },
  {
    "noun": "Bauxit",
    "article": "der"
  },
  {
    "noun": "Bauzaun",
    "article": "der"
  },
  {
    "noun": "Bauzeichner",
    "article": "der"
  },
  {
    "noun": "Bauzuschuss",
    "article": "der"
  },
  {
    "noun": "Bavaria",
    "article": "der"
  },
  {
    "noun": "Bayer",
    "article": "der"
  },
  {
    "noun": "Bazillus",
    "article": "der"
  },
  {
    "noun": "Beachvolleyball",
    "article": "der"
  },
  {
    "noun": "Beamer",
    "article": "der"
  },
  {
    "noun": "Bearbeiter",
    "article": "der"
  },
  {
    "noun": "Beat",
    "article": "der"
  },
  {
    "noun": "Beatnik",
    "article": "der"
  },
  {
    "noun": "Beau",
    "article": "der"
  },
  {
    "noun": "Bebauungsplan",
    "article": "der"
  },
  {
    "noun": "Becher",
    "article": "der"
  },
  {
    "noun": "Beckenbruch",
    "article": "der"
  },
  {
    "noun": "Beckengurt",
    "article": "der"
  },
  {
    "noun": "Beckenknochen",
    "article": "der"
  },
  {
    "noun": "Beckmesser",
    "article": "der"
  },
  {
    "noun": "Bedarf",
    "article": "der"
  },
  {
    "noun": "Bedarfsartikel",
    "article": "der"
  },
  {
    "noun": "Bedarfsfall",
    "article": "der"
  },
  {
    "noun": "Bedecktsamer",
    "article": "der"
  },
  {
    "noun": "Bedeutungsumfang",
    "article": "der"
  },
  {
    "noun": "Bedeutungswandel",
    "article": "der"
  },
  {
    "noun": "Bediener",
    "article": "der"
  },
  {
    "noun": "Bedienungsaufschlag",
    "article": "der"
  },
  {
    "noun": "Bedienungsfehler",
    "article": "der"
  },
  {
    "noun": "Bedingungssatz",
    "article": "der"
  },
  {
    "noun": "Beduine",
    "article": "der"
  },
  {
    "noun": "Befall",
    "article": "der"
  },
  {
    "noun": "Befehl",
    "article": "der"
  },
  {
    "noun": "Befehlshaber",
    "article": "der"
  },
  {
    "noun": "Befehlssatz",
    "article": "der"
  },
  {
    "noun": "Befehlston",
    "article": "der"
  },
  {
    "noun": "Befeuchter",
    "article": "der"
  },
  {
    "noun": "Befrachter",
    "article": "der"
  },
  {
    "noun": "Befreier",
    "article": "der"
  },
  {
    "noun": "Befreiungskampf",
    "article": "der"
  },
  {
    "noun": "Befreiungskrieg",
    "article": "der"
  },
  {
    "noun": "Befreiungsschlag",
    "article": "der"
  },
  {
    "noun": "Befreiungsversuch",
    "article": "der"
  },
  {
    "noun": "Befund",
    "article": "der"
  },
  {
    "noun": "Begeisterungstaumel",
    "article": "der"
  },
  {
    "noun": "Beginn",
    "article": "der"
  },
  {
    "noun": "Begleitbrief",
    "article": "der"
  },
  {
    "noun": "Begleiter",
    "article": "der"
  },
  {
    "noun": "Begleitschein",
    "article": "der"
  },
  {
    "noun": "Begleittext",
    "article": "der"
  },
  {
    "noun": "Begleitumstand",
    "article": "der"
  },
  {
    "noun": "Begrenzer",
    "article": "der"
  },
  {
    "noun": "Begriff",
    "article": "der"
  },
  {
    "noun": "Begriffsinhalt",
    "article": "der"
  },
  {
    "noun": "Begriffsumfang",
    "article": "der"
  },
  {
    "noun": "Begutachter",
    "article": "der"
  },
  {
    "noun": "Behandlungsplan",
    "article": "der"
  },
  {
    "noun": "Behandlungsraum",
    "article": "der"
  },
  {
    "noun": "Behang",
    "article": "der"
  },
  {
    "noun": "Behaviorismus",
    "article": "der"
  },
  {
    "noun": "Behelf",
    "article": "der"
  },
  {
    "noun": "Behelfsbau",
    "article": "der"
  },
  {
    "noun": "Beherbergungsbetrieb",
    "article": "der"
  },
  {
    "noun": "Beherrscher",
    "article": "der"
  },
  {
    "noun": "Behuf",
    "article": "der"
  },
  {
    "noun": "Beichtstuhl",
    "article": "der"
  },
  {
    "noun": "Beichtvater",
    "article": "der"
  },
  {
    "noun": "Beifahrer",
    "article": "der"
  },
  {
    "noun": "Beifahrersitz",
    "article": "der"
  },
  {
    "noun": "Beifall",
    "article": "der"
  },
  {
    "noun": "Beifallruf",
    "article": "der"
  },
  {
    "noun": "Beifallsruf",
    "article": "der"
  },
  {
    "noun": "Beifallssturm",
    "article": "der"
  },
  {
    "noun": "Beifang",
    "article": "der"
  },
  {
    "noun": "Beigeschmack",
    "article": "der"
  },
  {
    "noun": "Beignet",
    "article": "der"
  },
  {
    "noun": "Beiklang",
    "article": "der"
  },
  {
    "noun": "Beikoch",
    "article": "der"
  },
  {
    "noun": "Beileidsbrief",
    "article": "der"
  },
  {
    "noun": "Beiname",
    "article": "der"
  },
  {
    "noun": "Beinbruch",
    "article": "der"
  },
  {
    "noun": "Beinling",
    "article": "der"
  },
  {
    "noun": "Beinstumpf",
    "article": "der"
  },
  {
    "noun": "Beinwell",
    "article": "der"
  },
  {
    "noun": "Beipackzettel",
    "article": "der"
  },
  {
    "noun": "Beirat",
    "article": "der"
  },
  {
    "noun": "Beisatz",
    "article": "der"
  },
  {
    "noun": "Beischlaf",
    "article": "der"
  },
  {
    "noun": "Beisitzer",
    "article": "der"
  },
  {
    "noun": "Beispielsatz",
    "article": "der"
  },
  {
    "noun": "Beistand",
    "article": "der"
  },
  {
    "noun": "Beistandskredit",
    "article": "der"
  },
  {
    "noun": "Beistelltisch",
    "article": "der"
  },
  {
    "noun": "Beistrich",
    "article": "der"
  },
  {
    "noun": "Beitel",
    "article": "der"
  },
  {
    "noun": "Beitrag",
    "article": "der"
  },
  {
    "noun": "Beitragssatz",
    "article": "der"
  },
  {
    "noun": "Beitritt",
    "article": "der"
  },
  {
    "noun": "Beitrittsantrag",
    "article": "der"
  },
  {
    "noun": "Beitrittsgegner",
    "article": "der"
  },
  {
    "noun": "Beiwagen",
    "article": "der"
  },
  {
    "noun": "Bekanntenkreis",
    "article": "der"
  },
  {
    "noun": "Bekanntheitsgrad",
    "article": "der"
  },
  {
    "noun": "Bekehrungseifer",
    "article": "der"
  },
  {
    "noun": "Bekenner",
    "article": "der"
  },
  {
    "noun": "Bekleidungsgegenstand",
    "article": "der"
  },
  {
    "noun": "Belag",
    "article": "der"
  },
  {
    "noun": "Belagerer",
    "article": "der"
  },
  {
    "noun": "Belagerungszustand",
    "article": "der"
  },
  {
    "noun": "Belang",
    "article": "der"
  },
  {
    "noun": "Belarusse",
    "article": "der"
  },
  {
    "noun": "Belastungszeuge",
    "article": "der"
  },
  {
    "noun": "Belauf",
    "article": "der"
  },
  {
    "noun": "Belebungsversuch",
    "article": "der"
  },
  {
    "noun": "Beleg",
    "article": "der"
  },
  {
    "noun": "Belegarzt",
    "article": "der"
  },
  {
    "noun": "Belegleser",
    "article": "der"
  },
  {
    "noun": "Belegnagel",
    "article": "der"
  },
  {
    "noun": "Belegsortierer",
    "article": "der"
  },
  {
    "noun": "Beleidiger",
    "article": "der"
  },
  {
    "noun": "Beleidigungsprozess",
    "article": "der"
  },
  {
    "noun": "Beleihungswert",
    "article": "der"
  },
  {
    "noun": "Belemnit",
    "article": "der"
  },
  {
    "noun": "Beleuchter",
    "article": "der"
  },
  {
    "noun": "Beleuchtungsapparat",
    "article": "der"
  },
  {
    "noun": "Beleuchtungsmesser",
    "article": "der"
  },
  {
    "noun": "Belgier",
    "article": "der"
  },
  {
    "noun": "Belichtungsmesser",
    "article": "der"
  },
  {
    "noun": "Belletrist",
    "article": "der"
  },
  {
    "noun": "Belt",
    "article": "der"
  },
  {
    "noun": "Bendel",
    "article": "der"
  },
  {
    "noun": "Benediktiner",
    "article": "der"
  },
  {
    "noun": "Benediktinerorden",
    "article": "der"
  },
  {
    "noun": "Bengale",
    "article": "der"
  },
  {
    "noun": "Bengel",
    "article": "der"
  },
  {
    "noun": "Benimm",
    "article": "der"
  },
  {
    "noun": "Bentonit",
    "article": "der"
  },
  {
    "noun": "Benutzer",
    "article": "der"
  },
  {
    "noun": "Benutzerkreis",
    "article": "der"
  },
  {
    "noun": "Benzinabscheider",
    "article": "der"
  },
  {
    "noun": "Benziner",
    "article": "der"
  },
  {
    "noun": "Benzingutschein",
    "article": "der"
  },
  {
    "noun": "Benzinhahn",
    "article": "der"
  },
  {
    "noun": "Benzinkanister",
    "article": "der"
  },
  {
    "noun": "Benzinmotor",
    "article": "der"
  },
  {
    "noun": "Benzinpreis",
    "article": "der"
  },
  {
    "noun": "Benzintank",
    "article": "der"
  },
  {
    "noun": "Benzinverbrauch",
    "article": "der"
  },
  {
    "noun": "Beo",
    "article": "der"
  },
  {
    "noun": "Beobachter",
    "article": "der"
  },
  {
    "noun": "Beobachterstatus",
    "article": "der"
  },
  {
    "noun": "Beobachtungsfehler",
    "article": "der"
  },
  {
    "noun": "Beobachtungsposten",
    "article": "der"
  },
  {
    "noun": "Berater",
    "article": "der"
  },
  {
    "noun": "Beratervertrag",
    "article": "der"
  },
  {
    "noun": "Beratungsausschuss",
    "article": "der"
  },
  {
    "noun": "Beratungsdienst",
    "article": "der"
  },
  {
    "noun": "Beratungsvertrag",
    "article": "der"
  },
  {
    "noun": "Berechtigungsnachweis",
    "article": "der"
  },
  {
    "noun": "Berechtigungsschein",
    "article": "der"
  },
  {
    "noun": "Bereitschaftsarzt",
    "article": "der"
  },
  {
    "noun": "Bereitschaftsdienst",
    "article": "der"
  },
  {
    "noun": "Bergabhang",
    "article": "der"
  },
  {
    "noun": "Bergahorn",
    "article": "der"
  },
  {
    "noun": "Bergarbeiter",
    "article": "der"
  },
  {
    "noun": "Bergbau",
    "article": "der"
  },
  {
    "noun": "Bergbewohner",
    "article": "der"
  },
  {
    "noun": "Bergepanzer",
    "article": "der"
  },
  {
    "noun": "Bergfink",
    "article": "der"
  },
  {
    "noun": "Bergfried",
    "article": "der"
  },
  {
    "noun": "Berggeist",
    "article": "der"
  },
  {
    "noun": "Berggipfel",
    "article": "der"
  },
  {
    "noun": "Berghang",
    "article": "der"
  },
  {
    "noun": "Bergingenieur",
    "article": "der"
  },
  {
    "noun": "Bergkamm",
    "article": "der"
  },
  {
    "noun": "Bergkegel",
    "article": "der"
  },
  {
    "noun": "Bergkessel",
    "article": "der"
  },
  {
    "noun": "Bergknappe",
    "article": "der"
  },
  {
    "noun": "Bergkristall",
    "article": "der"
  },
  {
    "noun": "Bergler",
    "article": "der"
  },
  {
    "noun": "Bergmann",
    "article": "der"
  },
  {
    "noun": "Bergpfad",
    "article": "der"
  },
  {
    "noun": "Bergrettungsdienst",
    "article": "der"
  },
  {
    "noun": "Bergrutsch",
    "article": "der"
  },
  {
    "noun": "Bergsattel",
    "article": "der"
  },
  {
    "noun": "Bergschaden",
    "article": "der"
  },
  {
    "noun": "Bergschuh",
    "article": "der"
  },
  {
    "noun": "Bergsee",
    "article": "der"
  },
  {
    "noun": "Bergsteiger",
    "article": "der"
  },
  {
    "noun": "Bergstock",
    "article": "der"
  },
  {
    "noun": "Bergsturz",
    "article": "der"
  },
  {
    "noun": "Bergwohlverleih",
    "article": "der"
  },
  {
    "noun": "Bericht",
    "article": "der"
  },
  {
    "noun": "Berichter",
    "article": "der"
  },
  {
    "noun": "Berichterstatter",
    "article": "der"
  },
  {
    "noun": "Berichtsmonat",
    "article": "der"
  },
  {
    "noun": "Berichtszeitraum",
    "article": "der"
  },
  {
    "noun": "Bernhardiner",
    "article": "der"
  },
  {
    "noun": "Bernhardinerhund",
    "article": "der"
  },
  {
    "noun": "Bernstein",
    "article": "der"
  },
  {
    "noun": "Berserker",
    "article": "der"
  },
  {
    "noun": "Beruf",
    "article": "der"
  },
  {
    "noun": "Berufsberater",
    "article": "der"
  },
  {
    "noun": "Berufsboxer",
    "article": "der"
  },
  {
    "noun": "Berufskollege",
    "article": "der"
  },
  {
    "noun": "Berufskraftfahrer",
    "article": "der"
  },
  {
    "noun": "Berufsoffizier",
    "article": "der"
  },
  {
    "noun": "Berufsrichter",
    "article": "der"
  },
  {
    "noun": "Berufssoldat",
    "article": "der"
  },
  {
    "noun": "Berufsspieler",
    "article": "der"
  },
  {
    "noun": "Berufssportler",
    "article": "der"
  },
  {
    "noun": "Berufsstand",
    "article": "der"
  },
  {
    "noun": "Berufsverband",
    "article": "der"
  },
  {
    "noun": "Berufsverbrecher",
    "article": "der"
  },
  {
    "noun": "Berufsverkehr",
    "article": "der"
  },
  {
    "noun": "Berufswechsel",
    "article": "der"
  },
  {
    "noun": "Berufsweg",
    "article": "der"
  },
  {
    "noun": "Berufswunsch",
    "article": "der"
  },
  {
    "noun": "Beryll",
    "article": "der"
  },
  {
    "noun": "Besan",
    "article": "der"
  },
  {
    "noun": "Besanmast",
    "article": "der"
  },
  {
    "noun": "Besatz",
    "article": "der"
  },
  {
    "noun": "Besatzer",
    "article": "der"
  },
  {
    "noun": "Beschaffer",
    "article": "der"
  },
  {
    "noun": "Beschatter",
    "article": "der"
  },
  {
    "noun": "Beschauer",
    "article": "der"
  },
  {
    "noun": "Bescheid",
    "article": "der"
  },
  {
    "noun": "Beschiss",
    "article": "der"
  },
  {
    "noun": "Beschlag",
    "article": "der"
  },
  {
    "noun": "Beschleuniger",
    "article": "der"
  },
  {
    "noun": "Beschluss",
    "article": "der"
  },
  {
    "noun": "Beschmutzer",
    "article": "der"
  },
  {
    "noun": "Beschreiber",
    "article": "der"
  },
  {
    "noun": "Beschuss",
    "article": "der"
  },
  {
    "noun": "Besen",
    "article": "der"
  },
  {
    "noun": "Besenbinder",
    "article": "der"
  },
  {
    "noun": "Besenschrank",
    "article": "der"
  },
  {
    "noun": "Besenstiel",
    "article": "der"
  },
  {
    "noun": "Besetzer",
    "article": "der"
  },
  {
    "noun": "Besitz",
    "article": "der"
  },
  {
    "noun": "Besitzanspruch",
    "article": "der"
  },
  {
    "noun": "Besitzdiener",
    "article": "der"
  },
  {
    "noun": "Besitzer",
    "article": "der"
  },
  {
    "noun": "Besitzstand",
    "article": "der"
  },
  {
    "noun": "Besitztitel",
    "article": "der"
  },
  {
    "noun": "Besitzwechsel",
    "article": "der"
  },
  {
    "noun": "Besserwisser",
    "article": "der"
  },
  {
    "noun": "Bestand",
    "article": "der"
  },
  {
    "noun": "Bestandteil",
    "article": "der"
  },
  {
    "noun": "Bestatter",
    "article": "der"
  },
  {
    "noun": "Bestbieter",
    "article": "der"
  },
  {
    "noun": "Bestechungsskandal",
    "article": "der"
  },
  {
    "noun": "Bestechungsversuch",
    "article": "der"
  },
  {
    "noun": "Besteckkasten",
    "article": "der"
  },
  {
    "noun": "Bestelleingang",
    "article": "der"
  },
  {
    "noun": "Besteller",
    "article": "der"
  },
  {
    "noun": "Bestellschein",
    "article": "der"
  },
  {
    "noun": "Bestellzettel",
    "article": "der"
  },
  {
    "noun": "Bestimmungsbahnhof",
    "article": "der"
  },
  {
    "noun": "Bestimmungshafen",
    "article": "der"
  },
  {
    "noun": "Bestimmungsort",
    "article": "der"
  },
  {
    "noun": "Bestmann",
    "article": "der"
  },
  {
    "noun": "Bestseller",
    "article": "der"
  },
  {
    "noun": "Bestwert",
    "article": "der"
  },
  {
    "noun": "Bestzustand",
    "article": "der"
  },
  {
    "noun": "Besuch",
    "article": "der"
  },
  {
    "noun": "Besucher",
    "article": "der"
  },
  {
    "noun": "Besucherrekord",
    "article": "der"
  },
  {
    "noun": "Besucherstrom",
    "article": "der"
  },
  {
    "noun": "Betablocker",
    "article": "der"
  },
  {
    "noun": "Beton",
    "article": "der"
  },
  {
    "noun": "Betonbau",
    "article": "der"
  },
  {
    "noun": "Betonbauer",
    "article": "der"
  },
  {
    "noun": "Betonblock",
    "article": "der"
  },
  {
    "noun": "Betonboden",
    "article": "der"
  },
  {
    "noun": "Betonklotz",
    "article": "der"
  },
  {
    "noun": "Betonkopf",
    "article": "der"
  },
  {
    "noun": "Betonmischer",
    "article": "der"
  },
  {
    "noun": "Betonsockel",
    "article": "der"
  },
  {
    "noun": "Betonstein",
    "article": "der"
  },
  {
    "noun": "Betracht",
    "article": "der"
  },
  {
    "noun": "Betrachter",
    "article": "der"
  },
  {
    "noun": "Betrachtungswinkel",
    "article": "der"
  },
  {
    "noun": "Betrag",
    "article": "der"
  },
  {
    "noun": "Betreff",
    "article": "der"
  },
  {
    "noun": "Betreiber",
    "article": "der"
  },
  {
    "noun": "Betreuer",
    "article": "der"
  },
  {
    "noun": "Betrieb",
    "article": "der"
  },
  {
    "noun": "Betriebsarzt",
    "article": "der"
  },
  {
    "noun": "Betriebsausflug",
    "article": "der"
  },
  {
    "noun": "Betriebsfriede",
    "article": "der"
  },
  {
    "noun": "Betriebsingenieur",
    "article": "der"
  },
  {
    "noun": "Betriebsinhaber",
    "article": "der"
  },
  {
    "noun": "Betriebsleiter",
    "article": "der"
  },
  {
    "noun": "Betriebsobmann",
    "article": "der"
  },
  {
    "noun": "Betriebsrat",
    "article": "der"
  },
  {
    "noun": "Betriebsschluss",
    "article": "der"
  },
  {
    "noun": "Betriebsschutz",
    "article": "der"
  },
  {
    "noun": "Betriebsstoff",
    "article": "der"
  },
  {
    "noun": "Betriebsunfall",
    "article": "der"
  },
  {
    "noun": "Betriebswirt",
    "article": "der"
  },
  {
    "noun": "Betrug",
    "article": "der"
  },
  {
    "noun": "Betrugsversuch",
    "article": "der"
  },
  {
    "noun": "Bettbezug",
    "article": "der"
  },
  {
    "noun": "Bettel",
    "article": "der"
  },
  {
    "noun": "Bettelbrief",
    "article": "der"
  },
  {
    "noun": "Bettelsack",
    "article": "der"
  },
  {
    "noun": "Bettgenosse",
    "article": "der"
  },
  {
    "noun": "Betthase",
    "article": "der"
  },
  {
    "noun": "Betthimmel",
    "article": "der"
  },
  {
    "noun": "Bettler",
    "article": "der"
  },
  {
    "noun": "Bettpfosten",
    "article": "der"
  },
  {
    "noun": "Bettrand",
    "article": "der"
  },
  {
    "noun": "Bettzipfel",
    "article": "der"
  },
  {
    "noun": "Beugemuskel",
    "article": "der"
  },
  {
    "noun": "Beurteiler",
    "article": "der"
  },
  {
    "noun": "Beutel",
    "article": "der"
  },
  {
    "noun": "Beutezug",
    "article": "der"
  },
  {
    "noun": "Bewacher",
    "article": "der"
  },
  {
    "noun": "Bewahrer",
    "article": "der"
  },
  {
    "noun": "Beweggrund",
    "article": "der"
  },
  {
    "noun": "Bewegungsablauf",
    "article": "der"
  },
  {
    "noun": "Bewegungsapparat",
    "article": "der"
  },
  {
    "noun": "Bewegungsmelder",
    "article": "der"
  },
  {
    "noun": "Bewegungsnerv",
    "article": "der"
  },
  {
    "noun": "Beweis",
    "article": "der"
  },
  {
    "noun": "Beweisgrund",
    "article": "der"
  },
  {
    "noun": "Bewerber",
    "article": "der"
  },
  {
    "noun": "Bewerbungsbogen",
    "article": "der"
  },
  {
    "noun": "Bewilligungszeitraum",
    "article": "der"
  },
  {
    "noun": "Bewohner",
    "article": "der"
  },
  {
    "noun": "Bewuchs",
    "article": "der"
  },
  {
    "noun": "Bewunderer",
    "article": "der"
  },
  {
    "noun": "Bewurf",
    "article": "der"
  },
  {
    "noun": "Bewusstseinszustand",
    "article": "der"
  },
  {
    "noun": "Bezieher",
    "article": "der"
  },
  {
    "noun": "Beziehungswahn",
    "article": "der"
  },
  {
    "noun": "Bezirk",
    "article": "der"
  },
  {
    "noun": "Bezirksarzt",
    "article": "der"
  },
  {
    "noun": "Bezirksleiter",
    "article": "der"
  },
  {
    "noun": "Bezirksrichter",
    "article": "der"
  },
  {
    "noun": "Bezirksvorsteher",
    "article": "der"
  },
  {
    "noun": "Bezug",
    "article": "der"
  },
  {
    "noun": "Bezugspreis",
    "article": "der"
  },
  {
    "noun": "Bezugspunkt",
    "article": "der"
  },
  {
    "noun": "Bezugsschein",
    "article": "der"
  },
  {
    "noun": "Bezwinger",
    "article": "der"
  },
  {
    "noun": "Bibelspruch",
    "article": "der"
  },
  {
    "noun": "Bibelvers",
    "article": "der"
  },
  {
    "noun": "Biberpelz",
    "article": "der"
  },
  {
    "noun": "Biberschwanz",
    "article": "der"
  },
  {
    "noun": "Bibliograf",
    "article": "der"
  },
  {
    "noun": "Bibliothekar",
    "article": "der"
  },
  {
    "noun": "Biedermann",
    "article": "der"
  },
  {
    "noun": "Bienenfresser",
    "article": "der"
  },
  {
    "noun": "Bienenkorb",
    "article": "der"
  },
  {
    "noun": "Bienenschwarm",
    "article": "der"
  },
  {
    "noun": "Bienenstachel",
    "article": "der"
  },
  {
    "noun": "Bienenstich",
    "article": "der"
  },
  {
    "noun": "Bienenstock",
    "article": "der"
  },
  {
    "noun": "Bierbauch",
    "article": "der"
  },
  {
    "noun": "Bierbrauer",
    "article": "der"
  },
  {
    "noun": "Bierdeckel",
    "article": "der"
  },
  {
    "noun": "Biereifer",
    "article": "der"
  },
  {
    "noun": "Bierfilz",
    "article": "der"
  },
  {
    "noun": "Biergarten",
    "article": "der"
  },
  {
    "noun": "Bierkasten",
    "article": "der"
  },
  {
    "noun": "Bierkeller",
    "article": "der"
  },
  {
    "noun": "Bierkrug",
    "article": "der"
  },
  {
    "noun": "Bierkutscher",
    "article": "der"
  },
  {
    "noun": "Bierschaum",
    "article": "der"
  },
  {
    "noun": "Bierschinken",
    "article": "der"
  },
  {
    "noun": "Bieter",
    "article": "der"
  },
  {
    "noun": "Bigamist",
    "article": "der"
  },
  {
    "noun": "Bikini",
    "article": "der"
  },
  {
    "noun": "Bilanzbuchhalter",
    "article": "der"
  },
  {
    "noun": "Bilateralismus",
    "article": "der"
  },
  {
    "noun": "Bilch",
    "article": "der"
  },
  {
    "noun": "Bildabtaster",
    "article": "der"
  },
  {
    "noun": "Bildabzug",
    "article": "der"
  },
  {
    "noun": "Bildausschnitt",
    "article": "der"
  },
  {
    "noun": "Bildband",
    "article": "der"
  },
  {
    "noun": "Bildbericht",
    "article": "der"
  },
  {
    "noun": "Bilderrahmen",
    "article": "der"
  },
  {
    "noun": "Bildersturm",
    "article": "der"
  },
  {
    "noun": "Bildfunk",
    "article": "der"
  },
  {
    "noun": "Bildhauer",
    "article": "der"
  },
  {
    "noun": "Bildmischer",
    "article": "der"
  },
  {
    "noun": "Bildner",
    "article": "der"
  },
  {
    "noun": "Bildpunkt",
    "article": "der"
  },
  {
    "noun": "Bildschirm",
    "article": "der"
  },
  {
    "noun": "Bildschirmschoner",
    "article": "der"
  },
  {
    "noun": "Bildschirmtext",
    "article": "der"
  },
  {
    "noun": "Bildschnitzer",
    "article": "der"
  },
  {
    "noun": "Bildstreifen",
    "article": "der"
  },
  {
    "noun": "Bildsucher",
    "article": "der"
  },
  {
    "noun": "Bildteppich",
    "article": "der"
  },
  {
    "noun": "Bildungsauftrag",
    "article": "der"
  },
  {
    "noun": "Bildungsgang",
    "article": "der"
  },
  {
    "noun": "Bildungsgrad",
    "article": "der"
  },
  {
    "noun": "Bildungshunger",
    "article": "der"
  },
  {
    "noun": "Bildungsroman",
    "article": "der"
  },
  {
    "noun": "Bildungsstand",
    "article": "der"
  },
  {
    "noun": "Bildungsurlaub",
    "article": "der"
  },
  {
    "noun": "Bildungsweg",
    "article": "der"
  },
  {
    "noun": "Bildwerfer",
    "article": "der"
  },
  {
    "noun": "Bildwinkel",
    "article": "der"
  },
  {
    "noun": "Billardspieler",
    "article": "der"
  },
  {
    "noun": "Billardstock",
    "article": "der"
  },
  {
    "noun": "Billardtisch",
    "article": "der"
  },
  {
    "noun": "Billeteur",
    "article": "der"
  },
  {
    "noun": "Billigflieger",
    "article": "der"
  },
  {
    "noun": "Bimbo",
    "article": "der"
  },
  {
    "noun": "Bimetallismus",
    "article": "der"
  },
  {
    "noun": "Bims",
    "article": "der"
  },
  {
    "noun": "Bimsbeton",
    "article": "der"
  },
  {
    "noun": "Bimsstein",
    "article": "der"
  },
  {
    "noun": "Bindebogen",
    "article": "der"
  },
  {
    "noun": "Binder",
    "article": "der"
  },
  {
    "noun": "Bindestrich",
    "article": "der"
  },
  {
    "noun": "Bindfaden",
    "article": "der"
  },
  {
    "noun": "Binnenhafen",
    "article": "der"
  },
  {
    "noun": "Binnenhandel",
    "article": "der"
  },
  {
    "noun": "Binnenmarkt",
    "article": "der"
  },
  {
    "noun": "Binnensee",
    "article": "der"
  },
  {
    "noun": "Binnenstaat",
    "article": "der"
  },
  {
    "noun": "Binnentarif",
    "article": "der"
  },
  {
    "noun": "Binnenverkehr",
    "article": "der"
  },
  {
    "noun": "Binomialkoeffizient",
    "article": "der"
  },
  {
    "noun": "Biochemiker",
    "article": "der"
  },
  {
    "noun": "Biograf",
    "article": "der"
  },
  {
    "noun": "Bioladen",
    "article": "der"
  },
  {
    "noun": "Biologe",
    "article": "der"
  },
  {
    "noun": "Biologismus",
    "article": "der"
  },
  {
    "noun": "Biorhythmus",
    "article": "der"
  },
  {
    "noun": "Biotit",
    "article": "der"
  },
  {
    "noun": "Biotyp",
    "article": "der"
  },
  {
    "noun": "Biotypus",
    "article": "der"
  },
  {
    "noun": "Birkenhain",
    "article": "der"
  },
  {
    "noun": "Birmane",
    "article": "der"
  },
  {
    "noun": "Birnbaum",
    "article": "der"
  },
  {
    "noun": "Birnenbaum",
    "article": "der"
  },
  {
    "noun": "Bisam",
    "article": "der"
  },
  {
    "noun": "Bischof",
    "article": "der"
  },
  {
    "noun": "Bischofshut",
    "article": "der"
  },
  {
    "noun": "Bischofssitz",
    "article": "der"
  },
  {
    "noun": "Bischofsstab",
    "article": "der"
  },
  {
    "noun": "Biskuitteig",
    "article": "der"
  },
  {
    "noun": "Bismarckhering",
    "article": "der"
  },
  {
    "noun": "Bismutit",
    "article": "der"
  },
  {
    "noun": "Bison",
    "article": "der"
  },
  {
    "noun": "Biss",
    "article": "der"
  },
  {
    "noun": "Bissen",
    "article": "der"
  },
  {
    "noun": "Bittbrief",
    "article": "der"
  },
  {
    "noun": "Bitterling",
    "article": "der"
  },
  {
    "noun": "Bitterstoff",
    "article": "der"
  },
  {
    "noun": "Bittgesang",
    "article": "der"
  },
  {
    "noun": "Bittsteller",
    "article": "der"
  },
  {
    "noun": "Bizeps",
    "article": "der"
  },
  {
    "noun": "Blankoakzept",
    "article": "der"
  },
  {
    "noun": "Blankokredit",
    "article": "der"
  },
  {
    "noun": "Blankoscheck",
    "article": "der"
  },
  {
    "noun": "Blattgrund",
    "article": "der"
  },
  {
    "noun": "Blattnerv",
    "article": "der"
  },
  {
    "noun": "Blattsalat",
    "article": "der"
  },
  {
    "noun": "Blattschuss",
    "article": "der"
  },
  {
    "noun": "Blattstiel",
    "article": "der"
  },
  {
    "noun": "Blaufuchs",
    "article": "der"
  },
  {
    "noun": "Blauhelm",
    "article": "der"
  },
  {
    "noun": "Blaukohl",
    "article": "der"
  },
  {
    "noun": "Blauschimmel",
    "article": "der"
  },
  {
    "noun": "Blaustrumpf",
    "article": "der"
  },
  {
    "noun": "Blauton",
    "article": "der"
  },
  {
    "noun": "Blauwal",
    "article": "der"
  },
  {
    "noun": "Blazer",
    "article": "der"
  },
  {
    "noun": "Blecheimer",
    "article": "der"
  },
  {
    "noun": "Blechnapf",
    "article": "der"
  },
  {
    "noun": "Blechner",
    "article": "der"
  },
  {
    "noun": "Blechschaden",
    "article": "der"
  },
  {
    "noun": "Blechschmied",
    "article": "der"
  },
  {
    "noun": "Bleichplatz",
    "article": "der"
  },
  {
    "noun": "Bleigehalt",
    "article": "der"
  },
  {
    "noun": "Bleiglanz",
    "article": "der"
  },
  {
    "noun": "Bleisatz",
    "article": "der"
  },
  {
    "noun": "Bleisoldat",
    "article": "der"
  },
  {
    "noun": "Bleistift",
    "article": "der"
  },
  {
    "noun": "Bleistiftspitzer",
    "article": "der"
  },
  {
    "noun": "Blendbogen",
    "article": "der"
  },
  {
    "noun": "Blender",
    "article": "der"
  },
  {
    "noun": "Blick",
    "article": "der"
  },
  {
    "noun": "Blickfang",
    "article": "der"
  },
  {
    "noun": "Blickpunkt",
    "article": "der"
  },
  {
    "noun": "Blickwinkel",
    "article": "der"
  },
  {
    "noun": "Blinddarm",
    "article": "der"
  },
  {
    "noun": "Blindenhund",
    "article": "der"
  },
  {
    "noun": "Blindenstock",
    "article": "der"
  },
  {
    "noun": "Blindflug",
    "article": "der"
  },
  {
    "noun": "Blindschacht",
    "article": "der"
  },
  {
    "noun": "Blindversuch",
    "article": "der"
  },
  {
    "noun": "Blinker",
    "article": "der"
  },
  {
    "noun": "Blinkgeber",
    "article": "der"
  },
  {
    "noun": "Blitz",
    "article": "der"
  },
  {
    "noun": "Blitzableiter",
    "article": "der"
  },
  {
    "noun": "Blitzbesuch",
    "article": "der"
  },
  {
    "noun": "Blitzer",
    "article": "der"
  },
  {
    "noun": "Blitzkrieg",
    "article": "der"
  },
  {
    "noun": "Blitzschlag",
    "article": "der"
  },
  {
    "noun": "Blitzstrahl",
    "article": "der"
  },
  {
    "noun": "Blizzard",
    "article": "der"
  },
  {
    "noun": "Block",
    "article": "der"
  },
  {
    "noun": "Blockbau",
    "article": "der"
  },
  {
    "noun": "Blockunterricht",
    "article": "der"
  },
  {
    "noun": "Blockwart",
    "article": "der"
  },
  {
    "noun": "Blues",
    "article": "der"
  },
  {
    "noun": "Bluff",
    "article": "der"
  },
  {
    "noun": "Bluffer",
    "article": "der"
  },
  {
    "noun": "Blumenbinder",
    "article": "der"
  },
  {
    "noun": "Blumenflor",
    "article": "der"
  },
  {
    "noun": "Blumengarten",
    "article": "der"
  },
  {
    "noun": "Blumenkasten",
    "article": "der"
  },
  {
    "noun": "Blumenkohl",
    "article": "der"
  },
  {
    "noun": "Blumenkranz",
    "article": "der"
  },
  {
    "noun": "Blumenladen",
    "article": "der"
  },
  {
    "noun": "Blumenschmuck",
    "article": "der"
  },
  {
    "noun": "Blumenstock",
    "article": "der"
  },
  {
    "noun": "Blumentopf",
    "article": "der"
  },
  {
    "noun": "Blutalkohol",
    "article": "der"
  },
  {
    "noun": "Blutalkoholgehalt",
    "article": "der"
  },
  {
    "noun": "Blutandrang",
    "article": "der"
  },
  {
    "noun": "Blutdruck",
    "article": "der"
  },
  {
    "noun": "Blutegel",
    "article": "der"
  },
  {
    "noun": "Bluter",
    "article": "der"
  },
  {
    "noun": "Bluterguss",
    "article": "der"
  },
  {
    "noun": "Blutfleck",
    "article": "der"
  },
  {
    "noun": "Bluthochdruck",
    "article": "der"
  },
  {
    "noun": "Bluthund",
    "article": "der"
  },
  {
    "noun": "Bluthusten",
    "article": "der"
  },
  {
    "noun": "Blutkreislauf",
    "article": "der"
  },
  {
    "noun": "Blutmangel",
    "article": "der"
  },
  {
    "noun": "Blutpfropf",
    "article": "der"
  },
  {
    "noun": "Blutsauger",
    "article": "der"
  },
  {
    "noun": "Blutschwamm",
    "article": "der"
  },
  {
    "noun": "Blutspendedienst",
    "article": "der"
  },
  {
    "noun": "Blutspender",
    "article": "der"
  },
  {
    "noun": "Blutstatus",
    "article": "der"
  },
  {
    "noun": "Blutstrom",
    "article": "der"
  },
  {
    "noun": "Blutstropfen",
    "article": "der"
  },
  {
    "noun": "Blutsturz",
    "article": "der"
  },
  {
    "noun": "Blutverlust",
    "article": "der"
  },
  {
    "noun": "Blutzucker",
    "article": "der"
  },
  {
    "noun": "BMW",
    "article": "der"
  },
  {
    "noun": "Bob",
    "article": "der"
  },
  {
    "noun": "Bockmist",
    "article": "der"
  },
  {
    "noun": "Bockshornklee",
    "article": "der"
  },
  {
    "noun": "Bocksprung",
    "article": "der"
  },
  {
    "noun": "Bodden",
    "article": "der"
  },
  {
    "noun": "Boden",
    "article": "der"
  },
  {
    "noun": "Bodenbelag",
    "article": "der"
  },
  {
    "noun": "Bodendruck",
    "article": "der"
  },
  {
    "noun": "Bodenertrag",
    "article": "der"
  },
  {
    "noun": "Bodenfrost",
    "article": "der"
  },
  {
    "noun": "Bodenfund",
    "article": "der"
  },
  {
    "noun": "Bodenleger",
    "article": "der"
  },
  {
    "noun": "Bodennebel",
    "article": "der"
  },
  {
    "noun": "Bodensatz",
    "article": "der"
  },
  {
    "noun": "Bodenschatz",
    "article": "der"
  },
  {
    "noun": "Bodensee",
    "article": "der"
  },
  {
    "noun": "Body",
    "article": "der"
  },
  {
    "noun": "Bodybuilder",
    "article": "der"
  },
  {
    "noun": "Bodycheck",
    "article": "der"
  },
  {
    "noun": "Bofist",
    "article": "der"
  },
  {
    "noun": "Bogen",
    "article": "der"
  },
  {
    "noun": "Bogengang",
    "article": "der"
  },
  {
    "noun": "Bohemien",
    "article": "der"
  },
  {
    "noun": "Bohneneintopf",
    "article": "der"
  },
  {
    "noun": "Bohnenkaffee",
    "article": "der"
  },
  {
    "noun": "Bohnensalat",
    "article": "der"
  },
  {
    "noun": "Bohrer",
    "article": "der"
  },
  {
    "noun": "Bohrhammer",
    "article": "der"
  },
  {
    "noun": "Bohrturm",
    "article": "der"
  },
  {
    "noun": "Boiler",
    "article": "der"
  },
  {
    "noun": "Bojar",
    "article": "der"
  },
  {
    "noun": "Bolero",
    "article": "der"
  },
  {
    "noun": "Bolid",
    "article": "der"
  },
  {
    "noun": "Bolivianer",
    "article": "der"
  },
  {
    "noun": "Bolivier",
    "article": "der"
  },
  {
    "noun": "Bolschewik",
    "article": "der"
  },
  {
    "noun": "Bolschewismus",
    "article": "der"
  },
  {
    "noun": "Bolschewist",
    "article": "der"
  },
  {
    "noun": "Bolus",
    "article": "der"
  },
  {
    "noun": "Bolzplatz",
    "article": "der"
  },
  {
    "noun": "Bombast",
    "article": "der"
  },
  {
    "noun": "Bombenalarm",
    "article": "der"
  },
  {
    "noun": "Bombenangriff",
    "article": "der"
  },
  {
    "noun": "Bombenanschlag",
    "article": "der"
  },
  {
    "noun": "Bombenerfolg",
    "article": "der"
  },
  {
    "noun": "Bombenleger",
    "article": "der"
  },
  {
    "noun": "Bombentrichter",
    "article": "der"
  },
  {
    "noun": "Bomber",
    "article": "der"
  },
  {
    "noun": "Bon",
    "article": "der"
  },
  {
    "noun": "Bond",
    "article": "der"
  },
  {
    "noun": "Bonus",
    "article": "der"
  },
  {
    "noun": "Bonuspunkt",
    "article": "der"
  },
  {
    "noun": "Bonze",
    "article": "der"
  },
  {
    "noun": "Boom",
    "article": "der"
  },
  {
    "noun": "Booster",
    "article": "der"
  },
  {
    "noun": "Bootsbau",
    "article": "der"
  },
  {
    "noun": "Bootsbauer",
    "article": "der"
  },
  {
    "noun": "Bootshaken",
    "article": "der"
  },
  {
    "noun": "Bootsmann",
    "article": "der"
  },
  {
    "noun": "Bootsmannsmaat",
    "article": "der"
  },
  {
    "noun": "Bootsverleih",
    "article": "der"
  },
  {
    "noun": "Borax",
    "article": "der"
  },
  {
    "noun": "Bordcomputer",
    "article": "der"
  },
  {
    "noun": "Borddienst",
    "article": "der"
  },
  {
    "noun": "Bordeauxwein",
    "article": "der"
  },
  {
    "noun": "Bordfunk",
    "article": "der"
  },
  {
    "noun": "Bordfunker",
    "article": "der"
  },
  {
    "noun": "Bordmechaniker",
    "article": "der"
  },
  {
    "noun": "Bordstein",
    "article": "der"
  },
  {
    "noun": "Bornit",
    "article": "der"
  },
  {
    "noun": "Borretsch",
    "article": "der"
  },
  {
    "noun": "Borschtsch",
    "article": "der"
  },
  {
    "noun": "Borstenwurm",
    "article": "der"
  },
  {
    "noun": "Boskoop",
    "article": "der"
  },
  {
    "noun": "Bosniak",
    "article": "der"
  },
  {
    "noun": "Bosnier",
    "article": "der"
  },
  {
    "noun": "Bosporus",
    "article": "der"
  },
  {
    "noun": "Boss",
    "article": "der"
  },
  {
    "noun": "Bossierer",
    "article": "der"
  },
  {
    "noun": "Botaniker",
    "article": "der"
  },
  {
    "noun": "Botengang",
    "article": "der"
  },
  {
    "noun": "Botschafter",
    "article": "der"
  },
  {
    "noun": "Botschaftsrat",
    "article": "der"
  },
  {
    "noun": "Bottich",
    "article": "der"
  },
  {
    "noun": "Botulismus",
    "article": "der"
  },
  {
    "noun": "Boulevard",
    "article": "der"
  },
  {
    "noun": "Bourbon",
    "article": "der"
  },
  {
    "noun": "Bourbone",
    "article": "der"
  },
  {
    "noun": "Bovist",
    "article": "der"
  },
  {
    "noun": "Boxenstopp",
    "article": "der"
  },
  {
    "noun": "Boxer",
    "article": "der"
  },
  {
    "noun": "Boxermotor",
    "article": "der"
  },
  {
    "noun": "Boxhandschuh",
    "article": "der"
  },
  {
    "noun": "Boxkampf",
    "article": "der"
  },
  {
    "noun": "Boxring",
    "article": "der"
  },
  {
    "noun": "Boykott",
    "article": "der"
  },
  {
    "noun": "Brachacker",
    "article": "der"
  },
  {
    "noun": "Brachet",
    "article": "der"
  },
  {
    "noun": "Brachmonat",
    "article": "der"
  },
  {
    "noun": "Brachpieper",
    "article": "der"
  },
  {
    "noun": "Brachsen",
    "article": "der"
  },
  {
    "noun": "Brachvogel",
    "article": "der"
  },
  {
    "noun": "Branchenmix",
    "article": "der"
  },
  {
    "noun": "Brand",
    "article": "der"
  },
  {
    "noun": "Brandanschlag",
    "article": "der"
  },
  {
    "noun": "Brandfleck",
    "article": "der"
  },
  {
    "noun": "Brandgeruch",
    "article": "der"
  },
  {
    "noun": "Brandherd",
    "article": "der"
  },
  {
    "noun": "Brandleger",
    "article": "der"
  },
  {
    "noun": "Brandmeister",
    "article": "der"
  },
  {
    "noun": "Brandsatz",
    "article": "der"
  },
  {
    "noun": "Brandschaden",
    "article": "der"
  },
  {
    "noun": "Brandschiefer",
    "article": "der"
  },
  {
    "noun": "Brandschutz",
    "article": "der"
  },
  {
    "noun": "Brandstifter",
    "article": "der"
  },
  {
    "noun": "Brandteig",
    "article": "der"
  },
  {
    "noun": "Branntkalk",
    "article": "der"
  },
  {
    "noun": "Branntwein",
    "article": "der"
  },
  {
    "noun": "Branntweinbrenner",
    "article": "der"
  },
  {
    "noun": "Brasilianer",
    "article": "der"
  },
  {
    "noun": "Bratapfel",
    "article": "der"
  },
  {
    "noun": "Bratenfond",
    "article": "der"
  },
  {
    "noun": "Bratensaft",
    "article": "der"
  },
  {
    "noun": "Brathering",
    "article": "der"
  },
  {
    "noun": "Bratrost",
    "article": "der"
  },
  {
    "noun": "Brauch",
    "article": "der"
  },
  {
    "noun": "Brauer",
    "article": "der"
  },
  {
    "noun": "Brauneisenstein",
    "article": "der"
  },
  {
    "noun": "Braunkohl",
    "article": "der"
  },
  {
    "noun": "Braunkohlenteer",
    "article": "der"
  },
  {
    "noun": "Braunspat",
    "article": "der"
  },
  {
    "noun": "Braunstein",
    "article": "der"
  },
  {
    "noun": "Brausekopf",
    "article": "der"
  },
  {
    "noun": "Brautschleier",
    "article": "der"
  },
  {
    "noun": "Brautstand",
    "article": "der"
  },
  {
    "noun": "Brautvater",
    "article": "der"
  },
  {
    "noun": "Brautwerber",
    "article": "der"
  },
  {
    "noun": "Bravoruf",
    "article": "der"
  },
  {
    "noun": "Breakdance",
    "article": "der"
  },
  {
    "noun": "Brecher",
    "article": "der"
  },
  {
    "noun": "Brechreiz",
    "article": "der"
  },
  {
    "noun": "Brechungswinkel",
    "article": "der"
  },
  {
    "noun": "Brei",
    "article": "der"
  },
  {
    "noun": "Breitengrad",
    "article": "der"
  },
  {
    "noun": "Breitenkreis",
    "article": "der"
  },
  {
    "noun": "Breitensport",
    "article": "der"
  },
  {
    "noun": "Breitreifen",
    "article": "der"
  },
  {
    "noun": "Bremsbelag",
    "article": "der"
  },
  {
    "noun": "Bremsberg",
    "article": "der"
  },
  {
    "noun": "Bremser",
    "article": "der"
  },
  {
    "noun": "Bremsfallschirm",
    "article": "der"
  },
  {
    "noun": "Bremshebel",
    "article": "der"
  },
  {
    "noun": "Bremsklotz",
    "article": "der"
  },
  {
    "noun": "Bremsschuh",
    "article": "der"
  },
  {
    "noun": "Bremsweg",
    "article": "der"
  },
  {
    "noun": "Brenner",
    "article": "der"
  },
  {
    "noun": "Brennofen",
    "article": "der"
  },
  {
    "noun": "Brennpunkt",
    "article": "der"
  },
  {
    "noun": "Brennspiegel",
    "article": "der"
  },
  {
    "noun": "Brennspiritus",
    "article": "der"
  },
  {
    "noun": "Brennstab",
    "article": "der"
  },
  {
    "noun": "Brennstoff",
    "article": "der"
  },
  {
    "noun": "Brennstoffstab",
    "article": "der"
  },
  {
    "noun": "Brennwert",
    "article": "der"
  },
  {
    "noun": "Bretterboden",
    "article": "der"
  },
  {
    "noun": "Bretterzaun",
    "article": "der"
  },
  {
    "noun": "Brie",
    "article": "der"
  },
  {
    "noun": "Brief",
    "article": "der"
  },
  {
    "noun": "Briefadel",
    "article": "der"
  },
  {
    "noun": "Briefbeschwerer",
    "article": "der"
  },
  {
    "noun": "Briefbogen",
    "article": "der"
  },
  {
    "noun": "Briefeinwurf",
    "article": "der"
  },
  {
    "noun": "Brieffreund",
    "article": "der"
  },
  {
    "noun": "Briefkasten",
    "article": "der"
  },
  {
    "noun": "Briefkopf",
    "article": "der"
  },
  {
    "noun": "Briefkurs",
    "article": "der"
  },
  {
    "noun": "Briefmarkensammler",
    "article": "der"
  },
  {
    "noun": "Briefordner",
    "article": "der"
  },
  {
    "noun": "Briefpartner",
    "article": "der"
  },
  {
    "noun": "Briefroman",
    "article": "der"
  },
  {
    "noun": "Briefschlitz",
    "article": "der"
  },
  {
    "noun": "Briefumschlag",
    "article": "der"
  },
  {
    "noun": "Briefverkehr",
    "article": "der"
  },
  {
    "noun": "Briefwechsel",
    "article": "der"
  },
  {
    "noun": "Brigadegeneral",
    "article": "der"
  },
  {
    "noun": "Brigadeleiter",
    "article": "der"
  },
  {
    "noun": "Brigadier",
    "article": "der"
  },
  {
    "noun": "Brigant",
    "article": "der"
  },
  {
    "noun": "Brillant",
    "article": "der"
  },
  {
    "noun": "Brillantring",
    "article": "der"
  },
  {
    "noun": "Bringer",
    "article": "der"
  },
  {
    "noun": "Bristolkarton",
    "article": "der"
  },
  {
    "noun": "Brite",
    "article": "der"
  },
  {
    "noun": "Broccoli",
    "article": "der"
  },
  {
    "noun": "Broiler",
    "article": "der"
  },
  {
    "noun": "Brokat",
    "article": "der"
  },
  {
    "noun": "Broker",
    "article": "der"
  },
  {
    "noun": "Brombeerstrauch",
    "article": "der"
  },
  {
    "noun": "Bronchus",
    "article": "der"
  },
  {
    "noun": "Brontosaurus",
    "article": "der"
  },
  {
    "noun": "Bronzit",
    "article": "der"
  },
  {
    "noun": "Brotaufstrich",
    "article": "der"
  },
  {
    "noun": "Brotbaum",
    "article": "der"
  },
  {
    "noun": "Brotbelag",
    "article": "der"
  },
  {
    "noun": "Brotberuf",
    "article": "der"
  },
  {
    "noun": "Brotbeutel",
    "article": "der"
  },
  {
    "noun": "Broterwerb",
    "article": "der"
  },
  {
    "noun": "Brotfruchtbaum",
    "article": "der"
  },
  {
    "noun": "Brotherr",
    "article": "der"
  },
  {
    "noun": "Brotkasten",
    "article": "der"
  },
  {
    "noun": "Brotkorb",
    "article": "der"
  },
  {
    "noun": "Brotlaib",
    "article": "der"
  },
  {
    "noun": "Brottopf",
    "article": "der"
  },
  {
    "noun": "Brotverdiener",
    "article": "der"
  },
  {
    "noun": "Browser",
    "article": "der"
  },
  {
    "noun": "Bruchbau",
    "article": "der"
  },
  {
    "noun": "Bruchschaden",
    "article": "der"
  },
  {
    "noun": "Bruchstein",
    "article": "der"
  },
  {
    "noun": "Bruchstrich",
    "article": "der"
  },
  {
    "noun": "Bruchteil",
    "article": "der"
  },
  {
    "noun": "Bruder",
    "article": "der"
  },
  {
    "noun": "Bruderkrieg",
    "article": "der"
  },
  {
    "noun": "Bruderzwist",
    "article": "der"
  },
  {
    "noun": "Brummbass",
    "article": "der"
  },
  {
    "noun": "Brummer",
    "article": "der"
  },
  {
    "noun": "Brummkreisel",
    "article": "der"
  },
  {
    "noun": "Brummton",
    "article": "der"
  },
  {
    "noun": "Brunch",
    "article": "der"
  },
  {
    "noun": "Brunftschrei",
    "article": "der"
  },
  {
    "noun": "Brunnen",
    "article": "der"
  },
  {
    "noun": "Brunnenbauer",
    "article": "der"
  },
  {
    "noun": "Brunnenschacht",
    "article": "der"
  },
  {
    "noun": "Brustbeutel",
    "article": "der"
  },
  {
    "noun": "Brustharnisch",
    "article": "der"
  },
  {
    "noun": "Brustkasten",
    "article": "der"
  },
  {
    "noun": "Brustkorb",
    "article": "der"
  },
  {
    "noun": "Brustkrebs",
    "article": "der"
  },
  {
    "noun": "Brustpanzer",
    "article": "der"
  },
  {
    "noun": "Brustumfang",
    "article": "der"
  },
  {
    "noun": "Brustwirbel",
    "article": "der"
  },
  {
    "noun": "Brutapparat",
    "article": "der"
  },
  {
    "noun": "Brutkasten",
    "article": "der"
  },
  {
    "noun": "Brutplatz",
    "article": "der"
  },
  {
    "noun": "Brutreaktor",
    "article": "der"
  },
  {
    "noun": "Brutschrank",
    "article": "der"
  },
  {
    "noun": "Bruttobetrag",
    "article": "der"
  },
  {
    "noun": "Bruttoertrag",
    "article": "der"
  },
  {
    "noun": "Bruttogewinn",
    "article": "der"
  },
  {
    "noun": "Bruttolohn",
    "article": "der"
  },
  {
    "noun": "Bruttopreis",
    "article": "der"
  },
  {
    "noun": "Bruttoverdienst",
    "article": "der"
  },
  {
    "noun": "Bruxismus",
    "article": "der"
  },
  {
    "noun": "Bub",
    "article": "der"
  },
  {
    "noun": "Bube",
    "article": "der"
  },
  {
    "noun": "Bubenstreich",
    "article": "der"
  },
  {
    "noun": "Bubikopf",
    "article": "der"
  },
  {
    "noun": "Buchbestand",
    "article": "der"
  },
  {
    "noun": "Buchbinder",
    "article": "der"
  },
  {
    "noun": "Buchblock",
    "article": "der"
  },
  {
    "noun": "Buchclub",
    "article": "der"
  },
  {
    "noun": "Buchdruck",
    "article": "der"
  },
  {
    "noun": "Buchdrucker",
    "article": "der"
  },
  {
    "noun": "Bucheinband",
    "article": "der"
  },
  {
    "noun": "Buchfink",
    "article": "der"
  },
  {
    "noun": "Buchhalter",
    "article": "der"
  },
  {
    "noun": "Buchhandel",
    "article": "der"
  },
  {
    "noun": "Buchladen",
    "article": "der"
  },
  {
    "noun": "Buchmacher",
    "article": "der"
  },
  {
    "noun": "Buchsbaum",
    "article": "der"
  },
  {
    "noun": "Buchstabe",
    "article": "der"
  },
  {
    "noun": "Buchstabenglaube",
    "article": "der"
  },
  {
    "noun": "Buchtitel",
    "article": "der"
  },
  {
    "noun": "Buchumschlag",
    "article": "der"
  },
  {
    "noun": "Buchweizen",
    "article": "der"
  },
  {
    "noun": "Buchwert",
    "article": "der"
  },
  {
    "noun": "Buckel",
    "article": "der"
  },
  {
    "noun": "Buckelstein",
    "article": "der"
  },
  {
    "noun": "Buckelwal",
    "article": "der"
  },
  {
    "noun": "Buddhismus",
    "article": "der"
  },
  {
    "noun": "Buddhist",
    "article": "der"
  },
  {
    "noun": "Budenbesitzer",
    "article": "der"
  },
  {
    "noun": "Budgetentwurf",
    "article": "der"
  },
  {
    "noun": "Bug",
    "article": "der"
  },
  {
    "noun": "Buganker",
    "article": "der"
  },
  {
    "noun": "Buggy",
    "article": "der"
  },
  {
    "noun": "Buhler",
    "article": "der"
  },
  {
    "noun": "Buhmann",
    "article": "der"
  },
  {
    "noun": "Buhnenkopf",
    "article": "der"
  },
  {
    "noun": "Bulgare",
    "article": "der"
  },
  {
    "noun": "Bulldozer",
    "article": "der"
  },
  {
    "noun": "Bulle",
    "article": "der"
  },
  {
    "noun": "Bully",
    "article": "der"
  },
  {
    "noun": "Bumerang",
    "article": "der"
  },
  {
    "noun": "Bummelant",
    "article": "der"
  },
  {
    "noun": "Bummelstreik",
    "article": "der"
  },
  {
    "noun": "Bummler",
    "article": "der"
  },
  {
    "noun": "Bums",
    "article": "der"
  },
  {
    "noun": "Bundesgenosse",
    "article": "der"
  },
  {
    "noun": "Bundesgerichtshof",
    "article": "der"
  },
  {
    "noun": "Bundesgrenzschutz",
    "article": "der"
  },
  {
    "noun": "Bundeshaushalt",
    "article": "der"
  },
  {
    "noun": "Bundeskanzler",
    "article": "der"
  },
  {
    "noun": "Bundesminister",
    "article": "der"
  },
  {
    "noun": "Bundesnachrichtendienst",
    "article": "der"
  },
  {
    "noun": "Bundesrat",
    "article": "der"
  },
  {
    "noun": "Bundesstaat",
    "article": "der"
  },
  {
    "noun": "Bundestag",
    "article": "der"
  },
  {
    "noun": "Bundestrainer",
    "article": "der"
  },
  {
    "noun": "Bungalow",
    "article": "der"
  },
  {
    "noun": "Bunker",
    "article": "der"
  },
  {
    "noun": "Bunsenbrenner",
    "article": "der"
  },
  {
    "noun": "Buntbarsch",
    "article": "der"
  },
  {
    "noun": "Buntdruck",
    "article": "der"
  },
  {
    "noun": "Buntkupferkies",
    "article": "der"
  },
  {
    "noun": "Buntsandstein",
    "article": "der"
  },
  {
    "noun": "Buntspecht",
    "article": "der"
  },
  {
    "noun": "Buntstift",
    "article": "der"
  },
  {
    "noun": "Bure",
    "article": "der"
  },
  {
    "noun": "Burgfriede",
    "article": "der"
  },
  {
    "noun": "Burggraben",
    "article": "der"
  },
  {
    "noun": "Burgunder",
    "article": "der"
  },
  {
    "noun": "Burgunderwein",
    "article": "der"
  },
  {
    "noun": "Burgvogt",
    "article": "der"
  },
  {
    "noun": "Burmese",
    "article": "der"
  },
  {
    "noun": "Bursche",
    "article": "der"
  },
  {
    "noun": "Bus",
    "article": "der"
  },
  {
    "noun": "Busbahnhof",
    "article": "der"
  },
  {
    "noun": "Buschmann",
    "article": "der"
  },
  {
    "noun": "Busen",
    "article": "der"
  },
  {
    "noun": "Busenfreund",
    "article": "der"
  },
  {
    "noun": "Busfahrer",
    "article": "der"
  },
  {
    "noun": "Bussard",
    "article": "der"
  },
  {
    "noun": "Butler",
    "article": "der"
  },
  {
    "noun": "Butt",
    "article": "der"
  },
  {
    "noun": "Butterberg",
    "article": "der"
  },
  {
    "noun": "Butterpilz",
    "article": "der"
  },
  {
    "noun": "Butterteig",
    "article": "der"
  },
  {
    "noun": "Buttervogel",
    "article": "der"
  },
  {
    "noun": "Button",
    "article": "der"
  },
  {
    "noun": "Butzemann",
    "article": "der"
  },
  {
    "noun": "Butzen",
    "article": "der"
  },
  {
    "noun": "Bypass",
    "article": "der"
  },
  {
    "noun": "Byssus",
    "article": "der"
  },
  {
    "noun": "Byzantiner",
    "article": "der"
  },
  {
    "noun": "Kakerlak",
    "article": "der"
  },
  {
    "noun": "Kaktus",
    "article": "der"
  },
  {
    "noun": "Kalbsbraten",
    "article": "der"
  },
  {
    "noun": "Kalender",
    "article": "der"
  },
  {
    "noun": "Kalendertag",
    "article": "der"
  },
  {
    "noun": "Kalibergbau",
    "article": "der"
  },
  {
    "noun": "Kalif",
    "article": "der"
  },
  {
    "noun": "Kalifeldspat",
    "article": "der"
  },
  {
    "noun": "Kalifornier",
    "article": "der"
  },
  {
    "noun": "Kalisalpeter",
    "article": "der"
  },
  {
    "noun": "Kalk",
    "article": "der"
  },
  {
    "noun": "Kalkgehalt",
    "article": "der"
  },
  {
    "noun": "Kalkmangel",
    "article": "der"
  },
  {
    "noun": "Kalkofen",
    "article": "der"
  },
  {
    "noun": "Kalksandstein",
    "article": "der"
  },
  {
    "noun": "Kalkschiefer",
    "article": "der"
  },
  {
    "noun": "Kalkspat",
    "article": "der"
  },
  {
    "noun": "Kalkstein",
    "article": "der"
  },
  {
    "noun": "Kalksteinbruch",
    "article": "der"
  },
  {
    "noun": "Kalktuff",
    "article": "der"
  },
  {
    "noun": "Kalkulator",
    "article": "der"
  },
  {
    "noun": "Kalligraf",
    "article": "der"
  },
  {
    "noun": "Kalmar",
    "article": "der"
  },
  {
    "noun": "Kalorienverbrauch",
    "article": "der"
  },
  {
    "noun": "Kalorifer",
    "article": "der"
  },
  {
    "noun": "Kaltleim",
    "article": "der"
  },
  {
    "noun": "Kaltleiter",
    "article": "der"
  },
  {
    "noun": "Kaltstart",
    "article": "der"
  },
  {
    "noun": "Kalzit",
    "article": "der"
  },
  {
    "noun": "Kambodschaner",
    "article": "der"
  },
  {
    "noun": "Kambrik",
    "article": "der"
  },
  {
    "noun": "Kameldorn",
    "article": "der"
  },
  {
    "noun": "Kameltreiber",
    "article": "der"
  },
  {
    "noun": "Kamerad",
    "article": "der"
  },
  {
    "noun": "Kameradschaftsabend",
    "article": "der"
  },
  {
    "noun": "Kameradschaftsgeist",
    "article": "der"
  },
  {
    "noun": "Kameralismus",
    "article": "der"
  },
  {
    "noun": "Kameramann",
    "article": "der"
  },
  {
    "noun": "Kameraschwenk",
    "article": "der"
  },
  {
    "noun": "Kamikaze",
    "article": "der"
  },
  {
    "noun": "Kamikazeflieger",
    "article": "der"
  },
  {
    "noun": "Kamillentee",
    "article": "der"
  },
  {
    "noun": "Kamin",
    "article": "der"
  },
  {
    "noun": "Kaminfeger",
    "article": "der"
  },
  {
    "noun": "Kaminkehrer",
    "article": "der"
  },
  {
    "noun": "Kaminsims",
    "article": "der"
  },
  {
    "noun": "Kamm",
    "article": "der"
  },
  {
    "noun": "Kammerchor",
    "article": "der"
  },
  {
    "noun": "Kammerdiener",
    "article": "der"
  },
  {
    "noun": "Kammerherr",
    "article": "der"
  },
  {
    "noun": "Kammerofen",
    "article": "der"
  },
  {
    "noun": "Kammerton",
    "article": "der"
  },
  {
    "noun": "Kampf",
    "article": "der"
  },
  {
    "noun": "Kampfanzug",
    "article": "der"
  },
  {
    "noun": "Kampfauftrag",
    "article": "der"
  },
  {
    "noun": "Kampfflieger",
    "article": "der"
  },
  {
    "noun": "Kampfgeist",
    "article": "der"
  },
  {
    "noun": "Kampfgenosse",
    "article": "der"
  },
  {
    "noun": "Kampfhahn",
    "article": "der"
  },
  {
    "noun": "Kampfhund",
    "article": "der"
  },
  {
    "noun": "Kampfpanzer",
    "article": "der"
  },
  {
    "noun": "Kampfplatz",
    "article": "der"
  },
  {
    "noun": "Kampfpreis",
    "article": "der"
  },
  {
    "noun": "Kampfrichter",
    "article": "der"
  },
  {
    "noun": "Kampfschwimmer",
    "article": "der"
  },
  {
    "noun": "Kampfsport",
    "article": "der"
  },
  {
    "noun": "Kampfstoff",
    "article": "der"
  },
  {
    "noun": "Kampfverband",
    "article": "der"
  },
  {
    "noun": "Kanadabalsam",
    "article": "der"
  },
  {
    "noun": "Kanadier",
    "article": "der"
  },
  {
    "noun": "Kanake",
    "article": "der"
  },
  {
    "noun": "Kanal",
    "article": "der"
  },
  {
    "noun": "Kanalarbeiter",
    "article": "der"
  },
  {
    "noun": "Kanaldeckel",
    "article": "der"
  },
  {
    "noun": "Kanalschalter",
    "article": "der"
  },
  {
    "noun": "Kanarienvogel",
    "article": "der"
  },
  {
    "noun": "Kandelaber",
    "article": "der"
  },
  {
    "noun": "Kandidat",
    "article": "der"
  },
  {
    "noun": "Kandis",
    "article": "der"
  },
  {
    "noun": "Kaninchenbau",
    "article": "der"
  },
  {
    "noun": "Kaninchenstall",
    "article": "der"
  },
  {
    "noun": "Kanister",
    "article": "der"
  },
  {
    "noun": "Kannibale",
    "article": "der"
  },
  {
    "noun": "Kannibalismus",
    "article": "der"
  },
  {
    "noun": "Kanonenofen",
    "article": "der"
  },
  {
    "noun": "Kanonenschuss",
    "article": "der"
  },
  {
    "noun": "Kanonier",
    "article": "der"
  },
  {
    "noun": "Kanter",
    "article": "der"
  },
  {
    "noun": "Kantersieg",
    "article": "der"
  },
  {
    "noun": "Kantor",
    "article": "der"
  },
  {
    "noun": "Kanufahrer",
    "article": "der"
  },
  {
    "noun": "Kanuslalom",
    "article": "der"
  },
  {
    "noun": "Kanusport",
    "article": "der"
  },
  {
    "noun": "Kanute",
    "article": "der"
  },
  {
    "noun": "Kanzelredner",
    "article": "der"
  },
  {
    "noun": "Kanzler",
    "article": "der"
  },
  {
    "noun": "Kanzlerbonus",
    "article": "der"
  },
  {
    "noun": "Kanzlerkandidat",
    "article": "der"
  },
  {
    "noun": "Kaolinit",
    "article": "der"
  },
  {
    "noun": "Kapaun",
    "article": "der"
  },
  {
    "noun": "Kapellmeister",
    "article": "der"
  },
  {
    "noun": "Kapernstrauch",
    "article": "der"
  },
  {
    "noun": "Kapitalanteil",
    "article": "der"
  },
  {
    "noun": "Kapitalbedarf",
    "article": "der"
  },
  {
    "noun": "Kapitalbesitz",
    "article": "der"
  },
  {
    "noun": "Kapitaleigner",
    "article": "der"
  },
  {
    "noun": "Kapitalertrag",
    "article": "der"
  },
  {
    "noun": "Kapitalexport",
    "article": "der"
  },
  {
    "noun": "Kapitalgeber",
    "article": "der"
  },
  {
    "noun": "Kapitalhirsch",
    "article": "der"
  },
  {
    "noun": "Kapitalismus",
    "article": "der"
  },
  {
    "noun": "Kapitalist",
    "article": "der"
  },
  {
    "noun": "Kapitalmarkt",
    "article": "der"
  },
  {
    "noun": "Kapitalverkehr",
    "article": "der"
  },
  {
    "noun": "Kapitalwert",
    "article": "der"
  },
  {
    "noun": "Kapitalzins",
    "article": "der"
  },
  {
    "noun": "Kaplan",
    "article": "der"
  },
  {
    "noun": "Kapok",
    "article": "der"
  },
  {
    "noun": "Kapokbaum",
    "article": "der"
  },
  {
    "noun": "Kappzaum",
    "article": "der"
  },
  {
    "noun": "Kapuziner",
    "article": "der"
  },
  {
    "noun": "Karabiner",
    "article": "der"
  },
  {
    "noun": "Karabinerhaken",
    "article": "der"
  },
  {
    "noun": "Karabinier",
    "article": "der"
  },
  {
    "noun": "Karamellpudding",
    "article": "der"
  },
  {
    "noun": "Karateschlag",
    "article": "der"
  },
  {
    "noun": "Karbunkel",
    "article": "der"
  },
  {
    "noun": "Kardanantrieb",
    "article": "der"
  },
  {
    "noun": "Kardinal",
    "article": "der"
  },
  {
    "noun": "Kardinalbischof",
    "article": "der"
  },
  {
    "noun": "Kardinalfehler",
    "article": "der"
  },
  {
    "noun": "Kardinalpunkt",
    "article": "der"
  },
  {
    "noun": "Kardinalshut",
    "article": "der"
  },
  {
    "noun": "Kardiograf",
    "article": "der"
  },
  {
    "noun": "Kardiologe",
    "article": "der"
  },
  {
    "noun": "Karfiol",
    "article": "der"
  },
  {
    "noun": "Karfreitag",
    "article": "der"
  },
  {
    "noun": "Kargo",
    "article": "der"
  },
  {
    "noun": "Karikaturist",
    "article": "der"
  },
  {
    "noun": "Karmelit",
    "article": "der"
  },
  {
    "noun": "Karmeliter",
    "article": "der"
  },
  {
    "noun": "Karneol",
    "article": "der"
  },
  {
    "noun": "Karner",
    "article": "der"
  },
  {
    "noun": "Karneval",
    "article": "der"
  },
  {
    "noun": "Karnevalist",
    "article": "der"
  },
  {
    "noun": "Karnevalsumzug",
    "article": "der"
  },
  {
    "noun": "Karnevalszug",
    "article": "der"
  },
  {
    "noun": "Karosseriebau",
    "article": "der"
  },
  {
    "noun": "Karosseriebauer",
    "article": "der"
  },
  {
    "noun": "Karpfen",
    "article": "der"
  },
  {
    "noun": "Karpfenteich",
    "article": "der"
  },
  {
    "noun": "Karrengaul",
    "article": "der"
  },
  {
    "noun": "Karrieremacher",
    "article": "der"
  },
  {
    "noun": "Karrierestau",
    "article": "der"
  },
  {
    "noun": "Karrierismus",
    "article": "der"
  },
  {
    "noun": "Karrierist",
    "article": "der"
  },
  {
    "noun": "Karst",
    "article": "der"
  },
  {
    "noun": "Karteikasten",
    "article": "der"
  },
  {
    "noun": "Kartellvertrag",
    "article": "der"
  },
  {
    "noun": "Kartenblock",
    "article": "der"
  },
  {
    "noun": "Kartenbrief",
    "article": "der"
  },
  {
    "noun": "Kartenlocher",
    "article": "der"
  },
  {
    "noun": "Kartenspieler",
    "article": "der"
  },
  {
    "noun": "Kartentisch",
    "article": "der"
  },
  {
    "noun": "Kartenverkauf",
    "article": "der"
  },
  {
    "noun": "Kartenvorverkauf",
    "article": "der"
  },
  {
    "noun": "Kartenzeichner",
    "article": "der"
  },
  {
    "noun": "Kartoffelbranntwein",
    "article": "der"
  },
  {
    "noun": "Kartoffelbrei",
    "article": "der"
  },
  {
    "noun": "Kartoffelchip",
    "article": "der"
  },
  {
    "noun": "Kartoffelpuffer",
    "article": "der"
  },
  {
    "noun": "Kartoffelsack",
    "article": "der"
  },
  {
    "noun": "Kartoffelsalat",
    "article": "der"
  },
  {
    "noun": "Kartoffelteig",
    "article": "der"
  },
  {
    "noun": "Kartograf",
    "article": "der"
  },
  {
    "noun": "Karton",
    "article": "der"
  },
  {
    "noun": "Kasache",
    "article": "der"
  },
  {
    "noun": "Kasack",
    "article": "der"
  },
  {
    "noun": "Kaschmirschal",
    "article": "der"
  },
  {
    "noun": "Kascholong",
    "article": "der"
  },
  {
    "noun": "Kasernenhof",
    "article": "der"
  },
  {
    "noun": "Kasper",
    "article": "der"
  },
  {
    "noun": "Kasperl",
    "article": "der"
  },
  {
    "noun": "Kassakurs",
    "article": "der"
  },
  {
    "noun": "Kassamarkt",
    "article": "der"
  },
  {
    "noun": "Kassapreis",
    "article": "der"
  },
  {
    "noun": "Kassationshof",
    "article": "der"
  },
  {
    "noun": "Kassenarzt",
    "article": "der"
  },
  {
    "noun": "Kassenbericht",
    "article": "der"
  },
  {
    "noun": "Kassenbestand",
    "article": "der"
  },
  {
    "noun": "Kassenbon",
    "article": "der"
  },
  {
    "noun": "Kassenerfolg",
    "article": "der"
  },
  {
    "noun": "Kassenpatient",
    "article": "der"
  },
  {
    "noun": "Kassenraum",
    "article": "der"
  },
  {
    "noun": "Kassenschalter",
    "article": "der"
  },
  {
    "noun": "Kassenschein",
    "article": "der"
  },
  {
    "noun": "Kassenschlager",
    "article": "der"
  },
  {
    "noun": "Kassensturz",
    "article": "der"
  },
  {
    "noun": "Kassenverwalter",
    "article": "der"
  },
  {
    "noun": "Kassenwart",
    "article": "der"
  },
  {
    "noun": "Kassenzettel",
    "article": "der"
  },
  {
    "noun": "Kassettenfilm",
    "article": "der"
  },
  {
    "noun": "Kassettenrecorder",
    "article": "der"
  },
  {
    "noun": "Kassiber",
    "article": "der"
  },
  {
    "noun": "Kassierer",
    "article": "der"
  },
  {
    "noun": "Kassiterit",
    "article": "der"
  },
  {
    "noun": "Kastanienbaum",
    "article": "der"
  },
  {
    "noun": "Kastellan",
    "article": "der"
  },
  {
    "noun": "Kasten",
    "article": "der"
  },
  {
    "noun": "Kastengeist",
    "article": "der"
  },
  {
    "noun": "Kastenwagen",
    "article": "der"
  },
  {
    "noun": "Kastorhut",
    "article": "der"
  },
  {
    "noun": "Kastrat",
    "article": "der"
  },
  {
    "noun": "Kastrationskomplex",
    "article": "der"
  },
  {
    "noun": "Kasuar",
    "article": "der"
  },
  {
    "noun": "Kasus",
    "article": "der"
  },
  {
    "noun": "Katabolismus",
    "article": "der"
  },
  {
    "noun": "Katafalk",
    "article": "der"
  },
  {
    "noun": "Kataklysmus",
    "article": "der"
  },
  {
    "noun": "Katalog",
    "article": "der"
  },
  {
    "noun": "Katalogpreis",
    "article": "der"
  },
  {
    "noun": "Katalysator",
    "article": "der"
  },
  {
    "noun": "Katapultstart",
    "article": "der"
  },
  {
    "noun": "Katarr",
    "article": "der"
  },
  {
    "noun": "Katasterauszug",
    "article": "der"
  },
  {
    "noun": "Katasterplan",
    "article": "der"
  },
  {
    "noun": "Katastrophenalarm",
    "article": "der"
  },
  {
    "noun": "Katastrophenfall",
    "article": "der"
  },
  {
    "noun": "Katastrophenschutz",
    "article": "der"
  },
  {
    "noun": "Katechismus",
    "article": "der"
  },
  {
    "noun": "Kater",
    "article": "der"
  },
  {
    "noun": "Katheter",
    "article": "der"
  },
  {
    "noun": "Katheterismus",
    "article": "der"
  },
  {
    "noun": "Kathodenfall",
    "article": "der"
  },
  {
    "noun": "Kathodenstrahl",
    "article": "der"
  },
  {
    "noun": "Kathodenstrahloszillograf",
    "article": "der"
  },
  {
    "noun": "Katholik",
    "article": "der"
  },
  {
    "noun": "Katholizismus",
    "article": "der"
  },
  {
    "noun": "Katholyt",
    "article": "der"
  },
  {
    "noun": "Kattun",
    "article": "der"
  },
  {
    "noun": "Kattundruck",
    "article": "der"
  },
  {
    "noun": "Katzenfreund",
    "article": "der"
  },
  {
    "noun": "Katzenhai",
    "article": "der"
  },
  {
    "noun": "Katzenjammer",
    "article": "der"
  },
  {
    "noun": "Katzenkopf",
    "article": "der"
  },
  {
    "noun": "Katzenliebhaber",
    "article": "der"
  },
  {
    "noun": "Katzenschwanz",
    "article": "der"
  },
  {
    "noun": "Katzentisch",
    "article": "der"
  },
  {
    "noun": "Kauakt",
    "article": "der"
  },
  {
    "noun": "Kauer",
    "article": "der"
  },
  {
    "noun": "Kauerstart",
    "article": "der"
  },
  {
    "noun": "Kauf",
    "article": "der"
  },
  {
    "noun": "Kaufanreiz",
    "article": "der"
  },
  {
    "noun": "Kaufauftrag",
    "article": "der"
  },
  {
    "noun": "Kaufbetrag",
    "article": "der"
  },
  {
    "noun": "Kaufbrief",
    "article": "der"
  },
  {
    "noun": "Kaufentschluss",
    "article": "der"
  },
  {
    "noun": "Kauffahrer",
    "article": "der"
  },
  {
    "noun": "Kaufgegenstand",
    "article": "der"
  },
  {
    "noun": "Kaufhausdetektiv",
    "article": "der"
  },
  {
    "noun": "Kaufinteressent",
    "article": "der"
  },
  {
    "noun": "Kaufladen",
    "article": "der"
  },
  {
    "noun": "Kaufmann",
    "article": "der"
  },
  {
    "noun": "Kaufmannsladen",
    "article": "der"
  },
  {
    "noun": "Kaufpreis",
    "article": "der"
  },
  {
    "noun": "Kaufvertrag",
    "article": "der"
  },
  {
    "noun": "Kaufwert",
    "article": "der"
  },
  {
    "noun": "Kaufzwang",
    "article": "der"
  },
  {
    "noun": "Kaukasier",
    "article": "der"
  },
  {
    "noun": "Kaukasus",
    "article": "der"
  },
  {
    "noun": "Kaulbarsch",
    "article": "der"
  },
  {
    "noun": "Kausalnexus",
    "article": "der"
  },
  {
    "noun": "Kausalsatz",
    "article": "der"
  },
  {
    "noun": "Kausalzusammenhang",
    "article": "der"
  },
  {
    "noun": "Kaustobiolith",
    "article": "der"
  },
  {
    "noun": "Kautabak",
    "article": "der"
  },
  {
    "noun": "Kautschuk",
    "article": "der"
  },
  {
    "noun": "Kautschukbaum",
    "article": "der"
  },
  {
    "noun": "Kauz",
    "article": "der"
  },
  {
    "noun": "Kavalier",
    "article": "der"
  },
  {
    "noun": "Kavallerist",
    "article": "der"
  },
  {
    "noun": "Kaviar",
    "article": "der"
  },
  {
    "noun": "Kebab",
    "article": "der"
  },
  {
    "noun": "Keder",
    "article": "der"
  },
  {
    "noun": "Kefir",
    "article": "der"
  },
  {
    "noun": "Kegel",
    "article": "der"
  },
  {
    "noun": "Kegelbrecher",
    "article": "der"
  },
  {
    "noun": "Kegelclub",
    "article": "der"
  },
  {
    "noun": "Kegelschnitt",
    "article": "der"
  },
  {
    "noun": "Kegelstumpf",
    "article": "der"
  },
  {
    "noun": "Kegler",
    "article": "der"
  },
  {
    "noun": "Kehldeckel",
    "article": "der"
  },
  {
    "noun": "Kehlhobel",
    "article": "der"
  },
  {
    "noun": "Kehlkopf",
    "article": "der"
  },
  {
    "noun": "Kehlkopfkrebs",
    "article": "der"
  },
  {
    "noun": "Kehlkopflaut",
    "article": "der"
  },
  {
    "noun": "Kehllaut",
    "article": "der"
  },
  {
    "noun": "Kehrer",
    "article": "der"
  },
  {
    "noun": "Kehrichteimer",
    "article": "der"
  },
  {
    "noun": "Kehrichthaufen",
    "article": "der"
  },
  {
    "noun": "Kehrreim",
    "article": "der"
  },
  {
    "noun": "Kehrwert",
    "article": "der"
  },
  {
    "noun": "Kehrwisch",
    "article": "der"
  },
  {
    "noun": "Keil",
    "article": "der"
  },
  {
    "noun": "Keilabsatz",
    "article": "der"
  },
  {
    "noun": "Keiler",
    "article": "der"
  },
  {
    "noun": "Keilriemen",
    "article": "der"
  },
  {
    "noun": "Keim",
    "article": "der"
  },
  {
    "noun": "Keimling",
    "article": "der"
  },
  {
    "noun": "Kelch",
    "article": "der"
  },
  {
    "noun": "Keller",
    "article": "der"
  },
  {
    "noun": "Kellermeister",
    "article": "der"
  },
  {
    "noun": "Kellerraum",
    "article": "der"
  },
  {
    "noun": "Kellner",
    "article": "der"
  },
  {
    "noun": "Kelte",
    "article": "der"
  },
  {
    "noun": "Kenianer",
    "article": "der"
  },
  {
    "noun": "Kennbuchstabe",
    "article": "der"
  },
  {
    "noun": "Kenner",
    "article": "der"
  },
  {
    "noun": "Kennfaden",
    "article": "der"
  },
  {
    "noun": "Kenntnisstand",
    "article": "der"
  },
  {
    "noun": "Kennwert",
    "article": "der"
  },
  {
    "noun": "Kepheus",
    "article": "der"
  },
  {
    "noun": "Keramiker",
    "article": "der"
  },
  {
    "noun": "Keratophyr",
    "article": "der"
  },
  {
    "noun": "Kerbel",
    "article": "der"
  },
  {
    "noun": "Kerbnagel",
    "article": "der"
  },
  {
    "noun": "Kerbschlagbiegeversuch",
    "article": "der"
  },
  {
    "noun": "Kerbstift",
    "article": "der"
  },
  {
    "noun": "Kerker",
    "article": "der"
  },
  {
    "noun": "Kerkermeister",
    "article": "der"
  },
  {
    "noun": "Kerl",
    "article": "der"
  },
  {
    "noun": "Kernantrieb",
    "article": "der"
  },
  {
    "noun": "Kernbaustein",
    "article": "der"
  },
  {
    "noun": "Kernbereich",
    "article": "der"
  },
  {
    "noun": "Kernbohrer",
    "article": "der"
  },
  {
    "noun": "Kernbrennstoff",
    "article": "der"
  },
  {
    "noun": "Kerndurchmesser",
    "article": "der"
  },
  {
    "noun": "Kernenergieantrieb",
    "article": "der"
  },
  {
    "noun": "Kernfotoeffekt",
    "article": "der"
  },
  {
    "noun": "Kerngedanke",
    "article": "der"
  },
  {
    "noun": "Kerningenieur",
    "article": "der"
  },
  {
    "noun": "Kernkraftgegner",
    "article": "der"
  },
  {
    "noun": "Kernphysiker",
    "article": "der"
  },
  {
    "noun": "Kernpunkt",
    "article": "der"
  },
  {
    "noun": "Kernreaktor",
    "article": "der"
  },
  {
    "noun": "Kernschatten",
    "article": "der"
  },
  {
    "noun": "Kernspruch",
    "article": "der"
  },
  {
    "noun": "Kernsprung",
    "article": "der"
  },
  {
    "noun": "Kernwaffengegner",
    "article": "der"
  },
  {
    "noun": "Kernzerfall",
    "article": "der"
  },
  {
    "noun": "Kerzenhalter",
    "article": "der"
  },
  {
    "noun": "Kerzenleuchter",
    "article": "der"
  },
  {
    "noun": "Kerzenmacher",
    "article": "der"
  },
  {
    "noun": "Kerzenschein",
    "article": "der"
  },
  {
    "noun": "Kerzenzieher",
    "article": "der"
  },
  {
    "noun": "Kescher",
    "article": "der"
  },
  {
    "noun": "Kessel",
    "article": "der"
  },
  {
    "noun": "Kesselboden",
    "article": "der"
  },
  {
    "noun": "Kesseldruck",
    "article": "der"
  },
  {
    "noun": "Kesselflicker",
    "article": "der"
  },
  {
    "noun": "Kesselraum",
    "article": "der"
  },
  {
    "noun": "Kesselschmied",
    "article": "der"
  },
  {
    "noun": "Kesselstein",
    "article": "der"
  },
  {
    "noun": "Kesselwagen",
    "article": "der"
  },
  {
    "noun": "Kettenantrieb",
    "article": "der"
  },
  {
    "noun": "Kettenbrief",
    "article": "der"
  },
  {
    "noun": "Kettenbruch",
    "article": "der"
  },
  {
    "noun": "Kettenhandel",
    "article": "der"
  },
  {
    "noun": "Kettenhund",
    "article": "der"
  },
  {
    "noun": "Kettenladen",
    "article": "der"
  },
  {
    "noun": "Kettenpanzer",
    "article": "der"
  },
  {
    "noun": "Kettenraucher",
    "article": "der"
  },
  {
    "noun": "Kettenschutz",
    "article": "der"
  },
  {
    "noun": "Kettenstich",
    "article": "der"
  },
  {
    "noun": "Kettfaden",
    "article": "der"
  },
  {
    "noun": "Ketzer",
    "article": "der"
  },
  {
    "noun": "Keuchhusten",
    "article": "der"
  },
  {
    "noun": "Keulenpilz",
    "article": "der"
  },
  {
    "noun": "Keulenschlag",
    "article": "der"
  },
  {
    "noun": "Keuper",
    "article": "der"
  },
  {
    "noun": "Keynesianismus",
    "article": "der"
  },
  {
    "noun": "Kibbuz",
    "article": "der"
  },
  {
    "noun": "Kick",
    "article": "der"
  },
  {
    "noun": "Kicker",
    "article": "der"
  },
  {
    "noun": "Kickstarter",
    "article": "der"
  },
  {
    "noun": "Kiebitz",
    "article": "der"
  },
  {
    "noun": "Kiefer",
    "article": "der"
  },
  {
    "noun": "Kieferbruch",
    "article": "der"
  },
  {
    "noun": "Kieferchirurg",
    "article": "der"
  },
  {
    "noun": "Kieferknochen",
    "article": "der"
  },
  {
    "noun": "Kiefernkreuzschnabel",
    "article": "der"
  },
  {
    "noun": "Kiefernwald",
    "article": "der"
  },
  {
    "noun": "Kiefernzapfen",
    "article": "der"
  },
  {
    "noun": "Kielraum",
    "article": "der"
  },
  {
    "noun": "Kienapfel",
    "article": "der"
  },
  {
    "noun": "Kies",
    "article": "der"
  },
  {
    "noun": "Kiesbeton",
    "article": "der"
  },
  {
    "noun": "Kiesboden",
    "article": "der"
  },
  {
    "noun": "Kiesel",
    "article": "der"
  },
  {
    "noun": "Kieselstein",
    "article": "der"
  },
  {
    "noun": "Kieserit",
    "article": "der"
  },
  {
    "noun": "Kieshaufen",
    "article": "der"
  },
  {
    "noun": "Kiessand",
    "article": "der"
  },
  {
    "noun": "Kiesweg",
    "article": "der"
  },
  {
    "noun": "Kiez",
    "article": "der"
  },
  {
    "noun": "Kiffer",
    "article": "der"
  },
  {
    "noun": "Killer",
    "article": "der"
  },
  {
    "noun": "Killersatellit",
    "article": "der"
  },
  {
    "noun": "Killerwal",
    "article": "der"
  },
  {
    "noun": "Kilometer",
    "article": "der"
  },
  {
    "noun": "Kilometerstand",
    "article": "der"
  },
  {
    "noun": "Kilometerstein",
    "article": "der"
  },
  {
    "noun": "Kilt",
    "article": "der"
  },
  {
    "noun": "Kimberlit",
    "article": "der"
  },
  {
    "noun": "Kimono",
    "article": "der"
  },
  {
    "noun": "Kinematiker",
    "article": "der"
  },
  {
    "noun": "Kinematograf",
    "article": "der"
  },
  {
    "noun": "Kinnbart",
    "article": "der"
  },
  {
    "noun": "Kinnhaken",
    "article": "der"
  },
  {
    "noun": "Kinnriemen",
    "article": "der"
  },
  {
    "noun": "Kinobesitzer",
    "article": "der"
  },
  {
    "noun": "Kinobesucher",
    "article": "der"
  },
  {
    "noun": "Kiosk",
    "article": "der"
  },
  {
    "noun": "Kipper",
    "article": "der"
  },
  {
    "noun": "Kippkarren",
    "article": "der"
  },
  {
    "noun": "Kippschalter",
    "article": "der"
  },
  {
    "noun": "Kippwagen",
    "article": "der"
  },
  {
    "noun": "Kirchenbesuch",
    "article": "der"
  },
  {
    "noun": "Kirchenbesucher",
    "article": "der"
  },
  {
    "noun": "Kirchenchor",
    "article": "der"
  },
  {
    "noun": "Kirchendiener",
    "article": "der"
  },
  {
    "noun": "Kirchenrat",
    "article": "der"
  },
  {
    "noun": "Kirchenstaat",
    "article": "der"
  },
  {
    "noun": "Kirchenstuhl",
    "article": "der"
  },
  {
    "noun": "Kirchentag",
    "article": "der"
  },
  {
    "noun": "Kirchenvorstand",
    "article": "der"
  },
  {
    "noun": "Kirchgang",
    "article": "der"
  },
  {
    "noun": "Kirchhof",
    "article": "der"
  },
  {
    "noun": "Kirchtag",
    "article": "der"
  },
  {
    "noun": "Kirchturm",
    "article": "der"
  },
  {
    "noun": "Kirgise",
    "article": "der"
  },
  {
    "noun": "Kirschbaum",
    "article": "der"
  },
  {
    "noun": "Kirschkern",
    "article": "der"
  },
  {
    "noun": "Kirschkuchen",
    "article": "der"
  },
  {
    "noun": "Kirschsaft",
    "article": "der"
  },
  {
    "noun": "Kirschschnaps",
    "article": "der"
  },
  {
    "noun": "Kirschstein",
    "article": "der"
  },
  {
    "noun": "Kissenbezug",
    "article": "der"
  },
  {
    "noun": "Kitsch",
    "article": "der"
  },
  {
    "noun": "Kitt",
    "article": "der"
  },
  {
    "noun": "Kittel",
    "article": "der"
  },
  {
    "noun": "Kitzel",
    "article": "der"
  },
  {
    "noun": "Kitzler",
    "article": "der"
  },
  {
    "noun": "Klacks",
    "article": "der"
  },
  {
    "noun": "Kladderadatsch",
    "article": "der"
  },
  {
    "noun": "Klageweg",
    "article": "der"
  },
  {
    "noun": "Klamauk",
    "article": "der"
  },
  {
    "noun": "Klammeraffe",
    "article": "der"
  },
  {
    "noun": "Klammerausdruck",
    "article": "der"
  },
  {
    "noun": "Klang",
    "article": "der"
  },
  {
    "noun": "Klangeffekt",
    "article": "der"
  },
  {
    "noun": "Klapf",
    "article": "der"
  },
  {
    "noun": "Klappenschrank",
    "article": "der"
  },
  {
    "noun": "Klappentext",
    "article": "der"
  },
  {
    "noun": "Klapperkasten",
    "article": "der"
  },
  {
    "noun": "Klapperstorch",
    "article": "der"
  },
  {
    "noun": "Klapphut",
    "article": "der"
  },
  {
    "noun": "Klappsitz",
    "article": "der"
  },
  {
    "noun": "Klappstuhl",
    "article": "der"
  },
  {
    "noun": "Klapptisch",
    "article": "der"
  },
  {
    "noun": "Klarapfel",
    "article": "der"
  },
  {
    "noun": "Klarinettist",
    "article": "der"
  },
  {
    "noun": "Klartext",
    "article": "der"
  },
  {
    "noun": "Klassenaufsatz",
    "article": "der"
  },
  {
    "noun": "Klassenfeind",
    "article": "der"
  },
  {
    "noun": "Klassenkamerad",
    "article": "der"
  },
  {
    "noun": "Klassenkampf",
    "article": "der"
  },
  {
    "noun": "Klassenkonflikt",
    "article": "der"
  },
  {
    "noun": "Klassenlehrer",
    "article": "der"
  },
  {
    "noun": "Klassenraum",
    "article": "der"
  },
  {
    "noun": "Klassensprecher",
    "article": "der"
  },
  {
    "noun": "Klassenunterschied",
    "article": "der"
  },
  {
    "noun": "Klassiker",
    "article": "der"
  },
  {
    "noun": "Klassizismus",
    "article": "der"
  },
  {
    "noun": "Klassizist",
    "article": "der"
  },
  {
    "noun": "Klatsch",
    "article": "der"
  },
  {
    "noun": "Klatschkolumnist",
    "article": "der"
  },
  {
    "noun": "Klatschmohn",
    "article": "der"
  },
  {
    "noun": "Klausner",
    "article": "der"
  },
  {
    "noun": "Klavierauszug",
    "article": "der"
  },
  {
    "noun": "Klavierhocker",
    "article": "der"
  },
  {
    "noun": "Klavierlehrer",
    "article": "der"
  },
  {
    "noun": "Klavierspieler",
    "article": "der"
  },
  {
    "noun": "Klavierstuhl",
    "article": "der"
  },
  {
    "noun": "Kleber",
    "article": "der"
  },
  {
    "noun": "Klebestift",
    "article": "der"
  },
  {
    "noun": "Klebestreifen",
    "article": "der"
  },
  {
    "noun": "Klebezettel",
    "article": "der"
  },
  {
    "noun": "Klebstoff",
    "article": "der"
  },
  {
    "noun": "Klebstreifen",
    "article": "der"
  },
  {
    "noun": "Kleckerkram",
    "article": "der"
  },
  {
    "noun": "Klecks",
    "article": "der"
  },
  {
    "noun": "Klee",
    "article": "der"
  },
  {
    "noun": "Kleeblattbogen",
    "article": "der"
  },
  {
    "noun": "Klei",
    "article": "der"
  },
  {
    "noun": "Kleiber",
    "article": "der"
  },
  {
    "noun": "Kleiboden",
    "article": "der"
  },
  {
    "noun": "Kleiderhaken",
    "article": "der"
  },
  {
    "noun": "Kleiderkasten",
    "article": "der"
  },
  {
    "noun": "Kleiderrock",
    "article": "der"
  },
  {
    "noun": "Kleiderschrank",
    "article": "der"
  },
  {
    "noun": "Kleinbauer",
    "article": "der"
  },
  {
    "noun": "Kleinbetrieb",
    "article": "der"
  },
  {
    "noun": "Kleinbuchstabe",
    "article": "der"
  },
  {
    "noun": "Kleindarsteller",
    "article": "der"
  },
  {
    "noun": "Kleingeist",
    "article": "der"
  },
  {
    "noun": "Kleinhandel",
    "article": "der"
  },
  {
    "noun": "Kleinkram",
    "article": "der"
  },
  {
    "noun": "Kleinkredit",
    "article": "der"
  },
  {
    "noun": "Kleinkreis",
    "article": "der"
  },
  {
    "noun": "Kleinmotor",
    "article": "der"
  },
  {
    "noun": "Kleinrechner",
    "article": "der"
  },
  {
    "noun": "Kleinsparer",
    "article": "der"
  },
  {
    "noun": "Kleintransporter",
    "article": "der"
  },
  {
    "noun": "Kleinunternehmer",
    "article": "der"
  },
  {
    "noun": "Kleinverbraucher",
    "article": "der"
  },
  {
    "noun": "Kleinverkauf",
    "article": "der"
  },
  {
    "noun": "Kleinwagen",
    "article": "der"
  },
  {
    "noun": "Kleister",
    "article": "der"
  },
  {
    "noun": "Klempner",
    "article": "der"
  },
  {
    "noun": "Klepper",
    "article": "der"
  },
  {
    "noun": "Kleptomane",
    "article": "der"
  },
  {
    "noun": "Klerikalismus",
    "article": "der"
  },
  {
    "noun": "Kleriker",
    "article": "der"
  },
  {
    "noun": "Klerus",
    "article": "der"
  },
  {
    "noun": "Kletterer",
    "article": "der"
  },
  {
    "noun": "Klettergarten",
    "article": "der"
  },
  {
    "noun": "Klettverschluss",
    "article": "der"
  },
  {
    "noun": "Klick",
    "article": "der"
  },
  {
    "noun": "Klient",
    "article": "der"
  },
  {
    "noun": "Klimaschutz",
    "article": "der"
  },
  {
    "noun": "Klimawechsel",
    "article": "der"
  },
  {
    "noun": "Klimbim",
    "article": "der"
  },
  {
    "noun": "Klimmzug",
    "article": "der"
  },
  {
    "noun": "Klingelknopf",
    "article": "der"
  },
  {
    "noun": "Klingelton",
    "article": "der"
  },
  {
    "noun": "Klingklang",
    "article": "der"
  },
  {
    "noun": "Kliniker",
    "article": "der"
  },
  {
    "noun": "Klinker",
    "article": "der"
  },
  {
    "noun": "Klinograf",
    "article": "der"
  },
  {
    "noun": "Klipper",
    "article": "der"
  },
  {
    "noun": "Klips",
    "article": "der"
  },
  {
    "noun": "Klirrfaktor",
    "article": "der"
  },
  {
    "noun": "Kloben",
    "article": "der"
  },
  {
    "noun": "Klodeckel",
    "article": "der"
  },
  {
    "noun": "Klon",
    "article": "der"
  },
  {
    "noun": "Klopfer",
    "article": "der"
  },
  {
    "noun": "Klops",
    "article": "der"
  },
  {
    "noun": "Klosettdeckel",
    "article": "der"
  },
  {
    "noun": "Klosterbruder",
    "article": "der"
  },
  {
    "noun": "Klostergarten",
    "article": "der"
  },
  {
    "noun": "Klotz",
    "article": "der"
  },
  {
    "noun": "Klump",
    "article": "der"
  },
  {
    "noun": "Klunker",
    "article": "der"
  },
  {
    "noun": "Knabe",
    "article": "der"
  },
  {
    "noun": "Knabenchor",
    "article": "der"
  },
  {
    "noun": "Knack",
    "article": "der"
  },
  {
    "noun": "Knacker",
    "article": "der"
  },
  {
    "noun": "Knacki",
    "article": "der"
  },
  {
    "noun": "Knacklaut",
    "article": "der"
  },
  {
    "noun": "Knackpunkt",
    "article": "der"
  },
  {
    "noun": "Knacks",
    "article": "der"
  },
  {
    "noun": "Knall",
    "article": "der"
  },
  {
    "noun": "Knalleffekt",
    "article": "der"
  },
  {
    "noun": "Knaller",
    "article": "der"
  },
  {
    "noun": "Knallfrosch",
    "article": "der"
  },
  {
    "noun": "Knallkopf",
    "article": "der"
  },
  {
    "noun": "Knappe",
    "article": "der"
  },
  {
    "noun": "Knastbruder",
    "article": "der"
  },
  {
    "noun": "Knauf",
    "article": "der"
  },
  {
    "noun": "Knebel",
    "article": "der"
  },
  {
    "noun": "Knecht",
    "article": "der"
  },
  {
    "noun": "Kneifer",
    "article": "der"
  },
  {
    "noun": "Kneipenwirt",
    "article": "der"
  },
  {
    "noun": "Knetgummi",
    "article": "der"
  },
  {
    "noun": "Knick",
    "article": "der"
  },
  {
    "noun": "Knicker",
    "article": "der"
  },
  {
    "noun": "Knicks",
    "article": "der"
  },
  {
    "noun": "Kniefall",
    "article": "der"
  },
  {
    "noun": "Kniehebel",
    "article": "der"
  },
  {
    "noun": "Knieschoner",
    "article": "der"
  },
  {
    "noun": "Kniestrumpf",
    "article": "der"
  },
  {
    "noun": "Kniff",
    "article": "der"
  },
  {
    "noun": "Knigge",
    "article": "der"
  },
  {
    "noun": "Knilch",
    "article": "der"
  },
  {
    "noun": "Knirps",
    "article": "der"
  },
  {
    "noun": "Knittelvers",
    "article": "der"
  },
  {
    "noun": "Knobel",
    "article": "der"
  },
  {
    "noun": "Knoblauch",
    "article": "der"
  },
  {
    "noun": "Knochen",
    "article": "der"
  },
  {
    "noun": "Knochenbau",
    "article": "der"
  },
  {
    "noun": "Knochenbrand",
    "article": "der"
  },
  {
    "noun": "Knochenbruch",
    "article": "der"
  },
  {
    "noun": "Knochenfisch",
    "article": "der"
  },
  {
    "noun": "Knochenschwund",
    "article": "der"
  },
  {
    "noun": "Knochensplitter",
    "article": "der"
  },
  {
    "noun": "Knochentumor",
    "article": "der"
  },
  {
    "noun": "Knopf",
    "article": "der"
  },
  {
    "noun": "Knopfdruck",
    "article": "der"
  },
  {
    "noun": "Knopper",
    "article": "der"
  },
  {
    "noun": "Knorpel",
    "article": "der"
  },
  {
    "noun": "Knorpelfisch",
    "article": "der"
  },
  {
    "noun": "Knorren",
    "article": "der"
  },
  {
    "noun": "Knotenpunkt",
    "article": "der"
  },
  {
    "noun": "Knotenstock",
    "article": "der"
  },
  {
    "noun": "Knuff",
    "article": "der"
  },
  {
    "noun": "Knurrhahn",
    "article": "der"
  },
  {
    "noun": "Knurrlaut",
    "article": "der"
  },
  {
    "noun": "Knutschfleck",
    "article": "der"
  },
  {
    "noun": "Koala",
    "article": "der"
  },
  {
    "noun": "Koalitionspartner",
    "article": "der"
  },
  {
    "noun": "Kobaltglanz",
    "article": "der"
  },
  {
    "noun": "Kobel",
    "article": "der"
  },
  {
    "noun": "Koben",
    "article": "der"
  },
  {
    "noun": "Kobold",
    "article": "der"
  },
  {
    "noun": "Koch",
    "article": "der"
  },
  {
    "noun": "Kochapfel",
    "article": "der"
  },
  {
    "noun": "Kocher",
    "article": "der"
  },
  {
    "noun": "Kochherd",
    "article": "der"
  },
  {
    "noun": "Kochkessel",
    "article": "der"
  },
  {
    "noun": "Kochkurs",
    "article": "der"
  },
  {
    "noun": "Kochsalat",
    "article": "der"
  },
  {
    "noun": "Kochschinken",
    "article": "der"
  },
  {
    "noun": "Kochtopf",
    "article": "der"
  },
  {
    "noun": "Kodex",
    "article": "der"
  },
  {
    "noun": "Koeffizient",
    "article": "der"
  },
  {
    "noun": "Koffer",
    "article": "der"
  },
  {
    "noun": "Kofferraum",
    "article": "der"
  },
  {
    "noun": "Kog",
    "article": "der"
  },
  {
    "noun": "Kognak",
    "article": "der"
  },
  {
    "noun": "Kohl",
    "article": "der"
  },
  {
    "noun": "Kohlefaden",
    "article": "der"
  },
  {
    "noun": "Kohlenbergbau",
    "article": "der"
  },
  {
    "noun": "Kohleneimer",
    "article": "der"
  },
  {
    "noun": "Kohlengrus",
    "article": "der"
  },
  {
    "noun": "Kohlenkasten",
    "article": "der"
  },
  {
    "noun": "Kohlenkeller",
    "article": "der"
  },
  {
    "noun": "Kohlenmeiler",
    "article": "der"
  },
  {
    "noun": "Kohlenstaub",
    "article": "der"
  },
  {
    "noun": "Kohlenstift",
    "article": "der"
  },
  {
    "noun": "Kohlenstoff",
    "article": "der"
  },
  {
    "noun": "Kohlenstoffgehalt",
    "article": "der"
  },
  {
    "noun": "Kohlenstoffring",
    "article": "der"
  },
  {
    "noun": "Kohlenwasserstoff",
    "article": "der"
  },
  {
    "noun": "Kohlestift",
    "article": "der"
  },
  {
    "noun": "Kohlkopf",
    "article": "der"
  },
  {
    "noun": "Kohlrabi",
    "article": "der"
  },
  {
    "noun": "Koitus",
    "article": "der"
  },
  {
    "noun": "Koker",
    "article": "der"
  },
  {
    "noun": "Kokillenguss",
    "article": "der"
  },
  {
    "noun": "Kokkus",
    "article": "der"
  },
  {
    "noun": "Kokon",
    "article": "der"
  },
  {
    "noun": "Kokosraspel",
    "article": "der"
  },
  {
    "noun": "Koks",
    "article": "der"
  },
  {
    "noun": "Kokser",
    "article": "der"
  },
  {
    "noun": "Koksgrus",
    "article": "der"
  },
  {
    "noun": "Koksofen",
    "article": "der"
  },
  {
    "noun": "Kolabaum",
    "article": "der"
  },
  {
    "noun": "Kolben",
    "article": "der"
  },
  {
    "noun": "Kolbenblitz",
    "article": "der"
  },
  {
    "noun": "Kolbenfresser",
    "article": "der"
  },
  {
    "noun": "Kolbenmotor",
    "article": "der"
  },
  {
    "noun": "Kolbenring",
    "article": "der"
  },
  {
    "noun": "Kolbenschieber",
    "article": "der"
  },
  {
    "noun": "Kolibri",
    "article": "der"
  },
  {
    "noun": "Kolk",
    "article": "der"
  },
  {
    "noun": "Kolkrabe",
    "article": "der"
  },
  {
    "noun": "Kollaborateur",
    "article": "der"
  },
  {
    "noun": "Kollaborator",
    "article": "der"
  },
  {
    "noun": "Kollaps",
    "article": "der"
  },
  {
    "noun": "Kollateralschaden",
    "article": "der"
  },
  {
    "noun": "Kollege",
    "article": "der"
  },
  {
    "noun": "Kollektivismus",
    "article": "der"
  },
  {
    "noun": "Kollektivist",
    "article": "der"
  },
  {
    "noun": "Kollektivvertrag",
    "article": "der"
  },
  {
    "noun": "Kollektor",
    "article": "der"
  },
  {
    "noun": "Kollergang",
    "article": "der"
  },
  {
    "noun": "Kollimator",
    "article": "der"
  },
  {
    "noun": "Kollisionskurs",
    "article": "der"
  },
  {
    "noun": "Kolonialherr",
    "article": "der"
  },
  {
    "noun": "Kolonialismus",
    "article": "der"
  },
  {
    "noun": "Kolonialist",
    "article": "der"
  },
  {
    "noun": "Kolonialwarenladen",
    "article": "der"
  },
  {
    "noun": "Kolonist",
    "article": "der"
  },
  {
    "noun": "Koloratursopran",
    "article": "der"
  },
  {
    "noun": "Koloss",
    "article": "der"
  },
  {
    "noun": "Kolumbianer",
    "article": "der"
  },
  {
    "noun": "Kolumnentitel",
    "article": "der"
  },
  {
    "noun": "Kolumnist",
    "article": "der"
  },
  {
    "noun": "Kombiwagen",
    "article": "der"
  },
  {
    "noun": "Komet",
    "article": "der"
  },
  {
    "noun": "Komfort",
    "article": "der"
  },
  {
    "noun": "Komiker",
    "article": "der"
  },
  {
    "noun": "Kommandant",
    "article": "der"
  },
  {
    "noun": "Kommandeur",
    "article": "der"
  },
  {
    "noun": "Kommanditist",
    "article": "der"
  },
  {
    "noun": "Kommandoruf",
    "article": "der"
  },
  {
    "noun": "Kommandostab",
    "article": "der"
  },
  {
    "noun": "Kommandoton",
    "article": "der"
  },
  {
    "noun": "Kommentar",
    "article": "der"
  },
  {
    "noun": "Kommentator",
    "article": "der"
  },
  {
    "noun": "Kommerz",
    "article": "der"
  },
  {
    "noun": "Kommilitone",
    "article": "der"
  },
  {
    "noun": "Kommissar",
    "article": "der"
  },
  {
    "noun": "Kommissionsagent",
    "article": "der"
  },
  {
    "noun": "Kommissionshandel",
    "article": "der"
  },
  {
    "noun": "Kommissionsvertrag",
    "article": "der"
  },
  {
    "noun": "Kommittent",
    "article": "der"
  },
  {
    "noun": "Kommodore",
    "article": "der"
  },
  {
    "noun": "Kommunalkredit",
    "article": "der"
  },
  {
    "noun": "Kommunalverband",
    "article": "der"
  },
  {
    "noun": "Kommunikant",
    "article": "der"
  },
  {
    "noun": "Kommunikationsdienst",
    "article": "der"
  },
  {
    "noun": "Kommunikationselektroniker",
    "article": "der"
  },
  {
    "noun": "Kommunikationssatellit",
    "article": "der"
  },
  {
    "noun": "Kommunismus",
    "article": "der"
  },
  {
    "noun": "Kommunist",
    "article": "der"
  },
  {
    "noun": "Kommutator",
    "article": "der"
  },
  {
    "noun": "Kompagnon",
    "article": "der"
  },
  {
    "noun": "Kompaniechef",
    "article": "der"
  },
  {
    "noun": "Kompaniefeldwebel",
    "article": "der"
  },
  {
    "noun": "Komparativ",
    "article": "der"
  },
  {
    "noun": "Komparativsatz",
    "article": "der"
  },
  {
    "noun": "Komparator",
    "article": "der"
  },
  {
    "noun": "Komparse",
    "article": "der"
  },
  {
    "noun": "Kompass",
    "article": "der"
  },
  {
    "noun": "Kompensator",
    "article": "der"
  },
  {
    "noun": "Kompetenzbereich",
    "article": "der"
  },
  {
    "noun": "Kompetenzkonflikt",
    "article": "der"
  },
  {
    "noun": "Komplettpreis",
    "article": "der"
  },
  {
    "noun": "Komplex",
    "article": "der"
  },
  {
    "noun": "Komplice",
    "article": "der"
  },
  {
    "noun": "Komponist",
    "article": "der"
  },
  {
    "noun": "Kompost",
    "article": "der"
  },
  {
    "noun": "Komposthaufen",
    "article": "der"
  },
  {
    "noun": "Kompressor",
    "article": "der"
  },
  {
    "noun": "Kompromissler",
    "article": "der"
  },
  {
    "noun": "Kompromissvorschlag",
    "article": "der"
  },
  {
    "noun": "Kondensationspunkt",
    "article": "der"
  },
  {
    "noun": "Kondensator",
    "article": "der"
  },
  {
    "noun": "Kondensor",
    "article": "der"
  },
  {
    "noun": "Kondensstreifen",
    "article": "der"
  },
  {
    "noun": "Konditor",
    "article": "der"
  },
  {
    "noun": "Kondor",
    "article": "der"
  },
  {
    "noun": "Konduktor",
    "article": "der"
  },
  {
    "noun": "Konfektionsanzug",
    "article": "der"
  },
  {
    "noun": "Konferenzraum",
    "article": "der"
  },
  {
    "noun": "Konferenzsaal",
    "article": "der"
  },
  {
    "noun": "Konferenzteilnehmer",
    "article": "der"
  },
  {
    "noun": "Konferenztisch",
    "article": "der"
  },
  {
    "noun": "Konfirmand",
    "article": "der"
  },
  {
    "noun": "Konfirmandenunterricht",
    "article": "der"
  },
  {
    "noun": "Konflikt",
    "article": "der"
  },
  {
    "noun": "Konfliktstoff",
    "article": "der"
  },
  {
    "noun": "Konformismus",
    "article": "der"
  },
  {
    "noun": "Konformist",
    "article": "der"
  },
  {
    "noun": "Kongo",
    "article": "der"
  },
  {
    "noun": "Kongolese",
    "article": "der"
  },
  {
    "noun": "Kongress",
    "article": "der"
  },
  {
    "noun": "Kongressteilnehmer",
    "article": "der"
  },
  {
    "noun": "Konjunktiv",
    "article": "der"
  },
  {
    "noun": "Konjunkturanstieg",
    "article": "der"
  },
  {
    "noun": "Konjunkturaufschwung",
    "article": "der"
  },
  {
    "noun": "Konjunkturbarometer",
    "article": "der"
  },
  {
    "noun": "Konjunkturzuschlag",
    "article": "der"
  },
  {
    "noun": "Konjunkturzyklus",
    "article": "der"
  },
  {
    "noun": "Konkavspiegel",
    "article": "der"
  },
  {
    "noun": "Konkurrent",
    "article": "der"
  },
  {
    "noun": "Konkurrenzbetrieb",
    "article": "der"
  },
  {
    "noun": "Konkurrenzdruck",
    "article": "der"
  },
  {
    "noun": "Konkurrenzkampf",
    "article": "der"
  },
  {
    "noun": "Konkurrenzneid",
    "article": "der"
  },
  {
    "noun": "Konkurs",
    "article": "der"
  },
  {
    "noun": "Konkursverwalter",
    "article": "der"
  },
  {
    "noun": "Konnektor",
    "article": "der"
  },
  {
    "noun": "Konquistador",
    "article": "der"
  },
  {
    "noun": "Konsens",
    "article": "der"
  },
  {
    "noun": "Konservatismus",
    "article": "der"
  },
  {
    "noun": "Konservator",
    "article": "der"
  },
  {
    "noun": "Konservierungsstoff",
    "article": "der"
  },
  {
    "noun": "Konsignant",
    "article": "der"
  },
  {
    "noun": "Konsignatar",
    "article": "der"
  },
  {
    "noun": "Konsonant",
    "article": "der"
  },
  {
    "noun": "Konsorte",
    "article": "der"
  },
  {
    "noun": "Konspekt",
    "article": "der"
  },
  {
    "noun": "Konstrukteur",
    "article": "der"
  },
  {
    "noun": "Konstruktionsfehler",
    "article": "der"
  },
  {
    "noun": "Konstruktivismus",
    "article": "der"
  },
  {
    "noun": "Konstruktivist",
    "article": "der"
  },
  {
    "noun": "Konsul",
    "article": "der"
  },
  {
    "noun": "Konsularagent",
    "article": "der"
  },
  {
    "noun": "Konsum",
    "article": "der"
  },
  {
    "noun": "Konsumartikel",
    "article": "der"
  },
  {
    "noun": "Konsument",
    "article": "der"
  },
  {
    "noun": "Konsumentenkredit",
    "article": "der"
  },
  {
    "noun": "Konsumerismus",
    "article": "der"
  },
  {
    "noun": "Konsumverein",
    "article": "der"
  },
  {
    "noun": "Konsumverzicht",
    "article": "der"
  },
  {
    "noun": "Kontaktabzug",
    "article": "der"
  },
  {
    "noun": "Kontakter",
    "article": "der"
  },
  {
    "noun": "Kontakthebel",
    "article": "der"
  },
  {
    "noun": "Kontakthof",
    "article": "der"
  },
  {
    "noun": "Kontaktknopf",
    "article": "der"
  },
  {
    "noun": "Kontaktmann",
    "article": "der"
  },
  {
    "noun": "Kontaktstift",
    "article": "der"
  },
  {
    "noun": "Kontaktwiderstand",
    "article": "der"
  },
  {
    "noun": "Kontenplan",
    "article": "der"
  },
  {
    "noun": "Kontenrahmen",
    "article": "der"
  },
  {
    "noun": "Konter",
    "article": "der"
  },
  {
    "noun": "Konteradmiral",
    "article": "der"
  },
  {
    "noun": "Konterschlag",
    "article": "der"
  },
  {
    "noun": "Kontext",
    "article": "der"
  },
  {
    "noun": "Kontinent",
    "article": "der"
  },
  {
    "noun": "Kontinentalschelf",
    "article": "der"
  },
  {
    "noun": "Kontinentalsockel",
    "article": "der"
  },
  {
    "noun": "Kontoabschluss",
    "article": "der"
  },
  {
    "noun": "Kontoausgleich",
    "article": "der"
  },
  {
    "noun": "Kontoauszug",
    "article": "der"
  },
  {
    "noun": "Kontoinhaber",
    "article": "der"
  },
  {
    "noun": "Kontokorrentkredit",
    "article": "der"
  },
  {
    "noun": "Kontokorrentvertrag",
    "article": "der"
  },
  {
    "noun": "Kontostand",
    "article": "der"
  },
  {
    "noun": "Kontrabass",
    "article": "der"
  },
  {
    "noun": "Kontrahent",
    "article": "der"
  },
  {
    "noun": "Kontrahierungszwang",
    "article": "der"
  },
  {
    "noun": "Kontrapunkt",
    "article": "der"
  },
  {
    "noun": "Kontrast",
    "article": "der"
  },
  {
    "noun": "Kontrollabschnitt",
    "article": "der"
  },
  {
    "noun": "Kontroller",
    "article": "der"
  },
  {
    "noun": "Kontrolleur",
    "article": "der"
  },
  {
    "noun": "Kontrollgang",
    "article": "der"
  },
  {
    "noun": "Kontrollor",
    "article": "der"
  },
  {
    "noun": "Kontrollpunkt",
    "article": "der"
  },
  {
    "noun": "Kontrollturm",
    "article": "der"
  },
  {
    "noun": "Konus",
    "article": "der"
  },
  {
    "noun": "Konvektor",
    "article": "der"
  },
  {
    "noun": "Konverter",
    "article": "der"
  },
  {
    "noun": "Konvertit",
    "article": "der"
  },
  {
    "noun": "Konvexspiegel",
    "article": "der"
  },
  {
    "noun": "Konvoi",
    "article": "der"
  },
  {
    "noun": "Konzern",
    "article": "der"
  },
  {
    "noun": "Konzernchef",
    "article": "der"
  },
  {
    "noun": "Konzertbesucher",
    "article": "der"
  },
  {
    "noun": "Konzertmeister",
    "article": "der"
  },
  {
    "noun": "Konzertsaal",
    "article": "der"
  },
  {
    "noun": "Konzessionsinhaber",
    "article": "der"
  },
  {
    "noun": "Kooperationspartner",
    "article": "der"
  },
  {
    "noun": "Kooperationsvertrag",
    "article": "der"
  },
  {
    "noun": "Kooperator",
    "article": "der"
  },
  {
    "noun": "Koordinator",
    "article": "der"
  },
  {
    "noun": "Kopal",
    "article": "der"
  },
  {
    "noun": "Kopf",
    "article": "der"
  },
  {
    "noun": "Kopfbahnhof",
    "article": "der"
  },
  {
    "noun": "Kopfball",
    "article": "der"
  },
  {
    "noun": "Kopfkissenbezug",
    "article": "der"
  },
  {
    "noun": "Kopfputz",
    "article": "der"
  },
  {
    "noun": "Kopfsalat",
    "article": "der"
  },
  {
    "noun": "Kopfschmerz",
    "article": "der"
  },
  {
    "noun": "Kopfschmuck",
    "article": "der"
  },
  {
    "noun": "Kopfschuss",
    "article": "der"
  },
  {
    "noun": "Kopfschutz",
    "article": "der"
  },
  {
    "noun": "Kopfsprung",
    "article": "der"
  },
  {
    "noun": "Kopfstand",
    "article": "der"
  },
  {
    "noun": "Kopfstein",
    "article": "der"
  },
  {
    "noun": "Kopierer",
    "article": "der"
  },
  {
    "noun": "Kopierrahmen",
    "article": "der"
  },
  {
    "noun": "Kopierschutz",
    "article": "der"
  },
  {
    "noun": "Kopierstift",
    "article": "der"
  },
  {
    "noun": "Kopilot",
    "article": "der"
  },
  {
    "noun": "Kopist",
    "article": "der"
  },
  {
    "noun": "Koprolith",
    "article": "der"
  },
  {
    "noun": "Koprozessor",
    "article": "der"
  },
  {
    "noun": "Kops",
    "article": "der"
  },
  {
    "noun": "Korallenfischer",
    "article": "der"
  },
  {
    "noun": "Koran",
    "article": "der"
  },
  {
    "noun": "Korb",
    "article": "der"
  },
  {
    "noun": "Korbball",
    "article": "der"
  },
  {
    "noun": "Korbflechter",
    "article": "der"
  },
  {
    "noun": "Korbleger",
    "article": "der"
  },
  {
    "noun": "Korbmacher",
    "article": "der"
  },
  {
    "noun": "Korbsessel",
    "article": "der"
  },
  {
    "noun": "Korbstuhl",
    "article": "der"
  },
  {
    "noun": "Kordon",
    "article": "der"
  },
  {
    "noun": "Koreaner",
    "article": "der"
  },
  {
    "noun": "Koriander",
    "article": "der"
  },
  {
    "noun": "Korinthenkacker",
    "article": "der"
  },
  {
    "noun": "Korinther",
    "article": "der"
  },
  {
    "noun": "Kork",
    "article": "der"
  },
  {
    "noun": "Korkenzieher",
    "article": "der"
  },
  {
    "noun": "Kormoran",
    "article": "der"
  },
  {
    "noun": "Kornettist",
    "article": "der"
  },
  {
    "noun": "Kornspeicher",
    "article": "der"
  },
  {
    "noun": "Korporal",
    "article": "der"
  },
  {
    "noun": "Korral",
    "article": "der"
  },
  {
    "noun": "Korrektor",
    "article": "der"
  },
  {
    "noun": "Korrekturabzug",
    "article": "der"
  },
  {
    "noun": "Korrekturbogen",
    "article": "der"
  },
  {
    "noun": "Korrespondent",
    "article": "der"
  },
  {
    "noun": "Korrespondenzanwalt",
    "article": "der"
  },
  {
    "noun": "Korridor",
    "article": "der"
  },
  {
    "noun": "Korrosionsschutz",
    "article": "der"
  },
  {
    "noun": "Korruptionsskandal",
    "article": "der"
  },
  {
    "noun": "Korsar",
    "article": "der"
  },
  {
    "noun": "Korse",
    "article": "der"
  },
  {
    "noun": "Kortex",
    "article": "der"
  },
  {
    "noun": "Korund",
    "article": "der"
  },
  {
    "noun": "Kosekans",
    "article": "der"
  },
  {
    "noun": "Kosename",
    "article": "der"
  },
  {
    "noun": "Kosinus",
    "article": "der"
  },
  {
    "noun": "Kosmetiker",
    "article": "der"
  },
  {
    "noun": "Kosmetikkoffer",
    "article": "der"
  },
  {
    "noun": "Kosmetiksalon",
    "article": "der"
  },
  {
    "noun": "Kosmonaut",
    "article": "der"
  },
  {
    "noun": "Kosmopolit",
    "article": "der"
  },
  {
    "noun": "Kosmopolitismus",
    "article": "der"
  },
  {
    "noun": "Kosmos",
    "article": "der"
  },
  {
    "noun": "Kosovare",
    "article": "der"
  },
  {
    "noun": "Kostenanschlag",
    "article": "der"
  },
  {
    "noun": "Kostenanstieg",
    "article": "der"
  },
  {
    "noun": "Kostenanteil",
    "article": "der"
  },
  {
    "noun": "Kostenaufwand",
    "article": "der"
  },
  {
    "noun": "Kostendruck",
    "article": "der"
  },
  {
    "noun": "Kostenfaktor",
    "article": "der"
  },
  {
    "noun": "Kostenpunkt",
    "article": "der"
  },
  {
    "noun": "Kostenrahmen",
    "article": "der"
  },
  {
    "noun": "Kostenvoranschlag",
    "article": "der"
  },
  {
    "noun": "Kot",
    "article": "der"
  },
  {
    "noun": "Kotangens",
    "article": "der"
  },
  {
    "noun": "Kotzbrocken",
    "article": "der"
  },
  {
    "noun": "Krabbenkutter",
    "article": "der"
  },
  {
    "noun": "Krabbentaucher",
    "article": "der"
  },
  {
    "noun": "Krach",
    "article": "der"
  },
  {
    "noun": "Kracher",
    "article": "der"
  },
  {
    "noun": "Kradmelder",
    "article": "der"
  },
  {
    "noun": "Kraftantrieb",
    "article": "der"
  },
  {
    "noun": "Kraftarm",
    "article": "der"
  },
  {
    "noun": "Kraftaufwand",
    "article": "der"
  },
  {
    "noun": "Kraftausdruck",
    "article": "der"
  },
  {
    "noun": "Kraftfahrer",
    "article": "der"
  },
  {
    "noun": "Kraftfahrzeugbau",
    "article": "der"
  },
  {
    "noun": "Kraftfahrzeugbrief",
    "article": "der"
  },
  {
    "noun": "Kraftfahrzeugschein",
    "article": "der"
  },
  {
    "noun": "Kraftmeier",
    "article": "der"
  },
  {
    "noun": "Kraftmesser",
    "article": "der"
  },
  {
    "noun": "Kraftprotz",
    "article": "der"
  },
  {
    "noun": "Kraftspeicher",
    "article": "der"
  },
  {
    "noun": "Kraftstoff",
    "article": "der"
  },
  {
    "noun": "Kraftstoffmesser",
    "article": "der"
  },
  {
    "noun": "Kraftstoffverbrauch",
    "article": "der"
  },
  {
    "noun": "Kraftstrom",
    "article": "der"
  },
  {
    "noun": "Kraftverkehr",
    "article": "der"
  },
  {
    "noun": "Kraftverlust",
    "article": "der"
  },
  {
    "noun": "Kraftwagen",
    "article": "der"
  },
  {
    "noun": "Kragen",
    "article": "der"
  },
  {
    "noun": "Kragenknopf",
    "article": "der"
  },
  {
    "noun": "Kragenspiegel",
    "article": "der"
  },
  {
    "noun": "Krake",
    "article": "der"
  },
  {
    "noun": "Krakeeler",
    "article": "der"
  },
  {
    "noun": "Kram",
    "article": "der"
  },
  {
    "noun": "Krampf",
    "article": "der"
  },
  {
    "noun": "Krampus",
    "article": "der"
  },
  {
    "noun": "Kran",
    "article": "der"
  },
  {
    "noun": "Kranfahrer",
    "article": "der"
  },
  {
    "noun": "Kranich",
    "article": "der"
  },
  {
    "noun": "Kranwagen",
    "article": "der"
  },
  {
    "noun": "Kranz",
    "article": "der"
  },
  {
    "noun": "Kranzbinder",
    "article": "der"
  },
  {
    "noun": "Krapfen",
    "article": "der"
  },
  {
    "noun": "Krater",
    "article": "der"
  },
  {
    "noun": "Kratersee",
    "article": "der"
  },
  {
    "noun": "Kratzer",
    "article": "der"
  },
  {
    "noun": "Krauskopf",
    "article": "der"
  },
  {
    "noun": "Krautjunker",
    "article": "der"
  },
  {
    "noun": "Krautsalat",
    "article": "der"
  },
  {
    "noun": "Krautwickel",
    "article": "der"
  },
  {
    "noun": "Krawall",
    "article": "der"
  },
  {
    "noun": "Krawallmacher",
    "article": "der"
  },
  {
    "noun": "Krawattenmuffel",
    "article": "der"
  },
  {
    "noun": "Kreationismus",
    "article": "der"
  },
  {
    "noun": "Krebs",
    "article": "der"
  },
  {
    "noun": "Krebsspezialist",
    "article": "der"
  },
  {
    "noun": "Kredit",
    "article": "der"
  },
  {
    "noun": "Kreditantrag",
    "article": "der"
  },
  {
    "noun": "Kreditauftrag",
    "article": "der"
  },
  {
    "noun": "Kreditbetrag",
    "article": "der"
  },
  {
    "noun": "Kreditbrief",
    "article": "der"
  },
  {
    "noun": "Kreditgeber",
    "article": "der"
  },
  {
    "noun": "Kredithai",
    "article": "der"
  },
  {
    "noun": "Kreditkauf",
    "article": "der"
  },
  {
    "noun": "Kreditmarkt",
    "article": "der"
  },
  {
    "noun": "Kreditnehmer",
    "article": "der"
  },
  {
    "noun": "Kreditor",
    "article": "der"
  },
  {
    "noun": "Kreditzins",
    "article": "der"
  },
  {
    "noun": "Kreidefels",
    "article": "der"
  },
  {
    "noun": "Kreis",
    "article": "der"
  },
  {
    "noun": "Kreisabschnitt",
    "article": "der"
  },
  {
    "noun": "Kreisausschnitt",
    "article": "der"
  },
  {
    "noun": "Kreisbogen",
    "article": "der"
  },
  {
    "noun": "Kreisdurchmesser",
    "article": "der"
  },
  {
    "noun": "Kreisel",
    "article": "der"
  },
  {
    "noun": "Kreiselkompass",
    "article": "der"
  },
  {
    "noun": "Kreiskegel",
    "article": "der"
  },
  {
    "noun": "Kreiskolbenmotor",
    "article": "der"
  },
  {
    "noun": "Kreislauf",
    "article": "der"
  },
  {
    "noun": "Kreislaufstillstand",
    "article": "der"
  },
  {
    "noun": "Kreisprozess",
    "article": "der"
  },
  {
    "noun": "Kreisring",
    "article": "der"
  },
  {
    "noun": "Kreissektor",
    "article": "der"
  },
  {
    "noun": "Kreisstrom",
    "article": "der"
  },
  {
    "noun": "Kreisumfang",
    "article": "der"
  },
  {
    "noun": "Kreisverkehr",
    "article": "der"
  },
  {
    "noun": "Kreml",
    "article": "der"
  },
  {
    "noun": "Krempel",
    "article": "der"
  },
  {
    "noun": "Kren",
    "article": "der"
  },
  {
    "noun": "Kreole",
    "article": "der"
  },
  {
    "noun": "Kreter",
    "article": "der"
  },
  {
    "noun": "Kretikus",
    "article": "der"
  },
  {
    "noun": "Kreuzbiss",
    "article": "der"
  },
  {
    "noun": "Kreuzdorn",
    "article": "der"
  },
  {
    "noun": "Kreuzer",
    "article": "der"
  },
  {
    "noun": "Kreuzfahrer",
    "article": "der"
  },
  {
    "noun": "Kreuzgang",
    "article": "der"
  },
  {
    "noun": "Kreuzgriff",
    "article": "der"
  },
  {
    "noun": "Kreuzknoten",
    "article": "der"
  },
  {
    "noun": "Kreuzreim",
    "article": "der"
  },
  {
    "noun": "Kreuzritter",
    "article": "der"
  },
  {
    "noun": "Kreuzungspunkt",
    "article": "der"
  },
  {
    "noun": "Kreuzweg",
    "article": "der"
  },
  {
    "noun": "Kreuzzug",
    "article": "der"
  },
  {
    "noun": "Kricketball",
    "article": "der"
  },
  {
    "noun": "Kricketspieler",
    "article": "der"
  },
  {
    "noun": "Kridar",
    "article": "der"
  },
  {
    "noun": "Kriecher",
    "article": "der"
  },
  {
    "noun": "Kriechstrom",
    "article": "der"
  },
  {
    "noun": "Krieg",
    "article": "der"
  },
  {
    "noun": "Krieger",
    "article": "der"
  },
  {
    "noun": "Kriegsausbruch",
    "article": "der"
  },
  {
    "noun": "Kriegsberichterstatter",
    "article": "der"
  },
  {
    "noun": "Kriegsdienst",
    "article": "der"
  },
  {
    "noun": "Kriegsdienstverweigerer",
    "article": "der"
  },
  {
    "noun": "Kriegsgewinnler",
    "article": "der"
  },
  {
    "noun": "Kriegsherr",
    "article": "der"
  },
  {
    "noun": "Kriegshetzer",
    "article": "der"
  },
  {
    "noun": "Kriegsminister",
    "article": "der"
  },
  {
    "noun": "Kriegsrat",
    "article": "der"
  },
  {
    "noun": "Kriegsschaden",
    "article": "der"
  },
  {
    "noun": "Kriegsschauplatz",
    "article": "der"
  },
  {
    "noun": "Kriegstanz",
    "article": "der"
  },
  {
    "noun": "Kriegstreiber",
    "article": "der"
  },
  {
    "noun": "Kriegsverbrecher",
    "article": "der"
  },
  {
    "noun": "Kriegsverbrecherprozess",
    "article": "der"
  },
  {
    "noun": "Kriegszustand",
    "article": "der"
  },
  {
    "noun": "Krimi",
    "article": "der"
  },
  {
    "noun": "Kriminalfall",
    "article": "der"
  },
  {
    "noun": "Kriminalfilm",
    "article": "der"
  },
  {
    "noun": "Kriminalist",
    "article": "der"
  },
  {
    "noun": "Kriminalkommissar",
    "article": "der"
  },
  {
    "noun": "Kriminalpolizist",
    "article": "der"
  },
  {
    "noun": "Kriminalroman",
    "article": "der"
  },
  {
    "noun": "Kriminologe",
    "article": "der"
  },
  {
    "noun": "Krimskrams",
    "article": "der"
  },
  {
    "noun": "Kringel",
    "article": "der"
  },
  {
    "noun": "Krisenherd",
    "article": "der"
  },
  {
    "noun": "Krisenstab",
    "article": "der"
  },
  {
    "noun": "Kristallit",
    "article": "der"
  },
  {
    "noun": "Kritiker",
    "article": "der"
  },
  {
    "noun": "Kritikpunkt",
    "article": "der"
  },
  {
    "noun": "Krittler",
    "article": "der"
  },
  {
    "noun": "Kroate",
    "article": "der"
  },
  {
    "noun": "Krokant",
    "article": "der"
  },
  {
    "noun": "Krokus",
    "article": "der"
  },
  {
    "noun": "Kronenkorken",
    "article": "der"
  },
  {
    "noun": "Kronenkranich",
    "article": "der"
  },
  {
    "noun": "Kronenverschluss",
    "article": "der"
  },
  {
    "noun": "Kronkorken",
    "article": "der"
  },
  {
    "noun": "Kronleuchter",
    "article": "der"
  },
  {
    "noun": "Kronprinz",
    "article": "der"
  },
  {
    "noun": "Kronzeuge",
    "article": "der"
  },
  {
    "noun": "Kropf",
    "article": "der"
  },
  {
    "noun": "Krug",
    "article": "der"
  },
  {
    "noun": "Krummstab",
    "article": "der"
  },
  {
    "noun": "Kryolith",
    "article": "der"
  },
  {
    "noun": "Kryostat",
    "article": "der"
  },
  {
    "noun": "Kryptograf",
    "article": "der"
  },
  {
    "noun": "Kryptorchismus",
    "article": "der"
  },
  {
    "noun": "Kubaner",
    "article": "der"
  },
  {
    "noun": "Kubik",
    "article": "der"
  },
  {
    "noun": "Kubismus",
    "article": "der"
  },
  {
    "noun": "Kubist",
    "article": "der"
  },
  {
    "noun": "Kubus",
    "article": "der"
  },
  {
    "noun": "Kuchen",
    "article": "der"
  },
  {
    "noun": "Kuchenboden",
    "article": "der"
  },
  {
    "noun": "Kuckuck",
    "article": "der"
  },
  {
    "noun": "Kuder",
    "article": "der"
  },
  {
    "noun": "Kugelabschnitt",
    "article": "der"
  },
  {
    "noun": "Kugelblitz",
    "article": "der"
  },
  {
    "noun": "Kugelfisch",
    "article": "der"
  },
  {
    "noun": "Kugelhagel",
    "article": "der"
  },
  {
    "noun": "Kugelkopf",
    "article": "der"
  },
  {
    "noun": "Kugelschreiber",
    "article": "der"
  },
  {
    "noun": "Kuguar",
    "article": "der"
  },
  {
    "noun": "Kuhfladen",
    "article": "der"
  },
  {
    "noun": "Kuhhandel",
    "article": "der"
  },
  {
    "noun": "Kuhhirt",
    "article": "der"
  },
  {
    "noun": "Kuhmist",
    "article": "der"
  },
  {
    "noun": "Kuhstall",
    "article": "der"
  },
  {
    "noun": "Kuli",
    "article": "der"
  },
  {
    "noun": "Kulissenschieber",
    "article": "der"
  },
  {
    "noun": "Kulissenwechsel",
    "article": "der"
  },
  {
    "noun": "Kulminationspunkt",
    "article": "der"
  },
  {
    "noun": "Kult",
    "article": "der"
  },
  {
    "noun": "Kultfilm",
    "article": "der"
  },
  {
    "noun": "Kultivator",
    "article": "der"
  },
  {
    "noun": "Kulturanthropologe",
    "article": "der"
  },
  {
    "noun": "Kulturaustausch",
    "article": "der"
  },
  {
    "noun": "Kulturbanause",
    "article": "der"
  },
  {
    "noun": "Kulturbeutel",
    "article": "der"
  },
  {
    "noun": "Kulturkreis",
    "article": "der"
  },
  {
    "noun": "Kulturpalast",
    "article": "der"
  },
  {
    "noun": "Kulturraum",
    "article": "der"
  },
  {
    "noun": "Kulturschock",
    "article": "der"
  },
  {
    "noun": "Kulturverfall",
    "article": "der"
  },
  {
    "noun": "Kultusminister",
    "article": "der"
  },
  {
    "noun": "Kummer",
    "article": "der"
  },
  {
    "noun": "Kummerbund",
    "article": "der"
  },
  {
    "noun": "Kummerkasten",
    "article": "der"
  },
  {
    "noun": "Kumpan",
    "article": "der"
  },
  {
    "noun": "Kumpel",
    "article": "der"
  },
  {
    "noun": "Kundenberater",
    "article": "der"
  },
  {
    "noun": "Kundenbesuch",
    "article": "der"
  },
  {
    "noun": "Kundendienst",
    "article": "der"
  },
  {
    "noun": "Kundenfang",
    "article": "der"
  },
  {
    "noun": "Kundenkredit",
    "article": "der"
  },
  {
    "noun": "Kundenkreis",
    "article": "der"
  },
  {
    "noun": "Kundenstamm",
    "article": "der"
  },
  {
    "noun": "Kundenwerber",
    "article": "der"
  },
  {
    "noun": "Kundenwunsch",
    "article": "der"
  },
  {
    "noun": "Kundschafter",
    "article": "der"
  },
  {
    "noun": "Kunstbanause",
    "article": "der"
  },
  {
    "noun": "Kunstdarm",
    "article": "der"
  },
  {
    "noun": "Kunstdruck",
    "article": "der"
  },
  {
    "noun": "Kunsterzieher",
    "article": "der"
  },
  {
    "noun": "Kunstfehler",
    "article": "der"
  },
  {
    "noun": "Kunstflieger",
    "article": "der"
  },
  {
    "noun": "Kunstflug",
    "article": "der"
  },
  {
    "noun": "Kunstfreund",
    "article": "der"
  },
  {
    "noun": "Kunstgegenstand",
    "article": "der"
  },
  {
    "noun": "Kunstgewerbler",
    "article": "der"
  },
  {
    "noun": "Kunstgriff",
    "article": "der"
  },
  {
    "noun": "Kunsthandel",
    "article": "der"
  },
  {
    "noun": "Kunsthandwerker",
    "article": "der"
  },
  {
    "noun": "Kunsthistoriker",
    "article": "der"
  },
  {
    "noun": "Kunstkritiker",
    "article": "der"
  },
  {
    "noun": "Kunstliebhaber",
    "article": "der"
  },
  {
    "noun": "Kunstmaler",
    "article": "der"
  },
  {
    "noun": "Kunstmarkt",
    "article": "der"
  },
  {
    "noun": "Kunstrasen",
    "article": "der"
  },
  {
    "noun": "Kunstreiter",
    "article": "der"
  },
  {
    "noun": "Kunstsammler",
    "article": "der"
  },
  {
    "noun": "Kunstschatz",
    "article": "der"
  },
  {
    "noun": "Kunstschmied",
    "article": "der"
  },
  {
    "noun": "Kunstschnee",
    "article": "der"
  },
  {
    "noun": "Kunststein",
    "article": "der"
  },
  {
    "noun": "Kunststoff",
    "article": "der"
  },
  {
    "noun": "Kunstunterricht",
    "article": "der"
  },
  {
    "noun": "Kupferdraht",
    "article": "der"
  },
  {
    "noun": "Kupferdruck",
    "article": "der"
  },
  {
    "noun": "Kupferglanz",
    "article": "der"
  },
  {
    "noun": "Kupferkessel",
    "article": "der"
  },
  {
    "noun": "Kupferkies",
    "article": "der"
  },
  {
    "noun": "Kupferschiefer",
    "article": "der"
  },
  {
    "noun": "Kupferstecher",
    "article": "der"
  },
  {
    "noun": "Kupferstich",
    "article": "der"
  },
  {
    "noun": "Kupolofen",
    "article": "der"
  },
  {
    "noun": "Kuppler",
    "article": "der"
  },
  {
    "noun": "Kupplungsbelag",
    "article": "der"
  },
  {
    "noun": "Kurator",
    "article": "der"
  },
  {
    "noun": "Kuraufenthalt",
    "article": "der"
  },
  {
    "noun": "Kurbelinduktor",
    "article": "der"
  },
  {
    "noun": "Kurbeltrieb",
    "article": "der"
  },
  {
    "noun": "Kurde",
    "article": "der"
  },
  {
    "noun": "Kurgast",
    "article": "der"
  },
  {
    "noun": "Kurier",
    "article": "der"
  },
  {
    "noun": "Kurierdienst",
    "article": "der"
  },
  {
    "noun": "Kurort",
    "article": "der"
  },
  {
    "noun": "Kurpark",
    "article": "der"
  },
  {
    "noun": "Kurpfuscher",
    "article": "der"
  },
  {
    "noun": "Kurs",
    "article": "der"
  },
  {
    "noun": "Kursabschlag",
    "article": "der"
  },
  {
    "noun": "Kursanstieg",
    "article": "der"
  },
  {
    "noun": "Kursbericht",
    "article": "der"
  },
  {
    "noun": "Kurseinbruch",
    "article": "der"
  },
  {
    "noun": "Kursgewinn",
    "article": "der"
  },
  {
    "noun": "Kursleiter",
    "article": "der"
  },
  {
    "noun": "Kursmakler",
    "article": "der"
  },
  {
    "noun": "Kurssturz",
    "article": "der"
  },
  {
    "noun": "Kursteilnehmer",
    "article": "der"
  },
  {
    "noun": "Kursunterricht",
    "article": "der"
  },
  {
    "noun": "Kursus",
    "article": "der"
  },
  {
    "noun": "Kursverfall",
    "article": "der"
  },
  {
    "noun": "Kursverlust",
    "article": "der"
  },
  {
    "noun": "Kurswagen",
    "article": "der"
  },
  {
    "noun": "Kurswechsel",
    "article": "der"
  },
  {
    "noun": "Kurswert",
    "article": "der"
  },
  {
    "noun": "Kurszettel",
    "article": "der"
  },
  {
    "noun": "Kurvenschreiber",
    "article": "der"
  },
  {
    "noun": "Kurzbericht",
    "article": "der"
  },
  {
    "noun": "Kurzfilm",
    "article": "der"
  },
  {
    "noun": "Kurzschluss",
    "article": "der"
  },
  {
    "noun": "Kurzstreckenflug",
    "article": "der"
  },
  {
    "noun": "Kurzstreckenlauf",
    "article": "der"
  },
  {
    "noun": "Kurzstreckenverkehr",
    "article": "der"
  },
  {
    "noun": "Kurzurlaub",
    "article": "der"
  },
  {
    "noun": "Kurzwellensender",
    "article": "der"
  },
  {
    "noun": "Kurzzeitwecker",
    "article": "der"
  },
  {
    "noun": "Kuss",
    "article": "der"
  },
  {
    "noun": "Kustos",
    "article": "der"
  },
  {
    "noun": "Kutscher",
    "article": "der"
  },
  {
    "noun": "Kuttelfleck",
    "article": "der"
  },
  {
    "noun": "Kutter",
    "article": "der"
  },
  {
    "noun": "Kuwaiter",
    "article": "der"
  },
  {
    "noun": "Kux",
    "article": "der"
  },
  {
    "noun": "Kyanit",
    "article": "der"
  },
  {
    "noun": "Kybernetiker",
    "article": "der"
  },
  {
    "noun": "Abendanzug",
    "article": "der"
  },
  {
    "noun": "Abendgottesdienst",
    "article": "der"
  },
  {
    "noun": "Abendhimmel",
    "article": "der"
  },
  {
    "noun": "Abendkurs",
    "article": "der"
  },
  {
    "noun": "Abendmahlsgottesdienst",
    "article": "der"
  },
  {
    "noun": "Abendmahlskelch",
    "article": "der"
  },
  {
    "noun": "Abendmahlswein",
    "article": "der"
  },
  {
    "noun": "Abendschein",
    "article": "der"
  },
  {
    "noun": "Abendsegler",
    "article": "der"
  },
  {
    "noun": "Abendstern",
    "article": "der"
  },
  {
    "noun": "Abendunterricht",
    "article": "der"
  },
  {
    "noun": "Abendverkauf",
    "article": "der"
  },
  {
    "noun": "Abendwind",
    "article": "der"
  },
  {
    "noun": "Abendzug",
    "article": "der"
  },
  {
    "noun": "Abenteuerdrang",
    "article": "der"
  },
  {
    "noun": "Abenteuerfilm",
    "article": "der"
  },
  {
    "noun": "Abenteuerroman",
    "article": "der"
  },
  {
    "noun": "Abenteuerspielplatz",
    "article": "der"
  },
  {
    "noun": "Abenteuertourismus",
    "article": "der"
  },
  {
    "noun": "Abenteuerurlaub",
    "article": "der"
  },
  {
    "noun": "Abenteurer",
    "article": "der"
  },
  {
    "noun": "Abenteurerfilm",
    "article": "der"
  },
  {
    "noun": "Abenteurergeist",
    "article": "der"
  },
  {
    "noun": "Aberglaube",
    "article": "der"
  },
  {
    "noun": "Aberwitz",
    "article": "der"
  },
  {
    "noun": "Abfahrtslauf",
    "article": "der"
  },
  {
    "noun": "Abfahrtsort",
    "article": "der"
  },
  {
    "noun": "Abfahrtsplan",
    "article": "der"
  },
  {
    "noun": "Abfall",
    "article": "der"
  },
  {
    "noun": "Abfallcontainer",
    "article": "der"
  },
  {
    "noun": "Abfalleimer",
    "article": "der"
  },
  {
    "noun": "Abfallhaufen",
    "article": "der"
  },
  {
    "noun": "Abfallkalender",
    "article": "der"
  },
  {
    "noun": "Abfallkorb",
    "article": "der"
  },
  {
    "noun": "Abfallsack",
    "article": "der"
  },
  {
    "noun": "Abfallstoff",
    "article": "der"
  },
  {
    "noun": "Abfertigungsschalter",
    "article": "der"
  },
  {
    "noun": "Abfindungsvertrag",
    "article": "der"
  },
  {
    "noun": "Abflug",
    "article": "der"
  },
  {
    "noun": "Abflughafen",
    "article": "der"
  },
  {
    "noun": "Abflugort",
    "article": "der"
  },
  {
    "noun": "Abflugschalter",
    "article": "der"
  },
  {
    "noun": "Abfluss",
    "article": "der"
  },
  {
    "noun": "Abflussgraben",
    "article": "der"
  },
  {
    "noun": "Abflusshahn",
    "article": "der"
  },
  {
    "noun": "Abflusskanal",
    "article": "der"
  },
  {
    "noun": "Abflussreiniger",
    "article": "der"
  },
  {
    "noun": "Abgabepreis",
    "article": "der"
  },
  {
    "noun": "Abgabesatz",
    "article": "der"
  },
  {
    "noun": "Abgabetermin",
    "article": "der"
  },
  {
    "noun": "Abgang",
    "article": "der"
  },
  {
    "noun": "Abgaskatalysator",
    "article": "der"
  },
  {
    "noun": "Abgeordnetensitz",
    "article": "der"
  },
  {
    "noun": "Abgesang",
    "article": "der"
  },
  {
    "noun": "Abglanz",
    "article": "der"
  },
  {
    "noun": "Abgleich",
    "article": "der"
  },
  {
    "noun": "Abgott",
    "article": "der"
  },
  {
    "noun": "Abgriff",
    "article": "der"
  },
  {
    "noun": "Abgrund",
    "article": "der"
  },
  {
    "noun": "Abguss",
    "article": "der"
  },
  {
    "noun": "Abhang",
    "article": "der"
  },
  {
    "noun": "Abhitzekessel",
    "article": "der"
  },
  {
    "noun": "Abholdienst",
    "article": "der"
  },
  {
    "noun": "Abholer",
    "article": "der"
  },
  {
    "noun": "Abholservice",
    "article": "der"
  },
  {
    "noun": "Abholtag",
    "article": "der"
  },
  {
    "noun": "Abholtermin",
    "article": "der"
  },
  {
    "noun": "Abhub",
    "article": "der"
  },
  {
    "noun": "Abiturient",
    "article": "der"
  },
  {
    "noun": "Abklatsch",
    "article": "der"
  },
  {
    "noun": "Abladeplatz",
    "article": "der"
  },
  {
    "noun": "Ablader",
    "article": "der"
  },
  {
    "noun": "Ablagekorb",
    "article": "der"
  },
  {
    "noun": "Ablass",
    "article": "der"
  },
  {
    "noun": "Ablassbrief",
    "article": "der"
  },
  {
    "noun": "Ablasshahn",
    "article": "der"
  },
  {
    "noun": "Ablasshandel",
    "article": "der"
  },
  {
    "noun": "Ablativ",
    "article": "der"
  },
  {
    "noun": "Ablauf",
    "article": "der"
  },
  {
    "noun": "Ablaufplan",
    "article": "der"
  },
  {
    "noun": "Ablaut",
    "article": "der"
  },
  {
    "noun": "Ableger",
    "article": "der"
  },
  {
    "noun": "Ablehnungsbescheid",
    "article": "der"
  },
  {
    "noun": "Ableser",
    "article": "der"
  },
  {
    "noun": "Abluftkamin",
    "article": "der"
  },
  {
    "noun": "Abmarsch",
    "article": "der"
  },
  {
    "noun": "Abnehmer",
    "article": "der"
  },
  {
    "noun": "Abolitionist",
    "article": "der"
  },
  {
    "noun": "Abonnementspreis",
    "article": "der"
  },
  {
    "noun": "Abonnent",
    "article": "der"
  },
  {
    "noun": "Abort",
    "article": "der"
  },
  {
    "noun": "Abpackbetrieb",
    "article": "der"
  },
  {
    "noun": "Abprall",
    "article": "der"
  },
  {
    "noun": "Abpraller",
    "article": "der"
  },
  {
    "noun": "Abraum",
    "article": "der"
  },
  {
    "noun": "Abraumbagger",
    "article": "der"
  },
  {
    "noun": "Abrechnungsbeleg",
    "article": "der"
  },
  {
    "noun": "Abrechnungsbetrug",
    "article": "der"
  },
  {
    "noun": "Abrechnungstermin",
    "article": "der"
  },
  {
    "noun": "Abrechnungsverkehr",
    "article": "der"
  },
  {
    "noun": "Abreisetag",
    "article": "der"
  },
  {
    "noun": "Abreiteplatz",
    "article": "der"
  },
  {
    "noun": "Abrieb",
    "article": "der"
  },
  {
    "noun": "Abriss",
    "article": "der"
  },
  {
    "noun": "Abruf",
    "article": "der"
  },
  {
    "noun": "Abrutsch",
    "article": "der"
  },
  {
    "noun": "Absagebrief",
    "article": "der"
  },
  {
    "noun": "Absatz",
    "article": "der"
  },
  {
    "noun": "Absatzeinbruch",
    "article": "der"
  },
  {
    "noun": "Absatzkanal",
    "article": "der"
  },
  {
    "noun": "Absatzmarkt",
    "article": "der"
  },
  {
    "noun": "Absatzweg",
    "article": "der"
  },
  {
    "noun": "Abschaum",
    "article": "der"
  },
  {
    "noun": "Abscheider",
    "article": "der"
  },
  {
    "noun": "Abschied",
    "article": "der"
  },
  {
    "noun": "Abschiedsbesuch",
    "article": "der"
  },
  {
    "noun": "Abschiedsbrief",
    "article": "der"
  },
  {
    "noun": "Abschiedskuss",
    "article": "der"
  },
  {
    "noun": "Abschiedsschmerz",
    "article": "der"
  },
  {
    "noun": "Abschiedstrunk",
    "article": "der"
  },
  {
    "noun": "Abschirmdienst",
    "article": "der"
  },
  {
    "noun": "Abschlag",
    "article": "der"
  },
  {
    "noun": "Abschleppdienst",
    "article": "der"
  },
  {
    "noun": "Abschlepphaken",
    "article": "der"
  },
  {
    "noun": "Abschleppkran",
    "article": "der"
  },
  {
    "noun": "Abschleppwagen",
    "article": "der"
  },
  {
    "noun": "Abschliff",
    "article": "der"
  },
  {
    "noun": "Abschluss",
    "article": "der"
  },
  {
    "noun": "Abschlussball",
    "article": "der"
  },
  {
    "noun": "Abschlussbericht",
    "article": "der"
  },
  {
    "noun": "Abschneider",
    "article": "der"
  },
  {
    "noun": "Abschnitt",
    "article": "der"
  },
  {
    "noun": "Abschreiber",
    "article": "der"
  },
  {
    "noun": "Abschreibungsbetrag",
    "article": "der"
  },
  {
    "noun": "Abschreibungssatz",
    "article": "der"
  },
  {
    "noun": "Abschreibungszeitraum",
    "article": "der"
  },
  {
    "noun": "Abschuss",
    "article": "der"
  },
  {
    "noun": "Abschwung",
    "article": "der"
  },
  {
    "noun": "Absender",
    "article": "der"
  },
  {
    "noun": "Absenker",
    "article": "der"
  },
  {
    "noun": "Absentismus",
    "article": "der"
  },
  {
    "noun": "Absetzer",
    "article": "der"
  },
  {
    "noun": "Absinth",
    "article": "der"
  },
  {
    "noun": "Absolutheitsanspruch",
    "article": "der"
  },
  {
    "noun": "Absolutismus",
    "article": "der"
  },
  {
    "noun": "Absolutist",
    "article": "der"
  },
  {
    "noun": "Absolutwert",
    "article": "der"
  },
  {
    "noun": "Absolvent",
    "article": "der"
  },
  {
    "noun": "Absorber",
    "article": "der"
  },
  {
    "noun": "Abspannmast",
    "article": "der"
  },
  {
    "noun": "Abspanntransformator",
    "article": "der"
  },
  {
    "noun": "Absperrhahn",
    "article": "der"
  },
  {
    "noun": "Absperrschieber",
    "article": "der"
  },
  {
    "noun": "Absprung",
    "article": "der"
  },
  {
    "noun": "Absprungbalken",
    "article": "der"
  },
  {
    "noun": "Abstand",
    "article": "der"
  },
  {
    "noun": "Abstandhalter",
    "article": "der"
  },
  {
    "noun": "Abstauber",
    "article": "der"
  },
  {
    "noun": "Abstecher",
    "article": "der"
  },
  {
    "noun": "Absteigequartier",
    "article": "der"
  },
  {
    "noun": "Absteiger",
    "article": "der"
  },
  {
    "noun": "Abstellbahnhof",
    "article": "der"
  },
  {
    "noun": "Abstellhahn",
    "article": "der"
  },
  {
    "noun": "Abstellplatz",
    "article": "der"
  },
  {
    "noun": "Abstellraum",
    "article": "der"
  },
  {
    "noun": "Abstich",
    "article": "der"
  },
  {
    "noun": "Abstieg",
    "article": "der"
  },
  {
    "noun": "Abstinenzler",
    "article": "der"
  },
  {
    "noun": "Abstraktionsgrad",
    "article": "der"
  },
  {
    "noun": "Abstreicher",
    "article": "der"
  },
  {
    "noun": "Abstreifer",
    "article": "der"
  },
  {
    "noun": "Abstrich",
    "article": "der"
  },
  {
    "noun": "Absturz",
    "article": "der"
  },
  {
    "noun": "Absud",
    "article": "der"
  },
  {
    "noun": "Abt",
    "article": "der"
  },
  {
    "noun": "Abtaster",
    "article": "der"
  },
  {
    "noun": "Abtausch",
    "article": "der"
  },
  {
    "noun": "Abteilungsdirektor",
    "article": "der"
  },
  {
    "noun": "Abteilungsleiter",
    "article": "der"
  },
  {
    "noun": "Abteilwagen",
    "article": "der"
  },
  {
    "noun": "Abtrag",
    "article": "der"
  },
  {
    "noun": "Abtransport",
    "article": "der"
  },
  {
    "noun": "Abtreiber",
    "article": "der"
  },
  {
    "noun": "Abtreibungsarzt",
    "article": "der"
  },
  {
    "noun": "Abtreibungsgegner",
    "article": "der"
  },
  {
    "noun": "Abtreter",
    "article": "der"
  },
  {
    "noun": "Abtrieb",
    "article": "der"
  },
  {
    "noun": "Abtritt",
    "article": "der"
  },
  {
    "noun": "Abusus",
    "article": "der"
  },
  {
    "noun": "Abwasch",
    "article": "der"
  },
  {
    "noun": "Abwaschlappen",
    "article": "der"
  },
  {
    "noun": "Abwaschtisch",
    "article": "der"
  },
  {
    "noun": "Abwasserkanal",
    "article": "der"
  },
  {
    "noun": "Abweg",
    "article": "der"
  },
  {
    "noun": "Abwehrmechanismus",
    "article": "der"
  },
  {
    "noun": "Abwehrspieler",
    "article": "der"
  },
  {
    "noun": "Abwehrstoff",
    "article": "der"
  },
  {
    "noun": "Abweichler",
    "article": "der"
  },
  {
    "noun": "Abweiser",
    "article": "der"
  },
  {
    "noun": "Abwerber",
    "article": "der"
  },
  {
    "noun": "Abwickler",
    "article": "der"
  },
  {
    "noun": "Abwind",
    "article": "der"
  },
  {
    "noun": "Abwurf",
    "article": "der"
  },
  {
    "noun": "Abzahlungskauf",
    "article": "der"
  },
  {
    "noun": "Abzahlungsvertrag",
    "article": "der"
  },
  {
    "noun": "Abzieher",
    "article": "der"
  },
  {
    "noun": "Abziehstein",
    "article": "der"
  },
  {
    "noun": "Abzug",
    "article": "der"
  },
  {
    "noun": "Abzugsgraben",
    "article": "der"
  },
  {
    "noun": "Abzugshebel",
    "article": "der"
  },
  {
    "noun": "Abzugskanal",
    "article": "der"
  },
  {
    "noun": "Abzugsschacht",
    "article": "der"
  },
  {
    "noun": "Abzweig",
    "article": "der"
  },
  {
    "noun": "Abzweigkasten",
    "article": "der"
  },
  {
    "noun": "Acetaldehyd",
    "article": "der"
  },
  {
    "noun": "Achat",
    "article": "der"
  },
  {
    "noun": "Achsabstand",
    "article": "der"
  },
  {
    "noun": "Achsantrieb",
    "article": "der"
  },
  {
    "noun": "Achsbruch",
    "article": "der"
  },
  {
    "noun": "Achsdruck",
    "article": "der"
  },
  {
    "noun": "Achselgriff",
    "article": "der"
  },
  {
    "noun": "Achselspross",
    "article": "der"
  },
  {
    "noun": "Achsenbruch",
    "article": "der"
  },
  {
    "noun": "Achsnagel",
    "article": "der"
  },
  {
    "noun": "Achsschenkel",
    "article": "der"
  },
  {
    "noun": "Achsschenkelbolzen",
    "article": "der"
  },
  {
    "noun": "Achsstand",
    "article": "der"
  },
  {
    "noun": "Achtender",
    "article": "der"
  },
  {
    "noun": "Achter",
    "article": "der"
  },
  {
    "noun": "Achtersteven",
    "article": "der"
  },
  {
    "noun": "Achtkampf",
    "article": "der"
  },
  {
    "noun": "Achtstundentag",
    "article": "der"
  },
  {
    "noun": "Achtungserfolg",
    "article": "der"
  },
  {
    "noun": "Achtzylinder",
    "article": "der"
  },
  {
    "noun": "Acker",
    "article": "der"
  },
  {
    "noun": "Ackerbau",
    "article": "der"
  },
  {
    "noun": "Ackerbauer",
    "article": "der"
  },
  {
    "noun": "Ackerboden",
    "article": "der"
  },
  {
    "noun": "Ackergaul",
    "article": "der"
  },
  {
    "noun": "Ackersalat",
    "article": "der"
  },
  {
    "noun": "Ackerschlepper",
    "article": "der"
  },
  {
    "noun": "Ackersenf",
    "article": "der"
  },
  {
    "noun": "Ackerwagen",
    "article": "der"
  },
  {
    "noun": "Adamant",
    "article": "der"
  },
  {
    "noun": "Adamit",
    "article": "der"
  },
  {
    "noun": "Adamsapfel",
    "article": "der"
  },
  {
    "noun": "Adapter",
    "article": "der"
  },
  {
    "noun": "Addend",
    "article": "der"
  },
  {
    "noun": "Adduktor",
    "article": "der"
  },
  {
    "noun": "Adel",
    "article": "der"
  },
  {
    "noun": "Adelsbrief",
    "article": "der"
  },
  {
    "noun": "Adelsstand",
    "article": "der"
  },
  {
    "noun": "Adelstitel",
    "article": "der"
  },
  {
    "noun": "Adept",
    "article": "der"
  },
  {
    "noun": "Aderlass",
    "article": "der"
  },
  {
    "noun": "Adjutant",
    "article": "der"
  },
  {
    "noun": "Adlatus",
    "article": "der"
  },
  {
    "noun": "Adler",
    "article": "der"
  },
  {
    "noun": "Adlerblick",
    "article": "der"
  },
  {
    "noun": "Adlerfarn",
    "article": "der"
  },
  {
    "noun": "Adlerhorst",
    "article": "der"
  },
  {
    "noun": "Administrator",
    "article": "der"
  },
  {
    "noun": "Admiral",
    "article": "der"
  },
  {
    "noun": "Admiralstab",
    "article": "der"
  },
  {
    "noun": "Adobe",
    "article": "der"
  },
  {
    "noun": "Adoleszent",
    "article": "der"
  },
  {
    "noun": "Adoptivbruder",
    "article": "der"
  },
  {
    "noun": "Adoptivsohn",
    "article": "der"
  },
  {
    "noun": "Adressant",
    "article": "der"
  },
  {
    "noun": "Adressat",
    "article": "der"
  },
  {
    "noun": "Adular",
    "article": "der"
  },
  {
    "noun": "Advent",
    "article": "der"
  },
  {
    "noun": "Adventismus",
    "article": "der"
  },
  {
    "noun": "Adventivkrater",
    "article": "der"
  },
  {
    "noun": "Adventskalender",
    "article": "der"
  },
  {
    "noun": "Adventskranz",
    "article": "der"
  },
  {
    "noun": "Adventsstern",
    "article": "der"
  },
  {
    "noun": "Advokat",
    "article": "der"
  },
  {
    "noun": "Aerolith",
    "article": "der"
  },
  {
    "noun": "Aeroplan",
    "article": "der"
  },
  {
    "noun": "Affe",
    "article": "der"
  },
  {
    "noun": "Affekt",
    "article": "der"
  },
  {
    "noun": "Affektionswert",
    "article": "der"
  },
  {
    "noun": "Affenarsch",
    "article": "der"
  },
  {
    "noun": "Affenbrotbaum",
    "article": "der"
  },
  {
    "noun": "Affenmensch",
    "article": "der"
  },
  {
    "noun": "Affenzahn",
    "article": "der"
  },
  {
    "noun": "Affodill",
    "article": "der"
  },
  {
    "noun": "Affront",
    "article": "der"
  },
  {
    "noun": "Afghane",
    "article": "der"
  },
  {
    "noun": "Afrikaner",
    "article": "der"
  },
  {
    "noun": "Afrikanist",
    "article": "der"
  },
  {
    "noun": "After",
    "article": "der"
  },
  {
    "noun": "Agent",
    "article": "der"
  },
  {
    "noun": "Agenturbericht",
    "article": "der"
  },
  {
    "noun": "Aggregatzustand",
    "article": "der"
  },
  {
    "noun": "Aggressionskrieg",
    "article": "der"
  },
  {
    "noun": "Aggressionstrieb",
    "article": "der"
  },
  {
    "noun": "Aggressor",
    "article": "der"
  },
  {
    "noun": "Agitator",
    "article": "der"
  },
  {
    "noun": "Agnostiker",
    "article": "der"
  },
  {
    "noun": "Agnostizismus",
    "article": "der"
  },
  {
    "noun": "Agonist",
    "article": "der"
  },
  {
    "noun": "Agrarexport",
    "article": "der"
  },
  {
    "noun": "Agrarier",
    "article": "der"
  },
  {
    "noun": "Agrarimport",
    "article": "der"
  },
  {
    "noun": "Agrarkredit",
    "article": "der"
  },
  {
    "noun": "Agrarmarkt",
    "article": "der"
  },
  {
    "noun": "Agrarpreis",
    "article": "der"
  },
  {
    "noun": "Agrarstaat",
    "article": "der"
  },
  {
    "noun": "Agrarwissenschaftler",
    "article": "der"
  },
  {
    "noun": "Agrarzoll",
    "article": "der"
  },
  {
    "noun": "Agronom",
    "article": "der"
  },
  {
    "noun": "Ahn",
    "article": "der"
  },
  {
    "noun": "Ahnenkult",
    "article": "der"
  },
  {
    "noun": "Ahnenpass",
    "article": "der"
  },
  {
    "noun": "Ahnherr",
    "article": "der"
  },
  {
    "noun": "Ahorn",
    "article": "der"
  },
  {
    "noun": "Ahornsirup",
    "article": "der"
  },
  {
    "noun": "Airbag",
    "article": "der"
  },
  {
    "noun": "Airbus",
    "article": "der"
  },
  {
    "noun": "Akademiker",
    "article": "der"
  },
  {
    "noun": "Akanthus",
    "article": "der"
  },
  {
    "noun": "Akkord",
    "article": "der"
  },
  {
    "noun": "Akkordarbeiter",
    "article": "der"
  },
  {
    "noun": "Akkordeonist",
    "article": "der"
  },
  {
    "noun": "Akkordlohn",
    "article": "der"
  },
  {
    "noun": "Akkordrichtsatz",
    "article": "der"
  },
  {
    "noun": "Akku",
    "article": "der"
  },
  {
    "noun": "Akkumulator",
    "article": "der"
  },
  {
    "noun": "Akkusativ",
    "article": "der"
  },
  {
    "noun": "Akosmismus",
    "article": "der"
  },
  {
    "noun": "Akquisiteur",
    "article": "der"
  },
  {
    "noun": "Akrobat",
    "article": "der"
  },
  {
    "noun": "Akt",
    "article": "der"
  },
  {
    "noun": "Aktendeckel",
    "article": "der"
  },
  {
    "noun": "Aktenkoffer",
    "article": "der"
  },
  {
    "noun": "Aktenordner",
    "article": "der"
  },
  {
    "noun": "Aktenschrank",
    "article": "der"
  },
  {
    "noun": "Aktenvermerk",
    "article": "der"
  },
  {
    "noun": "Akteur",
    "article": "der"
  },
  {
    "noun": "Aktienbesitz",
    "article": "der"
  },
  {
    "noun": "Aktienindex",
    "article": "der"
  },
  {
    "noun": "Aktieninhaber",
    "article": "der"
  },
  {
    "noun": "Aktienkurs",
    "article": "der"
  },
  {
    "noun": "Aktienmarkt",
    "article": "der"
  },
  {
    "noun": "Aktinograf",
    "article": "der"
  },
  {
    "noun": "Aktinolith",
    "article": "der"
  },
  {
    "noun": "Aktionismus",
    "article": "der"
  },
  {
    "noun": "Aktionsbereich",
    "article": "der"
  },
  {
    "noun": "Aktionsradius",
    "article": "der"
  },
  {
    "noun": "Aktivismus",
    "article": "der"
  },
  {
    "noun": "Aktivist",
    "article": "der"
  },
  {
    "noun": "Aktivposten",
    "article": "der"
  },
  {
    "noun": "Aktivsaldo",
    "article": "der"
  },
  {
    "noun": "Aktivurlaub",
    "article": "der"
  },
  {
    "noun": "Aktivzins",
    "article": "der"
  },
  {
    "noun": "Aktor",
    "article": "der"
  },
  {
    "noun": "Aktuar",
    "article": "der"
  },
  {
    "noun": "Akupunkteur",
    "article": "der"
  },
  {
    "noun": "Akustiker",
    "article": "der"
  },
  {
    "noun": "Akut",
    "article": "der"
  },
  {
    "noun": "Akzelerator",
    "article": "der"
  },
  {
    "noun": "Akzent",
    "article": "der"
  },
  {
    "noun": "Akzentbuchstabe",
    "article": "der"
  },
  {
    "noun": "Akzentwechsel",
    "article": "der"
  },
  {
    "noun": "Akzeptant",
    "article": "der"
  },
  {
    "noun": "Akzeptkredit",
    "article": "der"
  },
  {
    "noun": "Akzeptor",
    "article": "der"
  },
  {
    "noun": "Akzidenzdruck",
    "article": "der"
  },
  {
    "noun": "Alabaster",
    "article": "der"
  },
  {
    "noun": "Alarm",
    "article": "der"
  },
  {
    "noun": "Alarmruf",
    "article": "der"
  },
  {
    "noun": "Alaun",
    "article": "der"
  },
  {
    "noun": "Alaunstein",
    "article": "der"
  },
  {
    "noun": "Alaunstift",
    "article": "der"
  },
  {
    "noun": "Alb",
    "article": "der"
  },
  {
    "noun": "Albaner",
    "article": "der"
  },
  {
    "noun": "Albatros",
    "article": "der"
  },
  {
    "noun": "Albdruck",
    "article": "der"
  },
  {
    "noun": "Albinismus",
    "article": "der"
  },
  {
    "noun": "Albino",
    "article": "der"
  },
  {
    "noun": "Albtraum",
    "article": "der"
  },
  {
    "noun": "Alchemist",
    "article": "der"
  },
  {
    "noun": "Alchimist",
    "article": "der"
  },
  {
    "noun": "Aldehyd",
    "article": "der"
  },
  {
    "noun": "Alemanne",
    "article": "der"
  },
  {
    "noun": "Alexandrit",
    "article": "der"
  },
  {
    "noun": "Algerier",
    "article": "der"
  },
  {
    "noun": "Algorithmus",
    "article": "der"
  },
  {
    "noun": "Alibibeweis",
    "article": "der"
  },
  {
    "noun": "Alk",
    "article": "der"
  },
  {
    "noun": "Alkohol",
    "article": "der"
  },
  {
    "noun": "Alkoholeinfluss",
    "article": "der"
  },
  {
    "noun": "Alkoholgegner",
    "article": "der"
  },
  {
    "noun": "Alkoholgehalt",
    "article": "der"
  },
  {
    "noun": "Alkoholgenuss",
    "article": "der"
  },
  {
    "noun": "Alkoholiker",
    "article": "der"
  },
  {
    "noun": "Alkoholismus",
    "article": "der"
  },
  {
    "noun": "Alkoholkonsum",
    "article": "der"
  },
  {
    "noun": "Alkoholtest",
    "article": "der"
  },
  {
    "noun": "Alkoven",
    "article": "der"
  },
  {
    "noun": "Alleinbesitz",
    "article": "der"
  },
  {
    "noun": "Alleinbesitzer",
    "article": "der"
  },
  {
    "noun": "Alleinerbe",
    "article": "der"
  },
  {
    "noun": "Alleingang",
    "article": "der"
  },
  {
    "noun": "Alleingesellschafter",
    "article": "der"
  },
  {
    "noun": "Alleinhandel",
    "article": "der"
  },
  {
    "noun": "Alleinherrscher",
    "article": "der"
  },
  {
    "noun": "Alleininhaber",
    "article": "der"
  },
  {
    "noun": "Alleinverkauf",
    "article": "der"
  },
  {
    "noun": "Alleinvertreter",
    "article": "der"
  },
  {
    "noun": "Alleinvertretungsanspruch",
    "article": "der"
  },
  {
    "noun": "Alleinvertrieb",
    "article": "der"
  },
  {
    "noun": "Allergiepass",
    "article": "der"
  },
  {
    "noun": "Allergiker",
    "article": "der"
  },
  {
    "noun": "Allergologe",
    "article": "der"
  },
  {
    "noun": "Allerseelentag",
    "article": "der"
  },
  {
    "noun": "Allesfresser",
    "article": "der"
  },
  {
    "noun": "Alleskleber",
    "article": "der"
  },
  {
    "noun": "Allgemeinarzt",
    "article": "der"
  },
  {
    "noun": "Allgemeinbegriff",
    "article": "der"
  },
  {
    "noun": "Allgemeinmediziner",
    "article": "der"
  },
  {
    "noun": "Allgemeinplatz",
    "article": "der"
  },
  {
    "noun": "Allgemeinzustand",
    "article": "der"
  },
  {
    "noun": "Alligator",
    "article": "der"
  },
  {
    "noun": "Allradantrieb",
    "article": "der"
  },
  {
    "noun": "Allrounder",
    "article": "der"
  },
  {
    "noun": "Allstrommotor",
    "article": "der"
  },
  {
    "noun": "Alltag",
    "article": "der"
  },
  {
    "noun": "Alltagstrott",
    "article": "der"
  },
  {
    "noun": "Allzweckreiniger",
    "article": "der"
  },
  {
    "noun": "Almabtrieb",
    "article": "der"
  },
  {
    "noun": "Almanach",
    "article": "der"
  },
  {
    "noun": "Almandin",
    "article": "der"
  },
  {
    "noun": "Alpenpass",
    "article": "der"
  },
  {
    "noun": "Alpenverein",
    "article": "der"
  },
  {
    "noun": "Alpinismus",
    "article": "der"
  },
  {
    "noun": "Alpinist",
    "article": "der"
  },
  {
    "noun": "Alraun",
    "article": "der"
  },
  {
    "noun": "Altai",
    "article": "der"
  },
  {
    "noun": "Altan",
    "article": "der"
  },
  {
    "noun": "Altar",
    "article": "der"
  },
  {
    "noun": "Altaraufsatz",
    "article": "der"
  },
  {
    "noun": "Altarraum",
    "article": "der"
  },
  {
    "noun": "Altarwein",
    "article": "der"
  },
  {
    "noun": "Altbau",
    "article": "der"
  },
  {
    "noun": "Altenpfleger",
    "article": "der"
  },
  {
    "noun": "Alternativvorschlag",
    "article": "der"
  },
  {
    "noun": "Alternator",
    "article": "der"
  },
  {
    "noun": "Altersaufbau",
    "article": "der"
  },
  {
    "noun": "Altersdiabetes",
    "article": "der"
  },
  {
    "noun": "Altersfreibetrag",
    "article": "der"
  },
  {
    "noun": "Altersgenosse",
    "article": "der"
  },
  {
    "noun": "Altersstarrsinn",
    "article": "der"
  },
  {
    "noun": "Altersunterschied",
    "article": "der"
  },
  {
    "noun": "Altertumsforscher",
    "article": "der"
  },
  {
    "noun": "Altertumswert",
    "article": "der"
  },
  {
    "noun": "Alterungsprozess",
    "article": "der"
  },
  {
    "noun": "Altglascontainer",
    "article": "der"
  },
  {
    "noun": "Altkanzler",
    "article": "der"
  },
  {
    "noun": "Altmeister",
    "article": "der"
  },
  {
    "noun": "Altphilologe",
    "article": "der"
  },
  {
    "noun": "Altruismus",
    "article": "der"
  },
  {
    "noun": "Altruist",
    "article": "der"
  },
  {
    "noun": "Altstoff",
    "article": "der"
  },
  {
    "noun": "Altvater",
    "article": "der"
  },
  {
    "noun": "Altwarenhandel",
    "article": "der"
  },
  {
    "noun": "Altweibersommer",
    "article": "der"
  },
  {
    "noun": "Alveolar",
    "article": "der"
  },
  {
    "noun": "Amateur",
    "article": "der"
  },
  {
    "noun": "Amateurfotograf",
    "article": "der"
  },
  {
    "noun": "Amateurfunker",
    "article": "der"
  },
  {
    "noun": "Amazonas",
    "article": "der"
  },
  {
    "noun": "Amber",
    "article": "der"
  },
  {
    "noun": "Amboss",
    "article": "der"
  },
  {
    "noun": "Ambulanzwagen",
    "article": "der"
  },
  {
    "noun": "Ameisenhaufen",
    "article": "der"
  },
  {
    "noun": "Ameisenstaat",
    "article": "der"
  },
  {
    "noun": "Amethyst",
    "article": "der"
  },
  {
    "noun": "Ami",
    "article": "der"
  },
  {
    "noun": "Amok",
    "article": "der"
  },
  {
    "noun": "Amoklauf",
    "article": "der"
  },
  {
    "noun": "Amor",
    "article": "der"
  },
  {
    "noun": "Amortisationsfonds",
    "article": "der"
  },
  {
    "noun": "Ampfer",
    "article": "der"
  },
  {
    "noun": "Amphibol",
    "article": "der"
  },
  {
    "noun": "Amphibolit",
    "article": "der"
  },
  {
    "noun": "Amphibrachys",
    "article": "der"
  },
  {
    "noun": "Amtmann",
    "article": "der"
  },
  {
    "noun": "Amtsantritt",
    "article": "der"
  },
  {
    "noun": "Amtsarzt",
    "article": "der"
  },
  {
    "noun": "Amtsbereich",
    "article": "der"
  },
  {
    "noun": "Amtsbezirk",
    "article": "der"
  },
  {
    "noun": "Amtsbruder",
    "article": "der"
  },
  {
    "noun": "Amtsdiener",
    "article": "der"
  },
  {
    "noun": "Amtseid",
    "article": "der"
  },
  {
    "noun": "Amtsinhaber",
    "article": "der"
  },
  {
    "noun": "Amtskollege",
    "article": "der"
  },
  {
    "noun": "Amtsnachfolger",
    "article": "der"
  },
  {
    "noun": "Amtsschimmel",
    "article": "der"
  },
  {
    "noun": "Amtssitz",
    "article": "der"
  },
  {
    "noun": "Amtston",
    "article": "der"
  },
  {
    "noun": "Amtsverzicht",
    "article": "der"
  },
  {
    "noun": "Amtsvorsteher",
    "article": "der"
  },
  {
    "noun": "Anachoret",
    "article": "der"
  },
  {
    "noun": "Anachronismus",
    "article": "der"
  },
  {
    "noun": "Analogismus",
    "article": "der"
  },
  {
    "noun": "Analogrechner",
    "article": "der"
  },
  {
    "noun": "Analphabet",
    "article": "der"
  },
  {
    "noun": "Analphabetismus",
    "article": "der"
  },
  {
    "noun": "Analverkehr",
    "article": "der"
  },
  {
    "noun": "Analysator",
    "article": "der"
  },
  {
    "noun": "Analyst",
    "article": "der"
  },
  {
    "noun": "Analytiker",
    "article": "der"
  },
  {
    "noun": "Anarchismus",
    "article": "der"
  },
  {
    "noun": "Anarchist",
    "article": "der"
  },
  {
    "noun": "Anarcho",
    "article": "der"
  },
  {
    "noun": "Anatom",
    "article": "der"
  },
  {
    "noun": "Anatomiesaal",
    "article": "der"
  },
  {
    "noun": "Anbau",
    "article": "der"
  },
  {
    "noun": "Anbaumotor",
    "article": "der"
  },
  {
    "noun": "Anbauschrank",
    "article": "der"
  },
  {
    "noun": "Anbeginn",
    "article": "der"
  },
  {
    "noun": "Anbeter",
    "article": "der"
  },
  {
    "noun": "Anbiederungsversuch",
    "article": "der"
  },
  {
    "noun": "Anbieter",
    "article": "der"
  },
  {
    "noun": "Anbiss",
    "article": "der"
  },
  {
    "noun": "Anblick",
    "article": "der"
  },
  {
    "noun": "Anbruch",
    "article": "der"
  },
  {
    "noun": "Andalusier",
    "article": "der"
  },
  {
    "noun": "Andalusit",
    "article": "der"
  },
  {
    "noun": "Andesin",
    "article": "der"
  },
  {
    "noun": "Andesit",
    "article": "der"
  },
  {
    "noun": "Andorraner",
    "article": "der"
  },
  {
    "noun": "Andrang",
    "article": "der"
  },
  {
    "noun": "Android",
    "article": "der"
  },
  {
    "noun": "Androide",
    "article": "der"
  },
  {
    "noun": "Androloge",
    "article": "der"
  },
  {
    "noun": "Andromedanebel",
    "article": "der"
  },
  {
    "noun": "Andruck",
    "article": "der"
  },
  {
    "noun": "Anfahrtsweg",
    "article": "der"
  },
  {
    "noun": "Anfall",
    "article": "der"
  },
  {
    "noun": "Anfangsbuchstabe",
    "article": "der"
  },
  {
    "noun": "Anfangserfolg",
    "article": "der"
  },
  {
    "noun": "Anfangspunkt",
    "article": "der"
  },
  {
    "noun": "Anfangsunterricht",
    "article": "der"
  },
  {
    "noun": "Anfangswert",
    "article": "der"
  },
  {
    "noun": "Anflug",
    "article": "der"
  },
  {
    "noun": "Anforderungskatalog",
    "article": "der"
  },
  {
    "noun": "Angeber",
    "article": "der"
  },
  {
    "noun": "Angebotskurs",
    "article": "der"
  },
  {
    "noun": "Angebotspreis",
    "article": "der"
  },
  {
    "noun": "Angelhaken",
    "article": "der"
  },
  {
    "noun": "Angelplatz",
    "article": "der"
  },
  {
    "noun": "Angelpunkt",
    "article": "der"
  },
  {
    "noun": "Angelsachse",
    "article": "der"
  },
  {
    "noun": "Angelschein",
    "article": "der"
  },
  {
    "noun": "Angelsport",
    "article": "der"
  },
  {
    "noun": "Anger",
    "article": "der"
  },
  {
    "noun": "Angler",
    "article": "der"
  },
  {
    "noun": "Anglist",
    "article": "der"
  },
  {
    "noun": "Anglizismus",
    "article": "der"
  },
  {
    "noun": "Anglomane",
    "article": "der"
  },
  {
    "noun": "Angolaner",
    "article": "der"
  },
  {
    "noun": "Angreifer",
    "article": "der"
  },
  {
    "noun": "Angriff",
    "article": "der"
  },
  {
    "noun": "Angriffskrieg",
    "article": "der"
  },
  {
    "noun": "Angriffspunkt",
    "article": "der"
  },
  {
    "noun": "Angstgegner",
    "article": "der"
  },
  {
    "noun": "Angsthase",
    "article": "der"
  },
  {
    "noun": "Angstmacher",
    "article": "der"
  },
  {
    "noun": "Angstschrei",
    "article": "der"
  },
  {
    "noun": "Angstzustand",
    "article": "der"
  },
  {
    "noun": "Anguss",
    "article": "der"
  },
  {
    "noun": "Anhalter",
    "article": "der"
  },
  {
    "noun": "Anhalteweg",
    "article": "der"
  },
  {
    "noun": "Anhaltspunkt",
    "article": "der"
  },
  {
    "noun": "Anhang",
    "article": "der"
  },
  {
    "noun": "Anhauch",
    "article": "der"
  },
  {
    "noun": "Anhydrit",
    "article": "der"
  },
  {
    "noun": "Animateur",
    "article": "der"
  },
  {
    "noun": "Animationsfilm",
    "article": "der"
  },
  {
    "noun": "Animismus",
    "article": "der"
  },
  {
    "noun": "Animist",
    "article": "der"
  },
  {
    "noun": "Anis",
    "article": "der"
  },
  {
    "noun": "Ankauf",
    "article": "der"
  },
  {
    "noun": "Ankaufspreis",
    "article": "der"
  },
  {
    "noun": "Anker",
    "article": "der"
  },
  {
    "noun": "Ankergrund",
    "article": "der"
  },
  {
    "noun": "Ankerplatz",
    "article": "der"
  },
  {
    "noun": "Anklagepunkt",
    "article": "der"
  },
  {
    "noun": "Anklagevertreter",
    "article": "der"
  },
  {
    "noun": "Anklang",
    "article": "der"
  },
  {
    "noun": "Ankleider",
    "article": "der"
  },
  {
    "noun": "Ankleideraum",
    "article": "der"
  },
  {
    "noun": "Ankunftsort",
    "article": "der"
  },
  {
    "noun": "Anlageberater",
    "article": "der"
  },
  {
    "noun": "Anlagefonds",
    "article": "der"
  },
  {
    "noun": "Anlagenbau",
    "article": "der"
  },
  {
    "noun": "Anlass",
    "article": "der"
  },
  {
    "noun": "Anlasser",
    "article": "der"
  },
  {
    "noun": "Anlauf",
    "article": "der"
  },
  {
    "noun": "Anlaufverlust",
    "article": "der"
  },
  {
    "noun": "Anlaut",
    "article": "der"
  },
  {
    "noun": "Anlegeplatz",
    "article": "der"
  },
  {
    "noun": "Anleger",
    "article": "der"
  },
  {
    "noun": "Anlegesteg",
    "article": "der"
  },
  {
    "noun": "Anleihemarkt",
    "article": "der"
  },
  {
    "noun": "Anleiter",
    "article": "der"
  },
  {
    "noun": "Anlernberuf",
    "article": "der"
  },
  {
    "noun": "Anlernling",
    "article": "der"
  },
  {
    "noun": "Anlieger",
    "article": "der"
  },
  {
    "noun": "Anliegerstaat",
    "article": "der"
  },
  {
    "noun": "Anliegerverkehr",
    "article": "der"
  },
  {
    "noun": "Anmarsch",
    "article": "der"
  },
  {
    "noun": "Anmeldeschein",
    "article": "der"
  },
  {
    "noun": "Anmeldeschluss",
    "article": "der"
  },
  {
    "noun": "Annex",
    "article": "der"
  },
  {
    "noun": "Anodenstrom",
    "article": "der"
  },
  {
    "noun": "Anorak",
    "article": "der"
  },
  {
    "noun": "Anorthit",
    "article": "der"
  },
  {
    "noun": "Anorthosit",
    "article": "der"
  },
  {
    "noun": "Anpasser",
    "article": "der"
  },
  {
    "noun": "Anpassungsmechanismus",
    "article": "der"
  },
  {
    "noun": "Anpassungsprozess",
    "article": "der"
  },
  {
    "noun": "Anpfiff",
    "article": "der"
  },
  {
    "noun": "Anprall",
    "article": "der"
  },
  {
    "noun": "Anrainer",
    "article": "der"
  },
  {
    "noun": "Anrainerstaat",
    "article": "der"
  },
  {
    "noun": "Anrechnungszeitraum",
    "article": "der"
  },
  {
    "noun": "Anrechtsschein",
    "article": "der"
  },
  {
    "noun": "Anredefall",
    "article": "der"
  },
  {
    "noun": "Anreger",
    "article": "der"
  },
  {
    "noun": "Anreisetag",
    "article": "der"
  },
  {
    "noun": "Anreisetermin",
    "article": "der"
  },
  {
    "noun": "Anreiseweg",
    "article": "der"
  },
  {
    "noun": "Anreiz",
    "article": "der"
  },
  {
    "noun": "Anriss",
    "article": "der"
  },
  {
    "noun": "Anruf",
    "article": "der"
  },
  {
    "noun": "Anrufbeantworter",
    "article": "der"
  },
  {
    "noun": "Anrufer",
    "article": "der"
  },
  {
    "noun": "Ansager",
    "article": "der"
  },
  {
    "noun": "Ansatz",
    "article": "der"
  },
  {
    "noun": "Ansatzpunkt",
    "article": "der"
  },
  {
    "noun": "Ansaugdruck",
    "article": "der"
  },
  {
    "noun": "Ansaugkanal",
    "article": "der"
  },
  {
    "noun": "Ansaugtakt",
    "article": "der"
  },
  {
    "noun": "Anschaffungspreis",
    "article": "der"
  },
  {
    "noun": "Anschaffungswert",
    "article": "der"
  },
  {
    "noun": "Anschauungsunterricht",
    "article": "der"
  },
  {
    "noun": "Anschein",
    "article": "der"
  },
  {
    "noun": "Anschlag",
    "article": "der"
  },
  {
    "noun": "Anschlagwinkel",
    "article": "der"
  },
  {
    "noun": "Anschluss",
    "article": "der"
  },
  {
    "noun": "Anschlussflug",
    "article": "der"
  },
  {
    "noun": "Anschlusstreffer",
    "article": "der"
  },
  {
    "noun": "Anschlusszug",
    "article": "der"
  },
  {
    "noun": "Anschnallgurt",
    "article": "der"
  },
  {
    "noun": "Anschnitt",
    "article": "der"
  },
  {
    "noun": "Anschwung",
    "article": "der"
  },
  {
    "noun": "Ansiedler",
    "article": "der"
  },
  {
    "noun": "Ansitz",
    "article": "der"
  },
  {
    "noun": "Anspitzer",
    "article": "der"
  },
  {
    "noun": "Ansporn",
    "article": "der"
  },
  {
    "noun": "Ansprechpartner",
    "article": "der"
  },
  {
    "noun": "Anspruch",
    "article": "der"
  },
  {
    "noun": "Anstand",
    "article": "der"
  },
  {
    "noun": "Anstandsbesuch",
    "article": "der"
  },
  {
    "noun": "Anstandshappen",
    "article": "der"
  },
  {
    "noun": "Anstellungsvertrag",
    "article": "der"
  },
  {
    "noun": "Anstieg",
    "article": "der"
  },
  {
    "noun": "Anstifter",
    "article": "der"
  },
  {
    "noun": "Anstreicher",
    "article": "der"
  },
  {
    "noun": "Anstrich",
    "article": "der"
  },
  {
    "noun": "Ansturm",
    "article": "der"
  },
  {
    "noun": "Ansucher",
    "article": "der"
  },
  {
    "noun": "Antagonismus",
    "article": "der"
  },
  {
    "noun": "Antagonist",
    "article": "der"
  },
  {
    "noun": "Anteil",
    "article": "der"
  },
  {
    "noun": "Anteilschein",
    "article": "der"
  },
  {
    "noun": "Anteilseigner",
    "article": "der"
  },
  {
    "noun": "Antennendraht",
    "article": "der"
  },
  {
    "noun": "Antennenmast",
    "article": "der"
  },
  {
    "noun": "Antennenwiderstand",
    "article": "der"
  },
  {
    "noun": "Anthropologe",
    "article": "der"
  },
  {
    "noun": "Antialkoholiker",
    "article": "der"
  },
  {
    "noun": "Antidumpingzoll",
    "article": "der"
  },
  {
    "noun": "Antifaschismus",
    "article": "der"
  },
  {
    "noun": "Antifaschist",
    "article": "der"
  },
  {
    "noun": "Antiheld",
    "article": "der"
  },
  {
    "noun": "Antikommunist",
    "article": "der"
  },
  {
    "noun": "Antilogarithmus",
    "article": "der"
  },
  {
    "noun": "Antipode",
    "article": "der"
  },
  {
    "noun": "Antiquar",
    "article": "der"
  },
  {
    "noun": "Antisemit",
    "article": "der"
  },
  {
    "noun": "Antisemitismus",
    "article": "der"
  },
  {
    "noun": "Antrag",
    "article": "der"
  },
  {
    "noun": "Antragsteller",
    "article": "der"
  },
  {
    "noun": "Antreiber",
    "article": "der"
  },
  {
    "noun": "Antrieb",
    "article": "der"
  },
  {
    "noun": "Antritt",
    "article": "der"
  },
  {
    "noun": "Antrittsbesuch",
    "article": "der"
  },
  {
    "noun": "Antwortschein",
    "article": "der"
  },
  {
    "noun": "Anus",
    "article": "der"
  },
  {
    "noun": "Anwachs",
    "article": "der"
  },
  {
    "noun": "Anwalt",
    "article": "der"
  },
  {
    "noun": "Anwaltsgehilfe",
    "article": "der"
  },
  {
    "noun": "Anwaltszwang",
    "article": "der"
  },
  {
    "noun": "Anwender",
    "article": "der"
  },
  {
    "noun": "Anwendungsbereich",
    "article": "der"
  },
  {
    "noun": "Anwohner",
    "article": "der"
  },
  {
    "noun": "Anwuchs",
    "article": "der"
  },
  {
    "noun": "Anwurf",
    "article": "der"
  },
  {
    "noun": "Anzeigenteil",
    "article": "der"
  },
  {
    "noun": "Anzeiger",
    "article": "der"
  },
  {
    "noun": "Anziehungspunkt",
    "article": "der"
  },
  {
    "noun": "Anzug",
    "article": "der"
  },
  {
    "noun": "Anzugstoff",
    "article": "der"
  },
  {
    "noun": "Aorist",
    "article": "der"
  },
  {
    "noun": "Aortenbogen",
    "article": "der"
  },
  {
    "noun": "Apache",
    "article": "der"
  },
  {
    "noun": "Apalliker",
    "article": "der"
  },
  {
    "noun": "Apatit",
    "article": "der"
  },
  {
    "noun": "Apennin",
    "article": "der"
  },
  {
    "noun": "Aperitif",
    "article": "der"
  },
  {
    "noun": "Aperwind",
    "article": "der"
  },
  {
    "noun": "Apex",
    "article": "der"
  },
  {
    "noun": "Apfel",
    "article": "der"
  },
  {
    "noun": "Apfelbaum",
    "article": "der"
  },
  {
    "noun": "Apfelkern",
    "article": "der"
  },
  {
    "noun": "Apfelkuchen",
    "article": "der"
  },
  {
    "noun": "Apfelmost",
    "article": "der"
  },
  {
    "noun": "Apfelpfannkuchen",
    "article": "der"
  },
  {
    "noun": "Apfelsaft",
    "article": "der"
  },
  {
    "noun": "Apfelschimmel",
    "article": "der"
  },
  {
    "noun": "Apfelsinenbaum",
    "article": "der"
  },
  {
    "noun": "Apfelsinensaft",
    "article": "der"
  },
  {
    "noun": "Apfelstrudel",
    "article": "der"
  },
  {
    "noun": "Apfelwein",
    "article": "der"
  },
  {
    "noun": "Aphorismus",
    "article": "der"
  },
  {
    "noun": "Aphoristiker",
    "article": "der"
  },
  {
    "noun": "Apologet",
    "article": "der"
  },
  {
    "noun": "Apophyllit",
    "article": "der"
  },
  {
    "noun": "Apostat",
    "article": "der"
  },
  {
    "noun": "Apostel",
    "article": "der"
  },
  {
    "noun": "Apostelbrief",
    "article": "der"
  },
  {
    "noun": "Apostroph",
    "article": "der"
  },
  {
    "noun": "Apotheker",
    "article": "der"
  },
  {
    "noun": "Apparat",
    "article": "der"
  },
  {
    "noun": "Apparatebau",
    "article": "der"
  },
  {
    "noun": "Apparatschik",
    "article": "der"
  },
  {
    "noun": "Appeal",
    "article": "der"
  },
  {
    "noun": "Appell",
    "article": "der"
  },
  {
    "noun": "Appellplatz",
    "article": "der"
  },
  {
    "noun": "Appendix",
    "article": "der"
  },
  {
    "noun": "Appetit",
    "article": "der"
  },
  {
    "noun": "Appetitbissen",
    "article": "der"
  },
  {
    "noun": "Appetithappen",
    "article": "der"
  },
  {
    "noun": "Applaus",
    "article": "der"
  },
  {
    "noun": "Applikator",
    "article": "der"
  },
  {
    "noun": "Apportierhund",
    "article": "der"
  },
  {
    "noun": "Appreteur",
    "article": "der"
  },
  {
    "noun": "Approach",
    "article": "der"
  },
  {
    "noun": "Aprikosenbaum",
    "article": "der"
  },
  {
    "noun": "Aprikosensaft",
    "article": "der"
  },
  {
    "noun": "April",
    "article": "der"
  },
  {
    "noun": "Aprilnarr",
    "article": "der"
  },
  {
    "noun": "Aprilscherz",
    "article": "der"
  },
  {
    "noun": "Aquamarin",
    "article": "der"
  },
  {
    "noun": "Aquanaut",
    "article": "der"
  },
  {
    "noun": "Aquavit",
    "article": "der"
  },
  {
    "noun": "Araber",
    "article": "der"
  },
  {
    "noun": "Arachnologe",
    "article": "der"
  },
  {
    "noun": "Aragonit",
    "article": "der"
  },
  {
    "noun": "Aralsee",
    "article": "der"
  },
  {
    "noun": "Arbeiter",
    "article": "der"
  },
  {
    "noun": "Arbeiterpriester",
    "article": "der"
  },
  {
    "noun": "Arbeiterrat",
    "article": "der"
  },
  {
    "noun": "Arbeitgeber",
    "article": "der"
  },
  {
    "noun": "Arbeitgeberanteil",
    "article": "der"
  },
  {
    "noun": "Arbeitgeberbeitrag",
    "article": "der"
  },
  {
    "noun": "Arbeitgeberverband",
    "article": "der"
  },
  {
    "noun": "Arbeitgebervertreter",
    "article": "der"
  },
  {
    "noun": "Arbeitnehmer",
    "article": "der"
  },
  {
    "noun": "Arbeitnehmervertreter",
    "article": "der"
  },
  {
    "noun": "Arbeitsablauf",
    "article": "der"
  },
  {
    "noun": "Arbeitsanfall",
    "article": "der"
  },
  {
    "noun": "Arbeitsanfang",
    "article": "der"
  },
  {
    "noun": "Arbeitsanzug",
    "article": "der"
  },
  {
    "noun": "Arbeitsaufwand",
    "article": "der"
  },
  {
    "noun": "Arbeitsausschuss",
    "article": "der"
  },
  {
    "noun": "Arbeitsbeginn",
    "article": "der"
  },
  {
    "noun": "Arbeitsbericht",
    "article": "der"
  },
  {
    "noun": "Arbeitsbesuch",
    "article": "der"
  },
  {
    "noun": "Arbeitsdienst",
    "article": "der"
  },
  {
    "noun": "Arbeitsdirektor",
    "article": "der"
  },
  {
    "noun": "Arbeitsertrag",
    "article": "der"
  },
  {
    "noun": "Arbeitsfriede",
    "article": "der"
  },
  {
    "noun": "Arbeitsgang",
    "article": "der"
  },
  {
    "noun": "Arbeitsimmigrant",
    "article": "der"
  },
  {
    "noun": "Arbeitskamerad",
    "article": "der"
  },
  {
    "noun": "Arbeitskampf",
    "article": "der"
  },
  {
    "noun": "Arbeitskittel",
    "article": "der"
  },
  {
    "noun": "Arbeitskollege",
    "article": "der"
  },
  {
    "noun": "Arbeitskonflikt",
    "article": "der"
  },
  {
    "noun": "Arbeitskreis",
    "article": "der"
  },
  {
    "noun": "Arbeitslohn",
    "article": "der"
  },
  {
    "noun": "Arbeitsmann",
    "article": "der"
  },
  {
    "noun": "Arbeitsmarkt",
    "article": "der"
  },
  {
    "noun": "Arbeitsmediziner",
    "article": "der"
  },
  {
    "noun": "Arbeitsminister",
    "article": "der"
  },
  {
    "noun": "Arbeitsnachweis",
    "article": "der"
  },
  {
    "noun": "Arbeitsort",
    "article": "der"
  },
  {
    "noun": "Arbeitsplan",
    "article": "der"
  },
  {
    "noun": "Arbeitsplatz",
    "article": "der"
  },
  {
    "noun": "Arbeitsplatzverlust",
    "article": "der"
  },
  {
    "noun": "Arbeitsplatzwechsel",
    "article": "der"
  },
  {
    "noun": "Arbeitsprozess",
    "article": "der"
  },
  {
    "noun": "Arbeitsraum",
    "article": "der"
  },
  {
    "noun": "Arbeitsrhythmus",
    "article": "der"
  },
  {
    "noun": "Arbeitsschluss",
    "article": "der"
  },
  {
    "noun": "Arbeitsschritt",
    "article": "der"
  },
  {
    "noun": "Arbeitsschutz",
    "article": "der"
  },
  {
    "noun": "Arbeitsspeicher",
    "article": "der"
  },
  {
    "noun": "Arbeitsstab",
    "article": "der"
  },
  {
    "noun": "Arbeitsstrom",
    "article": "der"
  },
  {
    "noun": "Arbeitstag",
    "article": "der"
  },
  {
    "noun": "Arbeitstakt",
    "article": "der"
  },
  {
    "noun": "Arbeitstisch",
    "article": "der"
  },
  {
    "noun": "Arbeitstitel",
    "article": "der"
  },
  {
    "noun": "Arbeitsunfall",
    "article": "der"
  },
  {
    "noun": "Arbeitsurlaub",
    "article": "der"
  },
  {
    "noun": "Arbeitsverlust",
    "article": "der"
  },
  {
    "noun": "Arbeitsvertrag",
    "article": "der"
  },
  {
    "noun": "Arbeitsvorbereiter",
    "article": "der"
  },
  {
    "noun": "Arbeitsvorgang",
    "article": "der"
  },
  {
    "noun": "Arbeitswille",
    "article": "der"
  },
  {
    "noun": "Arbeitszwang",
    "article": "der"
  },
  {
    "noun": "Archaismus",
    "article": "der"
  },
  {
    "noun": "Archetyp",
    "article": "der"
  },
  {
    "noun": "Archetypus",
    "article": "der"
  },
  {
    "noun": "Archidiakon",
    "article": "der"
  },
  {
    "noun": "Archipel",
    "article": "der"
  },
  {
    "noun": "Architekt",
    "article": "der"
  },
  {
    "noun": "Architekturwettbewerb",
    "article": "der"
  },
  {
    "noun": "Architrav",
    "article": "der"
  },
  {
    "noun": "Argentinier",
    "article": "der"
  },
  {
    "noun": "Argentit",
    "article": "der"
  },
  {
    "noun": "Argwohn",
    "article": "der"
  },
  {
    "noun": "Arier",
    "article": "der"
  },
  {
    "noun": "Aristokrat",
    "article": "der"
  },
  {
    "noun": "Arithmetiker",
    "article": "der"
  },
  {
    "noun": "Arkus",
    "article": "der"
  },
  {
    "noun": "Armbruch",
    "article": "der"
  },
  {
    "noun": "Armenier",
    "article": "der"
  },
  {
    "noun": "Armleuchter",
    "article": "der"
  },
  {
    "noun": "Armmuskel",
    "article": "der"
  },
  {
    "noun": "Armreif",
    "article": "der"
  },
  {
    "noun": "Armschutz",
    "article": "der"
  },
  {
    "noun": "Armsessel",
    "article": "der"
  },
  {
    "noun": "Armvoll",
    "article": "der"
  },
  {
    "noun": "Aronstab",
    "article": "der"
  },
  {
    "noun": "Arrak",
    "article": "der"
  },
  {
    "noun": "Arrangeur",
    "article": "der"
  },
  {
    "noun": "Arrest",
    "article": "der"
  },
  {
    "noun": "Arrestant",
    "article": "der"
  },
  {
    "noun": "Arsch",
    "article": "der"
  },
  {
    "noun": "Arschkriecher",
    "article": "der"
  },
  {
    "noun": "Arschlecker",
    "article": "der"
  },
  {
    "noun": "Arsenwasserstoff",
    "article": "der"
  },
  {
    "noun": "Artenreichtum",
    "article": "der"
  },
  {
    "noun": "Artenschutz",
    "article": "der"
  },
  {
    "noun": "Artgenosse",
    "article": "der"
  },
  {
    "noun": "Artikel",
    "article": "der"
  },
  {
    "noun": "Artillerist",
    "article": "der"
  },
  {
    "noun": "Artist",
    "article": "der"
  },
  {
    "noun": "Artname",
    "article": "der"
  },
  {
    "noun": "Arzt",
    "article": "der"
  },
  {
    "noun": "Arztberuf",
    "article": "der"
  },
  {
    "noun": "Arzthelfer",
    "article": "der"
  },
  {
    "noun": "Arzttermin",
    "article": "der"
  },
  {
    "noun": "Asbest",
    "article": "der"
  },
  {
    "noun": "Asbestzement",
    "article": "der"
  },
  {
    "noun": "Aschenbecher",
    "article": "der"
  },
  {
    "noun": "Aschenkasten",
    "article": "der"
  },
  {
    "noun": "Aschenkrug",
    "article": "der"
  },
  {
    "noun": "Aschenregen",
    "article": "der"
  },
  {
    "noun": "Ascher",
    "article": "der"
  },
  {
    "noun": "Aschermittwoch",
    "article": "der"
  },
  {
    "noun": "Aschkasten",
    "article": "der"
  },
  {
    "noun": "Asiat",
    "article": "der"
  },
  {
    "noun": "Asket",
    "article": "der"
  },
  {
    "noun": "Asketiker",
    "article": "der"
  },
  {
    "noun": "Aspekt",
    "article": "der"
  },
  {
    "noun": "Asphalt",
    "article": "der"
  },
  {
    "noun": "Asphaltlack",
    "article": "der"
  },
  {
    "noun": "Aspik",
    "article": "der"
  },
  {
    "noun": "Aspirant",
    "article": "der"
  },
  {
    "noun": "Aspirator",
    "article": "der"
  },
  {
    "noun": "Assekurant",
    "article": "der"
  },
  {
    "noun": "Assembler",
    "article": "der"
  },
  {
    "noun": "Assessor",
    "article": "der"
  },
  {
    "noun": "Assignant",
    "article": "der"
  },
  {
    "noun": "Assignat",
    "article": "der"
  },
  {
    "noun": "Assignatar",
    "article": "der"
  },
  {
    "noun": "Assistent",
    "article": "der"
  },
  {
    "noun": "Assistenzarzt",
    "article": "der"
  },
  {
    "noun": "Assistenzprofessor",
    "article": "der"
  },
  {
    "noun": "Ast",
    "article": "der"
  },
  {
    "noun": "Asteriskus",
    "article": "der"
  },
  {
    "noun": "Asteroid",
    "article": "der"
  },
  {
    "noun": "Astheniker",
    "article": "der"
  },
  {
    "noun": "Asthmaanfall",
    "article": "der"
  },
  {
    "noun": "Asthmatiker",
    "article": "der"
  },
  {
    "noun": "Astigmatismus",
    "article": "der"
  },
  {
    "noun": "Astrologe",
    "article": "der"
  },
  {
    "noun": "Astronaut",
    "article": "der"
  },
  {
    "noun": "Astronom",
    "article": "der"
  },
  {
    "noun": "Astrophysiker",
    "article": "der"
  },
  {
    "noun": "Asylant",
    "article": "der"
  },
  {
    "noun": "Asylantrag",
    "article": "der"
  },
  {
    "noun": "Asylbewerber",
    "article": "der"
  },
  {
    "noun": "Asylwerber",
    "article": "der"
  },
  {
    "noun": "Asynchronmotor",
    "article": "der"
  },
  {
    "noun": "Aszendent",
    "article": "der"
  },
  {
    "noun": "Aszites",
    "article": "der"
  },
  {
    "noun": "Atavismus",
    "article": "der"
  },
  {
    "noun": "Atem",
    "article": "der"
  },
  {
    "noun": "Atemstillstand",
    "article": "der"
  },
  {
    "noun": "Atemweg",
    "article": "der"
  },
  {
    "noun": "Atemzug",
    "article": "der"
  },
  {
    "noun": "Atheismus",
    "article": "der"
  },
  {
    "noun": "Atheist",
    "article": "der"
  },
  {
    "noun": "Athlet",
    "article": "der"
  },
  {
    "noun": "Athletiker",
    "article": "der"
  },
  {
    "noun": "Atlant",
    "article": "der"
  },
  {
    "noun": "Atlantik",
    "article": "der"
  },
  {
    "noun": "Atlantikwall",
    "article": "der"
  },
  {
    "noun": "Atombunker",
    "article": "der"
  },
  {
    "noun": "Atomismus",
    "article": "der"
  },
  {
    "noun": "Atomkern",
    "article": "der"
  },
  {
    "noun": "Atomkraftgegner",
    "article": "der"
  },
  {
    "noun": "Atomkrieg",
    "article": "der"
  },
  {
    "noun": "Atommeiler",
    "article": "der"
  },
  {
    "noun": "Atomphysiker",
    "article": "der"
  },
  {
    "noun": "Atompilz",
    "article": "der"
  },
  {
    "noun": "Atomreaktor",
    "article": "der"
  },
  {
    "noun": "Atomsprengkopf",
    "article": "der"
  },
  {
    "noun": "Atomtest",
    "article": "der"
  },
  {
    "noun": "Atomversuch",
    "article": "der"
  },
  {
    "noun": "Atomwaffensperrvertrag",
    "article": "der"
  },
  {
    "noun": "Atout",
    "article": "der"
  },
  {
    "noun": "Attributsatz",
    "article": "der"
  },
  {
    "noun": "Auditor",
    "article": "der"
  },
  {
    "noun": "Auenwald",
    "article": "der"
  },
  {
    "noun": "Auerhahn",
    "article": "der"
  },
  {
    "noun": "Auerochse",
    "article": "der"
  },
  {
    "noun": "Aufbau",
    "article": "der"
  },
  {
    "noun": "Aufbewahrungsort",
    "article": "der"
  },
  {
    "noun": "Aufbruch",
    "article": "der"
  },
  {
    "noun": "Aufdruck",
    "article": "der"
  },
  {
    "noun": "Aufenthalt",
    "article": "der"
  },
  {
    "noun": "Aufenthaltsort",
    "article": "der"
  },
  {
    "noun": "Aufenthaltsraum",
    "article": "der"
  },
  {
    "noun": "Auffahrunfall",
    "article": "der"
  },
  {
    "noun": "Aufforderungscharakter",
    "article": "der"
  },
  {
    "noun": "Aufgabenbereich",
    "article": "der"
  },
  {
    "noun": "Aufgang",
    "article": "der"
  },
  {
    "noun": "Aufgeber",
    "article": "der"
  },
  {
    "noun": "Aufguss",
    "article": "der"
  },
  {
    "noun": "Aufhebungsvertrag",
    "article": "der"
  },
  {
    "noun": "Aufkauf",
    "article": "der"
  },
  {
    "noun": "Aufkaufhandel",
    "article": "der"
  },
  {
    "noun": "Aufkleber",
    "article": "der"
  },
  {
    "noun": "Auflader",
    "article": "der"
  },
  {
    "noun": "Auflagedruck",
    "article": "der"
  },
  {
    "noun": "Auflauf",
    "article": "der"
  },
  {
    "noun": "Auflieferer",
    "article": "der"
  },
  {
    "noun": "Auflieger",
    "article": "der"
  },
  {
    "noun": "Aufmacher",
    "article": "der"
  },
  {
    "noun": "Aufmarsch",
    "article": "der"
  },
  {
    "noun": "Aufnehmer",
    "article": "der"
  },
  {
    "noun": "Aufpasser",
    "article": "der"
  },
  {
    "noun": "Aufprall",
    "article": "der"
  },
  {
    "noun": "Aufpreis",
    "article": "der"
  },
  {
    "noun": "Aufputz",
    "article": "der"
  },
  {
    "noun": "Aufriss",
    "article": "der"
  },
  {
    "noun": "Aufruf",
    "article": "der"
  },
  {
    "noun": "Aufruhr",
    "article": "der"
  },
  {
    "noun": "Aufsatz",
    "article": "der"
  },
  {
    "noun": "Aufschlag",
    "article": "der"
  },
  {
    "noun": "Aufschlagfehler",
    "article": "der"
  },
  {
    "noun": "Aufschluss",
    "article": "der"
  },
  {
    "noun": "Aufschneider",
    "article": "der"
  },
  {
    "noun": "Aufschnitt",
    "article": "der"
  },
  {
    "noun": "Aufschrei",
    "article": "der"
  },
  {
    "noun": "Aufschub",
    "article": "der"
  },
  {
    "noun": "Aufschwung",
    "article": "der"
  },
  {
    "noun": "Aufseher",
    "article": "der"
  },
  {
    "noun": "Aufsetzer",
    "article": "der"
  },
  {
    "noun": "Aufsichtsrat",
    "article": "der"
  },
  {
    "noun": "Aufsprung",
    "article": "der"
  },
  {
    "noun": "Aufstand",
    "article": "der"
  },
  {
    "noun": "Aufstau",
    "article": "der"
  },
  {
    "noun": "Aufsteiger",
    "article": "der"
  },
  {
    "noun": "Aufstieg",
    "article": "der"
  },
  {
    "noun": "Aufstrich",
    "article": "der"
  },
  {
    "noun": "Auftakt",
    "article": "der"
  },
  {
    "noun": "Auftrag",
    "article": "der"
  },
  {
    "noun": "Auftraggeber",
    "article": "der"
  },
  {
    "noun": "Auftragseingang",
    "article": "der"
  },
  {
    "noun": "Auftrieb",
    "article": "der"
  },
  {
    "noun": "Auftritt",
    "article": "der"
  },
  {
    "noun": "Aufwand",
    "article": "der"
  },
  {
    "noun": "Aufwasch",
    "article": "der"
  },
  {
    "noun": "Aufwiegler",
    "article": "der"
  },
  {
    "noun": "Aufwind",
    "article": "der"
  },
  {
    "noun": "Aufwuchs",
    "article": "der"
  },
  {
    "noun": "Aufwurf",
    "article": "der"
  },
  {
    "noun": "Aufzug",
    "article": "der"
  },
  {
    "noun": "Aufzugschacht",
    "article": "der"
  },
  {
    "noun": "Aufzugsschacht",
    "article": "der"
  },
  {
    "noun": "Augapfel",
    "article": "der"
  },
  {
    "noun": "Augenabstand",
    "article": "der"
  },
  {
    "noun": "Augenarzt",
    "article": "der"
  },
  {
    "noun": "Augenaufschlag",
    "article": "der"
  },
  {
    "noun": "Augenblick",
    "article": "der"
  },
  {
    "noun": "Augenblickserfolg",
    "article": "der"
  },
  {
    "noun": "Augenbrauenstift",
    "article": "der"
  },
  {
    "noun": "Augeninnendruck",
    "article": "der"
  },
  {
    "noun": "Augenmuskel",
    "article": "der"
  },
  {
    "noun": "Augenoptiker",
    "article": "der"
  },
  {
    "noun": "Augenrand",
    "article": "der"
  },
  {
    "noun": "Augenschein",
    "article": "der"
  },
  {
    "noun": "Augenschirm",
    "article": "der"
  },
  {
    "noun": "Augenschmaus",
    "article": "der"
  },
  {
    "noun": "Augenspiegel",
    "article": "der"
  },
  {
    "noun": "Augenstern",
    "article": "der"
  },
  {
    "noun": "Augentrost",
    "article": "der"
  },
  {
    "noun": "Augenwinkel",
    "article": "der"
  },
  {
    "noun": "Augenzahn",
    "article": "der"
  },
  {
    "noun": "Augenzeuge",
    "article": "der"
  },
  {
    "noun": "Augenzeugenbericht",
    "article": "der"
  },
  {
    "noun": "Augit",
    "article": "der"
  },
  {
    "noun": "August",
    "article": "der"
  },
  {
    "noun": "Auktionator",
    "article": "der"
  },
  {
    "noun": "Auktionskatalog",
    "article": "der"
  },
  {
    "noun": "Ausbau",
    "article": "der"
  },
  {
    "noun": "Ausbeuter",
    "article": "der"
  },
  {
    "noun": "Ausbilder",
    "article": "der"
  },
  {
    "noun": "Ausbildungsberuf",
    "article": "der"
  },
  {
    "noun": "Ausbildungsgang",
    "article": "der"
  },
  {
    "noun": "Ausbildungskurs",
    "article": "der"
  },
  {
    "noun": "Ausbildungslehrgang",
    "article": "der"
  },
  {
    "noun": "Ausbildungsstand",
    "article": "der"
  },
  {
    "noun": "Ausbildungsvertrag",
    "article": "der"
  },
  {
    "noun": "Ausblick",
    "article": "der"
  },
  {
    "noun": "Ausbruch",
    "article": "der"
  },
  {
    "noun": "Ausbund",
    "article": "der"
  },
  {
    "noun": "Ausdehnungskoeffizient",
    "article": "der"
  },
  {
    "noun": "Ausdruck",
    "article": "der"
  },
  {
    "noun": "Ausdruckstanz",
    "article": "der"
  },
  {
    "noun": "Ausfall",
    "article": "der"
  },
  {
    "noun": "Ausfallschritt",
    "article": "der"
  },
  {
    "noun": "Ausfallswinkel",
    "article": "der"
  },
  {
    "noun": "Ausfeger",
    "article": "der"
  },
  {
    "noun": "Ausflug",
    "article": "der"
  },
  {
    "noun": "Ausflugsort",
    "article": "der"
  },
  {
    "noun": "Ausfluss",
    "article": "der"
  },
  {
    "noun": "Ausfrager",
    "article": "der"
  },
  {
    "noun": "Ausfuhrartikel",
    "article": "der"
  },
  {
    "noun": "Ausfuhrkredit",
    "article": "der"
  },
  {
    "noun": "Ausfuhrzoll",
    "article": "der"
  },
  {
    "noun": "Ausgabekurs",
    "article": "der"
  },
  {
    "noun": "Ausgabenbeleg",
    "article": "der"
  },
  {
    "noun": "Ausgabepreis",
    "article": "der"
  },
  {
    "noun": "Ausgabewert",
    "article": "der"
  },
  {
    "noun": "Ausgang",
    "article": "der"
  },
  {
    "noun": "Ausgangspunkt",
    "article": "der"
  },
  {
    "noun": "Ausgangsstoff",
    "article": "der"
  },
  {
    "noun": "Ausgangswert",
    "article": "der"
  },
  {
    "noun": "Ausgangszustand",
    "article": "der"
  },
  {
    "noun": "Ausgleich",
    "article": "der"
  },
  {
    "noun": "Ausgleicher",
    "article": "der"
  },
  {
    "noun": "Ausgleichsfonds",
    "article": "der"
  },
  {
    "noun": "Ausgleichssport",
    "article": "der"
  },
  {
    "noun": "Ausgleichstreffer",
    "article": "der"
  },
  {
    "noun": "Ausgrabungsfund",
    "article": "der"
  },
  {
    "noun": "Ausguck",
    "article": "der"
  },
  {
    "noun": "Ausguss",
    "article": "der"
  },
  {
    "noun": "Aushang",
    "article": "der"
  },
  {
    "noun": "Ausheber",
    "article": "der"
  },
  {
    "noun": "Aushelfer",
    "article": "der"
  },
  {
    "noun": "Aushub",
    "article": "der"
  },
  {
    "noun": "Ausklang",
    "article": "der"
  },
  {
    "noun": "Auskundschafter",
    "article": "der"
  },
  {
    "noun": "Auskunftsdienst",
    "article": "der"
  },
  {
    "noun": "Auskunftsschalter",
    "article": "der"
  },
  {
    "noun": "Ausladebahnhof",
    "article": "der"
  },
  {
    "noun": "Auslass",
    "article": "der"
  },
  {
    "noun": "Auslassungssatz",
    "article": "der"
  },
  {
    "noun": "Auslauf",
    "article": "der"
  },
  {
    "noun": "Auslaut",
    "article": "der"
  },
  {
    "noun": "Ausleger",
    "article": "der"
  },
  {
    "noun": "Auslieferungsantrag",
    "article": "der"
  },
  {
    "noun": "Auslieferungsvertrag",
    "article": "der"
  },
  {
    "noun": "Auslug",
    "article": "der"
  },
  {
    "noun": "Ausnahmefall",
    "article": "der"
  },
  {
    "noun": "Ausnahmezustand",
    "article": "der"
  },
  {
    "noun": "Auspuff",
    "article": "der"
  },
  {
    "noun": "Auspufftopf",
    "article": "der"
  },
  {
    "noun": "Ausritt",
    "article": "der"
  },
  {
    "noun": "Ausruf",
    "article": "der"
  },
  {
    "noun": "Ausrufer",
    "article": "der"
  },
  {
    "noun": "Ausrufesatz",
    "article": "der"
  },
  {
    "noun": "Ausrutscher",
    "article": "der"
  },
  {
    "noun": "Aussagesatz",
    "article": "der"
  },
  {
    "noun": "Aussagewert",
    "article": "der"
  },
  {
    "noun": "Aussatz",
    "article": "der"
  },
  {
    "noun": "Ausschalter",
    "article": "der"
  },
  {
    "noun": "Ausschank",
    "article": "der"
  },
  {
    "noun": "Ausscheidungskampf",
    "article": "der"
  },
  {
    "noun": "Ausschlag",
    "article": "der"
  },
  {
    "noun": "Ausschlupf",
    "article": "der"
  },
  {
    "noun": "Ausschluss",
    "article": "der"
  },
  {
    "noun": "Ausschnitt",
    "article": "der"
  },
  {
    "noun": "Ausschuss",
    "article": "der"
  },
  {
    "noun": "Aussetzer",
    "article": "der"
  },
  {
    "noun": "Aussichtspunkt",
    "article": "der"
  },
  {
    "noun": "Aussichtsturm",
    "article": "der"
  },
  {
    "noun": "Aussiedler",
    "article": "der"
  },
  {
    "noun": "Ausspann",
    "article": "der"
  },
  {
    "noun": "Ausspruch",
    "article": "der"
  },
  {
    "noun": "Ausstand",
    "article": "der"
  },
  {
    "noun": "Aussteiger",
    "article": "der"
  },
  {
    "noun": "Aussteller",
    "article": "der"
  },
  {
    "noun": "Ausstellungskatalog",
    "article": "der"
  },
  {
    "noun": "Ausstellungsort",
    "article": "der"
  },
  {
    "noun": "Ausstellungsraum",
    "article": "der"
  },
  {
    "noun": "Ausstellungsstand",
    "article": "der"
  },
  {
    "noun": "Ausstellungstag",
    "article": "der"
  },
  {
    "noun": "Ausstieg",
    "article": "der"
  },
  {
    "noun": "Austausch",
    "article": "der"
  },
  {
    "noun": "Austauschmotor",
    "article": "der"
  },
  {
    "noun": "Austauschstudent",
    "article": "der"
  },
  {
    "noun": "Austenit",
    "article": "der"
  },
  {
    "noun": "Austernfischer",
    "article": "der"
  },
  {
    "noun": "Austernpilz",
    "article": "der"
  },
  {
    "noun": "Austernseitling",
    "article": "der"
  },
  {
    "noun": "Austrag",
    "article": "der"
  },
  {
    "noun": "Austragungsort",
    "article": "der"
  },
  {
    "noun": "Australier",
    "article": "der"
  },
  {
    "noun": "Austreiber",
    "article": "der"
  },
  {
    "noun": "Austriazismus",
    "article": "der"
  },
  {
    "noun": "Austritt",
    "article": "der"
  },
  {
    "noun": "Ausverkauf",
    "article": "der"
  },
  {
    "noun": "Auswanderer",
    "article": "der"
  },
  {
    "noun": "Auswechselspieler",
    "article": "der"
  },
  {
    "noun": "Ausweg",
    "article": "der"
  },
  {
    "noun": "Ausweis",
    "article": "der"
  },
  {
    "noun": "Auswerfer",
    "article": "der"
  },
  {
    "noun": "Auswuchs",
    "article": "der"
  },
  {
    "noun": "Auswurf",
    "article": "der"
  },
  {
    "noun": "Ausziehtisch",
    "article": "der"
  },
  {
    "noun": "Auszug",
    "article": "der"
  },
  {
    "noun": "Autismus",
    "article": "der"
  },
  {
    "noun": "Autobahnanschluss",
    "article": "der"
  },
  {
    "noun": "Autobahnzubringer",
    "article": "der"
  },
  {
    "noun": "Autobiograf",
    "article": "der"
  },
  {
    "noun": "Autobus",
    "article": "der"
  },
  {
    "noun": "Autobusbahnhof",
    "article": "der"
  },
  {
    "noun": "Autocorso",
    "article": "der"
  },
  {
    "noun": "Autofahrer",
    "article": "der"
  },
  {
    "noun": "Autofokus",
    "article": "der"
  },
  {
    "noun": "Autofriedhof",
    "article": "der"
  },
  {
    "noun": "Autoklav",
    "article": "der"
  },
  {
    "noun": "Autoknacker",
    "article": "der"
  },
  {
    "noun": "Autokran",
    "article": "der"
  },
  {
    "noun": "Autokrat",
    "article": "der"
  },
  {
    "noun": "Autolenker",
    "article": "der"
  },
  {
    "noun": "Automarder",
    "article": "der"
  },
  {
    "noun": "Automat",
    "article": "der"
  },
  {
    "noun": "Automatismus",
    "article": "der"
  },
  {
    "noun": "Automechaniker",
    "article": "der"
  },
  {
    "noun": "Automobilclub",
    "article": "der"
  },
  {
    "noun": "Automobilist",
    "article": "der"
  },
  {
    "noun": "Autonomist",
    "article": "der"
  },
  {
    "noun": "Autopilot",
    "article": "der"
  },
  {
    "noun": "Autor",
    "article": "der"
  },
  {
    "noun": "Autoreifen",
    "article": "der"
  },
  {
    "noun": "Autoreisezug",
    "article": "der"
  },
  {
    "noun": "Autorenkatalog",
    "article": "der"
  },
  {
    "noun": "Autoritarismus",
    "article": "der"
  },
  {
    "noun": "Autoschalter",
    "article": "der"
  },
  {
    "noun": "Autoschlosser",
    "article": "der"
  },
  {
    "noun": "Autounfall",
    "article": "der"
  },
  {
    "noun": "Autoverkehr",
    "article": "der"
  },
  {
    "noun": "Autoverleih",
    "article": "der"
  },
  {
    "noun": "Autoverwerter",
    "article": "der"
  },
  {
    "noun": "Avalist",
    "article": "der"
  },
  {
    "noun": "Avalkredit",
    "article": "der"
  },
  {
    "noun": "Avantgardismus",
    "article": "der"
  },
  {
    "noun": "Avantgardist",
    "article": "der"
  },
  {
    "noun": "Avatar",
    "article": "der"
  },
  {
    "noun": "Aventurin",
    "article": "der"
  },
  {
    "noun": "Avers",
    "article": "der"
  },
  {
    "noun": "Aviso",
    "article": "der"
  },
  {
    "noun": "Aware",
    "article": "der"
  },
  {
    "noun": "Axel",
    "article": "der"
  },
  {
    "noun": "Axinit",
    "article": "der"
  },
  {
    "noun": "Azteke",
    "article": "der"
  },
  {
    "noun": "Azur",
    "article": "der"
  },
  {
    "noun": "Azurit",
    "article": "der"
  },
  {
    "noun": "Sahnequark",
    "article": "der"
  },
  {
    "noun": "Saibling",
    "article": "der"
  },
  {
    "noun": "Saisonarbeiter",
    "article": "der"
  },
  {
    "noun": "Saisonausverkauf",
    "article": "der"
  },
  {
    "noun": "Saisonbetrieb",
    "article": "der"
  },
  {
    "noun": "Saisonindex",
    "article": "der"
  },
  {
    "noun": "Saisonkredit",
    "article": "der"
  },
  {
    "noun": "Saisonschlussverkauf",
    "article": "der"
  },
  {
    "noun": "Saitenhalter",
    "article": "der"
  },
  {
    "noun": "Saitenklang",
    "article": "der"
  },
  {
    "noun": "Sake",
    "article": "der"
  },
  {
    "noun": "Sakralbau",
    "article": "der"
  },
  {
    "noun": "Sakristan",
    "article": "der"
  },
  {
    "noun": "Salamander",
    "article": "der"
  },
  {
    "noun": "Salami",
    "article": "der"
  },
  {
    "noun": "Salat",
    "article": "der"
  },
  {
    "noun": "Salatkopf",
    "article": "der"
  },
  {
    "noun": "Salbader",
    "article": "der"
  },
  {
    "noun": "Salbeitee",
    "article": "der"
  },
  {
    "noun": "Saldo",
    "article": "der"
  },
  {
    "noun": "Saldovortrag",
    "article": "der"
  },
  {
    "noun": "Salinenkrebs",
    "article": "der"
  },
  {
    "noun": "Salm",
    "article": "der"
  },
  {
    "noun": "Salmiak",
    "article": "der"
  },
  {
    "noun": "Salmiakgeist",
    "article": "der"
  },
  {
    "noun": "Salmler",
    "article": "der"
  },
  {
    "noun": "Salmonide",
    "article": "der"
  },
  {
    "noun": "Salon",
    "article": "der"
  },
  {
    "noun": "Salonwagen",
    "article": "der"
  },
  {
    "noun": "Salpeter",
    "article": "der"
  },
  {
    "noun": "Salto",
    "article": "der"
  },
  {
    "noun": "Salut",
    "article": "der"
  },
  {
    "noun": "Salutschuss",
    "article": "der"
  },
  {
    "noun": "Salvadorianer",
    "article": "der"
  },
  {
    "noun": "Salzboden",
    "article": "der"
  },
  {
    "noun": "Salzbrunnen",
    "article": "der"
  },
  {
    "noun": "Salzgarten",
    "article": "der"
  },
  {
    "noun": "Salzgehalt",
    "article": "der"
  },
  {
    "noun": "Salzhering",
    "article": "der"
  },
  {
    "noun": "Salzsee",
    "article": "der"
  },
  {
    "noun": "Salzsieder",
    "article": "der"
  },
  {
    "noun": "Salzstreuer",
    "article": "der"
  },
  {
    "noun": "Salzteig",
    "article": "der"
  },
  {
    "noun": "Samariter",
    "article": "der"
  },
  {
    "noun": "Sambesi",
    "article": "der"
  },
  {
    "noun": "Sambier",
    "article": "der"
  },
  {
    "noun": "Same",
    "article": "der"
  },
  {
    "noun": "Samenerguss",
    "article": "der"
  },
  {
    "noun": "Samenfaden",
    "article": "der"
  },
  {
    "noun": "Samenkern",
    "article": "der"
  },
  {
    "noun": "Samenleiter",
    "article": "der"
  },
  {
    "noun": "Samenstrang",
    "article": "der"
  },
  {
    "noun": "Sammelauftrag",
    "article": "der"
  },
  {
    "noun": "Sammelband",
    "article": "der"
  },
  {
    "noun": "Sammelbegriff",
    "article": "der"
  },
  {
    "noun": "Sammelfahrschein",
    "article": "der"
  },
  {
    "noun": "Sammelname",
    "article": "der"
  },
  {
    "noun": "Sammelplatz",
    "article": "der"
  },
  {
    "noun": "Sammelpunkt",
    "article": "der"
  },
  {
    "noun": "Sammeltransport",
    "article": "der"
  },
  {
    "noun": "Sammeltrieb",
    "article": "der"
  },
  {
    "noun": "Sammler",
    "article": "der"
  },
  {
    "noun": "Samowar",
    "article": "der"
  },
  {
    "noun": "Sampan",
    "article": "der"
  },
  {
    "noun": "Sampler",
    "article": "der"
  },
  {
    "noun": "Samstag",
    "article": "der"
  },
  {
    "noun": "Samt",
    "article": "der"
  },
  {
    "noun": "Samthandschuh",
    "article": "der"
  },
  {
    "noun": "Samum",
    "article": "der"
  },
  {
    "noun": "Samurai",
    "article": "der"
  },
  {
    "noun": "Sand",
    "article": "der"
  },
  {
    "noun": "Sandboden",
    "article": "der"
  },
  {
    "noun": "Sanddorn",
    "article": "der"
  },
  {
    "noun": "Sandfang",
    "article": "der"
  },
  {
    "noun": "Sandfloh",
    "article": "der"
  },
  {
    "noun": "Sandguss",
    "article": "der"
  },
  {
    "noun": "Sandhase",
    "article": "der"
  },
  {
    "noun": "Sandhaufen",
    "article": "der"
  },
  {
    "noun": "Sandkasten",
    "article": "der"
  },
  {
    "noun": "Sandmann",
    "article": "der"
  },
  {
    "noun": "Sandsack",
    "article": "der"
  },
  {
    "noun": "Sandstein",
    "article": "der"
  },
  {
    "noun": "Sandstrand",
    "article": "der"
  },
  {
    "noun": "Sandsturm",
    "article": "der"
  },
  {
    "noun": "Sandweg",
    "article": "der"
  },
  {
    "noun": "Sang",
    "article": "der"
  },
  {
    "noun": "Sanguiniker",
    "article": "der"
  },
  {
    "noun": "Sanierungsplan",
    "article": "der"
  },
  {
    "noun": "Sanskritist",
    "article": "der"
  },
  {
    "noun": "Saphir",
    "article": "der"
  },
  {
    "noun": "Saprophyt",
    "article": "der"
  },
  {
    "noun": "Sarde",
    "article": "der"
  },
  {
    "noun": "Sardinier",
    "article": "der"
  },
  {
    "noun": "Sardonyx",
    "article": "der"
  },
  {
    "noun": "Sarg",
    "article": "der"
  },
  {
    "noun": "Sargdeckel",
    "article": "der"
  },
  {
    "noun": "Sargnagel",
    "article": "der"
  },
  {
    "noun": "Sari",
    "article": "der"
  },
  {
    "noun": "Sarkasmus",
    "article": "der"
  },
  {
    "noun": "Sarkophag",
    "article": "der"
  },
  {
    "noun": "Sarong",
    "article": "der"
  },
  {
    "noun": "Sass",
    "article": "der"
  },
  {
    "noun": "Sassafras",
    "article": "der"
  },
  {
    "noun": "Satan",
    "article": "der"
  },
  {
    "noun": "Satang",
    "article": "der"
  },
  {
    "noun": "Satanismus",
    "article": "der"
  },
  {
    "noun": "Satansbraten",
    "article": "der"
  },
  {
    "noun": "Satellit",
    "article": "der"
  },
  {
    "noun": "Satellitenempfang",
    "article": "der"
  },
  {
    "noun": "Satellitenstaat",
    "article": "der"
  },
  {
    "noun": "Satin",
    "article": "der"
  },
  {
    "noun": "Satiriker",
    "article": "der"
  },
  {
    "noun": "Satrap",
    "article": "der"
  },
  {
    "noun": "Sattdampf",
    "article": "der"
  },
  {
    "noun": "Sattel",
    "article": "der"
  },
  {
    "noun": "Sattelauflieger",
    "article": "der"
  },
  {
    "noun": "Satteldruck",
    "article": "der"
  },
  {
    "noun": "Sattelgurt",
    "article": "der"
  },
  {
    "noun": "Sattelkloben",
    "article": "der"
  },
  {
    "noun": "Sattelknopf",
    "article": "der"
  },
  {
    "noun": "Sattelpunkt",
    "article": "der"
  },
  {
    "noun": "Sattelschlepper",
    "article": "der"
  },
  {
    "noun": "Sattelzug",
    "article": "der"
  },
  {
    "noun": "Sattler",
    "article": "der"
  },
  {
    "noun": "Satyr",
    "article": "der"
  },
  {
    "noun": "Satz",
    "article": "der"
  },
  {
    "noun": "Satzakzent",
    "article": "der"
  },
  {
    "noun": "Satzball",
    "article": "der"
  },
  {
    "noun": "Satzbau",
    "article": "der"
  },
  {
    "noun": "Satzbruch",
    "article": "der"
  },
  {
    "noun": "Satzfehler",
    "article": "der"
  },
  {
    "noun": "Satzgegenstand",
    "article": "der"
  },
  {
    "noun": "Satzspiegel",
    "article": "der"
  },
  {
    "noun": "Satzteil",
    "article": "der"
  },
  {
    "noun": "Sauberkeitsfimmel",
    "article": "der"
  },
  {
    "noun": "Saubermann",
    "article": "der"
  },
  {
    "noun": "Sauerampfer",
    "article": "der"
  },
  {
    "noun": "Sauerbraten",
    "article": "der"
  },
  {
    "noun": "Sauerbrunnen",
    "article": "der"
  },
  {
    "noun": "Sauerdorn",
    "article": "der"
  },
  {
    "noun": "Sauerrahm",
    "article": "der"
  },
  {
    "noun": "Sauerstoff",
    "article": "der"
  },
  {
    "noun": "Sauerstoffapparat",
    "article": "der"
  },
  {
    "noun": "Sauerstoffgehalt",
    "article": "der"
  },
  {
    "noun": "Sauerstoffmangel",
    "article": "der"
  },
  {
    "noun": "Sauerteig",
    "article": "der"
  },
  {
    "noun": "Saufbold",
    "article": "der"
  },
  {
    "noun": "Saufbruder",
    "article": "der"
  },
  {
    "noun": "Saufkumpan",
    "article": "der"
  },
  {
    "noun": "Saugbagger",
    "article": "der"
  },
  {
    "noun": "Sauger",
    "article": "der"
  },
  {
    "noun": "Saugheber",
    "article": "der"
  },
  {
    "noun": "Saugkopf",
    "article": "der"
  },
  {
    "noun": "Saugkorb",
    "article": "der"
  },
  {
    "noun": "Saugmotor",
    "article": "der"
  },
  {
    "noun": "Saugnapf",
    "article": "der"
  },
  {
    "noun": "Saugreflex",
    "article": "der"
  },
  {
    "noun": "Saugwurm",
    "article": "der"
  },
  {
    "noun": "Sauhund",
    "article": "der"
  },
  {
    "noun": "Saukerl",
    "article": "der"
  },
  {
    "noun": "Saum",
    "article": "der"
  },
  {
    "noun": "Saumpfad",
    "article": "der"
  },
  {
    "noun": "Saumsattel",
    "article": "der"
  },
  {
    "noun": "Saurier",
    "article": "der"
  },
  {
    "noun": "Sausewind",
    "article": "der"
  },
  {
    "noun": "Saustall",
    "article": "der"
  },
  {
    "noun": "Savoyerkohl",
    "article": "der"
  },
  {
    "noun": "Saxofonist",
    "article": "der"
  },
  {
    "noun": "Sbirre",
    "article": "der"
  },
  {
    "noun": "Scampo",
    "article": "der"
  },
  {
    "noun": "Scanner",
    "article": "der"
  },
  {
    "noun": "Schaber",
    "article": "der"
  },
  {
    "noun": "Schabernack",
    "article": "der"
  },
  {
    "noun": "Schabrackenschakal",
    "article": "der"
  },
  {
    "noun": "Schabrackentapir",
    "article": "der"
  },
  {
    "noun": "Schacher",
    "article": "der"
  },
  {
    "noun": "Schacherer",
    "article": "der"
  },
  {
    "noun": "Schachmeister",
    "article": "der"
  },
  {
    "noun": "Schachspieler",
    "article": "der"
  },
  {
    "noun": "Schacht",
    "article": "der"
  },
  {
    "noun": "Schachtdeckel",
    "article": "der"
  },
  {
    "noun": "Schachtelhalm",
    "article": "der"
  },
  {
    "noun": "Schachtelsatz",
    "article": "der"
  },
  {
    "noun": "Schachtofen",
    "article": "der"
  },
  {
    "noun": "Schachzug",
    "article": "der"
  },
  {
    "noun": "Schadenbericht",
    "article": "der"
  },
  {
    "noun": "Schadenersatz",
    "article": "der"
  },
  {
    "noun": "Schadenersatzanspruch",
    "article": "der"
  },
  {
    "noun": "Schadenfall",
    "article": "der"
  },
  {
    "noun": "Schadenfreiheitsrabatt",
    "article": "der"
  },
  {
    "noun": "Schadennachweis",
    "article": "der"
  },
  {
    "noun": "Schadensbericht",
    "article": "der"
  },
  {
    "noun": "Schadensersatz",
    "article": "der"
  },
  {
    "noun": "Schadensersatzanspruch",
    "article": "der"
  },
  {
    "noun": "Schadensfall",
    "article": "der"
  },
  {
    "noun": "Schadensnachweis",
    "article": "der"
  },
  {
    "noun": "Schadstoff",
    "article": "der"
  },
  {
    "noun": "Schafbock",
    "article": "der"
  },
  {
    "noun": "Schaffensdrang",
    "article": "der"
  },
  {
    "noun": "Schaffer",
    "article": "der"
  },
  {
    "noun": "Schaffner",
    "article": "der"
  },
  {
    "noun": "Schafhirt",
    "article": "der"
  },
  {
    "noun": "Schafpelz",
    "article": "der"
  },
  {
    "noun": "Schafskopf",
    "article": "der"
  },
  {
    "noun": "Schafspelz",
    "article": "der"
  },
  {
    "noun": "Schafstall",
    "article": "der"
  },
  {
    "noun": "Schaft",
    "article": "der"
  },
  {
    "noun": "Schaftstiefel",
    "article": "der"
  },
  {
    "noun": "Schah",
    "article": "der"
  },
  {
    "noun": "Schakal",
    "article": "der"
  },
  {
    "noun": "Schalenbau",
    "article": "der"
  },
  {
    "noun": "Schalensitz",
    "article": "der"
  },
  {
    "noun": "Schalk",
    "article": "der"
  },
  {
    "noun": "Schall",
    "article": "der"
  },
  {
    "noun": "Schallbecher",
    "article": "der"
  },
  {
    "noun": "Schalldeckel",
    "article": "der"
  },
  {
    "noun": "Schalldruck",
    "article": "der"
  },
  {
    "noun": "Schallkasten",
    "article": "der"
  },
  {
    "noun": "Schallplattenspieler",
    "article": "der"
  },
  {
    "noun": "Schallschatten",
    "article": "der"
  },
  {
    "noun": "Schallschutz",
    "article": "der"
  },
  {
    "noun": "Schalltrichter",
    "article": "der"
  },
  {
    "noun": "Schallwandler",
    "article": "der"
  },
  {
    "noun": "Schalter",
    "article": "der"
  },
  {
    "noun": "Schalterdienst",
    "article": "der"
  },
  {
    "noun": "Schalthebel",
    "article": "der"
  },
  {
    "noun": "Schaltkasten",
    "article": "der"
  },
  {
    "noun": "Schaltkreis",
    "article": "der"
  },
  {
    "noun": "Schaltplan",
    "article": "der"
  },
  {
    "noun": "Schaltsatz",
    "article": "der"
  },
  {
    "noun": "Schaltschrank",
    "article": "der"
  },
  {
    "noun": "Schalttag",
    "article": "der"
  },
  {
    "noun": "Schalttisch",
    "article": "der"
  },
  {
    "noun": "Schaltweg",
    "article": "der"
  },
  {
    "noun": "Schamane",
    "article": "der"
  },
  {
    "noun": "Schamanismus",
    "article": "der"
  },
  {
    "noun": "Schamott",
    "article": "der"
  },
  {
    "noun": "Schamottestein",
    "article": "der"
  },
  {
    "noun": "Schampus",
    "article": "der"
  },
  {
    "noun": "Schandbube",
    "article": "der"
  },
  {
    "noun": "Schandfleck",
    "article": "der"
  },
  {
    "noun": "Schandpfahl",
    "article": "der"
  },
  {
    "noun": "Schani",
    "article": "der"
  },
  {
    "noun": "Schank",
    "article": "der"
  },
  {
    "noun": "Schankbetrieb",
    "article": "der"
  },
  {
    "noun": "Schanker",
    "article": "der"
  },
  {
    "noun": "Schankkellner",
    "article": "der"
  },
  {
    "noun": "Schankraum",
    "article": "der"
  },
  {
    "noun": "Schanktisch",
    "article": "der"
  },
  {
    "noun": "Schankwirt",
    "article": "der"
  },
  {
    "noun": "Schanzenbau",
    "article": "der"
  },
  {
    "noun": "Schanzenrekord",
    "article": "der"
  },
  {
    "noun": "Scharfblick",
    "article": "der"
  },
  {
    "noun": "Scharfmacher",
    "article": "der"
  },
  {
    "noun": "Scharfrichter",
    "article": "der"
  },
  {
    "noun": "Scharfsinn",
    "article": "der"
  },
  {
    "noun": "Scharlatan",
    "article": "der"
  },
  {
    "noun": "Scharm",
    "article": "der"
  },
  {
    "noun": "Scharwerker",
    "article": "der"
  },
  {
    "noun": "Schaschlik",
    "article": "der"
  },
  {
    "noun": "Schattenriss",
    "article": "der"
  },
  {
    "noun": "Schatz",
    "article": "der"
  },
  {
    "noun": "Schatzfund",
    "article": "der"
  },
  {
    "noun": "Schatzkanzler",
    "article": "der"
  },
  {
    "noun": "Schatzmeister",
    "article": "der"
  },
  {
    "noun": "Schatzsucher",
    "article": "der"
  },
  {
    "noun": "Schatzwechsel",
    "article": "der"
  },
  {
    "noun": "Schaub",
    "article": "der"
  },
  {
    "noun": "Schaubudenbesitzer",
    "article": "der"
  },
  {
    "noun": "Schauder",
    "article": "der"
  },
  {
    "noun": "Schauermann",
    "article": "der"
  },
  {
    "noun": "Schauerroman",
    "article": "der"
  },
  {
    "noun": "Schaufelbagger",
    "article": "der"
  },
  {
    "noun": "Schaufellader",
    "article": "der"
  },
  {
    "noun": "Schaufelradbagger",
    "article": "der"
  },
  {
    "noun": "Schaufelraddampfer",
    "article": "der"
  },
  {
    "noun": "Schaufensterbummel",
    "article": "der"
  },
  {
    "noun": "Schaufensterdekorateur",
    "article": "der"
  },
  {
    "noun": "Schaufenstereinbruch",
    "article": "der"
  },
  {
    "noun": "Schaukampf",
    "article": "der"
  },
  {
    "noun": "Schaukasten",
    "article": "der"
  },
  {
    "noun": "Schaukelstuhl",
    "article": "der"
  },
  {
    "noun": "Schaum",
    "article": "der"
  },
  {
    "noun": "Schaumbeton",
    "article": "der"
  },
  {
    "noun": "Schaumgummi",
    "article": "der"
  },
  {
    "noun": "Schaumstoff",
    "article": "der"
  },
  {
    "noun": "Schaumwein",
    "article": "der"
  },
  {
    "noun": "Schauplatz",
    "article": "der"
  },
  {
    "noun": "Schauprozess",
    "article": "der"
  },
  {
    "noun": "Schauspieldichter",
    "article": "der"
  },
  {
    "noun": "Schauspieler",
    "article": "der"
  },
  {
    "noun": "Schauspielunterricht",
    "article": "der"
  },
  {
    "noun": "Schausteller",
    "article": "der"
  },
  {
    "noun": "Schautanz",
    "article": "der"
  },
  {
    "noun": "Scheck",
    "article": "der"
  },
  {
    "noun": "Scheckbetrug",
    "article": "der"
  },
  {
    "noun": "Scheckverkehr",
    "article": "der"
  },
  {
    "noun": "Scheffel",
    "article": "der"
  },
  {
    "noun": "Scheibenhonig",
    "article": "der"
  },
  {
    "noun": "Scheibenwischer",
    "article": "der"
  },
  {
    "noun": "Scheich",
    "article": "der"
  },
  {
    "noun": "Scheidenausfluss",
    "article": "der"
  },
  {
    "noun": "Scheidenkrampf",
    "article": "der"
  },
  {
    "noun": "Scheideweg",
    "article": "der"
  },
  {
    "noun": "Scheidungsgrund",
    "article": "der"
  },
  {
    "noun": "Scheidungsrichter",
    "article": "der"
  },
  {
    "noun": "Schein",
    "article": "der"
  },
  {
    "noun": "Scheinangriff",
    "article": "der"
  },
  {
    "noun": "Scheinbeweis",
    "article": "der"
  },
  {
    "noun": "Scheinfriede",
    "article": "der"
  },
  {
    "noun": "Scheingesellschafter",
    "article": "der"
  },
  {
    "noun": "Scheingewinn",
    "article": "der"
  },
  {
    "noun": "Scheingrund",
    "article": "der"
  },
  {
    "noun": "Scheinkauf",
    "article": "der"
  },
  {
    "noun": "Scheinverlust",
    "article": "der"
  },
  {
    "noun": "Scheinvertrag",
    "article": "der"
  },
  {
    "noun": "Scheinwerfer",
    "article": "der"
  },
  {
    "noun": "Scheinwiderstand",
    "article": "der"
  },
  {
    "noun": "Scheitel",
    "article": "der"
  },
  {
    "noun": "Scheitelkreis",
    "article": "der"
  },
  {
    "noun": "Scheitelpunkt",
    "article": "der"
  },
  {
    "noun": "Scheitelwert",
    "article": "der"
  },
  {
    "noun": "Scheitelwinkel",
    "article": "der"
  },
  {
    "noun": "Scheiterhaufen",
    "article": "der"
  },
  {
    "noun": "Schellack",
    "article": "der"
  },
  {
    "noun": "Schelladler",
    "article": "der"
  },
  {
    "noun": "Schellenbaum",
    "article": "der"
  },
  {
    "noun": "Schellfisch",
    "article": "der"
  },
  {
    "noun": "Schellhammer",
    "article": "der"
  },
  {
    "noun": "Schelm",
    "article": "der"
  },
  {
    "noun": "Schelmenstreich",
    "article": "der"
  },
  {
    "noun": "Schemabrief",
    "article": "der"
  },
  {
    "noun": "Schematismus",
    "article": "der"
  },
  {
    "noun": "Schemel",
    "article": "der"
  },
  {
    "noun": "Schenk",
    "article": "der"
  },
  {
    "noun": "Schenkel",
    "article": "der"
  },
  {
    "noun": "Schenkelbruch",
    "article": "der"
  },
  {
    "noun": "Schenkelhals",
    "article": "der"
  },
  {
    "noun": "Schenkelknochen",
    "article": "der"
  },
  {
    "noun": "Schenker",
    "article": "der"
  },
  {
    "noun": "Schenkwirt",
    "article": "der"
  },
  {
    "noun": "Scherbenhaufen",
    "article": "der"
  },
  {
    "noun": "Scherenarm",
    "article": "der"
  },
  {
    "noun": "Scherenschlag",
    "article": "der"
  },
  {
    "noun": "Scherenschleifer",
    "article": "der"
  },
  {
    "noun": "Scherenschnabel",
    "article": "der"
  },
  {
    "noun": "Scherenschnitt",
    "article": "der"
  },
  {
    "noun": "Scherensprung",
    "article": "der"
  },
  {
    "noun": "Scherer",
    "article": "der"
  },
  {
    "noun": "Schergang",
    "article": "der"
  },
  {
    "noun": "Scherge",
    "article": "der"
  },
  {
    "noun": "Scherkopf",
    "article": "der"
  },
  {
    "noun": "Scherling",
    "article": "der"
  },
  {
    "noun": "Schersprung",
    "article": "der"
  },
  {
    "noun": "Scherz",
    "article": "der"
  },
  {
    "noun": "Scherzbold",
    "article": "der"
  },
  {
    "noun": "Scherzkeks",
    "article": "der"
  },
  {
    "noun": "Scheuerbesen",
    "article": "der"
  },
  {
    "noun": "Scheuerlappen",
    "article": "der"
  },
  {
    "noun": "Scheuersand",
    "article": "der"
  },
  {
    "noun": "Scheunendrescher",
    "article": "der"
  },
  {
    "noun": "Schi",
    "article": "der"
  },
  {
    "noun": "Schichtarbeiter",
    "article": "der"
  },
  {
    "noun": "Schichtpressstoff",
    "article": "der"
  },
  {
    "noun": "Schichtstoff",
    "article": "der"
  },
  {
    "noun": "Schichtunterricht",
    "article": "der"
  },
  {
    "noun": "Schichtwechsel",
    "article": "der"
  },
  {
    "noun": "Schick",
    "article": "der"
  },
  {
    "noun": "Schickimicki",
    "article": "der"
  },
  {
    "noun": "Schicksalsglaube",
    "article": "der"
  },
  {
    "noun": "Schicksalsschlag",
    "article": "der"
  },
  {
    "noun": "Schieber",
    "article": "der"
  },
  {
    "noun": "Schiebesitz",
    "article": "der"
  },
  {
    "noun": "Schiebewiderstand",
    "article": "der"
  },
  {
    "noun": "Schiedsmann",
    "article": "der"
  },
  {
    "noun": "Schiedsrichter",
    "article": "der"
  },
  {
    "noun": "Schiedsrichterball",
    "article": "der"
  },
  {
    "noun": "Schiedsspruch",
    "article": "der"
  },
  {
    "noun": "Schiedsvertrag",
    "article": "der"
  },
  {
    "noun": "Schiefer",
    "article": "der"
  },
  {
    "noun": "Schieferbruch",
    "article": "der"
  },
  {
    "noun": "Schieferdecker",
    "article": "der"
  },
  {
    "noun": "Schieferstift",
    "article": "der"
  },
  {
    "noun": "Schieferton",
    "article": "der"
  },
  {
    "noun": "Schiefhals",
    "article": "der"
  },
  {
    "noun": "Schiemann",
    "article": "der"
  },
  {
    "noun": "Schienbeinbruch",
    "article": "der"
  },
  {
    "noun": "Schienbeinschoner",
    "article": "der"
  },
  {
    "noun": "Schienenbus",
    "article": "der"
  },
  {
    "noun": "Schienenersatzverkehr",
    "article": "der"
  },
  {
    "noun": "Schienenomnibus",
    "article": "der"
  },
  {
    "noun": "Schienenstrang",
    "article": "der"
  },
  {
    "noun": "Schienenverkehr",
    "article": "der"
  },
  {
    "noun": "Schienenweg",
    "article": "der"
  },
  {
    "noun": "Schierling",
    "article": "der"
  },
  {
    "noun": "Schierlingsbecher",
    "article": "der"
  },
  {
    "noun": "Schiffahrtskaufmann",
    "article": "der"
  },
  {
    "noun": "Schiffbau",
    "article": "der"
  },
  {
    "noun": "Schiffbauer",
    "article": "der"
  },
  {
    "noun": "Schiffbauingenieur",
    "article": "der"
  },
  {
    "noun": "Schiffbruch",
    "article": "der"
  },
  {
    "noun": "Schiffer",
    "article": "der"
  },
  {
    "noun": "Schifferknoten",
    "article": "der"
  },
  {
    "noun": "Schiffsagent",
    "article": "der"
  },
  {
    "noun": "Schiffsarzt",
    "article": "der"
  },
  {
    "noun": "Schiffsbau",
    "article": "der"
  },
  {
    "noun": "Schiffsbauer",
    "article": "der"
  },
  {
    "noun": "Schiffsbohrwurm",
    "article": "der"
  },
  {
    "noun": "Schiffsbrief",
    "article": "der"
  },
  {
    "noun": "Schiffseigner",
    "article": "der"
  },
  {
    "noun": "Schiffsfrachtverkehr",
    "article": "der"
  },
  {
    "noun": "Schiffsjunge",
    "article": "der"
  },
  {
    "noun": "Schiffskoch",
    "article": "der"
  },
  {
    "noun": "Schiffsmakler",
    "article": "der"
  },
  {
    "noun": "Schiffsname",
    "article": "der"
  },
  {
    "noun": "Schiffsraum",
    "article": "der"
  },
  {
    "noun": "Schiffsrumpf",
    "article": "der"
  },
  {
    "noun": "Schiffsverkehr",
    "article": "der"
  },
  {
    "noun": "Schiffszettel",
    "article": "der"
  },
  {
    "noun": "Schiffszoll",
    "article": "der"
  },
  {
    "noun": "Schiffszwieback",
    "article": "der"
  },
  {
    "noun": "Schifter",
    "article": "der"
  },
  {
    "noun": "Schiismus",
    "article": "der"
  },
  {
    "noun": "Schiit",
    "article": "der"
  },
  {
    "noun": "Schildbogen",
    "article": "der"
  },
  {
    "noun": "Schildermaler",
    "article": "der"
  },
  {
    "noun": "Schildknappe",
    "article": "der"
  },
  {
    "noun": "Schilift",
    "article": "der"
  },
  {
    "noun": "Schilling",
    "article": "der"
  },
  {
    "noun": "Schimmel",
    "article": "der"
  },
  {
    "noun": "Schimmelbelag",
    "article": "der"
  },
  {
    "noun": "Schimmelpilz",
    "article": "der"
  },
  {
    "noun": "Schimmelreiter",
    "article": "der"
  },
  {
    "noun": "Schimmer",
    "article": "der"
  },
  {
    "noun": "Schimpanse",
    "article": "der"
  },
  {
    "noun": "Schimpf",
    "article": "der"
  },
  {
    "noun": "Schimpfname",
    "article": "der"
  },
  {
    "noun": "Schinder",
    "article": "der"
  },
  {
    "noun": "Schinderkarren",
    "article": "der"
  },
  {
    "noun": "Schinderknecht",
    "article": "der"
  },
  {
    "noun": "Schinken",
    "article": "der"
  },
  {
    "noun": "Schinkenspeck",
    "article": "der"
  },
  {
    "noun": "Schintoismus",
    "article": "der"
  },
  {
    "noun": "Schiri",
    "article": "der"
  },
  {
    "noun": "Schirm",
    "article": "der"
  },
  {
    "noun": "Schirmer",
    "article": "der"
  },
  {
    "noun": "Schirmgriff",
    "article": "der"
  },
  {
    "noun": "Schirmherr",
    "article": "der"
  },
  {
    "noun": "Schirmpilz",
    "article": "der"
  },
  {
    "noun": "Schirokko",
    "article": "der"
  },
  {
    "noun": "Schirting",
    "article": "der"
  },
  {
    "noun": "Schispringer",
    "article": "der"
  },
  {
    "noun": "Schiss",
    "article": "der"
  },
  {
    "noun": "Schisser",
    "article": "der"
  },
  {
    "noun": "Schistock",
    "article": "der"
  },
  {
    "noun": "Schlachtenbummler",
    "article": "der"
  },
  {
    "noun": "Schlachtgesang",
    "article": "der"
  },
  {
    "noun": "Schlachthof",
    "article": "der"
  },
  {
    "noun": "Schlachtkreuzer",
    "article": "der"
  },
  {
    "noun": "Schlachtplan",
    "article": "der"
  },
  {
    "noun": "Schlachtruf",
    "article": "der"
  },
  {
    "noun": "Schlack",
    "article": "der"
  },
  {
    "noun": "Schlaf",
    "article": "der"
  },
  {
    "noun": "Schlafanzug",
    "article": "der"
  },
  {
    "noun": "Schlafentzug",
    "article": "der"
  },
  {
    "noun": "Schlafgast",
    "article": "der"
  },
  {
    "noun": "Schlafgenosse",
    "article": "der"
  },
  {
    "noun": "Schlafplatz",
    "article": "der"
  },
  {
    "noun": "Schlafraum",
    "article": "der"
  },
  {
    "noun": "Schlafrock",
    "article": "der"
  },
  {
    "noun": "Schlafsaal",
    "article": "der"
  },
  {
    "noun": "Schlafsack",
    "article": "der"
  },
  {
    "noun": "Schlaftrunk",
    "article": "der"
  },
  {
    "noun": "Schlafwagen",
    "article": "der"
  },
  {
    "noun": "Schlafwagenplatz",
    "article": "der"
  },
  {
    "noun": "Schlafwagenschaffner",
    "article": "der"
  },
  {
    "noun": "Schlafwandler",
    "article": "der"
  },
  {
    "noun": "Schlafzustand",
    "article": "der"
  },
  {
    "noun": "Schlag",
    "article": "der"
  },
  {
    "noun": "Schlagabtausch",
    "article": "der"
  },
  {
    "noun": "Schlaganfall",
    "article": "der"
  },
  {
    "noun": "Schlagball",
    "article": "der"
  },
  {
    "noun": "Schlagbaum",
    "article": "der"
  },
  {
    "noun": "Schlagbohrer",
    "article": "der"
  },
  {
    "noun": "Schlagbolzen",
    "article": "der"
  },
  {
    "noun": "Schlagfluss",
    "article": "der"
  },
  {
    "noun": "Schlagmann",
    "article": "der"
  },
  {
    "noun": "Schlagrahm",
    "article": "der"
  },
  {
    "noun": "Schlagring",
    "article": "der"
  },
  {
    "noun": "Schlagstock",
    "article": "der"
  },
  {
    "noun": "Schlagwortkatalog",
    "article": "der"
  },
  {
    "noun": "Schlagzeuger",
    "article": "der"
  },
  {
    "noun": "Schlaks",
    "article": "der"
  },
  {
    "noun": "Schlamm",
    "article": "der"
  },
  {
    "noun": "Schlamp",
    "article": "der"
  },
  {
    "noun": "Schlangenbiss",
    "article": "der"
  },
  {
    "noun": "Schlangenmensch",
    "article": "der"
  },
  {
    "noun": "Schlangenstern",
    "article": "der"
  },
  {
    "noun": "Schlangentanz",
    "article": "der"
  },
  {
    "noun": "Schlapphut",
    "article": "der"
  },
  {
    "noun": "Schlappschwanz",
    "article": "der"
  },
  {
    "noun": "Schlaraffe",
    "article": "der"
  },
  {
    "noun": "Schlauberger",
    "article": "der"
  },
  {
    "noun": "Schlauch",
    "article": "der"
  },
  {
    "noun": "Schlauchpilz",
    "article": "der"
  },
  {
    "noun": "Schlauchreifen",
    "article": "der"
  },
  {
    "noun": "Schlaukopf",
    "article": "der"
  },
  {
    "noun": "Schlaumeier",
    "article": "der"
  },
  {
    "noun": "Schlawiner",
    "article": "der"
  },
  {
    "noun": "Schlecker",
    "article": "der"
  },
  {
    "noun": "Schlehdorn",
    "article": "der"
  },
  {
    "noun": "Schleicher",
    "article": "der"
  },
  {
    "noun": "Schleichhandel",
    "article": "der"
  },
  {
    "noun": "Schleichpfad",
    "article": "der"
  },
  {
    "noun": "Schleichweg",
    "article": "der"
  },
  {
    "noun": "Schleier",
    "article": "der"
  },
  {
    "noun": "Schleiertanz",
    "article": "der"
  },
  {
    "noun": "Schleifapparat",
    "article": "der"
  },
  {
    "noun": "Schleifautomat",
    "article": "der"
  },
  {
    "noun": "Schleifer",
    "article": "der"
  },
  {
    "noun": "Schleifring",
    "article": "der"
  },
  {
    "noun": "Schleifstein",
    "article": "der"
  },
  {
    "noun": "Schleim",
    "article": "der"
  },
  {
    "noun": "Schleimbeutel",
    "article": "der"
  },
  {
    "noun": "Schleimer",
    "article": "der"
  },
  {
    "noun": "Schleimfisch",
    "article": "der"
  },
  {
    "noun": "Schleimpilz",
    "article": "der"
  },
  {
    "noun": "Schlemmer",
    "article": "der"
  },
  {
    "noun": "Schlendergang",
    "article": "der"
  },
  {
    "noun": "Schlendrian",
    "article": "der"
  },
  {
    "noun": "Schlenzer",
    "article": "der"
  },
  {
    "noun": "Schleppdampfer",
    "article": "der"
  },
  {
    "noun": "Schlepper",
    "article": "der"
  },
  {
    "noun": "Schleppkahn",
    "article": "der"
  },
  {
    "noun": "Schlepplift",
    "article": "der"
  },
  {
    "noun": "Schleppzug",
    "article": "der"
  },
  {
    "noun": "Schlesier",
    "article": "der"
  },
  {
    "noun": "Schleuderbeton",
    "article": "der"
  },
  {
    "noun": "Schleuderer",
    "article": "der"
  },
  {
    "noun": "Schleuderhonig",
    "article": "der"
  },
  {
    "noun": "Schleuderpreis",
    "article": "der"
  },
  {
    "noun": "Schleudersitz",
    "article": "der"
  },
  {
    "noun": "Schleuderstart",
    "article": "der"
  },
  {
    "noun": "Schleuser",
    "article": "der"
  },
  {
    "noun": "Schlich",
    "article": "der"
  },
  {
    "noun": "Schlichter",
    "article": "der"
  },
  {
    "noun": "Schlichthobel",
    "article": "der"
  },
  {
    "noun": "Schlichtungsausschuss",
    "article": "der"
  },
  {
    "noun": "Schlichtungsversuch",
    "article": "der"
  },
  {
    "noun": "Schlick",
    "article": "der"
  },
  {
    "noun": "Schliff",
    "article": "der"
  },
  {
    "noun": "Schlingel",
    "article": "der"
  },
  {
    "noun": "Schlingensteller",
    "article": "der"
  },
  {
    "noun": "Schlips",
    "article": "der"
  },
  {
    "noun": "Schlittenhund",
    "article": "der"
  },
  {
    "noun": "Schlittschuh",
    "article": "der"
  },
  {
    "noun": "Schlitz",
    "article": "der"
  },
  {
    "noun": "Schlitzverschluss",
    "article": "der"
  },
  {
    "noun": "Schlosser",
    "article": "der"
  },
  {
    "noun": "Schlossgarten",
    "article": "der"
  },
  {
    "noun": "Schlossherr",
    "article": "der"
  },
  {
    "noun": "Schlosshof",
    "article": "der"
  },
  {
    "noun": "Schlosspark",
    "article": "der"
  },
  {
    "noun": "Schlot",
    "article": "der"
  },
  {
    "noun": "Schluchzer",
    "article": "der"
  },
  {
    "noun": "Schluck",
    "article": "der"
  },
  {
    "noun": "Schluckauf",
    "article": "der"
  },
  {
    "noun": "Schlucker",
    "article": "der"
  },
  {
    "noun": "Schluff",
    "article": "der"
  },
  {
    "noun": "Schlummer",
    "article": "der"
  },
  {
    "noun": "Schlummertrunk",
    "article": "der"
  },
  {
    "noun": "Schlumpf",
    "article": "der"
  },
  {
    "noun": "Schlund",
    "article": "der"
  },
  {
    "noun": "Schlupf",
    "article": "der"
  },
  {
    "noun": "Schlupfwinkel",
    "article": "der"
  },
  {
    "noun": "Schluss",
    "article": "der"
  },
  {
    "noun": "Schlussakt",
    "article": "der"
  },
  {
    "noun": "Schlussbericht",
    "article": "der"
  },
  {
    "noun": "Schlusskurs",
    "article": "der"
  },
  {
    "noun": "Schlussmann",
    "article": "der"
  },
  {
    "noun": "Schlusspfiff",
    "article": "der"
  },
  {
    "noun": "Schlusspunkt",
    "article": "der"
  },
  {
    "noun": "Schlussschein",
    "article": "der"
  },
  {
    "noun": "Schlusstermin",
    "article": "der"
  },
  {
    "noun": "Schlussverkauf",
    "article": "der"
  },
  {
    "noun": "Schmachtfetzen",
    "article": "der"
  },
  {
    "noun": "Schmachtlappen",
    "article": "der"
  },
  {
    "noun": "Schmachtriemen",
    "article": "der"
  },
  {
    "noun": "Schmalfilm",
    "article": "der"
  },
  {
    "noun": "Schmand",
    "article": "der"
  },
  {
    "noun": "Schmarotzer",
    "article": "der"
  },
  {
    "noun": "Schmarren",
    "article": "der"
  },
  {
    "noun": "Schmatz",
    "article": "der"
  },
  {
    "noun": "Schmatzer",
    "article": "der"
  },
  {
    "noun": "Schmauch",
    "article": "der"
  },
  {
    "noun": "Schmaus",
    "article": "der"
  },
  {
    "noun": "Schmeichler",
    "article": "der"
  },
  {
    "noun": "Schmelz",
    "article": "der"
  },
  {
    "noun": "Schmelzer",
    "article": "der"
  },
  {
    "noun": "Schmelzofen",
    "article": "der"
  },
  {
    "noun": "Schmelzpunkt",
    "article": "der"
  },
  {
    "noun": "Schmelztiegel",
    "article": "der"
  },
  {
    "noun": "Schmerbauch",
    "article": "der"
  },
  {
    "noun": "Schmerz",
    "article": "der"
  },
  {
    "noun": "Schmerzenslaut",
    "article": "der"
  },
  {
    "noun": "Schmerzensschrei",
    "article": "der"
  },
  {
    "noun": "Schmetterball",
    "article": "der"
  },
  {
    "noun": "Schmetterling",
    "article": "der"
  },
  {
    "noun": "Schmetterlingsstil",
    "article": "der"
  },
  {
    "noun": "Schmied",
    "article": "der"
  },
  {
    "noun": "Schmiedehammer",
    "article": "der"
  },
  {
    "noun": "Schmiedeofen",
    "article": "der"
  },
  {
    "noun": "Schmierblock",
    "article": "der"
  },
  {
    "noun": "Schmierer",
    "article": "der"
  },
  {
    "noun": "Schmierfink",
    "article": "der"
  },
  {
    "noun": "Schmiernippel",
    "article": "der"
  },
  {
    "noun": "Schmierplan",
    "article": "der"
  },
  {
    "noun": "Schmierstoff",
    "article": "der"
  },
  {
    "noun": "Schmierzettel",
    "article": "der"
  },
  {
    "noun": "Schminkkoffer",
    "article": "der"
  },
  {
    "noun": "Schminktisch",
    "article": "der"
  },
  {
    "noun": "Schminktopf",
    "article": "der"
  },
  {
    "noun": "Schmirgel",
    "article": "der"
  },
  {
    "noun": "Schmiss",
    "article": "der"
  },
  {
    "noun": "Schmitz",
    "article": "der"
  },
  {
    "noun": "Schmock",
    "article": "der"
  },
  {
    "noun": "Schmollmund",
    "article": "der"
  },
  {
    "noun": "Schmonzes",
    "article": "der"
  },
  {
    "noun": "Schmorbraten",
    "article": "der"
  },
  {
    "noun": "Schmortopf",
    "article": "der"
  },
  {
    "noun": "Schmuckkasten",
    "article": "der"
  },
  {
    "noun": "Schmuckstein",
    "article": "der"
  },
  {
    "noun": "Schmuddel",
    "article": "der"
  },
  {
    "noun": "Schmuggel",
    "article": "der"
  },
  {
    "noun": "Schmuggler",
    "article": "der"
  },
  {
    "noun": "Schmugglerpfad",
    "article": "der"
  },
  {
    "noun": "Schmus",
    "article": "der"
  },
  {
    "noun": "Schmusekater",
    "article": "der"
  },
  {
    "noun": "Schmutz",
    "article": "der"
  },
  {
    "noun": "Schmutzfink",
    "article": "der"
  },
  {
    "noun": "Schmutzfleck",
    "article": "der"
  },
  {
    "noun": "Schmutzgeier",
    "article": "der"
  },
  {
    "noun": "Schmutztitel",
    "article": "der"
  },
  {
    "noun": "Schnabel",
    "article": "der"
  },
  {
    "noun": "Schnabelhieb",
    "article": "der"
  },
  {
    "noun": "Schnack",
    "article": "der"
  },
  {
    "noun": "Schnallenschuh",
    "article": "der"
  },
  {
    "noun": "Schnalzer",
    "article": "der"
  },
  {
    "noun": "Schnapper",
    "article": "der"
  },
  {
    "noun": "Schnappschuss",
    "article": "der"
  },
  {
    "noun": "Schnappverschluss",
    "article": "der"
  },
  {
    "noun": "Schnaps",
    "article": "der"
  },
  {
    "noun": "Schnapsbrenner",
    "article": "der"
  },
  {
    "noun": "Schnarcher",
    "article": "der"
  },
  {
    "noun": "Schnauz",
    "article": "der"
  },
  {
    "noun": "Schnauzbart",
    "article": "der"
  },
  {
    "noun": "Schneckenbohrer",
    "article": "der"
  },
  {
    "noun": "Schneckengang",
    "article": "der"
  },
  {
    "noun": "Schnee",
    "article": "der"
  },
  {
    "noun": "Schneeball",
    "article": "der"
  },
  {
    "noun": "Schneeberg",
    "article": "der"
  },
  {
    "noun": "Schneebesen",
    "article": "der"
  },
  {
    "noun": "Schneefall",
    "article": "der"
  },
  {
    "noun": "Schneefink",
    "article": "der"
  },
  {
    "noun": "Schneehase",
    "article": "der"
  },
  {
    "noun": "Schneemann",
    "article": "der"
  },
  {
    "noun": "Schneematsch",
    "article": "der"
  },
  {
    "noun": "Schneemond",
    "article": "der"
  },
  {
    "noun": "Schneepflug",
    "article": "der"
  },
  {
    "noun": "Schneeregen",
    "article": "der"
  },
  {
    "noun": "Schneereifen",
    "article": "der"
  },
  {
    "noun": "Schneeschauer",
    "article": "der"
  },
  {
    "noun": "Schneeschuh",
    "article": "der"
  },
  {
    "noun": "Schneesturm",
    "article": "der"
  },
  {
    "noun": "Schneid",
    "article": "der"
  },
  {
    "noun": "Schneidbrenner",
    "article": "der"
  },
  {
    "noun": "Schneider",
    "article": "der"
  },
  {
    "noun": "Schneideraum",
    "article": "der"
  },
  {
    "noun": "Schneidergeselle",
    "article": "der"
  },
  {
    "noun": "Schneidermeister",
    "article": "der"
  },
  {
    "noun": "Schneidermuskel",
    "article": "der"
  },
  {
    "noun": "Schneidersitz",
    "article": "der"
  },
  {
    "noun": "Schneidezahn",
    "article": "der"
  },
  {
    "noun": "Schnellbus",
    "article": "der"
  },
  {
    "noun": "Schnelldampfer",
    "article": "der"
  },
  {
    "noun": "Schnelldienst",
    "article": "der"
  },
  {
    "noun": "Schnelldrucker",
    "article": "der"
  },
  {
    "noun": "Schnellfahrer",
    "article": "der"
  },
  {
    "noun": "Schnellgang",
    "article": "der"
  },
  {
    "noun": "Schnellhefter",
    "article": "der"
  },
  {
    "noun": "Schnellimbiss",
    "article": "der"
  },
  {
    "noun": "Schnellkocher",
    "article": "der"
  },
  {
    "noun": "Schnellkochtopf",
    "article": "der"
  },
  {
    "noun": "Schnellkurs",
    "article": "der"
  },
  {
    "noun": "Schnellschreiber",
    "article": "der"
  },
  {
    "noun": "Schnellschritt",
    "article": "der"
  },
  {
    "noun": "Schnellsegler",
    "article": "der"
  },
  {
    "noun": "Schnellsieder",
    "article": "der"
  },
  {
    "noun": "Schnellstahl",
    "article": "der"
  },
  {
    "noun": "Schnellverkehr",
    "article": "der"
  },
  {
    "noun": "Schnellzug",
    "article": "der"
  },
  {
    "noun": "Schnepper",
    "article": "der"
  },
  {
    "noun": "Schnickschnack",
    "article": "der"
  },
  {
    "noun": "Schniedelwutz",
    "article": "der"
  },
  {
    "noun": "Schnipser",
    "article": "der"
  },
  {
    "noun": "Schnitt",
    "article": "der"
  },
  {
    "noun": "Schnitter",
    "article": "der"
  },
  {
    "noun": "Schnittlauch",
    "article": "der"
  },
  {
    "noun": "Schnittmeister",
    "article": "der"
  },
  {
    "noun": "Schnittmusterbogen",
    "article": "der"
  },
  {
    "noun": "Schnittpunkt",
    "article": "der"
  },
  {
    "noun": "Schnitz",
    "article": "der"
  },
  {
    "noun": "Schnitzer",
    "article": "der"
  },
  {
    "noun": "Schnodder",
    "article": "der"
  },
  {
    "noun": "Schnorchel",
    "article": "der"
  },
  {
    "noun": "Schnorrer",
    "article": "der"
  },
  {
    "noun": "Schnuller",
    "article": "der"
  },
  {
    "noun": "Schnupftabak",
    "article": "der"
  },
  {
    "noun": "Schnupperkurs",
    "article": "der"
  },
  {
    "noun": "Schnurrbart",
    "article": "der"
  },
  {
    "noun": "Schober",
    "article": "der"
  },
  {
    "noun": "Schocker",
    "article": "der"
  },
  {
    "noun": "Schofel",
    "article": "der"
  },
  {
    "noun": "Schokoladenguss",
    "article": "der"
  },
  {
    "noun": "Schokoladenpudding",
    "article": "der"
  },
  {
    "noun": "Schokoriegel",
    "article": "der"
  },
  {
    "noun": "Scholar",
    "article": "der"
  },
  {
    "noun": "Scholastizismus",
    "article": "der"
  },
  {
    "noun": "Schonbezug",
    "article": "der"
  },
  {
    "noun": "Schoner",
    "article": "der"
  },
  {
    "noun": "Schonwaschgang",
    "article": "der"
  },
  {
    "noun": "Schopf",
    "article": "der"
  },
  {
    "noun": "Schorf",
    "article": "der"
  },
  {
    "noun": "Schornstein",
    "article": "der"
  },
  {
    "noun": "Schornsteinfeger",
    "article": "der"
  },
  {
    "noun": "Schote",
    "article": "der"
  },
  {
    "noun": "Schottenrock",
    "article": "der"
  },
  {
    "noun": "Schotter",
    "article": "der"
  },
  {
    "noun": "Schotterweg",
    "article": "der"
  },
  {
    "noun": "Schrank",
    "article": "der"
  },
  {
    "noun": "Schrankkoffer",
    "article": "der"
  },
  {
    "noun": "Schranze",
    "article": "der"
  },
  {
    "noun": "Schrapper",
    "article": "der"
  },
  {
    "noun": "Schrat",
    "article": "der"
  },
  {
    "noun": "Schraubdeckel",
    "article": "der"
  },
  {
    "noun": "Schraubenbolzen",
    "article": "der"
  },
  {
    "noun": "Schraubendampfer",
    "article": "der"
  },
  {
    "noun": "Schraubendreher",
    "article": "der"
  },
  {
    "noun": "Schraubengang",
    "article": "der"
  },
  {
    "noun": "Schraubenkopf",
    "article": "der"
  },
  {
    "noun": "Schraubenzieher",
    "article": "der"
  },
  {
    "noun": "Schraubstock",
    "article": "der"
  },
  {
    "noun": "Schraubverschluss",
    "article": "der"
  },
  {
    "noun": "Schrebergarten",
    "article": "der"
  },
  {
    "noun": "Schreck",
    "article": "der"
  },
  {
    "noun": "Schreckensruf",
    "article": "der"
  },
  {
    "noun": "Schreckensschrei",
    "article": "der"
  },
  {
    "noun": "Schreckschuss",
    "article": "der"
  },
  {
    "noun": "Schredder",
    "article": "der"
  },
  {
    "noun": "Schrei",
    "article": "der"
  },
  {
    "noun": "Schreiadler",
    "article": "der"
  },
  {
    "noun": "Schreibautomat",
    "article": "der"
  },
  {
    "noun": "Schreibbedarf",
    "article": "der"
  },
  {
    "noun": "Schreibblock",
    "article": "der"
  },
  {
    "noun": "Schreiber",
    "article": "der"
  },
  {
    "noun": "Schreiberling",
    "article": "der"
  },
  {
    "noun": "Schreibfehler",
    "article": "der"
  },
  {
    "noun": "Schreibgriffel",
    "article": "der"
  },
  {
    "noun": "Schreibkopf",
    "article": "der"
  },
  {
    "noun": "Schreibkrampf",
    "article": "der"
  },
  {
    "noun": "Schreibschutz",
    "article": "der"
  },
  {
    "noun": "Schreibstift",
    "article": "der"
  },
  {
    "noun": "Schreibtisch",
    "article": "der"
  },
  {
    "noun": "Schreibwarenladen",
    "article": "der"
  },
  {
    "noun": "Schreier",
    "article": "der"
  },
  {
    "noun": "Schreihals",
    "article": "der"
  },
  {
    "noun": "Schreikrampf",
    "article": "der"
  },
  {
    "noun": "Schrein",
    "article": "der"
  },
  {
    "noun": "Schreiner",
    "article": "der"
  },
  {
    "noun": "Schriftgrad",
    "article": "der"
  },
  {
    "noun": "Schriftkegel",
    "article": "der"
  },
  {
    "noun": "Schriftleiter",
    "article": "der"
  },
  {
    "noun": "Schriftsatz",
    "article": "der"
  },
  {
    "noun": "Schriftsetzer",
    "article": "der"
  },
  {
    "noun": "Schriftsteller",
    "article": "der"
  },
  {
    "noun": "Schriftstellername",
    "article": "der"
  },
  {
    "noun": "Schriftverkehr",
    "article": "der"
  },
  {
    "noun": "Schriftwechsel",
    "article": "der"
  },
  {
    "noun": "Schriftzug",
    "article": "der"
  },
  {
    "noun": "Schrimp",
    "article": "der"
  },
  {
    "noun": "Schrimpscocktail",
    "article": "der"
  },
  {
    "noun": "Schritt",
    "article": "der"
  },
  {
    "noun": "Schrittmacher",
    "article": "der"
  },
  {
    "noun": "Schrittmacherdienst",
    "article": "der"
  },
  {
    "noun": "Schrittmesser",
    "article": "der"
  },
  {
    "noun": "Schrott",
    "article": "der"
  },
  {
    "noun": "Schrotthandel",
    "article": "der"
  },
  {
    "noun": "Schrotthaufen",
    "article": "der"
  },
  {
    "noun": "Schrottplatz",
    "article": "der"
  },
  {
    "noun": "Schrottwert",
    "article": "der"
  },
  {
    "noun": "Schrubber",
    "article": "der"
  },
  {
    "noun": "Schrumpf",
    "article": "der"
  },
  {
    "noun": "Schrumpfkopf",
    "article": "der"
  },
  {
    "noun": "Schrumpfungsprozess",
    "article": "der"
  },
  {
    "noun": "Schrund",
    "article": "der"
  },
  {
    "noun": "Schrupphobel",
    "article": "der"
  },
  {
    "noun": "Schruppstahl",
    "article": "der"
  },
  {
    "noun": "Schub",
    "article": "der"
  },
  {
    "noun": "Schuber",
    "article": "der"
  },
  {
    "noun": "Schubkarren",
    "article": "der"
  },
  {
    "noun": "Schubkasten",
    "article": "der"
  },
  {
    "noun": "Schubmodul",
    "article": "der"
  },
  {
    "noun": "Schubs",
    "article": "der"
  },
  {
    "noun": "Schuft",
    "article": "der"
  },
  {
    "noun": "Schuh",
    "article": "der"
  },
  {
    "noun": "Schuhabsatz",
    "article": "der"
  },
  {
    "noun": "Schuhanzieher",
    "article": "der"
  },
  {
    "noun": "Schuhkarton",
    "article": "der"
  },
  {
    "noun": "Schuhladen",
    "article": "der"
  },
  {
    "noun": "Schuhlappen",
    "article": "der"
  },
  {
    "noun": "Schuhleisten",
    "article": "der"
  },
  {
    "noun": "Schuhmacher",
    "article": "der"
  },
  {
    "noun": "Schuhnagel",
    "article": "der"
  },
  {
    "noun": "Schuhplattler",
    "article": "der"
  },
  {
    "noun": "Schuhputzer",
    "article": "der"
  },
  {
    "noun": "Schuhriemen",
    "article": "der"
  },
  {
    "noun": "Schuhschnabel",
    "article": "der"
  },
  {
    "noun": "Schuhspanner",
    "article": "der"
  },
  {
    "noun": "Schukostecker",
    "article": "der"
  },
  {
    "noun": "Schulabschluss",
    "article": "der"
  },
  {
    "noun": "Schulanfang",
    "article": "der"
  },
  {
    "noun": "Schularzt",
    "article": "der"
  },
  {
    "noun": "Schulatlas",
    "article": "der"
  },
  {
    "noun": "Schulaufsatz",
    "article": "der"
  },
  {
    "noun": "Schulausflug",
    "article": "der"
  },
  {
    "noun": "Schulbau",
    "article": "der"
  },
  {
    "noun": "Schulbeginn",
    "article": "der"
  },
  {
    "noun": "Schulbesuch",
    "article": "der"
  },
  {
    "noun": "Schulbub",
    "article": "der"
  },
  {
    "noun": "Schulbus",
    "article": "der"
  },
  {
    "noun": "Schulchor",
    "article": "der"
  },
  {
    "noun": "Schuldbeitritt",
    "article": "der"
  },
  {
    "noun": "Schuldbetrag",
    "article": "der"
  },
  {
    "noun": "Schuldbeweis",
    "article": "der"
  },
  {
    "noun": "Schuldbrief",
    "article": "der"
  },
  {
    "noun": "Schuldenberg",
    "article": "der"
  },
  {
    "noun": "Schuldendienst",
    "article": "der"
  },
  {
    "noun": "Schuldenerlass",
    "article": "der"
  },
  {
    "noun": "Schuldiener",
    "article": "der"
  },
  {
    "noun": "Schuldienst",
    "article": "der"
  },
  {
    "noun": "Schuldiger",
    "article": "der"
  },
  {
    "noun": "Schuldirektor",
    "article": "der"
  },
  {
    "noun": "Schuldkomplex",
    "article": "der"
  },
  {
    "noun": "Schuldner",
    "article": "der"
  },
  {
    "noun": "Schuldschein",
    "article": "der"
  },
  {
    "noun": "Schuldspruch",
    "article": "der"
  },
  {
    "noun": "Schuldtitel",
    "article": "der"
  },
  {
    "noun": "Schuldturm",
    "article": "der"
  },
  {
    "noun": "Schuldwechsel",
    "article": "der"
  },
  {
    "noun": "Schuldzins",
    "article": "der"
  },
  {
    "noun": "Schulfreund",
    "article": "der"
  },
  {
    "noun": "Schulfuchs",
    "article": "der"
  },
  {
    "noun": "Schulfunk",
    "article": "der"
  },
  {
    "noun": "Schulgebrauch",
    "article": "der"
  },
  {
    "noun": "Schulhof",
    "article": "der"
  },
  {
    "noun": "Schuljunge",
    "article": "der"
  },
  {
    "noun": "Schulkamerad",
    "article": "der"
  },
  {
    "noun": "Schulkollege",
    "article": "der"
  },
  {
    "noun": "Schullehrer",
    "article": "der"
  },
  {
    "noun": "Schulleiter",
    "article": "der"
  },
  {
    "noun": "Schulmann",
    "article": "der"
  },
  {
    "noun": "Schulmeister",
    "article": "der"
  },
  {
    "noun": "Schulpsychologe",
    "article": "der"
  },
  {
    "noun": "Schulranzen",
    "article": "der"
  },
  {
    "noun": "Schulrat",
    "article": "der"
  },
  {
    "noun": "Schulraum",
    "article": "der"
  },
  {
    "noun": "Schulsport",
    "article": "der"
  },
  {
    "noun": "Schulsprecher",
    "article": "der"
  },
  {
    "noun": "Schulstress",
    "article": "der"
  },
  {
    "noun": "Schultag",
    "article": "der"
  },
  {
    "noun": "Schulterschluss",
    "article": "der"
  },
  {
    "noun": "Schultersieg",
    "article": "der"
  },
  {
    "noun": "Schulunterricht",
    "article": "der"
  },
  {
    "noun": "Schulvorsteher",
    "article": "der"
  },
  {
    "noun": "Schulweg",
    "article": "der"
  },
  {
    "noun": "Schulze",
    "article": "der"
  },
  {
    "noun": "Schummel",
    "article": "der"
  },
  {
    "noun": "Schummer",
    "article": "der"
  },
  {
    "noun": "Schund",
    "article": "der"
  },
  {
    "noun": "Schundroman",
    "article": "der"
  },
  {
    "noun": "Schuppenbaum",
    "article": "der"
  },
  {
    "noun": "Schuppenpanzer",
    "article": "der"
  },
  {
    "noun": "Schurke",
    "article": "der"
  },
  {
    "noun": "Schurkenstreich",
    "article": "der"
  },
  {
    "noun": "Schurz",
    "article": "der"
  },
  {
    "noun": "Schuss",
    "article": "der"
  },
  {
    "noun": "Schussbereich",
    "article": "der"
  },
  {
    "noun": "Schussfaden",
    "article": "der"
  },
  {
    "noun": "Schusswaffengebrauch",
    "article": "der"
  },
  {
    "noun": "Schusswechsel",
    "article": "der"
  },
  {
    "noun": "Schuster",
    "article": "der"
  },
  {
    "noun": "Schusterjunge",
    "article": "der"
  },
  {
    "noun": "Schusterlehrling",
    "article": "der"
  },
  {
    "noun": "Schusterschemel",
    "article": "der"
  },
  {
    "noun": "Schutt",
    "article": "der"
  },
  {
    "noun": "Schuttabladeplatz",
    "article": "der"
  },
  {
    "noun": "Schuttkarren",
    "article": "der"
  },
  {
    "noun": "Schuttplatz",
    "article": "der"
  },
  {
    "noun": "Schutzanstrich",
    "article": "der"
  },
  {
    "noun": "Schutzanzug",
    "article": "der"
  },
  {
    "noun": "Schutzbereich",
    "article": "der"
  },
  {
    "noun": "Schutzbrief",
    "article": "der"
  },
  {
    "noun": "Schutzdamm",
    "article": "der"
  },
  {
    "noun": "Schutzengel",
    "article": "der"
  },
  {
    "noun": "Schutzgeist",
    "article": "der"
  },
  {
    "noun": "Schutzgott",
    "article": "der"
  },
  {
    "noun": "Schutzhafen",
    "article": "der"
  },
  {
    "noun": "Schutzhelm",
    "article": "der"
  },
  {
    "noun": "Schutzhund",
    "article": "der"
  },
  {
    "noun": "Schutzkontaktstecker",
    "article": "der"
  },
  {
    "noun": "Schutzlack",
    "article": "der"
  },
  {
    "noun": "Schutzmann",
    "article": "der"
  },
  {
    "noun": "Schutzmantel",
    "article": "der"
  },
  {
    "noun": "Schutzort",
    "article": "der"
  },
  {
    "noun": "Schutzpatron",
    "article": "der"
  },
  {
    "noun": "Schutzpolizist",
    "article": "der"
  },
  {
    "noun": "Schutzraum",
    "article": "der"
  },
  {
    "noun": "Schutzschild",
    "article": "der"
  },
  {
    "noun": "Schutzstaat",
    "article": "der"
  },
  {
    "noun": "Schutzstoff",
    "article": "der"
  },
  {
    "noun": "Schutzumschlag",
    "article": "der"
  },
  {
    "noun": "Schutzverband",
    "article": "der"
  },
  {
    "noun": "Schutzwald",
    "article": "der"
  },
  {
    "noun": "Schutzwall",
    "article": "der"
  },
  {
    "noun": "Schutzzoll",
    "article": "der"
  },
  {
    "noun": "Schwabe",
    "article": "der"
  },
  {
    "noun": "Schwabenstreich",
    "article": "der"
  },
  {
    "noun": "Schwachkopf",
    "article": "der"
  },
  {
    "noun": "Schwachpunkt",
    "article": "der"
  },
  {
    "noun": "Schwachsinn",
    "article": "der"
  },
  {
    "noun": "Schwachstrom",
    "article": "der"
  },
  {
    "noun": "Schwaden",
    "article": "der"
  },
  {
    "noun": "Schwadroneur",
    "article": "der"
  },
  {
    "noun": "Schwafler",
    "article": "der"
  },
  {
    "noun": "Schwager",
    "article": "der"
  },
  {
    "noun": "Schwalbenschwanz",
    "article": "der"
  },
  {
    "noun": "Schwall",
    "article": "der"
  },
  {
    "noun": "Schwamm",
    "article": "der"
  },
  {
    "noun": "Schwan",
    "article": "der"
  },
  {
    "noun": "Schwanengesang",
    "article": "der"
  },
  {
    "noun": "Schwanenhals",
    "article": "der"
  },
  {
    "noun": "Schwangerschaftsabbruch",
    "article": "der"
  },
  {
    "noun": "Schwangerschaftsstreifen",
    "article": "der"
  },
  {
    "noun": "Schwangerschaftstest",
    "article": "der"
  },
  {
    "noun": "Schwangerschaftsurlaub",
    "article": "der"
  },
  {
    "noun": "Schwanz",
    "article": "der"
  },
  {
    "noun": "Schwanzlurch",
    "article": "der"
  },
  {
    "noun": "Schwanzwirbel",
    "article": "der"
  },
  {
    "noun": "Schwarm",
    "article": "der"
  },
  {
    "noun": "Schwartenmagen",
    "article": "der"
  },
  {
    "noun": "Schwarzarbeiter",
    "article": "der"
  },
  {
    "noun": "Schwarzbrenner",
    "article": "der"
  },
  {
    "noun": "Schwarzdorn",
    "article": "der"
  },
  {
    "noun": "Schwarzfahrer",
    "article": "der"
  },
  {
    "noun": "Schwarzhandel",
    "article": "der"
  },
  {
    "noun": "Schwarzmaler",
    "article": "der"
  },
  {
    "noun": "Schwarzmarkt",
    "article": "der"
  },
  {
    "noun": "Schwarzmarktpreis",
    "article": "der"
  },
  {
    "noun": "Schwarzseher",
    "article": "der"
  },
  {
    "noun": "Schwarzsender",
    "article": "der"
  },
  {
    "noun": "Schwarzspecht",
    "article": "der"
  },
  {
    "noun": "Schwarzstorch",
    "article": "der"
  },
  {
    "noun": "Schwarzwald",
    "article": "der"
  },
  {
    "noun": "Schwatz",
    "article": "der"
  },
  {
    "noun": "Schwebebalken",
    "article": "der"
  },
  {
    "noun": "Schwebebaum",
    "article": "der"
  },
  {
    "noun": "Schwebeflug",
    "article": "der"
  },
  {
    "noun": "Schwebestoff",
    "article": "der"
  },
  {
    "noun": "Schwebezustand",
    "article": "der"
  },
  {
    "noun": "Schwebstoff",
    "article": "der"
  },
  {
    "noun": "Schwede",
    "article": "der"
  },
  {
    "noun": "Schwefel",
    "article": "der"
  },
  {
    "noun": "Schwefeldampf",
    "article": "der"
  },
  {
    "noun": "Schwefelfarbstoff",
    "article": "der"
  },
  {
    "noun": "Schwefelkies",
    "article": "der"
  },
  {
    "noun": "Schwefelkohlenstoff",
    "article": "der"
  },
  {
    "noun": "Schwefelwasserstoff",
    "article": "der"
  },
  {
    "noun": "Schweif",
    "article": "der"
  },
  {
    "noun": "Schweifstern",
    "article": "der"
  },
  {
    "noun": "Schweigemarsch",
    "article": "der"
  },
  {
    "noun": "Schweiger",
    "article": "der"
  },
  {
    "noun": "Schweinebraten",
    "article": "der"
  },
  {
    "noun": "Schweinehirt",
    "article": "der"
  },
  {
    "noun": "Schweinekoben",
    "article": "der"
  },
  {
    "noun": "Schweinekopf",
    "article": "der"
  },
  {
    "noun": "Schweinestall",
    "article": "der"
  },
  {
    "noun": "Schweinigel",
    "article": "der"
  },
  {
    "noun": "Schweinsbraten",
    "article": "der"
  },
  {
    "noun": "Schweinskopf",
    "article": "der"
  },
  {
    "noun": "Schwelger",
    "article": "der"
  },
  {
    "noun": "Schwellenwert",
    "article": "der"
  },
  {
    "noun": "Schwemmkegel",
    "article": "der"
  },
  {
    "noun": "Schwemmsand",
    "article": "der"
  },
  {
    "noun": "Schwengel",
    "article": "der"
  },
  {
    "noun": "Schwenk",
    "article": "der"
  },
  {
    "noun": "Schwenkarm",
    "article": "der"
  },
  {
    "noun": "Schwenkbereich",
    "article": "der"
  },
  {
    "noun": "Schwenkkran",
    "article": "der"
  },
  {
    "noun": "Schwerarbeiter",
    "article": "der"
  },
  {
    "noun": "Schwerathlet",
    "article": "der"
  },
  {
    "noun": "Schwerbeton",
    "article": "der"
  },
  {
    "noun": "Schweregrad",
    "article": "der"
  },
  {
    "noun": "Schwergewichtler",
    "article": "der"
  },
  {
    "noun": "Schwerpunkt",
    "article": "der"
  },
  {
    "noun": "Schwerpunktstreik",
    "article": "der"
  },
  {
    "noun": "Schwertfisch",
    "article": "der"
  },
  {
    "noun": "Schwertkampf",
    "article": "der"
  },
  {
    "noun": "Schwertkasten",
    "article": "der"
  },
  {
    "noun": "Schwertknauf",
    "article": "der"
  },
  {
    "noun": "Schwertransport",
    "article": "der"
  },
  {
    "noun": "Schwertstreich",
    "article": "der"
  },
  {
    "noun": "Schwertwal",
    "article": "der"
  },
  {
    "noun": "Schwerverbrecher",
    "article": "der"
  },
  {
    "noun": "Schwerwasserreaktor",
    "article": "der"
  },
  {
    "noun": "Schwesternorden",
    "article": "der"
  },
  {
    "noun": "Schwibbogen",
    "article": "der"
  },
  {
    "noun": "Schwiegersohn",
    "article": "der"
  },
  {
    "noun": "Schwiegervater",
    "article": "der"
  },
  {
    "noun": "Schwierigkeitsgrad",
    "article": "der"
  },
  {
    "noun": "Schwimmanzug",
    "article": "der"
  },
  {
    "noun": "Schwimmbagger",
    "article": "der"
  },
  {
    "noun": "Schwimmer",
    "article": "der"
  },
  {
    "noun": "Schwimmkran",
    "article": "der"
  },
  {
    "noun": "Schwimmlehrer",
    "article": "der"
  },
  {
    "noun": "Schwimmpanzer",
    "article": "der"
  },
  {
    "noun": "Schwimmsand",
    "article": "der"
  },
  {
    "noun": "Schwimmunterricht",
    "article": "der"
  },
  {
    "noun": "Schwimmvogel",
    "article": "der"
  },
  {
    "noun": "Schwindel",
    "article": "der"
  },
  {
    "noun": "Schwindelanfall",
    "article": "der"
  },
  {
    "noun": "Schwindler",
    "article": "der"
  },
  {
    "noun": "Schwinger",
    "article": "der"
  },
  {
    "noun": "Schwingkreis",
    "article": "der"
  },
  {
    "noun": "Schwingquarz",
    "article": "der"
  },
  {
    "noun": "Schwingungskreis",
    "article": "der"
  },
  {
    "noun": "Schwingungszustand",
    "article": "der"
  },
  {
    "noun": "Schwips",
    "article": "der"
  },
  {
    "noun": "Schwirrvogel",
    "article": "der"
  },
  {
    "noun": "Schwitzkasten",
    "article": "der"
  },
  {
    "noun": "Schwulentreff",
    "article": "der"
  },
  {
    "noun": "Schwulst",
    "article": "der"
  },
  {
    "noun": "Schwund",
    "article": "der"
  },
  {
    "noun": "Schwung",
    "article": "der"
  },
  {
    "noun": "Schwur",
    "article": "der"
  },
  {
    "noun": "Scoop",
    "article": "der"
  },
  {
    "noun": "Scotchterrier",
    "article": "der"
  },
  {
    "noun": "Scout",
    "article": "der"
  },
  {
    "noun": "Sechsachser",
    "article": "der"
  },
  {
    "noun": "Sechser",
    "article": "der"
  },
  {
    "noun": "Sechskant",
    "article": "der"
  },
  {
    "noun": "Sechsling",
    "article": "der"
  },
  {
    "noun": "See",
    "article": "der"
  },
  {
    "noun": "Seeaal",
    "article": "der"
  },
  {
    "noun": "Seeadler",
    "article": "der"
  },
  {
    "noun": "Seeblick",
    "article": "der"
  },
  {
    "noun": "Seefahrer",
    "article": "der"
  },
  {
    "noun": "Seefisch",
    "article": "der"
  },
  {
    "noun": "Seefrachtbrief",
    "article": "der"
  },
  {
    "noun": "Seegang",
    "article": "der"
  },
  {
    "noun": "Seehafen",
    "article": "der"
  },
  {
    "noun": "Seehandel",
    "article": "der"
  },
  {
    "noun": "Seehecht",
    "article": "der"
  },
  {
    "noun": "Seehund",
    "article": "der"
  },
  {
    "noun": "Seeigel",
    "article": "der"
  },
  {
    "noun": "Seekadett",
    "article": "der"
  },
  {
    "noun": "Seekanal",
    "article": "der"
  },
  {
    "noun": "Seekrieg",
    "article": "der"
  },
  {
    "noun": "Seelachs",
    "article": "der"
  },
  {
    "noun": "Seelenadel",
    "article": "der"
  },
  {
    "noun": "Seelenforscher",
    "article": "der"
  },
  {
    "noun": "Seelenfriede",
    "article": "der"
  },
  {
    "noun": "Seelenhirt",
    "article": "der"
  },
  {
    "noun": "Seelenzustand",
    "article": "der"
  },
  {
    "noun": "Seeleopard",
    "article": "der"
  },
  {
    "noun": "Seelsorger",
    "article": "der"
  },
  {
    "noun": "Seemann",
    "article": "der"
  },
  {
    "noun": "Seenotrettungsdienst",
    "article": "der"
  },
  {
    "noun": "Seeoffizier",
    "article": "der"
  },
  {
    "noun": "Seeraub",
    "article": "der"
  },
  {
    "noun": "Seersucker",
    "article": "der"
  },
  {
    "noun": "Seesack",
    "article": "der"
  },
  {
    "noun": "Seesand",
    "article": "der"
  },
  {
    "noun": "Seeschaden",
    "article": "der"
  },
  {
    "noun": "Seestern",
    "article": "der"
  },
  {
    "noun": "Seetang",
    "article": "der"
  },
  {
    "noun": "Seetaucher",
    "article": "der"
  },
  {
    "noun": "Seeteufel",
    "article": "der"
  },
  {
    "noun": "Seetransport",
    "article": "der"
  },
  {
    "noun": "Seeunfall",
    "article": "der"
  },
  {
    "noun": "Seeverkehr",
    "article": "der"
  },
  {
    "noun": "Seeweg",
    "article": "der"
  },
  {
    "noun": "Seewetterdienst",
    "article": "der"
  },
  {
    "noun": "Seewind",
    "article": "der"
  },
  {
    "noun": "Seewolf",
    "article": "der"
  },
  {
    "noun": "Segelclub",
    "article": "der"
  },
  {
    "noun": "Segelflieger",
    "article": "der"
  },
  {
    "noun": "Segelflosser",
    "article": "der"
  },
  {
    "noun": "Segelflug",
    "article": "der"
  },
  {
    "noun": "Segelmacher",
    "article": "der"
  },
  {
    "noun": "Segelsport",
    "article": "der"
  },
  {
    "noun": "Segen",
    "article": "der"
  },
  {
    "noun": "Segerkegel",
    "article": "der"
  },
  {
    "noun": "Segler",
    "article": "der"
  },
  {
    "noun": "Seher",
    "article": "der"
  },
  {
    "noun": "Seherblick",
    "article": "der"
  },
  {
    "noun": "Sehfehler",
    "article": "der"
  },
  {
    "noun": "Sehkreis",
    "article": "der"
  },
  {
    "noun": "Sehnerv",
    "article": "der"
  },
  {
    "noun": "Sehschlitz",
    "article": "der"
  },
  {
    "noun": "Sehtest",
    "article": "der"
  },
  {
    "noun": "Sehwinkel",
    "article": "der"
  },
  {
    "noun": "Seidelbast",
    "article": "der"
  },
  {
    "noun": "Seidenatlas",
    "article": "der"
  },
  {
    "noun": "Seidenbau",
    "article": "der"
  },
  {
    "noun": "Seidenbrokat",
    "article": "der"
  },
  {
    "noun": "Seidenfaden",
    "article": "der"
  },
  {
    "noun": "Seidenglanz",
    "article": "der"
  },
  {
    "noun": "Seidenreiher",
    "article": "der"
  },
  {
    "noun": "Seidenschal",
    "article": "der"
  },
  {
    "noun": "Seidenschwanz",
    "article": "der"
  },
  {
    "noun": "Seidenspinner",
    "article": "der"
  },
  {
    "noun": "Seidenstoff",
    "article": "der"
  },
  {
    "noun": "Seidenstrumpf",
    "article": "der"
  },
  {
    "noun": "Seifenschaum",
    "article": "der"
  },
  {
    "noun": "Seifensieder",
    "article": "der"
  },
  {
    "noun": "Seiher",
    "article": "der"
  },
  {
    "noun": "Seiler",
    "article": "der"
  },
  {
    "noun": "Seiltanz",
    "article": "der"
  },
  {
    "noun": "Seilzug",
    "article": "der"
  },
  {
    "noun": "Seismiker",
    "article": "der"
  },
  {
    "noun": "Seismograf",
    "article": "der"
  },
  {
    "noun": "Seismologe",
    "article": "der"
  },
  {
    "noun": "Seitenangriff",
    "article": "der"
  },
  {
    "noun": "Seitenarm",
    "article": "der"
  },
  {
    "noun": "Seitenaufprallschutz",
    "article": "der"
  },
  {
    "noun": "Seitenausgang",
    "article": "der"
  },
  {
    "noun": "Seitenbau",
    "article": "der"
  },
  {
    "noun": "Seitenblick",
    "article": "der"
  },
  {
    "noun": "Seitendruck",
    "article": "der"
  },
  {
    "noun": "Seiteneingang",
    "article": "der"
  },
  {
    "noun": "Seitenhieb",
    "article": "der"
  },
  {
    "noun": "Seitenkanal",
    "article": "der"
  },
  {
    "noun": "Seitenriss",
    "article": "der"
  },
  {
    "noun": "Seitenscheitel",
    "article": "der"
  },
  {
    "noun": "Seitenschneider",
    "article": "der"
  },
  {
    "noun": "Seitenschritt",
    "article": "der"
  },
  {
    "noun": "Seitensprung",
    "article": "der"
  },
  {
    "noun": "Seitenstreifen",
    "article": "der"
  },
  {
    "noun": "Seitenteil",
    "article": "der"
  },
  {
    "noun": "Seitentrieb",
    "article": "der"
  },
  {
    "noun": "Seitenwagen",
    "article": "der"
  },
  {
    "noun": "Seitenwechsel",
    "article": "der"
  },
  {
    "noun": "Seitenweg",
    "article": "der"
  },
  {
    "noun": "Seitenwind",
    "article": "der"
  },
  {
    "noun": "Sekans",
    "article": "der"
  },
  {
    "noun": "Sekt",
    "article": "der"
  },
  {
    "noun": "Sektierer",
    "article": "der"
  },
  {
    "noun": "Sektor",
    "article": "der"
  },
  {
    "noun": "Sekundararzt",
    "article": "der"
  },
  {
    "noun": "Sekundenzeiger",
    "article": "der"
  },
  {
    "noun": "Selbstanschluss",
    "article": "der"
  },
  {
    "noun": "Selbstbedarf",
    "article": "der"
  },
  {
    "noun": "Selbstbedienungsladen",
    "article": "der"
  },
  {
    "noun": "Selbstbehalt",
    "article": "der"
  },
  {
    "noun": "Selbstbetrug",
    "article": "der"
  },
  {
    "noun": "Selbstdarsteller",
    "article": "der"
  },
  {
    "noun": "Selbsteintritt",
    "article": "der"
  },
  {
    "noun": "Selbstentlader",
    "article": "der"
  },
  {
    "noun": "Selbsterhaltungstrieb",
    "article": "der"
  },
  {
    "noun": "Selbstfahrer",
    "article": "der"
  },
  {
    "noun": "Selbstkostenpreis",
    "article": "der"
  },
  {
    "noun": "Selbstlader",
    "article": "der"
  },
  {
    "noun": "Selbstlaut",
    "article": "der"
  },
  {
    "noun": "Selbstschalter",
    "article": "der"
  },
  {
    "noun": "Selbstschreiber",
    "article": "der"
  },
  {
    "noun": "Selbstschuss",
    "article": "der"
  },
  {
    "noun": "Selbstschutz",
    "article": "der"
  },
  {
    "noun": "Selbstunterricht",
    "article": "der"
  },
  {
    "noun": "Selbstverlag",
    "article": "der"
  },
  {
    "noun": "Selbstversorger",
    "article": "der"
  },
  {
    "noun": "Selbstversuch",
    "article": "der"
  },
  {
    "noun": "Selbstvorwurf",
    "article": "der"
  },
  {
    "noun": "Selbstzahler",
    "article": "der"
  },
  {
    "noun": "Selbstzweck",
    "article": "der"
  },
  {
    "noun": "Selleriesalat",
    "article": "der"
  },
  {
    "noun": "Seltenheitswert",
    "article": "der"
  },
  {
    "noun": "Semantiker",
    "article": "der"
  },
  {
    "noun": "Seminarist",
    "article": "der"
  },
  {
    "noun": "Semit",
    "article": "der"
  },
  {
    "noun": "Semivokal",
    "article": "der"
  },
  {
    "noun": "Senat",
    "article": "der"
  },
  {
    "noun": "Senator",
    "article": "der"
  },
  {
    "noun": "Senatsbeschluss",
    "article": "der"
  },
  {
    "noun": "Sendbote",
    "article": "der"
  },
  {
    "noun": "Sendebereich",
    "article": "der"
  },
  {
    "noun": "Sendeleiter",
    "article": "der"
  },
  {
    "noun": "Sendemast",
    "article": "der"
  },
  {
    "noun": "Sendeplan",
    "article": "der"
  },
  {
    "noun": "Sendeplatz",
    "article": "der"
  },
  {
    "noun": "Sender",
    "article": "der"
  },
  {
    "noun": "Senderaum",
    "article": "der"
  },
  {
    "noun": "Sendeschluss",
    "article": "der"
  },
  {
    "noun": "Sendeturm",
    "article": "der"
  },
  {
    "noun": "Senegaler",
    "article": "der"
  },
  {
    "noun": "Senegalese",
    "article": "der"
  },
  {
    "noun": "Senf",
    "article": "der"
  },
  {
    "noun": "Senftopf",
    "article": "der"
  },
  {
    "noun": "Senior",
    "article": "der"
  },
  {
    "noun": "Seniorchef",
    "article": "der"
  },
  {
    "noun": "Senkel",
    "article": "der"
  },
  {
    "noun": "Senker",
    "article": "der"
  },
  {
    "noun": "Senkkasten",
    "article": "der"
  },
  {
    "noun": "Senkrechtstart",
    "article": "der"
  },
  {
    "noun": "Senkrechtstarter",
    "article": "der"
  },
  {
    "noun": "Senn",
    "article": "der"
  },
  {
    "noun": "Senner",
    "article": "der"
  },
  {
    "noun": "Sensenbaum",
    "article": "der"
  },
  {
    "noun": "Sensenmann",
    "article": "der"
  },
  {
    "noun": "Sensenstein",
    "article": "der"
  },
  {
    "noun": "Sensibilisator",
    "article": "der"
  },
  {
    "noun": "Sensibilismus",
    "article": "der"
  },
  {
    "noun": "Sensor",
    "article": "der"
  },
  {
    "noun": "Sensualismus",
    "article": "der"
  },
  {
    "noun": "Sensualist",
    "article": "der"
  },
  {
    "noun": "Separatabdruck",
    "article": "der"
  },
  {
    "noun": "Separateingang",
    "article": "der"
  },
  {
    "noun": "Separatismus",
    "article": "der"
  },
  {
    "noun": "Separatist",
    "article": "der"
  },
  {
    "noun": "Separator",
    "article": "der"
  },
  {
    "noun": "September",
    "article": "der"
  },
  {
    "noun": "Sequenzer",
    "article": "der"
  },
  {
    "noun": "Seraph",
    "article": "der"
  },
  {
    "noun": "Serbe",
    "article": "der"
  },
  {
    "noun": "Serienbau",
    "article": "der"
  },
  {
    "noun": "Serienschalter",
    "article": "der"
  },
  {
    "noun": "Serienwagen",
    "article": "der"
  },
  {
    "noun": "Sermon",
    "article": "der"
  },
  {
    "noun": "Serval",
    "article": "der"
  },
  {
    "noun": "Server",
    "article": "der"
  },
  {
    "noun": "Serviertisch",
    "article": "der"
  },
  {
    "noun": "Servierwagen",
    "article": "der"
  },
  {
    "noun": "Serviettenring",
    "article": "der"
  },
  {
    "noun": "Servilismus",
    "article": "der"
  },
  {
    "noun": "Servomotor",
    "article": "der"
  },
  {
    "noun": "Sesam",
    "article": "der"
  },
  {
    "noun": "Sessel",
    "article": "der"
  },
  {
    "noun": "Sessellift",
    "article": "der"
  },
  {
    "noun": "Setter",
    "article": "der"
  },
  {
    "noun": "Setzer",
    "article": "der"
  },
  {
    "noun": "Setzhammer",
    "article": "der"
  },
  {
    "noun": "Setzkasten",
    "article": "der"
  },
  {
    "noun": "Setzling",
    "article": "der"
  },
  {
    "noun": "Seuchenherd",
    "article": "der"
  },
  {
    "noun": "Seufzer",
    "article": "der"
  },
  {
    "noun": "Sezessionist",
    "article": "der"
  },
  {
    "noun": "Sezessionskrieg",
    "article": "der"
  },
  {
    "noun": "Shag",
    "article": "der"
  },
  {
    "noun": "Sheriff",
    "article": "der"
  },
  {
    "noun": "Sherpa",
    "article": "der"
  },
  {
    "noun": "Sherry",
    "article": "der"
  },
  {
    "noun": "Shilling",
    "article": "der"
  },
  {
    "noun": "Shintoismus",
    "article": "der"
  },
  {
    "noun": "Showmaster",
    "article": "der"
  },
  {
    "noun": "Shunt",
    "article": "der"
  },
  {
    "noun": "Sibilant",
    "article": "der"
  },
  {
    "noun": "Sicherheitsabstand",
    "article": "der"
  },
  {
    "noun": "Sicherheitsdienst",
    "article": "der"
  },
  {
    "noun": "Sicherheitsfaktor",
    "article": "der"
  },
  {
    "noun": "Sicherheitsgurt",
    "article": "der"
  },
  {
    "noun": "Sicherheitsingenieur",
    "article": "der"
  },
  {
    "noun": "Sicherheitsrat",
    "article": "der"
  },
  {
    "noun": "Sicherheitsverschluss",
    "article": "der"
  },
  {
    "noun": "Sicherungsautomat",
    "article": "der"
  },
  {
    "noun": "Sicherungskasten",
    "article": "der"
  },
  {
    "noun": "Sicherungsnehmer",
    "article": "der"
  },
  {
    "noun": "Sichtbeton",
    "article": "der"
  },
  {
    "noun": "Sichtschutz",
    "article": "der"
  },
  {
    "noun": "Sichtvermerk",
    "article": "der"
  },
  {
    "noun": "Sichtwechsel",
    "article": "der"
  },
  {
    "noun": "Sickerschacht",
    "article": "der"
  },
  {
    "noun": "Siderit",
    "article": "der"
  },
  {
    "noun": "Siebdruck",
    "article": "der"
  },
  {
    "noun": "Siebdrucker",
    "article": "der"
  },
  {
    "noun": "Siebener",
    "article": "der"
  },
  {
    "noun": "Siebenkampf",
    "article": "der"
  },
  {
    "noun": "Siebenmeter",
    "article": "der"
  },
  {
    "noun": "Siebenpunkt",
    "article": "der"
  },
  {
    "noun": "Siebmacher",
    "article": "der"
  },
  {
    "noun": "Siedegrad",
    "article": "der"
  },
  {
    "noun": "Siedepunkt",
    "article": "der"
  },
  {
    "noun": "Sieder",
    "article": "der"
  },
  {
    "noun": "Siedewasserreaktor",
    "article": "der"
  },
  {
    "noun": "Sieg",
    "article": "der"
  },
  {
    "noun": "Siegelabdruck",
    "article": "der"
  },
  {
    "noun": "Siegelbewahrer",
    "article": "der"
  },
  {
    "noun": "Siegelbruch",
    "article": "der"
  },
  {
    "noun": "Siegellack",
    "article": "der"
  },
  {
    "noun": "Siegelring",
    "article": "der"
  },
  {
    "noun": "Sieger",
    "article": "der"
  },
  {
    "noun": "Siegeslauf",
    "article": "der"
  },
  {
    "noun": "Siegespreis",
    "article": "der"
  },
  {
    "noun": "Siegesrausch",
    "article": "der"
  },
  {
    "noun": "Siegestaumel",
    "article": "der"
  },
  {
    "noun": "Siegeswille",
    "article": "der"
  },
  {
    "noun": "Siegeszug",
    "article": "der"
  },
  {
    "noun": "Signalmast",
    "article": "der"
  },
  {
    "noun": "Signalton",
    "article": "der"
  },
  {
    "noun": "Signatar",
    "article": "der"
  },
  {
    "noun": "Signifikant",
    "article": "der"
  },
  {
    "noun": "Signifikanztest",
    "article": "der"
  },
  {
    "noun": "Sigrist",
    "article": "der"
  },
  {
    "noun": "Silbenstecher",
    "article": "der"
  },
  {
    "noun": "Silberbarren",
    "article": "der"
  },
  {
    "noun": "Silberbeschlag",
    "article": "der"
  },
  {
    "noun": "Silberbisam",
    "article": "der"
  },
  {
    "noun": "Silberblick",
    "article": "der"
  },
  {
    "noun": "Silberbrokat",
    "article": "der"
  },
  {
    "noun": "Silberdraht",
    "article": "der"
  },
  {
    "noun": "Silberfaden",
    "article": "der"
  },
  {
    "noun": "Silberfisch",
    "article": "der"
  },
  {
    "noun": "Silberfuchs",
    "article": "der"
  },
  {
    "noun": "Silberglanz",
    "article": "der"
  },
  {
    "noun": "Silberklang",
    "article": "der"
  },
  {
    "noun": "Silberreiher",
    "article": "der"
  },
  {
    "noun": "Silberring",
    "article": "der"
  },
  {
    "noun": "Silberschmied",
    "article": "der"
  },
  {
    "noun": "Silberschmuck",
    "article": "der"
  },
  {
    "noun": "Silberstreifen",
    "article": "der"
  },
  {
    "noun": "Silomais",
    "article": "der"
  },
  {
    "noun": "Silvaner",
    "article": "der"
  },
  {
    "noun": "Silvesterabend",
    "article": "der"
  },
  {
    "noun": "Simbabwer",
    "article": "der"
  },
  {
    "noun": "Simshobel",
    "article": "der"
  },
  {
    "noun": "Simulant",
    "article": "der"
  },
  {
    "noun": "Simulator",
    "article": "der"
  },
  {
    "noun": "Simultandolmetscher",
    "article": "der"
  },
  {
    "noun": "Sinfoniker",
    "article": "der"
  },
  {
    "noun": "Singleton",
    "article": "der"
  },
  {
    "noun": "Singsang",
    "article": "der"
  },
  {
    "noun": "Singschwan",
    "article": "der"
  },
  {
    "noun": "Singular",
    "article": "der"
  },
  {
    "noun": "Singultus",
    "article": "der"
  },
  {
    "noun": "Singvogel",
    "article": "der"
  },
  {
    "noun": "Sinn",
    "article": "der"
  },
  {
    "noun": "Sinnengenuss",
    "article": "der"
  },
  {
    "noun": "Sinnenmensch",
    "article": "der"
  },
  {
    "noun": "Sinnenrausch",
    "article": "der"
  },
  {
    "noun": "Sinnenreiz",
    "article": "der"
  },
  {
    "noun": "Sinnentaumel",
    "article": "der"
  },
  {
    "noun": "Sinneseindruck",
    "article": "der"
  },
  {
    "noun": "Sinnesreiz",
    "article": "der"
  },
  {
    "noun": "Sinneswandel",
    "article": "der"
  },
  {
    "noun": "Sinnierer",
    "article": "der"
  },
  {
    "noun": "Sinnspruch",
    "article": "der"
  },
  {
    "noun": "Sinologe",
    "article": "der"
  },
  {
    "noun": "Sinter",
    "article": "der"
  },
  {
    "noun": "Sinus",
    "article": "der"
  },
  {
    "noun": "Sinussatz",
    "article": "der"
  },
  {
    "noun": "Sinuston",
    "article": "der"
  },
  {
    "noun": "Sioux",
    "article": "der"
  },
  {
    "noun": "Siphon",
    "article": "der"
  },
  {
    "noun": "Sirenengesang",
    "article": "der"
  },
  {
    "noun": "Sirius",
    "article": "der"
  },
  {
    "noun": "Sirtaki",
    "article": "der"
  },
  {
    "noun": "Sirup",
    "article": "der"
  },
  {
    "noun": "Sisal",
    "article": "der"
  },
  {
    "noun": "Sitar",
    "article": "der"
  },
  {
    "noun": "Sittenkodex",
    "article": "der"
  },
  {
    "noun": "Sittenprediger",
    "article": "der"
  },
  {
    "noun": "Sittenrichter",
    "article": "der"
  },
  {
    "noun": "Sittenroman",
    "article": "der"
  },
  {
    "noun": "Sittenstrolch",
    "article": "der"
  },
  {
    "noun": "Sittenverfall",
    "article": "der"
  },
  {
    "noun": "Sittich",
    "article": "der"
  },
  {
    "noun": "Sittlichkeitsverbrecher",
    "article": "der"
  },
  {
    "noun": "Sitz",
    "article": "der"
  },
  {
    "noun": "Sitzplatz",
    "article": "der"
  },
  {
    "noun": "Sitzstreik",
    "article": "der"
  },
  {
    "noun": "Sitzungsbericht",
    "article": "der"
  },
  {
    "noun": "Sitzungsraum",
    "article": "der"
  },
  {
    "noun": "Sitzungssaal",
    "article": "der"
  },
  {
    "noun": "Sizilianer",
    "article": "der"
  },
  {
    "noun": "Skalde",
    "article": "der"
  },
  {
    "noun": "Skalp",
    "article": "der"
  },
  {
    "noun": "Skandal",
    "article": "der"
  },
  {
    "noun": "Skandinavier",
    "article": "der"
  },
  {
    "noun": "Skat",
    "article": "der"
  },
  {
    "noun": "Skater",
    "article": "der"
  },
  {
    "noun": "Skatspieler",
    "article": "der"
  },
  {
    "noun": "Skeleton",
    "article": "der"
  },
  {
    "noun": "Skelettbau",
    "article": "der"
  },
  {
    "noun": "Skeptiker",
    "article": "der"
  },
  {
    "noun": "Skeptizismus",
    "article": "der"
  },
  {
    "noun": "Sketch",
    "article": "der"
  },
  {
    "noun": "Ski",
    "article": "der"
  },
  {
    "noun": "Skianzug",
    "article": "der"
  },
  {
    "noun": "Skibob",
    "article": "der"
  },
  {
    "noun": "Skicircus",
    "article": "der"
  },
  {
    "noun": "Skifahrer",
    "article": "der"
  },
  {
    "noun": "Skiflug",
    "article": "der"
  },
  {
    "noun": "Skikurs",
    "article": "der"
  },
  {
    "noun": "Skilanglauf",
    "article": "der"
  },
  {
    "noun": "Skilauf",
    "article": "der"
  },
  {
    "noun": "Skilehrer",
    "article": "der"
  },
  {
    "noun": "Skilift",
    "article": "der"
  },
  {
    "noun": "Skin",
    "article": "der"
  },
  {
    "noun": "Skinhead",
    "article": "der"
  },
  {
    "noun": "Skink",
    "article": "der"
  },
  {
    "noun": "Skipass",
    "article": "der"
  },
  {
    "noun": "Skipper",
    "article": "der"
  },
  {
    "noun": "Skischuh",
    "article": "der"
  },
  {
    "noun": "Skisport",
    "article": "der"
  },
  {
    "noun": "Skispringer",
    "article": "der"
  },
  {
    "noun": "Skisprung",
    "article": "der"
  },
  {
    "noun": "Skistiefel",
    "article": "der"
  },
  {
    "noun": "Skistock",
    "article": "der"
  },
  {
    "noun": "Skiurlaub",
    "article": "der"
  },
  {
    "noun": "Skorbut",
    "article": "der"
  },
  {
    "noun": "Skorpion",
    "article": "der"
  },
  {
    "noun": "Skribent",
    "article": "der"
  },
  {
    "noun": "Skrupel",
    "article": "der"
  },
  {
    "noun": "Skulpteur",
    "article": "der"
  },
  {
    "noun": "Skunk",
    "article": "der"
  },
  {
    "noun": "Slalom",
    "article": "der"
  },
  {
    "noun": "Slang",
    "article": "der"
  },
  {
    "noun": "Slapstick",
    "article": "der"
  },
  {
    "noun": "Slawe",
    "article": "der"
  },
  {
    "noun": "Slawismus",
    "article": "der"
  },
  {
    "noun": "Slawist",
    "article": "der"
  },
  {
    "noun": "Slawonier",
    "article": "der"
  },
  {
    "noun": "Slibowitz",
    "article": "der"
  },
  {
    "noun": "Slick",
    "article": "der"
  },
  {
    "noun": "Slip",
    "article": "der"
  },
  {
    "noun": "Sliwowitz",
    "article": "der"
  },
  {
    "noun": "Slogan",
    "article": "der"
  },
  {
    "noun": "Slot",
    "article": "der"
  },
  {
    "noun": "Slowake",
    "article": "der"
  },
  {
    "noun": "Slowene",
    "article": "der"
  },
  {
    "noun": "Slum",
    "article": "der"
  },
  {
    "noun": "Smaragd",
    "article": "der"
  },
  {
    "noun": "Smog",
    "article": "der"
  },
  {
    "noun": "Smogalarm",
    "article": "der"
  },
  {
    "noun": "Smoking",
    "article": "der"
  },
  {
    "noun": "Snack",
    "article": "der"
  },
  {
    "noun": "Snob",
    "article": "der"
  },
  {
    "noun": "Snobismus",
    "article": "der"
  },
  {
    "noun": "Sockel",
    "article": "der"
  },
  {
    "noun": "Sockelbetrag",
    "article": "der"
  },
  {
    "noun": "Sockenhalter",
    "article": "der"
  },
  {
    "noun": "Softball",
    "article": "der"
  },
  {
    "noun": "Softie",
    "article": "der"
  },
  {
    "noun": "Sog",
    "article": "der"
  },
  {
    "noun": "Sohn",
    "article": "der"
  },
  {
    "noun": "Sokratiker",
    "article": "der"
  },
  {
    "noun": "Solarplexus",
    "article": "der"
  },
  {
    "noun": "Solawechsel",
    "article": "der"
  },
  {
    "noun": "Sold",
    "article": "der"
  },
  {
    "noun": "Soldat",
    "article": "der"
  },
  {
    "noun": "Soldatenfriedhof",
    "article": "der"
  },
  {
    "noun": "Solidarbeitrag",
    "article": "der"
  },
  {
    "noun": "Solidarpakt",
    "article": "der"
  },
  {
    "noun": "Solidarschuldner",
    "article": "der"
  },
  {
    "noun": "Solipsismus",
    "article": "der"
  },
  {
    "noun": "Solist",
    "article": "der"
  },
  {
    "noun": "Sologesang",
    "article": "der"
  },
  {
    "noun": "Solotanz",
    "article": "der"
  },
  {
    "noun": "Somalier",
    "article": "der"
  },
  {
    "noun": "Sombrero",
    "article": "der"
  },
  {
    "noun": "Sommelier",
    "article": "der"
  },
  {
    "noun": "Sommer",
    "article": "der"
  },
  {
    "noun": "Sommerabend",
    "article": "der"
  },
  {
    "noun": "Sommeranfang",
    "article": "der"
  },
  {
    "noun": "Sommeranzug",
    "article": "der"
  },
  {
    "noun": "Sommeraufenthalt",
    "article": "der"
  },
  {
    "noun": "Sommerfahrplan",
    "article": "der"
  },
  {
    "noun": "Sommerfrischler",
    "article": "der"
  },
  {
    "noun": "Sommermonat",
    "article": "der"
  },
  {
    "noun": "Sommerregen",
    "article": "der"
  },
  {
    "noun": "Sommerreifen",
    "article": "der"
  },
  {
    "noun": "Sommersanfang",
    "article": "der"
  },
  {
    "noun": "Sommerschlaf",
    "article": "der"
  },
  {
    "noun": "Sommerschlussverkauf",
    "article": "der"
  },
  {
    "noun": "Sommertag",
    "article": "der"
  },
  {
    "noun": "Sommerurlaub",
    "article": "der"
  },
  {
    "noun": "Sommervogel",
    "article": "der"
  },
  {
    "noun": "Sommerweizen",
    "article": "der"
  },
  {
    "noun": "Somnambulismus",
    "article": "der"
  },
  {
    "noun": "Sonderabdruck",
    "article": "der"
  },
  {
    "noun": "Sonderabfall",
    "article": "der"
  },
  {
    "noun": "Sonderauftrag",
    "article": "der"
  },
  {
    "noun": "Sonderbeitrag",
    "article": "der"
  },
  {
    "noun": "Sonderberichterstatter",
    "article": "der"
  },
  {
    "noun": "Sonderbotschafter",
    "article": "der"
  },
  {
    "noun": "Sonderbund",
    "article": "der"
  },
  {
    "noun": "Sonderdruck",
    "article": "der"
  },
  {
    "noun": "Sonderfall",
    "article": "der"
  },
  {
    "noun": "Sonderfriede",
    "article": "der"
  },
  {
    "noun": "Sonderling",
    "article": "der"
  },
  {
    "noun": "Sonderpreis",
    "article": "der"
  },
  {
    "noun": "Sonderschullehrer",
    "article": "der"
  },
  {
    "noun": "Sonderstempel",
    "article": "der"
  },
  {
    "noun": "Sonderurlaub",
    "article": "der"
  },
  {
    "noun": "Sonderverkauf",
    "article": "der"
  },
  {
    "noun": "Sonderwunsch",
    "article": "der"
  },
  {
    "noun": "Sonderzug",
    "article": "der"
  },
  {
    "noun": "Song",
    "article": "der"
  },
  {
    "noun": "Sonnabend",
    "article": "der"
  },
  {
    "noun": "Sonnenanbeter",
    "article": "der"
  },
  {
    "noun": "Sonnenaufgang",
    "article": "der"
  },
  {
    "noun": "Sonnenbrand",
    "article": "der"
  },
  {
    "noun": "Sonnenbrenner",
    "article": "der"
  },
  {
    "noun": "Sonnenfleck",
    "article": "der"
  },
  {
    "noun": "Sonnenglanz",
    "article": "der"
  },
  {
    "noun": "Sonnengott",
    "article": "der"
  },
  {
    "noun": "Sonnenhut",
    "article": "der"
  },
  {
    "noun": "Sonnenkollektor",
    "article": "der"
  },
  {
    "noun": "Sonnenkult",
    "article": "der"
  },
  {
    "noun": "Sonnenschein",
    "article": "der"
  },
  {
    "noun": "Sonnenschirm",
    "article": "der"
  },
  {
    "noun": "Sonnenschutz",
    "article": "der"
  },
  {
    "noun": "Sonnenstand",
    "article": "der"
  },
  {
    "noun": "Sonnenstich",
    "article": "der"
  },
  {
    "noun": "Sonnenstrahl",
    "article": "der"
  },
  {
    "noun": "Sonnentag",
    "article": "der"
  },
  {
    "noun": "Sonnenuntergang",
    "article": "der"
  },
  {
    "noun": "Sonnenvogel",
    "article": "der"
  },
  {
    "noun": "Sonnenwind",
    "article": "der"
  },
  {
    "noun": "Sonntag",
    "article": "der"
  },
  {
    "noun": "Sonntagsanzug",
    "article": "der"
  },
  {
    "noun": "Sonntagsausflug",
    "article": "der"
  },
  {
    "noun": "Sonntagsbraten",
    "article": "der"
  },
  {
    "noun": "Sonntagsfahrer",
    "article": "der"
  },
  {
    "noun": "Sonntagsgottesdienst",
    "article": "der"
  },
  {
    "noun": "Sonntagsmaler",
    "article": "der"
  },
  {
    "noun": "Sonntagsstaat",
    "article": "der"
  },
  {
    "noun": "Soor",
    "article": "der"
  },
  {
    "noun": "Sophismus",
    "article": "der"
  },
  {
    "noun": "Sophist",
    "article": "der"
  },
  {
    "noun": "Sopran",
    "article": "der"
  },
  {
    "noun": "Sorbe",
    "article": "der"
  },
  {
    "noun": "Sortenhandel",
    "article": "der"
  },
  {
    "noun": "Sortierer",
    "article": "der"
  },
  {
    "noun": "Sortimenter",
    "article": "der"
  },
  {
    "noun": "Souffleur",
    "article": "der"
  },
  {
    "noun": "Souffleurkasten",
    "article": "der"
  },
  {
    "noun": "Sound",
    "article": "der"
  },
  {
    "noun": "Souvenirladen",
    "article": "der"
  },
  {
    "noun": "Sowjet",
    "article": "der"
  },
  {
    "noun": "Sozialarbeiter",
    "article": "der"
  },
  {
    "noun": "Sozialbeitrag",
    "article": "der"
  },
  {
    "noun": "Sozialbericht",
    "article": "der"
  },
  {
    "noun": "Sozialdarwinismus",
    "article": "der"
  },
  {
    "noun": "Sozialdemokrat",
    "article": "der"
  },
  {
    "noun": "Sozialfonds",
    "article": "der"
  },
  {
    "noun": "Sozialimperialismus",
    "article": "der"
  },
  {
    "noun": "Sozialismus",
    "article": "der"
  },
  {
    "noun": "Sozialist",
    "article": "der"
  },
  {
    "noun": "Soziallohn",
    "article": "der"
  },
  {
    "noun": "Sozialpartner",
    "article": "der"
  },
  {
    "noun": "Sozialplan",
    "article": "der"
  },
  {
    "noun": "Sozialstaat",
    "article": "der"
  },
  {
    "noun": "Sozialversicherungsausweis",
    "article": "der"
  },
  {
    "noun": "Sozialversicherungsbeitrag",
    "article": "der"
  },
  {
    "noun": "Soziologe",
    "article": "der"
  },
  {
    "noun": "Soziopath",
    "article": "der"
  },
  {
    "noun": "Sozius",
    "article": "der"
  },
  {
    "noun": "Soziusfahrer",
    "article": "der"
  },
  {
    "noun": "Soziussitz",
    "article": "der"
  },
  {
    "noun": "Spalierbaum",
    "article": "der"
  },
  {
    "noun": "Spalt",
    "article": "der"
  },
  {
    "noun": "Spaltpilz",
    "article": "der"
  },
  {
    "noun": "Span",
    "article": "der"
  },
  {
    "noun": "Spaniel",
    "article": "der"
  },
  {
    "noun": "Spanier",
    "article": "der"
  },
  {
    "noun": "Spaniole",
    "article": "der"
  },
  {
    "noun": "Spankorb",
    "article": "der"
  },
  {
    "noun": "Spann",
    "article": "der"
  },
  {
    "noun": "Spannbeton",
    "article": "der"
  },
  {
    "noun": "Spanndienst",
    "article": "der"
  },
  {
    "noun": "Spanner",
    "article": "der"
  },
  {
    "noun": "Spannstahl",
    "article": "der"
  },
  {
    "noun": "Spannteppich",
    "article": "der"
  },
  {
    "noun": "Spannungsabfall",
    "article": "der"
  },
  {
    "noun": "Spannungsbogen",
    "article": "der"
  },
  {
    "noun": "Spannungsfall",
    "article": "der"
  },
  {
    "noun": "Spannungsmesser",
    "article": "der"
  },
  {
    "noun": "Spannungsregler",
    "article": "der"
  },
  {
    "noun": "Spannungsstabilisator",
    "article": "der"
  },
  {
    "noun": "Spannungssucher",
    "article": "der"
  },
  {
    "noun": "Spannungsteiler",
    "article": "der"
  },
  {
    "noun": "Spannungsverlust",
    "article": "der"
  },
  {
    "noun": "Spannungswandler",
    "article": "der"
  },
  {
    "noun": "Spannungszustand",
    "article": "der"
  },
  {
    "noun": "Sparbetrag",
    "article": "der"
  },
  {
    "noun": "Sparer",
    "article": "der"
  },
  {
    "noun": "Sparerfreibetrag",
    "article": "der"
  },
  {
    "noun": "Spargel",
    "article": "der"
  },
  {
    "noun": "Spargelstecher",
    "article": "der"
  },
  {
    "noun": "Spargroschen",
    "article": "der"
  },
  {
    "noun": "Sparpreis",
    "article": "der"
  },
  {
    "noun": "Sparring",
    "article": "der"
  },
  {
    "noun": "Sparringkampf",
    "article": "der"
  },
  {
    "noun": "Sparringpartner",
    "article": "der"
  },
  {
    "noun": "Sparstrumpf",
    "article": "der"
  },
  {
    "noun": "Spartakus",
    "article": "der"
  },
  {
    "noun": "Sparvertrag",
    "article": "der"
  },
  {
    "noun": "Spasmus",
    "article": "der"
  },
  {
    "noun": "Spastiker",
    "article": "der"
  },
  {
    "noun": "Spat",
    "article": "der"
  },
  {
    "noun": "Spatenstich",
    "article": "der"
  },
  {
    "noun": "Spatz",
    "article": "der"
  },
  {
    "noun": "Spaziergang",
    "article": "der"
  },
  {
    "noun": "Spazierstock",
    "article": "der"
  },
  {
    "noun": "Spazierweg",
    "article": "der"
  },
  {
    "noun": "Specht",
    "article": "der"
  },
  {
    "noun": "Speck",
    "article": "der"
  },
  {
    "noun": "Speckbauch",
    "article": "der"
  },
  {
    "noun": "Speckkuchen",
    "article": "der"
  },
  {
    "noun": "Speckstein",
    "article": "der"
  },
  {
    "noun": "Spediteur",
    "article": "der"
  },
  {
    "noun": "Speditionskaufmann",
    "article": "der"
  },
  {
    "noun": "Speer",
    "article": "der"
  },
  {
    "noun": "Speerschaft",
    "article": "der"
  },
  {
    "noun": "Speerwerfer",
    "article": "der"
  },
  {
    "noun": "Speichel",
    "article": "der"
  },
  {
    "noun": "Speichelfluss",
    "article": "der"
  },
  {
    "noun": "Speichellecker",
    "article": "der"
  },
  {
    "noun": "Speichennippel",
    "article": "der"
  },
  {
    "noun": "Speichenreflektor",
    "article": "der"
  },
  {
    "noun": "Speicher",
    "article": "der"
  },
  {
    "noun": "Speicherofen",
    "article": "der"
  },
  {
    "noun": "Speicherplatz",
    "article": "der"
  },
  {
    "noun": "Speil",
    "article": "der"
  },
  {
    "noun": "Speisebrei",
    "article": "der"
  },
  {
    "noun": "Speisefisch",
    "article": "der"
  },
  {
    "noun": "Speisenaufzug",
    "article": "der"
  },
  {
    "noun": "Speisepilz",
    "article": "der"
  },
  {
    "noun": "Speiseplan",
    "article": "der"
  },
  {
    "noun": "Speisequark",
    "article": "der"
  },
  {
    "noun": "Speiseraum",
    "article": "der"
  },
  {
    "noun": "Speiserest",
    "article": "der"
  },
  {
    "noun": "Speisesaal",
    "article": "der"
  },
  {
    "noun": "Speiseschrank",
    "article": "der"
  },
  {
    "noun": "Speisewagen",
    "article": "der"
  },
  {
    "noun": "Speisezettel",
    "article": "der"
  },
  {
    "noun": "Spektralapparat",
    "article": "der"
  },
  {
    "noun": "Spektrograf",
    "article": "der"
  },
  {
    "noun": "Spekulant",
    "article": "der"
  },
  {
    "noun": "Spekulationsgewinn",
    "article": "der"
  },
  {
    "noun": "Spekulatius",
    "article": "der"
  },
  {
    "noun": "Spelz",
    "article": "der"
  },
  {
    "noun": "Spendenaufruf",
    "article": "der"
  },
  {
    "noun": "Spendensammler",
    "article": "der"
  },
  {
    "noun": "Spender",
    "article": "der"
  },
  {
    "noun": "Spengler",
    "article": "der"
  },
  {
    "noun": "Sperber",
    "article": "der"
  },
  {
    "noun": "Sperling",
    "article": "der"
  },
  {
    "noun": "Sperrbalken",
    "article": "der"
  },
  {
    "noun": "Sperrballon",
    "article": "der"
  },
  {
    "noun": "Sperrbaum",
    "article": "der"
  },
  {
    "noun": "Sperrdruck",
    "article": "der"
  },
  {
    "noun": "Sperrhebel",
    "article": "der"
  },
  {
    "noun": "Sperrkreis",
    "article": "der"
  },
  {
    "noun": "Sperrwert",
    "article": "der"
  },
  {
    "noun": "Sperrzoll",
    "article": "der"
  },
  {
    "noun": "Spezialarzt",
    "article": "der"
  },
  {
    "noun": "Spezialfall",
    "article": "der"
  },
  {
    "noun": "Spezialist",
    "article": "der"
  },
  {
    "noun": "Speziallack",
    "article": "der"
  },
  {
    "noun": "Spezialslalom",
    "article": "der"
  },
  {
    "noun": "Spezifikationskauf",
    "article": "der"
  },
  {
    "noun": "Sphalerit",
    "article": "der"
  },
  {
    "noun": "Sphinkter",
    "article": "der"
  },
  {
    "noun": "Sphinx",
    "article": "der"
  },
  {
    "noun": "Spickaal",
    "article": "der"
  },
  {
    "noun": "Spickzettel",
    "article": "der"
  },
  {
    "noun": "Spiegel",
    "article": "der"
  },
  {
    "noun": "Spiegelfechter",
    "article": "der"
  },
  {
    "noun": "Spiegelkarpfen",
    "article": "der"
  },
  {
    "noun": "Spiegelsaal",
    "article": "der"
  },
  {
    "noun": "Spiegelschleifer",
    "article": "der"
  },
  {
    "noun": "Spiegelschrank",
    "article": "der"
  },
  {
    "noun": "Spiegelstrich",
    "article": "der"
  },
  {
    "noun": "Spielanzug",
    "article": "der"
  },
  {
    "noun": "Spielautomat",
    "article": "der"
  },
  {
    "noun": "Spielball",
    "article": "der"
  },
  {
    "noun": "Spielbeginn",
    "article": "der"
  },
  {
    "noun": "Spieleinsatz",
    "article": "der"
  },
  {
    "noun": "Spieler",
    "article": "der"
  },
  {
    "noun": "Spielfilm",
    "article": "der"
  },
  {
    "noun": "Spielgestalter",
    "article": "der"
  },
  {
    "noun": "Spielkamerad",
    "article": "der"
  },
  {
    "noun": "Spielleiter",
    "article": "der"
  },
  {
    "noun": "Spielmacher",
    "article": "der"
  },
  {
    "noun": "Spielmann",
    "article": "der"
  },
  {
    "noun": "Spielmannszug",
    "article": "der"
  },
  {
    "noun": "Spielplan",
    "article": "der"
  },
  {
    "noun": "Spielplatz",
    "article": "der"
  },
  {
    "noun": "Spielraum",
    "article": "der"
  },
  {
    "noun": "Spielsaal",
    "article": "der"
  },
  {
    "noun": "Spielstand",
    "article": "der"
  },
  {
    "noun": "Spielstein",
    "article": "der"
  },
  {
    "noun": "Spieltag",
    "article": "der"
  },
  {
    "noun": "Spieltisch",
    "article": "der"
  },
  {
    "noun": "Spieltrieb",
    "article": "der"
  },
  {
    "noun": "Spielverderber",
    "article": "der"
  },
  {
    "noun": "Spielzug",
    "article": "der"
  },
  {
    "noun": "Spike",
    "article": "der"
  },
  {
    "noun": "Spikereifen",
    "article": "der"
  },
  {
    "noun": "Spikesreifen",
    "article": "der"
  },
  {
    "noun": "Spin",
    "article": "der"
  },
  {
    "noun": "Spinat",
    "article": "der"
  },
  {
    "noun": "Spinell",
    "article": "der"
  },
  {
    "noun": "Spinnaker",
    "article": "der"
  },
  {
    "noun": "Spinner",
    "article": "der"
  },
  {
    "noun": "Spinnfaden",
    "article": "der"
  },
  {
    "noun": "Spinnrocken",
    "article": "der"
  },
  {
    "noun": "Spinnwirtel",
    "article": "der"
  },
  {
    "noun": "Spion",
    "article": "der"
  },
  {
    "noun": "Spionagering",
    "article": "der"
  },
  {
    "noun": "Spiralbohrer",
    "article": "der"
  },
  {
    "noun": "Spiralnebel",
    "article": "der"
  },
  {
    "noun": "Spirant",
    "article": "der"
  },
  {
    "noun": "Spirit",
    "article": "der"
  },
  {
    "noun": "Spiritismus",
    "article": "der"
  },
  {
    "noun": "Spiritist",
    "article": "der"
  },
  {
    "noun": "Spiritualismus",
    "article": "der"
  },
  {
    "noun": "Spiritualist",
    "article": "der"
  },
  {
    "noun": "Spiritus",
    "article": "der"
  },
  {
    "noun": "Spiritusbrenner",
    "article": "der"
  },
  {
    "noun": "Spirituskocher",
    "article": "der"
  },
  {
    "noun": "Spirituslack",
    "article": "der"
  },
  {
    "noun": "Spitz",
    "article": "der"
  },
  {
    "noun": "Spitzahorn",
    "article": "der"
  },
  {
    "noun": "Spitzbart",
    "article": "der"
  },
  {
    "noun": "Spitzbogen",
    "article": "der"
  },
  {
    "noun": "Spitzbohrer",
    "article": "der"
  },
  {
    "noun": "Spitzbube",
    "article": "der"
  },
  {
    "noun": "Spitzel",
    "article": "der"
  },
  {
    "noun": "Spitzenfilm",
    "article": "der"
  },
  {
    "noun": "Spitzenkandidat",
    "article": "der"
  },
  {
    "noun": "Spitzenkoch",
    "article": "der"
  },
  {
    "noun": "Spitzenpolitiker",
    "article": "der"
  },
  {
    "noun": "Spitzenreiter",
    "article": "der"
  },
  {
    "noun": "Spitzenspieler",
    "article": "der"
  },
  {
    "noun": "Spitzensport",
    "article": "der"
  },
  {
    "noun": "Spitzensportler",
    "article": "der"
  },
  {
    "noun": "Spitzentanz",
    "article": "der"
  },
  {
    "noun": "Spitzenverdiener",
    "article": "der"
  },
  {
    "noun": "Spitzenwert",
    "article": "der"
  },
  {
    "noun": "Spitzer",
    "article": "der"
  },
  {
    "noun": "Spitzname",
    "article": "der"
  },
  {
    "noun": "Spitzwegerich",
    "article": "der"
  },
  {
    "noun": "Spleen",
    "article": "der"
  },
  {
    "noun": "Splint",
    "article": "der"
  },
  {
    "noun": "Spliss",
    "article": "der"
  },
  {
    "noun": "Splitt",
    "article": "der"
  },
  {
    "noun": "Splitter",
    "article": "der"
  },
  {
    "noun": "Spoiler",
    "article": "der"
  },
  {
    "noun": "Spondeus",
    "article": "der"
  },
  {
    "noun": "Sponsor",
    "article": "der"
  },
  {
    "noun": "Spor",
    "article": "der"
  },
  {
    "noun": "Sporn",
    "article": "der"
  },
  {
    "noun": "Sporophyt",
    "article": "der"
  },
  {
    "noun": "Sport",
    "article": "der"
  },
  {
    "noun": "Sportanzug",
    "article": "der"
  },
  {
    "noun": "Sportarzt",
    "article": "der"
  },
  {
    "noun": "Sportbericht",
    "article": "der"
  },
  {
    "noun": "Sportclub",
    "article": "der"
  },
  {
    "noun": "Sportdress",
    "article": "der"
  },
  {
    "noun": "Sportflieger",
    "article": "der"
  },
  {
    "noun": "Sportfreund",
    "article": "der"
  },
  {
    "noun": "Sportlehrer",
    "article": "der"
  },
  {
    "noun": "Sportler",
    "article": "der"
  },
  {
    "noun": "Sportmediziner",
    "article": "der"
  },
  {
    "noun": "Sportplatz",
    "article": "der"
  },
  {
    "noun": "Sportreporter",
    "article": "der"
  },
  {
    "noun": "Sportschuh",
    "article": "der"
  },
  {
    "noun": "Sportteil",
    "article": "der"
  },
  {
    "noun": "Sportunterricht",
    "article": "der"
  },
  {
    "noun": "Sportverband",
    "article": "der"
  },
  {
    "noun": "Sportverein",
    "article": "der"
  },
  {
    "noun": "Sportwagen",
    "article": "der"
  },
  {
    "noun": "Spotmarkt",
    "article": "der"
  },
  {
    "noun": "Spott",
    "article": "der"
  },
  {
    "noun": "Spottname",
    "article": "der"
  },
  {
    "noun": "Spottpreis",
    "article": "der"
  },
  {
    "noun": "Spottvogel",
    "article": "der"
  },
  {
    "noun": "Spracherwerb",
    "article": "der"
  },
  {
    "noun": "Sprachfehler",
    "article": "der"
  },
  {
    "noun": "Sprachgebrauch",
    "article": "der"
  },
  {
    "noun": "Sprachinhalt",
    "article": "der"
  },
  {
    "noun": "Sprachkenner",
    "article": "der"
  },
  {
    "noun": "Sprachkurs",
    "article": "der"
  },
  {
    "noun": "Sprachlaut",
    "article": "der"
  },
  {
    "noun": "Sprachlehrer",
    "article": "der"
  },
  {
    "noun": "Sprachraum",
    "article": "der"
  },
  {
    "noun": "Sprachschatz",
    "article": "der"
  },
  {
    "noun": "Sprachschnitzer",
    "article": "der"
  },
  {
    "noun": "Sprachunterricht",
    "article": "der"
  },
  {
    "noun": "Sprachwissenschaftler",
    "article": "der"
  },
  {
    "noun": "Sprechchor",
    "article": "der"
  },
  {
    "noun": "Sprecher",
    "article": "der"
  },
  {
    "noun": "Sprechfunk",
    "article": "der"
  },
  {
    "noun": "Sprechgesang",
    "article": "der"
  },
  {
    "noun": "Sprechtakt",
    "article": "der"
  },
  {
    "noun": "Sprechverkehr",
    "article": "der"
  },
  {
    "noun": "Sprengel",
    "article": "der"
  },
  {
    "noun": "Sprenger",
    "article": "der"
  },
  {
    "noun": "Sprengkopf",
    "article": "der"
  },
  {
    "noun": "Sprengmeister",
    "article": "der"
  },
  {
    "noun": "Sprengsatz",
    "article": "der"
  },
  {
    "noun": "Sprengstoff",
    "article": "der"
  },
  {
    "noun": "Sprengtrichter",
    "article": "der"
  },
  {
    "noun": "Sprengtrupp",
    "article": "der"
  },
  {
    "noun": "Sprenkel",
    "article": "der"
  },
  {
    "noun": "Spriegel",
    "article": "der"
  },
  {
    "noun": "Spring",
    "article": "der"
  },
  {
    "noun": "Springbock",
    "article": "der"
  },
  {
    "noun": "Springbrunnen",
    "article": "der"
  },
  {
    "noun": "Springer",
    "article": "der"
  },
  {
    "noun": "Springerstiefel",
    "article": "der"
  },
  {
    "noun": "Springinsfeld",
    "article": "der"
  },
  {
    "noun": "Springquell",
    "article": "der"
  },
  {
    "noun": "Springreiter",
    "article": "der"
  },
  {
    "noun": "Sprinkler",
    "article": "der"
  },
  {
    "noun": "Sprint",
    "article": "der"
  },
  {
    "noun": "Sprinter",
    "article": "der"
  },
  {
    "noun": "Sprit",
    "article": "der"
  },
  {
    "noun": "Spritzbeton",
    "article": "der"
  },
  {
    "noun": "Spritzbeutel",
    "article": "der"
  },
  {
    "noun": "Spritzer",
    "article": "der"
  },
  {
    "noun": "Spritzguss",
    "article": "der"
  },
  {
    "noun": "Spritzkuchen",
    "article": "der"
  },
  {
    "noun": "Spritzlack",
    "article": "der"
  },
  {
    "noun": "Spritzschutz",
    "article": "der"
  },
  {
    "noun": "Spritzversteller",
    "article": "der"
  },
  {
    "noun": "Spross",
    "article": "der"
  },
  {
    "noun": "Sprossenkohl",
    "article": "der"
  },
  {
    "noun": "Sprosser",
    "article": "der"
  },
  {
    "noun": "Spruch",
    "article": "der"
  },
  {
    "noun": "Sprudel",
    "article": "der"
  },
  {
    "noun": "Sprudelstein",
    "article": "der"
  },
  {
    "noun": "Sprung",
    "article": "der"
  },
  {
    "noun": "Sprunglauf",
    "article": "der"
  },
  {
    "noun": "Sprungturm",
    "article": "der"
  },
  {
    "noun": "Spucknapf",
    "article": "der"
  },
  {
    "noun": "Spuk",
    "article": "der"
  },
  {
    "noun": "Spukgeist",
    "article": "der"
  },
  {
    "noun": "Spuler",
    "article": "der"
  },
  {
    "noun": "Spulwurm",
    "article": "der"
  },
  {
    "noun": "Spund",
    "article": "der"
  },
  {
    "noun": "Spurennachweis",
    "article": "der"
  },
  {
    "noun": "Spurenstoff",
    "article": "der"
  },
  {
    "noun": "Spurkranz",
    "article": "der"
  },
  {
    "noun": "Spurt",
    "article": "der"
  },
  {
    "noun": "Spurwechsel",
    "article": "der"
  },
  {
    "noun": "Sputnik",
    "article": "der"
  },
  {
    "noun": "Staat",
    "article": "der"
  },
  {
    "noun": "Staatenbund",
    "article": "der"
  },
  {
    "noun": "Staatsakt",
    "article": "der"
  },
  {
    "noun": "Staatsanwalt",
    "article": "der"
  },
  {
    "noun": "Staatsapparat",
    "article": "der"
  },
  {
    "noun": "Staatsbankrott",
    "article": "der"
  },
  {
    "noun": "Staatsbesitz",
    "article": "der"
  },
  {
    "noun": "Staatsbesuch",
    "article": "der"
  },
  {
    "noun": "Staatsbetrieb",
    "article": "der"
  },
  {
    "noun": "Staatschef",
    "article": "der"
  },
  {
    "noun": "Staatsdiener",
    "article": "der"
  },
  {
    "noun": "Staatsdienst",
    "article": "der"
  },
  {
    "noun": "Staatsempfang",
    "article": "der"
  },
  {
    "noun": "Staatsetat",
    "article": "der"
  },
  {
    "noun": "Staatsfeiertag",
    "article": "der"
  },
  {
    "noun": "Staatsfeind",
    "article": "der"
  },
  {
    "noun": "Staatsgerichtshof",
    "article": "der"
  },
  {
    "noun": "Staatshandel",
    "article": "der"
  },
  {
    "noun": "Staatshaushalt",
    "article": "der"
  },
  {
    "noun": "Staatshaushaltsplan",
    "article": "der"
  },
  {
    "noun": "Staatskapitalismus",
    "article": "der"
  },
  {
    "noun": "Staatsmann",
    "article": "der"
  },
  {
    "noun": "Staatsminister",
    "article": "der"
  },
  {
    "noun": "Staatsrat",
    "article": "der"
  },
  {
    "noun": "Staatsschatz",
    "article": "der"
  },
  {
    "noun": "Staatsschutz",
    "article": "der"
  },
  {
    "noun": "Staatssicherheitsdienst",
    "article": "der"
  },
  {
    "noun": "Staatsstreich",
    "article": "der"
  },
  {
    "noun": "Staatsvertrag",
    "article": "der"
  },
  {
    "noun": "Staatswald",
    "article": "der"
  },
  {
    "noun": "Staatszuschuss",
    "article": "der"
  },
  {
    "noun": "Stab",
    "article": "der"
  },
  {
    "noun": "Stabhochspringer",
    "article": "der"
  },
  {
    "noun": "Stabhochsprung",
    "article": "der"
  },
  {
    "noun": "Stabilisator",
    "article": "der"
  },
  {
    "noun": "Stabmagnet",
    "article": "der"
  },
  {
    "noun": "Stabreim",
    "article": "der"
  },
  {
    "noun": "Stabsarzt",
    "article": "der"
  },
  {
    "noun": "Stabschef",
    "article": "der"
  },
  {
    "noun": "Stabsfeldwebel",
    "article": "der"
  },
  {
    "noun": "Stabsoffizier",
    "article": "der"
  },
  {
    "noun": "Stabstahl",
    "article": "der"
  },
  {
    "noun": "Stabsunteroffizier",
    "article": "der"
  },
  {
    "noun": "Stabwechsel",
    "article": "der"
  },
  {
    "noun": "Stachel",
    "article": "der"
  },
  {
    "noun": "Stachelbeerstrauch",
    "article": "der"
  },
  {
    "noun": "Stacheldraht",
    "article": "der"
  },
  {
    "noun": "Stachelflosser",
    "article": "der"
  },
  {
    "noun": "Stachelpilz",
    "article": "der"
  },
  {
    "noun": "Stachelrochen",
    "article": "der"
  },
  {
    "noun": "Stadel",
    "article": "der"
  },
  {
    "noun": "Stadionsprecher",
    "article": "der"
  },
  {
    "noun": "Stadtbewohner",
    "article": "der"
  },
  {
    "noun": "Stadtbezirk",
    "article": "der"
  },
  {
    "noun": "Stadtbummel",
    "article": "der"
  },
  {
    "noun": "Stadtgraben",
    "article": "der"
  },
  {
    "noun": "Stadtkern",
    "article": "der"
  },
  {
    "noun": "Stadtklatsch",
    "article": "der"
  },
  {
    "noun": "Stadtkommandant",
    "article": "der"
  },
  {
    "noun": "Stadtpark",
    "article": "der"
  },
  {
    "noun": "Stadtplan",
    "article": "der"
  },
  {
    "noun": "Stadtplaner",
    "article": "der"
  },
  {
    "noun": "Stadtrand",
    "article": "der"
  },
  {
    "noun": "Stadtrat",
    "article": "der"
  },
  {
    "noun": "Stadtschreiber",
    "article": "der"
  },
  {
    "noun": "Stadtstaat",
    "article": "der"
  },
  {
    "noun": "Stadtstreicher",
    "article": "der"
  },
  {
    "noun": "Stadtteil",
    "article": "der"
  },
  {
    "noun": "Stadtvater",
    "article": "der"
  },
  {
    "noun": "Stadtverkehr",
    "article": "der"
  },
  {
    "noun": "Stadtwald",
    "article": "der"
  },
  {
    "noun": "Stafettenlauf",
    "article": "der"
  },
  {
    "noun": "Staffellauf",
    "article": "der"
  },
  {
    "noun": "Staffelpreis",
    "article": "der"
  },
  {
    "noun": "Staffelstab",
    "article": "der"
  },
  {
    "noun": "Stahl",
    "article": "der"
  },
  {
    "noun": "Stahlarbeiter",
    "article": "der"
  },
  {
    "noun": "Stahlbau",
    "article": "der"
  },
  {
    "noun": "Stahlbeton",
    "article": "der"
  },
  {
    "noun": "Stahlbetonbau",
    "article": "der"
  },
  {
    "noun": "Stahlblock",
    "article": "der"
  },
  {
    "noun": "Stahldraht",
    "article": "der"
  },
  {
    "noun": "Stahlguss",
    "article": "der"
  },
  {
    "noun": "Stahlhelm",
    "article": "der"
  },
  {
    "noun": "Stahlstab",
    "article": "der"
  },
  {
    "noun": "Stahlstich",
    "article": "der"
  },
  {
    "noun": "Staketenzaun",
    "article": "der"
  },
  {
    "noun": "Stalagmit",
    "article": "der"
  },
  {
    "noun": "Stalaktit",
    "article": "der"
  },
  {
    "noun": "Stall",
    "article": "der"
  },
  {
    "noun": "Stallbursche",
    "article": "der"
  },
  {
    "noun": "Stalldung",
    "article": "der"
  },
  {
    "noun": "Stallgeruch",
    "article": "der"
  },
  {
    "noun": "Stallhase",
    "article": "der"
  },
  {
    "noun": "Stallknecht",
    "article": "der"
  },
  {
    "noun": "Stallmeister",
    "article": "der"
  },
  {
    "noun": "Stamm",
    "article": "der"
  },
  {
    "noun": "Stammbaum",
    "article": "der"
  },
  {
    "noun": "Stammgast",
    "article": "der"
  },
  {
    "noun": "Stammhalter",
    "article": "der"
  },
  {
    "noun": "Stammkunde",
    "article": "der"
  },
  {
    "noun": "Stammplatz",
    "article": "der"
  },
  {
    "noun": "Stammsitz",
    "article": "der"
  },
  {
    "noun": "Stammspieler",
    "article": "der"
  },
  {
    "noun": "Stammtisch",
    "article": "der"
  },
  {
    "noun": "Stammtischpolitiker",
    "article": "der"
  },
  {
    "noun": "Stammvater",
    "article": "der"
  },
  {
    "noun": "Stammvokal",
    "article": "der"
  },
  {
    "noun": "Stampfasphalt",
    "article": "der"
  },
  {
    "noun": "Stampfbeton",
    "article": "der"
  },
  {
    "noun": "Stampfer",
    "article": "der"
  },
  {
    "noun": "Stand",
    "article": "der"
  },
  {
    "noun": "Standard",
    "article": "der"
  },
  {
    "noun": "Standardbrief",
    "article": "der"
  },
  {
    "noun": "Standardpreis",
    "article": "der"
  },
  {
    "noun": "Standardtyp",
    "article": "der"
  },
  {
    "noun": "Standardwert",
    "article": "der"
  },
  {
    "noun": "Standesunterschied",
    "article": "der"
  },
  {
    "noun": "Standort",
    "article": "der"
  },
  {
    "noun": "Standortfaktor",
    "article": "der"
  },
  {
    "noun": "Standortkatalog",
    "article": "der"
  },
  {
    "noun": "Standortkommandant",
    "article": "der"
  },
  {
    "noun": "Standortnachteil",
    "article": "der"
  },
  {
    "noun": "Standortvorteil",
    "article": "der"
  },
  {
    "noun": "Standortwechsel",
    "article": "der"
  },
  {
    "noun": "Standplatz",
    "article": "der"
  },
  {
    "noun": "Standpunkt",
    "article": "der"
  },
  {
    "noun": "Standquartier",
    "article": "der"
  },
  {
    "noun": "Standvogel",
    "article": "der"
  },
  {
    "noun": "Stangenbohrer",
    "article": "der"
  },
  {
    "noun": "Stangenspargel",
    "article": "der"
  },
  {
    "noun": "Stapel",
    "article": "der"
  },
  {
    "noun": "Stapelbetrieb",
    "article": "der"
  },
  {
    "noun": "Stapellauf",
    "article": "der"
  },
  {
    "noun": "Stapelplatz",
    "article": "der"
  },
  {
    "noun": "Stapes",
    "article": "der"
  },
  {
    "noun": "Staphylokokkus",
    "article": "der"
  },
  {
    "noun": "Stapler",
    "article": "der"
  },
  {
    "noun": "Staplerfahrer",
    "article": "der"
  },
  {
    "noun": "Star",
    "article": "der"
  },
  {
    "noun": "Staranwalt",
    "article": "der"
  },
  {
    "noun": "Stargast",
    "article": "der"
  },
  {
    "noun": "Starkstrom",
    "article": "der"
  },
  {
    "noun": "Starmatz",
    "article": "der"
  },
  {
    "noun": "Starrkopf",
    "article": "der"
  },
  {
    "noun": "Starrkrampf",
    "article": "der"
  },
  {
    "noun": "Starrsinn",
    "article": "der"
  },
  {
    "noun": "Start",
    "article": "der"
  },
  {
    "noun": "Startblock",
    "article": "der"
  },
  {
    "noun": "Starter",
    "article": "der"
  },
  {
    "noun": "Startplatz",
    "article": "der"
  },
  {
    "noun": "Startschuss",
    "article": "der"
  },
  {
    "noun": "Startsprung",
    "article": "der"
  },
  {
    "noun": "Statiker",
    "article": "der"
  },
  {
    "noun": "Stationsarzt",
    "article": "der"
  },
  {
    "noun": "Stationsdienst",
    "article": "der"
  },
  {
    "noun": "Statist",
    "article": "der"
  },
  {
    "noun": "Statistiker",
    "article": "der"
  },
  {
    "noun": "Stator",
    "article": "der"
  },
  {
    "noun": "Statthalter",
    "article": "der"
  },
  {
    "noun": "Status",
    "article": "der"
  },
  {
    "noun": "Stau",
    "article": "der"
  },
  {
    "noun": "Staub",
    "article": "der"
  },
  {
    "noun": "Staubbach",
    "article": "der"
  },
  {
    "noun": "Staubbesen",
    "article": "der"
  },
  {
    "noun": "Staubfaden",
    "article": "der"
  },
  {
    "noun": "Staubkamm",
    "article": "der"
  },
  {
    "noun": "Staubsand",
    "article": "der"
  },
  {
    "noun": "Staubsauger",
    "article": "der"
  },
  {
    "noun": "Staubsturm",
    "article": "der"
  },
  {
    "noun": "Staubwedel",
    "article": "der"
  },
  {
    "noun": "Staubzucker",
    "article": "der"
  },
  {
    "noun": "Staudamm",
    "article": "der"
  },
  {
    "noun": "Staudensellerie",
    "article": "der"
  },
  {
    "noun": "Stauer",
    "article": "der"
  },
  {
    "noun": "Staupunkt",
    "article": "der"
  },
  {
    "noun": "Stauraum",
    "article": "der"
  },
  {
    "noun": "Stausee",
    "article": "der"
  },
  {
    "noun": "Steamer",
    "article": "der"
  },
  {
    "noun": "Steatit",
    "article": "der"
  },
  {
    "noun": "Stechbeitel",
    "article": "der"
  },
  {
    "noun": "Stecher",
    "article": "der"
  },
  {
    "noun": "Stechheber",
    "article": "der"
  },
  {
    "noun": "Stechzirkel",
    "article": "der"
  },
  {
    "noun": "Steckbrief",
    "article": "der"
  },
  {
    "noun": "Stecker",
    "article": "der"
  },
  {
    "noun": "Steckling",
    "article": "der"
  },
  {
    "noun": "Stecknadelkopf",
    "article": "der"
  },
  {
    "noun": "Steckplatz",
    "article": "der"
  },
  {
    "noun": "Steckverbinder",
    "article": "der"
  },
  {
    "noun": "Steg",
    "article": "der"
  },
  {
    "noun": "Stegosaurus",
    "article": "der"
  },
  {
    "noun": "Stehkragen",
    "article": "der"
  },
  {
    "noun": "Stehler",
    "article": "der"
  },
  {
    "noun": "Stehplatz",
    "article": "der"
  },
  {
    "noun": "Steiger",
    "article": "der"
  },
  {
    "noun": "Steigerer",
    "article": "der"
  },
  {
    "noun": "Steigungswinkel",
    "article": "der"
  },
  {
    "noun": "Steilflug",
    "article": "der"
  },
  {
    "noun": "Steilhang",
    "article": "der"
  },
  {
    "noun": "Stein",
    "article": "der"
  },
  {
    "noun": "Steinadler",
    "article": "der"
  },
  {
    "noun": "Steinbau",
    "article": "der"
  },
  {
    "noun": "Steinbaukasten",
    "article": "der"
  },
  {
    "noun": "Steinblock",
    "article": "der"
  },
  {
    "noun": "Steinbock",
    "article": "der"
  },
  {
    "noun": "Steinboden",
    "article": "der"
  },
  {
    "noun": "Steinbrech",
    "article": "der"
  },
  {
    "noun": "Steinbrecher",
    "article": "der"
  },
  {
    "noun": "Steinbruch",
    "article": "der"
  },
  {
    "noun": "Steinbutt",
    "article": "der"
  },
  {
    "noun": "Steindamm",
    "article": "der"
  },
  {
    "noun": "Steindruck",
    "article": "der"
  },
  {
    "noun": "Steingarten",
    "article": "der"
  },
  {
    "noun": "Steinhaufen",
    "article": "der"
  },
  {
    "noun": "Steinkauz",
    "article": "der"
  },
  {
    "noun": "Steinkohlenbergbau",
    "article": "der"
  },
  {
    "noun": "Steinmarder",
    "article": "der"
  },
  {
    "noun": "Steinmetz",
    "article": "der"
  },
  {
    "noun": "Steinpilz",
    "article": "der"
  },
  {
    "noun": "Steinsarg",
    "article": "der"
  },
  {
    "noun": "Steinschlag",
    "article": "der"
  },
  {
    "noun": "Steinschneider",
    "article": "der"
  },
  {
    "noun": "Steinsetzer",
    "article": "der"
  },
  {
    "noun": "Steintopf",
    "article": "der"
  },
  {
    "noun": "Steinzeitmensch",
    "article": "der"
  },
  {
    "noun": "Steirer",
    "article": "der"
  },
  {
    "noun": "Stek",
    "article": "der"
  },
  {
    "noun": "Stellenabbau",
    "article": "der"
  },
  {
    "noun": "Stelleninhaber",
    "article": "der"
  },
  {
    "noun": "Stellenmarkt",
    "article": "der"
  },
  {
    "noun": "Stellenplan",
    "article": "der"
  },
  {
    "noun": "Stellenwechsel",
    "article": "der"
  },
  {
    "noun": "Stellenwert",
    "article": "der"
  },
  {
    "noun": "Steller",
    "article": "der"
  },
  {
    "noun": "Stellhebel",
    "article": "der"
  },
  {
    "noun": "Stellknopf",
    "article": "der"
  },
  {
    "noun": "Stellmacher",
    "article": "der"
  },
  {
    "noun": "Stellmotor",
    "article": "der"
  },
  {
    "noun": "Stellplatz",
    "article": "der"
  },
  {
    "noun": "Stellungskrieg",
    "article": "der"
  },
  {
    "noun": "Stellungswechsel",
    "article": "der"
  },
  {
    "noun": "Stellvertreter",
    "article": "der"
  },
  {
    "noun": "Stelzvogel",
    "article": "der"
  },
  {
    "noun": "Stempel",
    "article": "der"
  },
  {
    "noun": "Stempelschneider",
    "article": "der"
  },
  {
    "noun": "Stengel",
    "article": "der"
  },
  {
    "noun": "Stenotypist",
    "article": "der"
  },
  {
    "noun": "Step",
    "article": "der"
  },
  {
    "noun": "Steppenadler",
    "article": "der"
  },
  {
    "noun": "Steppenfuchs",
    "article": "der"
  },
  {
    "noun": "Stepptanz",
    "article": "der"
  },
  {
    "noun": "Sterbefall",
    "article": "der"
  },
  {
    "noun": "Sterbetag",
    "article": "der"
  },
  {
    "noun": "Stereoempfang",
    "article": "der"
  },
  {
    "noun": "Stereofilm",
    "article": "der"
  },
  {
    "noun": "Stereokomparator",
    "article": "der"
  },
  {
    "noun": "Stereoplattenspieler",
    "article": "der"
  },
  {
    "noun": "Sterilisator",
    "article": "der"
  },
  {
    "noun": "Stern",
    "article": "der"
  },
  {
    "noun": "Sterndeuter",
    "article": "der"
  },
  {
    "noun": "Sternenhimmel",
    "article": "der"
  },
  {
    "noun": "Sternforscher",
    "article": "der"
  },
  {
    "noun": "Sterngucker",
    "article": "der"
  },
  {
    "noun": "Sternhaufen",
    "article": "der"
  },
  {
    "noun": "Sternhimmel",
    "article": "der"
  },
  {
    "noun": "Sternmotor",
    "article": "der"
  },
  {
    "noun": "Sternort",
    "article": "der"
  },
  {
    "noun": "Sterz",
    "article": "der"
  },
  {
    "noun": "Steuerabzug",
    "article": "der"
  },
  {
    "noun": "Steuerausgleich",
    "article": "der"
  },
  {
    "noun": "Steuerausschuss",
    "article": "der"
  },
  {
    "noun": "Steuerbefehl",
    "article": "der"
  },
  {
    "noun": "Steuerberater",
    "article": "der"
  },
  {
    "noun": "Steuerbescheid",
    "article": "der"
  },
  {
    "noun": "Steuerbetrag",
    "article": "der"
  },
  {
    "noun": "Steuerbetrug",
    "article": "der"
  },
  {
    "noun": "Steuereinnehmer",
    "article": "der"
  },
  {
    "noun": "Steuererlass",
    "article": "der"
  },
  {
    "noun": "Steuerfachgehilfe",
    "article": "der"
  },
  {
    "noun": "Steuerfahnder",
    "article": "der"
  },
  {
    "noun": "Steuerfreibetrag",
    "article": "der"
  },
  {
    "noun": "Steuerhebel",
    "article": "der"
  },
  {
    "noun": "Steuermann",
    "article": "der"
  },
  {
    "noun": "Steuernachlass",
    "article": "der"
  },
  {
    "noun": "Steuersatz",
    "article": "der"
  },
  {
    "noun": "Steuerschalter",
    "article": "der"
  },
  {
    "noun": "Steuerschuldner",
    "article": "der"
  },
  {
    "noun": "Steuertarif",
    "article": "der"
  },
  {
    "noun": "Steuerungsmechanismus",
    "article": "der"
  },
  {
    "noun": "Steuervorteil",
    "article": "der"
  },
  {
    "noun": "Steuerzahler",
    "article": "der"
  },
  {
    "noun": "Steuerzettel",
    "article": "der"
  },
  {
    "noun": "Steuerzuschlag",
    "article": "der"
  },
  {
    "noun": "Steward",
    "article": "der"
  },
  {
    "noun": "Stich",
    "article": "der"
  },
  {
    "noun": "Stichel",
    "article": "der"
  },
  {
    "noun": "Stichler",
    "article": "der"
  },
  {
    "noun": "Stichling",
    "article": "der"
  },
  {
    "noun": "Stichpunkt",
    "article": "der"
  },
  {
    "noun": "Stichtag",
    "article": "der"
  },
  {
    "noun": "Stickhusten",
    "article": "der"
  },
  {
    "noun": "Stickstoff",
    "article": "der"
  },
  {
    "noun": "Stiefbruder",
    "article": "der"
  },
  {
    "noun": "Stiefel",
    "article": "der"
  },
  {
    "noun": "Stiefelknecht",
    "article": "der"
  },
  {
    "noun": "Stiefsohn",
    "article": "der"
  },
  {
    "noun": "Stiefvater",
    "article": "der"
  },
  {
    "noun": "Stiegenabsatz",
    "article": "der"
  },
  {
    "noun": "Stieglitz",
    "article": "der"
  },
  {
    "noun": "Stiel",
    "article": "der"
  },
  {
    "noun": "Stier",
    "article": "der"
  },
  {
    "noun": "Stierkampf",
    "article": "der"
  },
  {
    "noun": "Stiernacken",
    "article": "der"
  },
  {
    "noun": "Stifter",
    "article": "der"
  },
  {
    "noun": "Stil",
    "article": "der"
  },
  {
    "noun": "Stilist",
    "article": "der"
  },
  {
    "noun": "Stillstand",
    "article": "der"
  },
  {
    "noun": "Stilus",
    "article": "der"
  },
  {
    "noun": "Stimmbruch",
    "article": "der"
  },
  {
    "noun": "Stimmenanteil",
    "article": "der"
  },
  {
    "noun": "Stimmenkauf",
    "article": "der"
  },
  {
    "noun": "Stimmenverlust",
    "article": "der"
  },
  {
    "noun": "Stimmer",
    "article": "der"
  },
  {
    "noun": "Stimmungsumschwung",
    "article": "der"
  },
  {
    "noun": "Stimmwechsel",
    "article": "der"
  },
  {
    "noun": "Stimmzettel",
    "article": "der"
  },
  {
    "noun": "Stimulator",
    "article": "der"
  },
  {
    "noun": "Stimulus",
    "article": "der"
  },
  {
    "noun": "Stinker",
    "article": "der"
  },
  {
    "noun": "Stint",
    "article": "der"
  },
  {
    "noun": "Stipendiat",
    "article": "der"
  },
  {
    "noun": "Stipp",
    "article": "der"
  },
  {
    "noun": "Stippbesuch",
    "article": "der"
  },
  {
    "noun": "Stirnreif",
    "article": "der"
  },
  {
    "noun": "Stock",
    "article": "der"
  },
  {
    "noun": "Stockfisch",
    "article": "der"
  },
  {
    "noun": "Stockhieb",
    "article": "der"
  },
  {
    "noun": "Stockpunkt",
    "article": "der"
  },
  {
    "noun": "Stockschirm",
    "article": "der"
  },
  {
    "noun": "Stockzahn",
    "article": "der"
  },
  {
    "noun": "Stoff",
    "article": "der"
  },
  {
    "noun": "Stoffaustausch",
    "article": "der"
  },
  {
    "noun": "Stoffel",
    "article": "der"
  },
  {
    "noun": "Stoffrest",
    "article": "der"
  },
  {
    "noun": "Stofftransport",
    "article": "der"
  },
  {
    "noun": "Stoffwechsel",
    "article": "der"
  },
  {
    "noun": "Stoizismus",
    "article": "der"
  },
  {
    "noun": "Stollen",
    "article": "der"
  },
  {
    "noun": "Stollengang",
    "article": "der"
  },
  {
    "noun": "Stolperstein",
    "article": "der"
  },
  {
    "noun": "Stolz",
    "article": "der"
  },
  {
    "noun": "Stop",
    "article": "der"
  },
  {
    "noun": "Stopfer",
    "article": "der"
  },
  {
    "noun": "Stoppelbart",
    "article": "der"
  },
  {
    "noun": "Stopper",
    "article": "der"
  },
  {
    "noun": "Storch",
    "article": "der"
  },
  {
    "noun": "Storchschnabel",
    "article": "der"
  },
  {
    "noun": "Store",
    "article": "der"
  },
  {
    "noun": "Stotterer",
    "article": "der"
  },
  {
    "noun": "Strabismus",
    "article": "der"
  },
  {
    "noun": "Straddlesprung",
    "article": "der"
  },
  {
    "noun": "Strafantrag",
    "article": "der"
  },
  {
    "noun": "Strafaufschub",
    "article": "der"
  },
  {
    "noun": "Straferlass",
    "article": "der"
  },
  {
    "noun": "Straffall",
    "article": "der"
  },
  {
    "noun": "Strafprozess",
    "article": "der"
  },
  {
    "noun": "Straftatbestand",
    "article": "der"
  },
  {
    "noun": "Strafverteidiger",
    "article": "der"
  },
  {
    "noun": "Strafvollzug",
    "article": "der"
  },
  {
    "noun": "Strafzettel",
    "article": "der"
  },
  {
    "noun": "Strahl",
    "article": "der"
  },
  {
    "noun": "Strahlantrieb",
    "article": "der"
  },
  {
    "noun": "Strahlengang",
    "article": "der"
  },
  {
    "noun": "Strahlenschaden",
    "article": "der"
  },
  {
    "noun": "Strahlenschutz",
    "article": "der"
  },
  {
    "noun": "Strahler",
    "article": "der"
  },
  {
    "noun": "Strahlstrom",
    "article": "der"
  },
  {
    "noun": "Strahlungsfluss",
    "article": "der"
  },
  {
    "noun": "Strand",
    "article": "der"
  },
  {
    "noun": "Strang",
    "article": "der"
  },
  {
    "noun": "Strangguss",
    "article": "der"
  },
  {
    "noun": "Strass",
    "article": "der"
  },
  {
    "noun": "Stratege",
    "article": "der"
  },
  {
    "noun": "Strauch",
    "article": "der"
  },
  {
    "noun": "Strauchdieb",
    "article": "der"
  },
  {
    "noun": "Streamer",
    "article": "der"
  },
  {
    "noun": "Strebebalken",
    "article": "der"
  },
  {
    "noun": "Strebebogen",
    "article": "der"
  },
  {
    "noun": "Strebepfeiler",
    "article": "der"
  },
  {
    "noun": "Streber",
    "article": "der"
  },
  {
    "noun": "Streckenabschnitt",
    "article": "der"
  },
  {
    "noun": "Streckenflug",
    "article": "der"
  },
  {
    "noun": "Strecker",
    "article": "der"
  },
  {
    "noun": "Streich",
    "article": "der"
  },
  {
    "noun": "Streicher",
    "article": "der"
  },
  {
    "noun": "Streifendienst",
    "article": "der"
  },
  {
    "noun": "Streifenwagen",
    "article": "der"
  },
  {
    "noun": "Streifzug",
    "article": "der"
  },
  {
    "noun": "Streik",
    "article": "der"
  },
  {
    "noun": "Streikaufruf",
    "article": "der"
  },
  {
    "noun": "Streikbrecher",
    "article": "der"
  },
  {
    "noun": "Streikposten",
    "article": "der"
  },
  {
    "noun": "Streit",
    "article": "der"
  },
  {
    "noun": "Streiter",
    "article": "der"
  },
  {
    "noun": "Streitfall",
    "article": "der"
  },
  {
    "noun": "Streitgegenstand",
    "article": "der"
  },
  {
    "noun": "Streithahn",
    "article": "der"
  },
  {
    "noun": "Streithammel",
    "article": "der"
  },
  {
    "noun": "Streitkolben",
    "article": "der"
  },
  {
    "noun": "Streitpunkt",
    "article": "der"
  },
  {
    "noun": "Streitwert",
    "article": "der"
  },
  {
    "noun": "Stress",
    "article": "der"
  },
  {
    "noun": "Streuer",
    "article": "der"
  },
  {
    "noun": "Streuner",
    "article": "der"
  },
  {
    "noun": "Streusand",
    "article": "der"
  },
  {
    "noun": "Streuungskoeffizient",
    "article": "der"
  },
  {
    "noun": "Streuzucker",
    "article": "der"
  },
  {
    "noun": "Strich",
    "article": "der"
  },
  {
    "noun": "Strichcode",
    "article": "der"
  },
  {
    "noun": "Strichjunge",
    "article": "der"
  },
  {
    "noun": "Strichpunkt",
    "article": "der"
  },
  {
    "noun": "Strick",
    "article": "der"
  },
  {
    "noun": "Stricker",
    "article": "der"
  },
  {
    "noun": "Striegel",
    "article": "der"
  },
  {
    "noun": "Striemen",
    "article": "der"
  },
  {
    "noun": "String",
    "article": "der"
  },
  {
    "noun": "Strip",
    "article": "der"
  },
  {
    "noun": "Strohhalm",
    "article": "der"
  },
  {
    "noun": "Strohhut",
    "article": "der"
  },
  {
    "noun": "Strohkopf",
    "article": "der"
  },
  {
    "noun": "Strohmann",
    "article": "der"
  },
  {
    "noun": "Strohsack",
    "article": "der"
  },
  {
    "noun": "Strohschober",
    "article": "der"
  },
  {
    "noun": "Strohschuh",
    "article": "der"
  },
  {
    "noun": "Strohwisch",
    "article": "der"
  },
  {
    "noun": "Strohwitwer",
    "article": "der"
  },
  {
    "noun": "Strolch",
    "article": "der"
  },
  {
    "noun": "Strom",
    "article": "der"
  },
  {
    "noun": "Stromabnehmer",
    "article": "der"
  },
  {
    "noun": "Stromanschluss",
    "article": "der"
  },
  {
    "noun": "Stromausfall",
    "article": "der"
  },
  {
    "noun": "Strombedarf",
    "article": "der"
  },
  {
    "noun": "Stromer",
    "article": "der"
  },
  {
    "noun": "Stromerzeuger",
    "article": "der"
  },
  {
    "noun": "Stromkreis",
    "article": "der"
  },
  {
    "noun": "Strommesser",
    "article": "der"
  },
  {
    "noun": "Strompreis",
    "article": "der"
  },
  {
    "noun": "Stromregler",
    "article": "der"
  },
  {
    "noun": "Stromrichter",
    "article": "der"
  },
  {
    "noun": "Stromschlag",
    "article": "der"
  },
  {
    "noun": "Stromtarif",
    "article": "der"
  },
  {
    "noun": "Stromunterbrecher",
    "article": "der"
  },
  {
    "noun": "Stromverbrauch",
    "article": "der"
  },
  {
    "noun": "Stromwandler",
    "article": "der"
  },
  {
    "noun": "Stromwender",
    "article": "der"
  },
  {
    "noun": "Stropp",
    "article": "der"
  },
  {
    "noun": "Strubbelkopf",
    "article": "der"
  },
  {
    "noun": "Strudel",
    "article": "der"
  },
  {
    "noun": "Strudelwurm",
    "article": "der"
  },
  {
    "noun": "Strukturalismus",
    "article": "der"
  },
  {
    "noun": "Strukturplan",
    "article": "der"
  },
  {
    "noun": "Strukturwandel",
    "article": "der"
  },
  {
    "noun": "Strumpf",
    "article": "der"
  },
  {
    "noun": "Strumpfhalter",
    "article": "der"
  },
  {
    "noun": "Strunk",
    "article": "der"
  },
  {
    "noun": "Stubben",
    "article": "der"
  },
  {
    "noun": "Stubenarrest",
    "article": "der"
  },
  {
    "noun": "Stubenhocker",
    "article": "der"
  },
  {
    "noun": "Stubenkamerad",
    "article": "der"
  },
  {
    "noun": "Stuckateur",
    "article": "der"
  },
  {
    "noun": "Stuckgips",
    "article": "der"
  },
  {
    "noun": "Student",
    "article": "der"
  },
  {
    "noun": "Studentenaustausch",
    "article": "der"
  },
  {
    "noun": "Studentenausweis",
    "article": "der"
  },
  {
    "noun": "Studienaufenthalt",
    "article": "der"
  },
  {
    "noun": "Studienberater",
    "article": "der"
  },
  {
    "noun": "Studiendirektor",
    "article": "der"
  },
  {
    "noun": "Studiengang",
    "article": "der"
  },
  {
    "noun": "Studienplan",
    "article": "der"
  },
  {
    "noun": "Studienplatz",
    "article": "der"
  },
  {
    "noun": "Studienrat",
    "article": "der"
  },
  {
    "noun": "Stuhl",
    "article": "der"
  },
  {
    "noun": "Stuhlgang",
    "article": "der"
  },
  {
    "noun": "Stulpenstiefel",
    "article": "der"
  },
  {
    "noun": "Stummel",
    "article": "der"
  },
  {
    "noun": "Stummelschwanz",
    "article": "der"
  },
  {
    "noun": "Stummfilm",
    "article": "der"
  },
  {
    "noun": "Stumpf",
    "article": "der"
  },
  {
    "noun": "Stumpfsinn",
    "article": "der"
  },
  {
    "noun": "Stundenlohn",
    "article": "der"
  },
  {
    "noun": "Stundenplan",
    "article": "der"
  },
  {
    "noun": "Stundenzeiger",
    "article": "der"
  },
  {
    "noun": "Stunk",
    "article": "der"
  },
  {
    "noun": "Stupor",
    "article": "der"
  },
  {
    "noun": "Sturm",
    "article": "der"
  },
  {
    "noun": "Sturmbock",
    "article": "der"
  },
  {
    "noun": "Sturmriemen",
    "article": "der"
  },
  {
    "noun": "Sturmwind",
    "article": "der"
  },
  {
    "noun": "Sturz",
    "article": "der"
  },
  {
    "noun": "Sturzacker",
    "article": "der"
  },
  {
    "noun": "Sturzbach",
    "article": "der"
  },
  {
    "noun": "Sturzflug",
    "article": "der"
  },
  {
    "noun": "Sturzhelm",
    "article": "der"
  },
  {
    "noun": "Stuss",
    "article": "der"
  },
  {
    "noun": "Stutzbart",
    "article": "der"
  },
  {
    "noun": "Stutzer",
    "article": "der"
  },
  {
    "noun": "Subjektivismus",
    "article": "der"
  },
  {
    "noun": "Subjektsatz",
    "article": "der"
  },
  {
    "noun": "Subkontinent",
    "article": "der"
  },
  {
    "noun": "Subskriptionspreis",
    "article": "der"
  },
  {
    "noun": "Substanzverlust",
    "article": "der"
  },
  {
    "noun": "Substanzwert",
    "article": "der"
  },
  {
    "noun": "Subtrahend",
    "article": "der"
  },
  {
    "noun": "Subunternehmer",
    "article": "der"
  },
  {
    "noun": "Sucher",
    "article": "der"
  },
  {
    "noun": "Suchlauf",
    "article": "der"
  },
  {
    "noun": "Suchscheinwerfer",
    "article": "der"
  },
  {
    "noun": "Sud",
    "article": "der"
  },
  {
    "noun": "Sudan",
    "article": "der"
  },
  {
    "noun": "Sudaner",
    "article": "der"
  },
  {
    "noun": "Sudanese",
    "article": "der"
  },
  {
    "noun": "Sudler",
    "article": "der"
  },
  {
    "noun": "Suezkanal",
    "article": "der"
  },
  {
    "noun": "Suff",
    "article": "der"
  },
  {
    "noun": "Sultan",
    "article": "der"
  },
  {
    "noun": "Sumach",
    "article": "der"
  },
  {
    "noun": "Summand",
    "article": "der"
  },
  {
    "noun": "Summenausdruck",
    "article": "der"
  },
  {
    "noun": "Summer",
    "article": "der"
  },
  {
    "noun": "Summton",
    "article": "der"
  },
  {
    "noun": "Sumpf",
    "article": "der"
  },
  {
    "noun": "Sumpfboden",
    "article": "der"
  },
  {
    "noun": "Sums",
    "article": "der"
  },
  {
    "noun": "Sund",
    "article": "der"
  },
  {
    "noun": "Sunnit",
    "article": "der"
  },
  {
    "noun": "Superlativ",
    "article": "der"
  },
  {
    "noun": "Supermarkt",
    "article": "der"
  },
  {
    "noun": "Supervisor",
    "article": "der"
  },
  {
    "noun": "Suppenteller",
    "article": "der"
  },
  {
    "noun": "Suppentopf",
    "article": "der"
  },
  {
    "noun": "Supplementband",
    "article": "der"
  },
  {
    "noun": "Supplementwinkel",
    "article": "der"
  },
  {
    "noun": "Support",
    "article": "der"
  },
  {
    "noun": "Supraleiter",
    "article": "der"
  },
  {
    "noun": "Suprastrom",
    "article": "der"
  },
  {
    "noun": "Surfer",
    "article": "der"
  },
  {
    "noun": "Surrealismus",
    "article": "der"
  },
  {
    "noun": "Swing",
    "article": "der"
  },
  {
    "noun": "Syllogismus",
    "article": "der"
  },
  {
    "noun": "Symbolismus",
    "article": "der"
  },
  {
    "noun": "Sympathisant",
    "article": "der"
  },
  {
    "noun": "Synchronismus",
    "article": "der"
  },
  {
    "noun": "Synchronmotor",
    "article": "der"
  },
  {
    "noun": "Synchronsprecher",
    "article": "der"
  },
  {
    "noun": "Syndikalist",
    "article": "der"
  },
  {
    "noun": "Syndikus",
    "article": "der"
  },
  {
    "noun": "Synergieeffekt",
    "article": "der"
  },
  {
    "noun": "Synoptiker",
    "article": "der"
  },
  {
    "noun": "Synthesizer",
    "article": "der"
  },
  {
    "noun": "Syphilitiker",
    "article": "der"
  },
  {
    "noun": "Syrer",
    "article": "der"
  },
  {
    "noun": "Syrier",
    "article": "der"
  },
  {
    "noun": "Systemabsturz",
    "article": "der"
  },
  {
    "noun": "Systemanalytiker",
    "article": "der"
  },
  {
    "noun": "Systemzusammenbruch",
    "article": "der"
  },
  {
    "noun": "Systemzwang",
    "article": "der"
  },
  {
    "noun": "Szenejargon",
    "article": "der"
  },
  {
    "noun": "Szenenwechsel",
    "article": "der"
  },
  {
    "noun": "Quereinsteigerin",
    "article": "die"
  },
  {
    "noun": "Querfrage",
    "article": "die"
  },
  {
    "noun": "Querlage",
    "article": "die"
  },
  {
    "noun": "Querlatte",
    "article": "die"
  },
  {
    "noun": "Querlinie",
    "article": "die"
  },
  {
    "noun": "Querneigung",
    "article": "die"
  },
  {
    "noun": "Querpfeife",
    "article": "die"
  },
  {
    "noun": "Querrichtung",
    "article": "die"
  },
  {
    "noun": "Quersumme",
    "article": "die"
  },
  {
    "noun": "Quertreiberin",
    "article": "die"
  },
  {
    "noun": "Querverbindung",
    "article": "die"
  },
  {
    "noun": "Quese",
    "article": "die"
  },
  {
    "noun": "Quetsche",
    "article": "die"
  },
  {
    "noun": "Quetschfalte",
    "article": "die"
  },
  {
    "noun": "Quetschgrenze",
    "article": "die"
  },
  {
    "noun": "Quetschkommode",
    "article": "die"
  },
  {
    "noun": "Quetschung",
    "article": "die"
  },
  {
    "noun": "Quiche",
    "article": "die"
  },
  {
    "noun": "Quinte",
    "article": "die"
  },
  {
    "noun": "Quintessenz",
    "article": "die"
  },
  {
    "noun": "Quintilliarde",
    "article": "die"
  },
  {
    "noun": "Quintillion",
    "article": "die"
  },
  {
    "noun": "Quitte",
    "article": "die"
  },
  {
    "noun": "Quittierung",
    "article": "die"
  },
  {
    "noun": "Quittung",
    "article": "die"
  },
  {
    "noun": "Quizfrage",
    "article": "die"
  },
  {
    "noun": "Quizsendung",
    "article": "die"
  },
  {
    "noun": "Quotation",
    "article": "die"
  },
  {
    "noun": "Quote",
    "article": "die"
  },
  {
    "noun": "Quotenfrau",
    "article": "die"
  },
  {
    "noun": "Quotenregelung",
    "article": "die"
  },
  {
    "noun": "Quotierung",
    "article": "die"
  },
  {
    "noun": "Jammergestalt",
    "article": "die"
  },
  {
    "noun": "Jammermiene",
    "article": "die"
  },
  {
    "noun": "Japanerin",
    "article": "die"
  },
  {
    "noun": "Japanseide",
    "article": "die"
  },
  {
    "noun": "Jastimme",
    "article": "die"
  },
  {
    "noun": "Jauche",
    "article": "die"
  },
  {
    "noun": "Jauchegrube",
    "article": "die"
  },
  {
    "noun": "Jauchengrube",
    "article": "die"
  },
  {
    "noun": "Jause",
    "article": "die"
  },
  {
    "noun": "Javanerin",
    "article": "die"
  },
  {
    "noun": "Jazzband",
    "article": "die"
  },
  {
    "noun": "Jazzkapelle",
    "article": "die"
  },
  {
    "noun": "Jeans",
    "article": "die"
  },
  {
    "noun": "Jeanshose",
    "article": "die"
  },
  {
    "noun": "Jeansjacke",
    "article": "die"
  },
  {
    "noun": "Jemenitin",
    "article": "die"
  },
  {
    "noun": "Jenseitigkeit",
    "article": "die"
  },
  {
    "noun": "Jetztzeit",
    "article": "die"
  },
  {
    "noun": "Jiddistik",
    "article": "die"
  },
  {
    "noun": "Jodhpurhose",
    "article": "die"
  },
  {
    "noun": "Jodlerin",
    "article": "die"
  },
  {
    "noun": "Jodtinktur",
    "article": "die"
  },
  {
    "noun": "Jodvergiftung",
    "article": "die"
  },
  {
    "noun": "Jodzahl",
    "article": "die"
  },
  {
    "noun": "Joggerin",
    "article": "die"
  },
  {
    "noun": "Johannisbeere",
    "article": "die"
  },
  {
    "noun": "Jolle",
    "article": "die"
  },
  {
    "noun": "Jordanierin",
    "article": "die"
  },
  {
    "noun": "Journalistik",
    "article": "die"
  },
  {
    "noun": "Journalistin",
    "article": "die"
  },
  {
    "noun": "Jubilarin",
    "article": "die"
  },
  {
    "noun": "Judikative",
    "article": "die"
  },
  {
    "noun": "Jugend",
    "article": "die"
  },
  {
    "noun": "Jugendarbeit",
    "article": "die"
  },
  {
    "noun": "Jugendarbeitslosigkeit",
    "article": "die"
  },
  {
    "noun": "Jugendbewegung",
    "article": "die"
  },
  {
    "noun": "Jugendbrigade",
    "article": "die"
  },
  {
    "noun": "Jugenderinnerung",
    "article": "die"
  },
  {
    "noun": "Jugendgruppe",
    "article": "die"
  },
  {
    "noun": "Jugendherberge",
    "article": "die"
  },
  {
    "noun": "Jugendhilfe",
    "article": "die"
  },
  {
    "noun": "Jugendlichkeit",
    "article": "die"
  },
  {
    "noun": "Jugendliebe",
    "article": "die"
  },
  {
    "noun": "Jugendorganisation",
    "article": "die"
  },
  {
    "noun": "Jugendpflege",
    "article": "die"
  },
  {
    "noun": "Jugendpsychologie",
    "article": "die"
  },
  {
    "noun": "Jugendrichterin",
    "article": "die"
  },
  {
    "noun": "Jugendsendung",
    "article": "die"
  },
  {
    "noun": "Jugendsprache",
    "article": "die"
  },
  {
    "noun": "Jugendstrafanstalt",
    "article": "die"
  },
  {
    "noun": "Jugendweihe",
    "article": "die"
  },
  {
    "noun": "Jugendzeit",
    "article": "die"
  },
  {
    "noun": "Jugendzeitschrift",
    "article": "die"
  },
  {
    "noun": "Jugoslawin",
    "article": "die"
  },
  {
    "noun": "Jukebox",
    "article": "die"
  },
  {
    "noun": "Jungenschule",
    "article": "die"
  },
  {
    "noun": "Jungfer",
    "article": "die"
  },
  {
    "noun": "Jungfernfahrt",
    "article": "die"
  },
  {
    "noun": "Jungfernrede",
    "article": "die"
  },
  {
    "noun": "Jungfernreise",
    "article": "die"
  },
  {
    "noun": "Jungfernschaft",
    "article": "die"
  },
  {
    "noun": "Jungfrau",
    "article": "die"
  },
  {
    "noun": "Junggesellenbude",
    "article": "die"
  },
  {
    "noun": "Junggesellenwohnung",
    "article": "die"
  },
  {
    "noun": "Junggesellenzeit",
    "article": "die"
  },
  {
    "noun": "Junggesellin",
    "article": "die"
  },
  {
    "noun": "Jungmannschaft",
    "article": "die"
  },
  {
    "noun": "Jungschar",
    "article": "die"
  },
  {
    "noun": "Jungsozialistin",
    "article": "die"
  },
  {
    "noun": "Jungsteinzeit",
    "article": "die"
  },
  {
    "noun": "Juniorin",
    "article": "die"
  },
  {
    "noun": "Juniorpartnerin",
    "article": "die"
  },
  {
    "noun": "Junta",
    "article": "die"
  },
  {
    "noun": "Jupiterlampe",
    "article": "die"
  },
  {
    "noun": "Jurastudentin",
    "article": "die"
  },
  {
    "noun": "Jurisdiktion",
    "article": "die"
  },
  {
    "noun": "Jurisprudenz",
    "article": "die"
  },
  {
    "noun": "Juristin",
    "article": "die"
  },
  {
    "noun": "Jurte",
    "article": "die"
  },
  {
    "noun": "Jury",
    "article": "die"
  },
  {
    "noun": "Justage",
    "article": "die"
  },
  {
    "noun": "Justierung",
    "article": "die"
  },
  {
    "noun": "Justierwaage",
    "article": "die"
  },
  {
    "noun": "Justifikation",
    "article": "die"
  },
  {
    "noun": "Justiz",
    "article": "die"
  },
  {
    "noun": "Justizbeamtin",
    "article": "die"
  },
  {
    "noun": "Justizministerin",
    "article": "die"
  },
  {
    "noun": "Justizverwaltung",
    "article": "die"
  },
  {
    "noun": "Justizvollzugsanstalt",
    "article": "die"
  },
  {
    "noun": "Jute",
    "article": "die"
  },
  {
    "noun": "Juxtaposition",
    "article": "die"
  },
  {
    "noun": "Charismatikerin",
    "article": "die"
  },
  {
    "noun": "Charmeuse",
    "article": "die"
  },
  {
    "noun": "Charta",
    "article": "die"
  },
  {
    "noun": "Chartergesellschaft",
    "article": "die"
  },
  {
    "noun": "Chartermaschine",
    "article": "die"
  },
  {
    "noun": "Chartreuse",
    "article": "die"
  },
  {
    "noun": "Chaussee",
    "article": "die"
  },
  {
    "noun": "Checkliste",
    "article": "die"
  },
  {
    "noun": "Chefin",
    "article": "die"
  },
  {
    "noun": "Chefingenieurin",
    "article": "die"
  },
  {
    "noun": "Chefredakteurin",
    "article": "die"
  },
  {
    "noun": "Chefsache",
    "article": "die"
  },
  {
    "noun": "Cheiloplastik",
    "article": "die"
  },
  {
    "noun": "Cheiloschisis",
    "article": "die"
  },
  {
    "noun": "Chemie",
    "article": "die"
  },
  {
    "noun": "Chemieanlage",
    "article": "die"
  },
  {
    "noun": "Chemiefaser",
    "article": "die"
  },
  {
    "noun": "Chemieindustrie",
    "article": "die"
  },
  {
    "noun": "Chemielaborantin",
    "article": "die"
  },
  {
    "noun": "Chemiewaffe",
    "article": "die"
  },
  {
    "noun": "Chemikalie",
    "article": "die"
  },
  {
    "noun": "Chemikantin",
    "article": "die"
  },
  {
    "noun": "Chemikerin",
    "article": "die"
  },
  {
    "noun": "Chemosis",
    "article": "die"
  },
  {
    "noun": "Chemotechnik",
    "article": "die"
  },
  {
    "noun": "Chemotechnikerin",
    "article": "die"
  },
  {
    "noun": "Chemotherapie",
    "article": "die"
  },
  {
    "noun": "Chenille",
    "article": "die"
  },
  {
    "noun": "Chiffre",
    "article": "die"
  },
  {
    "noun": "Chiffreschrift",
    "article": "die"
  },
  {
    "noun": "Chiffriermaschine",
    "article": "die"
  },
  {
    "noun": "Chiffrierung",
    "article": "die"
  },
  {
    "noun": "Chilenin",
    "article": "die"
  },
  {
    "noun": "Chinesin",
    "article": "die"
  },
  {
    "noun": "Chipkarte",
    "article": "die"
  },
  {
    "noun": "Chiromantie",
    "article": "die"
  },
  {
    "noun": "Chiropraktik",
    "article": "die"
  },
  {
    "noun": "Chirotherapie",
    "article": "die"
  },
  {
    "noun": "Chirurgie",
    "article": "die"
  },
  {
    "noun": "Chirurgin",
    "article": "die"
  },
  {
    "noun": "Chlorierung",
    "article": "die"
  },
  {
    "noun": "Chloritisierung",
    "article": "die"
  },
  {
    "noun": "Chlorverbindung",
    "article": "die"
  },
  {
    "noun": "Cholangitis",
    "article": "die"
  },
  {
    "noun": "Cholelithiasis",
    "article": "die"
  },
  {
    "noun": "Cholera",
    "article": "die"
  },
  {
    "noun": "Choleraepidemie",
    "article": "die"
  },
  {
    "noun": "Cholestase",
    "article": "die"
  },
  {
    "noun": "Cholezystitis",
    "article": "die"
  },
  {
    "noun": "Choreografie",
    "article": "die"
  },
  {
    "noun": "Choreografin",
    "article": "die"
  },
  {
    "noun": "Chorleiterin",
    "article": "die"
  },
  {
    "noun": "Chormusik",
    "article": "die"
  },
  {
    "noun": "Chorprobe",
    "article": "die"
  },
  {
    "noun": "Chose",
    "article": "die"
  },
  {
    "noun": "Chromatik",
    "article": "die"
  },
  {
    "noun": "Chromatografie",
    "article": "die"
  },
  {
    "noun": "Chromolithografie",
    "article": "die"
  },
  {
    "noun": "Chromosomenanomalie",
    "article": "die"
  },
  {
    "noun": "Chronik",
    "article": "die"
  },
  {
    "noun": "Chronistin",
    "article": "die"
  },
  {
    "noun": "Chronologie",
    "article": "die"
  },
  {
    "noun": "Chrysantheme",
    "article": "die"
  },
  {
    "noun": "Chuzpe",
    "article": "die"
  },
  {
    "noun": "Cineastin",
    "article": "die"
  },
  {
    "noun": "Cinemathek",
    "article": "die"
  },
  {
    "noun": "Circusartistin",
    "article": "die"
  },
  {
    "noun": "Circusdirektorin",
    "article": "die"
  },
  {
    "noun": "Circuskuppel",
    "article": "die"
  },
  {
    "noun": "Circusmanege",
    "article": "die"
  },
  {
    "noun": "Circusnummer",
    "article": "die"
  },
  {
    "noun": "Circusreiterin",
    "article": "die"
  },
  {
    "noun": "Circusvorstellung",
    "article": "die"
  },
  {
    "noun": "City",
    "article": "die"
  },
  {
    "noun": "Clementine",
    "article": "die"
  },
  {
    "noun": "Cleverness",
    "article": "die"
  },
  {
    "noun": "Clique",
    "article": "die"
  },
  {
    "noun": "Clubjacke",
    "article": "die"
  },
  {
    "noun": "Cocktailparty",
    "article": "die"
  },
  {
    "noun": "Codierung",
    "article": "die"
  },
  {
    "noun": "Collage",
    "article": "die"
  },
  {
    "noun": "Combo",
    "article": "die"
  },
  {
    "noun": "Compagnie",
    "article": "die"
  },
  {
    "noun": "Computeranimation",
    "article": "die"
  },
  {
    "noun": "Computergeneration",
    "article": "die"
  },
  {
    "noun": "Computergrafik",
    "article": "die"
  },
  {
    "noun": "Computerisierung",
    "article": "die"
  },
  {
    "noun": "Computerlinguistik",
    "article": "die"
  },
  {
    "noun": "Computermaus",
    "article": "die"
  },
  {
    "noun": "Computersimulation",
    "article": "die"
  },
  {
    "noun": "Computersprache",
    "article": "die"
  },
  {
    "noun": "Computertechnik",
    "article": "die"
  },
  {
    "noun": "Computertomografie",
    "article": "die"
  },
  {
    "noun": "Confiserie",
    "article": "die"
  },
  {
    "noun": "Copilotin",
    "article": "die"
  },
  {
    "noun": "Cordhose",
    "article": "die"
  },
  {
    "noun": "Corona",
    "article": "die"
  },
  {
    "noun": "Couch",
    "article": "die"
  },
  {
    "noun": "Couchgarnitur",
    "article": "die"
  },
  {
    "noun": "Courage",
    "article": "die"
  },
  {
    "noun": "Courtage",
    "article": "die"
  },
  {
    "noun": "Courtoisie",
    "article": "die"
  },
  {
    "noun": "Cousine",
    "article": "die"
  },
  {
    "noun": "Coverversion",
    "article": "die"
  },
  {
    "noun": "Cremetorte",
    "article": "die"
  },
  {
    "noun": "Crew",
    "article": "die"
  },
  {
    "noun": "Croquette",
    "article": "die"
  },
  {
    "noun": "Crux",
    "article": "die"
  },
  {
    "noun": "Curcuma",
    "article": "die"
  },
  {
    "noun": "Currywurst",
    "article": "die"
  },
  {
    "noun": "Cutis",
    "article": "die"
  },
  {
    "noun": "Cutterin",
    "article": "die"
  },
  {
    "noun": "Obsternte",
    "article": "die"
  },
  {
    "noun": "Obstetrik",
    "article": "die"
  },
  {
    "noun": "Obstination",
    "article": "die"
  },
  {
    "noun": "Obstipation",
    "article": "die"
  },
  {
    "noun": "Obstplantage",
    "article": "die"
  },
  {
    "noun": "Obstruktion",
    "article": "die"
  },
  {
    "noun": "Obstruktionspolitik",
    "article": "die"
  },
  {
    "noun": "Obsttorte",
    "article": "die"
  },
  {
    "noun": "Ochsenbrust",
    "article": "die"
  },
  {
    "noun": "Ochsentour",
    "article": "die"
  },
  {
    "noun": "Odaliske",
    "article": "die"
  },
  {
    "noun": "Odontologie",
    "article": "die"
  },
  {
    "noun": "Odyssee",
    "article": "die"
  },
  {
    "noun": "Ofenbank",
    "article": "die"
  },
  {
    "noun": "Ofenheizung",
    "article": "die"
  },
  {
    "noun": "Ofenkachel",
    "article": "die"
  },
  {
    "noun": "Ofensetzerin",
    "article": "die"
  },
  {
    "noun": "Offenbarung",
    "article": "die"
  },
  {
    "noun": "Offenheit",
    "article": "die"
  },
  {
    "noun": "Offenherzigkeit",
    "article": "die"
  },
  {
    "noun": "Offenkundigkeit",
    "article": "die"
  },
  {
    "noun": "Offenlegung",
    "article": "die"
  },
  {
    "noun": "Offenlegungspflicht",
    "article": "die"
  },
  {
    "noun": "Offenmarktpolitik",
    "article": "die"
  },
  {
    "noun": "Offensichtlichkeit",
    "article": "die"
  },
  {
    "noun": "Offensive",
    "article": "die"
  },
  {
    "noun": "Offensivwaffe",
    "article": "die"
  },
  {
    "noun": "Offerte",
    "article": "die"
  },
  {
    "noun": "Offizierin",
    "article": "die"
  },
  {
    "noun": "Offizierslaufbahn",
    "article": "die"
  },
  {
    "noun": "Offiziersmesse",
    "article": "die"
  },
  {
    "noun": "Offizin",
    "article": "die"
  },
  {
    "noun": "Ohnmacht",
    "article": "die"
  },
  {
    "noun": "Ohrenbeichte",
    "article": "die"
  },
  {
    "noun": "Ohrenheilkunde",
    "article": "die"
  },
  {
    "noun": "Ohrenklappe",
    "article": "die"
  },
  {
    "noun": "Ohrenrobbe",
    "article": "die"
  },
  {
    "noun": "Ohrfeige",
    "article": "die"
  },
  {
    "noun": "Ohrmarke",
    "article": "die"
  },
  {
    "noun": "Ohrmuschel",
    "article": "die"
  },
  {
    "noun": "Ohrtrompete",
    "article": "die"
  },
  {
    "noun": "Okarina",
    "article": "die"
  },
  {
    "noun": "Okkasion",
    "article": "die"
  },
  {
    "noun": "Okklusion",
    "article": "die"
  },
  {
    "noun": "Okkultistin",
    "article": "die"
  },
  {
    "noun": "Okkupation",
    "article": "die"
  },
  {
    "noun": "Okra",
    "article": "die"
  },
  {
    "noun": "Oktanzahl",
    "article": "die"
  },
  {
    "noun": "Oktave",
    "article": "die"
  },
  {
    "noun": "Oktoberrevolution",
    "article": "die"
  },
  {
    "noun": "Okulation",
    "article": "die"
  },
  {
    "noun": "Oligarchie",
    "article": "die"
  },
  {
    "noun": "Oligophrenie",
    "article": "die"
  },
  {
    "noun": "Oligurie",
    "article": "die"
  },
  {
    "noun": "Olive",
    "article": "die"
  },
  {
    "noun": "Olympiade",
    "article": "die"
  },
  {
    "noun": "Olympianorm",
    "article": "die"
  },
  {
    "noun": "Oma",
    "article": "die"
  },
  {
    "noun": "Omi",
    "article": "die"
  },
  {
    "noun": "Omnipotenz",
    "article": "die"
  },
  {
    "noun": "Omophagie",
    "article": "die"
  },
  {
    "noun": "Onanie",
    "article": "die"
  },
  {
    "noun": "Ondulation",
    "article": "die"
  },
  {
    "noun": "Oneiromantie",
    "article": "die"
  },
  {
    "noun": "Onkologie",
    "article": "die"
  },
  {
    "noun": "Onkologin",
    "article": "die"
  },
  {
    "noun": "Onomasiologie",
    "article": "die"
  },
  {
    "noun": "Onomatologie",
    "article": "die"
  },
  {
    "noun": "Ontogenese",
    "article": "die"
  },
  {
    "noun": "Ontogenie",
    "article": "die"
  },
  {
    "noun": "Ontologie",
    "article": "die"
  },
  {
    "noun": "Oophorektomie",
    "article": "die"
  },
  {
    "noun": "Oozyte",
    "article": "die"
  },
  {
    "noun": "Opaleszenz",
    "article": "die"
  },
  {
    "noun": "Oper",
    "article": "die"
  },
  {
    "noun": "Operateurin",
    "article": "die"
  },
  {
    "noun": "Operation",
    "article": "die"
  },
  {
    "noun": "Operationsbasis",
    "article": "die"
  },
  {
    "noun": "Operationsnarbe",
    "article": "die"
  },
  {
    "noun": "Operationsschwester",
    "article": "die"
  },
  {
    "noun": "Operette",
    "article": "die"
  },
  {
    "noun": "Operettenmusik",
    "article": "die"
  },
  {
    "noun": "Operndiva",
    "article": "die"
  },
  {
    "noun": "Opernmelodie",
    "article": "die"
  },
  {
    "noun": "Opferbereitschaft",
    "article": "die"
  },
  {
    "noun": "Opfergabe",
    "article": "die"
  },
  {
    "noun": "Opferkerze",
    "article": "die"
  },
  {
    "noun": "Opferrolle",
    "article": "die"
  },
  {
    "noun": "Opferschale",
    "article": "die"
  },
  {
    "noun": "Opferung",
    "article": "die"
  },
  {
    "noun": "Opferwilligkeit",
    "article": "die"
  },
  {
    "noun": "Ophthalmie",
    "article": "die"
  },
  {
    "noun": "Ophthalmologie",
    "article": "die"
  },
  {
    "noun": "Ophthalmoskopie",
    "article": "die"
  },
  {
    "noun": "Opportunistin",
    "article": "die"
  },
  {
    "noun": "Opposition",
    "article": "die"
  },
  {
    "noun": "Oppositionspartei",
    "article": "die"
  },
  {
    "noun": "Optik",
    "article": "die"
  },
  {
    "noun": "Optikerin",
    "article": "die"
  },
  {
    "noun": "Optimierung",
    "article": "die"
  },
  {
    "noun": "Optimistin",
    "article": "die"
  },
  {
    "noun": "Option",
    "article": "die"
  },
  {
    "noun": "Optionsanleihe",
    "article": "die"
  },
  {
    "noun": "Optionsschuldverschreibung",
    "article": "die"
  },
  {
    "noun": "Optoelektronik",
    "article": "die"
  },
  {
    "noun": "Optometrie",
    "article": "die"
  },
  {
    "noun": "Opulenz",
    "article": "die"
  },
  {
    "noun": "Orangeade",
    "article": "die"
  },
  {
    "noun": "Orangenmarmelade",
    "article": "die"
  },
  {
    "noun": "Orangenschale",
    "article": "die"
  },
  {
    "noun": "Orangenscheibe",
    "article": "die"
  },
  {
    "noun": "Orangerie",
    "article": "die"
  },
  {
    "noun": "Orbita",
    "article": "die"
  },
  {
    "noun": "Orbitalbahn",
    "article": "die"
  },
  {
    "noun": "Orchesterbegleitung",
    "article": "die"
  },
  {
    "noun": "Orchestermusik",
    "article": "die"
  },
  {
    "noun": "Orchestrierung",
    "article": "die"
  },
  {
    "noun": "Orchidee",
    "article": "die"
  },
  {
    "noun": "Orchitis",
    "article": "die"
  },
  {
    "noun": "Ordensregel",
    "article": "die"
  },
  {
    "noun": "Ordensschwester",
    "article": "die"
  },
  {
    "noun": "Ordentlichkeit",
    "article": "die"
  },
  {
    "noun": "Order",
    "article": "die"
  },
  {
    "noun": "Orderklausel",
    "article": "die"
  },
  {
    "noun": "Ordinalzahl",
    "article": "die"
  },
  {
    "noun": "Ordinate",
    "article": "die"
  },
  {
    "noun": "Ordinatenachse",
    "article": "die"
  },
  {
    "noun": "Ordination",
    "article": "die"
  },
  {
    "noun": "Ordnerin",
    "article": "die"
  },
  {
    "noun": "Ordnung",
    "article": "die"
  },
  {
    "noun": "Ordnungsliebe",
    "article": "die"
  },
  {
    "noun": "Ordnungspolitik",
    "article": "die"
  },
  {
    "noun": "Ordnungsstrafe",
    "article": "die"
  },
  {
    "noun": "Ordnungswidrigkeit",
    "article": "die"
  },
  {
    "noun": "Ordnungszahl",
    "article": "die"
  },
  {
    "noun": "Ordonanz",
    "article": "die"
  },
  {
    "noun": "Orfe",
    "article": "die"
  },
  {
    "noun": "Organbank",
    "article": "die"
  },
  {
    "noun": "Organelle",
    "article": "die"
  },
  {
    "noun": "Organentnahme",
    "article": "die"
  },
  {
    "noun": "Organisation",
    "article": "die"
  },
  {
    "noun": "Organisationsform",
    "article": "die"
  },
  {
    "noun": "Organisationsstruktur",
    "article": "die"
  },
  {
    "noun": "Organisatorin",
    "article": "die"
  },
  {
    "noun": "Organisierung",
    "article": "die"
  },
  {
    "noun": "Organistin",
    "article": "die"
  },
  {
    "noun": "Organkonservierung",
    "article": "die"
  },
  {
    "noun": "Organografie",
    "article": "die"
  },
  {
    "noun": "Organologie",
    "article": "die"
  },
  {
    "noun": "Organotherapie",
    "article": "die"
  },
  {
    "noun": "Organspenderin",
    "article": "die"
  },
  {
    "noun": "Organverpflanzung",
    "article": "die"
  },
  {
    "noun": "Orgel",
    "article": "die"
  },
  {
    "noun": "Orgelmusik",
    "article": "die"
  },
  {
    "noun": "Orgelpfeife",
    "article": "die"
  },
  {
    "noun": "Orgie",
    "article": "die"
  },
  {
    "noun": "Orientalin",
    "article": "die"
  },
  {
    "noun": "Orientalistik",
    "article": "die"
  },
  {
    "noun": "Orientalistin",
    "article": "die"
  },
  {
    "noun": "Orientierung",
    "article": "die"
  },
  {
    "noun": "Orientierungshilfe",
    "article": "die"
  },
  {
    "noun": "Orientierungsstufe",
    "article": "die"
  },
  {
    "noun": "Originalausgabe",
    "article": "die"
  },
  {
    "noun": "Originalfassung",
    "article": "die"
  },
  {
    "noun": "Originalsprache",
    "article": "die"
  },
  {
    "noun": "Originalverpackung",
    "article": "die"
  },
  {
    "noun": "Originalzeichnung",
    "article": "die"
  },
  {
    "noun": "Ornis",
    "article": "die"
  },
  {
    "noun": "Ornithologie",
    "article": "die"
  },
  {
    "noun": "Ornithologin",
    "article": "die"
  },
  {
    "noun": "Orogenese",
    "article": "die"
  },
  {
    "noun": "Orografie",
    "article": "die"
  },
  {
    "noun": "Orometrie",
    "article": "die"
  },
  {
    "noun": "Orphik",
    "article": "die"
  },
  {
    "noun": "Orthese",
    "article": "die"
  },
  {
    "noun": "Orthodontie",
    "article": "die"
  },
  {
    "noun": "Orthodoxie",
    "article": "die"
  },
  {
    "noun": "Orthodrome",
    "article": "die"
  },
  {
    "noun": "Orthoepie",
    "article": "die"
  },
  {
    "noun": "Orthofonie",
    "article": "die"
  },
  {
    "noun": "Orthogenese",
    "article": "die"
  },
  {
    "noun": "Orthografie",
    "article": "die"
  },
  {
    "noun": "Ortsangabe",
    "article": "die"
  },
  {
    "noun": "Ortsbeschreibung",
    "article": "die"
  },
  {
    "noun": "Ortsbesichtigung",
    "article": "die"
  },
  {
    "noun": "Ortsbestimmung",
    "article": "die"
  },
  {
    "noun": "Ortschaft",
    "article": "die"
  },
  {
    "noun": "Ortsgruppe",
    "article": "die"
  },
  {
    "noun": "Ortskenntnis",
    "article": "die"
  },
  {
    "noun": "Ortsnetzkennzahl",
    "article": "die"
  },
  {
    "noun": "Ortspolizei",
    "article": "die"
  },
  {
    "noun": "Ortstafel",
    "article": "die"
  },
  {
    "noun": "Ortsvorwahl",
    "article": "die"
  },
  {
    "noun": "Ortszeit",
    "article": "die"
  },
  {
    "noun": "Ortszulage",
    "article": "die"
  },
  {
    "noun": "Ortung",
    "article": "die"
  },
  {
    "noun": "Osmanin",
    "article": "die"
  },
  {
    "noun": "Osmose",
    "article": "die"
  },
  {
    "noun": "Osmotherapie",
    "article": "die"
  },
  {
    "noun": "Ossi",
    "article": "die"
  },
  {
    "noun": "Ossifikation",
    "article": "die"
  },
  {
    "noun": "Osteologie",
    "article": "die"
  },
  {
    "noun": "Osteolyse",
    "article": "die"
  },
  {
    "noun": "Osteomalazie",
    "article": "die"
  },
  {
    "noun": "Osteomyelitis",
    "article": "die"
  },
  {
    "noun": "Osteopathie",
    "article": "die"
  },
  {
    "noun": "Osteoporose",
    "article": "die"
  },
  {
    "noun": "Osteotomie",
    "article": "die"
  },
  {
    "noun": "Osterblume",
    "article": "die"
  },
  {
    "noun": "Ostereiersuche",
    "article": "die"
  },
  {
    "noun": "Osterglocke",
    "article": "die"
  },
  {
    "noun": "Osteria",
    "article": "die"
  },
  {
    "noun": "Osterwoche",
    "article": "die"
  },
  {
    "noun": "Osterzeit",
    "article": "die"
  },
  {
    "noun": "Ostpolitik",
    "article": "die"
  },
  {
    "noun": "Ostsee",
    "article": "die"
  },
  {
    "noun": "Ostwestrichtung",
    "article": "die"
  },
  {
    "noun": "Oszillation",
    "article": "die"
  },
  {
    "noun": "Otalgie",
    "article": "die"
  },
  {
    "noun": "Otiatrie",
    "article": "die"
  },
  {
    "noun": "Otitis",
    "article": "die"
  },
  {
    "noun": "Otologie",
    "article": "die"
  },
  {
    "noun": "Otologin",
    "article": "die"
  },
  {
    "noun": "Otosklerose",
    "article": "die"
  },
  {
    "noun": "Otoskopie",
    "article": "die"
  },
  {
    "noun": "Otter",
    "article": "die"
  },
  {
    "noun": "Ovation",
    "article": "die"
  },
  {
    "noun": "Ovoviviparie",
    "article": "die"
  },
  {
    "noun": "Ovulation",
    "article": "die"
  },
  {
    "noun": "Oxidation",
    "article": "die"
  },
  {
    "noun": "Oxidierung",
    "article": "die"
  },
  {
    "noun": "Ozeanografie",
    "article": "die"
  },
  {
    "noun": "Ozonbelastung",
    "article": "die"
  },
  {
    "noun": "Ozonschicht",
    "article": "die"
  },
  {
    "noun": "Immoralistin",
    "article": "die"
  },
  {
    "noun": "Immortelle",
    "article": "die"
  },
  {
    "noun": "Immunologie",
    "article": "die"
  },
  {
    "noun": "Immunreaktion",
    "article": "die"
  },
  {
    "noun": "Impedanz",
    "article": "die"
  },
  {
    "noun": "Impertinenz",
    "article": "die"
  },
  {
    "noun": "Impetigo",
    "article": "die"
  },
  {
    "noun": "Impfpflicht",
    "article": "die"
  },
  {
    "noun": "Impfung",
    "article": "die"
  },
  {
    "noun": "Implantation",
    "article": "die"
  },
  {
    "noun": "Implementation",
    "article": "die"
  },
  {
    "noun": "Implementierung",
    "article": "die"
  },
  {
    "noun": "Implikation",
    "article": "die"
  },
  {
    "noun": "Implosion",
    "article": "die"
  },
  {
    "noun": "Imponderabilie",
    "article": "die"
  },
  {
    "noun": "Importe",
    "article": "die"
  },
  {
    "noun": "Importfirma",
    "article": "die"
  },
  {
    "noun": "Importlizenz",
    "article": "die"
  },
  {
    "noun": "Importquote",
    "article": "die"
  },
  {
    "noun": "Importware",
    "article": "die"
  },
  {
    "noun": "Impotenz",
    "article": "die"
  },
  {
    "noun": "Impression",
    "article": "die"
  },
  {
    "noun": "Impressionistin",
    "article": "die"
  },
  {
    "noun": "Improvisation",
    "article": "die"
  },
  {
    "noun": "Improvisatorin",
    "article": "die"
  },
  {
    "noun": "Impulsdauer",
    "article": "die"
  },
  {
    "noun": "Impulsfrequenz",
    "article": "die"
  },
  {
    "noun": "Impulstechnik",
    "article": "die"
  },
  {
    "noun": "Inaktivierung",
    "article": "die"
  },
  {
    "noun": "Inangriffnahme",
    "article": "die"
  },
  {
    "noun": "Inanition",
    "article": "die"
  },
  {
    "noun": "Inanspruchnahme",
    "article": "die"
  },
  {
    "noun": "Inauguraldissertation",
    "article": "die"
  },
  {
    "noun": "Inauguration",
    "article": "die"
  },
  {
    "noun": "Inaugurierung",
    "article": "die"
  },
  {
    "noun": "Inbesitznahme",
    "article": "die"
  },
  {
    "noun": "Inbetriebnahme",
    "article": "die"
  },
  {
    "noun": "Inbetriebsetzung",
    "article": "die"
  },
  {
    "noun": "Inbrunst",
    "article": "die"
  },
  {
    "noun": "Inbusschraube",
    "article": "die"
  },
  {
    "noun": "Independenz",
    "article": "die"
  },
  {
]