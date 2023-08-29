package logos

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Payload(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path d="M116.204 0L256 84.031v159.5l-105.265 60.896v-159.5L10.772 61.008L116.204 0ZM105.49 171.121v124.463L0 232.13l105.489-61.008Z"/>`),
		g.Group(children),
	)
}