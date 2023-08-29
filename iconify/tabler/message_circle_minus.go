package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MessageCircleMinus(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12.023 19.98A9.87 9.87 0 0 1 7.7 19L3 20l1.3-3.9C1.976 12.663 2.874 8.228 6.4 5.726c3.526-2.501 8.59-2.296 11.845.48c2.718 2.319 3.473 5.832 2.096 8.811M16 19h6"/>`),
		g.Group(children),
	)
}