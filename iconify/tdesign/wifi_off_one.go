package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WifiOffOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 1.586L22.414 21l-1.415 1.414l-5.361-5.361l-3.639 4.548L.6 7.351l.779-.626c.747-.6 1.534-1.127 2.35-1.582L1.586 3L3 1.586Zm2.207 5.036c-.611.31-1.205.665-1.778 1.064l8.57 10.713l2.216-2.77l-9.008-9.007Zm15.362 1.064A15.002 15.002 0 0 0 9.41 5.224l-.985.172l-.343-1.97l.985-.173a16.99 16.99 0 0 1 13.555 3.472l.779.625l-5.467 6.833l-1.562-1.25l4.198-5.247Z"/>`),
		g.Group(children),
	)
}