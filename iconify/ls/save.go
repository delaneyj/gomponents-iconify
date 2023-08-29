package ls

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Save(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 717 717"),
		g.Raw(`<path fill="currentColor" d="M390 0v149h78L312 337L156 149h82V0h152zm132 241l91 144c5 12 11 35 11 49v277c0 15-12 26-26 26H27c-15 0-27-11-27-26V434c0-14 5-37 10-49l91-144c5-13 23-24 36-24h41l42 51h-84L55 416c-1 0-1 1-1 2c0 2 0 3-1 4h517v-4c0-1-1-2-1-3l-81-147h-84l42-51h40c14 0 31 11 36 24zM239 530h147c13 0 25-12 25-26c0-13-12-25-25-25H239c-14 0-26 12-26 25c0 14 12 26 26 26z"/>`),
		g.Group(children),
	)
}