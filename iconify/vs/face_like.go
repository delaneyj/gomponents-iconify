package vs

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FaceLike(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1792 1792"),
		g.Raw(`<path fill="currentColor" d="M832 0q169 0 323 66t265.5 177.5T1598 509t66 323t-66 323.5t-177.5 265.5t-265.5 177t-323 66t-323-66t-265.5-177T66 1155.5T0 832t66-323t177.5-265.5T509 66T832 0zM384 640q0 53 37.5 90.5T512 768t90.5-37.5T640 640t-37.5-90.5T512 512t-90.5 37.5T384 640zm448 704q137 0 243.5-89.5T1216 1025H448q34 141 140 230t244 89zm320-576q53 0 90.5-37.5T1280 640t-37.5-90.5T1152 512t-90.5 37.5T1024 640t37.5 90.5T1152 768z"/>`),
		g.Group(children),
	)
}