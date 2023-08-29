package subway

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FileTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M325.5 0v128h128L325.5 0zm-21.3 0H69.5v512h384V149.3H304.2V0zm64 341.3L261.5 448L154.8 341.3h64v-128h85.3v128h64.1z"/>`),
		g.Group(children),
	)
}