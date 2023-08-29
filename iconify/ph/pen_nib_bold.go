package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PenNibBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M244 100.68a19.86 19.86 0 0 0-5.86-14.14l-68.68-68.69a20 20 0 0 0-28.29 0l-27.82 27.82l-57.24 21.47A20 20 0 0 0 43.4 82.58L20.16 222A12 12 0 0 0 32 236a11.9 11.9 0 0 0 2-.16l139.45-23.25a20.07 20.07 0 0 0 15.44-12.7l21.46-57.25l27.82-27.82a19.85 19.85 0 0 0 5.83-14.14Zm-76.8 88.62L67 206l33-33a34 34 0 1 0-17-17l-33 33L66.7 88.8L117 69.94L186.05 139ZM104 142a10 10 0 1 1 10 10a10 10 0 0 1-10-10Zm96-23l-63-63l18.34-18.34l63 63Z"/>`),
		g.Group(children),
	)
}