package gridicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bell(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m6.14 14.969l2.828 2.828a2 2 0 1 1-2.828-2.828zm8.867 5.325l-.706.706L3 9.699l.706-.706l1.102.157c.754.108 1.689-.122 2.077-.51l3.885-3.884a5.993 5.993 0 0 1 8.475 8.475l-3.885 3.885c-.388.388-.618 1.323-.51 2.077l.157 1.101z"/>`),
		g.Group(children),
	)
}