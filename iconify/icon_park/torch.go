package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Torch(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="2" stroke-width="4"><path d="M27 34H21V44H27V34Z"/><path d="M31 34H17L14 26H34L31 34Z"/><path fill="#2F88FF" d="M20.9996 14C22.9996 9.39 26.9996 4 26.9996 4C26.9996 4 33.3596 15.49 33.3596 19.23C33.3596 22.97 30.7396 26 26.9996 26H18.9996C16.3596 26 14.6396 23.86 14.6396 21.22C14.6396 18.58 18.9996 10 18.9996 10C18.9996 10 19.9596 12.01 20.9996 14Z"/></g>`),
		g.Group(children),
	)
}