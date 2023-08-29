package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BrandNeteaseMusic(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 4c-2.93 1.346-5 5.046-5 8.492C4 17 8 20 12 20s8-3 8-7c0-3.513-3.5-5.513-6-5.513S9 9 9 12c0 2 1.5 3 3 3s3-1 3-3c0-3.513-2-4.508-2-6.515c0-3.504 3.5-2.603 4-1.502"/>`),
		g.Group(children),
	)
}