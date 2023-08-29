package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PenNewRoundLineDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="M16.652 3.455s.081 1.379 1.298 2.595c1.216 1.217 2.595 1.298 2.595 1.298M10.1 15.588L8.413 13.9" opacity=".5"/><path d="m16.652 3.455l.649-.649A2.753 2.753 0 0 1 21.194 6.7l-.65.649l-5.964 5.965c-.404.404-.606.606-.829.78a4.59 4.59 0 0 1-.848.524c-.255.121-.526.211-1.068.392l-1.735.579l-1.123.374a.742.742 0 0 1-.939-.94l.374-1.122l.579-1.735c.18-.542.27-.813.392-1.068c.144-.301.32-.586.524-.848c.174-.223.376-.425.78-.83l5.965-5.964Z"/><path stroke-linecap="round" d="M22 12c0 5.523-4.477 10-10 10S2 17.523 2 12S6.477 2 12 2" opacity=".5"/></g>`),
		g.Group(children),
	)
}