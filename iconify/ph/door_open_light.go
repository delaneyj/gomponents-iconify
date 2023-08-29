package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DoorOpenLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M232 218h-26V40a14 14 0 0 0-14-14H64a14 14 0 0 0-14 14v178H24a6 6 0 0 0 0 12h208a6 6 0 0 0 0-12ZM194 40v178h-20V40a14.71 14.71 0 0 0-.16-2H192a2 2 0 0 1 2 2ZM62 40a2 2 0 0 1 2-2h96a2 2 0 0 1 2 2v178H62Zm82 92a12 12 0 1 1-12-12a12 12 0 0 1 12 12Z"/>`),
		g.Group(children),
	)
}