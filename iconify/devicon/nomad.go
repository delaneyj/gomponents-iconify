package devicon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Nomad(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 128 128"),
		g.Raw(`<path fill="#00ca8e" d="M64 0L8.569 31.984v64.013L64 128l55.431-32.003V31.984Zm24.697 70.393L73.924 78.9L56.08 69.212v20.424l-16.776 10.615V57.605l13.308-8.146l18.467 9.734V38.4l17.628-10.584Z"/>`),
		g.Group(children),
	)
}