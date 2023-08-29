package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Weight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSWeight0"><g fill="none"><path fill="#fff" stroke="#fff" stroke-linejoin="round" stroke-width="4" d="M41 4H7a3 3 0 0 0-3 3v34a3 3 0 0 0 3 3h34a3 3 0 0 0 3-3V7a3 3 0 0 0-3-3Z"/><path stroke="#000" stroke-linecap="round" stroke-width="4" d="M12 19.054c3.325-4 7.325-6 12-6s8.675 2 12 6"/><path fill="#000" d="M24 31a3 3 0 1 0 0-6a3 3 0 0 0 0 6Z"/><path stroke="#000" stroke-linecap="round" stroke-width="4" d="m19 21l5.008 7"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSWeight0)"/>`),
		g.Group(children),
	)
}