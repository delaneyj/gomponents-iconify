package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SelectionAllThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M108 40a4 4 0 0 1 4-4h32a4 4 0 0 1 0 8h-32a4 4 0 0 1-4-4Zm36 172h-32a4 4 0 0 0 0 8h32a4 4 0 0 0 0-8Zm64-176h-24a4 4 0 0 0 0 8h24a4 4 0 0 1 4 4v24a4 4 0 0 0 8 0V48a12 12 0 0 0-12-12Zm8 72a4 4 0 0 0-4 4v32a4 4 0 0 0 8 0v-32a4 4 0 0 0-4-4Zm0 72a4 4 0 0 0-4 4v24a4 4 0 0 1-4 4h-24a4 4 0 0 0 0 8h24a12 12 0 0 0 12-12v-24a4 4 0 0 0-4-4ZM40 148a4 4 0 0 0 4-4v-32a4 4 0 0 0-8 0v32a4 4 0 0 0 4 4Zm32 64H48a4 4 0 0 1-4-4v-24a4 4 0 0 0-8 0v24a12 12 0 0 0 12 12h24a4 4 0 0 0 0-8ZM40 76a4 4 0 0 0 4-4V48a4 4 0 0 1 4-4h24a4 4 0 0 0 0-8H48a12 12 0 0 0-12 12v24a4 4 0 0 0 4 4Zm140 100a4 4 0 0 1-4 4H80a4 4 0 0 1-4-4V80a4 4 0 0 1 4-4h96a4 4 0 0 1 4 4Zm-8-92H84v88h88Z"/>`),
		g.Group(children),
	)
}