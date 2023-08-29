package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CircleFiveLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSCircleFiveLine0"><g fill="none" stroke="#fff" stroke-linejoin="round" stroke-width="4"><path fill="#fff" d="M40 12a4 4 0 1 0 0-8a4 4 0 0 0 0 8ZM24 28a4 4 0 1 0 0-8a4 4 0 0 0 0 8Zm16 16a4 4 0 1 0 0-8a4 4 0 0 0 0 8ZM8 12a4 4 0 1 0 0-8a4 4 0 0 0 0 8Zm0 32a4 4 0 1 0 0-8a4 4 0 0 0 0 8Z"/><path stroke-linecap="round" d="M20 40h8M20 8h8M8 20v8m32-8v8"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSCircleFiveLine0)"/>`),
		g.Group(children),
	)
}