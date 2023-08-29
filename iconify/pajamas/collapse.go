package pajamas

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Collapse(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M11.78 3.16a.75.75 0 0 0-1.06 0L8.75 5.13V.75a.75.75 0 0 0-1.5 0v4.38L5.28 3.16a.75.75 0 0 0-1.06 1.06l3.25 3.25L8 8l-.53.53l-3.25 3.25a.75.75 0 0 0 1.06 1.061l1.97-1.97v4.379a.75.75 0 0 0 1.5 0v-4.379l1.97 1.97a.75.75 0 1 0 1.06-1.06L8.53 8.53L8 8l.53-.53l3.25-3.25a.75.75 0 0 0 0-1.061Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}