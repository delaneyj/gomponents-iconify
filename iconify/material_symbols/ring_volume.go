package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RingVolume(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m3.4 20.4l-2.3-2.25q-.3-.3-.3-.7t.3-.7q2.2-2.375 5.075-3.562T12 12q2.95 0 5.813 1.188T22.9 16.75q.3.3.3.7t-.3.7l-2.3 2.25q-.275.275-.638.3t-.662-.2l-2.9-2.2q-.2-.15-.3-.35t-.1-.45v-2.85q-.95-.3-1.95-.475T12 14q-1.05 0-2.05.175T8 14.65v2.85q0 .25-.1.45t-.3.35l-2.9 2.2q-.3.225-.663.2t-.637-.3ZM11 7V2h2v5h-2Zm6.6 2.85L16.2 8.4l3.55-3.55l1.4 1.45l-3.55 3.55Zm-11.2 0L2.85 6.3l1.4-1.45L7.8 8.4L6.4 9.85Z"/>`),
		g.Group(children),
	)
}