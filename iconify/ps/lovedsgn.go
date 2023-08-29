package ps

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Lovedsgn(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M334 2q-44-1-69 12q-32 16-56 72q-1-3-4-7q-11-20-17-29t-16.5-20T148 13Q118-2 75 4Q17 18 8 82q-11 41 23 92q24 37 28 43q7 10 35.5 47t38.5 52q49 77 61 107q1 2 7.5 19.5T209 462q20-88 118-204q114-131 73-205q-17-38-66-51z"/>`),
		g.Group(children),
	)
}