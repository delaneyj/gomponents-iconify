package radix_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Pause(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M6.05 2.75a.55.55 0 0 0-1.1 0v9.5a.55.55 0 0 0 1.1 0v-9.5Zm4 0a.55.55 0 0 0-1.1 0v9.5a.55.55 0 0 0 1.1 0v-9.5Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}