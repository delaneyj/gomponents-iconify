package ci

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HomeAltCheck(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M19 22H5a1 1 0 0 1-1-1v-9.643a1.01 1.01 0 0 1 .293-.65l7-7a1 1 0 0 1 1.415 0l7 7a.994.994 0 0 1 .292.707V21a1 1 0 0 1-1 1ZM12 5.828l-6 6V20h12v-8.172l-6-6Zm-1.5 12.731l-2.706-2.7L9.2 14.441l1.3 1.292l4.3-4.292l1.412 1.416l-5.712 5.701v.001Z"/>`),
		g.Group(children),
	)
}