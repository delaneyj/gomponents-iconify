package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CheckboxComposite(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M2048 0v2048H0V0h2048zm-128 128H128v1792h1792V128zM768 1627l-429-430l90-90l339 338l851-850l90 90l-941 942z"/>`),
		g.Group(children),
	)
}