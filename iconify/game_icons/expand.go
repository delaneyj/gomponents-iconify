package game_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Expand(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M64 64v128l48-48l48 48l32-32l-48-48l48-48H64zm256 0l48 48l-48 48l32 32l48-48l48 48V64H320zM64 320v128h128l-48-48l48-48l-32-32l-48 48l-48-48zm288 0l-32 32l48 48l-48 48h128V320l-48 48l-48-48z"/>`),
		g.Group(children),
	)
}