package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func IntersectSquareThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M216 92h-52V40a4 4 0 0 0-4-4H40a4 4 0 0 0-4 4v120a4 4 0 0 0 4 4h52v52a4 4 0 0 0 4 4h120a4 4 0 0 0 4-4V96a4 4 0 0 0-4-4ZM44 156V44h112v48H96a4 4 0 0 0-4 4v60Zm56-50.34L150.34 156H100Zm56 44.68L105.66 100H156ZM212 212H100v-48h60a4 4 0 0 0 4-4v-60h48Z"/>`),
		g.Group(children),
	)
}