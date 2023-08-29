package icomoon_free

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Pilcrow(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M6 0h8v2h-2v14h-2V2H8v14H6V8a4 4 0 0 1 0-8z"/>`),
		g.Group(children),
	)
}