package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RazerAudio(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M11.84 18.97v10.163m6.056-18.602v26.816M24.006 3.5v41m6.101-30.944v21.01m6.053-12.67v4.145"/>`),
		g.Group(children),
	)
}