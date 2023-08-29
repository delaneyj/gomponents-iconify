package foundation

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Braille(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 100 100"),
		g.Raw(`<circle cx="31.102" cy="23.979" r="6.978" fill="currentColor"/><circle cx="68.898" cy="23.979" r="6.978" fill="currentColor"/><circle cx="31.102" cy="50" r="6.978" fill="currentColor"/><circle cx="68.898" cy="50" r="6.978" fill="currentColor"/><circle cx="31.102" cy="76.02" r="6.978" fill="currentColor"/><circle cx="68.898" cy="76.02" r="6.978" fill="currentColor"/>`),
		g.Group(children),
	)
}