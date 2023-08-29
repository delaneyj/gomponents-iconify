package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Dw(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M12.897 29.016V18.984h2.257a4.39 4.39 0 0 1 4.39 4.389v1.254a4.39 4.39 0 0 1-4.39 4.39Zm25.014-10.032l-2.508 10.032l-2.508-10.032l-2.508 10.032l-2.508-10.032"/><circle cx="15.105" cy="24" r="11.605" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M23.999 16.547a11.605 11.605 0 1 1 0 14.906"/>`),
		g.Group(children),
	)
}