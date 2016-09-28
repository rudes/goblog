autocmd BufWritePost * silent !docker cp static/css/base.css otherlettersnet_web_1:/go/src/github.com/rudes/otherletters.net/static/css/base.css
autocmd BufWritePost * silent !docker cp static/js/base.js otherlettersnet_web_1:/go/src/github.com/rudes/otherletters.net/static/js/base.js
autocmd BufWritePost * silent !docker cp static/templates/index.tmpl otherlettersnet_web_1:/go/src/github.com/rudes/otherletters.net/static/templates/index.tmpl
autocmd BufWritePost * silent !docker cp static/templates/header.tmpl otherlettersnet_web_1:/go/src/github.com/rudes/otherletters.net/static/templates/header.tmpl
autocmd BufWritePost * silent !docker cp static/templates/base.tmpl otherlettersnet_web_1:/go/src/github.com/rudes/otherletters.net/static/templates/base.tmpl
autocmd BufWritePost * silent !docker cp static/templates/post.tmpl otherlettersnet_web_1:/go/src/github.com/rudes/otherletters.net/static/templates/post.tmpl
