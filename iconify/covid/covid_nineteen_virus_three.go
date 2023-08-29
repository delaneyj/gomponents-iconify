package covid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CovidNineteenVirusThree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" stroke-linejoin="round" d="M7.851 4.958L6.215 3.322a2.046 2.046 0 0 0-2.893 2.893l1.6 1.694m11.227 11.133l1.636 1.636a2.046 2.046 0 0 0 2.893-2.893l-1.636-1.636m-14.084 0l-1.636 1.636a2.046 2.046 0 0 0 2.893 2.893l1.636-1.637M19.075 7.909l1.589-1.694a2.047 2.047 0 1 0-2.893-2.893L16.14 4.953"/><path stroke-linecap="round" stroke-linejoin="round" d="M21.205 9.955h-1.291a8.181 8.181 0 0 0-5.869-5.869V2.8a2.045 2.045 0 0 0-4.09 0v1.286a8.181 8.181 0 0 0-5.869 5.869H2.8a2.045 2.045 0 0 0 0 4.09h1.286a8.181 8.181 0 0 0 5.869 5.869v1.291a2.045 2.045 0 0 0 4.09 0v-1.291a8.182 8.182 0 0 0 5.869-5.869h1.291a2.045 2.045 0 1 0 0-4.09v0Z"/><path stroke-linecap="round" stroke-linejoin="round" d="M13.534 13.023a2.557 2.557 0 1 0 0-5.114a2.557 2.557 0 0 0 0 5.114Z"/><path d="M9.443 15.955a.375.375 0 0 1 0-.75m0 .75a.375.375 0 0 0 0-.75"/></g>`),
		g.Group(children),
	)
}