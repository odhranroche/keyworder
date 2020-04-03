## Description
This program is intended to extract the minimum number of words needed to understand the text. Duplicate words, similar words, short words, and proper nouns are removed. The resulting list forms the foundational words of the text.

## Usage
<pre>> go run main.go textUtils.go jaro.go</pre>
- input file should be in the same directory 
- input file should be specified in main.go
- number of times a word is capitalized to count as a proper noun

## Issues
- Word similarity algorithm sometimes deletes words with different meanings
- Suggested proper nouns sometimes picks words that are at the start of a sentence
- Words with apostrophies are split haven't -> haven t  
 
## Example (167 words):
It was on the corner of the street that he noticed the first sign of something peculiar -- a cat reading a map. For a second, Mr. Dursley didn't realize what he had seen -- then he jerked his head around to look again. There was a tabby cat standing on the corner of Privet Drive, but there wasn't a map in sight. What could he have been thinking of? It must have been a trick of the light. Mr. Dursley blinked and stared at the cat. It stared back. As Mr. Dursley drove around the corner and up the road, he watched the cat in his mirror. It was now reading the sign that said Privet Drive -- no, looking at the sign; cats couldn't read maps or signs. Mr. Dursley gave himself a little shake and put the cat out of his mind. As he drove toward town he thought of nothing except a large order of drills he was hoping to get that day. 
 
## Output (56 words):
noticed
except
first
reading
then
standing
have
said
maps
street
realize
drove
toward
nothing
back
second
sight
town
large
tabby
shake
around
look
again
wasn
himself
thinking
watched
order
corner
jerked
could
that
something
cats
been
light
mind
drills
blinked
looking
signs
peculiar
seen
must
gave
hoping
didn
trick
road
little
thought
head
stared
mirror
read

## Example of extracted similar words
<pre>
packages   package
shape      shapes
shape      shaped
boil       boils
leapt      leap
objects    object
punishment punishments
stacked    stack
toilets    toilet
furiously  furious
tight      tighter
tight      tightly
drone      droned
brave      bravery
brave      braver
brave      bravely
crush      crushed
doughnuts  doughnut</pre>


The program also removed duplicate words, short words and proper nouns like Dursley, Grunnings, Privet Drive.
