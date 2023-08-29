package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DocumentImageSixteenRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M6 1a2 2 0 0 0-2 2v3h1V3a1 1 0 0 1 1-1h3v2.5A1.5 1.5 0 0 0 10.5 6H13v7a1 1 0 0 1-1 1h-1.035c-.05.353-.154.69-.302 1H12a2 2 0 0 0 2-2V5.414a1.5 1.5 0 0 0-.44-1.06l-2.914-2.915A1.5 1.5 0 0 0 9.586 1H6Zm6.793 4H10.5a.5.5 0 0 1-.5-.5V2.207L12.793 5ZM1 9.5A2.5 2.5 0 0 1 3.5 7h4A2.5 2.5 0 0 1 10 9.5v4c0 .51-.152.983-.414 1.379L6.56 11.854a1.5 1.5 0 0 0-2.122 0l-3.025 3.025A2.488 2.488 0 0 1 1 13.5v-4Zm7 .25a.75.75 0 1 0-1.5 0a.75.75 0 0 0 1.5 0Zm-5.879 5.836c.396.262.87.414 1.379.414h4c.51 0 .983-.152 1.379-.414L5.854 12.56a.5.5 0 0 0-.708 0l-3.025 3.025Z"/>`),
		g.Group(children),
	)
}