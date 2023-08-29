package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Sidebar(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M896.428 1024h-768q-53 0-90.5-37.5T.428 896V128q0-53 37.5-90.5t90.5-37.5h768q53 0 90.5 37.5t37.5 90.5v768q0 53-37.5 90.5t-90.5 37.5zm-768-160q0 13 9.5 22.5t22.5 9.5h352V448h-384v416zm384-736h-352q-13 0-22.5 9.5t-9.5 22.5v160h384V128zm384 32q0-13-9.5-22.5t-22.5-9.5h-224v768h224q13 0 22.5-9.5t9.5-22.5V160z"/>`),
		g.Group(children),
	)
}