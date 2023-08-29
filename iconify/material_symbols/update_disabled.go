package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UpdateDisabled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m19.775 22.625l-3-3q-1.05.65-2.25 1.012T12 21q-1.875 0-3.513-.713t-2.85-1.924q-1.212-1.213-1.924-2.85T3 12q0-1.325.363-2.525t1.012-2.25l-3-3L2.8 2.8l18.4 18.4l-1.425 1.425ZM12 19q.9 0 1.738-.213t1.587-.612l-9.5-9.5q-.4.75-.612 1.588T5 12q0 2.925 2.038 4.963T12 19Zm3-9V8h2.75q-1.025-1.4-2.525-2.2T12 5q-.9 0-1.738.213t-1.587.612l-1.45-1.45q1.05-.65 2.25-1.012T12 3q2.05 0 3.888.875T19 6.35V4h2v6h-6Zm-2 .15l-2-2V7h2v3.15Zm6.625 6.625l-1.45-1.45q.275-.55.463-1.125T18.9 13h2.05q-.125 1.05-.463 2t-.862 1.775Z"/>`),
		g.Group(children),
	)
}