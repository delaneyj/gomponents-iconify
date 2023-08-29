package file_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Nginx(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="m512 236.329l-121.905 246.58h-262.65L0 259.048L126.892 29.09h273.731l72.59 170.112h-94.754l-37.126-75.359H182.857l-74.25 134.65l73.696 130.216h149.056l28.26-57.628H243.393c-24.13 0-47.746-18.68-47.746-46.822c0-26.166 26.087-47.931 47.746-47.931H512z"/>`),
		g.Group(children),
	)
}