package basil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ColumnsSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M10.542 19.267a74.67 74.67 0 0 0 0-14.534c-.043-.435-.34-.8-.753-.935a8.258 8.258 0 0 0-5.078 0c-.412.134-.71.5-.753.935a74.662 74.662 0 0 0 0 14.534c.043.434.341.8.753.934a8.259 8.259 0 0 0 5.078 0c.412-.133.71-.5.753-.934Zm9.5 0a74.67 74.67 0 0 0 0-14.534c-.043-.435-.34-.8-.753-.935a8.258 8.258 0 0 0-5.078 0c-.412.134-.71.5-.753.935a74.67 74.67 0 0 0 0 14.534c.043.434.341.8.753.934a8.259 8.259 0 0 0 5.078 0c.412-.133.71-.5.753-.934Z"/>`),
		g.Group(children),
	)
}