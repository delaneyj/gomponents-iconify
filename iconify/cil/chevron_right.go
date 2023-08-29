package cil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChevronRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="m183.505 496l-97.268-97.27l143.175-143.174L86.237 112.38l97.268-97.27l143.227 143.228l11.316 11.209l85.9 85.9l.051.05l-11.311 11.419l-85.9 85.9l-.055-.054Zm-52.013-97.27l52.013 52.014L326.629 307.62l.055.054l52.116-52.118l-52.127-52.128l-11.308-11.2l-131.86-131.862l-52.013 52.014l143.175 143.176Z"/>`),
		g.Group(children),
	)
}