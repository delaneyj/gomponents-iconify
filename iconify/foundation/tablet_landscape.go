package foundation

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TabletLandscape(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 100 100"),
		g.Raw(`<path fill="currentColor" d="M92.566 22.081a3.035 3.035 0 0 0-3.035-3.035H10.469a3.035 3.035 0 0 0-3.035 3.035c0 .043.011.084.013.127h-.013v55.838h.013a3.028 3.028 0 0 0 3.022 2.908h79.062a3.028 3.028 0 0 0 3.022-2.908h.013V22.208h-.013c.002-.044.013-.084.013-.127zm-9.957 48.916H17.391V29.003H82.61v41.994zm4.978-18.386a2.484 2.484 0 1 1 0-4.968a2.484 2.484 0 0 1 0 4.968z"/>`),
		g.Group(children),
	)
}