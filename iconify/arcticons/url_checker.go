package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UrlChecker(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M12.852 32.708a8.297 8.297 0 0 1-7.236-4.355a9.012 9.012 0 0 1 .009-8.704a8.296 8.296 0 0 1 7.245-4.338m0-.001l9.001-.035m-9.457 17.46l8.639.09m14.095-17.651a8.297 8.297 0 0 1 7.236 4.355a9.012 9.012 0 0 1-.009 8.704a8.296 8.296 0 0 1-7.245 4.338m.09-17.369l-9 .036m8.91 17.333l-8.64-.09m-10.984-8.749h16.733v.182H15.506z"/>`),
		g.Group(children),
	)
}