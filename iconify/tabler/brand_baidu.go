package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BrandBaidu(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 9.5a1 1.5 0 1 0 2 0a1 1.5 0 1 0-2 0m10.463 2.096c1.282 1.774 3.476 3.416 3.476 3.416s1.921 1.574.593 3.636C17.204 20.711 13.64 19.8 13.64 19.8s-1.416-.44-3.06-.088c-1.644.356-3.06.22-3.06.22s-2.055-.22-2.47-2.304c-.416-2.084 1.918-3.638 2.102-3.858c.182-.222 1.409-.966 2.284-2.394c.875-1.428 3.337-2.287 5.027.221zM8 4.5a1 1.5 0 1 0 2 0a1 1.5 0 1 0-2 0m6 0a1 1.5 0 1 0 2 0a1 1.5 0 1 0-2 0m4 5a1 1.5 0 1 0 2 0a1 1.5 0 1 0-2 0"/>`),
		g.Group(children),
	)
}