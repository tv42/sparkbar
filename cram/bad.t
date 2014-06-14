  $ sparkbar bad
  sparkbar: cannot parse number: strconv.ParseFloat: parsing "bad": invalid syntax
  [1]

  $ echo bad | sparkbar
  sparkbar: cannot parse number: strconv.ParseFloat: parsing "bad": invalid syntax
  [1]

  $ sparkbar </dev/null
  sparkbar: no data in stdin
  [1]
