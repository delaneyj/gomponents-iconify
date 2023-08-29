package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WbSunnyOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11 4V1h2v3h-2Zm0 19v-3h2v3h-2Zm9-10v-2h3v2h-3ZM1 13v-2h3v2H1Zm17.7-6.3l-1.4-1.4l1.75-1.8l1.45 1.45l-1.8 1.75ZM4.95 20.5L3.5 19.05l1.8-1.75l1.4 1.4l-1.75 1.8Zm14.1 0l-1.75-1.8l1.4-1.4l1.8 1.75l-1.45 1.45ZM5.3 6.7L3.5 4.95L4.95 3.5L6.7 5.3L5.3 6.7ZM12 18q-2.5 0-4.25-1.75T6 12q0-2.5 1.75-4.25T12 6q2.5 0 4.25 1.75T18 12q0 2.5-1.75 4.25T12 18Zm0-2q1.675 0 2.838-1.163T16 12q0-1.675-1.163-2.838T12 8q-1.675 0-2.838 1.163T8 12q0 1.675 1.163 2.838T12 16Zm0-4Z"/>`),
		g.Group(children),
	)
}