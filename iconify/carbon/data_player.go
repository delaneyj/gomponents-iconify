package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DataPlayer(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M28 24v2H15.816a2.983 2.983 0 0 0-5.632 0H4v-2H2v6h2v-2h6.184a2.983 2.983 0 0 0 5.632 0H28v2h2v-6zM13 7.5v8.999L21 12l-8-4.5z"/><path fill="currentColor" d="M16 22a10 10 0 1 1 10-10a10.011 10.011 0 0 1-10 10Zm0-18a8 8 0 1 0 8 8a8.01 8.01 0 0 0-8-8Z"/>`),
		g.Group(children),
	)
}