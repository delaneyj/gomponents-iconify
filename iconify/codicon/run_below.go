package codicon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RunBelow(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="m1.8 1.01l-.78.41v12l.78.42l9-6v-.83l-9-6zm.22 11.48V2.36l7.6 5.07l-7.6 5.06zM12.85 15h-.71l-2.5-2.5l.71-.71L12 13.44V8h1v5.45l1.65-1.65l.71.71L12.85 15z"/>`),
		g.Group(children),
	)
}