package ic

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TwotoneAutoAwesomeMotion(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 12h8v8h-8z" opacity=".3"/><path fill="currentColor" d="M14 2H4c-1.1 0-2 .9-2 2v10h2V4h10V2zm6 8h-8c-1.1 0-2 .9-2 2v8c0 1.1.9 2 2 2h8c1.1 0 2-.9 2-2v-8c0-1.1-.9-2-2-2zm0 10h-8v-8h8v8z"/><path fill="currentColor" d="M18 6H8c-1.1 0-2 .9-2 2v10h2V8h10V6z"/>`),
		g.Group(children),
	)
}