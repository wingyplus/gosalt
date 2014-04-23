var MainCtrl = function($scope, $http) {
	$scope.loading = true;
  $http.get('http://dth2732042t01:8080/SaveSpkdAddWS', { 
		'USER_CODE': 'LLTHUNYADAP',
		'BLPD_INDC': 'PCN',
		'CS_SPKD_PCN__CUST_NUMB': '536672462',
		'CS_SPKD_PCN__SUBR_NUMB': '66900010040',
		'CS_SPKD_PCN__PACK_CODE': '31001501',
		'RD_TELP__TELP_TYPE': 'TEL',
		'SAVE_FLAG': '1',
	})
	.success(function(data, status, headers, config) {
		$scope.CS_SPKD_PCN__PACK_CODE = data.Body.CS_SPKD_PCN__PACK_CODE;
		$scope.CS_PKPL_PCN__PACK_DESC = data.Body.CS_PKPL_PCN__PACK_DESC;
		$scope.CS_PACK_TYPE__PACK_TYPE_DESC = data.Body.CS_PACK_TYPE__PACK_TYPE_DESC;
		$scope.CS_SPKD_PCN__PACK_STRT_DTTM = data.Body.CS_SPKD_PCN__PACK_STRT_DTTM;
		$scope.CS_SPKD_PCN__PACK_END_DTTM = data.Body.CS_SPKD_PCN__PACK_END_DTTM;
		$scope.CS_SPKD_PCN__DISC_CODE = data.Body.CS_SPKD_PCN__DISC_CODE;
		$scope.TBL_OCCR = data.Body.TBL_OCCR;
		scope.Status = data.Status;
		$scope.loading = false;
	})
	.error(function(err, status, headers, config) {
		console.log("Well, this is embarassing.");
		scope.Status = "Error ZZZZZ";
		$scope.loading = false;
	});
	
}