package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagTwoBoldDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path fill-rule="evenodd" d="M6.5 1.75a.75.75 0 0 0-1.5 0v20a.75.75 0 0 0 1.5 0v-20Z" clip-rule="evenodd" opacity=".5"/><path d="m13.558 3.873l-.413-.165a8.677 8.677 0 0 0-4.924-.452L6.5 3.6v10l1.72-.344a8.677 8.677 0 0 1 4.925.452a8.677 8.677 0 0 0 5.327.361l.1-.025a.9.9 0 0 0 .553-1.335l-1.56-2.601c-.342-.57-.513-.854-.553-1.163a1.508 1.508 0 0 1 0-.39c.04-.309.211-.594.553-1.163l1.278-2.13a.73.73 0 0 0-.803-1.084a7.3 7.3 0 0 1-4.482-.305Z"/></g>`),
		g.Group(children),
	)
}