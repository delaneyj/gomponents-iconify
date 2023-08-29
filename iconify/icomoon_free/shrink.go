package icomoon_free

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Shrink(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M9 7h6.5L13 4.5l3-3L14.5 0l-3 3L9 .5zm0 2v6.5l2.5-2.5l3 3l1.5-1.5l-3-3L15.5 9zM7 9H.5L3 11.5l-3 3L1.5 16l3-3L7 15.5zm0-2V.5L4.5 3l-3-3L0 1.5l3 3L.5 7z"/>`),
		g.Group(children),
	)
}