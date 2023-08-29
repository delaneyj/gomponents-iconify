package la

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SortNumericDown(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="m8.594 5l-.156.781s-.165.574-.563 1.157C7.477 7.52 6.98 8 6 8v2c1.375 0 2.32-.676 3-1.406V15h2V5zM22 5v18.688l-2.594-2.594L18 22.5l4.281 4.313l.719.687l.719-.688L28 22.5l-1.406-1.406L24 23.687V5zM8.5 17A3.514 3.514 0 0 0 5 20.5v.5h2v-.5c0-.875.625-1.5 1.5-1.5h1c.875 0 1.5.625 1.5 1.5c0 .457-.352.98-.938 1.344c-1.234.758-2.316 1.242-3.218 1.75c-.453.254-.867.496-1.219.875C5.273 24.848 5 25.418 5 26v1h8v-2H8.437c.735-.379 1.583-.758 2.688-1.438C12.141 22.927 13 21.845 13 20.5c0-1.922-1.578-3.5-3.5-3.5z"/>`),
		g.Group(children),
	)
}