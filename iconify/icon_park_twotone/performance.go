package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Performance(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTPerformance0"><g fill="none" stroke="#fff" stroke-linejoin="round" stroke-width="4"><path stroke-linecap="round" d="M24 4C12.954 4 4 12.954 4 24s8.954 20 20 20s20-8.954 20-20a19.94 19.94 0 0 0-2.358-9.43"/><path fill="#555" d="M35 10c2.21 0 4-1.343 4-3s-1.79-3-4-3s-4 1.343-4 3s1.79 3 4 3ZM24 31a7 7 0 1 0 0-14a7 7 0 0 0 0 14Z"/><path stroke-linecap="round" d="M31 6.5V24"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTPerformance0)"/>`),
		g.Group(children),
	)
}