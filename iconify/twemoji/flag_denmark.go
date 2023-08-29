package twemoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagDenmark(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 36 36"),
		g.Raw(`<path fill="#C60C30" d="M32 5H15v11h21V9a4 4 0 0 0-4-4zM15 31h17c2.209 0 4-1.791 4-4.5V20H15v11zM0 20v6.5C0 29.209 1.791 31 4 31h7V20H0zM11 5H4a4 4 0 0 0-4 4v7h11V5z"/><path fill="#EEE" d="M15 5h-4v11H0v4h11v11h4V20h21v-4H15z"/>`),
		g.Group(children),
	)
}