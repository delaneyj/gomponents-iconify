package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MagnifierTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<g fill="currentColor" fill-rule="evenodd"><path d="M16.025 7.5c0-4.143-3.356-7.5-7.499-7.5a7.499 7.499 0 0 0-7.499 7.5a7.5 7.5 0 0 0 7.5 7.5c2.219 0 7.5-.052 7.5-.052l-.002-7.449zm-7.553 5.529a5.506 5.506 0 1 1 .002-11.012a5.506 5.506 0 0 1-.002 11.012zm6.487.929h-1v-1h1v1z"/><path d="M7.844 3.044c-2.119 0-3.839 1.616-3.839 3.608c0 .25.026.496.077.73c.186.84.529.691.529-.158c0-1.998 1.719-3.609 3.84-3.609c.905 0 .608-.571-.607-.571z"/></g>`),
		g.Group(children),
	)
}