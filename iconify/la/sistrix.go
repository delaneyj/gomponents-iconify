package la

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Sistrix(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M12.5 3C7.262 3 3 7.262 3 12.5S7.262 22 12.5 22a9.45 9.45 0 0 0 5.967-2.12l9.16 9.161l1.414-1.414l-9.16-9.16A9.45 9.45 0 0 0 22 12.5C22 7.262 17.738 3 12.5 3zm0 2c4.136 0 7.5 3.364 7.5 7.5S16.636 20 12.5 20S5 16.636 5 12.5S8.364 5 12.5 5z"/>`),
		g.Group(children),
	)
}