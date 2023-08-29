package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HeartCircleSixteenRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="m8 7l-.421-.492a1.465 1.465 0 1 0-2.157 1.98l2.4 2.44c.098.1.259.1.357 0l2.4-2.44a1.465 1.465 0 1 0-2.157-1.98L8 7ZM2 8a6 6 0 1 1 12 0A6 6 0 0 1 2 8Zm6-5a5 5 0 1 0 0 10A5 5 0 0 0 8 3Z"/>`),
		g.Group(children),
	)
}