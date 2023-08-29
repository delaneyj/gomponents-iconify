package subway

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ParagraphSix(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M256 456.5h256v-64H256v64zM0 349.8h512v-64H0v64zm0-106.6h512v-64H0v64zM0 72.5v64h512v-64H0z"/>`),
		g.Group(children),
	)
}