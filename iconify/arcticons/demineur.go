package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Demineur(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M18.993 19.86h10v10h-10z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M44.493 19.86h-4.25a16.882 16.882 0 0 0-2.3-4.707l2.132-2.133a.757.757 0 0 0 0-1.07l-3.173-3.172a.756.756 0 0 0-1.07 0l-2.136 2.136a16.977 16.977 0 0 0-4.703-2.294V4.36a1 1 0 0 0-1-1h-8a1 1 0 0 0-1 1v4.26a16.977 16.977 0 0 0-4.703 2.294l-2.137-2.136a.756.756 0 0 0-1.07 0L7.911 11.95a.756.756 0 0 0 0 1.07l2.132 2.132a16.88 16.88 0 0 0-2.3 4.707h-4.25a1 1 0 0 0-1 1v8a1 1 0 0 0 1 1h4.25a16.88 16.88 0 0 0 2.3 4.707L7.911 36.7a.756.756 0 0 0 0 1.07l3.172 3.173a.756.756 0 0 0 1.07 0l2.137-2.136a16.977 16.977 0 0 0 4.703 2.293v4.261a1 1 0 0 0 1 1h8a1 1 0 0 0 1-1V41.1a16.977 16.977 0 0 0 4.703-2.295l2.136 2.137a.756.756 0 0 0 1.07 0l3.173-3.173a.757.757 0 0 0 0-1.07l-2.132-2.132a16.882 16.882 0 0 0 2.3-4.707h4.25a1 1 0 0 0 1-1v-8a1 1 0 0 0-1-1Z"/>`),
		g.Group(children),
	)
}