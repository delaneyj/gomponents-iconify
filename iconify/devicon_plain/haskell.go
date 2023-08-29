package devicon_plain

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Haskell(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 128 128"),
		g.Raw(`<path fill="currentColor" d="M0 110.2L30.1 65L0 19.9h22.6L52.7 65l-30.1 45.1H0z"/><path fill="currentColor" d="M30.1 110.2L60.2 65L30.1 19.9h22.6l60.2 90.3H90.4L71.5 81.9l-18.8 28.2H30.1zm72.8-26.4l-10-15.1H128v15.1h-25.1zM87.8 61.3l-10-15.1H128v15.1H87.8z"/>`),
		g.Group(children),
	)
}