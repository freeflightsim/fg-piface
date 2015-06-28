
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

int DTime = 3;

bool on = 0;
bool off = 1;
int curr_num = 0;

long previousMillis = 0; 
long interval = 500;

String in_str = "";

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
  
  digitalWrite( GND1, LOW);
  digitalWrite( GND2, LOW);
  digitalWrite( GND3, LOW);
  digitalWrite( GND4, LOW);


  //delay(1000);
  if (Serial.available() > 0) {
    
    int c = Serial.read();  //gets one byte from serial buffer
    if( isDigit(c) ){
      in_str += (char)c;
    }
    if( c == '\n'){
     // Serial.print("----");
      //Serial.println(in_str);
      curr_num = in_str.toInt();
      in_str = "";
    }
    
  //unsigned long currentMillis = millis();
  //if(currentMillis - previousMillis > interval) {
    //previousMillis = currentMillis;
    //curr_num = random(0, 9999);
     //Serial.print(curr_num);
    

 
    num = curr_num;
    dig1 = num / 1000;
    num = num - (dig1 * 1000);
    dig2 = num / 100;
    num = num - (dig2 * 100);
    dig3 = num / 10;
    dig4 = num - (dig3 *10);
    /*
    Serial.print(" = ");
    Serial.print(dig1);
    Serial.print(" ");
    Serial.print(dig2);
    Serial.print(" ");
    Serial.print(dig3);
    Serial.print(" ");
    Serial.println(dig4);
    */
   }
   
   int pinz[] = {GND4, GND3, GND2, GND1};
   int vals[] = {dig4, dig3, dig2, dig1};
   int i;
   for (i = 0; i < 4; i = i + 1) {
         digitalWrite( pinz[i], off);    //digit 4
         pickNumber(vals[i]);
        /// Serial.print(i );
         ///Serial.print(pinz[i]);
         delay(DTime);
         digitalWrite( pinz[i], on);
   }
   
   /*
   
  digitalWrite( GND4, off);    //digit 4
  pickNumber(dig4);
  delay(DTime);
  digitalWrite( GND4, on);
 
  
  digitalWrite( GND3, off);    //digit 3
  pickNumber(dig3);
  delay(DTime);
  digitalWrite( GND3, on);
 
  digitalWrite( GND2, off);   //digit 2
  pickNumber(dig2);
  delay(DTime);
  digitalWrite( GND2, on);
 
  digitalWrite( GND1, off);   //digit 1
  pickNumber(dig1);
  delay(DTime);
  digitalWrite( GND1, on);
  */
  
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
  digitalWrite(  aPin, off); // A
  digitalWrite(  bPin, off); // B
  digitalWrite(  cPin, off); // C
  digitalWrite(  dPin, off); // D
  digitalWrite(  ePin, off); // E
  digitalWrite(  fPin, off); // F
  digitalWrite(  gPin, off); // G
}
void dash()
{  
  digitalWrite(  aPin, off); // A
  digitalWrite(  bPin, off); // B
  digitalWrite(  cPin, off); // C
  digitalWrite(  dPin, off); // D
  digitalWrite(  ePin, off); // E
  digitalWrite(  fPin, off); // F
  digitalWrite(  gPin, on); // G
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


