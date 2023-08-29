package fe

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BirthdayCake(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g id="feBirthdayCake0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feBirthdayCake1" fill="currentColor" fill-rule="nonzero"><path id="feBirthdayCake2" d="M17 10a4 4 0 0 1 4 4v5a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2v-5a4 4 0 0 1 4-4V7h2v3h2V7h2v3h2V7h2v3ZM7 12a2 2 0 0 0-2 2v1h14v-1a2 2 0 0 0-2-2H7Zm-2 5v2h14v-2H5ZM7 4a1 1 0 0 0 1-1a1 1 0 0 1 1 1v1a1 1 0 1 1-2 0V4Zm4 0a1 1 0 0 0 1-1a1 1 0 0 1 1 1v1a1 1 0 0 1-2 0V4Zm4 0a1 1 0 0 0 1-1a1 1 0 0 1 1 1v1a1 1 0 0 1-2 0V4Z"/></g></g>`),
		g.Group(children),
	)
}