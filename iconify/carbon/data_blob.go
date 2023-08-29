package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DataBlob(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M28 30H4a2.002 2.002 0 0 1-2-2V4a2.002 2.002 0 0 1 2-2h24a2.002 2.002 0 0 1 2 2v24a2.002 2.002 0 0 1-2 2ZM4 4v24h24V4Z"/><path fill="currentColor" d="M13 7h2v7h-2zM8 7h2v7H8zm14 7h-2a2.002 2.002 0 0 1-2-2V9a2.002 2.002 0 0 1 2-2h2a2.002 2.002 0 0 1 2 2v3a2.002 2.002 0 0 1-2 2zm-2-5v3h2V9zm2 9h2v7h-2zM8 18h2v7H8zm9 7h-2a2.002 2.002 0 0 1-2-2v-3a2.002 2.002 0 0 1 2-2h2a2.002 2.002 0 0 1 2 2v3a2.002 2.002 0 0 1-2 2zm-2-5v3h2v-3z"/>`),
		g.Group(children),
	)
}