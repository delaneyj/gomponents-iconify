package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Spectrum(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 16h4a8 8 0 0 0-8-8v4a4 4 0 0 1 4 4Z"/>`),
		g.Group(children),
	)
}