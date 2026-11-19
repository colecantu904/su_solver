# su_solver

### TODO

#### Sudoku Solver
- Determine a good approach

##### Plan
Here is my initial instinct. We first start by going open space by open space. With each open space, we basically determine every possible value that can be in that space depending on the rows the columns and the box. Any square that has only 1 possible value, is just set to that value. Then you go to the next open square and do the same. Keep looping over all of the squares until solved. I think there is however some sort of standstill with this method. But honestly, lets try this method and see if it works.

### Large Scale Goals
- for puzzles we implement, lets have solver and generator functionality/commands for each! you can solve a puzzle you input, or generate a puzzle just for fun, then even check your solution.
- it would be cool, to make it a cli application, I also have some vision to make some way to expose it for public connectivity and use. That would be cool too.
- For starters, working sudoko solver/checker, with cli tool input and ouput, and if possible, even some http exposure for some web apps to send the data to!
- In addition, a solver/solution generator for kanoodle like game!
