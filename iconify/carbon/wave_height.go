package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WaveHeight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M24 2h6v2h-6zm0 6h4v2h-4zm0 6h6v2h-6zm0 6h4v2h-4z"/><path fill="currentColor" d="M30 28h-6a10.035 10.035 0 0 1-6.927-17.262a11.963 11.963 0 0 0-4.08-.738a6.903 6.903 0 0 0-6.03 3.42C4.997 16.435 4 21.34 4 28H2c0-7.054 1.106-12.327 3.287-15.673A8.906 8.906 0 0 1 12.994 8H13a14.762 14.762 0 0 1 6.461 1.592a1 1 0 0 1 .087 1.722A8.025 8.025 0 0 0 24 26h6Z"/>`),
		g.Group(children),
	)
}