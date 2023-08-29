package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MallBag(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSMallBag0"><g fill="none" stroke-linejoin="round" stroke-width="4"><path fill="#fff" stroke="#fff" d="M6 12.6V41a2 2 0 0 0 2 2h32a2 2 0 0 0 2-2V12.6H6Z"/><path stroke="#fff" stroke-linecap="round" d="M42 12.6L36.333 5H11.667L6 12.6v0"/><path stroke="#000" stroke-linecap="round" d="M31.555 19.2c0 4.198-3.382 7.6-7.555 7.6s-7.556-3.402-7.556-7.6"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSMallBag0)"/>`),
		g.Group(children),
	)
}