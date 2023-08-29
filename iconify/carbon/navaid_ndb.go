package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NavaidNdb(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M16 24a8 8 0 1 1 8-8a8.01 8.01 0 0 1-8 8Zm0-14a6 6 0 1 0 6 6a6.007 6.007 0 0 0-6-6Z"/><circle cx="16" cy="16" r="2" fill="currentColor"/><circle cx="16" cy="4" r="2" fill="currentColor"/><circle cx="16" cy="28" r="2" fill="currentColor"/><circle cx="28" cy="16" r="2" fill="currentColor"/><circle cx="4" cy="16" r="2" fill="currentColor"/><circle cx="7.515" cy="7.515" r="2" fill="currentColor"/><circle cx="24.485" cy="24.485" r="2" fill="currentColor"/><circle cx="24.485" cy="7.515" r="2" fill="currentColor"/><circle cx="7.515" cy="24.485" r="2" fill="currentColor"/>`),
		g.Group(children),
	)
}