package uim

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SignInAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M21 12c0-.34-.02-.67-.05-1H12.5V9.5a1 1 0 0 0-1.707-.707l-2.5 2.5a1 1 0 0 0 0 1.414l2.5 2.5A1 1 0 0 0 12.5 14.5V13h8.45c.03-.33.05-.66.05-1Z"/><path fill="currentColor" d="M12.5 13v1.5a1 1 0 0 1-1.707.707l-2.5-2.5a1 1 0 0 1 0-1.414l2.5-2.5A1 1 0 0 1 12.5 9.5V11h8.45a10 10 0 1 0 0 2Z" opacity=".5"/>`),
		g.Group(children),
	)
}