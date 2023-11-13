document.addEventListener('DOMContentLoaded', function() {
    const loginButton = document.querySelector('.btn-login');
    loginButton.addEventListener('click', function() {
        const email = document.querySelector('input[name="email"]').value;
        const password = document.querySelector('input[name="password"]').value;

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
                document.cookie = `token=${data.token};path=/;max-age=1800`;
                window.location.href = '/saidas';
            } else {
                alert('Login falhou: ' + data.error);
            }
        })
        .catch((error) => {
            console.error('Error:', error);
        });
    });
});
