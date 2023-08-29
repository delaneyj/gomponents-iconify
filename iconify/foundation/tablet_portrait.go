package foundation

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TabletPortrait(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 100 100"),
		g.Raw(`<path fill="currentColor" d="M77.919 7.434c-.043 0-.084.011-.127.013v-.013H21.954v.013a3.028 3.028 0 0 0-2.908 3.022v79.062a3.028 3.028 0 0 0 2.908 3.022v.013h55.838v-.013c.043.002.083.013.127.013a3.035 3.035 0 0 0 3.035-3.035V10.469a3.035 3.035 0 0 0-3.035-3.035zM49.873 90.072a2.485 2.485 0 1 1 .001-4.97a2.485 2.485 0 0 1-.001 4.97zm21.124-7.463H29.003V17.391h41.995v65.218z"/>`),
		g.Group(children),
	)
}