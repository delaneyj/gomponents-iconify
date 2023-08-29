package guidance

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Hairdresser(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" d="m12 16.61l-4.276 4.589C7 22 6.102 22.5 5.002 22.5a3.5 3.5 0 1 1 3.162-5M12 16.61l9.641-10.35a3.2 3.2 0 0 0-.078-4.444l-.316-.316h-.25l-.106.133A55.371 55.371 0 0 1 8.6 12.961L2.36 6.26a3.2 3.2 0 0 1 .078-4.444l.316-.316M12 16.61l4.276 4.589C17 22 17.898 22.5 18.998 22.5a3.5 3.5 0 1 0-3.162-5M2.753 1.5h.079m-.08 0h.25l.107.133A55.374 55.374 0 0 0 12 10.477"/>`),
		g.Group(children),
	)
}