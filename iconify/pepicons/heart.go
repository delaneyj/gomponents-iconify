package pepicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Heart(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M10 5.722c1.69-3.023 7.5-1.968 7.5 2.4c0 2.918-2.5 5.582-7.5 7.993c-5-2.41-7.5-5.075-7.5-7.993c0-4.368 5.81-5.423 7.5-2.4Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}