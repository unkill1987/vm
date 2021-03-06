//SPDX-License-Identifier: Apache-2.0
'use strict'
// nodejs server setup 

// call the packages we need
var express       = require('express');        // call express
var app           = express();                 // define our app using express
var bodyParser    = require('body-parser');
var http          = require('http')
var fs            = require('fs');
var Fabric_Client = require('fabric-client');
var path          = require('path');
var util          = require('util');
var os            = require('os');
var ipfilter      = require('express-ipfilter').IpFilter;
var IpDeniedError = require('express-ipfilter').IpDeniedError;
// Load all of our middleware
// configure app to use bodyParser()
// this will let us get the data from a POST
app.use(bodyParser.urlencoded({ extended: true }));
app.use(bodyParser.json());

var ips = ['222.239.231.247','210.107.78.158','210.107.78.156'];
app.use(ipfilter(ips, {mode:'allow'}));
app.use(function(err,req,res,_next){
    res.send('Access Denied');
    if (err instanceof IpDeniedError){
        res.status(401).end();
    }else{
        res.status(err.status || 500).end();
    }
});
// instantiate the app
// this line requires and runs the code from our routes.js file and passes it app
require('./routes.js')(app);

// Save our port
var port = process.env.PORT || 8001;

// Start the server and listen on port 
app.listen(port,function(){
  console.log("Live on port: " + port);
});
