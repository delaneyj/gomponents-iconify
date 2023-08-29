package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CircleHalfTiltBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M204.37 51.6A108.08 108.08 0 1 0 236 128a108.09 108.09 0 0 0-31.63-76.4ZM176 197a83.43 83.43 0 0 1-16 8.75V113l16-16ZM68.6 68.58a84.08 84.08 0 0 1 109.7-7.88L60.72 178.33A84.08 84.08 0 0 1 68.6 68.58ZM96 177v28.69a83.63 83.63 0 0 1-18.3-10.39Zm24 34.62V153l16-16v74.64a84.68 84.68 0 0 1-16-.02Zm80-40.27v-86.7a84.24 84.24 0 0 1 0 86.7Z"/>`),
		g.Group(children),
	)
}