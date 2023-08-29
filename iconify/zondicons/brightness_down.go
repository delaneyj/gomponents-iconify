package zondicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BrightnessDown(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M10 13a3 3 0 1 1 0-6a3 3 0 0 1 0 6zM9 4a1 1 0 1 1 2 0a1 1 0 1 1-2 0zm4.54 1.05a1 1 0 1 1 1.41 1.41a1 1 0 1 1-1.41-1.41zM16 9a1 1 0 1 1 0 2a1 1 0 1 1 0-2zm-1.05 4.54a1 1 0 1 1-1.41 1.41a1 1 0 1 1 1.41-1.41zM11 16a1 1 0 1 1-2 0a1 1 0 1 1 2 0zm-4.54-1.05a1 1 0 1 1-1.41-1.41a1 1 0 1 1 1.41 1.41zM4 11a1 1 0 1 1 0-2a1 1 0 1 1 0 2zm1.05-4.54a1 1 0 1 1 1.41-1.41a1 1 0 1 1-1.41 1.41z"/>`),
		g.Group(children),
	)
}