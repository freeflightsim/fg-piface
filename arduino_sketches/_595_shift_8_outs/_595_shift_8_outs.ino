

// Pin connected to SH_CP pin 11 of 74HC595
int clockPin = 2;

// Pin connected to ST_CP pin 12 of 74HC595
int latchPin = 3; 

// Pin connected to DS pin 14  of 74HC595
int dataPin = 4; 


byte pinStates = B00000000;

int d = 200;

void setup() {
    // set pins to output so you can control the shift register
    pinMode(latchPin, OUTPUT);
    pinMode(clockPin, OUTPUT);
    pinMode(dataPin, OUTPUT);
    Serial.begin(9600);
}

void loop() {
  //if (Serial.available() > 0) {
    // ASCII '0' through '9' characters are
    // represented by the values 48 through 57.
    // so if the user types a number from 0 through 9 in ASCII, 
    // you can subtract 48 to get the actual value:
    //int bitToSet = Serial.read() - 48;
    //Serial.println(bitToSet);
  // write to the shift register with the correct bit set high:
    //registerWrite(bitToSet, HIGH);
  //}

  sweepUp();
//  sweepUp();
  //sweepDown();
  Serial.println("-");
  sweepDown();
  Serial.println("-----");
}

void sweepUp(){
  int i;
   for(i = 0; i < 8; i = i + 1){
    togglePin(i);
    //ssbitWrite(pinStates, npin, !bitRead(pinStates, npin));
    //pinStates =  pinStates << 1;
    Serial.println(i);
    update();
    delay(d);
  }
}
void sweepDown(){
  int i;
   for(i = 7; i != -1; i = i - 1){
    togglePin(i);
    Serial.println(i);
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
  shiftOut(dataPin, clockPin, MSBFIRST, pinStates);
  shiftOut(dataPin, clockPin, MSBFIRST, pinStates);
  
  // turn on the output so the LEDs can light up:
  digitalWrite(latchPin, HIGH);

}

// Toggles state of npin
void togglePin(int npin) {

  // turn off the output so the pins don't light up while you're shifting bits:
  //digitalWrite(latchPin, LOW);
  
  // toogle
  bitWrite(pinStates, npin, !bitRead(pinStates, npin));
  
  // shift the bits out:
  //shiftOut(dataPin, clockPin, MSBFIRST, pinStates);
  
  // turn on the output so the LEDs can light up:
  //digitalWrite(latchPin, HIGH);

}
