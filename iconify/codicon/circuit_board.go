package codicon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CircuitBoard(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M14.5 1h-13l-.5.5v13l.5.5h13l.5-.5v-13l-.5-.5zM14 14H5v-2h2.3c.3.6 1 1 1.7 1c1.1 0 2-.9 2-2s-.9-2-2-2s-2 .9-2 2H4v3H2V2h2v2.3c-.6.3-1 1-1 1.7c0 1.1.9 2 2 2s2-.9 2-2h2c0 1.1.9 2 2 2s2-.9 2-2s-.9-2-2-2c-.7 0-1.4.4-1.7 1H6.7c-.3-.6-1-1-1.7-1V2h9v12zm-6-3c0-.6.4-1 1-1s1 .4 1 1s-.4 1-1 1s-1-.4-1-1zM5 5c.6 0 1 .4 1 1s-.4 1-1 1s-1-.4-1-1s.4-1 1-1zm6 0c.6 0 1 .4 1 1s-.4 1-1 1s-1-.4-1-1s.4-1 1-1z"/>`),
		g.Group(children),
	)
}