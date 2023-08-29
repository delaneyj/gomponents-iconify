package bx

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Run(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<circle cx="17" cy="4" r="2" fill="currentColor"/><path fill="currentColor" d="M15.777 10.969a2.007 2.007 0 0 0 2.148.83l3.316-.829l-.483-1.94l-3.316.829l-1.379-2.067a2.01 2.01 0 0 0-1.272-.854l-3.846-.77a1.998 1.998 0 0 0-2.181 1.067l-1.658 3.316l1.789.895l1.658-3.317l1.967.394L7.434 17H3v2h4.434c.698 0 1.355-.372 1.715-.971l1.918-3.196l5.169 1.034l1.816 5.449l1.896-.633l-1.815-5.448a2.007 2.007 0 0 0-1.506-1.33l-3.039-.607l1.772-2.954l.417.625z"/>`),
		g.Group(children),
	)
}