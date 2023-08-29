package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CassetteTapeLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M224 50H32a14 14 0 0 0-14 14v128a14 14 0 0 0 14 14h192a14 14 0 0 0 14-14V64a14 14 0 0 0-14-14ZM76 194l14.4-19.2a2 2 0 0 1 1.6-.8h72a2 2 0 0 1 1.6.8L180 194Zm150-2a2 2 0 0 1-2 2h-29l-19.8-26.4A14.06 14.06 0 0 0 164 162H92a14.06 14.06 0 0 0-11.2 5.6L61 194H32a2 2 0 0 1-2-2V64a2 2 0 0 1 2-2h192a2 2 0 0 1 2 2ZM176 82H80a30 30 0 0 0 0 60h96a30 30 0 0 0 0-60Zm-24 12a29.92 29.92 0 0 0 0 36h-48a29.92 29.92 0 0 0 0-36Zm-90 18a18 18 0 1 1 18 18a18 18 0 0 1-18-18Zm114 18a18 18 0 1 1 18-18a18 18 0 0 1-18 18Z"/>`),
		g.Group(children),
	)
}