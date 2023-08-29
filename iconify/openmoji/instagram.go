package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Instagram(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<path fill="#FFF" d="M46.785 12.614h-21.57c-6.96 0-12.601 5.642-12.601 12.601v21.57c0 6.96 5.642 12.601 12.602 12.601h21.568c6.96 0 12.602-5.642 12.602-12.602V25.215c0-6.96-5.642-12.601-12.602-12.601z"/><path fill="none" stroke="#000" stroke-miterlimit="10" stroke-width="2" d="M46.785 12.614h-21.57c-6.96 0-12.601 5.642-12.601 12.601v21.57c0 6.96 5.642 12.601 12.602 12.601h21.568c6.96 0 12.602-5.642 12.602-12.602V25.215c0-6.96-5.642-12.601-12.602-12.601z"/><circle cx="36" cy="36" r="11.29" fill="none" stroke="#000" stroke-miterlimit="10" stroke-width="2"/><circle cx="49.709" cy="22.291" r="2.419"/>`),
		g.Group(children),
	)
}