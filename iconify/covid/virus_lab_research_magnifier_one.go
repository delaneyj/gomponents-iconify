package covid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VirusLabResearchMagnifierOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M12.017.75H8.983m1.517 0v2.6m-5.822-.817L3.606 3.606L2.533 4.678m1.073-1.072l1.838 1.838M.75 8.983v3.034m0-1.517h2.6m-.817 5.822l1.073 1.072l1.072 1.073m-1.072-1.073l1.838-1.838M18.467 4.678l-1.073-1.072l-1.072-1.073m1.072 1.073l-1.838 1.838M8.983 9.85a1.517 1.517 0 1 0 0-3.034a1.517 1.517 0 0 0 0 3.034Z"/><path d="M17.435 8.755A7.15 7.15 0 1 0 8.5 17.361"/><path d="M16.264 21.154a4.89 4.89 0 1 0 0-9.78a4.89 4.89 0 0 0 0 9.78Zm6.986 2.096l-3.528-3.528"/></g>`),
		g.Group(children),
	)
}