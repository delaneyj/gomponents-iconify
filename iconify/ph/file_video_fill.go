package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FileVideoFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M148.23 153.21a8 8 0 0 0-7.75-.39l-21.22 10.39A16 16 0 0 0 104 152H48a16 16 0 0 0-16 16v40a16 16 0 0 0 16 16h56a16 16 0 0 0 15.44-11.87l20.84 11A8 8 0 0 0 152 216v-56a8 8 0 0 0-3.77-6.79ZM136 202.76l-16-8.42v-13.68l16-7.83Zm77.66-120.42l-56-56A8 8 0 0 0 152 24H56a16 16 0 0 0-16 16v88a8 8 0 0 0 16 0V40h88v48a8 8 0 0 0 8 8h48v120h-24a8 8 0 0 0 0 16h24a16 16 0 0 0 16-16V88a8 8 0 0 0-2.34-5.66ZM160 51.31L188.69 80H160Z"/>`),
		g.Group(children),
	)
}