package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Combine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M2048 896v128h-677l210 211l-90 90l-365-365l365-365l90 90l-210 211h677zM467 685l90-90l365 365l-365 365l-90-90l210-211H0V896h677L467 685z"/>`),
		g.Group(children),
	)
}