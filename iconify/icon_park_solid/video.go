package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Video(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSVideo0"><g fill="none" stroke-linejoin="round" stroke-width="4"><path fill="#fff" stroke="#fff" d="M4 10a2 2 0 0 1 2-2h36a2 2 0 0 1 2 2v28a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2V10Z"/><path stroke="#000" stroke-linecap="round" d="M36 8v32M12 8v32m26-22h6m-6 12h6M4 18h6"/><path stroke="#fff" stroke-linecap="round" d="M4 16v4M9 8h6M9 40h6M33 8h6m-6 32h6"/><path stroke="#000" stroke-linecap="round" d="M4 30h6"/><path stroke="#fff" stroke-linecap="round" d="M4 28v4m40-4v4m0-16v4"/><path fill="#000" stroke="#000" d="m21 19l8 5l-8 5V19Z"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSVideo0)"/>`),
		g.Group(children),
	)
}