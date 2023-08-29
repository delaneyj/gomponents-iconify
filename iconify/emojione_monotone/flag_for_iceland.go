package emojione_monotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagForIceland(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="currentColor" d="M32 2C15.432 2 2 15.432 2 32s13.432 30 30 30s30-13.432 30-30S48.568 2 32 2m-.983 36.883v21.092a27.845 27.845 0 0 1-2.95-.256V35.934h31.652a27.89 27.89 0 0 1-.577 2.949H31.017m-26.153.023a27.792 27.792 0 0 1-.582-2.973H20.2v21.45a28.021 28.021 0 0 1-2.95-1.599V38.906H4.864m12.386-13.79V8.215a28.021 28.021 0 0 1 2.95-1.599v21.45H4.282c.142-1.001.332-1.986.577-2.95H17.25m10.816 2.95V4.281a27.845 27.845 0 0 1 2.95-.256v21.091h28.125c.245.964.435 1.949.577 2.95H28.066"/>`),
		g.Group(children),
	)
}