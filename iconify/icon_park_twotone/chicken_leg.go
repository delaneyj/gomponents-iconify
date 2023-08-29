package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChickenLeg(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTChickenLeg0"><g fill="none"><path fill="#555" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M33.375 33.874c4.242-4.242 1.414-18.384-4.95-24.748c-2.828-2.829-10.96-8.84-19.799 0c-8.839 8.838-2.828 16.97 0 19.799c6.364 6.364 20.506 9.192 24.749 4.95Z"/><path stroke="#fff" stroke-width="4" d="m41 41l-7-7"/><circle cx="42.193" cy="40.071" r="2.5" fill="#fff" transform="rotate(135 42.193 40.071)"/><circle cx="40.072" cy="42.192" r="2.5" fill="#fff" transform="rotate(135 40.072 42.192)"/><circle cx="17" cy="18" r="2" fill="#fff"/><circle cx="12" cy="21" r="2" fill="#fff"/><circle cx="17" cy="24" r="2" fill="#fff"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTChickenLeg0)"/>`),
		g.Group(children),
	)
}