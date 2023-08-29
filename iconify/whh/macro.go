package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Macro(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M576 737q41-73 116-120q38-23 84-33.5t93.5-5.5t89.5 37t65 89q-47 0-88.5 14T864 754.5t-59 51t-52 58.5t-50 59t-53 51.5t-61.5 36T513 1024h-2q-40 0-75.5-13.5t-61.5-36t-53-51.5t-50-59t-52-58.5t-59-51T88.5 718T0 704q23-57 65-89t89.5-37t93.5 5.5t84 33.5q75 47 116 120V511q-87 5-172-48q-71-44-111-110t-37-130q47-34 117.5-31T385 234q5-80 39.5-143.5T512 0q53 27 87.5 90.5T639 234q69-39 139.5-41.5T896 223q3 64-37 130T748 463q-85 53-172 48v226z"/>`),
		g.Group(children),
	)
}