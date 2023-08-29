package simple_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Googlenearby(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6.546 12L12 6.545L17.455 12l-5.454 5.454zm16.976-1.154L13.158.48a1.635 1.635 0 0 0-2.314 0L.478 10.846a1.629 1.629 0 0 0 0 2.305l10.37 10.372a1.629 1.629 0 0 0 2.304 0l10.37-10.372a1.629 1.629 0 0 0 0-2.305zM12 20.726l-8.727-8.728L12 3.27l8.727 8.728z"/>`),
		g.Group(children),
	)
}