package emojione_v_1

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UpLeftArrow(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="#1b75bb" d="M63.873 56.802a6.863 6.863 0 0 1-6.862 6.868H7.067a6.864 6.864 0 0 1-6.86-6.868V6.864A6.863 6.863 0 0 1 7.067 0H57.01a6.862 6.862 0 0 1 6.862 6.864v49.938z"/><path fill="#fff" d="m16.264 40.451l-2.192-24.21l24.2 1.263c1.705 1.917 1.964 4.541.519 6.05l-2.452 2.556l11.997 11.506a5.069 5.069 0 0 1 .15 7.166l-4.552 4.743a5.067 5.067 0 0 1-7.165.146L24.775 38.164l-2.45 2.554c-1.444 1.502-4.068 1.359-6.06-.262"/>`),
		g.Group(children),
	)
}