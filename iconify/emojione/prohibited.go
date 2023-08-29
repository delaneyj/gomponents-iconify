package emojione

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Prohibited(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="#ff5a79" d="M32 2C15.4 2 2 15.4 2 32s13.4 30 30 30s30-13.4 30-30S48.6 2 32 2m22 30c0 4.6-1.4 8.9-3.9 12.5L19.5 13.9C23.1 11.4 27.4 10 32 10c12.2 0 22 9.9 22 22m-44 0c0-4.6 1.4-8.9 3.9-12.5l30.6 30.6C40.9 52.6 36.6 54 32 54c-12.1 0-22-9.9-22-22"/>`),
		g.Group(children),
	)
}