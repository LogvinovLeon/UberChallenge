"use strict";

var renderMessage = function(message){
    var renderedMessage =
        "<div class='message'>" +
            "<h4 class='to'>" + message.to + "</h4>" +
            "<h5 class='subject'>" + message.subject + "</h5>" +
            "<p class='body'>" + message.body + "</p>" +
        "</div>";
    var wrapper = document.createElement('div');
    wrapper.innerHTML = renderedMessage;
    return wrapper;
};

var prependToMessageLog = function (payload, response) {
    document.getElementById("message_log").appendChild(renderMessage(payload));
};

var sendEmail = function () {
    var SEND_ENDPOINT = 'https://api.uberchallenge.email/email';
    var payload = {
        to: document.getElementById("to").value,
        subject: document.getElementById("subject").value,
        body: document.getElementById("body").value
    };
    fetch(SEND_ENDPOINT, {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify(payload)
    }).then(function (response) {
        prependToMessageLog(payload, response);
    }).catch(function () {
        console.log("FAIL");
        console.log(arguments);
    });
    return false
};

document.getElementById("send_email").onclick = sendEmail;

var rewriteMultiLinePlaceholders = function () {
    var placeholder = document.getElementById("body").placeholder;
    document.getElementById("body").placeholder = placeholder.split("\\n").join("\n");
};
rewriteMultiLinePlaceholders();
