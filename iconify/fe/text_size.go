package fe

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TextSize(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g id="feTextSize0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feTextSize1" fill="currentColor"><path id="feTextSize2" d="m13.5 16.494l2.408-7.224h2.182L21 17.995h-1.968l-.569-1.927h-2.966l-.643 1.927h-.853l.002.005h-2.707l-.782-2.65h-4.08L5.55 18H3L7 6h3l3.5 10.494ZM7 13h3L8.496 9L7 13Zm8.908 1.36h2.182l-1.094-2.909l-1.088 2.909Z"/></g></g>`),
		g.Group(children),
	)
}