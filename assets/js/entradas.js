document.addEventListener('DOMContentLoaded', function() {
  initializeMenu();
  fetchMonthlyExpensesTotal();
  fetchMonthlyIncomesTotal();
  fetchRecentExpenses();
  initializeFormSubmission();
});

function initializeMenu() {
  const btn = document.querySelector(".mobile-menu-button");
  const menu = document.querySelector(".mobile-menu");

  btn.addEventListener("click", function() {
      toggleMenu(menu);
  });
}

function toggleMenu(menu) {
  if (menu.classList.contains("hidden")) {
      menu.classList.remove("hidden");
      menu.classList.add("entering");
      setTimeout(() => menu.classList.remove("entering"), 500);
  } else {
      menu.classList.add("leaving");
      setTimeout(() => {
          menu.classList.add("hidden");
          menu.classList.remove("leaving");
      }, 500);
  }
}

function fetchMonthlyExpensesTotal() {
  const formattedMonth = getFormattedCurrentMonth();
  fetch(`/api/v1/spendings/sumByMonth?month=${formattedMonth}`)
      .then(response => response.json())
      .then(data => updatePanelValue('totalSaidas', data.total))
      .catch(error => console.error('Erro ao buscar total de gastos do mês:', error));
}

function fetchMonthlyIncomesTotal() {
  const formattedMonth = getFormattedCurrentMonth();
  fetch(`/api/v1/incomes/sumByMonth?month=${formattedMonth}`)
      .then(response => response.json())
      .then(data => updatePanelValue('totalEntradas', data.total))
      .catch(error => console.error('Erro ao buscar total de entradas do mês:', error));
}

function updatePanelValue(elementId, total) {
  const element = document.getElementById(elementId);
  if (element) {
      element.textContent = `R$ ${total.toFixed(2)}`;
      if (elementId === 'totalSaidas') {
          compareExpensesAndIncomes();
      }
  }
}

function compareExpensesAndIncomes() {
  const totalSaidas = parseFloat(document.getElementById('totalSaidas').textContent.replace('R$', '').trim());
  const totalEntradas = parseFloat(document.getElementById('totalEntradas').textContent.replace('R$', '').trim());

  const saidaAlerta = document.getElementById('saidaAlerta');
  const panelSaidas = document.getElementById('panelSaidas');
  
  if (totalSaidas > totalEntradas) {
      saidaAlerta.classList.remove('hidden');
      panelSaidas.classList.add('bg-red-600');
      panelSaidas.classList.remove('bg-red-500');
  } else {
      saidaAlerta.classList.add('hidden');
      panelSaidas.classList.remove('bg-red-600');
      panelSaidas.classList.add('bg-red-600');
  }
}


function fetchRecentExpenses() {
  fetch('/api/v1/incomes/recent')
      .then(response => response.json())
      .then(data => {
          const incomes = data.recent_incomes;
          if (incomes.length === 0) {
              showNoExpensesMessage();
              return;
          }
          incomes.forEach(income => addIncomeToList(income));
      })
      .catch(error => {
          console.error('Erro ao buscar entradas recentes:', error);
          showNoExpensesMessage();
      });
}

function showNoExpensesMessage() {
  const expensesList = document.getElementById('expensesList');
  expensesList.innerHTML = '<div class="text-center py-4 bg-white rounded-lg shadow-md"><p class="text-gray-600 font-semibold">Sem entradas esse mês</p></div>';
}

function addIncomeToList(income) {
  const incomeItem = createIncomeItem(income);
  document.getElementById('expensesList').appendChild(incomeItem);
}

function createIncomeItem(income) {
  const incomeItem = document.createElement('div');
  incomeItem.className = 'bg-white p-4 rounded-lg shadow-md mb-2';
  incomeItem.innerHTML = getIncomeItemHTML(income);
  return incomeItem;
}

function getIncomeItemHTML(income) {
  return `
      <div class="flex justify-between items-center">
          <div class="flex-1">
              <h3 class="font-bold">${income.title}</h3>
              <p>Valor: R$ ${income.value.toFixed(2)}</p>
              <p>Método de Recebimento: ${income.incomeMethod}</p>
          </div>
          <div class="details-button" onclick="toggleDetails(this)">+</div>
      </div>
      <div class="details">
          <p>Cadastrado por: ${income.author}</p>
          <p>Data: ${new Date(income.date).toLocaleDateString('pt-BR')}</p>
      </div>
  `;
}

window.toggleDetails = function(button) {
  const details = button.parentElement.nextElementSibling;
  details.style.display = details.style.display === 'none' ? 'block' : 'none';
  button.textContent = details.style.display === 'none' ? '+' : '-';
};

function getFormattedCurrentMonth() {
  const today = new Date();
  const month = today.getMonth() + 1;
  const year = today.getFullYear();
  return `${month}/${year}`;
}

function initializeFormSubmission() {
  const form = document.querySelector('form');
  form.addEventListener('submit', submitForm);
}

function submitForm(event) {
  event.preventDefault();
  const income = getFormData();
  postIncome(income)
      .then(response => {
          const messageType = response.ok ? "success" : "error";
          const messageText = response.ok ? "Entrada registrada com sucesso!" : "Falha ao registrar a entrada.";
          showNotification(messageText, messageType);
          return response.json();
      })
      .catch(() => showNotification("Erro ao conectar ao servidor.", "error"));
}

function getFormData() {
  return {
      title: document.getElementById('titulo').value,
      date: formatDate(document.getElementById('data').value),
      value: parseFloat(document.getElementById('valor').value.replace('R$', '').replace(',', '.')),
      incomeMethod: document.getElementById('incomeMethod').value,
      description: document.getElementById('descricao').value
  };
}

function postIncome(income) {
  const token = getCookie('token');
  return fetch('/api/v1/incomes/register', {
      method: 'POST',
      headers: {
          'Content-Type': 'application/json',
          'Authorization': `Bearer ${token}`
      },
      body: JSON.stringify(income)
  });
}

function formatDate(inputDate) {
  const parts = inputDate.split('-');
  return `${parts[2]}/${parts[1]}/${parts[0]}`;
}

function showNotification(message, type) {
  const notification = document.createElement("div");
  notification.className = `notification ${type === "success" ? "success" : ""}`;
  notification.innerText = message;
  notification.style.animation = 'slideIn 0.5s, fadeOut 3s 2.5s forwards';
  document.body.appendChild(notification);
  setTimeout(() => notification.remove(), 6000);
}

function getCookie(name) {
  const cookieArr = document.cookie.split(";");
  for (let i = 0; i < cookieArr.length; i++) {
      const cookiePair = cookieArr[i].split("=");
      if (name === cookiePair[0].trim()) {
          return decodeURIComponent(cookiePair[1]);
      }
  }
  return null;
}
