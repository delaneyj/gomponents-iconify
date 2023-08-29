package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AdvertisementsOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m12.2 9l-2-2H13c1.1 0 2 .9 2 2v2.8l-2-2V9h-.8M23 9V7h-4c-1.1 0-2 .9-2 2v2c0 1.1.9 2 2 2h2v2h-2.8l2 2h.8c1.1 0 2-.9 2-2v-2c0-1.1-.9-2-2-2h-2V9h4m-.9 12.5l-1.3 1.3l-6.4-6.4c-.3.3-.8.6-1.4.6H9v-6.1l-2-2V17H5v-4H3v4H1V9c0-1.1.9-2 2-2h2.1l-4-4l1.3-1.3l19.7 19.8M5 9H3v2h2V9m8 5.9l-2-2V15h2v-.1Z"/>`),
		g.Group(children),
	)
}