package fontisto

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Save(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M23.681 6.158L17.843.32a1.093 1.093 0 0 0-.771-.32H1.092C.489 0 .001.489.001 1.091v21.817c0 .603.489 1.091 1.091 1.091h21.817c.603 0 1.091-.489 1.091-1.091V6.928c0-.301-.122-.574-.32-.771zM6.549 2.182h6.546v5.819H6.546zm0 19.635v-5.818h10.905v5.818zm15.273 0h-2.185v-6.908c0-.603-.489-1.091-1.091-1.091H5.455c-.603 0-1.091.489-1.091 1.091v6.909H2.182V2.181h2.182V9.09c0 .603.489 1.091 1.091 1.091h8.728c.603 0 1.091-.489 1.091-1.091V2.181h1.344l5.199 5.199z"/>`),
		g.Group(children),
	)
}