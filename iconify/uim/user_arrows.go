package uim

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UserArrows(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6 16.5A3.5 3.5 0 1 1 9.5 13A3.504 3.504 0 0 1 6 16.5Z" opacity=".25"/><path fill="currentColor" d="M8.64 15.272a3.452 3.452 0 0 1-5.28 0A4.988 4.988 0 0 0 1 19.5a1 1 0 0 0 1 1h8a1 1 0 0 0 1-1a4.988 4.988 0 0 0-2.36-4.228zM18 16.5a3.5 3.5 0 1 1 3.5-3.5a3.504 3.504 0 0 1-3.5 3.5z" opacity=".5"/><path fill="currentColor" d="M20.64 15.272a3.452 3.452 0 0 1-5.28 0A4.988 4.988 0 0 0 13 19.5a1 1 0 0 0 1 1h8a1 1 0 0 0 1-1a4.988 4.988 0 0 0-2.36-4.228Z" opacity=".7"/><path fill="currentColor" d="m17.207 5.793l-2-2a1 1 0 0 0-1.414 1.414l.293.293H9.914l.293-.293a1 1 0 0 0-1.414-1.414l-2 2a1 1 0 0 0 0 1.414l2 2a1 1 0 0 0 1.414-1.414L9.914 7.5h4.172l-.293.293a1 1 0 1 0 1.414 1.414l2-2a1 1 0 0 0 0-1.414Z"/>`),
		g.Group(children),
	)
}