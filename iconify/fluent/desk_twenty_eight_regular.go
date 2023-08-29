package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DeskTwentyEightRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M4.754 4a2.75 2.75 0 0 0-2.75 2.75v13.5A3.75 3.75 0 0 0 5.754 24h4.496A3.75 3.75 0 0 0 14 20.25V10.5h10.5v12.75a.75.75 0 0 0 1.5 0V6.75A2.75 2.75 0 0 0 23.25 4H4.755Zm-1.25 6.5H12.5v9.75a2.25 2.25 0 0 1-2.25 2.25H5.754a2.25 2.25 0 0 1-2.25-2.25V10.5Zm0-1.5V6.75c0-.69.56-1.25 1.25-1.25H23.25c.69 0 1.25.56 1.25 1.25V9H3.505Zm2.746 4a.75.75 0 0 0 0 1.5h3.5a.75.75 0 1 0 0-1.5h-3.5Z"/>`),
		g.Group(children),
	)
}