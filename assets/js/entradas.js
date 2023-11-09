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
  
    function addExpenseToList(expense) {
      const expenseItem = document.createElement('div');
      expenseItem.className = 'bg-white p-4 rounded-lg shadow-md mb-2';
      const expenseContent = `
        <div class="flex justify-between items-center">
          <div class="flex-1">
            <h3 class="font-bold">${expense.title}</h3>
            <p>Valor: ${expense.value}</p>
            <p>Categoria: ${expense.category}</p>
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
  
    addExpenseToList({
      title: 'Restaurante Família Silva',
      value: 'R$ 150,00',
      category: 'Alimentação',
      user: 'João',
      date: '2023-11-07 19:30'
    });
    addExpenseToList({
      title: 'Cinema "A Jornada"',
      value: 'R$ 50,00',
      category: 'Lazer',
      user: 'Maria',
      date: '2023-11-06 15:20'
    });
    addExpenseToList({
      title: 'Supermercado CompreBem',
      value: 'R$ 300,00',
      category: 'Mercado',
      user: 'Ana',
      date: '2023-11-05 10:45'
    });


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
    const formaPagamento = document.getElementById('formaPagamento').value;
    const categoria = document.getElementById('categoria').value;
    const descricao = document.getElementById('descricao').value;
    const valorNumerico = parseFloat(valor.replace('R$', '').replace(',', '.'));
    const spending = {
      title: titulo,
      date: new Date(data),
      value: valorNumerico,
      paymentMethod: formaPagamento.toLowerCase(),
      category: categoria.toLowerCase(),
      description: descricao
    };
    fetch('/api/v1/user/register/spendings', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
        'Authorization': `Bearer ${token}`
      },
      body: JSON.stringify(spending)
    }).then(response => response.json()).then(data => {
      console.log(data);
    }).catch((error) => {
      console.error('Error:', error);
    });
})
})