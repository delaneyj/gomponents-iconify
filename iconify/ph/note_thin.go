package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NoteThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M92 96a4 4 0 0 1 4-4h64a4 4 0 0 1 0 8H96a4 4 0 0 1-4-4Zm4 36h64a4 4 0 0 0 0-8H96a4 4 0 0 0 0 8Zm32 24H96a4 4 0 0 0 0 8h32a4 4 0 0 0 0-8Zm92-108v108.69a11.9 11.9 0 0 1-3.52 8.48l-51.31 51.32a11.93 11.93 0 0 1-8.48 3.51H48a12 12 0 0 1-12-12V48a12 12 0 0 1 12-12h160a12 12 0 0 1 12 12ZM48 212h108v-52a4 4 0 0 1 4-4h52V48a4 4 0 0 0-4-4H48a4 4 0 0 0-4 4v160a4 4 0 0 0 4 4Zm158.35-48H164v42.35Z"/>`),
		g.Group(children),
	)
}