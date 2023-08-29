package la

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Yahoo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M7.34 6.027s.785.172 1.441.172s1.426-.183 1.426-.183L16 15.66L21.82 6s.63.215 1.414.215c.782 0 1.426-.2 1.426-.2l-7.445 12.583l.2 9.375s-.786-.215-1.415-.215c-.625 0-1.441.242-1.441.242l.226-9.414z"/>`),
		g.Group(children),
	)
}