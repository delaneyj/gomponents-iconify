package file_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Galaxy(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M0 166.565V54.935h381.034v111.63H0zm0 146.348V201.284h284.347v111.63H0zm130.088 144.151V345.435H512v111.63H130.088z"/>`),
		g.Group(children),
	)
}