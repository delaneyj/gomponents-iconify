package uil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func YinYang(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 15.5a1 1 0 1 0 1 1a1 1 0 0 0-1-1ZM12 2a5.545 5.545 0 0 0-.562.029A9.993 9.993 0 0 0 12 22a5.545 5.545 0 0 0 .562-.029A9.993 9.993 0 0 0 12 2Zm0 18A7.989 7.989 0 0 1 6.71 6.015A5.484 5.484 0 0 0 12 13a3.5 3.5 0 0 1 0 7Zm5.29-2.015A5.484 5.484 0 0 0 12 11a3.5 3.5 0 0 1 0-7a7.989 7.989 0 0 1 5.29 13.985ZM12 6.5a1 1 0 1 0 1 1a1 1 0 0 0-1-1Z"/>`),
		g.Group(children),
	)
}