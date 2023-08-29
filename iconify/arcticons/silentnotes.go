package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Silentnotes(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<ellipse cx="19.05" cy="23.636" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="4.95" ry="5.508"/><circle cx="19.592" cy="24.691" r="2.014" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><ellipse cx="28.95" cy="23.636" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="4.95" ry="5.508"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M26.72 30.259c-.624 1.337-1.412 1.783-2.72 1.783c-1.308 0-2.096-.446-2.72-1.784m9.655-11.668L24 3.5l-6.935 15.09"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M31.613 28.277V44.5H16.387V28.277"/><circle cx="28.408" cy="24.691" r="2.014" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M27.078 10.175s-.922.49-3.078.49c-2.155 0-3.078-.49-3.078-.49"/>`),
		g.Group(children),
	)
}