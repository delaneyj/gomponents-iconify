package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func OfficeChairBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M252 128a12 12 0 0 1-12 12h-13.4a52.09 52.09 0 0 1-50.6 40h-36v16h20a36 36 0 0 1 36 36a12 12 0 0 1-24 0a12 12 0 0 0-12-12h-20v12a12 12 0 0 1-24 0v-12H96a12 12 0 0 0-12 12a12 12 0 0 1-24 0a36 36 0 0 1 36-36h20v-16H80a52.09 52.09 0 0 1-50.6-40H16a12 12 0 0 1 0-24h24a12 12 0 0 1 12 12a28 28 0 0 0 28 28h96a28 28 0 0 0 28-28a12 12 0 0 1 12-12h24a12 12 0 0 1 12 12Zm-179.18 5a20 20 0 0 1-4.59-16L81.8 29a19.91 19.91 0 0 1 19.79-17h52.82a19.91 19.91 0 0 1 19.79 17l13.57 88A20 20 0 0 1 168 140H88a20 20 0 0 1-15.18-7Zm19.85-17h70.66L151 36h-46Z"/>`),
		g.Group(children),
	)
}