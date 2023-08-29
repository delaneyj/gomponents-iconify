package basil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RowsSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M19.267 3.958a74.662 74.662 0 0 0-14.534 0c-.435.043-.8.34-.934.753a8.258 8.258 0 0 0 0 5.078c.133.412.5.71.934.753c4.833.472 9.7.472 14.534 0c.435-.043.8-.341.934-.753a8.259 8.259 0 0 0 0-5.078c-.133-.412-.5-.71-.934-.753Zm0 9.5a74.67 74.67 0 0 0-14.534 0c-.435.043-.8.34-.934.753a8.258 8.258 0 0 0 0 5.078c.133.412.5.71.934.753c4.833.472 9.7.472 14.534 0c.435-.043.8-.341.934-.753a8.259 8.259 0 0 0 0-5.078c-.133-.412-.5-.71-.934-.753Z"/>`),
		g.Group(children),
	)
}