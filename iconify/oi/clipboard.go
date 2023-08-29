package oi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Clipboard(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 8 8"),
		g.Raw(`<path fill="currentColor" d="M3.5 0c-.28 0-.5.22-.5.5V1h-.75c-.14 0-.25.11-.25.25V2h3v-.75C5 1.11 4.89 1 4.75 1H4V.5c0-.28-.22-.5-.5-.5zM.25 1C.11 1 0 1.11 0 1.25v6.5c0 .14.11.25.25.25h6.5c.14 0 .25-.11.25-.25v-6.5C7 1.11 6.89 1 6.75 1H6v2H1V1H.25z"/>`),
		g.Group(children),
	)
}