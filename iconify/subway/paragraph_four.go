package subway

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ParagraphFour(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M128 464h256v-64H128v64zM0 357.3h512v-64H0v64zm128-170.6v64h256v-64H128zM0 80v64h512V80H0z"/>`),
		g.Group(children),
	)
}