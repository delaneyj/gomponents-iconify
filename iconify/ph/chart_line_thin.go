package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChartLineThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M228 208a4 4 0 0 1-4 4H32a4 4 0 0 1-4-4V48a4 4 0 0 1 8 0v103.19L93.37 101a4 4 0 0 1 5-.19l61.41 46.05L221.37 93a4 4 0 0 1 5.26 6l-64 56a4 4 0 0 1-5 .19l-61.41-46L36 161.81V204h188a4 4 0 0 1 4 4Z"/>`),
		g.Group(children),
	)
}