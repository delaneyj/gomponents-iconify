package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LogoAdobePhotoshopOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 2h20v20H2V2Zm2 2v16h16v-6.5h-1c-1.015 0-1.99.403-2.707 1.121l-.293.293V18h-2v-7h2v1.331a5.828 5.828 0 0 1 3-.831h1V4H4Zm4 2v10h5v2H6V6h2Z"/>`),
		g.Group(children),
	)
}