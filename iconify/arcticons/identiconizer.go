package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Identiconizer(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M35.1 35.1h7.4v7.4h-7.4zm-29.6 0h7.4v7.4H5.5z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M42.5 27.7V12.9h-7.4v7.4h-7.4v-7.4h7.4V5.5H12.9v7.4h7.4v7.4h-7.4v-7.4H5.5v14.8h7.4v7.4h7.4v7.4h7.4v-7.4h7.4v-7.4h7.4z"/>`),
		g.Group(children),
	)
}