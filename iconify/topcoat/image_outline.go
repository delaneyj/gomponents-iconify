package topcoat

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ImageOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 42 42"),
		g.Raw(`<path fill="currentColor" d="M35.79 31.5c-.05-.04-5.181-7.971-6.72-7.971c-1.51 0-4.391 3.851-4.391 3.851c-2.24-2.31-7.2-8.87-8.661-8.87c-1.509 0-7.449 8.12-10.789 12.99H35.79zm-8.979-17c0 2.04 1.649 3.69 3.689 3.69s3.689-1.65 3.689-3.69s-1.649-3.69-3.689-3.69s-3.689 1.65-3.689 3.69zM.5 7.5v27c0 2.52.51 3 3 3h34c2.471 0 3-.46 3-3v-27c0-2.46-.471-3-3-3h-34c-2.48 0-3 .43-3 3zm3 0h34v27h-34v-27z"/>`),
		g.Group(children),
	)
}