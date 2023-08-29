package icomoon_free

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Icomoon(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M4.055 8a1.851 1.851 0 1 1 3.703 0a1.851 1.851 0 0 1-3.703 0zM8 0a8 8 0 1 0 0 16A8 8 0 0 0 8 0zM5.928 14.989C3.522 13.589 1.905 10.984 1.905 8s1.617-5.589 4.023-6.989C8.334 2.41 9.953 5.016 9.953 8s-1.618 5.589-4.025 6.989z"/>`),
		g.Group(children),
	)
}