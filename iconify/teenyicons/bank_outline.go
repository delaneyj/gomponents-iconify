package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BankOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="m7.5.5l.224-.447a.5.5 0 0 0-.448 0L7.5.5ZM0 15h15v-1H0v1ZM7.276.053l-6 3l.448.894l6-3l-.448-.894ZM0 6h15V5H0v1Zm13.724-2.947l-6-3l-.448.894l6 3l.448-.894ZM5 8v4h1V8H5Zm4 0v4h1V8H9ZM1 5.5v9h1v-9H1Zm12 0v9h1v-9h-1Z"/>`),
		g.Group(children),
	)
}