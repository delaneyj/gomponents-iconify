package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PreviewClose(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M6 16C6.63472 17.2193 7.59646 18.3504 8.82276 19.3554C12.261 22.1733 17.779 24 24 24C30.221 24 35.739 22.1733 39.1772 19.3554C40.4035 18.3504 41.3653 17.2193 42 16"/><path d="M28.9775 24L31.048 31.7274"/><path d="M37.3535 21.3536L43.0103 27.0104"/><path d="M5.00004 27.0103L10.6569 21.3534"/><path d="M16.9278 31.7276L18.9983 24.0001"/></g>`),
		g.Group(children),
	)
}