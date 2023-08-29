package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MyNotes(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M39.488 23.42A10.283 10.283 0 0 0 24.946 8.879"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24.426 35.818a5.127 5.127 0 0 1 1.077 1.588L39.488 23.42L24.946 8.878L10.831 22.994a5.042 5.042 0 0 1 2.025 1.254c1.597 1.597 1.918 3.867.715 5.07c1.203-1.203 3.473-.883 5.07.715s1.918 3.867.715 5.07c1.203-1.203 3.472-.883 5.07.715Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M10.831 22.994L5.5 42.134l20.003-4.728"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M12.409 40.212a7.182 7.182 0 0 0-4.725-4.91"/>`),
		g.Group(children),
	)
}