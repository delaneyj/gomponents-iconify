package game_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SleevelessTop(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="m320 32l32 16c0 32 16 48 32 96c0 0 16 32.7 16 48c0 48-16 96-32 144s32 96 48 160H96c16-64 64-112 48-160s-32-96-32-144c0-16 16-48 16-48c16-48 32-64 32-96l32-16c0 64 32 144 64 144s64-80 64-144z"/>`),
		g.Group(children),
	)
}