package icons_8

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GpsDevice(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="m16 4.438l-.906 2.187l-8 19l-.907 2.125l2.157-.813L16 24.063l7.656 2.875l2.157.813l-.907-2.124l-8-19L16 4.436zm0 5.093l6.188 14.72l-5.844-2.187l-.344-.125l-.344.125l-5.844 2.188L16 9.53z"/>`),
		g.Group(children),
	)
}