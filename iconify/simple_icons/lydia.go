package simple_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Lydia(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 0C5.373 0 0 5.373 0 12s5.373 12 12 12s12-5.373 12-12S18.627 0 12 0zm5.895 17.611a.421.421 0 0 1-.168.035h-1.155a.608.608 0 0 1-.56-.377l-4-9.613l-3.991 9.607a.61.61 0 0 1-.56.377H6.273a.42.42 0 0 1-.385-.59L10.91 5.575a.793.793 0 0 1 .726-.475h.748a.792.792 0 0 1 .726.48l5.003 11.482a.42.42 0 0 1-.218.549z"/>`),
		g.Group(children),
	)
}