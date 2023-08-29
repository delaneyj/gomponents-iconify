package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LockingPicture(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSLockingPicture0"><g fill="none"><path fill="#fff" d="M44 26a2 2 0 1 0-4 0h4ZM24 8a2 2 0 1 0 0-4v4Zm15 32H9v4h30v-4ZM8 39V9H4v30h4Zm32-13v13h4V26h-4ZM9 8h15V4H9v4Zm0 32a1 1 0 0 1-1-1H4a5 5 0 0 0 5 5v-4Zm30 4a5 5 0 0 0 5-5h-4a1 1 0 0 1-1 1v4ZM8 9a1 1 0 0 1 1-1V4a5 5 0 0 0-5 5h4Z"/><path stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="m6 35l10.693-9.802a2 2 0 0 1 2.653-.044L32 36m-4-5l4.773-4.773a2 2 0 0 1 2.615-.186L42 31"/><rect width="12" height="8" x="30" y="12" fill="#fff" stroke="#fff" stroke-linejoin="round" stroke-width="4" rx="3"/><path fill="#fff" d="M36 6a3 3 0 0 1 3 3v3h-6V9a3 3 0 0 1 3-3Z"/><path fill="#fff" d="M39 12v2a2 2 0 0 0 2-2h-2Zm-6 0h-2a2 2 0 0 0 2 2v-2Zm4-3v3h4V9h-4Zm2 1h-6v4h6v-4Zm-4 2V9h-4v3h4Zm0-3a1 1 0 0 1 1-1V4a5 5 0 0 0-5 5h4Zm6 0a5 5 0 0 0-5-5v4a1 1 0 0 1 1 1h4Z"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSLockingPicture0)"/>`),
		g.Group(children),
	)
}