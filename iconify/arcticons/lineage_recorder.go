package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LineageRecorder(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="34.554" cy="24" r="8.945" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M18.679 24h2.538"/><rect width="3.31" height="9.606" x="11.396" y="19.197" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="1.655"/><rect width="3.31" height="17.84" x="4.5" y="15.08" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="1.655"/>`),
		g.Group(children),
	)
}