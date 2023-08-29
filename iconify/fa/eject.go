package fa

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Eject(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1536 1536"),
		g.Raw(`<path fill="currentColor" d="M14 755L724 45q19-19 45-19t45 19l710 710q19 19 13 32t-32 13H33q-26 0-32-13t13-32zm1459 557H65q-26 0-45-19t-19-45V992q0-26 19-45t45-19h1408q26 0 45 19t19 45v256q0 26-19 45t-45 19z"/>`),
		g.Group(children),
	)
}