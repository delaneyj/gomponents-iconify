package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LockingComputer(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSLockingComputer0"><g fill="none" stroke="#fff" stroke-linejoin="round" stroke-width="4"><path stroke-linecap="round" d="M24 6H9a3 3 0 0 0-3 3v22a3 3 0 0 0 3 3h30a3 3 0 0 0 3-3v-4m-18 7v8"/><rect width="12" height="8" x="30" y="12" fill="#fff" rx="3"/><path fill="#fff" d="M36 6a3 3 0 0 1 3 3v3h-6V9a3 3 0 0 1 3-3Z"/><path stroke-linecap="round" d="M14 42h20"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSLockingComputer0)"/>`),
		g.Group(children),
	)
}