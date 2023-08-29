package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SearchPropertySolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="M5 8V6.207l1.5-1.5l1.5 1.5V8H5Z"/><path fill="currentColor" fill-rule="evenodd" d="M0 6.5a6.5 6.5 0 1 1 11.436 4.23l3.418 3.416l-.707.707l-3.418-3.417A6.5 6.5 0 0 1 0 6.5Zm8.854-.854l-2-2a.5.5 0 0 0-.708 0l-2 2A.5.5 0 0 0 4 6v2.5a.5.5 0 0 0 .5.5h4a.5.5 0 0 0 .5-.5V6a.5.5 0 0 0-.146-.354Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}