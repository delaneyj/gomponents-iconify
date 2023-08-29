package fa

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ThumbTack(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1536 1536"),
		g.Raw(`<path fill="currentColor" d="M480 736V288q0-14-9-23t-23-9t-23 9t-9 23v448q0 14 9 23t23 9t23-9t9-23zm672 352q0 26-19 45t-45 19H659l-51 483q-2 12-10.5 20.5T577 1664h-1q-27 0-32-27l-76-485H64q-26 0-45-19t-19-45q0-123 78.5-221.5T256 768V256q-52 0-90-38t-38-90t38-90t90-38h640q52 0 90 38t38 90t-38 90t-90 38v512q99 0 177.5 98.5T1152 1088z"/>`),
		g.Group(children),
	)
}