Current benchmarks (as of 7/23/2018)
```
goos: darwin
goarch: amd64
pkg: github.com/rictorlome/chess
BenchmarkParseFen-4                     	  200000	      8282 ns/op
BenchmarkGenerateFen-4                  	  500000	      3722 ns/op
BenchmarkGetColoredPieces-4             	2000000000	         0.35 ns/op
BenchmarkGetColoredKing-4               	2000000000	         0.36 ns/op
BenchmarkHasColoredPieceThere-4         	50000000	        25.8 ns/op
BenchmarkToPos-4                        	100000000	        13.2 ns/op
BenchmarkToPiece-4                      	100000000	        21.7 ns/op
BenchmarkFindPiece-4                    	30000000	        57.6 ns/op
BenchmarkParseMove-4                    	50000000	        26.1 ns/op
BenchmarkNaiveMove-4                    	10000000	       177 ns/op
BenchmarkGetAttackingSquaresInitial-4   	 2000000	       713 ns/op
BenchmarkGetAttackingSquaresMiddle-4    	 2000000	       911 ns/op
BenchmarkGetPseudoLegalMovesInitial-4   	 1000000	      1307 ns/op
BenchmarkGetPseudoLegalMovesMiddle-4    	 1000000	      1625 ns/op
BenchmarkGetLegalMovesInitial-4         	  100000	     17009 ns/op
BenchmarkGetLegalMovesMiddle-4          	   30000	     44893 ns/op
BenchmarkGetAllLegalMovesInitial-4      	    5000	    264244 ns/op
BenchmarkGetAllLegalMovesMiddle-4       	    2000	    730208 ns/op
PASS
ok  	github.com/rictorlome/chess	29.545s
```
