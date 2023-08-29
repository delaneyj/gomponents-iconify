package fa_6_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NotEqual(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M369.8 37.4c14.7 9.8 18.7 29.7 8.9 44.4L337.1 144H400c17.7 0 32 14.3 32 32s-14.3 32-32 32H294.5l-64 96H400c17.7 0 32 14.3 32 32s-14.3 32-32 32H187.8l-65.2 97.7c-9.8 14.7-29.7 18.7-44.4 8.9s-18.7-29.7-8.9-44.4l41.6-62.2H48c-17.7 0-32-14.3-32-32s14.3-32 32-32h105.5l64-96H48c-17.7 0-32-14.3-32-32s14.3-32 32-32h212.2l65.2-97.7c9.8-14.7 29.7-18.7 44.4-8.9z"/>`),
		g.Group(children),
	)
}