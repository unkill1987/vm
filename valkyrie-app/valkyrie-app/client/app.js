// SPDX-License-Identifier: Apache-2.0

'use strict';

var app = angular.module('application', []);

// Angular Controller
app.controller('appController', function($scope, appFactory){

	$("#success_CI").hide();
	$("#error_CI").hide();
	$("#success_LCR").hide();
	$("#error_LCR").hide();
	$("#success_LC").hide();
	$("#error_LC").hide();
	$("#success_SR").hide();
	$("#error_SR").hide();
	$("#success_BL").hide();
	$("#error_BL").hide();
	$("#success_DO").hide();
	$("#error_DO").hide();
	$("#error_query").hide();
	
	
	
	$scope.keyHistory = function(){
		var id = $scope.contract_id;
		
		appFactory.keyHistory(id, function(data){
		
			$scope.key_History = data;			
		});
	}
	
	$scope.queryTrade = function(){

		var id = $scope.contract_id;
		
		appFactory.queryTrade(id, function(data){
			$scope.query_trade = data;

			if ($scope.query_trade == "Could not found Contract_ID"){
				console.log()
				$("#error_query").show();
			} else{
				$("#error_query").hide();
			}
		});
	}

	// $scope.recordCI = function(){

	// 	appFactory.recordCI($scope.record_ci, function(data){
	// 		$scope.record_ci = data;
	// 		$("#success_CI").show();
	// 	});
	// }

	$scope.recordCI = function(){

		appFactory.recordCI($scope.record_ci, function(data){
			$scope.record_ci = data;
			if ($scope.record_ci == "The contract already exists"){
				$("#error_CI").show();
				$("#success_CI").hide();
			} else{
				$("#success_CI").show();
				$("#error_CI").hide();
			}
		});
	}

	$scope.recordLCR = function(){

		appFactory.recordLCR($scope.record_lcr, function(data){
			$scope.record_lcr = data;
			if ($scope.record_lcr == "Could not found Contract_ID"){
				$("#error_LCR").show();
				$("#success_LCR").hide();
			} else{
				$("#success_LCR").show();
				$("#error_LCR").hide();
			}
		});
	}

	$scope.recordLC = function(){

		appFactory.recordLC($scope.record_lc, function(data){
			$scope.record_lc = data;
			if ($scope.record_lc == "Could not found Contract_ID"){
				$("#error_LC").show();
				$("#success_LC").hide();
			} else{
				$("#success_LC").show();
				$("#error_LC").hide();
			} 
		});
	}

	$scope.recordSR = function(){

		appFactory.recordSR($scope.record_sr, function(data){
			$scope.record_sr = data;
			if ($scope.record_sr == "Could not found Contract_ID"){
				$("#error_SR").show();
				$("#success_SR").hide();
			} else{
				$("#success_SR").show();
				$("#error_SR").hide();
			}
		});
	}

	$scope.recordBL = function(){

		appFactory.recordBL($scope.record_bl, function(data){
			$scope.record_bl = data;
			if ($scope.record_bl == "Could not found Contract_ID"){
				$("#error_BL").show();
				$("#success_BL").hide();
			} else{
				$("#success_BL").show();
				$("#error_BL").hide();
			}
		});
	}

	$scope.recordDO = function(){

		appFactory.recordDO($scope.record_do, function(data){
			$scope.record_do = data;
			if ($scope.record_do == "Could not found Contract_ID"){
				$("#error_DO").show();
				$("#success_DO").hide();
			} else{
				$("#success_DO").show();
				$("#error_DO").hide();
			}
		});
	}

});

// Angular Factory
app.factory('appFactory', function($http){
	
	var factory = {};

	factory.keyHistory = function(id, callback){
    	$http.get('/keyHistory/'+id).success(function(output){
			callback(output)
		});
	}

   	factory.queryTrade = function(id, callback){
    	$http.get('/get_trade/'+id).success(function(output){
			callback(output)
		});
	}

	factory.recordCI = function(data, callback){

		var record_ci = data.id + "-" + data.exporter + "-" + data.ci_hash;

    	$http.get('/add_CI/'+record_ci).success(function(output){
			callback(output)
		});
	}

	factory.recordLCR = function(data, callback){

		var record_lcr = data.id + "-" + data.importer + "-" + data.lcr_hash;

    	$http.get('/add_LCR/'+record_lcr).success(function(output){
			callback(output)
		});
	}

	factory.recordLC = function(data, callback){

		var record_lc = data.id + "-" + data.bank + "-" + data.lc_hash;

    	$http.get('/add_LC/'+record_lc).success(function(output){
			callback(output)
		});
	}

	factory.recordSR = function(data, callback){

		var record_sr = data.id + "-" + data.exporter + "-" + data.sr_hash;

    	$http.get('/add_SR/'+record_sr).success(function(output){
			callback(output)
		});
	}

	factory.recordBL = function(data, callback){

		var record_bl = data.id + "-" + data.shipper + "-" + data.bl_hash;

    	$http.get('/add_BL/'+ record_bl).success(function(output){
			callback(output)
		});
	}

	factory.recordDO = function(data, callback){

		var record_do = data.id + "-" + data.shipper + "-" + data.do_hash;

    	$http.get('/add_DO/'+record_do).success(function(output){
			callback(output)
		});
	}

	return factory;
});


