# Checkout Demo

### Notes
* I haven't implemented a main entrypoint for the task it just has tools that
  can be used. 
* Spent 2 hours on it so far and wasted 20-30 mins refactoring to use price
  rather than decimal, which was probably the wrong decision.
* I'll make a branch off main where I'll continue with the task adding CLI
  interaction because I'm invested in it now and enjoying myself.

### Dev Notes
* Made a cache so I could have prices saved per transaction so they don't
  change while in a sale.
* Started adding the change to transaction/tally but had a thought that my
  current price data structure isn't very useful as it requires checking
  two places rather than one.
* Started refactoring the Tally test to just count items and multibuys so
  only transaction cared about prices an realised tally could just be a
  function that does the counting rather than a separate pkg.
* Realised that having a tally for each item that inluded a count and a
  running total meant that when a new item is added the summation would
  only need done with the new item. Therefore reinstated tally but in
  transaction as it isn't a unique pkg rather a helper datatype.