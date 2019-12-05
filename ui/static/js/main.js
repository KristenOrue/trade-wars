const BASE_URL = "https://stranger-trade-wars.herokuapp.com"

var navLinks = document.querySelectorAll("nav a");
for (var i = 0; i < navLinks.length; i++) {
	var link = navLinks[i]
	if (link.getAttribute('href') == window.location.pathname) {
		link.classList.add("live");
		break;
	}
}

var canvas = document.getElementsByClassName("grid");
var elem = document.getElementById("ship");

// ctx.moveTo(0,0);
// ctx.lineTo(200,100);
// ctx.stroke();  

function moveRight() {
	var elem = document.getElementById("ship"); 
    // elem.style.left = '0px';   
	// elem.style.left = parseInt(elem.style.left) + 10 + 'px'; 

	var pos = 0;
	var id = setInterval(frame, 0);
	function frame() {
	  if (pos == 675) {
		clearInterval(id);
	  } else {
		pos++; 
		// elem.style.top = pos + 'px'; 
		elem.style.left = pos + 'px'; 
	  }
	}
  }

function moveDown() {
	var elem = document.getElementById("ship");
	var id = setInterval(frame, 0);
	function frame() {
		if (pos == 675) {
		clearInterval(id);
		} else {
		pos++; 
		elem.style.top = pos + 'px';  
		}
	}
}

var objImage= null;
function init(){
	objImage=document.getElementById("ship");				
	objImage.style.position='relative';
	objImage.style.left='0px';
	objImage.style.top='0px';
}

function getMove(e){				
	var key_code=e.which||e.keyCode;
		switch(key_code){
			case 65: // A arrow key
				moveLeft();
				break;
			case 87: // W arrow key
				moveUp();
				break;
			case 68: // D arrow key
				moveRight();
				break;
			case 83: // S arrow key
				moveDown();
				break;						
		}
	}

	function moveLeft(){
		objImage.style.left=parseInt(objImage.style.left)-5 +'px';
	}

	function moveUp(){
		objImage.style.top=parseInt(objImage.style.top)-5 +'px';
	}

	function moveRight(){
		objImage.style.left=parseInt(objImage.style.left)+5 +'px';
	}

	function moveDown(){
		objImage.style.top=parseInt(objImage.style.top)+5 +'px';
	}

	window.onload=init;
