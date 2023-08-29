package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Francomanca(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<ellipse cx="24" cy="24" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="20.04" ry="10.47" transform="rotate(-15.69 24.017 24.017)"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M30.18 20.13c-8.87-.92-21.08 2.56-19.47 7.52s10.6 4.61 18.9 1s10.25-9 7.26-11.35s-11.3-1.21-14.18-.52"/>`),
		g.Group(children),
	)
}