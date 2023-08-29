package la

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HandPointRightSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="m14.688 3l-.313.281L7.594 10H2v16h16.844c1.41 0 2.64-.996 2.937-2.375L23.438 16H27c1.645 0 3-1.355 3-3s-1.355-3-3-3H16.25l.188-.75c.203-.156.332-.223.625-.625c.468-.64.937-1.633.937-2.969C18 4.23 16.71 3 15.094 3zm.718 2.094c.422.082.594.254.594.562c0 .903-.273 1.461-.531 1.813c-.258.351-.438.437-.438.437l-.344.188l-.124.406l-.594 2.25l-.313 1.25H27c.566 0 1 .434 1 1c0 .566-.434 1-1 1h-5.188l-.187.781l-1.781 8.438a1.008 1.008 0 0 1-1 .781H9V11.406zM4 12h3v12H4z"/>`),
		g.Group(children),
	)
}