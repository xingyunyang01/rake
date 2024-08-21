package rake

// stop word list from SMART (Salton,1971).  Available at ftp://ftp.cs.cornell.edu/pub/smart/english.stop
var StopWordsSlice = []string{
	"'s",
	"a",
	"able",
	"about",
	"above",
	"according",
	"accordingly",
	"across",
	"actually",
	"after",
	"afterwards",
	"again",
	"against",
	"all",
	"allow",
	"allows",
	"almost",
	"alone",
	"along",
	"already",
	"also",
	"although",
	"always",
	"am",
	"among",
	"amongst",
	"an",
	"and",
	"another",
	"any",
	"anybody",
	"anyhow",
	"anyone",
	"anything",
	"anyway",
	"anyways",
	"anywhere",
	"apart",
	"appear",
	"appreciate",
	"appropriate",
	"are",
	"aren't",
	"around",
	"as",
	"aside",
	"ask",
	"asking",
	"associated",
	"at",
	"autoscaler",
	"available",
	"away",
	"awfully",
	"be",
	"became",
	"because",
	"become",
	"becomes",
	"becoming",
	"been",
	"before",
	"beforehand",
	"behind",
	"being",
	"believe",
	"below",
	"beside",
	"besides",
	"best",
	"better",
	"between",
	"beyond",
	"both",
	"brief",
	"but",
	"by",
	"came",
	"can",
	"can't",
	"cannot",
	"cant",
	"cause",
	"causes",
	"certain",
	"certainly",
	"changes",
	"clearly",
	"codename",
	"com",
	"come",
	"comes",
	"concerning",
	"concise",
	"configuration",
	"consequently",
	"consider",
	"considering",
	"contain",
	"containing",
	"contains",
	"corresponding",
	"could",
	"couldn't",
	"course",
	"currently",
	"definitely",
	"described",
	"despite",
	"did",
	"didn't",
	"different",
	"do",
	"does",
	"doesn't",
	"doing",
	"don't",
	"done",
	"down",
	"downwards",
	"during",
	"each",
	"eight",
	"either",
	"else",
	"elsewhere",
	"enough",
	"entirely",
	"especially",
	"entry",
	"even",
	"ever",
	"every",
	"everybody",
	"everyone",
	"everything",
	"everywhere",
	"exactly",
	"example",
	"except",
	"far",
	"few",
	"feature",
	"field",
	"followed",
	"following",
	"follows",
	"for",
	"former",
	"formerly",
	"forth",
	"four",
	"from",
	"function",
	"functionality",
	"further",
	"furthermore",
	"get",
	"gets",
	"getting",
	"given",
	"gives",
	"go",
	"goes",
	"going",
	"gone",
	"got",
	"gotten",
	"greetings",
	"had",
	"hadn't",
	"happens",
	"hardly",
	"has",
	"hasn't",
	"have",
	"haven't",
	"having",
	"he",
	"he's",
	"hello",
	"help",
	"hence",
	"her",
	"here",
	"here's",
	"hereafter",
	"hereby",
	"herein",
	"hereupon",
	"hers",
	"herself",
	"hi",
	"him",
	"himself",
	"his",
	"hither",
	"hopefully",
	"how",
	"however",
	"horizontal",
	"i",
	"if",
	"ignored",
	"immediate",
	"in",
	"inasmuch",
	"inc",
	"indeed",
	"indicate",
	"indicated",
	"indicates",
	"information",
	"inner",
	"insofar",
	"instead",
	"into",
	"inward",
	"is",
	"isn't",
	"issuing",
	"issuance",
	"it",
	"it's",
	"its",
	"itself",
	"just",
	"keep",
	"keeps",
	"kept",
	"know",
	"knows",
	"known",
	"last",
	"lately",
	"later",
	"latter",
	"latterly",
	"least",
	"less",
	"lest",
	"let",
	"let's",
	"like",
	"liked",
	"likely",
	"little",
	"login",
	"look",
	"looking",
	"looks",
	"mainly",
	"many",
	"may",
	"maybe",
	"me",
	"message",
	"mean",
	"meanwhile",
	"merely",
	"might",
	"more",
	"moreover",
	"most",
	"mostly",
	"much",
	"must",
	"my",
	"myself",
	"namely",
	"near",
	"nearly",
	"necessary",
	"need",
	"needs",
	"neither",
	"never",
	"nevertheless",
	"new",
	"next",
	"nine",
	"nutshell",
	"no",
	"nobody",
	"non",
	"none",
	"noone",
	"nor",
	"normally",
	"not",
	"nothing",
	"novel",
	"now",
	"nowhere",
	"obviously",
	"of",
	"off",
	"often",
	"oh",
	"ok",
	"okay",
	"old",
	"on",
	"once",
	"one",
	"ones",
	"only",
	"onto",
	"or",
	"other",
	"others",
	"otherwise",
	"ought",
	"outline",
	"our",
	"ours",
	"ourselves",
	"out",
	"outside",
	"over",
	"overall",
	"own",
	"particular",
	"particularly",
	"per",
	"perhaps",
	"placed",
	"please",
	"plus",
	"pod",
	"pods",
	"possible",
	"presumably",
	"probably",
	"provides",
	"que",
	"quite",
	"rather",
	"really",
	"reasonably",
	"regarding",
	"regardless",
	"regards",
	"relatively",
	"respectively",
	"right",
	"s",
	"said",
	"same",
	"saw",
	"say",
	"saying",
	"says",
	"second",
	"secondly",
	"see",
	"seeing",
	"seem",
	"seemed",
	"seeming",
	"seems",
	"seen",
	"self",
	"selves",
	"sensible",
	"sentence",
	"sent",
	"serious",
	"seriously",
	"seven",
	"several",
	"shall",
	"sky",
	"skyWalking",
	"she",
	"should",
	"shouldn't",
	"shortest",
	"since",
	"six",
	"so",
	"some",
	"somebody",
	"somehow",
	"someone",
	"something",
	"sometime",
	"sometimes",
	"somewhat",
	"somewhere",
	"soon",
	"sorry",
	"specified",
	"step",
	"steps",
	"still",
	"sub",
	"such",
	"sup",
	"sure",
	"system",
	"take",
	"taken",
	"tell",
	"tends",
	"than",
	"thanx",
	"that",
	"that's",
	"thats",
	"the",
	"their",
	"theirs",
	"them",
	"themselves",
	"then",
	"thence",
	"there",
	"there's",
	"thereafter",
	"thereby",
	"therefore",
	"therein",
	"theres",
	"thereupon",
	"these",
	"they",
	"think",
	"third",
	"this",
	"thorough",
	"thoroughly",
	"those",
	"though",
	"three",
	"through",
	"throughout",
	"thus",
	"to",
	"together",
	"too",
	"took",
	"toward",
	"towards",
	"tried",
	"tries",
	"truly",
	"try",
	"trying",
	"twice",
	"un",
	"under",
	"unfortunately",
	"unless",
	"unlikely",
	"until",
	"unto",
	"up",
	"upon",
	"us",
	"use",
	"used",
	"useful",
	"uses",
	"using",
	"usually",
	"value",
	"various",
	"very",
	"via",
	"viz",
	"want",
	"wants",
	"was",
	"wasn't",
	"way",
	"walking",
	"welcome",
	"well",
	"went",
	"were",
	"weren't",
	"what",
	"what's",
	"whatever",
	"when",
	"whence",
	"whenever",
	"where",
	"where's",
	"whereafter",
	"whereas",
	"whereby",
	"wherein",
	"whereupon",
	"wherever",
	"whether",
	"which",
	"while",
	"whither",
	"why",
	"will",
	"willing",
	"wish",
	"with",
	"within",
	"without",
	"won't",
	"wonder",
	"would",
	"would",
	"wouldn't",
	"yes",
	"yet",
	"briefest",
	"consice",
	"description",
	"name",
	"simplest",
	"file",
	"files",
	"settings",
	"entries",
	"length",
	"designation",
	"code",
}
