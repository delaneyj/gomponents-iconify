package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bus(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M184 32H72a32 32 0 0 0-32 32v144a16 16 0 0 0 16 16h24a16 16 0 0 0 16-16v-16h64v16a16 16 0 0 0 16 16h24a16 16 0 0 0 16-16V64a32 32 0 0 0-32-32ZM56 176v-56h144v56Zm0-96h144v24H56Zm16-32h112a16 16 0 0 1 16 16H56a16 16 0 0 1 16-16Zm8 160H56v-16h24Zm96 0v-16h24v16Zm-72-60a12 12 0 1 1-12-12a12 12 0 0 1 12 12Zm72 0a12 12 0 1 1-12-12a12 12 0 0 1 12 12Zm72-68v24a8 8 0 0 1-16 0V80a8 8 0 0 1 16 0ZM24 80v24a8 8 0 0 1-16 0V80a8 8 0 0 1 16 0Z"/>`),
		g.Group(children),
	)
}