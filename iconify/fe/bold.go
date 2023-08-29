package fe

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g id="feBold0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feBold1" fill="currentColor"><path id="feBold2" d="M5 20V4c0-1.1.9-2 2-2h5c3 0 5.966 1.4 6 4.919c0 1.838-.931 3.73-2.532 4.324v.135c2.033.514 3.532 2.027 3.532 4.73c0 1.022-.203 1.905-.573 2.653C17.337 20.96 15.09 22 12 22H7c-1.1 0-2-.9-2-2Zm3-1h4c2.566 0 4-1.2 4-3s-1.408-3-4-3H8v6Zm0-9h3c2.388 0 4-.85 4-2.55C15 5.75 13.5 5 10.996 5H8v5Z"/></g></g>`),
		g.Group(children),
	)
}