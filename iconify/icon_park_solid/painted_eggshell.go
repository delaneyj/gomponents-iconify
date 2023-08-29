package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PaintedEggshell(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSPaintedEggshell0"><g fill="none"><path fill="#fff" fill-rule="evenodd" stroke="#fff" stroke-width="4" d="M24 44c10.252 0 16-6.954 16-18S31.132 4 24 4S8 14.954 8 26s5.748 18 16 18Z" clip-rule="evenodd"/><path fill="#000" d="M21 38a3 3 0 1 0 0-6a3 3 0 0 0 0 6Z"/><path fill="#000" fill-rule="evenodd" d="M16 29.668a2 2 0 1 0 0-4a2 2 0 0 0 0 4Z" clip-rule="evenodd"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSPaintedEggshell0)"/>`),
		g.Group(children),
	)
}