package map_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RoutePin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 50 50"),
		g.Raw(`<path fill="currentColor" d="M47.495 21.575c-.225-12.158-7.278-16.989-7.567-17.18l-1.091-.746l-1.099.744c-1.826 1.239-3.719 1.868-5.626 1.868c-3.158 0-5.302-1.7-5.393-1.774l-1.23-.987l-1.221.99c-.09.072-2.233 1.771-5.392 1.771c-1.905 0-3.797-.628-5.626-1.871l-1.094-.738l-1.102.745c-.748.514-7.314 5.364-7.552 17.118c-.095 1.024 2.729 20.374 21.997 25.985c20.948-6.122 22.104-24.648 21.996-25.925z"/>`),
		g.Group(children),
	)
}