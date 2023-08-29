package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Conditioner(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTConditioner0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path fill="#555" d="M8 24h32v20H8z"/><path d="M37 24v-7H11v7m6-7c2-2.167 3-6.5 3-13c3 0 10 5.417 10 12.733M16 31h16v6H16z"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTConditioner0)"/>`),
		g.Group(children),
	)
}