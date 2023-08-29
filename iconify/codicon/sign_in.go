package codicon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SignIn(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="m11.02 3.77l.01-.01l.99.99V2.5l-.5-.5h-9l-.51.5v.493L2 3v10.29l.36.46l5 1.72L8 15v-1h3.52l.5-.5v-2.25l-1 1V13H8V4.71l-.33-.46L4.036 3h6.984v.77zM7 14.28l-4-1.34V3.72l4 1.34v9.22zm3.09-6.75h4.97v1h-4.93l1.59 1.6l-.71.7l-2.47-2.46v-.71l2.49-2.48l.7.7l-1.64 1.65z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}