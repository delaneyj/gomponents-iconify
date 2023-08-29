package logos

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlickrIcon(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="#FF0084" d="M199.111 113.778c31.42 0 56.889-25.47 56.889-56.89C256 25.476 230.53 0 199.111 0c-31.42 0-56.889 25.475-56.889 56.889c0 31.42 25.47 56.889 56.89 56.889"/><path fill="#0063DC" d="M56.889 113.778c31.42 0 56.889-25.47 56.889-56.89C113.778 25.476 88.308 0 56.888 0C25.47 0 0 25.475 0 56.889c0 31.42 25.47 56.889 56.889 56.889"/>`),
		g.Group(children),
	)
}