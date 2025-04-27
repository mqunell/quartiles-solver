![Quartiles example image](https://images.macrumors.com/article-new/2024/04/iOS-17.5-Quartiles-Feature.jpg)

# Apple News+ Quartiles Solver

This is a short Go script for solving the quartiles game. It generates a list of ~100,000 word guesses by looping through the word segments and making every possible combination. Based on the puzzle in the image above, the guesses would look something like:

- tro•dis•gas•in
- tro•dis•gas•ti
- tro•dis•gas•zan
- ...
- ra•ma•dan•cor
- ra•ma•dan•vit
- ra•ma•dan•ic

Somewhere in that list would be the solution words:

- dis•cor•dan•ce
- gas•tro•nom•ic
- un•in•vit•ed
- by•zan•ti•ne
- ma•ra•thon•er

Next, the script checks each guess against a map of all ~400,000 English words. If it finds a match, it prints the word to the console. _(It may find extra words not considered valid by the game, like below. I would need the official list of words to fix this.)_

```
⚡️ found discordance
⚡️ found gastronomer
⚡️ found gastronomic
⚡️ found uninvited
⚡️ found untiered
⚡️ found byzantine
⚡️ found marathoner
```

> Implementation note: Structuring the valid words as a map is an extremely important optimization step. Looking up the guesses this way compared to simply doing `slices.Contains(allWords, guess)` for each guess reduces runtime by 99.955%, from roughly 51 seconds to 23 *milli*seconds! This is because the number of operations is reduced from O(nm) to O(n+m); ~40 billion to ~500k. 🤯
