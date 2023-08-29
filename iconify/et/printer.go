package et

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Printer(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<g fill="currentColor"><path d="M40.5 10H38V6.5c0-.827-.673-1.5-1.5-1.5h-2a.5.5 0 0 0 0 1h2a.5.5 0 0 1 .5.5V10h-4V.5a.5.5 0 0 0-.5-.5h-23a.5.5 0 0 0-.5.5V10H5V6.5a.5.5 0 0 1 .5-.5h2a.5.5 0 0 0 0-1h-2C4.673 5 4 5.673 4 6.5V10H1.5c-.827 0-1.5.673-1.5 1.5v14c0 .827.673 1.5 1.5 1.5h6a.5.5 0 0 0 0-1h-6a.5.5 0 0 1-.5-.5V22h40v3.5a.5.5 0 0 1-.5.5h-6a.5.5 0 0 0 0 1h6c.827 0 1.5-.673 1.5-1.5v-14c0-.827-.673-1.5-1.5-1.5zM10 1h22v9H10V1zM1 21v-9.5a.5.5 0 0 1 .5-.5h39a.5.5 0 0 1 .5.5V21H1z"/><path d="M32.5 23a.5.5 0 0 0-.5.5V31H10v-7.5a.5.5 0 0 0-1 0v8a.5.5 0 0 0 .5.5h23a.5.5 0 0 0 .5-.5v-8a.5.5 0 0 0-.5-.5z"/><circle cx="38" cy="18" r="1"/><circle cx="34" cy="18" r="1"/></g>`),
		g.Group(children),
	)
}