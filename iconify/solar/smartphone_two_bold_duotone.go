package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SmartphoneTwoBoldDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path d="M5.172 3.172C4 4.343 4 6.229 4 10v4c0 3.771 0 5.657 1.172 6.828C6.343 22 8.229 22 12 22c3.771 0 5.657 0 6.828-1.172C20 19.657 20 17.771 20 14v-4c0-3.771 0-5.657-1.172-6.828C17.657 2 15.771 2 12 2C8.229 2 6.343 2 5.172 3.172Z" opacity=".5"/><path d="M9 4.25a.75.75 0 0 0 0 1.5h6a.75.75 0 0 0 0-1.5H9ZM12 19a2 2 0 1 0 0-4a2 2 0 0 0 0 4Z"/></g>`),
		g.Group(children),
	)
}