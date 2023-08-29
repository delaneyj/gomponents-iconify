package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Trunk(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTTrunk0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><rect width="24" height="28" x="12" y="12" fill="#555" rx="4"/><path d="M20 12V6m8 6V6M16 4h16M18 40v4m12-4v4M20 25h8"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTTrunk0)"/>`),
		g.Group(children),
	)
}