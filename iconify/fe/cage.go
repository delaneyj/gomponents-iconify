package fe

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Cage(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g id="feCage0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feCage1" fill="currentColor"><path id="feCage2" d="M14 4.341c2.33.824 4 3.047 4 5.659v9h.5a1.5 1.5 0 0 1 0 3h-13a1.5 1.5 0 0 1 0-3H6v-9a6.002 6.002 0 0 1 4-5.659V4a2 2 0 1 1 4 0v.341ZM16 11v3h-3v-3h3Zm0-2h-3V6a4.183 4.183 0 0 1 3 3Zm-8 2h3v3H8v-3Zm0-2a4.183 4.183 0 0 1 3-3v3H8Zm8 7v3h-3v-3h3Zm-8 0h3v3H8v-3Z"/></g></g>`),
		g.Group(children),
	)
}