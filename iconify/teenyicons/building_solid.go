package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BuildingSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="M7.724.053a.5.5 0 0 0-.448 0l-6 3l.448.894L7.5 1.06l5.776 2.888l.448-.894l-6-3ZM14 6h1V5H0v1h1v8H0v1h4V8h5v7h6v-1h-1V6Z"/><path fill="currentColor" d="M8 15V9H5v6h3Z"/>`),
		g.Group(children),
	)
}