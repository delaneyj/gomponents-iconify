package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AccountPinBoxLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m14 21l-2 2l-2-2H4.995A1.995 1.995 0 0 1 3 19.005V4.995C3 3.893 3.893 3 4.995 3h14.01C20.107 3 21 3.893 21 4.995v14.01A1.995 1.995 0 0 1 19.005 21H14Zm5-2V5H5v14h5.828L12 20.172L13.172 19H19Zm-11.028-.82a9.977 9.977 0 0 1-1.751-.978A6.994 6.994 0 0 1 12.102 14c2.4 0 4.517 1.207 5.778 3.047a9.987 9.987 0 0 1-1.724 1.025A4.993 4.993 0 0 0 12.101 16c-1.716 0-3.23.864-4.13 2.18ZM12 13a3.5 3.5 0 1 1 0-7a3.5 3.5 0 0 1 0 7Zm0-2a1.5 1.5 0 1 0 0-3a1.5 1.5 0 0 0 0 3Z"/>`),
		g.Group(children),
	)
}