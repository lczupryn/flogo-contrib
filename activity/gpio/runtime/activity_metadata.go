package salesforce

var jsonMetadata = `{
  "name": "tibco-gpio",
  "version": "0.0.1",
  "title": "tibco-gpio",
  "description": "Control raspberry gpio",
  "homepage": "https://github.com/TIBCOSoftware/flogo-contrib/tree/master/activity/gpio",

  "inputs":[
    {
      "name": "method",
      "type": "string",
      "required": true,
      "allowed" : ["Direction", "Set State", "Read State", "Pull"]
    },
    {
      "name": "pinNumber",
      "type": "int",
      "required": true
    },
    {
      "name": "direction",
      "type": "string",
      "allowed" : ["Input", "Output"]
    },
    {
      "name": "state",
      "type": "string",
      "allowed" : ["High", "Low"]
    },

    {
      "name": "Pull",
      "type": "string",
      "allowed" : ["Up", "Down", "Off"]
    }
  ],
  "outputs": [
    {
      "name": "result",
      "type": "integer"
    }
  ]
}

`