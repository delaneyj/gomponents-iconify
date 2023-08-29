package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PaletteThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M198 56.74A99.31 99.31 0 0 0 128 28h-1a100 100 0 0 0-99 100c0 41.22 25.55 75.85 66.69 90.38a28.34 28.34 0 0 0 9.42 1.63A28 28 0 0 0 132 192a20 20 0 0 1 20-20h46.21a27.84 27.84 0 0 0 27.3-21.76a100.37 100.37 0 0 0 2.49-23.1a99.26 99.26 0 0 0-30-70.4Zm19.74 91.72A19.89 19.89 0 0 1 198.21 164H152a28 28 0 0 0-28 28a20 20 0 0 1-26.64 18.83C59.51 197.46 36 165.72 36 128a92 92 0 0 1 91.05-92h.95a92 92 0 0 1 89.72 112.46ZM136 76a8 8 0 1 1-8-8a8 8 0 0 1 8 8Zm-44 24a8 8 0 1 1-8-8a8 8 0 0 1 8 8Zm0 56a8 8 0 1 1-8-8a8 8 0 0 1 8 8Zm88-56a8 8 0 1 1-8-8a8 8 0 0 1 8 8Z"/>`),
		g.Group(children),
	)
}