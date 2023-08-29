package la

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LayerGroup(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="m16 5.938l-.375.125l-10 4L3.312 11l2.313.938L9.531 13.5l-3.906 1.563L3.312 16l2.313.938L9.531 18.5l-3.906 1.563L3.312 21l2.313.938l10 4l.375.125l.375-.125l10-4L28.688 21l-2.313-.938l-3.906-1.562l3.906-1.563L28.688 16l-2.313-.938l-3.906-1.562l3.906-1.563L28.688 11l-2.313-.938l-10-4zm0 2.156L23.281 11L16 13.906L8.719 11zm-3.75 6.5l3.375 1.344l.375.124l.375-.125l3.375-1.343L23.281 16L16 18.906L8.719 16zm0 5l3.375 1.343l.375.125l.375-.125l3.375-1.343L23.281 21L16 23.906L8.719 21z"/>`),
		g.Group(children),
	)
}