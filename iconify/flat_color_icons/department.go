package flat_color_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Department(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="#C5CAE9" d="M42 42H6V9l18-7l18 7z"/><path fill="#9FA8DA" d="M6 42h36v2H6z"/><path fill="#BF360C" d="M20 35h8v9h-8z"/><path fill="#1565C0" d="M31 27h6v5h-6zm-10 0h6v5h-6zm-10 0h6v5h-6zm20 8h6v5h-6zm-20 0h6v5h-6zm20-16h6v5h-6zm-10 0h6v5h-6zm-10 0h6v5h-6zm20-8h6v5h-6zm-10 0h6v5h-6zm-10 0h6v5h-6z"/>`),
		g.Group(children),
	)
}