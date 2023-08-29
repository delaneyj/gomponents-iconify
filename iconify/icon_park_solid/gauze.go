package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Gauze(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSGauze0"><g fill="none" stroke-width="4"><circle cx="26" cy="24" r="17" fill="#fff" stroke="#fff"/><circle cx="26" cy="24" r="7" fill="#000" stroke="#000"/><path stroke="#fff" stroke-linecap="round" stroke-linejoin="round" d="M5 41h21"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSGauze0)"/>`),
		g.Group(children),
	)
}