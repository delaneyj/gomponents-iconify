package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LockingWeb(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSLockingWeb0"><g fill="none"><path stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M24 40H7a3 3 0 0 1-3-3V11a3 3 0 0 1 3-3h34a3 3 0 0 1 3 3v12.059"/><path fill="#fff" stroke="#fff" stroke-width="4" d="M4 11a3 3 0 0 1 3-3h34a3 3 0 0 1 3 3v9H4v-9Z"/><rect width="12" height="8" x="32" y="33" stroke="#fff" stroke-linejoin="round" stroke-width="4" rx="3"/><path stroke="#fff" stroke-linejoin="round" stroke-width="4" d="M38 27a3 3 0 0 1 3 3v3h-6v-3a3 3 0 0 1 3-3Z"/><circle r="2" fill="#000" transform="matrix(0 -1 -1 0 10 14)"/><circle r="2" fill="#000" transform="matrix(0 -1 -1 0 16 14)"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSLockingWeb0)"/>`),
		g.Group(children),
	)
}