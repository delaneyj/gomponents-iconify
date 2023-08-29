package fa_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Sink(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M32 416a96 96 0 0 0 96 96h256a96 96 0 0 0 96-96v-32H32Zm464-128h-96v-32h64a16 16 0 0 0 16-16v-16a16 16 0 0 0-16-16h-80a32 32 0 0 0-32 32v48h-64V96a32 32 0 0 1 64 0v16a16 16 0 0 0 16 16h32a16 16 0 0 0 16-16V96A96.16 96.16 0 0 0 300.87 1.86C255.29 10.71 224 53.36 224 99.79V288h-64v-48a32 32 0 0 0-32-32H48a16 16 0 0 0-16 16v16a16 16 0 0 0 16 16h64v32H16a16 16 0 0 0-16 16v32a16 16 0 0 0 16 16h480a16 16 0 0 0 16-16v-32a16 16 0 0 0-16-16Z"/>`),
		g.Group(children),
	)
}