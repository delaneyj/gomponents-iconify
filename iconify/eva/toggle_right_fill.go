package eva

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ToggleRightFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g id="evaToggleRightFill0"><g id="evaToggleRightFill1"><g id="evaToggleRightFill2" fill="currentColor"><circle cx="15" cy="12" r="1"/><path d="M15 5H9a7 7 0 0 0 0 14h6a7 7 0 0 0 0-14Zm0 10a3 3 0 1 1 3-3a3 3 0 0 1-3 3Z"/></g></g></g>`),
		g.Group(children),
	)
}