package icomoon_free

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Eyedropper(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M15.414.586a2 2 0 0 0-2.828 0L9.897 3.275L8.001 1.379L5.88 3.5l1.663 1.663L.166 12.54a.56.56 0 0 0-.161.46H.001v2.5a.5.5 0 0 0 .5.5h2.563a.561.561 0 0 0 .398-.165l7.377-7.377l1.663 1.663L14.623 8l-1.896-1.896l2.689-2.689a2 2 0 0 0 0-2.828zM2.705 15H1v-1.705l7.337-7.337l1.704 1.704l-7.337 7.337z"/>`),
		g.Group(children),
	)
}