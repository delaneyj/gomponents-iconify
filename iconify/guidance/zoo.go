package guidance

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Zoo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" d="M6 3h-.5a5 5 0 0 0-5 5v13h3l1.072-2.679a3.693 3.693 0 0 1 6.856 0L12.5 21h3v-6.5a1.5 1.5 0 1 1 3 0v4A2.5 2.5 0 0 0 21 21h2.5v-3h-2V8a5 5 0 0 0-5-5h-8v6a2.5 2.5 0 0 0 2.5 2.5h.5l.867-2.603A2.775 2.775 0 0 1 15 7m3.5 2.5v-2"/>`),
		g.Group(children),
	)
}