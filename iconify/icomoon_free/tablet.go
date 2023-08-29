package icomoon_free

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Tablet(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M12.5 0h-10C1.675 0 1 .675 1 1.5v13c0 .825.675 1.5 1.5 1.5h10c.825 0 1.5-.675 1.5-1.5v-13c0-.825-.675-1.5-1.5-1.5zm-5 15.5a.5.5 0 1 1 0-1a.5.5 0 0 1 0 1zM12 14H3V2h9v12z"/>`),
		g.Group(children),
	)
}