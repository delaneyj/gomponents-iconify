package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CameraSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="M2 1h5v1H2V1Zm6 7.5a1.5 1.5 0 1 1 3 0a1.5 1.5 0 0 1-3 0Z"/><path fill="currentColor" fill-rule="evenodd" d="M0 12.5A1.5 1.5 0 0 0 1.5 14h12a1.5 1.5 0 0 0 1.5-1.5v-8A1.5 1.5 0 0 0 13.5 3h-12A1.5 1.5 0 0 0 0 4.5v8ZM9.5 6a2.5 2.5 0 1 0 0 5a2.5 2.5 0 0 0 0-5Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}