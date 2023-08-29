package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CouchBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M244 104V72a20 20 0 0 0-20-20H32a20 20 0 0 0-20 20v32a20 20 0 0 0-8 16v48a20 20 0 0 0 20 20h4v12a12 12 0 0 0 24 0v-12h152v12a12 12 0 0 0 24 0v-12h4a20 20 0 0 0 20-20v-48a20 20 0 0 0-8-16Zm-24-4h-12a20 20 0 0 0-20 20v4h-48V76h80ZM116 76v48H68v-4a20 20 0 0 0-20-20H36V76Zm112 88H28v-40h16v4a20 20 0 0 0 20 20h128a20 20 0 0 0 20-20v-4h16Z"/>`),
		g.Group(children),
	)
}