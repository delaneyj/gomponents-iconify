package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bib(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSBib0"><g fill="none" stroke-linejoin="round" stroke-width="4"><path fill="#fff" stroke="#fff" d="M31 14c0 8-14 8-14 0c0-5 5-6 3-9S8 7 8 16v16c0 8.5 8.5 12 15.5 12S40 41 40 32V16C40 7 29 2 27 5s4 4 4 9Z"/><path fill="#000" stroke="#000" stroke-linecap="round" d="m19 32l5-5l5 5l-5 5l-5-5Z"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSBib0)"/>`),
		g.Group(children),
	)
}