package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Canva(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M33.563 22.611L31.341 28.5l-2.223-5.889M20.9 26.278a2.229 2.229 0 0 1-2.223 2.222h0a2.229 2.229 0 0 1-2.222-2.222v-1.445a2.229 2.229 0 0 1 2.222-2.222h0a2.229 2.229 0 0 1 2.223 2.222m0 3.667v-5.889m6.63 5.889v-3.667a2.229 2.229 0 0 0-2.222-2.222h0a2.229 2.229 0 0 0-2.222 2.222V28.5m0-3.667v-2.222m-8.67 2.889h0a2.987 2.987 0 0 1-3 3h0a2.987 2.987 0 0 1-3-3v-3a2.987 2.987 0 0 1 3-3h0a2.895 2.895 0 0 1 2.89 3h0m25.278 3.778A2.229 2.229 0 0 1 37.36 28.5h0a2.229 2.229 0 0 1-2.222-2.222v-1.445a2.229 2.229 0 0 1 2.222-2.222h0a2.229 2.229 0 0 1 2.223 2.222m.001 3.667v-5.889"/><circle cx="24" cy="24" r="21.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/>`),
		g.Group(children),
	)
}