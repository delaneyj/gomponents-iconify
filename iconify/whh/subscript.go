package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Subscript(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M992.31 896h-160v64h160q13 0 22.5 9.5t9.5 22.5t-9.5 22.5t-22.5 9.5h-192q-12 0-20.5-7.5t-10.5-18.5V858q2-11 10.5-18.5t20.5-7.5h160v-64h-160q-13 0-22.5-9.5t-9.5-22.5t9.5-22.5t22.5-9.5h192q12 0 20.5 7.5t10.5 18.5v140q-2 11-10.5 18.5t-20.5 7.5zm-369.5-145.5q-16.5 16.5-39.5 16.5t-40-16l-223-278l-224 278q-16 16-39.5 16t-40-16.5T.31 711t17-40l231-287l-231-288q-17-16-17-39.5t16.5-40t40-16.5t39.5 17l224 277l223-277q17-17 40-17t39.5 16.5t16.5 40t-16 39.5l-231 288l231 287q16 17 16 40t-16.5 39.5z"/>`),
		g.Group(children),
	)
}