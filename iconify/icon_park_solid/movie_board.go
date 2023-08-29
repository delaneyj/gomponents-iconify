package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MovieBoard(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSMovieBoard0"><g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path fill="#fff" stroke="#fff" d="M44 16H4v26h40V16Z"/><path stroke="#fff" d="M44 16V6H4v10h40ZM26 6l4 10M18 6l4 10M10 6l4 10M34 6l4 10"/><path stroke="#000" d="M12 24h24m-24 8h12"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSMovieBoard0)"/>`),
		g.Group(children),
	)
}