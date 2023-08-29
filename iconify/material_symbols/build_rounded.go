package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BuildRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M17.875 20.375q-.5 0-.95-.188t-.825-.562l-5-5.025q-.5.2-1.012.3T9 15q-2.5 0-4.25-1.75T3 9q0-.9.25-1.712t.7-1.538L7.6 9.4l1.8-1.8l-3.65-3.65q.725-.45 1.538-.7T9 3q2.5 0 4.25 1.75T15 9q0 .575-.1 1.088t-.3 1.012l5.05 5q.375.375.563.825t.187.95q0 .5-.2.963t-.55.812q-.375.375-.825.55t-.95.175Z"/>`),
		g.Group(children),
	)
}