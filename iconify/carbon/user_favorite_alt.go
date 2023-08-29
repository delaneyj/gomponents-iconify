package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UserFavoriteAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M28.766 4.256A4.212 4.212 0 0 0 23 4.032a4.212 4.212 0 0 0-5.766.224a4.319 4.319 0 0 0 0 6.044l5.764 5.84l.002-.002l.002.001l5.764-5.839a4.319 4.319 0 0 0 0-6.044zm-1.424 4.639l-4.34 4.397L23 13.29l-.002.002l-4.34-4.397a2.308 2.308 0 0 1 0-3.234a2.264 2.264 0 0 1 3.156 0l1.181 1.207l.005-.005l.005.005l1.18-1.207a2.264 2.264 0 0 1 3.157 0a2.308 2.308 0 0 1 0 3.234zM16 30h-2v-5a3.003 3.003 0 0 0-3-3H7a3.003 3.003 0 0 0-3 3v5H2v-5a5.006 5.006 0 0 1 5-5h4a5.006 5.006 0 0 1 5 5zM9 10a3 3 0 1 1-3 3a3 3 0 0 1 3-3m0-2a5 5 0 1 0 5 5a5 5 0 0 0-5-5z"/>`),
		g.Group(children),
	)
}