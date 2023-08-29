package fa_6_brands

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CssThree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 448 512"),
		g.Raw(`<path fill="currentColor" d="m480 32l-64 368l-223.3 80L0 400l19.6-94.8h82l-8 40.6L210 390.2l134.1-44.4l18.8-97.1H29.5l16-82h333.7l10.5-52.7H56.3l16.3-82H480z"/>`),
		g.Group(children),
	)
}