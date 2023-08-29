package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HeaderOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M1792 384v1408h-128V646q-66 59-149 90t-171 32V640q54 0 103-15t92-44t75-69t55-89l17-39h106zm-896 0h128v1408H896v-640H256v640H128V384h128v640h640V384z"/>`),
		g.Group(children),
	)
}