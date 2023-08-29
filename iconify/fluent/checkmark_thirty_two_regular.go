package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CheckmarkThirtyTwoRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M29.726 5.312a1 1 0 0 1-.038 1.414l-19 18a1 1 0 0 1-1.42-.044l-7-7.5a1 1 0 1 1 1.463-1.364l6.313 6.763L28.312 5.274a1 1 0 0 1 1.414.038Z"/>`),
		g.Group(children),
	)
}