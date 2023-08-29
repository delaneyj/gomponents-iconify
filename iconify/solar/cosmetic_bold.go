package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CosmeticBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 12.5a.5.5 0 0 1 .5-.5h5a.5.5 0 0 1 .5.5V18a3 3 0 1 1-6 0v-5.5Zm1-2h4V6a1 1 0 0 0-1.447-.895l-2 1A1 1 0 0 0 3 7v3.5Zm8 0a5.5 5.5 0 1 1 11 0a5.5 5.5 0 0 1-11 0Zm4.75 7.21a7.091 7.091 0 0 0 1.5 0v1.79h2.25a.75.75 0 0 1 0 1.5h-6a.75.75 0 0 1 0-1.5h2.25v-1.79Z"/>`),
		g.Group(children),
	)
}