package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StoragePool(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M28 30H4a2.002 2.002 0 0 1-2-2V4a2.002 2.002 0 0 1 2-2h24a2.002 2.002 0 0 1 2 2v24a2.002 2.002 0 0 1-2 2ZM4 4v24h24V4Z"/><path fill="currentColor" d="M17.5 13A3.5 3.5 0 1 1 21 9.5a3.504 3.504 0 0 1-3.5 3.5zm0-5A1.5 1.5 0 1 0 19 9.5A1.502 1.502 0 0 0 17.5 8zm-3 18a3.5 3.5 0 1 1 3.5-3.5a3.504 3.504 0 0 1-3.5 3.5zm0-5a1.5 1.5 0 1 0 1.5 1.5a1.502 1.502 0 0 0-1.5-1.5zm-5-3a3.5 3.5 0 1 1 3.5-3.5A3.504 3.504 0 0 1 9.5 18zm0-5a1.5 1.5 0 1 0 1.5 1.5A1.502 1.502 0 0 0 9.5 13zm13 8a3.5 3.5 0 1 1 3.5-3.5a3.504 3.504 0 0 1-3.5 3.5zm0-5a1.5 1.5 0 1 0 1.5 1.5a1.502 1.502 0 0 0-1.5-1.5z"/>`),
		g.Group(children),
	)
}