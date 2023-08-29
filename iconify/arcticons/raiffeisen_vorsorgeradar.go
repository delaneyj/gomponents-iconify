package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RaiffeisenVorsorgeradar(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M6.824 36.934A21.504 21.504 0 0 1 21 2.708m12.431 40.619a21.517 21.517 0 0 1-24.634-4.124M24 2.5a21.509 21.509 0 0 1 19.326 30.931"/><circle cx="38.5" cy="38.5" r="7" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M35.85 42.5v-8h2.619a2.687 2.687 0 0 1 0 5.373H35.85m2.619 0l2.619 2.625"/>`),
		g.Group(children),
	)
}