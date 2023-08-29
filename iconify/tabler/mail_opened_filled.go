package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MailOpenedFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M0 0h24v24H0z"/><path fill="currentColor" d="m14.872 14.287l6.522 6.52a2.996 2.996 0 0 1-2.218 1.188L19 22H5a2.995 2.995 0 0 1-2.394-1.191l6.521-6.522l2.318 1.545l.116.066a1 1 0 0 0 .878 0l.116-.066l2.317-1.545zM2 9.535l5.429 3.62L2 18.585zm20 0v9.05l-5.43-5.43zm-9.56-7.433l.115.066l8.444 5.629l-8.999 6l-9-6l8.445-5.63a1 1 0 0 1 .994-.065z"/></g>`),
		g.Group(children),
	)
}