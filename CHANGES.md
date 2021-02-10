# Changes

available at: https://shakespearch.herokuapp.com/

## Initial ideas

1. better search backend
2. highlight search result
3. link to full context of search result
4. run data through sentiment analysis
5. incremental search

## Implemented

Due to 1 and 4, the search backend should support structure data:
selected [Algolia](https://www.algolia.com/) for ease of use and included highlights for 2.

Algolia requires structured data:
easy way out is to search for preparsed data, found in [Elastic Kibana tutorial](https://www.elastic.co/guide/en/kibana/6.8/tutorial-load-dataset.html).

- TODO: This might not include the sonnets
- Alternative source (split plain text): https://github.com/okfn/shakespeare-material/tree/master/texts/gutenberg

Basic styling with [milligram](https://milligram.io/):
I've used it (a long time ago), doesn't require too much work to look half decent.

## Not implemented

in order of priority

**Serve full text and link from search result:**
Data in [static/texts](./static/texts/) but needs restructuring / crosslinking with the indexed data,
or reindex using the plain text.

**Better UI:**
too long since I touched frontend code, need more time (than 2h) to ramp back up.

**Hint user about possible search operators:**
Algolia doesn't have easily findable docs on it.

**Incremental search:**
more frontend work

**Sentiment analysis:**
Run the data through [GCP Natural Language](https://cloud.google.com/natural-language), one time cost (Shakespeare isn't going to revive and start writing more stuff). Not done due to:

- lack of time
- no effective way of presenting the data (need frontend work)

maybe making google crawl the site would have been faster...
