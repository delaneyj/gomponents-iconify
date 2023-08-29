package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StackoverflowSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="m12.59 5.787l-3.5-5l.82-.574l3.5 5l-.82.574Zm-.925 1.085l-5-4.5l.67-.744l5 4.5l-.67.744Zm-.491 1.174l-6.1-3.1l.453-.892l6.1 3.1l-.454.892Zm-.283 1.442l-6.7-1.5l.218-.976l6.7 1.5l-.218.976ZM2 9h1v5h9V9h1v6H2V9Zm8.95 2.197l-7-.7l.1-.995l7 .7l-.1.995ZM11 13H4v-1h7v1Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}