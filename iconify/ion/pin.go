package ion

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Pin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M336 96a80 80 0 1 0-96 78.39v283.17a32.09 32.09 0 0 0 2.49 12.38l10.07 24a3.92 3.92 0 0 0 6.88 0l10.07-24a32.09 32.09 0 0 0 2.49-12.38V174.39A80.13 80.13 0 0 0 336 96Zm-56 0a24 24 0 1 1 24-24a24 24 0 0 1-24 24Z"/>`),
		g.Group(children),
	)
}