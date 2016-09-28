/*jslint browser: true*/

function getcookie(cname) {
    "use strict";
    var i = 0, name = cname + "=", ca = document.cookie.split(';');
    for(i; i < ca.length; i++) {
	var c = ca[i];
	while(c.charAt(0)==' ') {
	    c = c.substring(1);
	}
	if (c.indexOf(name) == 0) {
	    return c.substring(name.length, c.length);
	}
    }
    return "";
}

function enter() {
    "use strict";
    window.document.getElementById("postcontent").value = unescape(getcookie("content"));
    window.document.getElementById("postkey").value = getcookie("key");
    window.document.getElementById("posttitle").value = getcookie("title");
}

function leave() {
    "use strict";
    document.cookie = "content=" + escape(window.document.getElementById("postcontent").value) + ";";
    document.cookie = "key=" + window.document.getElementById("postkey").value + ";";
    document.cookie = "title=" + window.document.getElementById("posttitle").value + ";";
}

function send() {
    "use strict";
    var key, title, cnt, xhr;
    cnt = window.document.getElementById("postcontent").value;
    key = window.document.getElementById("postkey").value;
    title = window.document.getElementById("posttitle").value;

    xhr = new XMLHttpRequest();
    xhr.open("POST", "/post/");
    xhr.setRequestHeader('Content-Type', 'application/json; charset=UTF-8');
    xhr.send(JSON.stringify({Key: key, Title: title, Content: cnt}));
}
