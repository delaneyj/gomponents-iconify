package flat_color_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Internal(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="24" cy="30" r="15" fill="#B3E5FC"/><g fill="#1565C0"><path d="M24 38.7L15 28h18z"/><path d="M21 5h6v26h-6z"/></g>`),
		g.Group(children),
	)
}