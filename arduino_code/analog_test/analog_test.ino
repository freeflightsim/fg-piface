
#include <Servo.h>




// Throttle Pot
const int analogInPin = 0;  
int led = 13;
int led2 = 12;
bool state = 0;
//Servo mservo;

int sensorValue = 0;        // value read from the pot
int outputValue = 0;        // value output as %
//int servoValue = 0;

void setup() {

  pinMode(led, OUTPUT);
  pinMode(led2, OUTPUT);
  // initialize serial communications at 9600 bps:
  Serial.begin(9600); 
  //mservo.attach(3);
}


void loop() {

  // read the analog in value:
  sensorValue = analogRead(analogInPin);     
  
  // map it to the range of the analog out:
  outputValue = map(sensorValue, 0, 1023, 0, 100);  
  //servoValue =  map(sensorValue, 0, 1023, 0, 170);  
  
  //## NOTE ## Not using prinln() == 10,13.. using \n = 10 instead    
  Serial.print(outputValue);   
  Serial.print("\n");   

  state = !state;
  digitalWrite(led, state);
  digitalWrite(led2, !state);
  delay(400);

// mservo.write(servoValue);
 //delay(10);                     
}
