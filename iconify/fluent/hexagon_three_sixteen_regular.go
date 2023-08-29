package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HexagonThreeSixteenRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M3.473 1.998a1 1 0 0 1 .865-.498h2.324a1 1 0 0 1 .865.498l1.16 2a1 1 0 0 1 0 1.004l-1.16 2a1 1 0 0 1-.865.498H4.338a1 1 0 0 1-.865-.498l-1.16-2a1 1 0 0 1 0-1.004l1.16-2Zm3.189.502H4.338l-1.16 2l1.16 2h2.324l1.16-2l-1.16-2Zm-3.19 6.498a1 1 0 0 1 .866-.498h2.324a1 1 0 0 1 .865.498l1.16 2a1 1 0 0 1 0 1.004l-1.16 2a1 1 0 0 1-.865.498H4.338a1 1 0 0 1-.865-.498l-1.16-2a1 1 0 0 1 0-1.004l1.16-2Zm3.19.502H4.338l-1.16 2l1.16 2h2.324l1.16-2l-1.16-2ZM10.338 5a1 1 0 0 0-.865.498l-1.16 2a1 1 0 0 0 0 1.004l1.16 2a1 1 0 0 0 .865.498h2.324a1 1 0 0 0 .865-.498l1.16-2a1 1 0 0 0 0-1.004l-1.16-2A1 1 0 0 0 12.662 5h-2.324Zm0 1h2.324l1.16 2l-1.16 2h-2.324l-1.16-2l1.16-2Z"/>`),
		g.Group(children),
	)
}