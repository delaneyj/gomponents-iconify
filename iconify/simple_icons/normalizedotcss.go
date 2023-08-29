package simple_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Normalizedotcss(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8.427 6.361v5.064l-5.381.705l7.058.924v-1.915l5.469 6.5v-5.064l5.382-.705l-7.059-.924v1.914zM12 0l12 12l-12 12L0 12Z"/>`),
		g.Group(children),
	)
}