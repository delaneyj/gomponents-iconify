package circum

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Power(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6.638 5.209a8.782 8.782 0 1 0 13.917 8.96a8.871 8.871 0 0 0-3.189-8.96c-.5-.39-1.214.312-.707.707a7.93 7.93 0 0 1 3.082 7.113a7.787 7.787 0 0 1-15.308.956a7.9 7.9 0 0 1 2.912-8.069c.507-.394-.205-1.1-.707-.707Z"/><path fill="currentColor" d="M12.5 12.519a.5.5 0 0 1-1 0V3.548a.5.5 0 0 1 1 0Z"/>`),
		g.Group(children),
	)
}