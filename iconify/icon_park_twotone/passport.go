package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Passport(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTPassport0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path fill="#555" d="M10 10h28v34H10V10Z"/><path d="m10 10l22-6v6"/><circle cx="24" cy="24" r="4" fill="#555"/><path d="M20 34h8"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTPassport0)"/>`),
		g.Group(children),
	)
}