package pepicons_print

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Watch(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor"><g opacity=".2"><path d="M16.1 11.25a4.5 4.5 0 1 1-9 0a4.5 4.5 0 0 1 9 0Z"/><path fill-rule="evenodd" d="M11.6 9a.75.75 0 0 1 .75.75v1.5a.75.75 0 0 1-1.5 0v-1.5A.75.75 0 0 1 11.6 9Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M13.85 11.25a.75.75 0 0 1-.75.75h-1.5a.75.75 0 0 1 0-1.5h1.5a.75.75 0 0 1 .75.75ZM9.6 3.75a.5.5 0 0 1 .5-.5h3a.5.5 0 0 1 0 1h-3a.5.5 0 0 1-.5-.5Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M10.244 3.271a.5.5 0 0 1 .335.623l-1.5 5a.5.5 0 1 1-.958-.288l1.5-5a.5.5 0 0 1 .623-.335Zm0 15.958a.5.5 0 0 0 .335-.623l-1.5-5a.5.5 0 1 0-.958.288l1.5 5a.5.5 0 0 0 .623.335Zm2.712 0a.5.5 0 0 1-.335-.623l1.5-5a.5.5 0 1 1 .958.288l-1.5 5a.5.5 0 0 1-.623.335Zm0-15.958a.5.5 0 0 0-.335.623l1.5 5a.5.5 0 1 0 .958-.288l-1.5-5a.5.5 0 0 0-.623-.335Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M9.6 18.75a.5.5 0 0 1 .5-.5h3a.5.5 0 0 1 0 1h-3a.5.5 0 0 1-.5-.5Z" clip-rule="evenodd"/><path d="M10.1 4.25h3l1 3h-5l1-3Zm3 14h-3l-1-3h5l-1 3Z"/></g><path fill-rule="evenodd" d="M10 6a4 4 0 1 0 0 8a4 4 0 0 0 0-8Zm-5 4a5 5 0 1 1 10 0a5 5 0 0 1-10 0Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M10 8a.5.5 0 0 1 .5.5V10a.5.5 0 0 1-1 0V8.5A.5.5 0 0 1 10 8Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M12 10a.5.5 0 0 1-.5.5H10a.5.5 0 0 1 0-1h1.5a.5.5 0 0 1 .5.5ZM8 2.5a.5.5 0 0 1 .5-.5h3a.5.5 0 0 1 0 1h-3a.5.5 0 0 1-.5-.5Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M8.637 2.02a.5.5 0 0 1 .344.617l-1 3.5a.5.5 0 0 1-.962-.274l1-3.5a.5.5 0 0 1 .618-.344Zm0 15.96a.5.5 0 0 0 .344-.617l-1-3.5a.5.5 0 0 0-.962.274l1 3.5a.5.5 0 0 0 .618.344Zm2.726 0a.5.5 0 0 1-.344-.617l1-3.5a.5.5 0 0 1 .962.274l-1 3.5a.5.5 0 0 1-.618.344Zm0-15.96a.5.5 0 0 0-.344.617l1 3.5a.5.5 0 0 0 .962-.274l-1-3.5a.5.5 0 0 0-.618-.344Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M8 17.5a.5.5 0 0 1 .5-.5h3a.5.5 0 0 1 0 1h-3a.5.5 0 0 1-.5-.5Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}