package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VoiceAccess(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="14.355" cy="19.496" r="4.934" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24.734 33.276v1.69a1.641 1.641 0 0 1-1.645 1.645H6.145A1.641 1.641 0 0 1 4.5 34.966v-1.69c2.159-7.02 18.973-5.817 20.234 0Zm.524-12.399v5.032m4.561-9.697v14.363M34.38 11.39v24.007m9.12-14.52v5.032m-4.56-9.697v14.363"/>`),
		g.Group(children),
	)
}