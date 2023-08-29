package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Fotmob(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="31.746" cy="23.887" r="6.223" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M19.991 35.989v5.297M43.5 6.714H19.991v16.959m7.786 5.006L6.971 39.039M4.5 31.462l21.197-6.135"/>`),
		g.Group(children),
	)
}