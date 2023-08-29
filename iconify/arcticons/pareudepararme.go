package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Pareudepararme(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m36.973 37.872l-18.5 4.299L5.5 28.299l5.527-18.171l18.5-4.299L42.5 19.701ZM14.69 18.605l18.43 11.058m-18.43 0l18.43-11.058"/>`),
		g.Group(children),
	)
}