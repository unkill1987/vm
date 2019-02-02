//SPDX-License-Identifier: Apache-2.0
 
var trade = require('./controller.js');
 
module.exports = function(app){

  app.post('/keyHistory/:id', function(req, res){
    trade.keyHistory(req, res);
  });
  app.post('/get_trade/:id', function(req, res){
    trade.querytrade(req, res);
  });
  app.post('/add_OS/:record_os', function(req, res){
    trade.recordOS(req, res);
  });
  app.post('/add_LCR/:record_lcr', function(req, res){
    trade.recordLCR(req, res);
  }); 
  app.post('/add_LC/:record_lc', function(req, res){
    trade.recordLC(req, res);
  });
  app.post('/add_SR/:record_sr', function(req, res){
    trade.recordSR(req, res);
  });
  app.post('/add_BL/:record_bl', function(req, res){
    trade.recordBL(req, res);
  });
  app.post('/add_CI/:record_ci', function(req, res){
    trade.recordCI(req, res);
  });
  app.post('/add_DO/:record_do', function(req, res){
    trade.recordDO(req, res);
  });
}
