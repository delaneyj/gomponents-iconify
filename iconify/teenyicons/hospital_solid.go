package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HospitalSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M7.724.053a.5.5 0 0 0-.448 0l-5.99 2.995A.5.5 0 0 0 1 3.5V14H0v1h5v-5h5v5h5v-1h-1V3.5a.5.5 0 0 0-.286-.452L7.724.053ZM7 5V3h1v2h2v1H8v2H7V6H5V5h2Z" clip-rule="evenodd"/><path fill="currentColor" d="M9 15v-4H6v4h3Z"/>`),
		g.Group(children),
	)
}