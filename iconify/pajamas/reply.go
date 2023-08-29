package pajamas

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Reply(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M6.84 9.72a.75.75 0 1 1-1.06 1.06L2.53 7.53L2 7l.53-.53l3.25-3.25a.75.75 0 0 1 1.06 1.06L4.87 6.25h4.378a4.75 4.75 0 0 1 4.75 4.75v3.25a.75.75 0 0 1-1.5 0V11a3.25 3.25 0 0 0-3.25-3.25H4.871l1.97 1.97Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}