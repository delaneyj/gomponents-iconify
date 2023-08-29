package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Razor(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSRazor0"><g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><rect width="38" height="20" x="5" y="5" fill="#fff" stroke="#fff" rx="2"/><path stroke="#000" d="M14 5v20M34 5v20"/><path stroke="#fff" d="M30 5h8"/><path stroke="#000" d="M5 11h38M5 19h38"/><path stroke="#fff" d="M5 21V9m38 12V9M10 5h8m12 20h8m-28 0h8m14 0a8 8 0 1 1-16 0"/><path stroke="#fff" d="m28 32l-1 5v7m-7-12l1 5v7"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSRazor0)"/>`),
		g.Group(children),
	)
}