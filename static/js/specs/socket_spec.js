describe('An EventDispatcher', function() {
    var socket;
    var ed;
    
    beforeEach(function() {
        //socket = jasmine.createSpyObj('socket', ['onmessage', 'send']);
        socket = jasmine.createSpy('socket');
        ed = new EventDispatcher(socket);
    });

    it('should be able to bind callbacks to events.', function() {
        function x() {};
        function y() {};
        ed.bind('some_event', x);
        ed.bind('some_event', y);

        expect(ed.callbacks['some_event']).toEqual([x, y]);
    });

    xit('should be able to send events to server.', function() {
        ed.send("pin_state", {pin_id: 3, state: 0})
        expect(socket.send).toHaveBeenCalledWith(JSON.stringify({
            'name': 'pin_state',
            'data': {
                'pin_id': 3,
                'state': 0
            }
        }));
    });

    it('should dispatch events when it receives them.', function() {
        var callback = jasmine.createSpy('callback');
        // Mock of MessageEvent aka the message send over the socket.
        var event = {'data': 
            JSON.stringify({
                'event': 'pin_state',
                'data': {
                    'pin_id': 3,
                    'state': 0
                }
            })
        };

        ed.bind('pin_state', callback);
        socket.onmessage(event);
        expect(callback).toHaveBeenCalledWith({'pin_id': 3, 'state': 0});
    });
});
