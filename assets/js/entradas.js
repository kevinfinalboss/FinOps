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
    fetchMonthlyExpensesTotal();

    function fetchMonthlyExpensesTotal() {
        const today = new Date();
        const month = today.getMonth() + 1;
        const year = today.getFullYear();
        const formattedMonth = `${month}/${year}`;

        fetch(`/api/v1/spendings/sumByMonth?month=${formattedMonth}`)
            .then(response => response.json())
            .then(data => {
                updateExpensesPanel(data.total);
            })
            .catch(error => {
                console.error('Erro ao buscar total de gastos do mês:', error);
            });
    }

    function updateExpensesPanel(total) {
        const panelSaidas = document.getElementById('panelSaidas');
        const totalValueElement = panelSaidas.querySelector('p');
        totalValueElement.textContent = `R$ ${total.toFixed(2)}`;
    }
    fetchRecentExpenses();

    function fetchRecentExpenses() {
      fetch('/api/v1/incomes/recent')
          .then(response => response.json())
          .then(data => {
              const incomes = data.recent_incomes;
              if (incomes.length === 0) {
                  showNoExpensesMessage();
                  return;
              }
              incomes.forEach(income => {
                  addIncomeToList({
                      title: income.title,
                      value: `R$ ${income.value.toFixed(2)}`,
                      user: income.author,
                      incomeMethod: income.incomeMethod,
                      date: new Date(income.date).toLocaleDateString('pt-BR')
                  });
              });
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
      const incomeItem = document.createElement('div');
      incomeItem.className = 'bg-white p-4 rounded-lg shadow-md mb-2';
      const incomeContent = `
          <div class="flex justify-between items-center">
              <div class="flex-1">
                  <h3 class="font-bold">${income.title}</h3>
                  <p>Valor: ${income.value}</p>
                  <p>Método de Recebimento: ${income.incomeMethod}</p>
              </div>
              <div class="details-button" onclick="toggleDetails(this)">+</div>
          </div>
          <div class="details">
              <p>Cadastrado por: ${income.user}</p>
              <p>Data: ${income.date}</p>
          </div>
      `;
      incomeItem.innerHTML = incomeContent;
      document.getElementById('expensesList').appendChild(incomeItem);
  }

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
      const dataElement = document.getElementById('data').value;
      const valor = document.getElementById('valor').value;
      const incomeMethod = document.getElementById('incomeMethod').value;
      const descricao = document.getElementById('descricao').value;
      const valorNumerico = parseFloat(valor.replace('R$', '').replace(',', '.'));

      const data = formatDate(dataElement);

      const income = {
          title: titulo,
          date: data,
          value: valorNumerico,
          incomeMethod: incomeMethod,
          description: descricao
      };

      fetch('/api/v1/incomes/register', {
          method: 'POST',
          headers: {
              'Content-Type': 'application/json',
              'Authorization': `Bearer ${token}`
          },
          body: JSON.stringify(income)
      }).then(response => {
          if (response.ok) {
              showNotification("Entrada registrada com sucesso!", "success");
          } else {
              showNotification("Falha ao registrar a entrada.", "error");
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
  
  setTimeout(() => {
      notification.remove();
  }, 6000);
}