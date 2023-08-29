package fa_6_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HandHoldingDroplet(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M275.5 6.6C278.3 2.5 283 0 288 0s9.7 2.5 12.5 6.6l66.3 96.4c11.2 16.3 17.2 35.6 17.2 55.3v1.7c0 53-43 96-96 96s-96-43-96-96v-1.7c0-19.8 6-39 17.2-55.3l66.3-96.4zm292.7 329.7c13.1 17.8 9.3 42.8-8.5 55.9l-126.6 93.3c-23.4 17.2-51.6 26.5-80.7 26.5H32c-17.7 0-32-14.3-32-32v-64c0-17.7 14.3-32 32-32h36.8l44.9-36c22.7-18.2 50.9-28 80-28H352c17.7 0 32 14.3 32 32s-14.3 32-32 32h-80c-8.8 0-16 7.2-16 16s7.2 16 16 16h120.6l119.7-88.2c17.8-13.1 42.8-9.3 55.9 8.5zM193.6 384h-.9h.9z"/>`),
		g.Group(children),
	)
}