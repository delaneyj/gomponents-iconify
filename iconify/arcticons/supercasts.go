package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Supercasts(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m8.442 11.148l-4.387 4.407a1.917 1.917 0 0 0 .001 2.7l18.6 18.598a1.898 1.898 0 0 0 2.683.005l.006-.005l18.599-18.599a1.917 1.917 0 0 0 .001-2.7l-4.387-4.406a1.908 1.908 0 0 0-1.345-.56H9.791a1.908 1.908 0 0 0-1.345.56Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M17.894 27.363a8.595 8.595 0 0 1 12.113 0m-15.451-3.339a13.402 13.402 0 0 1 18.885 0"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M11.217 20.686a18.226 18.226 0 0 1 25.657 0"/>`),
		g.Group(children),
	)
}