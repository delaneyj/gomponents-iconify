package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HereWego(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M10.365 35.125a2.321 2.321 0 0 1-1.641-3.963L26.33 13.555a2.321 2.321 0 0 1 3.283 3.283L12.006 34.445a2.313 2.313 0 0 1-1.641.68Zm13.205 0a2.321 2.321 0 0 1-1.64-3.963l17.607-17.607a2.321 2.321 0 1 1 3.283 3.283L25.212 34.445a2.313 2.313 0 0 1-1.641.68ZM6.822 25.462A2.321 2.321 0 0 1 5.18 21.5l7.945-7.945a2.321 2.321 0 1 1 3.283 3.283l-7.945 7.944a2.313 2.313 0 0 1-1.641.68Zm1.291 6.778l2.369-9.477m6.537-7.003l-2.369 9.476m6.669 7.005l2.369-9.477m6.537-7.004l-2.37 9.476"/>`),
		g.Group(children),
	)
}