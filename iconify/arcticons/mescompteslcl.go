package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Mescompteslcl(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M43.5 30.81c-24.66 25.9-56.92-8.48-26.87-23a9.1 9.1 0 0 0-.85 2.19C4.84 17 7.11 27.46 13.4 32.42a22.12 22.12 0 0 0 15 4.34c8.33-.46 15.14-5.95 15.14-5.95Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M15.78 15.74v10.79h5.4m13.44-10.79v10.79h5.39M31 22.91a3.62 3.62 0 0 1-3.63 3.63h0a3.62 3.62 0 0 1-3.62-3.63v-3.56a3.62 3.62 0 0 1 3.62-3.62h0A3.62 3.62 0 0 1 31 19.35"/>`),
		g.Group(children),
	)
}