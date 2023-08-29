package covid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VaccineProtectionFaceMaskOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M12 4.563a3 3 0 0 0-1.8.6c-.779.584-2.726 2.15-3.7 2.15L5 7.438v9l3.285 1.971A7.223 7.223 0 0 0 12 19.438m-7-3l-.357-.168A6.782 6.782 0 0 1 .75 10.135A2.981 2.981 0 0 1 5 7.438m7-2.875a3 3 0 0 1 1.8.6c.779.584 2.726 2.15 3.7 2.15l1.5.125v9l-3.285 1.971A7.223 7.223 0 0 1 12 19.438m7-3l.357-.168a6.782 6.782 0 0 0 3.893-6.135A2.981 2.981 0 0 0 19 7.438m-9.5 4.75h5m-2.5 2.5v-5"/>`),
		g.Group(children),
	)
}