package ion

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func IosMore(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path d="M255.8 218c-21 0-38 17-38 38s17 38 38 38 38-17 38-38-17-38-38-38z" fill="currentColor"/><path d="M102 218c-21 0-38 17-38 38s17 38 38 38 38-17 38-38-17-38-38-38z" fill="currentColor"/><path d="M410 218c-21 0-38 17-38 38s17 38 38 38 38-17 38-38-17-38-38-38z" fill="currentColor"/>`),
		g.Group(children),
	)
}