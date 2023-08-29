package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RulePartial(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M23 16c-3.9 0-7 3.1-7 7s3.1 7 7 7s7-3.1 7-7s-3.1-7-7-7zm0 12V18c2.8 0 5 2.2 5 5s-2.2 5-5 5zM8 16h6v2H8zm0-6h12v2H8z"/><path fill="currentColor" d="M26 4c0-1.1-.9-2-2-2H4c-1.1 0-2 .9-2 2v13c0 4.1 2.2 7.8 5.8 9.7l5.2 2.8v-2.3l-4.2-2.3C5.8 23.4 4 20.3 4 17V4h20v9h2V4z"/>`),
		g.Group(children),
	)
}