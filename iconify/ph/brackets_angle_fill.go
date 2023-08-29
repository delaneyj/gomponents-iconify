package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BracketsAngleFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M216 40H40a16 16 0 0 0-16 16v144a16 16 0 0 0 16 16h176a16 16 0 0 0 16-16V56a16 16 0 0 0-16-16ZM103 180a8 8 0 0 1-13.95 8l-32-56a8 8 0 0 1 0-7.94l32-56A8 8 0 0 1 103 76l-29.79 52Zm96-48l-32 56a8 8 0 0 1-13.9-7.94l29.74-52L153.05 76A8 8 0 1 1 167 68l32 56a8 8 0 0 1 0 8Z"/>`),
		g.Group(children),
	)
}