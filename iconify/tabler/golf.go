package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Golf(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M12 18V3l7 4l-7 4"/><path d="M9 17.67c-.62.36-1 .82-1 1.33c0 1.1 1.8 2 4 2s4-.9 4-2c0-.5-.38-.97-1-1.33"/></g>`),
		g.Group(children),
	)
}