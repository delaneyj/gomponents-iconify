package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InvalidFiles(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSInvalidFiles0"><g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path fill="#fff" stroke="#fff" d="M8 44V4h23l9 10.5V44H8Z"/><path fill="#000" stroke="#000" d="M34 25c0 5.523-4.477 10-10 10s-10-4.477-10-10s4.477-10 10-10a9.965 9.965 0 0 1 6.865 2.729A9.972 9.972 0 0 1 34 25Z"/><path stroke="#000" d="m17 18l14 14m3-7c0 5.523-4.477 10-10 10M14 25c0-5.523 4.477-10 10-10"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSInvalidFiles0)"/>`),
		g.Group(children),
	)
}