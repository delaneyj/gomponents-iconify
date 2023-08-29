package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MovieBoard(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTMovieBoard0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path fill="#555" d="M44 16H4v26h40V16Z"/><path d="M44 16V6H4v10h40ZM26 6l4 10M18 6l4 10M10 6l4 10M34 6l4 10m-26 8h24m-24 8h12"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTMovieBoard0)"/>`),
		g.Group(children),
	)
}