package subway

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ParagraphFive(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M0 462h256v-64H0v64zm0-106.7h512v-64H0v64zm0-106.6h512v-64H0v64zM0 78v64h512V78H0z"/>`),
		g.Group(children),
	)
}