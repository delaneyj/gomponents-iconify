package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShieldSlashLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M52.44 36a6 6 0 0 0-9.63 7A14 14 0 0 0 34 56v58.77c0 88.25 74.68 117.52 89.65 122.49a13.5 13.5 0 0 0 8.7 0c7.76-2.58 33.48-12.37 55.29-34.76l15.92 17.5a6 6 0 0 0 4.44 2a6 6 0 0 0 4.44-10Zm76.12 189.93a1.57 1.57 0 0 1-1.13 0C113.84 221.38 46 194.8 46 114.79V56a2 2 0 0 1 2-2h4.62l126.93 139.62c-20.01 20.82-43.83 29.9-50.99 32.27ZM222 56v58.77c0 19.38-3.67 37.37-10.92 53.47a6 6 0 0 1-11-4.93c6.55-14.54 9.87-30.87 9.87-48.54V56a2 2 0 0 0-2-2H98.52a6 6 0 1 1 0-12H208a14 14 0 0 1 14 14Z"/>`),
		g.Group(children),
	)
}