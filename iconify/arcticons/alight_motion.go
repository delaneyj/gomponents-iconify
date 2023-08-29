package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AlightMotion(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M43.5 28.33a19.5 19.5 0 0 0-39 0"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M18.47 32.593a6.923 6.923 0 0 0 9.79-9.79"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M23.25 39.17a11.227 11.227 0 1 0 0-22.454"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M34.602 38.222a14.637 14.637 0 1 0-20.7-20.7"/>`),
		g.Group(children),
	)
}