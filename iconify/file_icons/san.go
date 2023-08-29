package file_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func San(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M159.594 373.61L0 405.94V512l424.344-85.966v-.53zM0 295.948v1.117l2.804-.568zM156.79 158.62L0 190.384l266.08 52.155l155.46-31.495v-.53zm267.554 30.614L0 106.059V0l424.344 83.176zM0 214.935l424.344 83.176v106.058L0 320.992z"/>`),
		g.Group(children),
	)
}