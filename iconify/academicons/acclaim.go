package academicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Acclaim(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 448 512"),
		g.Raw(`<path fill="currentColor" d="M192 64C176.376 74.86 34.879 208.266 27.429 221.373v72.22c10.325.253 147.565-143.552 164.571-150.382c17.006 6.83 154.246 150.635 164.571 150.382v-72.22C349.121 208.266 207.624 74.859 192 64Zm0 154.406C176.376 229.266 34.879 362.672 27.429 375.78V448c10.325.254 147.565-143.551 164.571-150.383C209.006 304.448 346.246 448.253 356.571 448v-72.221C349.121 362.672 207.624 229.266 192 218.406Z"/>`),
		g.Group(children),
	)
}