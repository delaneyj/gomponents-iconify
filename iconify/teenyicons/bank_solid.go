package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BankSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="M7.724.053a.5.5 0 0 0-.448 0l-6 3l.448.894L7.5 1.06l5.776 2.888l.448-.894l-6-3Z"/><path fill="currentColor" fill-rule="evenodd" d="M14 6h1V5H0v1h1v8H0v1h15v-1h-1V6Zm-9 6V8h1v4H5Zm4 0V8h1v4H9Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}