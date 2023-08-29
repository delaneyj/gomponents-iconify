package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Eyebrow(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSEyebrow0"><g fill="none" stroke-linejoin="round" stroke-width="4"><path fill="#fff" stroke="#fff" d="M24 40c9.941 0 18-10 18-10s-8.059-10-18-10S6 30 6 30s8.059 10 18 10Z"/><path fill="#000" stroke="#000" d="M24 34a4 4 0 1 0 0-8a4 4 0 0 0 0 8Z"/><path fill="#fff" stroke="#fff" d="M28 6c-7 0-18 3.5-21 6s-1 7 1 6s15.2-5.8 20-7c4.8-1.2 11.667.833 14 2c-2.333-2-7-7-14-7Z"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSEyebrow0)"/>`),
		g.Group(children),
	)
}