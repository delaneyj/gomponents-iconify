package fa_6_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PersonThroughWindow(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M64 64h224v9.8c0 39-23.7 74-59.9 88.4C167.6 186.5 128 245 128 310.2V384H64V64zm288 0h224v320h-67.7l-3.7-4.5l-75.2-90.2c-9.1-10.9-22.6-17.3-36.9-17.3h-71.1l-41-63.1c-.3-.5-.6-1-1-1.4c44.7-29 72.5-79 72.5-133.6v-9.8zm73 320h-45.8l42.7 64H592c26.5 0 48-21.5 48-48V48c0-26.5-21.5-48-48-48H48C21.5 0 0 21.5 0 48v352c0 26.5 21.5 48 48 48h260.2l33.2 49.8c9.8 14.7 29.7 18.7 44.4 8.9s18.7-29.7 8.9-44.4L310.5 336h74.6l40 48zm-159.5 0H192v-73.8c0-10.2 1.6-20.1 4.7-29.5L265.5 384zM192 128a48 48 0 1 0-96 0a48 48 0 1 0 96 0z"/>`),
		g.Group(children),
	)
}