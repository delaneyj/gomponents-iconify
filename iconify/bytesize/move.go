package bytesize

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Move(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 16h26M16 3v26M12 7l4-4l4 4m-8 18l4 4l4-4m5-13l4 4l-4 4M7 12l-4 4l4 4"/>`),
		g.Group(children),
	)
}