package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Magnifier(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<g fill="currentColor" fill-rule="evenodd"><path d="M17 5.954C17 2.665 14.317 0 11.009 0C7.698 0 5.016 2.665 5.016 5.954c0 3.287 2.683 5.952 5.993 5.952c3.308 0 5.991-2.665 5.991-5.952zm-11.066.065A5.082 5.082 0 0 1 11.026.943a5.08 5.08 0 0 1 5.088 5.076a5.08 5.08 0 0 1-5.088 5.075c-2.813 0-5.092-2.272-5.092-5.075zm-3.112 9.945L1 14.142l4.037-4.038s.096.765.58 1.247c.482.484 1.242.576 1.242.576l-4.037 4.037z"/><path d="M14.398 5.073c0 .572.44.356.44-.439c0-1.37-1.109-2.48-2.479-2.48c-.797 0-1.012.439-.439.439a2.479 2.479 0 0 1 2.478 2.48z"/></g>`),
		g.Group(children),
	)
}