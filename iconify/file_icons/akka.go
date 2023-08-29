package file_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Akka(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M497.415 321.528c38.317 64.495-4.645 164.82-61.901 116.147c-111.172-94.397-295.017-24.019-297.357-23.192c111.96-130.69 290.54-195.023 359.258-92.955zm10.908-17.466L396.64 78.704c-15.972-25.51-56.79-20.187-79.861-.665L21.292 325.609c-26.842 23.072-28.617 63.89-3.55 88.957c21.963 21.962 65.307 28.764 98.563-8.15C264.192 232.65 447.668 213.292 508.323 304.063z"/>`),
		g.Group(children),
	)
}