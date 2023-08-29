package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LogoRedHatAnsible(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="m21.91 20.586l-5-11a1.001 1.001 0 0 0-1.82 0l-5 11a1 1 0 0 0 1.82.828l1.999-4.396l6.498 4.787a1 1 0 0 0 1.504-1.22ZM16 12.417l2.499 5.498l-3.744-2.759L16 12.416Z"/><path fill="currentColor" d="M16 30C8.28 30 2 23.72 2 16S8.28 2 16 2s14 6.28 14 14s-6.28 14-14 14Zm0-26C9.383 4 4 9.383 4 16s5.383 12 12 12s12-5.383 12-12S22.617 4 16 4Z"/>`),
		g.Group(children),
	)
}