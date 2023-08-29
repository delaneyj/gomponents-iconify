package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Messages(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSMessages0"><g fill="none" stroke-width="4"><path fill="#fff" stroke="#fff" d="M39 6H9a3 3 0 0 0-3 3v30a3 3 0 0 0 3 3h30a3 3 0 0 0 3-3V9a3 3 0 0 0-3-3Z"/><path fill="#000" stroke="#000" stroke-linecap="round" stroke-linejoin="round" d="M35 23c0 5.523-4.477 10-10 10H15V23c0-5.523 4.477-10 10-10s10 4.477 10 10Z"/><path stroke="#000" stroke-linecap="round" stroke-linejoin="round" d="M22 21h6m-6 6h2"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSMessages0)"/>`),
		g.Group(children),
	)
}