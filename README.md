Foresee how many days are left to accomplish a mission in a week with 7 days.

Parameters:
- `work`: how many days to work per week.
- `left`: how many items are left to be accomplished (e.g. questions in a book).
- `velocity`: how many items per day will be done.

Example:

```
go run .

262 5 2
work 1 1 260
work 2 2 258
work 3 3 256
work 4 4 254
work 5 5 252
rest 6 1 252
rest 7 2 252
work 8 1 250
work 9 2 248
work 10 3 246
work 11 4 244
...
work 173 5 12
rest 174 1 12
rest 175 2 12
work 176 1 10
work 177 2 8
work 178 3 6
work 179 4 4
work 180 5 2
rest 181 1 2
rest 182 2 2
work 183 1 0
2026-05-26 --> 183 --> 2026-11-25 (ETA)
```