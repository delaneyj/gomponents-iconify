package covid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VirusLabResearchMagnifierTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M10.617 13.837a3.22 3.22 0 1 0 0-6.44a3.22 3.22 0 0 0 0 6.44Zm-.537-8.855h1.074m-.537 0v2.415m3.605-1.144l.759.76m-.38-.38L12.894 8.34m3.358 1.74v1.074m0-.537h-2.415m1.144 3.605l-.759.759m.379-.38l-1.707-1.707m-1.74 3.358H10.08m.537 0v-2.415m-3.605 1.144l-.759-.759m.38.379l1.707-1.707m-3.358-1.74V10.08m0 .537h2.415M6.253 7.012l.759-.759m-.379.38L8.34 8.34"/><path d="M10.617 20.484c5.45 0 9.867-4.418 9.867-9.867c0-5.45-4.418-9.867-9.867-9.867C5.167.75.75 5.168.75 10.617c0 5.45 4.418 9.867 9.867 9.867ZM23.25 23.25l-5.656-5.656"/></g>`),
		g.Group(children),
	)
}