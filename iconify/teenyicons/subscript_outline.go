package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SubscriptOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="m12.5 14.5l-.354-.354A.5.5 0 0 0 12.5 15v-.5ZM15 14h-2.5v1H15v-1Zm-2.146.854l1.792-1.793l-.707-.707l-1.793 1.792l.708.708ZM13.793 11H13.5v1h.293v-1Zm-.293 0a1.5 1.5 0 0 0-1.5 1.5h1a.5.5 0 0 1 .5-.5v-1Zm1.5 1.207C15 11.54 14.46 11 13.793 11v1c.114 0 .207.093.207.207h1Zm-.354.854c.227-.227.354-.534.354-.854h-1a.207.207 0 0 1-.06.147l.706.707ZM1.9 13.8l9-12l-.8-.6l-9 12l.8.6Zm-.8-12l9 12l.8-.6l-9-12l-.8.6Z"/>`),
		g.Group(children),
	)
}