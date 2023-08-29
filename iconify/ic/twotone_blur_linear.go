package ic

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TwotoneBlurLinear(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M17 16.5c.28 0 .5-.22.5-.5s-.22-.5-.5-.5s-.5.22-.5.5s.22.5.5.5z"/><circle cx="9" cy="12" r="1" fill="currentColor"/><circle cx="13" cy="8" r="1" fill="currentColor"/><circle cx="13" cy="16" r="1" fill="currentColor"/><path fill="currentColor" d="M17 12.5c.28 0 .5-.22.5-.5s-.22-.5-.5-.5s-.5.22-.5.5s.22.5.5.5z"/><circle cx="13" cy="12" r="1" fill="currentColor"/><path fill="currentColor" d="M3 3h18v2H3z"/><circle cx="5" cy="8" r="1.5" fill="currentColor"/><circle cx="5" cy="12" r="1.5" fill="currentColor"/><circle cx="5" cy="16" r="1.5" fill="currentColor"/><path fill="currentColor" d="M17 8.5c.28 0 .5-.22.5-.5s-.22-.5-.5-.5s-.5.22-.5.5s.22.5.5.5z"/><circle cx="9" cy="16" r="1" fill="currentColor"/><circle cx="9" cy="8" r="1" fill="currentColor"/><path fill="currentColor" d="M3 19h18v2H3z"/>`),
		g.Group(children),
	)
}