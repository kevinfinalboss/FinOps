console.log("Oi")
document.addEventListener('DOMContentLoaded', function() {
    const loginButton = document.querySelector('.btn-login');
    loginButton.addEventListener('click', function() {
        const email = document.querySelector('input[name="email"]').value;
        const password = document.querySelector('input[name="senha"]').value;

        fetch('/api/v1/user/login', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({ email, password }),
        })
        .then(response => response.json())
        .then(data => {
            if (data.token) {
                localStorage.setItem('token', data.token);
                window.location.href = '/entradas.html';
            } else {
                alert('Login falhou: ' + data.error);
            }
        })
        .catch((error) => {
            console.error('Error:', error);
        });
    });
});
