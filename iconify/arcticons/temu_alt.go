package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TemuAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<rect width="37" height="37" x="5.5" y="5.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="2" ry="2"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M22.895 28.031V19.96l4.039 8.08l4.04-8.068v8.068M9.5 19.96h5.353m-2.676 8.08v-8.08m20.97 0v5.404a2.676 2.676 0 1 0 5.353 0V19.96M16.813 24h2.634m1.406 4.04h-4.04v-8.08h4.04"/>`),
		g.Group(children),
	)
}