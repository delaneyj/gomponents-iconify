package file_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Asymptote(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M376.595 345.736c37.216-40.707 63.484-88.782 84.134-138.704v138.704zM463.18 12.962c-23.372 97.456-51.867 237.008-151.63 326.092C204.857 434.33 23.998 457.4 11.657 463.848c-23.45 17.976-8.277 50.994 16.894 47.956c142.277-29.12 229.65-65.352 292.852-115.325H460.73v91.263c-.7 14.406 12.406 22.587 24.705 23.422c12.778.868 25.969-7.942 26.566-21.85V23.708c-.538-24.867-34.758-32.925-48.819-10.745z"/>`),
		g.Group(children),
	)
}