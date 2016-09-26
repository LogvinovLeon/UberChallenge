"use strict";

var prependToMessageLog = function (payload, response) {
    var logEntry = document.createElement("div");
    logEntry.textContent = JSON.stringify(payload);
    document.getElementById("message_log").appendChild(logEntry);
};

var sendEmail = function () {
    var SEND_ENDPOINT = 'https://api.uberchallenge.email/email/';
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
        console.log("OK");
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
    console.log(placeholder.split("\\n"));
    document.getElementById("body").placeholder = placeholder.split("\\n").join("\n");
};
rewriteMultiLinePlaceholders();