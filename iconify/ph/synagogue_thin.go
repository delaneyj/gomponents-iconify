package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SynagogueThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M204 60.4V32a4 4 0 0 0-8 0v28.4A20 20 0 0 0 180 80v49.11l-48-27.43V72a4 4 0 0 0-8 0v29.68l-48 27.43V80a20 20 0 0 0-16-19.6V32a4 4 0 0 0-8 0v28.4A20 20 0 0 0 36 80v136a4 4 0 0 0 4 4h72a4 4 0 0 0 4-4v-40a12 12 0 0 1 24 0v40a4 4 0 0 0 4 4h72a4 4 0 0 0 4-4V80a20 20 0 0 0-16-19.6Zm-4 7.6a12 12 0 0 1 12 12v28h-24V80a12 12 0 0 1 12-12ZM56 68a12 12 0 0 1 12 12v28H44V80a12 12 0 0 1 12-12Zm-12 48h24v96H44Zm84 40a20 20 0 0 0-20 20v36H76v-73.68l52-29.71l52 29.71V212h-32v-36a20 20 0 0 0-20-20Zm60 56v-96h24v96Z"/>`),
		g.Group(children),
	)
}