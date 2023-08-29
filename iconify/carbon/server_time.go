package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ServerTime(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M28 30H4a2.002 2.002 0 0 1-2-2v-6a2.002 2.002 0 0 1 2-2h24a2.002 2.002 0 0 1 2 2v6a2.002 2.002 0 0 1-2 2ZM4 22v6h24v-6Z"/><circle cx="7" cy="25" r="1" fill="currentColor"/><path fill="currentColor" d="m19 11.586l-2-2V6h-2v4.414L17.586 13L19 11.586z"/><path fill="currentColor" d="M16 18a8 8 0 1 1 8-8a8.01 8.01 0 0 1-8 8Zm0-14a6 6 0 1 0 6 6a6.007 6.007 0 0 0-6-6Z"/>`),
		g.Group(children),
	)
}