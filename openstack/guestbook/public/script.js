$(document).ready(function() {
  var headerTitleElement = $("#header h1");
  var entriesElement = $("#guestbook-entries");
  var formElement = $("#guestbook-form");
  var submitElement = $("#guestbook-submit");
  var entryContentElement = $("#guestbook-entry-content");
  var hostAddressElement = $("#guestbook-host-address");
  var errorElement = $("#error");

  var appendGuestbookEntries = function(data) {
    entriesElement.empty();
    errorElement.empty();
    $.each(data, function(key, val) {
      entriesElement.append("<p>" + val + "</p>");
    });
  }

  var showError = function(value) {
    console.log("show error");
    errorElement.empty();
    errorElement.append("<p>" + value +" already exists </p>");
  }

  var handleSubmission = function(e) {
    e.preventDefault();
    var entryValue = entryContentElement.val()
    if (entryValue.length > 0) {
      $.getJSON("push/" + entryValue, function(data, status, xhr) {
          appendGuestbookEntries(data);
      }).fail(function() {
          showError(entryValue);
      });
    }
    return false;
  }

  // colors = purple, blue, red, green, yellow
  var colors = ["#549", "#18d", "#2a4", "#db1"];
  var randomColor = colors[Math.floor(5 * Math.random())];
  (function setElementsColor(color) {
    headerTitleElement.css("color", color);
    entryContentElement.css("box-shadow", "inset 0 0 0 2px " + color);
    submitElement.css("background-color", color);
  })(randomColor);

  submitElement.click(handleSubmission);
  formElement.submit(handleSubmission);
  hostAddressElement.append(document.URL);

  (function fetchGuestbook() {
    $.getJSON("list").done(appendGuestbookEntries);
  })();
});
