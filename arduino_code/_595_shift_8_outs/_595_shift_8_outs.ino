#include <SPI.h>


#include <Servo.h>

#include <rotary.h>
int csPin = 10; //You can use any IO pin but for this example we use 10

int cycles = 0;


// Clock Pin  to SH_CP pin 11 of 74HC595
const int clockPin = 5;

// Latch Pin to ST_CP pin 12 of 74HC595
const int latchPin = 6; 

// Data Pin to DS pin 14  of 74HC595
const int dataPin = 7; 

// The array for 16 outputs
byte pinz[] = {B00000000, B00000000};


// delay
int d = 100;

// the string from serial
String in_str = "";

//const int ENC_A = 2;
//const int ENC_B = 3;


Rotary rot = Rotary(2, 3);

int servo_pos =0;


ISR(PCINT2_vect) {
  unsigned char result = rot.process();
  if (result) {
      if( result == DIR_CW ){
        Serial.print("enc=1\n");
        servo_pos = servo_pos + 1;
      }else{
        Serial.print("enc=-1\n");
                servo_pos = servo_pos - 1;
      }
  }
}


void setup() {
    // set pins to output so you can control the shift register
    pinMode(latchPin, OUTPUT);
    pinMode(clockPin, OUTPUT);
    pinMode(dataPin, OUTPUT);


 
   

    pinMode(csPin, OUTPUT);
    digitalWrite(csPin, HIGH); //By default, don't be selecting OpenSegment
    SPI.begin(); //Start the SPI hardware
    SPI.setClockDivider(SPI_CLOCK_DIV64); //Slow down the master a bit

    //Send the reset command to the display - this forces the cursor to return to the beginning of the display
    digitalWrite(csPin, LOW); //Drive the CS pin low to select OpenSegment
    SPI.transfer('v'); //Reset command
  
    Serial.begin(9600);
  
    PCICR |= (1 << PCIE2);
    PCMSK2 |= (1 << PCINT18) | (1 << PCINT19);
    sei();
   
}

void spiSendValue(int tempCycles)
{
  digitalWrite(csPin, LOW); //Drive the CS pin low to select OpenSegment

  SPI.transfer(tempCycles / 1000); //Send the left most digit
  tempCycles %= 1000; //Now remove the left most digit from the number we want to display
  SPI.transfer(tempCycles / 100);
  tempCycles %= 100;
  SPI.transfer(tempCycles / 10);
  tempCycles %= 10;
  SPI.transfer(tempCycles); //Send the right most digit

  digitalWrite(csPin, HIGH); //Release the CS pin to de-select OpenSegment
}

void loop() {
    //Serial.println(servo_pos);
    //int msv =  map(servo_pos, 0, 180, 1000, 2000);
    //myservo.write( servo_pos );
      spiSendValue(servo_pos); 
    /*
    n = digitalRead(ENC_A);
     if( (encoder0PinALast == LOW) && (n == HIGH) ) {
     if( digitalRead(ENC_B) == LOW ) {
     Serial.print("enc=+1\n");
     if( x > 0) {
     x--;
     
     
     }
     } else {
     Serial.print("enc=-1\n");
     if( x < 15 ){
     x++;
     
     }
     }
     //Serial.print (x);
     //Serial.print ("\n");
     } 
     encoder0PinALast = n;
     */

    if (Serial.available() > 0) {
        int c = Serial.read();
        if( c == '=' || isDigit(c) ){
            in_str += (char)c;
        }
        if( c == '\n'){            
            int eq_pos = in_str.indexOf('=');

            String pin_str = in_str.substring(0, eq_pos);
            int pin = pin_str.toInt();

            String state_str = in_str.substring(eq_pos + 1);
            int xstate = state_str.toInt();

            setPin(pin, xstate);
            in_str = "";
        }
    }


    //setPin(0, true);
    //setPin(8, true);
    //setPin(9, true);
    // setVal(x);
    update();
    togglePin(4);
    togglePin(5);
    togglePin(8);
    togglePin(9);
    delay(d);
    //walkUp();
    //sweepUp();
    //  sweepUp();
    //sweepDown();
    //Serial.println("-");
    //sweepDown();
    //Serial.println("-----");
}

void setVal(int v){
    int i;
    for(i = 0; i < 16; i = i + 1){

        setPin(i, i <= v ? HIGH: LOW);
        //ssbitWrite(pinStates, npin, !bitRead(pinStates, npin));
        //pinStates =  pinStates << 1;
        //Serial.println(i);
        //update();
        //delay(d);
    }
}

void sweepUp(){
    int i;
    for(i = 0; i < 16; i = i + 1){

        togglePin(i);
        //ssbitWrite(pinStates, npin, !bitRead(pinStates, npin));
        //pinStates =  pinStates << 1;
        Serial.println(i);
        update();
        delay(d);
    }
}
void walkUp(){
    int x = 0;
    int i;
    for(i = 0; i < 15; i = i + 1){
        //togglePin(i);
        setPin(i, x == i );
        //Serial.println(i);
        //pinStates =  pinStates << 1;
        x += 1;
        update();
        delay(d);
    }
}


void sweepDown(){
    int i;
    for(i = 15; i != -1; i = i - 1){
        togglePin(i);
        //Serial.println(i);
        //pinStates =  pinStates << 1;
        update();
        delay(d);
    }
}
// Toggles state of npin
void update() {

    // turn off the output so the pins don't light up while you're shifting bits:
    digitalWrite(latchPin, LOW);

    // shift the bits out:
    shiftOut(dataPin, clockPin, MSBFIRST, pinz[0]);
    shiftOut(dataPin, clockPin, MSBFIRST, pinz[1]);

    // turn on the output so the LEDs can light up:
    digitalWrite(latchPin, HIGH);

}

// Toggles state of npin
void togglePin(int xpin) {

    int idx = xpin > 7 ? 1 : 0;
    int npin = xpin > 7 ? xpin - 8 : xpin;
    // toogle
    bitWrite(pinz[idx], npin, !bitRead(pinz[idx], npin));
}
// Sets a pin to state
void setPin(int xpin, int state) {

    int idx = xpin > 7 ? 1 : 0;
    int npin = xpin > 7 ? xpin - 8 : xpin;
    bitWrite(pinz[idx], npin, state);
}


