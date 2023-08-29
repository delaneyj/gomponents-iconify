package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CircleFourLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSCircleFourLine0"><g fill="none"><path fill="#fff" stroke="#fff" stroke-linejoin="round" stroke-width="4" d="M40 12a4 4 0 1 0 0-8a4 4 0 0 0 0 8Zm0 32a4 4 0 1 0 0-8a4 4 0 0 0 0 8ZM8 12a4 4 0 1 0 0-8a4 4 0 0 0 0 8Zm0 32a4 4 0 1 0 0-8a4 4 0 0 0 0 8Z"/><path stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M20 8h8"/><path fill="#fff" fill-rule="evenodd" d="M32 16L16 32l16-16Z" clip-rule="evenodd"/><path stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M32 16L16 32m24-12v8"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSCircleFourLine0)"/>`),
		g.Group(children),
	)
}