const body = document.querySelector('body');                        // Выбираем тег "body" в разметке
const containerForButtons = document.createElement('div');          // Создаём контейнер с тегом div 
const paragraph = document.createElement("p");                      // Создаём парарграф
containerForButtons.classList.add('btn-group-vertical');
//containerForButtons.role.add('group');
containerForButtons.append(paragraph);                              // Помещаем парарграф в контейнер
//containerForButtons.      
body.append(containerForButtons);                                   // Помещаем контейнер с div в body  

//containerForButtons.innerText = 'Hello buttons'                   // Заносим надпись в контейнер

/**** Добавляем кнопки в цикле *****/
for (let i = 1; i < 6; i++) {
  const button = document.createElement('button');                  // Создаём в контейнер c элементом "button"
  //button.setAttribute('class', 'btn-group-vertical');
  //document.querySelector('button');
  //button.type.button;
  button.classList.add('.btn-primary');
  //button.addEventListener('.click', (e) => {                          // Обработчик события "левый клик мыши"
  //button.querySelector('.btn-primary');
  button.addEventListener('click', () => {
    paragraph.innerText = `Click ${i}`;                             // Присваиваем значение клика по кнопке
  })
  containerForButtons.append(button);                               // Помещаем контейнер с body в контейнер с div (containerForButtons)
  //document.querySelector('button');
  //button.type.add('button')
  //button.classList.add('btn btn-primery');
  button.innerText = `Кнопка ${i}`;                                  // Делаем на кнопке надпись с порядковым номером
}
/*** Конец добавления копок в цикле ****/

