package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ContrastViewCircle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSContrastViewCircle0"><g fill="none" stroke="#fff" stroke-linejoin="round" stroke-width="4"><path stroke-linecap="round" d="M24 4C12.954 4 4 12.954 4 24s8.954 20 20 20V4Z" clip-rule="evenodd"/><path fill="#fff" d="M24 4c11.046 0 20 8.954 20 20s-8.954 20-20 20V4Z"/><path stroke-linecap="round" d="M24 36H9m15-8H5m19-8H5m19-8H9"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSContrastViewCircle0)"/>`),
		g.Group(children),
	)
}