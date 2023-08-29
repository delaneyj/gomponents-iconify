package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Liqueur(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSLiqueur0"><g fill="none" stroke-linecap="round" stroke-width="4"><path fill="#fff" stroke="#fff" stroke-linejoin="round" d="M23 31L7.41 13h31.177L22.999 31Z"/><path stroke="#fff" stroke-linejoin="round" d="M23 32v10"/><path stroke="#fff" d="M17 44h12"/><path stroke="#000" stroke-linejoin="round" d="M16 23h14"/><path stroke="#fff" stroke-linejoin="round" d="M35 18h3a6 6 0 1 0-5.917-5M11 17.144L19.534 27M35 17.143l-8.79 10.15"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSLiqueur0)"/>`),
		g.Group(children),
	)
}