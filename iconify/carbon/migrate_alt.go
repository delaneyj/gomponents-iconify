package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MigrateAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M28 4H4a2 2 0 0 0-2 2v20a2 2 0 0 0 2 2h24a2 2 0 0 0 2-2V6a2 2 0 0 0-2-2Zm0 22H12v-6h-2v6H4v-9h16.172l-3.586 3.586L18 22l6-6l-6-6l-1.414 1.414L20.172 15H4V6h6v6h2V6h16Z"/>`),
		g.Group(children),
	)
}