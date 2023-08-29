package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Leaves(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M31 43C31 43 18 44 11 36C4 28 4 4 4 4C4 4 28 3 36 9C44 15 42 32 42 32"/><path d="M44 44C44 44 32.8207 35.5515 26 28C19.1793 20.4485 16 13 16 13"/><path d="M26 28L27 15"/><path d="M26 28L16 27"/></g>`),
		g.Group(children),
	)
}