package fontisto

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StarHalf(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11.855.029a1.305 1.305 0 0 0-1.454.721l-.003.008l-2.511 5.669a.953.953 0 0 1-.805.728h-.005L1.084 8.289a1.66 1.66 0 0 0-.73.406l.001-.001a1.201 1.201 0 0 0 .079 1.78l.002.002l4.535 4.212a1.243 1.243 0 0 1 .405 1.058v-.005l-.972 7.532a.72.72 0 0 0 .083.408l-.002-.004a.58.58 0 0 0 .812.242l-.003.001l6.398-3.726c.081 0 .162-.081.243-.081V.027z"/>`),
		g.Group(children),
	)
}