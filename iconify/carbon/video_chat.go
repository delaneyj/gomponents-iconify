package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VideoChat(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M8 12h9v2H8zm0 5h5v2H8z"/><path fill="currentColor" d="M21 26H4a2.002 2.002 0 0 1-2-2V8a2.002 2.002 0 0 1 2-2h17a2.002 2.002 0 0 1 2 2v4.057l5.419-3.87A1 1 0 0 1 30 9v14a1 1 0 0 1-1.581.814L23 19.944V24a2.002 2.002 0 0 1-2 2ZM4 8v16.001L21 24v-6a1 1 0 0 1 1.581-.814L28 21.056V10.944l-5.419 3.87A1 1 0 0 1 21 14V8Z"/>`),
		g.Group(children),
	)
}