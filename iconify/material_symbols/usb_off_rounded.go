package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UsbOffRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 22q-.825 0-1.413-.588T10 20q0-.525.275-.975T11 18.3V16H8q-.825 0-1.413-.588T6 14v-2.3q-.45-.225-.725-.675T5 10q0-.45.175-.825t.5-.65l-3.6-3.6Q1.8 4.65 1.8 4.238t.3-.713q.3-.3.713-.3t.712.3L20.5 20.5q.275.275.288.688t-.288.712q-.3.3-.713.3t-.712-.3L13.15 16H13v2.3q.475.25.738.7T14 20q0 .825-.588 1.413T12 22Zm5.775-7.075L16 13.15V12h-.5q-.2 0-.35-.15T15 11.5v-3q0-.2.15-.35T15.5 8h3q.2 0 .35.15t.15.35v3q0 .2-.15.35t-.35.15H18v2q0 .25-.05.488t-.175.437ZM8 14h3v-.15l-2.525-2.525q-.1.125-.225.213T8 11.7V14Zm5-3.85l-2-2V6h-1q-.3 0-.45-.275T9.6 5.2l2-2.675q.15-.2.4-.2t.4.2l2 2.675q.2.25.05.525T14 6h-1v4.15Z"/>`),
		g.Group(children),
	)
}