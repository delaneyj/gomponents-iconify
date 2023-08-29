package iwwa

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MiddleO(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 40 40"),
		g.Raw(`<path fill="currentColor" d="M20.001 2.683C10.452 2.683 2.684 10.451 2.684 20s7.769 17.317 17.317 17.317S37.316 29.548 37.316 20S29.549 2.683 20.001 2.683zm0 33.134c-8.722 0-15.817-7.096-15.817-15.817S11.279 4.183 20.001 4.183c8.721 0 15.815 7.096 15.815 15.817s-7.094 15.817-15.815 15.817z"/><circle cx="15.431" cy="16.424" r="1.044" fill="currentColor"/><circle cx="24.569" cy="16.424" r="1.044" fill="currentColor"/><path fill="currentColor" d="M24.828 25.312h-9.656a.75.75 0 0 0 0 1.5h9.656a.75.75 0 0 0 0-1.5z"/>`),
		g.Group(children),
	)
}