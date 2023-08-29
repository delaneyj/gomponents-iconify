package uil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AdjustAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M20.83 7.32a.2.2 0 0 0 0-.08a10 10 0 0 0-3.38-3.65A9.89 9.89 0 0 0 12 2a1 1 0 0 0-1 1v18a1 1 0 0 0 1 1h.28a10 10 0 0 0 8.55-14.68ZM13 4.06a8 8 0 0 1 2.49.74L13 9.12Zm0 9.06l4.17-7.22a7.89 7.89 0 0 1 1.58 1.83L13 17.69Zm1.16 6.57L19.75 10a8.36 8.36 0 0 1 .25 2a7.94 7.94 0 0 1-5.84 7.69Z"/>`),
		g.Group(children),
	)
}