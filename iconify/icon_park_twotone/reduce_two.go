package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ReduceTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTReduceTwo0"><g fill="none"><rect width="36" height="36" x="6" y="6" fill="#555" stroke="#fff" stroke-width="4" rx="3"/><path fill="#fff" d="M17 30h-4.586c-.89 0-1.337 1.077-.707 1.707l4.586 4.586c.63.63 1.707.184 1.707-.707V31a1 1 0 0 0-1-1Zm13 1v4.586c0 .89 1.077 1.337 1.707.707l4.586-4.586c.63-.63.184-1.707-.707-1.707H31a1 1 0 0 0-1 1Zm1-13h4.586c.89 0 1.337-1.077.707-1.707l-4.586-4.586c-.63-.63-1.707-.184-1.707.707V17a1 1 0 0 0 1 1Zm-13-1v-4.586c0-.89-1.077-1.337-1.707-.707l-4.586 4.586c-.63.63-.184 1.707.707 1.707H17a1 1 0 0 0 1-1Z"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTReduceTwo0)"/>`),
		g.Group(children),
	)
}