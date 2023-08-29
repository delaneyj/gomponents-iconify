package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ReadingListTwentyFourFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M20 18a1 1 0 0 1 .117 1.993L20 20H7a1 1 0 0 1-.117-1.993L7 18h13Zm-3-3a1 1 0 0 1 .117 1.993L17 17H4a1 1 0 0 1-.117-1.993L4 15h13Zm3-3a1 1 0 0 1 .117 1.993L20 14H7a1 1 0 0 1-.117-1.993L7 12h13ZM6 5a3 3 0 0 1 2.78 1.873a1 1 0 0 1-1.803.857l-.05-.105A1 1 0 1 0 6 9h11.5a1 1 0 0 1 .117 1.993L17.5 11H6a3 3 0 0 1 0-6Zm14 1a1 1 0 0 1 .117 1.993L20 8h-9a1 1 0 0 1-.117-1.993L11 6h9Z"/>`),
		g.Group(children),
	)
}