package ion

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func IosInformationEmpty(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<circle cx="251.5" cy="172" r="20" fill="currentColor"/><path d="M272 344V216h-48v8h16v120h-16v8h64v-8z" fill="currentColor"/>`),
		g.Group(children),
	)
}