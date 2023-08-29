package bytesize

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Music(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M11 25V6l13-3v20M11 13l13-3"/><ellipse cx="7" cy="25" rx="4" ry="5"/><ellipse cx="20" cy="23" rx="4" ry="5"/></g>`),
		g.Group(children),
	)
}