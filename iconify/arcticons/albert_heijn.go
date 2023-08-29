package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AlbertHeijn(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24.297 31.477V16.129m0 8.624l-5.004-5.887a2.689 2.689 0 0 0-4.737 1.742v8.29a2.689 2.689 0 0 0 4.737 1.741l10.008-11.773a2.689 2.689 0 0 1 4.737 1.742v10.87"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M42.04 23.175L33.023 6.213a3.227 3.227 0 0 0-4.425-1.301l-16.764 9.383a2.688 2.688 0 0 0-1.266 1.588L5.755 32.256a2.689 2.689 0 0 0 1.82 3.338l26.52 7.796a2.689 2.689 0 0 0 3.338-1.821l4.813-16.373a2.689 2.689 0 0 0-.206-2.02Z"/>`),
		g.Group(children),
	)
}