package codicon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Newline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M12 5.5v1.984a.5.5 0 0 1-.5.5H4.618l1.633-1.633l-.707-.707l-2.121 2.121L3 8.188v.568L5.544 11.3l.707-.707l-1.61-1.609H11.5a1.5 1.5 0 0 0 1.5-1.5V5.5h-1Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}