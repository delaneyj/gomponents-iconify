package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Hypocam(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M6.5 21.813h4.375v4.375H6.5zm4.375 0h4.375v4.375h-4.375zm4.375 0h4.375v4.375H15.25zm4.375 0H24v4.375h-4.375zm4.375 0h4.375v4.375H24zm4.375 0h4.375v4.375h-4.375zm4.375 0h4.375v4.375H32.75zm4.375 0H41.5v4.375h-4.375z"/><circle cx="24" cy="24" r="21.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/>`),
		g.Group(children),
	)
}