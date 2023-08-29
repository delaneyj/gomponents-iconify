package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Openlp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="24" cy="24" r="21.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M44.823 29.359L13.047 5.495M10.43 40.677l34.393-11.319m-38.839 6.38l38.84-6.38l-42.02-1.734m.822-10.509l41.198 12.243l-36.77-19.78"/>`),
		g.Group(children),
	)
}