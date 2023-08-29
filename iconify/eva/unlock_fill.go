package eva

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UnlockFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g id="evaUnlockFill0"><g id="evaUnlockFill1"><g id="evaUnlockFill2" fill="currentColor"><circle cx="12" cy="15" r="1"/><path d="M17 8h-7V6a2 2 0 0 1 4 0a1 1 0 0 0 2 0a4 4 0 0 0-8 0v2H7a3 3 0 0 0-3 3v8a3 3 0 0 0 3 3h10a3 3 0 0 0 3-3v-8a3 3 0 0 0-3-3Zm-5 10a3 3 0 1 1 3-3a3 3 0 0 1-3 3Z"/></g></g></g>`),
		g.Group(children),
	)
}