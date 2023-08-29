package pajamas

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Diagram(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M5.5 5.5v-3h5v3h-5ZM4 2a1 1 0 0 1 1-1h6a1 1 0 0 1 1 1v4a1 1 0 0 1-1 1h-.103l1.535 2H14a1 1 0 0 1 1 1v4a1 1 0 0 1-1 1h-4a1 1 0 0 1-1-1v-4a1 1 0 0 1 1-1h.541L9.007 7H6.993L5.46 9H6a1 1 0 0 1 1 1v4a1 1 0 0 1-1 1H2a1 1 0 0 1-1-1v-4a1 1 0 0 1 1-1h1.568l1.535-2H5a1 1 0 0 1-1-1V2Zm-1.5 8.5v3h3v-3h-3Zm8 3v-3h3v3h-3Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}