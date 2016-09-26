document.getElementById("send_email").onclick = function() {
    fetch('https://api.uberchallenge.email/email/', {
        method: 'POST',
        headers: {
            'Accept': 'application/json',
            'Content-Type': 'application/json'
        },
        body: JSON.stringify({
            to: 'logvinov.leon@gmail.com',
            subject: 'First email from web interface',
            body: 'Lorem ipsum'
        })
    })
    return false
}