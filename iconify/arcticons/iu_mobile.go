package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func IuMobile(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M18.15 43.5h11.7m-11.7-39h11.7M24 4.5v39M11.325 11.813v12.024C11.325 31.15 17.175 37 24 37s12.675-5.85 12.675-13.163V11.813m-3.9-.271h7.8m-33.15 0h7.8"/>`),
		g.Group(children),
	)
}