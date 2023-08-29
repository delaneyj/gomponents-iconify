package entypo

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Palette(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M15.74 2.608c-3.528-1.186-7.066-.961-10.72 1.274C2.167 5.625.302 9.958.917 13.064c.728 3.671 4.351 5.995 9.243 4.651c5.275-1.449 6.549-4.546 6.379-5.334c-.17-.788-2.665-1.652-1.718-3.498c1.188-2.313 3.129-1.149 3.982-1.622c.855-.472.539-3.442-3.063-4.653zm-3.646 10.706a1.504 1.504 0 0 1-1.843-1.059a1.5 1.5 0 0 1 1.046-1.849a1.503 1.503 0 0 1 1.843 1.059a1.501 1.501 0 0 1-1.046 1.849z"/>`),
		g.Group(children),
	)
}