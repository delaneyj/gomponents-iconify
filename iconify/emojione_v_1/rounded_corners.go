package emojione_v_1

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RoundedCorners(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="#be1e2d" d="M56.628 62.09H7.288c-2.545 0-4.864-1.352-6.204-3.616c-1.387-2.346-1.445-5.268-.152-7.628l24.664-45.04a7.262 7.262 0 0 1 6.361-3.797c2.64 0 5.079 1.456 6.36 3.799l24.67 45.05c1.311 2.388 1.287 5.146-.055 7.423c-1.241 2.396-3.58 3.818-6.303 3.818m-48.37-7.988h47.4l-23.701-43.28l-23.697 43.28"/>`),
		g.Group(children),
	)
}