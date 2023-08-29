package fa_6_brands

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Yahoo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 448 512"),
		g.Raw(`<path fill="currentColor" d="M223.69 141.06L167 284.23l-56-143.17H14.93l105.83 249.13L82.19 480h94.17l140.91-338.94Zm105.4 135.79a58.22 58.22 0 1 0 58.22 58.22a58.22 58.22 0 0 0-58.22-58.22ZM394.65 32l-93 223.47h104.79L499.07 32Z"/>`),
		g.Group(children),
	)
}