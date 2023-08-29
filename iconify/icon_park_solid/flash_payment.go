package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlashPayment(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSFlashPayment0"><g fill="none" stroke-linecap="round" stroke-width="4"><path fill="#fff" stroke="#fff" stroke-linejoin="round" d="M31 4H16l-6 23h8l-4 17l26-28H28l3-12Z"/><path stroke="#000" d="m21 11l-2 8"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSFlashPayment0)"/>`),
		g.Group(children),
	)
}