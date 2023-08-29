package guidance

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LeftLuggage(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" d="M7.5 7.5v11m8-11v11m-1-11a3 3 0 1 0-6 0M1 1.5h21.5v21H1m18.5-4h-16v-.177l.202-1.345a26.737 26.737 0 0 0 0-7.956L3.5 7.676V7.5h16v.176l-.203 1.346a26.748 26.748 0 0 0 0 7.956l.203 1.345v.177Z"/>`),
		g.Group(children),
	)
}