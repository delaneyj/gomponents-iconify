package bi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UsbPlug(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M6 .5a.5.5 0 0 1 .5-.5h4a.5.5 0 0 1 .5.5v4H6v-4ZM7 1v1h1V1H7Zm2 0v1h1V1H9ZM6 5a1 1 0 0 0-1 1v4.394c0 .494.146.976.42 1.387l1.038 1.558c.354.53.542 1.152.542 1.789c0 .481.39.872.872.872h1.256c.481 0 .872-.39.872-.872c0-.637.188-1.26.541-1.789l1.04-1.558A2.5 2.5 0 0 0 12 10.394V6a1 1 0 0 0-1-1H6Zm0 1h5v4.394a1.5 1.5 0 0 1-.252.832L9.71 12.784A4.224 4.224 0 0 0 9.002 15H7.998a4.224 4.224 0 0 0-.707-2.216l-1.04-1.558A1.5 1.5 0 0 1 6 10.394V6Z"/>`),
		g.Group(children),
	)
}