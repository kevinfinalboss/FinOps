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
    })})