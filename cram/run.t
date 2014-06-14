# TODO stupid cram doesn't understand utf-8
#
#   $ sparkbar 1 2 3 4 5 6 7 8
#   ▁▂▃▄▅▆▇█
#
#   $ sparkbar .270 .272 .293 .310 .274 .239 .237 .238 .111
#   ▇▇██▇▆▆▆▁
#
#   $ sparkbar 67 71 77 85 95 104 106 105 100 89 76 66
#   ▁▂▃▄▆███▇▅▃▁
#
#   $ sparkbar 42
#   ❄
#
#   $ sparkbar -- 0 -100 0
#   █▁█
#
#   $ printf '1\n2\n3\n4\n5\n6\n7\n8\n' | sparkbar
#   ▁▂▃▄▅▆▇█
#
#   # empty lines are ignored
#   $ printf '1\n\n2\n3\n4\n5\n6\n7\n8\n' | sparkbar
#   ▁▂▃▄▅▆▇█
#
#   # whitespace separates number for ignored columns
#   $ printf '1\n2 junk\n3\n4\tnoise\n5\n6\n7\n8\n' | sparkbar
#   ▁▂▃▄▅▆▇█

