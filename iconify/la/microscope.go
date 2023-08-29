package la

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Microscope(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M17 4v2h4V4zm-1 3v2h-2c-2.746 0-5 2.254-5 5v4.188c-1.156.417-2 1.519-2 2.812s.844 2.395 2 2.813V26H7v2h17v-2H11v-2.188A3.016 3.016 0 0 0 12.813 22H24v-2H12.812A3.036 3.036 0 0 0 11 18.187V14a3 3 0 0 1 3-3h2v6h6V7zm2 2h2v6h-2zm-8 11c.563 0 1 .438 1 1c0 .563-.438 1-1 1c-.563 0-1-.438-1-1c0-.563.438-1 1-1z"/>`),
		g.Group(children),
	)
}