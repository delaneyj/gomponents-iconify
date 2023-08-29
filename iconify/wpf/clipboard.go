package wpf

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Clipboard(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<path fill="currentColor" d="M13 0c-.798 0-1.448.523-1.719 1H10c-1.094 0-2 .906-2 2H6C4.344 3 3 4.344 3 6v17c0 1.656 1.344 3 3 3h14c1.656 0 3-1.344 3-3V6c0-1.656-1.344-3-3-3h-2c0-1.094-.906-2-2-2h-1.281c-.271-.477-.92-1-1.719-1zm-3 3h6v2h-6V3zM6 6h2.281c.349.597.986 1 1.719 1h6c.733 0 1.37-.403 1.719-1H20v17H6V6zm2 4v2h10v-2H8zm0 4v2h7v-2H8zm0 4v2h10v-2H8z"/>`),
		g.Group(children),
	)
}