console.log('XD')
const finput = document.getElementById('first_form');
const sinput = document.getElementById('second_form');
const rate = document.getElementById('rate').value

finput.addEventListener('input', secondUpdateValue);
sinput.addEventListener('input', firstUpdateValue);

function secondUpdateValue(e) {
  sinput.value = (parseInt(e.target.value) * rate).toString();
  console.log('xd')
}

function firstUpdateValue(e) {
    finput.value = (parseInt(e.target.value) * (1./rate)).toString()
}