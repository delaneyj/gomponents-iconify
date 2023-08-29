package oi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Brush(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 8 8"),
		g.Raw(`<path fill="currentColor" d="M7.44.03c-.03 0-.04.02-.06.03L3.63 2.72c-.04.03-.1.11-.13.16l-.13.25c.72.23 1.27.78 1.5 1.5l.25-.13c.05-.03.12-.08.16-.13L7.94.62c.03-.05.04-.09 0-.13L7.5.05C7.48.03 7.46.02 7.44.02zM2.66 4c-.74 0-1.31.61-1.31 1.34c0 .99-.55 1.85-1.34 2.31c.39.22.86.34 1.34.34c1.47 0 2.66-1.18 2.66-2.66c0-.74-.61-1.34-1.34-1.34z"/>`),
		g.Group(children),
	)
}