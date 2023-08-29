package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Sandwich(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSSandwich0"><g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path fill="#fff" stroke="#fff" d="M42 36V22H6v14a4 4 0 0 0 4 4h28a4 4 0 0 0 4-4Z"/><path stroke="#000" d="M6 30h36"/><path stroke="#fff" d="M6 26v8m36-8v8M6.067 22H42c0-2.016-11.807-10.582-17.967-14.613c-6.16-4.03-11.293 1.512-15.4 6.55C5.348 17.97 6.067 19.986 6.067 22Z"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSSandwich0)"/>`),
		g.Group(children),
	)
}