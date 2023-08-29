package radix_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Section(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M2 1.5a.5.5 0 1 1-1 0a.5.5 0 0 1 1 0ZM2 5v5h11V5H2Zm0-1a1 1 0 0 0-1 1v5a1 1 0 0 0 1 1h11a1 1 0 0 0 1-1V5a1 1 0 0 0-1-1H2Zm-.5 10a.5.5 0 1 0 0-1a.5.5 0 0 0 0 1ZM4 1.5a.5.5 0 1 1-1 0a.5.5 0 0 1 1 0ZM3.5 14a.5.5 0 1 0 0-1a.5.5 0 0 0 0 1ZM6 1.5a.5.5 0 1 1-1 0a.5.5 0 0 1 1 0ZM5.5 14a.5.5 0 1 0 0-1a.5.5 0 0 0 0 1ZM8 1.5a.5.5 0 1 1-1 0a.5.5 0 0 1 1 0ZM7.5 14a.5.5 0 1 0 0-1a.5.5 0 0 0 0 1ZM10 1.5a.5.5 0 1 1-1 0a.5.5 0 0 1 1 0ZM9.5 14a.5.5 0 1 0 0-1a.5.5 0 0 0 0 1ZM12 1.5a.5.5 0 1 1-1 0a.5.5 0 0 1 1 0ZM11.5 14a.5.5 0 1 0 0-1a.5.5 0 0 0 0 1ZM14 1.5a.5.5 0 1 1-1 0a.5.5 0 0 1 1 0ZM13.5 14a.5.5 0 1 0 0-1a.5.5 0 0 0 0 1Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}