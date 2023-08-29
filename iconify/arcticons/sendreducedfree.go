package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Sendreducedfree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<ellipse cx="24" cy="23.54" fill="none" stroke="currentColor" stroke-linecap="round" stroke-miterlimit="10" stroke-width="1.21" rx="13" ry="2.77"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-miterlimit="10" stroke-width="1.21" d="M11.63 17.17h24.74M11.63 4.5h24.74m-24.74 6.37h24.74"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-miterlimit="10" stroke-width="1.17" d="M27.21 43.5a7 7 0 0 0-6.42 0m9.21-6a14 14 0 0 0-12 0m14.66-5.75a20.69 20.69 0 0 0-17.32 0"/>`),
		g.Group(children),
	)
}