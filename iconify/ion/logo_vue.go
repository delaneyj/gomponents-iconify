package ion

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LogoVue(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="m256 144.03l-55.49-96.11h-79.43L256 281.61L390.92 47.92h-79.43L256 144.03z"/><path fill="currentColor" d="M409.4 47.92L256 313.61L102.6 47.92H15.74L256 464.08L496.26 47.92H409.4z"/>`),
		g.Group(children),
	)
}