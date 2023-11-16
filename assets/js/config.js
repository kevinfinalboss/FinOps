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
    fetch('/api/v1/users/list')
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
    const usersList = document.getElementById('usersList');
    users.forEach(user => {
        const userCard = document.createElement('div');
        userCard.className = 'p-4 bg-white rounded-lg border border-gray-200 shadow';
        userCard.innerHTML = `
            <h3 class="text-lg font-semibold">${user.full_name}</h3>
            <p class="text-gray-600">${user.email}</p>
            <div class="mt-3">
                <a href="#" class="text-blue-600 hover:text-blue-800">Editar</a> | 
                <a href="#" class="text-green-600 hover:text-green-800">Ver</a> | 
                <a href="#" class="text-red-600 hover:text-red-800">Excluir</a>
            </div>`;
        usersList.appendChild(userCard);
    });
}
