package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Flatbread(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<circle cx="36" cy="36" r="23" fill="#fcea2b"/><path fill="#f1b31c" d="M54.695 22.61A23 23 0 0 1 14.02 42.792A23.001 23.001 0 1 0 54.695 22.61Z"/><g fill="none" stroke="#000" stroke-miterlimit="10" stroke-width="2"><circle cx="36" cy="36" r="23"/><path stroke-linecap="round" d="M46.264 25.894L25.051 47.107m22.984-12.371L33.893 48.879m-6.791-13.863l-3.536 3.535m14.142-14.142l-7.071 7.071m-1.484-8.555l-7.071 7.071"/></g>`),
		g.Group(children),
	)
}