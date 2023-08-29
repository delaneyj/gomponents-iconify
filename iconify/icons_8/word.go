package icons_8

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Word(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M6 3v26h20V3H6zm2 2h16v22H8V5zm10 7v6.5c0 .217-.283.5-.5.5c-.042 0 .02.048-.063-.063c-.082-.11-.206-.388-.28-.687C17.006 17.652 17 17 17 17v-2h-2v4.5c0 .217-.283.5-.5.5c-.217 0-.5-.283-.5-.5V13h-4v2h2v4.5c0 1.383 1.117 2.5 2.5 2.5c.984 0 1.686-.644 2.094-1.47c.302.187.52.47.906.47c1.383 0 2.5-1.117 2.5-2.5V14h2v-2h-4z"/>`),
		g.Group(children),
	)
}