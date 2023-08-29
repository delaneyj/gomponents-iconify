package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StepInsert(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M1920 256v1408H128v-640h512V640h512V256h768zm-128 128h-512v384H768v384H256v384h1536V384zm-1274 0H0V256h518L355 93l90-90l317 317l-317 317l-90-90l163-163z"/>`),
		g.Group(children),
	)
}