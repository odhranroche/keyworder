<pre>
/*
 * Description: This program is intended to extract the minimum number of words
 *              needed to understand the text. Duplicate words, similar words,
 *              short words, and proper nouns are removed. The resulting list 
 *              forms the foundational words of the text.
 *
 * Usage: go run main.go textUtils.go jaro.go
 *        - input file should be in the same directory 
 *        - input file should be specified in main.go
 *        - number of times a word is capitalized to count as a proper noun
 *
 * Issues: - Word similarity algorithm sometimes deletes words
             with different meanings
           - Suggested proper nouns sometimes picks words that
             are at the start of a sentence
           - Words with apostrophies are split haven't -> haven t  
 */
 </pre>
 
## Example (167 words):
It was on the corner of the street that he noticed the first sign of something peculiar -- a cat reading a map. For a second, Mr. Dursley didn't realize what he had seen -- then he jerked his head around to look again. There was a tabby cat standing on the corner of Privet Drive, but there wasn't a map in sight. What could he have been thinking of? It must have been a trick of the light. Mr. Dursley blinked and stared at the cat. It stared back. As Mr. Dursley drove around the corner and up the road, he watched the cat in his mirror. It was now reading the sign that said Privet Drive -- no, looking at the sign; cats couldn't read maps or signs. Mr. Dursley gave himself a little shake and put the cat out of his mind. As he drove toward town he thought of nothing except a large order of drills he was hoping to get that day. 
 
## Output (54 words):
realize
thinking
blinked
stared
himself
road
gave
corner
been
wasn
sight
mirror
looking
something
reading
large
noticed
didn
have
trick
street
that
head
read
nothing
second
look
cats
watched
light
order
around
little
seen
drills
peculiar
then
again
first
toward
standing
maps
must
drove
jerked
back
said
shake
thought
hoping
tabby
town
except
mind

The program removed duplicate words, short words, similar words like:
* couldn't could
* sign signs
* could couldn't

And proper nouns like Dursley, Grunnings, Privet Drive 
