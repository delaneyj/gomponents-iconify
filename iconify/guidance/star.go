package guidance

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Star(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" d="M11.761 1.5h.478l.619 2.243a8 8 0 0 0 7.711 5.87H22.5V10l-1.656 1.143a8 8 0 0 0-3.1 8.946l.745 2.411H18l-.904-.747a8 8 0 0 0-10.192 0L6 22.5h-.49l.746-2.41a8 8 0 0 0-3.1-8.947L1.5 10v-.386h1.93a8 8 0 0 0 7.712-5.871l.62-2.243Z"/>`),
		g.Group(children),
	)
}