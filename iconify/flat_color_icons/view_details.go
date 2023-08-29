package flat_color_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ViewDetails(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="#BBDEFB" d="M7 4h34v40H7z"/><path fill="#2196F3" d="M13 26h4v4h-4zm0-8h4v4h-4zm0 16h4v4h-4zm0-24h4v4h-4zm8 16h14v4H21zm0-8h14v4H21zm0 16h14v4H21zm0-24h14v4H21z"/>`),
		g.Group(children),
	)
}