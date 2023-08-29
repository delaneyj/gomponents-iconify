package bxs

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Dizzy(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 2C6.486 2 2 6.486 2 12s4.486 10 10 10s10-4.486 10-10S17.514 2 12 2zM8 12.414l-1.293 1.293l-1.414-1.414L6.586 11L5.293 9.707l1.414-1.414L8 9.586l1.293-1.293l1.414 1.414L9.414 11l1.293 1.293l-1.414 1.414L8 12.414zM14 18h-4v-2h4v2zm4.707-5.707l-1.414 1.414L16 12.414l-1.293 1.293l-1.414-1.414L14.586 11l-1.293-1.293l1.414-1.414L16 9.586l1.293-1.293l1.414 1.414L17.414 11l1.293 1.293z"/>`),
		g.Group(children),
	)
}