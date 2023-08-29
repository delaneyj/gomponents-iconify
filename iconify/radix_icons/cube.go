package radix_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Cube(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M7.289.797a.5.5 0 0 1 .422 0l6 2.8A.5.5 0 0 1 14 4.05v6.9a.5.5 0 0 1-.289.453l-6 2.8a.5.5 0 0 1-.422 0l-6-2.8A.5.5 0 0 1 1 10.95v-6.9a.5.5 0 0 1 .289-.453l6-2.8ZM2 4.806L7 6.93v6.034l-5-2.333V4.806Zm6 8.159l5-2.333V4.806L8 6.93v6.034Zm-.5-6.908l4.772-2.028L7.5 1.802L2.728 4.029L7.5 6.057Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}