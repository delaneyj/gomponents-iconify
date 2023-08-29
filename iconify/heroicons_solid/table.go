package heroicons_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Table(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M5 4a3 3 0 0 0-3 3v6a3 3 0 0 0 3 3h10a3 3 0 0 0 3-3V7a3 3 0 0 0-3-3H5Zm-1 9v-1h5v2H5a1 1 0 0 1-1-1Zm7 1h4a1 1 0 0 0 1-1v-1h-5v2Zm0-4h5V8h-5v2ZM9 8H4v2h5V8Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}