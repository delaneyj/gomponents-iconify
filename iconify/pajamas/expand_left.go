package pajamas

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ExpandLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M0 12.25a.75.75 0 0 0 1.5 0v-8.5a.75.75 0 0 0-1.5 0v8.5Zm11.159-8.03a.75.75 0 0 0 0 1.06l1.97 1.97H3.75a.75.75 0 1 0 0 1.5h9.379l-1.97 1.97a.75.75 0 1 0 1.06 1.06l3.25-3.25L16 8l-.53-.53l-3.25-3.25a.75.75 0 0 0-1.061 0Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}