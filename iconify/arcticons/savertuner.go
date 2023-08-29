package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Savertuner(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24 43.5h-9.75V8.4h5.85V4.5h7.8v3.9h5.85v35.1H24Zm0-21.45v-7.8m-3.9 3.9h7.8M24 37.65v-7.8m-3.9 3.9h7.8"/>`),
		g.Group(children),
	)
}