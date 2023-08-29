package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DiffInline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M1792 0v2048H256V0h1536zM384 128v384h1280V128H384zm1280 1408H384v384h1280v-384zm0-256V640H384v640h1280z"/>`),
		g.Group(children),
	)
}