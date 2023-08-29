package subway

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Massage(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M0 59.1v393.8h512V59.1H0zm433.2 39.4L256 275.7L78.8 98.5h354.4zm39.4 315H39.4V118.2L256 334.8l216.6-216.6v295.3z"/>`),
		g.Group(children),
	)
}