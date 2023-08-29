package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Checkerboard(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSCheckerboard0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M17 4H7a3 3 0 0 0-3 3v34a3 3 0 0 0 3 3h34a3 3 0 0 0 3-3V7a3 3 0 0 0-3-3H17Zm6 13h21M4 17h9m22 14h9M6 31h19m-8-10v23M31 4v23m0 8v9M17 4v9"/><path fill="#fff" d="M35 31a4 4 0 1 1-8 0a4 4 0 0 1 8 0ZM21 17a4 4 0 1 1-8 0a4 4 0 0 1 8 0Z"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSCheckerboard0)"/>`),
		g.Group(children),
	)
}