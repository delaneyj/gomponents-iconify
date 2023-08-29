package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ComponentRadio(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 4a8 8 0 1 0 0 16a8 8 0 0 0 0-16ZM2 12C2 6.477 6.477 2 12 2s10 4.477 10 10s-4.477 10-10 10S2 17.523 2 12Zm10-4a4 4 0 1 0 0 8a4 4 0 0 0 0-8Zm-6 4a6 6 0 1 1 12 0a6 6 0 0 1-12 0Z"/>`),
		g.Group(children),
	)
}