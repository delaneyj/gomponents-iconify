package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CosmeticBrush(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTCosmeticBrush0"><g fill="none" stroke="#fff" stroke-linejoin="round" stroke-width="4"><path fill="#555" d="m16.1 25.829l22.627-19.8s2.12-2.12 4.242 0c2.121 2.122 0 4.243 0 4.243L23.17 32.9l-7.07-7.071Z"/><path stroke-linecap="round" d="m22.464 20.879l5.657 5.657"/><path d="m5.493 30.778l10.607-4.95l7.07 7.072l-4.95 10.606s-4.95.707-9.192-3.535c-4.243-4.243-3.536-9.193-3.536-9.193Z"/><path stroke-linecap="round" d="m6.908 36.435l4.95-2.121m.707 7.778l1.415-2.828"/><path d="m18.93 23.354l2.828-2.475l2.828-2.475m6.01 6.01l-2.475 2.828l-2.475 2.829"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTCosmeticBrush0)"/>`),
		g.Group(children),
	)
}