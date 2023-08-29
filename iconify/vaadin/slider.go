package vaadin

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Slider(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M16 6h-3.6c-.7-1.2-2-2-3.4-2s-2.8.8-3.4 2H0v4h5.6c.7 1.2 2 2 3.4 2s2.8-.8 3.4-2H16V6zM1 9V7h4.1c0 .3-.1.7-.1 1s.1.7.1 1H1zm8 2c-1.7 0-3-1.3-3-3s1.3-3 3-3s3 1.3 3 3s-1.3 3-3 3z"/>`),
		g.Group(children),
	)
}