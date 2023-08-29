package iconoir

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Jellyfish(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M12 2c4.97 0 9 4.104 9 9.167c0 .068 0 .136-.002.204c-.02.954-.865 1.629-1.819 1.629H4.821c-.954 0-1.798-.675-1.819-1.629A9.52 9.52 0 0 1 3 11.167C3 6.104 7.03 2 12 2ZM6 13l1 1.125c.57.642.57 1.608 0 2.25v0a1.693 1.693 0 0 0 0 2.25v0c.57.642.57 1.608 0 2.25L6 22m5-9l1 1.125c.57.642.57 1.608 0 2.25v0a1.693 1.693 0 0 0 0 2.25v0c.57.642.57 1.608 0 2.25L11 22m5-9l1 1.125c.57.642.57 1.608 0 2.25v0a1.693 1.693 0 0 0 0 2.25v0c.57.642.57 1.608 0 2.25L16 22"/>`),
		g.Group(children),
	)
}