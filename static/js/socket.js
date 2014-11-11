var EventDispatcher = function(socket) {
    var dit = this;

    this.socket = socket;

    this.callbacks = {};

    this.bind = function(event_name, callback) {
        this.callbacks[event_name] = this.callbacks[event_name] || [];
        this.callbacks[event_name].push(callback);
        return this;
    };

    this.send = function(event_name, event_data) {
        var payload = JSON.stringify({name: event_name, data: event_data});
        socket.send(payload); 
        return this;
    };

    var dispatch = function(event_name, event_data) {
        var callbacks = dit.callbacks[event_name];
        if (typeof callbacks === 'undefined') {
            return;
        }

        for (var i = 0; i < callbacks.length; i++) {
            callbacks[i](event_data);
        }
    };

    socket.onmessage = function(event){
        var json = JSON.parse(event.data);
        dispatch(json.event, json.data);
    };
};
