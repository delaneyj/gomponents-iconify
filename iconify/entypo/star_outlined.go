package entypo

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StarOutlined(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M18.8 8.022h-6.413L10 1.3L7.611 8.022H1.199l5.232 3.947l-1.871 6.929L10 14.744l5.438 4.154l-1.869-6.929L18.8 8.022zM10 12.783l-3.014 2.5l1.243-3.562l-2.851-2.3l3.522.101l1.1-4.04l1.099 4.04l3.521-.101l-2.851 2.3l1.243 3.562l-3.012-2.5z"/>`),
		g.Group(children),
	)
}