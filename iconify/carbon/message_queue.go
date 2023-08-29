package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MessageQueue(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M28 28H4a2.002 2.002 0 0 1-2-2v-5h2v5h24v-5h2v5a2.002 2.002 0 0 1-2 2Z"/><path fill="currentColor" d="M7 21h18v2H7zm0-5h18v2H7zm0-5h18v2H7zm0-5h18v2H7z"/>`),
		g.Group(children),
	)
}