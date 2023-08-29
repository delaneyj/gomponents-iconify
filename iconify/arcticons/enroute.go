package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Enroute(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M34.986 15.35c-.063 0-.124.008-.186.01a10.697 10.697 0 0 0-19.988-3.212c-.065-.001-.127-.01-.192-.01a9.12 9.12 0 0 0 0 18.24h20.366a7.514 7.514 0 1 0 0-15.028Zm-6.499 17.395l12.609 3.211l-10.39 5.573v-4.392l-2.219-4.392zm-6.565 4.392h3.495m-11.004 0h3.495m-11.004 0h3.495"/>`),
		g.Group(children),
	)
}