import './style.css'


// menu toggles
(() => {
    const menu = document.querySelector<HTMLDivElement>(".menu");
    const openMenuBtn = document.querySelector<HTMLButtonElement>(".open-menu-btn");
    const closeMenuBtn = document.querySelector<HTMLButtonElement>(".close-menu-btn");
    if (!menu || !openMenuBtn || !closeMenuBtn) return;
    openMenuBtn.addEventListener("click", () => {
        menu.classList.add("open");
        menu.style.transition = "transform 0.5s ease"
    })
    closeMenuBtn.addEventListener("click", () => {
        menu.classList.remove("open");
        menu.style.transition = "transform 0.5s ease"
    })
    menu.addEventListener("transitionend", function () {
        this.removeAttribute("style")
    })
    const chevrons = menu.querySelectorAll<HTMLSpanElement>(".dropdown > i")
    chevrons.forEach(arrow => {
        arrow.addEventListener("click", function () {
            const dropdown = this.closest(".dropdown")
            if (!dropdown) return;
            dropdown.classList.toggle("active")
        })
    });
})();


// theme
(() => {
    //determines if the user has a set theme
    const defaultScheme = window.matchMedia("(prefers-color-scheme: dark)").matches ? "dark" : "light";
    const theme = localStorage.getItem("theme") || defaultScheme;
    //dark theme preferred, set document with a `data-theme` attribute
    document.documentElement.setAttribute("data-theme", theme);


    const toggleSwitch = document.getElementById('checkbox-theme') as HTMLInputElement;
    if (!toggleSwitch) return;
    //listener for changing themes
    //function that changes the theme, and sets a localStorage variable
    toggleSwitch.addEventListener('change', function () {
        if (this.checked) {
            localStorage.setItem('theme', 'dark');
            document.documentElement.setAttribute('data-theme', 'dark');
        } else {
            localStorage.setItem('theme', 'light');
            document.documentElement.setAttribute('data-theme', 'light');
        }
    }, false);

    //pre-check the dark-theme checkbox if dark-theme is set
    if (theme == "dark") toggleSwitch.checked = true;
})()