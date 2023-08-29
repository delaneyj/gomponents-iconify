package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowCounterClockwiseDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<g fill="currentColor"><path d="M72 104H24V56Z" opacity=".2"/><path d="M195.88 60.08A96.08 96.08 0 0 0 60.25 60L49.31 70L29.66 50.3A8 8 0 0 0 16 56v48a8 8 0 0 0 8 8h48a8 8 0 0 0 5.66-13.66l-17-17l10.54-9.65a3.07 3.07 0 0 0 .26-.25a80 80 0 1 1 1.65 114.78a8 8 0 0 0-11 11.63A95.38 95.38 0 0 0 128 224h1.32a96 96 0 0 0 66.56-163.92ZM32 96V75.28L52.69 96Z"/></g>`),
		g.Group(children),
	)
}