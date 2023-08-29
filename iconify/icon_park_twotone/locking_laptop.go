package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LockingLaptop(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTLockingLaptop0"><g fill="none" stroke="#fff" stroke-linejoin="round" stroke-width="4"><path stroke-linecap="round" d="M24 9H11a3 3 0 0 0-3 3v21h32v-5"/><path fill="#555" d="M4 33h40v2a6 6 0 0 1-6 6H10a6 6 0 0 1-6-6v-2Z"/><rect width="12" height="8" x="30" y="13" fill="#555" rx="3"/><path fill="#555" d="M36 7a3 3 0 0 1 3 3v3h-6v-3a3 3 0 0 1 3-3Z"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTLockingLaptop0)"/>`),
		g.Group(children),
	)
}