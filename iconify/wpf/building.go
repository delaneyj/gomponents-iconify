package wpf

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Building(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<path fill="currentColor" d="M6 0C4.355 0 3 1.355 3 3v23h20V3c0-1.645-1.355-3-3-3H6zm0 2h14c.555 0 1 .445 1 1v21h-2v-5h-5v5H5V3c0-.555.445-1 1-1zm1 2v3h5V4H7zm7 0v3h5V4h-5zM7 9v3h5V9H7zm7 0v3h5V9h-5zm-7 5v3h5v-3H7zm7 0v3h5v-3h-5zm-7 5v3h5v-3H7z"/>`),
		g.Group(children),
	)
}