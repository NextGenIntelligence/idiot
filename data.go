package main

var nonterminals = map[string][]string{
	"sentence": {
		"{in_adverbial_modifier} {pos=verb number=@2} {extended_objects case=@1 number=@2}, {participle_phrase case=@1 number=@2}.",
	},
	"in_adverbial_modifier": {
		"в {extended_objects case=loct}",
	},
	"extended_objects": {
		"{extended_object !case !number}",
		"@{number=plur}{extended_object !case=@1}, {extended_objects case=@1}",
		"@{number=plur}{extended_object !case=@1} и {extended_objects case=@1}",
	},
	"extended_object": {
		"{pos=noun !case !number}",
		"{pos=adjf case=@1 number=@2} {pos=noun !case=@1 !number=@2}",
		"{pos=adjf case=@1 number=@2} {pos=noun !case=@1 !number=@2} {extended_objects case=gent}",
	},
	"participle_phrase": {
		"{pos=prtf !case !number} {extended_objects case=ablt}",
	},
}

func FindNonterminalRules(nonterminal string) []string {
	return nonterminals[nonterminal]
}
