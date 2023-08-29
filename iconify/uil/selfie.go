package uil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Selfie(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M16 2H8a3 3 0 0 0-3 3v14a3 3 0 0 0 3 3h8a3 3 0 0 0 3-3V5a3 3 0 0 0-3-3Zm1 17a1 1 0 0 1-1 1H8a1 1 0 0 1-1-1v-1h10Zm-5-5a3 3 0 0 1 2.82 2H9.18A3 3 0 0 1 12 14Zm-1-3a1 1 0 1 1 1 1a1 1 0 0 1-1-1Zm6 5h-.1a5 5 0 0 0-2.42-3.32A3 3 0 0 0 15 11a3 3 0 0 0-6 0a3 3 0 0 0 .52 1.68A5 5 0 0 0 7.1 16H7V5a1 1 0 0 1 1-1h8a1 1 0 0 1 1 1ZM12 5a1 1 0 1 0 1 1a1 1 0 0 0-1-1Z"/>`),
		g.Group(children),
	)
}