package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RewindBackCircleBroken(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path d="m11 8.071l-3.222 2.302a2 2 0 0 0 0 3.254L11 15.93m5.5-.872V8.943a1 1 0 0 0-1.581-.814l-4.28 3.057a1 1 0 0 0 0 1.628l4.28 3.056a1 1 0 0 0 1.581-.813Z"/><path d="M7 3.338A9.954 9.954 0 0 1 12 2c5.523 0 10 4.477 10 10s-4.477 10-10 10S2 17.523 2 12c0-1.821.487-3.53 1.338-5"/></g>`),
		g.Group(children),
	)
}