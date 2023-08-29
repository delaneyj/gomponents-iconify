package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CollectionFiles(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSCollectionFiles0"><g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path fill="#fff" stroke="#fff" d="M8 44V4h23l9 10.5V44H8Z"/><path fill="#000" stroke="#000" d="m24 15l3.084 6.755l7.378.846l-5.472 5.02l1.476 7.278L24 31.247l-6.466 3.652l1.476-7.278l-5.472-5.02l7.378-.846L24 15Z"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSCollectionFiles0)"/>`),
		g.Group(children),
	)
}