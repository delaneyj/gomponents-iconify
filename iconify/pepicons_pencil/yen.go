package pepicons_pencil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Yen(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor" fill-rule="evenodd" clip-rule="evenodd"><path d="M9.336 8.87a.5.5 0 0 1-.706-.034l-5-5.512a.5.5 0 1 1 .74-.672l5 5.512a.5.5 0 0 1-.034.706Z"/><path d="M8.664 8.87a.5.5 0 0 0 .706-.034l5-5.512a.5.5 0 1 0-.74-.672l-5 5.512a.5.5 0 0 0 .034.706Z"/><path d="M3.5 8.988a.5.5 0 0 1 .5-.5h10a.5.5 0 0 1 0 1H4a.5.5 0 0 1-.5-.5Zm0 3.5a.5.5 0 0 1 .5-.5h10a.5.5 0 0 1 0 1H4a.5.5 0 0 1-.5-.5Z"/><path d="M9 8a.5.5 0 0 1 .5.5v8.488a.5.5 0 0 1-1 0V8.5A.5.5 0 0 1 9 8Z"/></g>`),
		g.Group(children),
	)
}