package eva

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LockFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g id="evaLockFill0"><g id="evaLockFill1"><g id="evaLockFill2" fill="currentColor"><circle cx="12" cy="15" r="1"/><path d="M17 8h-1V6.11a4 4 0 1 0-8 0V8H7a3 3 0 0 0-3 3v8a3 3 0 0 0 3 3h10a3 3 0 0 0 3-3v-8a3 3 0 0 0-3-3Zm-7-1.89A2.06 2.06 0 0 1 12 4a2.06 2.06 0 0 1 2 2.11V8h-4ZM12 18a3 3 0 1 1 3-3a3 3 0 0 1-3 3Z"/></g></g></g>`),
		g.Group(children),
	)
}