
//bubble sort
PROGRAM bubbleSort;

CONST
  n=100;
VAR
  stala: array[1..100] of Integer;
  i,zm,temp: Integer; //zm - flag


PROCEDURE sortuj;
 BEGIN
  zm:=0;
  While zm=0 Do
    Begin
       zm:=1; //no more swaps - end
       For i:=1 To n-1 Do
          Begin
              if stala[i] > stala[i+1] then
                  Begin
                     zm:=0;
                     temp := stala[i]; //temporary variable to swap
                     stala[i] := stala[i+1];
                     stala[i+1] := temp;

                   End;

          End;

      End;
   END;





BEGIN
  
  For i:=1 To n Do
  Begin
  
   stala[i] := Random(n)+1;
   write(' | ', stala[i] );
  End;
   sortuj();
   writeln();
   writeln(' Posortowane! ');
   
   for i:= 1 to n Do
     Begin
       write(' | ', stala[i] );
     End;
   
 END.
