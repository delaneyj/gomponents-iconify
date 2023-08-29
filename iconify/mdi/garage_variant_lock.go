package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GarageVariantLock(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M21.8 16v-1.5c0-1.4-1.4-2.5-2.8-2.5s-2.8 1.1-2.8 2.5V16c-.6 0-1.2.6-1.2 1.2v3.5c0 .7.6 1.3 1.2 1.3h5.5c.7 0 1.3-.6 1.3-1.2v-3.5c0-.7-.6-1.3-1.2-1.3m-1.3 0h-3v-1.5c0-.8.7-1.3 1.5-1.3s1.5.5 1.5 1.3V16M5 12h10.04c-.43.59-.69 1.27-.78 2H5v-2m11.06-1H4v9H2V9l10-4l10 4v2.04A4.911 4.911 0 0 0 19 10c-1.1 0-2.12.39-2.94 1M13 20H5v-2h8v2m-8-5h8.95c-.53.54-.87 1.24-.95 2H5v-2Z"/>`),
		g.Group(children),
	)
}