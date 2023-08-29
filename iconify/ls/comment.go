package ls

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Comment(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 717 717"),
		g.Raw(`<path fill="currentColor" d="M80 9h557c44 0 80 36 80 80v338c0 44-36 80-80 80H531l5 137c0 33-18 41-42 18L337 507H80c-44 0-80-36-80-80V89C0 45 36 9 80 9z"/>`),
		g.Group(children),
	)
}