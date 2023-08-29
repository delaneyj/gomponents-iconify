package ls

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Tilemenu(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 717 717"),
		g.Raw(`<path fill="currentColor" d="M0 204h190V14H0v190zm238 0h190V14H238v190zm237 0h191V14H475v190zM0 442h190V252H0v190zm238 0h190V252H238v190zm237 0h191V252H475v190zM0 680h190V489H0v191zm238 0h190V489H238v191zm237 0h191V489H475v191z"/>`),
		g.Group(children),
	)
}