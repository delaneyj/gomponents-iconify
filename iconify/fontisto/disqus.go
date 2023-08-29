package fontisto

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Disqus(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12.807 24h-.014a11.742 11.742 0 0 1-7.699-2.861l.015.013l-5.109.7l1.974-4.874a11.83 11.83 0 0 1-1.073-4.964C.901 5.416 6.22.06 12.804-.001h.006c6.627 0 12 5.373 12 12s-5.373 12-12 12zm6.503-12.034v-.034c0-3.463-2.443-5.931-6.654-5.931H8.111v12.002h4.479c4.242 0 6.719-2.574 6.719-6.037zm-6.606 3.089h-1.328V8.952h1.328a2.92 2.92 0 0 1 3.256 3.038v-.006v.031a2.914 2.914 0 0 1-3.264 3.036l.014.001z"/>`),
		g.Group(children),
	)
}