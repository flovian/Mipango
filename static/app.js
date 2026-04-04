document.addEventListener("DOMContentLoaded", () => {
    const modal = document.getElementById("newObjectiveModal");
    const newBtn = document.getElementById("newObjectiveBtn");
    const closeBtn = document.querySelector(".close-btn");

    if (newBtn && modal) {
        newBtn.onclick = () => { modal.classList.add("active"); }
    }
    if (closeBtn && modal) {
        closeBtn.onclick = () => { modal.classList.remove("active"); }
    }
    window.onclick = (e) => {
        if (e.target == modal) modal.classList.remove("active");
    }
});
