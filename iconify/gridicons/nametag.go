package gridicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Nametag(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 6a1 1 0 1 1-1 1a1 1 0 0 1 1-1zm-6 8h12v3H6zm14-8h-4V3H8v3H4a2 2 0 0 0-2 2v11a2 2 0 0 0 2 2h16a2 2 0 0 0 2-2V8a2 2 0 0 0-2-2zM10 5h4v5h-4zm10 14H4v-9h4a2 2 0 0 0 2 2h4a2 2 0 0 0 2-2h4z"/>`),
		g.Group(children),
	)
}