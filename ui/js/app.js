const submitForm = () => {
    const textInput = document.getElementById('text-input').value;
    document.getElementById('result-text').innerHTML = '<div class="loader"></div>';
    fetch('/api', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify({text: textInput})
    })
        .then(response => response.json())
        .then(data => {
            document.getElementById('result-text').innerHTML = data.message;
            console.log(data);
        })
        .catch(err => {
            console.log(err);
        });
}

document.addEventListener('DOMContentLoaded', () => {
    const input = document.getElementById('text-input');
    input.addEventListener("keypress", function (event) {
        if (event.shiftKey && event.key === "Enter") {
            event.preventDefault();
            submitForm();
        }
    });
});
