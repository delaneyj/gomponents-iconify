package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PaperPlaneTiltBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M230.14 25.86a20 20 0 0 0-19.57-5.11l-.22.07L18.44 79a20 20 0 0 0-3 37.28l84.32 40l40 84.32a19.81 19.81 0 0 0 18 11.44c.57 0 1.15 0 1.73-.07A19.82 19.82 0 0 0 177 237.56l58.18-191.91a1.42 1.42 0 0 0 .07-.22a20 20 0 0 0-5.11-19.57ZM157 220.92l-33.72-71.19l45.25-45.25a12 12 0 0 0-17-17l-45.25 45.25L35.08 99L210 46Z"/>`),
		g.Group(children),
	)
}