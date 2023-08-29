package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Sunrise(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M2 27h27.998v2H2zm14-7a4.005 4.005 0 0 1 4 4h2a6 6 0 0 0-12 0h2a4.005 4.005 0 0 1 4-4zm9 2h5v2h-5zm-3.313-5.1l3.506-3.507l1.414 1.414l-3.506 3.506zM16 4l-5 5l1.41 1.41L15 7.83V15h2V7.83l2.59 2.58L21 9l-5-5zM5.393 14.807l1.414-1.414l3.506 3.506L8.9 18.313zM2 22h5v2H2z"/>`),
		g.Group(children),
	)
}