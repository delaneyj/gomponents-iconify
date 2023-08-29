package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HorseShoe(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M7.502 1.042c-4.545 0-7.424 4.236-7.424 9.462c0 2.405 1.523 5.554 2.245 6.265c.723.711 2.324-.822 2.05-1.323a8.363 8.363 0 0 1-1.014-3.969c0-3.647 1.605-6.604 4.143-6.604c2.535 0 4.142 2.957 4.142 6.604c0 1.486-.41 2.855-1.011 3.957c-.272.507 1.469 2.019 2.045 1.335c.577-.684 2.247-3.859 2.247-6.265c0-5.226-2.88-9.462-7.423-9.462z"/>`),
		g.Group(children),
	)
}