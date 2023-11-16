document.addEventListener('DOMContentLoaded', function() {
    const form = document.getElementById('formUsuario');
    
    form.addEventListener('submit', function(event) {
        event.preventDefault();

        const formData = new FormData(form);
        const userData = Object.fromEntries(formData.entries());

        fetch('/api/v1/user/register', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(userData)
        }).then(response => {
            if (response.ok) {
                return response.json();
            } else {
                throw new Error('Erro ao criar usuário');
            }
        }).then(data => {
            console.log(data);
            alert('Usuário criado com sucesso!');
            form.reset();
        }).catch((error) => {
            console.error('Error:', error);
            alert('Erro ao criar usuário.');
        });
    });
});
