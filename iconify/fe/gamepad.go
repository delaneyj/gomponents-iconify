package fe

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Gamepad(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g id="feGamepad0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feGamepad1" fill="currentColor"><path id="feGamepad2" d="M13 15h-2a5 5 0 1 1-4-8h10a5 5 0 1 1-4 8Zm-5-4v-1a1 1 0 1 0-2 0v1H5a1 1 0 0 0 0 2h1v1a1 1 0 0 0 2 0v-1h1a1 1 0 0 0 0-2H8Zm9 0a1 1 0 1 0 0-2a1 1 0 0 0 0 2Zm-1 3a1 1 0 1 0 0-2a1 1 0 0 0 0 2Zm3-1a1 1 0 1 0 0-2a1 1 0 0 0 0 2Z"/></g></g>`),
		g.Group(children),
	)
}