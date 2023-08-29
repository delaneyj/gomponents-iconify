package simple_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Powerautomate(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m20.079 7.235l3.666 4.097c.34.38.34.956 0 1.336l-7.63 8.528c-.19.213-.462.334-.748.334H7.289l12.79-14.295ZM6.789 21.53H1.005c-.867 0-1.326-1.025-.748-1.671L15.748 2.545c.139.057.265.145.367.259l3.631 4.058a.478.478 0 0 0-.039.039L6.916 21.197a.497.497 0 0 0-.127.333Zm8.356-19.06l-8.192 9.155L.257 4.141c-.578-.646-.119-1.671.748-1.671h14.14Z"/>`),
		g.Group(children),
	)
}