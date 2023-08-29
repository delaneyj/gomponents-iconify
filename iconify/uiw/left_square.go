package uiw

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LeftSquare(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M18.182 0C19.186 0 20 .814 20 1.818v16.364A1.818 1.818 0 0 1 18.182 20H1.818A1.818 1.818 0 0 1 0 18.182V1.818C0 .814.814 0 1.818 0h16.364Zm-7.394 5.627l-3.87 3.971a.682.682 0 0 0 .005.957l3.87 3.895a.682.682 0 1 0 .967-.962L8.363 10.07l3.402-3.49a.682.682 0 1 0-.977-.952Z"/>`),
		g.Group(children),
	)
}