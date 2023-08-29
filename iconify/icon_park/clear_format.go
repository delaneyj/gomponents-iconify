package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ClearFormat(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000"><path fill="#2F88FF" stroke-linejoin="round" stroke-width="4.302" d="M44.7818 24.1702L31.918 7.09935L14.1348 20.5L27.5 37L30.8556 34.6643L44.7818 24.1702Z"/><path stroke-linejoin="round" stroke-width="4.302" d="M27.4998 37L23.6613 40.0748L13.0978 40.074L10.4973 36.6231L4.06543 28.0876L14.4998 20.2248"/><path stroke-linecap="round" stroke-width="4.5" d="M13.2056 40.072L44.5653 40.072"/></g>`),
		g.Group(children),
	)
}