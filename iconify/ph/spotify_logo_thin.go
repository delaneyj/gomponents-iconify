package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SpotifyLogoThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M128 28a100 100 0 1 0 100 100A100.11 100.11 0 0 0 128 28Zm0 192a92 92 0 1 1 92-92a92.1 92.1 0 0 1-92 92Zm27.54-52.13a4 4 0 0 1-5.41 1.67a46.73 46.73 0 0 0-44.26 0a4.07 4.07 0 0 1-1.87.46a4 4 0 0 1-1.87-7.54a55.39 55.39 0 0 1 51.74 0a4 4 0 0 1 1.67 5.41Zm32-56a4 4 0 0 1-5.41 1.65a114.67 114.67 0 0 0-108.24 0a4 4 0 0 1-3.76-7.06a122.65 122.65 0 0 1 115.76 0a4 4 0 0 1 1.64 5.41Zm-16 28a4 4 0 0 1-5.41 1.65a81 81 0 0 0-76.24 0a4 4 0 0 1-3.76-7.06a89 89 0 0 1 83.76 0a4 4 0 0 1 1.64 5.42Z"/>`),
		g.Group(children),
	)
}