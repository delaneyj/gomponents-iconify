package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NavaidDme(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M16 20a4 4 0 1 1 4-4a4.005 4.005 0 0 1-4 4Zm0-6a2 2 0 1 0 2 2a2.002 2.002 0 0 0-2-2Z"/><path fill="currentColor" d="M24 24H8a2.002 2.002 0 0 1-2-2V10a2.002 2.002 0 0 1 2-2h16a2.002 2.002 0 0 1 2 2v12a2.002 2.002 0 0 1-2 2ZM8 10v12h16V10Z"/><path fill="currentColor" d="M28 28H4a2.002 2.002 0 0 1-2-2V6a2.002 2.002 0 0 1 2-2h24a2.002 2.002 0 0 1 2 2v20a2.002 2.002 0 0 1-2 2ZM4 6v20h24V6Z"/>`),
		g.Group(children),
	)
}