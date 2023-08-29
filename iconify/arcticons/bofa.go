package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bofa(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M17.541 31.494L36.98 20.271M24.061 35.258L43.5 24.035m-39-.07l6.48-3.741m.04 7.505l6.48-3.741m6.439-11.246l6.52 3.764m-13-.023l6.521 3.764"/>`),
		g.Group(children),
	)
}