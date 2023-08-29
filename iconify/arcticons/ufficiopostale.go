package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Ufficiopostale(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="24" cy="24" r="21.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M12.346 32V16h5.342c3.082 0 5.547 2.4 5.547 5.4s-2.465 5.4-5.547 5.4h-5.342M24.765 16h10.889m-5.342 15.999V16"/>`),
		g.Group(children),
	)
}