package pajamas

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SearchResults(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M1.75 1a.75.75 0 0 0 0 1.5h8.5a.75.75 0 0 0 0-1.5h-8.5ZM1 4.75A.75.75 0 0 1 1.75 4H7a.75.75 0 0 1 0 1.5H1.75A.75.75 0 0 1 1 4.75Zm9 7.75a2.5 2.5 0 1 0 0-5a2.5 2.5 0 0 0 0 5Zm0 1.5c.834 0 1.607-.255 2.248-.691l1.472 1.471a.75.75 0 1 0 1.06-1.06l-1.471-1.472A4 4 0 1 0 10 14ZM1.75 7a.75.75 0 0 0 0 1.5H4A.75.75 0 0 0 4 7H1.75Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}