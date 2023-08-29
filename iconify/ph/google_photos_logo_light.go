package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GooglePhotosLogoLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M232 122h-42.82A70 70 0 0 0 128 18a6 6 0 0 0-6 6v42.82A70 70 0 0 0 18 128a6 6 0 0 0 6 6h42.82A70 70 0 0 0 128 238a6 6 0 0 0 6-6v-42.82A70 70 0 0 0 238 128a6 6 0 0 0-6-6Zm-46-34a57.3 57.3 0 0 1-11 34h-41V30.31A58.08 58.08 0 0 1 186 88ZM88 70a57.3 57.3 0 0 1 34 11v41H30.31A58.08 58.08 0 0 1 88 70Zm-18 98a57.3 57.3 0 0 1 11-34h41v91.69A58.08 58.08 0 0 1 70 168Zm98 18a57.3 57.3 0 0 1-34-11v-41h91.69A58.08 58.08 0 0 1 168 186Z"/>`),
		g.Group(children),
	)
}