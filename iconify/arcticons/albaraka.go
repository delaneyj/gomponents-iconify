package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Albaraka(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-miterlimit="3.864" d="M7.654 13.81a9.31 9.31 0 1 1 17.506 4.42c-.125.23-7.26 12.103-8.112 13.52a.097.097 0 0 1-.168-.003l-1.79-2.979a.595.595 0 0 1-.047-.515l4.96-13.14a3.308 3.308 0 1 0-6.346-1.305v12.378a.055.055 0 0 1-.101.03a2517.365 2517.365 0 0 1-4.787-7.987a9.259 9.259 0 0 1-1.115-4.419Zm32.692 20.38a9.31 9.31 0 1 1-17.505-4.42c.122-.226 7.02-11.703 8.065-13.44a.155.155 0 0 1 .266.005l1.76 2.928a.513.513 0 0 1 .04.447l-4.975 13.175a3.308 3.308 0 1 0 6.346 1.306v-12.38a.054.054 0 0 1 .1-.03c2.22 3.694 4.716 7.854 4.789 7.99a9.27 9.27 0 0 1 1.114 4.419Z"/>`),
		g.Group(children),
	)
}