package la

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Rev(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M15 3v2.05C9.402 5.559 5 10.274 5 16c0 6.065 4.935 11 11 11h11V16a10.97 10.97 0 0 0-3.914-8.4l-1.863 1.086C23.505 10.32 25 12.986 25 16c0 4.962-4.037 9-9 9s-9-4.038-9-9c0-4.624 3.506-8.442 8-8.941V10l6-3.5L15 3zm1 10a3 3 0 0 0 0 6a3 3 0 0 0 0-6z"/>`),
		g.Group(children),
	)
}