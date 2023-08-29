package subway

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Magic(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M349.1 0v116.4L116.4 349.1H0V512h162.9V395.6l232.7-232.7H512V0z"/>`),
		g.Group(children),
	)
}