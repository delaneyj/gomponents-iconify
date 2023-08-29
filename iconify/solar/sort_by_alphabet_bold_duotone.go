package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SortByAlphabetBoldDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path fill-rule="evenodd" d="M17.184 6.692a.75.75 0 0 0-1.368 0l-4.5 10a.75.75 0 0 0 1.368.616l1.438-3.194h4.757l1.437 3.194a.75.75 0 0 0 1.368-.616l-4.5-10ZM16.5 8.828l-1.704 3.786h3.408L16.5 8.828Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M2.25 7A.75.75 0 0 1 3 6.25h10a.75.75 0 0 1 0 1.5H3A.75.75 0 0 1 2.25 7Z" clip-rule="evenodd" opacity=".5"/><path d="M2.25 12a.75.75 0 0 1 .75-.75h7a.75.75 0 0 1 0 1.5H3a.75.75 0 0 1-.75-.75Zm0 5a.75.75 0 0 1 .75-.75h5a.75.75 0 0 1 0 1.5H3a.75.75 0 0 1-.75-.75Z" opacity=".5"/></g>`),
		g.Group(children),
	)
}