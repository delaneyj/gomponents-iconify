package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FemurAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M13.5 13q.425 0 .713-.288T14.5 12q0-.15-.15-.575q.275-.1.463-.35T15 10.5q0-.425-.288-.712T14 9.5q-.325 0-.575.175t-.35.475L9.425 8.4q.025-.075.075-.4q0-.425-.288-.713T8.5 7q-.425 0-.713.288T7.5 8q0 .15.175.55q-.275.1-.475.35t-.2.6q0 .425.288.713T8 10.5q.35 0 .6-.188t.35-.487l3.65 1.75l-.1.425q0 .425.288.713T13.5 13ZM2.975 11.5L9.05 3.025l6.425 2.925q1.6.725 2.563 2.2T19 11.425V22H9.85q-.4-1.175-.6-2.313T9 17.538q-.05-1.012-.012-1.862t.112-1.45q-.025 0 0 0q-.55-.125-1.262-.312T6.313 13.4q-.813-.325-1.663-.787T2.975 11.5Z"/>`),
		g.Group(children),
	)
}