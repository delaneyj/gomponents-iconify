package game_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Contract(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M96 64L64 96l48 48l-48 48h128V64l-48 48l-48-48zm224 0v128h128l-48-48l48-48l-32-32l-48 48l-48-48zM64 320l48 48l-48 48l32 32l48-48l48 48V320H64zm256 0v128l48-48l48 48l32-32l-48-48l48-48H320z"/>`),
		g.Group(children),
	)
}