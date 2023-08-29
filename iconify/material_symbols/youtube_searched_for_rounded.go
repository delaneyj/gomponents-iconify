package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func YoutubeSearchedForRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11 16q-.6 0-1.163-.1t-1.087-.3q-.4-.15-.55-.55t.05-.8q.2-.375.575-.488t.8.013q.325.125.675.175t.7.05q1.875 0 3.188-1.313T15.5 9.5q0-1.875-1.313-3.188T11 5Q9.275 5 8.012 6.163T6.55 9.05l.55-.55q.275-.275.7-.275t.7.275q.275.275.275.7t-.275.7l-2.3 2.3q-.3.3-.7.3t-.7-.3L2.5 9.9q-.275-.275-.275-.7t.3-.725Q2.8 8.2 3.2 8.2t.7.275l.65.625q.15-2.575 2-4.338T11 3q2.725 0 4.612 1.888T17.5 9.5q0 1.05-.325 2.05T16.2 13.3l5.05 5.05q.275.275.275.7t-.275.7q-.275.275-.7.275t-.7-.275L14.8 14.7q-.8.65-1.775.975T11 16Z"/>`),
		g.Group(children),
	)
}