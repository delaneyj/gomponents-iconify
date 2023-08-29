package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ScreencastLinear(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path stroke-linejoin="round" d="M14 20c3.771 0 5.657 0 6.828-1.172C22 17.657 22 15.771 22 12c0-3.771 0-5.657-1.172-6.828C19.657 4 17.771 4 14 4H6.5c-.464 0-.697 0-.892.02a4 4 0 0 0-3.589 3.588C2 7.803 2 8.036 2 8.5"/><path d="M11 20a9 9 0 0 0-9-9"/><path d="M8 20a6 6 0 0 0-6-6m3 6a3 3 0 0 0-3-3"/></g>`),
		g.Group(children),
	)
}