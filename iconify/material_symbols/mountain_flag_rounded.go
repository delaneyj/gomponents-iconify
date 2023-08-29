package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MountainFlagRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m11.55 13.1l-1.475.75l-3.05-2.025l.825-1.7q.25-.5.738-.813T9.65 9H11V3q0-.425.288-.713T12 2h5.2q.275 0 .425.238t.025.487l-.55 1.05q-.05.125-.05.225t.05.225l.55 1.05q.125.25-.025.487T17.2 6H13v3h1.25q.575 0 1.05.3t.75.8l.875 1.75l-3 2l-1.475-.75q-.2-.125-.45-.125t-.45.125ZM2.7 20.55l3.425-6.925l3.325 2.2q.225.15.488.175t.512-.1l1.55-.775l1.55.775q.25.125.513.113t.487-.163l3.275-2.175l3.45 6.875q.25.5-.037.975t-.863.475H3.6q-.575 0-.862-.475T2.7 20.55Z"/>`),
		g.Group(children),
	)
}