package icons_8

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NumericalSortingTwelve(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="m8.594 5l-.156.78s-.166.576-.563 1.157C7.478 7.52 6.98 8 6 8v2c1.376 0 2.32-.675 3-1.406V15h2V5H8.594zM22 5v18.688l-2.594-2.594L18 22.5l4.28 4.313l.72.687l.72-.688L28 22.5l-1.406-1.406L24 23.687V5h-2zM8.5 17A3.512 3.512 0 0 0 5 20.5v.5h2v-.5c0-.876.624-1.5 1.5-1.5h1c.876 0 1.5.624 1.5 1.5c0 .456-.353.98-.938 1.344c-1.234.76-2.316 1.244-3.218 1.75c-.452.253-.867.496-1.22.875c-.35.377-.624.95-.624 1.53v1h8v-2H8.437c.736-.378 1.58-.756 2.688-1.438C12.14 22.928 13 21.845 13 20.5c0-1.924-1.576-3.5-3.5-3.5h-1z"/>`),
		g.Group(children),
	)
}