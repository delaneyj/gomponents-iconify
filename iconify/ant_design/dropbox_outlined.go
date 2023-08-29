package ant_design

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DropboxOutlined(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="m64 556.9l264.2 173.5L512.5 577L246.8 412.7zm896-290.3zm0 0L696.8 95L512.5 248.5l265.2 164.2L512.5 577l184.3 153.4L960 558.8L777.7 412.7zM513 609.8L328.2 763.3l-79.4-51.5v57.8L513 928l263.7-158.4v-57.8l-78.9 51.5zM328.2 95L64 265.1l182.8 147.6l265.7-164.2zM64 556.9z"/>`),
		g.Group(children),
	)
}