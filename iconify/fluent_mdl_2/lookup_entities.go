package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LookupEntities(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="m2048 835l-128-128V256h-512v512h-128V256H768v1408h512v-512h128v512h512v-451l128-128v707H0V128h2048v707zM640 768H128v384h512V768zm0 896v-384H128v384h512zm0-1408H128v384h512V256zm899 931l162-163h-677V896h677l-162-163l90-90l317 317l-317 317l-90-90z"/>`),
		g.Group(children),
	)
}