package pajamas

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ExpandRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M14.5 12.25a.75.75 0 0 0 1.5 0v-8.5a.75.75 0 0 0-1.5 0v8.5ZM4.841 4.22a.75.75 0 0 1 0 1.06l-1.97 1.97h9.379a.75.75 0 0 1 0 1.5H2.871l1.97 1.97a.75.75 0 1 1-1.06 1.06L.53 8.53L0 8l.53-.53l3.25-3.25a.75.75 0 0 1 1.061 0Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}