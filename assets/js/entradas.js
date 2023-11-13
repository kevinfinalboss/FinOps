document.addEventListener('DOMContentLoaded', function() {
    const btn = document.querySelector(".mobile-menu-button");
    const menu = document.querySelector(".mobile-menu");
  
    btn.addEventListener("click", function() {
      if (menu.classList.contains("hidden")) {
        menu.classList.remove("hidden");
        menu.classList.add("entering");
        setTimeout(function() {
          menu.classList.remove("entering");
        }, 500);
      } else {
        menu.classList.add("leaving");
        setTimeout(function() {
          menu.classList.add("hidden");
          menu.classList.remove("leaving");
        }, 500);
      }
    });
    const entradas = 0;
    const saidas = 0;
    const panelSaidas = document.getElementById('panelSaidas');
  
    if (panelSaidas) {
      if (saidas > entradas) {
        panelSaidas.classList.remove('bg-green-500');
        panelSaidas.classList.add('bg-red-500');
      } else {
        panelSaidas.classList.remove('bg-red-500');
        panelSaidas.classList.add('bg-green-500');
      }
    }
  
    window.toggleDetails = function(button) {
      const details = button.parentElement.nextElementSibling;
      if (details.style.display === 'none') {
        details.style.display = 'block';
        button.textContent = '-';
      } else {
        details.style.display = 'none';
        button.textContent = '+';
      }
    };

    fetchRecentExpenses();

    function fetchRecentExpenses() {
        fetch('/api/v1/spendings/recent')
            .then(response => response.json())
            .then(data => {
                const expenses = data.recent_spendings;
                expenses.forEach(expense => {
                    addExpenseToList({
                        title: expense.title,
                        value: `R$ ${expense.value.toFixed(2)}`,
                        category: expense.category,
                        user: expense.author,
                        paymentMethod: expense.paymentMethod,
                        date: new Date(expense.date).toLocaleString()
                    });
                });
            })
            .catch(error => {
              console.error('Erro ao buscar gastos recentes:', error);
              showNoExpensesMessage();
          });
  }

    function showNoExpensesMessage() {
      const expensesList = document.getElementById('expensesList');
      expensesList.innerHTML = '<div class="text-center py-4 bg-white rounded-lg shadow-md"><p class="text-gray-600 font-semibold">Sem saídas esse mês</p></div>';
  }
  
    function addExpenseToList(expense) {
        const expenseItem = document.createElement('div');
        expenseItem.className = 'bg-white p-4 rounded-lg shadow-md mb-2';
        const expenseContent = `
            <div class="flex justify-between items-center">
                <div class="flex-1">
                    <h3 class="font-bold">${expense.title}</h3>
                    <p>Valor: ${expense.value}</p>
                    <p>Categoria: ${expense.category}</p>
                    <p>Pagamento: ${expense.paymentMethod}</p>
                </div>
                <div class="details-button" onclick="toggleDetails(this)">+</div>
            </div>
            <div class="details">
                <p>Cadastrado por: ${expense.user}</p>
                <p>Data e hora: ${expense.date}</p>
            </div>
        `;
        expenseItem.innerHTML = expenseContent;
        document.getElementById('expensesList').appendChild(expenseItem);
    }


  const dataExpenses = {
    labels: ['Janeiro', 'Fevereiro', 'Março', 'Abril', 'Maio', 'Junho', 'Julho'],
    datasets: [{
      label: 'Despesas por mês',
      data: [500, 400, 300, 700, 200, 300, 400],
      backgroundColor: 'rgba(255, 99, 132, 0.2)',
      borderColor: 'rgba(255, 99, 132, 1)',
      borderWidth: 1
    }]
  };
  const dataIncome = {
    labels: ['Janeiro', 'Fevereiro', 'Março', 'Abril', 'Maio', 'Junho'],
    datasets: [{
      label: 'Entrada por mês',
      data: [600, 700, 800, 200, 300, 400],
      backgroundColor: 'rgba(54, 162, 235, 0.2)',
      borderColor: 'rgba(54, 162, 235, 1)',
      borderWidth: 1
    }]
  };
  let currentChart = new Chart(document.getElementById('myChart'), {
    type: 'bar',
    data: dataExpenses,
    options: {
      scales: {
        y: {
          beginAtZero: true
        }
      }
    },
  });
  document.getElementById('chartType').addEventListener('change', function() {
    let selectedOption = this.value;
    let newData;
    switch (selectedOption) {
      case 'expenses':
        newData = dataExpenses;
        break;
      case 'income':
        newData = dataIncome;
        break;
      case 'category':
        break;
    }
    currentChart.destroy();
    currentChart = new Chart(document.getElementById('myChart'), {
      type: 'bar',
      data: newData,
      options: {
        scales: {
          y: {
            beginAtZero: true
          }
        }
      },
    });
  });

  function getCookie(name) {
    let cookieArr = document.cookie.split(";");
    for (let i = 0; i < cookieArr.length; i++) {
      let cookiePair = cookieArr[i].split("=");
      if (name == cookiePair[0].trim()) {
        return decodeURIComponent(cookiePair[1]);
      }
    }
    return null;
  }
  const form = document.querySelector('form');
  form.addEventListener('submit', function(event) {
      event.preventDefault();
      const token = getCookie('token');
      const titulo = document.getElementById('titulo').value;
      const data = document.getElementById('data').value;
      const valor = document.getElementById('valor').value;
      const categoria = document.getElementById('categoria').value;
      const descricao = document.getElementById('descricao').value;
      const valorNumerico = parseFloat(valor.replace('R$', '').replace(',', '.'));
      const formaPagamento = document.getElementById('formaPagamento').value;
      const spending = {
          title: titulo,
          date: new Date(data),
          value: valorNumerico,
          category: categoria.toLowerCase(),
          paymentMethod: formaPagamento,
          description: descricao
      };
      fetch('/api/v1/user/register/spendings', {
          method: 'POST',
          headers: {
              'Content-Type': 'application/json',
              'Authorization': `Bearer ${token}`
          },
          body: JSON.stringify(spending)
      }).then(response => {
          if (response.ok) {
              showNotification("Gasto registrado com sucesso!", "success");
              clearFormFields();
          } else {
              showNotification("Falha ao registrar o gasto.", "error");
          }
          return response.json();
      }).then(data => {
          console.log(data);
      }).catch((error) => {
          console.error('Error:', error);
          showNotification("Erro ao conectar ao servidor.", "error");
      });
  });
});

function showNotification(message, type) {
  const notification = document.createElement("div");
  notification.className = `notification ${type === "success" ? "success" : ""}`;
  notification.innerText = message;

  notification.style.animation = 'slideIn 0.5s, fadeOut 3s 2.5s forwards';

  document.body.appendChild(notification);
  
  setTimeout(() => {
      notification.remove();
  }, 6000);
}

function clearFormFields() {
  document.getElementById('titulo').value = '';
  document.getElementById('data').value = '';
  document.getElementById('valor').value = '';
  document.getElementById('categoria').selectedIndex = 0;
  document.getElementById('descricao').value = '';
}