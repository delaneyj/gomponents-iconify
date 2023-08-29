package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MountainFlagOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m8.7 12.925l1.375.925l1.475-.75q.2-.125.45-.125t.45.125l1.475.75l1.325-.875l-1-1.975h-4.6l-.95 1.925ZM5.225 20H18.75l-2.6-5.225l-1.6 1.075q-.225.15-.487.163t-.513-.113L12 15.125l-1.55.775q-.25.125-.513.1t-.487-.175L7.8 14.75L5.225 20ZM3.6 22q-.575 0-.862-.475T2.7 20.55l5.15-10.425q.25-.5.738-.813T9.65 9H11V3q0-.425.288-.713T12 2h5.2q.275 0 .425.238t.025.487l-.55 1.05q-.05.125-.05.225t.05.225l.55 1.05q.125.25-.025.487T17.2 6H13v3h1.25q.575 0 1.05.3t.75.8l5.225 10.45q.25.5-.037.975t-.863.475H3.6Zm8.4-6.875Z"/>`),
		g.Group(children),
	)
}