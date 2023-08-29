package fa_6_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Taxi(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M192 0c-17.7 0-32 14.3-32 32v32.2c-38.6 2.2-72.3 27.3-85.2 64.1L39.6 228.8C16.4 238.4 0 261.3 0 288v192c0 17.7 14.3 32 32 32h32c17.7 0 32-14.3 32-32v-48h320v48c0 17.7 14.3 32 32 32h32c17.7 0 32-14.3 32-32V288c0-26.7-16.4-49.6-39.6-59.2l-35.2-100.5c-12.9-36.8-46.6-62-85.2-64.1V32c0-17.7-14.3-32-32-32H192zm-26.6 128h181.2c13.6 0 25.7 8.6 30.2 21.4l26.1 74.6H109.1l26.1-74.6c4.5-12.8 16.6-21.4 30.2-21.4zM96 288a32 32 0 1 1 0 64a32 32 0 1 1 0-64zm288 32a32 32 0 1 1 64 0a32 32 0 1 1-64 0z"/>`),
		g.Group(children),
	)
}