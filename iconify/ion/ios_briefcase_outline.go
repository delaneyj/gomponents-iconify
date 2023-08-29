package ion

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func IosBriefcaseOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path d="M416 128v-16h-32v16h-48V96c-.5-18.2-13.6-32-32.2-32h-99.1C186.1 64 176 77.7 176 96v32h-48v-16H96v16H48v320h416V128h-48zm-224-.5V98.1c0-10.2 2.9-18.1 13.7-18.1h97.1c10.4 0 17.2 7.7 17.2 18.1V128H192v-.5zM448 432H64V208h384v224zm0-240H64v-48h32v16h32v-16h256v16h32v-16h32v48z" fill="currentColor"/>`),
		g.Group(children),
	)
}