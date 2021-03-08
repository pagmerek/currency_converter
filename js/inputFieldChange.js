const finput = document.getElementById('first_form');
const sinput = document.getElementById('second_form');
const rate = document.getElementById('rate').textContent
finput.addEventListener('input', secondUpdateValue);
sinput.addEventListener('input', firstUpdateValue);

function secondUpdateValue(e) {
  sinput.value = (parseFloat(e.target.value) * rate).toFixed(2).toString();
}

function firstUpdateValue(e) {
    finput.value = ((parseFloat(e.target.value) * (1./rate)).toFixed(2)).toString()
}