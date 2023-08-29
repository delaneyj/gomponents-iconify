package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FishOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTFishOne0"><g fill="none"><path fill="#555" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M44 24c-1.215 4.69-7.962 8.467-11 9c-2.43 5.97-8.962 6.533-12 6l4-6c-4.456-.427-9.975-5.046-12-7c-2.614 2.85-6.806 5.08-9 5.969C7.646 24.294 5.519 17.309 4 15c2.835 0 7.143 3.224 9 5c2.025-2.132 8.962-5.112 12-6l-4-5c7.696-.853 11.156 2.868 12 5c7.696 1.706 10.662 7.69 11 10Z"/><circle cx="36" cy="24" r="2" fill="#fff"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTFishOne0)"/>`),
		g.Group(children),
	)
}