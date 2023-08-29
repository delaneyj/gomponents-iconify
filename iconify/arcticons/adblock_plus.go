package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AdblockPlus(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M30.532 30.015V17.988h3.909a4.059 4.059 0 0 1 0 8.118h-3.909"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width=".866" d="M24.975 23.999a3.007 3.007 0 0 1 0 6.013h-4.96V17.985h4.96a3.007 3.007 0 1 1 0 6.014Zm0 0h-4.961"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m17.468 30.015l-3.909-12.027L9.5 30.015m1.353-4.059h5.262"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M31.663 5.5H16.337L5.5 16.337v15.326L16.337 42.5h15.326L42.5 31.663V16.337L31.663 5.5z"/>`),
		g.Group(children),
	)
}