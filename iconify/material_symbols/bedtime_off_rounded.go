package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BedtimeOffRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11.075 8.25L6.5 3.65q.975-.65 1.938-1t2.012-.5q.725-.1 1.05.35t.05 1.25q-.4 1.125-.512 2.25t.037 2.25Zm8 13.65L17.5 20.325q-1.2.8-2.575 1.238T12 22q-2.075 0-3.9-.788t-3.175-2.137q-1.35-1.35-2.137-3.175T2 12q0-1.525.425-2.913T3.65 6.5L2.1 4.925q-.275-.275-.275-.688t.275-.712q.3-.3.713-.3t.712.3L20.5 20.5q.3.3.288.7t-.313.7q-.3.275-.7.288t-.7-.288Z"/>`),
		g.Group(children),
	)
}