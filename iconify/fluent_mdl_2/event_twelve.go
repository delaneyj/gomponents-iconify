package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EventTwelve(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M1536 171h341v1877H0V171h341V0h171v171h853V0h171v171zm171 1706V683H171v1194h1536zm0-1365V341H171v171h1536z"/>`),
		g.Group(children),
	)
}