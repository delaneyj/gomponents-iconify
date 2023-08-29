package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Extend(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTExtend0"><g fill="none"><rect width="36" height="36" x="6" y="6" fill="#555" stroke="#fff" stroke-width="4" rx="3"/><path fill="#fff" d="M35 12h-4.586c-.89 0-1.337 1.077-.707 1.707l4.586 4.586c.63.63 1.707.184 1.707-.707V13a1 1 0 0 0-1-1Zm-23 1v4.586c0 .89 1.077 1.337 1.707.707l4.586-4.586c.63-.63.184-1.707-.707-1.707H13a1 1 0 0 0-1 1Zm1 23h4.586c.89 0 1.337-1.077.707-1.707l-4.586-4.586c-.63-.63-1.707-.184-1.707.707V35a1 1 0 0 0 1 1Zm23-1v-4.586c0-.89-1.077-1.337-1.707-.707l-4.586 4.586c-.63.63-.184 1.707.707 1.707H35a1 1 0 0 0 1-1Z"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTExtend0)"/>`),
		g.Group(children),
	)
}