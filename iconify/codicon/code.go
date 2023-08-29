package codicon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Code(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M4.708 5.578L2.061 8.224l2.647 2.646l-.708.708l-3-3V7.87l3-3l.708.708zm7-.708L11 5.578l2.647 2.646L11 10.87l.708.708l3-3V7.87l-3-3zM4.908 13l.894.448l5-10L9.908 3l-5 10z"/>`),
		g.Group(children),
	)
}