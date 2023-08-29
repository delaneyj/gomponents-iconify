package vaadin

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Glasses(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M15.5 7h-.7c-.4-1.2-1.5-2-2.8-2s-2.4.9-2.8 2.1c-.3-.4-.7-.6-1.2-.6s-.9.2-1.2.6C6.4 5.9 5.3 5 4 5s-2.4.9-2.8 2H.5c-.3 0-.5.2-.5.5s.2.5.5.5H1c0 1.7 1.3 3 3 3c1.5 0 2.7-1.1 3-2.5c.3 0 .5-.2.5-.5s.2-.5.5-.5s.5.2.5.5s.2.5.5.5c.2 1.4 1.5 2.5 3 2.5c1.7 0 3-1.3 3-3h.5c.3 0 .5-.2.5-.5s-.2-.5-.5-.5zM4 10c-1.1 0-2-.9-2-2s.9-2 2-2s2 .9 2 2s-.9 2-2 2zm8 0c-1.1 0-2-.9-2-2s.9-2 2-2s2 .9 2 2s-.9 2-2 2z"/>`),
		g.Group(children),
	)
}