package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FilmStripThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M216 44H40a12 12 0 0 0-12 12v144a12 12 0 0 0 12 12h176a12 12 0 0 0 12-12V56a12 12 0 0 0-12-12ZM36 84h88v88H36Zm96-8V52h40v24Zm-8 0H84V52h40Zm0 104v24H84v-24Zm8 0h40v24h-40Zm0-8V84h88v88Zm88-116v20h-40V52h36a4 4 0 0 1 4 4ZM40 52h36v24H36V56a4 4 0 0 1 4-4Zm-4 148v-20h40v24H40a4 4 0 0 1-4-4Zm180 4h-36v-24h40v20a4 4 0 0 1-4 4Z"/>`),
		g.Group(children),
	)
}