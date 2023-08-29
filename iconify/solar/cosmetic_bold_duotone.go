package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CosmeticBoldDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path d="M2 11.5a.5.5 0 0 1 .5-.5h5a.5.5 0 0 1 .5.5V18a3 3 0 1 1-6 0v-6.5Z" opacity=".5"/><path d="M3 11h4V6a1 1 0 0 0-1.447-.895l-2 1A1 1 0 0 0 3 7v4Zm8-.5a5.5 5.5 0 1 1 11 0a5.5 5.5 0 0 1-11 0Z"/><path d="M15.75 15.95v3.55H13.5a.75.75 0 0 0 0 1.5h6a.75.75 0 0 0 0-1.5h-2.25v-3.55a5.539 5.539 0 0 1-1.5 0Z" opacity=".5"/></g>`),
		g.Group(children),
	)
}