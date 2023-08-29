package uil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Inbox(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M19.056 2h-14a3.003 3.003 0 0 0-3 3v14a3.003 3.003 0 0 0 3 3h14a3.003 3.003 0 0 0 3-3V5a3.003 3.003 0 0 0-3-3Zm-14 2h14a1.001 1.001 0 0 1 1 1v8H17.59a1.997 1.997 0 0 0-1.664.89L14.52 16H9.59l-1.406-2.11A1.997 1.997 0 0 0 6.52 13H4.056V5a1.001 1.001 0 0 1 1-1Zm14 16h-14a1.001 1.001 0 0 1-1-1v-4H6.52l1.406 2.11A1.997 1.997 0 0 0 9.59 18h4.93a1.997 1.997 0 0 0 1.664-.89L17.59 15h2.465v4a1.001 1.001 0 0 1-1 1Z"/>`),
		g.Group(children),
	)
}