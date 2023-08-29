package ion

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LogoAngular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M213.57 256h84.85l-42.43-89.36L213.57 256z"/><path fill="currentColor" d="M256 32L32 112l46.12 272L256 480l177.75-96L480 112Zm88 320l-26.59-56H194.58L168 352h-40L256 72l128 280Z"/>`),
		g.Group(children),
	)
}