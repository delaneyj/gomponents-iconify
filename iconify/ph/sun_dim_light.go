package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SunDimLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M122 40v-8a6 6 0 0 1 12 0v8a6 6 0 0 1-12 0Zm68 88a62 62 0 1 1-62-62a62.07 62.07 0 0 1 62 62Zm-12 0a50 50 0 1 0-50 50a50.06 50.06 0 0 0 50-50ZM59.76 68.24a6 6 0 1 0 8.48-8.48l-8-8a6 6 0 0 0-8.48 8.48Zm0 119.52l-8 8a6 6 0 1 0 8.48 8.48l8-8a6 6 0 1 0-8.48-8.48Zm136-136l-8 8a6 6 0 1 0 8.48 8.48l8-8a6 6 0 0 0-8.48-8.48Zm.48 136a6 6 0 0 0-8.48 8.48l8 8a6 6 0 0 0 8.48-8.48ZM40 122h-8a6 6 0 0 0 0 12h8a6 6 0 0 0 0-12Zm88 88a6 6 0 0 0-6 6v8a6 6 0 0 0 12 0v-8a6 6 0 0 0-6-6Zm96-88h-8a6 6 0 0 0 0 12h8a6 6 0 0 0 0-12Z"/>`),
		g.Group(children),
	)
}