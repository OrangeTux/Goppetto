function InvalidPinException(msg) {
  this.msg = msg;
  this.name = "InvalidPinException";
}

/**
 * Return the <input> element of a pin.
 */
function getPinSwitch(pin_id) {
  var input = $(':checkbox[data-id="' + pin_id + '"]').first();
  if (input.length === 0) {
    throw new InvalidPinException("Pin " + pin_id + " does not exists.");
  }

  return input;
}

/**
 * Set state of a pin switch, 0 and 1.
 */
function setPinState(pin_id, state) {
    var state = state === 1 ? true : false
    getPinSwitch(pin_id).bootstrapSwitch('state', state);
}

/**
 * Get state of a pin switch, either 0 or 1.
 */
function getPinState(pin_id) {
    var state = getPinSwitch(pin_id).bootstrapSwitch('state');
    return state ? 1 : 0
}
