package codicon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Indent(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M4 5v1.984a.5.5 0 0 0 .5.5h6.882L9.749 5.851l.707-.707l2.121 2.121l.423.423v.568L10.456 10.8l-.707-.707l1.61-1.609H4.5a1.5 1.5 0 0 1-1.5-1.5V5h1Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}