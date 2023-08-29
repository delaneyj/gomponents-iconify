package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RemoveContent(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="m2045 1245l-355 355l355 355l-90 90l-355-355l-355 355l-90-90l355-355l-355-355l90-90l355 355l355-355l90 90zM256 1792h864l-128 128H128V128h1792v864l-128 128V256H256v1536z"/>`),
		g.Group(children),
	)
}