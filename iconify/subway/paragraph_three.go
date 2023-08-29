package subway

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ParagraphThree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M256 458.5h256v-64H256v64zM0 351.8h512v-64H0v64zm256-106.6h256v-64H256v64zM0 74.5v64h512v-64H0z"/>`),
		g.Group(children),
	)
}