package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DeathStar(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTDeathStar0"><g fill="none" stroke="#fff" stroke-width="4"><circle cx="24" cy="24" r="20" fill="#555"/><path stroke-linecap="round" d="M5 25h39"/><circle cx="19" cy="15" r="4" fill="#555"/><path stroke-linecap="round" d="M31 31h12m-9-14h8M25 37h14"/><path d="M40 11.998A19.91 19.91 0 0 1 44 24c0 7.808-4.475 14.572-11 17.865M4.4 20c-.262 1.292-.4 2.63-.4 4c0 1.727.219 3.402.63 5"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTDeathStar0)"/>`),
		g.Group(children),
	)
}