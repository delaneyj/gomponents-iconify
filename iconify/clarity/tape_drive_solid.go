package clarity

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TapeDriveSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 36 36"),
		g.Raw(`<g id="clarityTapeDriveSolid0" fill="currentColor"><path d="M32 6H4a2 2 0 0 0-2 2v20a2 2 0 0 0 2 2h28a2 2 0 0 0 2-2V8a2 2 0 0 0-2-2Zm-2 18H6V12h24Z"/><path d="M12.21 23a5 5 0 1 0-5-5a5 5 0 0 0 5 5Zm0-7a2 2 0 1 1-2 2a2 2 0 0 1 2-2Zm11.58 7a5 5 0 1 0-5-5a5 5 0 0 0 5 5Zm0-7a2 2 0 1 1-2 2a2 2 0 0 1 2-2Z"/></g>`),
		g.Group(children),
	)
}