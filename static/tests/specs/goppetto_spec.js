describe("A pin", function() {
  beforeEach(function() {
    affix('input[type="checkbox"][data-id="3"][checked="checked"]');
    affix('input[type="checkbox"][data-id="4"]');
  });

  it("should have a checkbox.", function() {
    expect(getPinSwitch(3).data('id')).toBe(3);
  });

  it("is invalid without checkbox.", function() {
    expect(function() {getPinSwitch(18);}).toThrow();
  });

  it("has a state which is either 0 or 1.", function() {
    affix('input[type="checkbox"][data-id="4"]');
    expect(getPinState(3)).toBe(1);
    expect(getPinState(4)).toBe(0);
  });

  it("can be set to a state which either 0 or 1.", function() {
    setPinState(3, 0);
    expect(getPinState(3)).toBe(0);

    setPinState(4, 1);
    expect(getPinState(4)).toBe(1);
  });
});
