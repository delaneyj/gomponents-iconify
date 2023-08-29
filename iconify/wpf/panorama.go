package wpf

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Panorama(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<path fill="currentColor" d="M.875 3A1 1 0 0 0 0 4v18a1 1 0 0 0 1.375.938S6.245 21 13 21s11.625 1.938 11.625 1.938A1 1 0 0 0 26 22V4a1 1 0 0 0-1.438-.906S20.826 5 13 5C5.174 5 1.437 3.094 1.437 3.094A1 1 0 0 0 .875 3zM2 5.375c.785.301 2.063.719 4 1.063v13.187c-1.869.35-3.222.767-4 1.031V5.375zm22 0v15.281c-.782-.265-2.13-.68-4-1.031V6.437c1.937-.343 3.216-.761 4-1.062zM8 6.719C9.407 6.878 11.03 7 13 7c1.97 0 3.593-.122 5-.281v12.593A39.245 39.245 0 0 0 13 19c-1.898 0-3.55.128-5 .313V6.719z"/>`),
		g.Group(children),
	)
}