package sudoku_solver 

//import "fmt"

func (s *Sudoku) solveByNakedAndLockedSubsets(similarCellCount int) bool {
   var markersChanged bool

   var markerCount               [9][9]uint8
   var blockSolvedCellCountTable [9][9]uint8
   var cellMarkers                  [9]bool

   // in int belows I'll be storing number of found cells that have the same cellMarkers
   var similarRowCellCounter, similarColumnCellCounter, blockSimilarCellCounter int
   var rowSolvedCellCounter,  columnSolvedCellCounter   [9]int

   for a := range s.solution {
      a_min := (a/3)*3     // minimal index of a block row
      a_max := a_min + 2   // maximal index of a block row

      for b := range s.solution[a] {
         b_min := (b/3)*3   // minimal index of a block column
         b_max := b_min + 2 // maximal index of a block column

         if s.solution[a][b] == 0 {
            for c := range s.markerTable[a][b] {
               if s.markerTable[a][b][c] {
                  markerCount[a][b]++
               }
            }
         } else {
            rowSolvedCellCounter[a]++
            columnSolvedCellCounter[b]++
            for a2 := a_min; a2 <= a_max; a2++ {
               for b2 := b_min; b2 <= b_max; b2++ {
                  blockSolvedCellCountTable[a2][b2]++
               }
            }
         }
      }
   }

/*
   fmt.Println("markerCount table\n")
   print9x9(markerCount)
   fmt.Println("blockSolvedCellCountTable table\n")
   print9x9(blockSolvedCellCountTable)
   print9x9x9(s.solution, s.markerTable)
   */ 

   for a := range s.solution {
      a_min := (a/3)*3     // minimal index of a block row
      a_max := a_min + 2   // maximal index of a block row

      for b := range s.solution[a] {
         b_min := (b/3)*3     // minimal index of a block row
         b_max := b_min + 2   // maximal index of a block row

         if s.solution[a][b] == 0 {
            if markerCount[a][b] == uint8(similarCellCount) {
               cellMarkers = s.markerTable[a][b]

               //fmt.Printf("(a,b)=(%d,%d), cellMarkers = %v\n", a+1, b+1, cellMarkers)

               // finding similar cells in blocks ->
               // >>>>>>>>>>>>>>> Locked Subsets <<<<<<<<<<<<<<<
               if blockSolvedCellCountTable[a][b] != uint8(9 - similarCellCount) {
                  for a2 := a_min; a2 <= a_max; a2++ {
                     for b2 := b_min; b2 <= b_max; b2++ {

                        if (a != a2 || b != b2) && s.markerTable[a2][b2] == cellMarkers {
                           blockSimilarCellCounter++
                        }

                        if blockSimilarCellCounter == similarCellCount - 1 {
                           blockSimilarCellCounter = 0
                           /*
                           fmt.Printf("block:  cell (%d:%d) is the same like  (%d:%d)\n", a+1, b+1, a2+1, b2+1)
                           fmt.Println(cellMarkers)
                           fmt.Println(s.markerTable[a2][b2]) 
                           */
                           for a3 := a_min; a3 <= a_max; a3++ {
                              for b3 := b_min; b3 <= b_max; b3++ {
                                 if s.markerTable[a3][b3] != cellMarkers {
                                    for marker := range s.markerTable[a3][b3] {
                                       if cellMarkers[marker] && s.markerTable[a3][b3][marker] {
                                          s.markerTable[a3][b3][marker] = false
                                          markersChanged = true
                                       }
                                    }
                                 }
                              }
                           }
                           //print9x9x9(s.solution, s.markerTable)
                        }
                     }
                  }
               }

               if markersChanged {s.solveBasingOnMarkers()}
 
               // finding similar cells in rows/columns -> 
               // >>>>>>>>>>>>>>> Naked Subsets <<<<<<<<<<<<<<<
               for c := range s.solution[a] {
                  if c != b && s.markerTable[a][c] == cellMarkers {
                     similarRowCellCounter++
                  }

                  if similarRowCellCounter == similarCellCount - 1  {
                     similarRowCellCounter = 0
                     if rowSolvedCellCounter[a] != 9 - similarCellCount {
                        /*
                        fmt.Printf("row:  cell (%d:%d) is the same like  (%d:%d)\n", a+1, b+1, a+1, c+1)
                        fmt.Println(cellMarkers)
                        fmt.Println(s.markerTable[a][c])
                        */
                        for col := range s.solution[a] {
                           if s.solution[a][col] == 0 && s.markerTable[a][col] != cellMarkers {
                              for numberMarker := range s.markerTable[a][col] {
                                 if cellMarkers[numberMarker] && s.markerTable[a][col][numberMarker] {
                                    s.markerTable[a][col][numberMarker] = false
                                    markersChanged = true   
                                 }
                              }
                           }
                        }
                        //print9x9x9(s.solution, s.markerTable)
                     }
                  }
                  
                  if markersChanged {s.solveBasingOnMarkers()}

                  if c != a && s.markerTable[c][b] == cellMarkers {
                     similarColumnCellCounter++
                  }

                  if similarColumnCellCounter == similarCellCount - 1  {
                     similarColumnCellCounter = 0
                     if columnSolvedCellCounter[b] != 9 - similarCellCount {
                        /*
                        fmt.Printf("column: cell (%d:%d) is the same like (%d:%d)\n", a+1, b+1, c+1, b+1)
                        fmt.Println(cellMarkers)
                        fmt.Println(s.markerTable[c][b])
                        */
                        for row := range s.solution[a] {
                           if s.solution[row][b] == 0 && s.markerTable[row][b] != cellMarkers {
                              for numberMarker := range s.markerTable[row][b] {
                                 if cellMarkers[numberMarker] && s.markerTable[row][b][numberMarker] {
                                    markersChanged = true
                                    s.markerTable[row][b][numberMarker] = false
                                 }                                 
                              }
                           }
                        }
                        //print9x9x9(s.solution, s.markerTable)
                     }
                  }
                  if markersChanged {s.solveBasingOnMarkers()}
               }
            }
         }
      }
   }
  return markersChanged
}