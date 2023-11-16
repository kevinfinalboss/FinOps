document.addEventListener('DOMContentLoaded', function() {
    fetchUsers();
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

function fetchUsers() {
    fetch('/api/v1/users')
        .then(response => response.json())
        .then(data => {
            if (data && data.users) {
                renderUsers(data.users);
            }
        })
        .catch(error => {
            console.error('Erro ao buscar usuários:', error);
        });
}

function renderUsers(users) {
    const usersTable = document.getElementById('usersTable');
    users.forEach(user => {
        const tr = document.createElement('tr');
        tr.innerHTML = `
            <td>${user.full_name}</td>
            <td>${user.email}</td>
            <td>
                <a href="#">Editar</a> | 
                <a href="#">Ver</a> | 
                <a href="#">Excluir</a>
            </td>`;
        usersTable.appendChild(tr);
    });
}
