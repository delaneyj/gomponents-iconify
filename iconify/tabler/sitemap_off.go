package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SitemapOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 17a2 2 0 0 1 2-2h2a2 2 0 0 1 2 2v2a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2zm16-2a2 2 0 0 1 2 2m-.591 3.42c-.362.358-.86.58-1.409.58h-2a2 2 0 0 1-2-2v-2c0-.549.221-1.046.579-1.407M9 5a2 2 0 0 1 2-2h2a2 2 0 0 1 2 2v2a2 2 0 0 1-2 2m-7 6v-1a2 2 0 0 1 2-2h4m4 0a2 2 0 0 1 2 2M3 3l18 18"/>`),
		g.Group(children),
	)
}