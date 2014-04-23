var ws;

function pocWebSocketController($scope, $http) {
	//$scope.msgs = [];

	$scope.currentImage = 0;
	$scope.availableImages = [
	{
        src: "img/xigKMAejT.png"
	},
	{
        src: "img/bTy5ayeTL.png"
	}
	];


	$scope.Room= 'CRM';
	$scope.Message= 'Hello';

    function createSocket(roomName) {
        var ws = new WebSocket("ws://127.0.0.1:12345/echo");
        ws.onopen = function(){
            console.log("CONNECTION opened..." + ws.readyState);
            ws.send(JSON.stringify({event: "ADD", roomName: $scope.Room}));
        };

        ws.onmessage = function(message) {
            console.log("MESSAGE RECEIVE: " + message.data);
            //$scope.Msgs.push(message.data);
            $scope.Msgs = message.data;
            $scope.currentImage = 1;
            $scope.$apply();
        };

        ws.onclose  = function(m) {
            console.log("CONNECTION CLOSE");
            setTimeout(function() {
                createSocket();
            }, 1000);
        }
        $scope.ws = ws;
    }

    createSocket();

	$scope.Connect= function() {
        console.log("RECONNECTING...");
        createSocket();
	}

	$scope.Hungup= function() {
        console.log("Hungup");
		$scope.currentImage = 0;
		$scope.Msgs = "";
	}

	/*$scope.submitMessage= function() {
      console.log("Websocket - status: " + ws.readyState);
      if (ws.readyState == 1) {
      ws.send(JSON.stringify({ event: "ECHO", roomName: $scope.Room, message: $scope.Message }));
      }

	}*/
}
