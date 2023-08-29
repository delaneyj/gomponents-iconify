package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CropTwentyFourFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M7 15.5a1.5 1.5 0 0 0 1.356 1.493L8.5 17H21a1 1 0 0 1 .117 1.993L21 19h-2v2a1 1 0 0 1-1.993.117L17 21v-2H8.5a3.5 3.5 0 0 1-3.495-3.308L5 15.5V7H3a1 1 0 0 1-.117-1.993L3 5h2V3a1 1 0 0 1 1.993-.117L7 3v12.5ZM8 5h7.5a3.5 3.5 0 0 1 3.495 3.308L19 8.5V16h-2V8.5a1.5 1.5 0 0 0-1.355-1.493L15.5 7H8V5Z"/>`),
		g.Group(children),
	)
}