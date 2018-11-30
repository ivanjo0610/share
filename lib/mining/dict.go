// Copyright 2018 Shulhan <ms@kilabit.info>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mining

// VulgarWords contain list of vulgar and offensive words in informal and
// slangs.
// nolint: misspell
var VulgarWords = []string{ // nolint: gochecknoglobals
	"$#!+", "$1ut", "$h1t", "$hit", "$lut", "'f*ck'", "'ho", "'hobag",
	"@ss", "@sshole", "a$$", "a$$h0!e", "a$$h01e", "a$$h0le", "a$$hole",
	"a55", "a55hole", "aeolus", "ahole", "analprobe", "anilingus",
	"anorexia", "anorexic", "anus", "areola", "areole", "arian", "aryan",
	"ass", "assbang", "assbanged", "assbangs", "asses", "assfuck",
	"assfucker", "assfuckers", "assh0le", "assho!e", "assho1e", "asshole",
	"assholes", "assmaster", "assmunch", "asswipe", "asswipes", "azazel",
	"azz", "b1tch", "b1tch", "b@lls", "baal", "babe", "babes", "ballsack",
	"bang", "banger", "barf", "bastard", "bastards", "bawdy", "beaner",
	"beardedclam", "beastiality", "beatch", "beater", "beaver", "beer",
	"beeyotch", "beotch", "biatch", "bigtits", "bimbo", "bitch",
	"bitched", "bitches", "bitchface", "bitchy", "blew", "blow", "blow",
	"blowjob", "blowjobs", "blowup", "bod", "bodily", "boink", "bollock",
	"bollocks", "bollok", "bone", "boned", "boner", "boners", "bong",
	"boob", "boobies", "boobs", "booby", "booger", "bookie", "booky",
	"bootee", "bootie", "booty", "booze", "boozer", "boozy", "bosom",
	"bosomy", "bowel", "bowels", "bra", "brassiere", "bugger", "bukkake",
	"bulimia", "bulimiic", "bullshit", "bullshits", "bullshitted",
	"bullturds", "bung", "burp", "bush", "busty", "butt", "buttfuck",
	"buttfucker", "buttplug", "c-0-c-k", "c-o-c-k", "c-u-n-t", "c.0.c.k",
	"c.o.c.k.", "c.u.n.", "c0ck", "caca", "cahone", "cameltoe", "carnal",
	"carpetmuncher", "cawk", "cervix", "chinc", "chincs", "chink",
	"chink", "chode", "chodes", "cl1t", "climax", "clit", "clit",
	"clitoris", "clitorus", "clits", "clitty", "cocain", "cocaine",
	"cock", "cockblock", "cockholster", "cockknocker", "cocks",
	"cocksucker", "coital", "coke", "commie", "condom", "coon", "coons",
	"copulator", "corksucker", "corpse", "coven", "crabs", "crack",
	"cracker", "crackwhore", "crap", "crappy", "cuervo", "cum", "cummin",
	"cumming", "cumshot", "cumshots", "cumstain", "cunilingus",
	"cunnilingus", "cunny", "cunt", "cuntface", "cunthunter", "cuntlick",
	"cuntlicker", "cunts", "d0ng", "d0uch3", "d0uche", "d1ck", "d1ck",
	"d1ld0", "d1ldo", "dago", "dagos", "dammit", "damn", "damned",
	"damnit", "dawgie-style", "dick", "dick-ish", "dickbag", "dickdipper",
	"dickflipper", "dickhead", "dickheads", "dickish", "dickripper",
	"dicksipper", "dickweed", "dickwhipper", "dickzipper", "diddle",
	"dike", "diligaf", "dillweed", "dimwit", "dingle", "doggie",
	"doggie-style", "doggy", "doggy-style", "dong", "doofus", "doosh",
	"dopey", "douch3", "douche", "douche", "douchebag", "douchebags",
	"douchey", "drunk", "dumass", "dumbass", "dumbasses", "dummy", "dyke",
	"dykes", "ejaculate", "enlargement", "erect", "erotic", "essohbee",
	"exotic", "extacy", "extasy", "f-u-c-k", "f.u.c.k", "f@g", "f@gg0t",
	"f@ggot", "fack", "faerie", "faery", "fag", "fagg", "fagged",
	"faggit", "faggot", "fagot", "fags", "faig", "faigt", "fairy",
	"fanny", "fannybandit", "fart", "fartknocker", "felch", "felcher",
	"felching", "fellate", "fellatio", "fellatio", "feltch", "feltcher",
	"fisted", "fisting", "fisty", "floozy", "foad", "fondle", "foobar",
	"foreskin", "frack", "freex", "frigg", "frigga", "fu*ck", "fubar",
	"fuck", "fuck-tard", "fucka$$", "fuckass", "fucked", "fucker",
	"fucker", "fuckers", "fuckface", "fucking", "fucknugget", "fucknut",
	"fuckoff", "fucks", "fucktard", "fuckup", "fuckwad", "fuckwit",
	"fudgepacker", "fuk", "futs", "fvck", "fxck", "g-spot", "g@y", "gae",
	"gai", "ganja", "gay", "gays", "gey", "gfy", "ghay", "ghey", "gigolo",
	"glans", "goatse", "godamn", "godamnit", "godamnit", "goddam",
	"goddammit", "goddamn", "goldenshower", "gonad", "gonads", "gonads",
	"gook", "gooks", "gspot", "gtfo", "h0m0", "h0mo", "handjob", "he11",
	"hebe", "hell", "hemp", "heroin", "herp", "herpes", "herpy", "hijack",
	"hitler", "hiv", "hobag", "hom0", "homey", "homo", "homoey", "honky",
	"hooch", "hookah", "hooker", "hoor", "hootch", "hooter", "hooters",
	"horny", "hump", "humped", "humper", "humping", "hussy", "hymen",
	"idiot", "idiots", "inbred", "incest", "injun", "j3rk0ff", "jackass",
	"jackhole", "jackoff", "jap", "japs", "jerk", "jerk", "jerk0ff",
	"jerked", "jerkoff", "jism", "jiz", "jizm", "jizm", "jizz", "jizzed",
	"junkie", "junky", "kike", "kikes", "kill", "kinky", "kkk", "klan",
	"knobend", "knobend", "kooch", "kooches", "kootch", "kraut", "kyke",
	"labia", "lech", "leper", "lesbians", "lesbos", "lez", "lezbian",
	"lezbians", "lezbo", "lezbos", "lezzie", "lezzies", "lezzy", "licker",
	"lickers", "lmao", "lmfao", "loin", "loins", "lsd", "lube", "lust",
	"lusty", "m-fucking", "mams", "marijuana", "massa", "masterbate",
	"masterbating", "masterbation", "masturbate", "masturbating",
	"masturbation", "maxi", "menses", "menstruate", "menstruation",
	"meth", "mofo", "molest", "moolie", "moron", "motherf*cker",
	"motherfucka", "motherfucker", "motherfucking", "mtherfucker",
	"mthrf*cker", "mthrfucker", "mthrfucking", "muff", "murder",
	"muthafuckaz", "muthafucker", "mutherfucker", "mutherfucking",
	"muthrfucking", "nad", "nads", "naked", "napalm", "nappy", "nazi",
	"nazism", "negro", "nigga", "niggah", "niggas", "niggaz", "nigger",
	"niggers", "niggle", "nimrod", "ninny", "nooky", "nucking", "nympho",
	"opiate", "opium", "oral", "orally", "organ", "orgasm", "orgasmic",
	"orgies", "orgy", "ovary", "ovum", "ovums", "p*ssy", "p.u.s.s.y.",
	"p.u.s.s.y.", "paddy", "paedophile", "pantie", "panties", "panty",
	"pastie", "pasty", "pcp", "pecker", "pedo", "pedophile", "pedophile",
	"pedophiles", "pedophilia", "pedophiliac", "pee", "peepee",
	"penetrate", "penetration", "penial", "penile", "perversion",
	"peyote", "phalli", "phallic", "phuck", "pillowbiter", "pimp",
	"pinko", "piss", "piss-off", "pissed", "pissoff", "pms", "polack",
	"porn", "porno", "pornography", "pot", "potty", "prick", "prig",
	"prostitute", "prude", "pu$$y", "pube", "pubic", "pubis", "punkass",
	"punky", "puss", "pussies", "pussy", "pussypounder", "pussys",
	"queaf", "queef", "queefing", "queer", "queero", "queers", "quicky",
	"quife", "quim", "r-tard", "racist", "racy", "rape", "raped", "raper",
	"rapist", "raunch", "rectal", "rectum", "rectus", "reefer", "reetard",
	"reich", "retard", "retarded", "revue", "rim", "risque", "ritard",
	"rtard", "rum", "rump", "rumprammer", "s***", "s*o*b", "s-h-1-t",
	"s-h-i-t", "s-o-b", "s.h.i.t.", "s.o.b.", "s0b", "sadism", "sadist",
	"satan", "scag", "scantily", "schizo", "screw", "screwed", "scrog",
	"scrot", "scrote", "scrotum", "scrud", "scum", "seaman", "seamen",
	"seduce", "semen", "sex_story", "sexual", "sh!t", "sh*t", "sh1t",
	"shamedame", "shit", "shite", "shiteater", "shitface", "shithead",
	"shithole", "shithouse", "shits", "shitt", "shitted", "shitter",
	"shitty", "shiz", "sissy", "skag", "slave", "sleaze", "sleazy",
	"slut", "slutdumper", "slutkiss", "sluts", "smegma", "smut", "smutty",
	"snatch", "sniper", "snuff", "sodom", "souse", "soused", "sperm",
	"spic", "spick", "spik", "spiks", "spooge", "spunk", "stab", "steamy",
	"stfu", "stiffy", "stoned", "strip", "stroke", "suck", "sucked",
	"sucking", "sucks", "sumofabiatch", "t1t", "tampon", "tard", "tawdry",
	"teabagging", "teat", "terd", "teste", "testee", "testes", "testis",
	"thrust", "thug", "tinkle", "tit", "titfuck", "titi", "tits",
	"tittiefucker", "titties", "titty", "tittyfuck", "tittyfucker",
	"toke", "toots", "tosser", "tossers", "tramp", "transsexual",
	"trashy", "tubgirl", "turd", "tush", "twat", "twats", "undies",
	"unwed", "urinal", "urine", "uterus", "uzi", "vag", "valium",
	"viagra", "virgin", "vixen", "vodka", "vomit", "voyeur", "vulgar",
	"vulva", "w@ng", "wad", "wang", "wank", "wazoo", "wedgie", "weed",
	"weenie", "weewee", "weiner", "weirdo", "wench", "wetback",
	"wh0r3f@ce", "wh0re", "wh0ref@ce", "wh0reface", "whack", "whacked",
	"whacking", "whitey", "whiz", "whoralicious", "whore",
	"whorealicious", "whored", "whoreface", "whorehopper", "whorehouse",
	"whores", "whoring", "wigger", "womb", "woody", "wop", "wtf",
	"x-rated", "xxx", "yeasty", "yobbo", "zoophile",
}

// PronounWords contains list of first and second person pronouns including
// slangs.
var PronounWords = []string{ // nolint: gochecknoglobals
	"i", "me", "mine", "my", "myself", "our", "ours", "ourself",
	"ourselves", "selves", "thee", "thine", "thou", "thy", "thyself",
	"us", "we", "y'all", "y'all", "y'all's", "yis", "you", "you-uns",
	"your", "yours", "yourself", "yourselves", "yourselves", "yous",
	"yous's", "youse", "youse",
}

// BiasedWords contain list of colloquial words with high bias.
var BiasedWords = []string{ // nolint: gochecknoglobals
	"cutting-edge", "single-handedly", "well-established", "well-known",
	"world-class", "absolute", "acclaimed", "amazing", "astonishing",
	"authoritative", "beautiful", "best", "boreing", "boring",
	"brilliant", "canonical", "cares", "celebrated", "charismatic",
	"classic", "coolest", "defining", "definitive", "eminent", "enigma",
	"ever", "everyone", "exciting", "extraordinary", "fabulous", "famous",
	"fantastic", "fat", "fully", "genius", "global", "great", "greatest",
	"hate", "huge", "iconic", "idiotic", "immensely", "impactful",
	"incendiary", "indisputable", "infamous", "influential", "innovative",
	"inspired", "intriguing", "lame", "leader", "leading", "legendary",
	"like", "major", "masterly", "mature", "memorable", "most", "notable",
	"outstanding", "pioneer", "popular", "prestigious", "probably",
	"really", "remarkable", "renowned", "respected", "seminal",
	"significant", "skillful", "solution", "staunch", "strange", "super",
	"talented", "top", "total", "totally", "transcendent", "ugly",
	"undoubtedly", "unique", "virtually", "virtuoso", "visionary",
	"weird", "worst",
}

// SexWords contain list of non-vulgar sex-related words.
var SexWords = []string{ // nolint: gochecknoglobals
	"anal", "breast", "breasts", "buttocks", "dildo", "dildos", "erect",
	"nipple", "nipples", "penis", "sex", "sodomized", "sodomy", "vagina",
	"vibrator", "vibrators",
}

// BadWords contain list of colloquial words or bad writing words.
var BadWords = []string{ // nolint: gochecknoglobals
	"666", "da", "dont", "dosent", "whatever", "guy", "hi", "nazi", "sup",
	"guise", "loser", "thats", "ugly", "wanna", "whats", "wont", "gotta",
	"bloody", "fart", "pot", "prick", "stink", "smells", "smelly", "alot",
	"dunno", "gotcha",
}
