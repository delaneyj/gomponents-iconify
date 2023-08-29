package ion

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChevronCollapseSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M102.145 108.514L256 241.855l153.855-133.341l-31.437-36.273L256 178.337L133.582 72.241l-31.437 36.273zm0 294.972L256 270.145l153.855 133.341l-31.437 36.273L256 333.663L133.582 439.759l-31.437-36.273z"/>`),
		g.Group(children),
	)
}