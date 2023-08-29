package fa_brands

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Deviantart(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 448 512"),
		g.Raw(`<path fill="currentColor" d="m320 93.2l-98.2 179.1l7.4 9.5H320v127.7H159.1l-13.5 9.2l-43.7 84c-.3 0-8.6 8.6-9.2 9.2H0v-93.2l93.2-179.4l-7.4-9.2H0V102.5h156l13.5-9.2l43.7-84c.3 0 8.6-8.6 9.2-9.2H320v93.1z"/>`),
		g.Group(children),
	)
}