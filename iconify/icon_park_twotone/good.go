package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Good(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTGood0"><path fill="#555" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M8.006 19.197C12.443 10.325 22 7 28.506 4.197c2.958-1.274 3.69 3.526 2.5 6.5c-1 2.5-3 5.303-3 5.303h8.5a3.5 3.5 0 1 1 0 7h2a3.5 3.5 0 1 1 0 7h-4a3.5 3.5 0 1 1 0 7h-3a3.5 3.5 0 0 1 .002 7H19.006c-3.5 0-8-1.803-11-6.803c-2.875-4.793-3-12 0-18Z"/></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTGood0)"/>`),
		g.Group(children),
	)
}