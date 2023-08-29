package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Viper(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24.22 42.26L5.5 5.74h13.9l9.15 17.85l8.38-17.75h5.57Zm4.33-18.67l2.55 4.96"/>`),
		g.Group(children),
	)
}