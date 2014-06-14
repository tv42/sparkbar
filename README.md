# sparkbar -- Draw a sparkline in a terminal with UTF-8 block characters

This is a simple command line wrapper for
[github.com/joliv/spark](https://github.com/joliv/spark).

``` console
$ sparkbar 1 2 3 4 5 6 7 8
▁▂▃▄▅▆▇█

$ sparkbar .270 .272 .293 .310 .274 .239 .237 .238 .111
▇▇██▇▆▆▆▁

# relative memory sizes of 20 largest processes
$ ps -eo rss --sort rss|tail -20|sparkbar
▁▁▁▁▁▁▁▁▁▁▁▁▁▂▂▂▂▆██
```
