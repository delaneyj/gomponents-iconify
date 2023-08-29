package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NavaidVordme(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<circle cx="16" cy="16" r="2" fill="currentColor"/><path fill="currentColor" d="M30 6a2.002 2.002 0 0 0-2-2H4a2.002 2.002 0 0 0-2 2v20a2.002 2.002 0 0 0 2 2h24a2.002 2.002 0 0 0 2-2Zm-2 6.926L22.964 6H28ZM27.764 16L20.49 26h-8.982L4.236 16L11.51 6h8.982ZM9.036 6L4 12.925V6ZM4 19.075L9.036 26H4ZM22.964 26L28 19.074V26Z"/>`),
		g.Group(children),
	)
}