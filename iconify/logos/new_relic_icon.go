package logos

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NewRelicIcon(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="#00AC69" d="M206.745 102.322v90.942L128 238.745v56.841l128-73.892V73.892z"/><path fill="#1CE783" d="m128 56.86l78.745 45.462L256 73.892L128 0L0 73.892l49.236 28.43z"/><path fill="#1D252C" d="M78.764 176.232v90.943L128 295.586V147.802L0 73.892v56.86z"/>`),
		g.Group(children),
	)
}