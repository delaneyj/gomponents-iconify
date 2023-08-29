package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowKeys(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSArrowKeys0"><g fill="none"><path fill="#fff" d="M14 4h20v20H14V4ZM4 24h20v20H4V24Zm20 0h20v20H24V24Z"/><path stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M24 24H4v20h20V24Zm0 0v20v-20Zm0 0h20v20H24V24ZM14 4h20v20H14V4Z"/><path stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M24 10v8m-4-4l4-4l4 4M10 34h8m-4 4l-4-4l4-4m24 4h-8m4-4l4 4l-4 4"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSArrowKeys0)"/>`),
		g.Group(children),
	)
}