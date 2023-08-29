package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M13 18a9 9 0 0 1 9-9h4a9.003 9.003 0 0 1 8.487 6a3 3 0 1 1-5.657 2A3.003 3.003 0 0 0 26 15h-4a3 3 0 1 0 0 6h4a9 9 0 1 1 0 18h-4a9.003 9.003 0 0 1-8.487-6a3 3 0 1 1 5.657-2A3.003 3.003 0 0 0 22 33h4a3 3 0 1 0 0-6h-4a9 9 0 0 1-9-9Zm9-7a7 7 0 1 0 0 14h4a5 5 0 0 1 0 10h-4a5.002 5.002 0 0 1-4.716-3.333a1 1 0 1 0-1.885.666A7.002 7.002 0 0 0 22 37h4a7 7 0 1 0 0-14h-4a5 5 0 0 1 0-10h4a5.002 5.002 0 0 1 4.716 3.333a1 1 0 1 0 1.885-.666A7.002 7.002 0 0 0 26 11h-4Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}