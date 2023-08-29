package maki

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MarkerEleven(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path d="M5.5-.018c-1.787 0-3.871 1.092-3.871 3.872C1.629 5.739 4.607 9.908 5.5 11c.794-1.092 3.871-5.161 3.871-7.147c0-2.779-2.084-3.87-3.871-3.87z" fill="currentColor"/>`),
		g.Group(children),
	)
}