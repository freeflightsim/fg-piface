
int aPin = 2;  //                     A
int bPin = 3;  //             ________
int cPin = 4;  //           |                   |
int dPin = 5;  //       F  |                   |  B
int ePin = 6;  //           |         G       |
int fPin = 7;  //            |________|
int gPin = 8;  //           |                   |
int GND1 = 9;  //        |                   |
int GND2 = 10; //   E |                   |   C
int GND3 = 11; //       |________|
int GND4 = 12; //       
int num = 0;  //         D
int val = 0;
int dig1 = 0;
int dig2 = 0;
int dig3 = 0;
int dig4 = 0;
int DTime = 4;

bool on = 0;
bool off = 1;


void setup()
{
  pinMode(aPin, OUTPUT);
  pinMode(bPin, OUTPUT);
  pinMode(cPin, OUTPUT);
  pinMode(dPin, OUTPUT);
  pinMode(ePin, OUTPUT); 
  pinMode(fPin, OUTPUT);
  pinMode(gPin, OUTPUT);
  pinMode(GND1, OUTPUT);
  pinMode(GND2, OUTPUT);
  pinMode(GND3, OUTPUT);
  pinMode(GND4, OUTPUT);
  Serial.begin(9600);
}
void loop()
{
  digitalWrite( GND1, HIGH);
  digitalWrite( GND2, HIGH);
  digitalWrite( GND3, HIGH);
  digitalWrite( GND4, HIGH);


  delay(400);
  if ( val == 9 ) {
    val = 0;
    num = 2;
  } else {
    val = val + 1;
    num = 3;
  }
  //num = 0;
  //val = 1;
  //num = val;
  //num = random(0, 9999);
  Serial.print(num);
  dig1 = val;// val; //num / 1000;
  //num = num - (dig1 * 1000);
  dig2 = val; //num / 100;
  //num = num - (dig2 * 100);
  dig3 = num; //num / 10;
  dig4 = num; //num - (dig3 *10);
  
  Serial.print(" ");
  Serial.print(dig1);
  Serial.print(" ");
  Serial.print(dig2);
  Serial.print(" ");
  Serial.println(dig3);
 
 
  digitalWrite( GND4, LOW);    //digit 4
  pickNumber(dig4);
  delay(DTime);
  digitalWrite( GND4, HIGH);
 
  
  digitalWrite( GND3, LOW);    //digit 3
  pickNumber(dig3);
  delay(DTime);
  digitalWrite( GND3, HIGH);
 
  digitalWrite( GND2, LOW);   //digit 2
  pickNumber(dig2);
  delay(DTime);
  digitalWrite( GND2, HIGH);
 
  digitalWrite( GND1, LOW);   //digit 1
  pickNumber(dig1);
  delay(DTime);
  digitalWrite( GND1, HIGH);

}
 
void pickNumber(int x){
   switch(x){
     case 1: one(); break;
     case 2: two(); break;
     case 3: three(); break;
     case 4: four(); break;
     case 5: five(); break;
     case 6: six(); break;
     case 7: seven(); break;
     case 8: eight(); break;
     case 9: nine(); break;
     default: zero(); break;
   }
}

void clearLEDs()
{  
  digitalWrite(  aPin, LOW); // A
  digitalWrite(  bPin, LOW); // B
  digitalWrite(  cPin, LOW); // C
  digitalWrite(  dPin, LOW); // D
  digitalWrite(  ePin, LOW); // E
  digitalWrite(  fPin, LOW); // F
  digitalWrite(  gPin, LOW); // G
}
void zero()
{
  digitalWrite( aPin, on);
  digitalWrite( bPin, on);
  digitalWrite( cPin, on);
  digitalWrite( dPin, on);
  digitalWrite( ePin, on);
  digitalWrite( fPin, on);
  digitalWrite( gPin, off);
}
void one()
{
  digitalWrite( aPin, off);
  digitalWrite( bPin, on);
  digitalWrite( cPin, on);
  digitalWrite( dPin, off);
  digitalWrite( ePin, off);
  digitalWrite( fPin, off);
  digitalWrite( gPin, off);
}

void two()
{
  digitalWrite( aPin, on);
  digitalWrite( bPin, on);
  digitalWrite( cPin, off);
  digitalWrite( dPin, on);
  digitalWrite( ePin, on);
  digitalWrite( fPin, off);
  digitalWrite( gPin, on);
}

void three()
{
  digitalWrite( aPin, on);
  digitalWrite( bPin, on);
  digitalWrite( cPin, on);
  digitalWrite( dPin, on);
  digitalWrite( ePin, off);
  digitalWrite( fPin, off);
  digitalWrite( gPin, on);
}

void four()
{
  digitalWrite( aPin, off);
  digitalWrite( bPin, on);
  digitalWrite( cPin, on);
  digitalWrite( dPin, off);
  digitalWrite( ePin, off);
  digitalWrite( fPin, on);
  digitalWrite( gPin, on);
}

void five()
{
  digitalWrite( aPin, on);
  digitalWrite( bPin, off);
  digitalWrite( cPin, on);
  digitalWrite( dPin, on);
  digitalWrite( ePin, off);
  digitalWrite( fPin, on);
  digitalWrite( gPin, on);
}

void six()
{
  digitalWrite( aPin, off);
  digitalWrite( bPin, off);
  digitalWrite( cPin, on);
  digitalWrite( dPin, on);
  digitalWrite( ePin, on);
  digitalWrite( fPin, on);
  digitalWrite( gPin, on);
}

void seven()
{
  digitalWrite( aPin, on);
  digitalWrite( bPin, on);
  digitalWrite( cPin, on);
  digitalWrite( dPin, off);
  digitalWrite( ePin, off);
  digitalWrite( fPin, off);
  digitalWrite( gPin, off);
}

void eight()
{
  digitalWrite( aPin, on);
  digitalWrite( bPin, on);
  digitalWrite( cPin, on);
  digitalWrite( dPin, on);
  digitalWrite( ePin, on);
  digitalWrite( fPin, on);
  digitalWrite( gPin, on);
}

void nine()
{
  digitalWrite( aPin, on);
  digitalWrite( bPin, on);
  digitalWrite( cPin, on);
  digitalWrite( dPin, on);
  digitalWrite( ePin, off);
  digitalWrite( fPin, on);
  digitalWrite( gPin, on);
}


