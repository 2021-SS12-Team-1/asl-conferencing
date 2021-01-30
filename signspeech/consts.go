package signspeech

// CC conjunction: a word used to connect clauses or sentences
//			   or to coordinate words in the same clause
//			   e.g. and, but, if
// coordinating:  conjunction placed between words, phrases,
//				 clauses, or sentences of equal rank,
//				 e.g. and, but, or
var CC string = "CC"

// CD cardinal number: natural number
var CD string = "CD"

// DT determiner: a modifying word that determines the kind of reference
// 			  a noun or noun group has
// 			  e.g a, the, every
var DT string = "DT"

// EX existential there: a there that comes at the start of a sentence
var EX string = "EX"

// FW foreign word
var FW string = "FW"

// IN subordinating conjunction: A subordinating conjunction is a word or
// 			phrase that links a dependent clause to an independent clause
//			e.g. once, while, when, whenever, where, wherever, before, after
// preoposition: a word governing, and usually preceding, a noun or pronoun
//				and expressing a relation to another word or element in the
//				clause
//				e.g. “the man on the platform,” “she arrived after dinner,”
//					 “what did you do it for ?”
var IN string = "IN"

// JJ adjective
var JJ string = "JJ"

// JJR comparitive adjective: adjectives ending in -er
//						 e.g. better
var JJR string = "JJR"

// JJS superlative adjective: adjectives ending in -est
//						 e.g. best
var JJS string = "JJS"

// LS list item marker: an item marker for things in lists,
//					e.g. Things to do: 1) ... 2) ...
var LS string = "LS"

// MD modal auxillary verb: a verb that is usually used with another verb to express
//						ideas such as possibility, necessity, and permission
//						e.g. can, could, shall, should, ought to, will, or would
var MD string = "MD"

// NN noun, singular or mass
var NN string = "NN"

// NNP noun, proper singular
var NNP string = "NNP"

// NNPS noun, proper plural
var NNPS string = "NNPS"

// NNS noun, plural
var NNS string = "NNS"

// PDT predeterminer: a word or phrase that occurs before a determiner, typically
//				 quantifying the noun phrase
//				 e.g. both or a lot of.
var PDT string = "PDT"

// POS possessive ending: an apostrophe at the end of plural nouns
var POS string = "POS"

// PRP personal pronoun: I, you, he, she, it, we, they, me, him, her, us, them
var PRP string = "PRP"

// PRPS possessive pronoun: mine, yours, hers, theirs
var PRPS string = "PRPS"

// RB adverb: a word or phrase that modifies or qualifies an adjective, verb, or
//		  other adverb or a word group, expressing a relation of place, time,
//		  circumstance, manner, cause, degree, etc.
//		  e.g., gently, quite, then, there
var RB string = "RB"

// RBR comparative adverb: comparing two actions
var RBR string = "RBR"

// RBS superlative adverb: comparing three or more actions
var RBS string = "RBS"

// RP adverb particle: some verbs are followed by adverb particles
//				   e.g. put on, take off, give away, bring up, call in.
var RP string = "RP"

// SYM symbol
var SYM string = "SYM"

// TO infinitival to
var TO string = "TO"

// UH interjection: demonstrates the emotion or feeling of the author.
//				These words or phrases can stand alone, or be placed before or
//				after a sentenc
var UH string = "UH"

// VB verb, base form: the version of the verb without any endings (-s, -ing, and ed)
//				   The base form is the same as the infinitive (e.g., to walk,
//					to paint, to think) but without the to
var VB string = "VB"

// VBD verb, past tense (-ed)
var VBD string = "VBD"

// VBG verb, gerund or present participle (-ing)
var VBG string = "VBG"

// VBN verb, past participle: The past participle, also sometimes called the passive or
//		perfect participle, is identical to the past tense form (ending in -ed) in
//		the case of regular verbs, for example "loaded", "boiled", "mounted", but
// 		takes various forms in the case of irregular verbs, such as done, sung,
//		written, put, gone, etc.
var VBN string = "VBN"

// VBP verb, non-3rd person singular present: When a verb is used with subject like
//		I, You, They, We, if it is not in a "infinitive", "imperative", and
//		"subjunctive" condition
var VBP string = "VBP"

// VBZ verb, 3rd person singular present: In present simple, the verb changes only
//		in third person singular (he, she, it, a person, a thing), where it
// 		gets the suffix -s or -es
var VBZ string = "VBZ"

// WDT wh-determiner: what, which
var WDT string = "WDT"

// WP wh-pronoun, personal: who, whether, to what, which
var WP string = "WP"

// WPS wh-pronoun, possessive: whose
var WPS string = "WP$"

// WRB wh-adverb: how, when, whence, where, why
var WRB string = "WRB"
