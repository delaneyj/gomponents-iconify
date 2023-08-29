package mdi_light

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GridOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m.79 2.46l.71-.71L22.25 22.5l-.71.71l-1.78-1.78c-.49.36-1.1.57-1.76.57H5a3 3 0 0 1-3-3V6c0-.66.21-1.27.57-1.76L.79 2.46M5 3h13a3 3 0 0 1 3 3v13l-.09.74l-.91-.9V16h-2.84l-1-1H20v-5h-5v3.84l-1-1V10h-2.84l-1-1H14V4H9v3.84l-1-1V4H5.16l-.9-.91L5 3M3 6v3h4.34L3.29 4.96C3.11 5.26 3 5.62 3 6m6 9h4.34L9 10.66V15m6 6h3c.38 0 .74-.11 1.04-.29L15 16.66V21M3 19a2 2 0 0 0 2 2h3v-5H3v3m5-9H3v5h5v-5m12-4a2 2 0 0 0-2-2h-3v5h5V6M9 21h5v-5H9v5Z"/>`),
		g.Group(children),
	)
}