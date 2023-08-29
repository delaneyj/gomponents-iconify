package wpf

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DefineLocation(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<path fill="currentColor" d="M12.906-.031a1 1 0 0 0-.125.031A1 1 0 0 0 12 1v1.063C6.737 2.54 2.54 6.737 2.062 12H1a1 1 0 0 0-.094 0A1.001 1.001 0 0 0 1 14h1.063c.477 5.263 4.674 9.46 9.937 9.938V25a1 1 0 1 0 2 0v-1.063c5.263-.477 9.46-4.674 9.938-9.937H25a1 1 0 1 0 0-2h-1.063C23.46 6.737 19.264 2.54 14 2.062V1a1 1 0 0 0-1.094-1.031zM12 5.062V6a1 1 0 1 0 2 0v-.938A7.986 7.986 0 0 1 20.938 12H20a1 1 0 0 0-.094 0A1.001 1.001 0 0 0 20 14h.938A7.986 7.986 0 0 1 14 20.938V20a1 1 0 0 0-1.219-1A1 1 0 0 0 12 20v.938A7.986 7.986 0 0 1 5.062 14H6a1 1 0 1 0 0-2h-.938A7.986 7.986 0 0 1 12 5.062z"/>`),
		g.Group(children),
	)
}