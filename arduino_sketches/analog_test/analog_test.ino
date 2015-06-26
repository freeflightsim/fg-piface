


// Throttle Pot
const int analogInPin = 0;  


int sensorValue = 0;        // value read from the pot
int outputValue = 0;        // value output as %


void setup() {
  // initialize serial communications at 9600 bps:
  Serial.begin(9600); 
}


void loop() {

  // read the analog in value:
  sensorValue = analogRead(analogInPin);     
  
  // map it to the range of the analog out:
  outputValue = map(sensorValue, 0, 1023, 0, 100);  
  
  
  //## NOTE ## Not using prinln() == 10,13.. using \n = 10 instead    
  Serial.print(outputValue);   
  Serial.print("\n");   

  delay(100);                     
}
