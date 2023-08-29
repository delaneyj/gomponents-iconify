package fa_brands

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Maxcdn(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 448 512"),
		g.Raw(`<path fill="currentColor" d="M461.1 442.7h-97.4L415.6 200c2.3-10.2.9-19.5-4.4-25.7c-5-6.1-13.7-9.6-24.2-9.6h-49.3l-59.5 278h-97.4l59.5-278h-83.4l-59.5 278H0l59.5-278l-44.6-95.4H387c39.4 0 75.3 16.3 98.3 44.9c23.3 28.6 31.8 67.4 23.6 105.9l-47.8 222.6z"/>`),
		g.Group(children),
	)
}