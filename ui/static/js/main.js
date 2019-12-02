const BASE_URL = "https://stranger-trade-wars.herokuapp.com"

var navLinks = document.querySelectorAll("nav a");
for (var i = 0; i < navLinks.length; i++) {
	var link = navLinks[i]
	if (link.getAttribute('href') == window.location.pathname) {
		link.classList.add("live");
		break;
	}
}

var grid = document.getElementsByClassName("grid");
var ctx = grid.getContext("2d");
ctx.moveTo(0,0);
ctx.lineTo(200,100);
ctx.stroke();



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
