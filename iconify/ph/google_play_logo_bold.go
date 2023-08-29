package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GooglePlayLogoBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M225.79 110.7L58 14.65a20.24 20.24 0 0 0-20.12.06A19.62 19.62 0 0 0 28 31.84v192.32a19.62 19.62 0 0 0 9.91 17.13a20.22 20.22 0 0 0 20.12.06l167.76-96a19.76 19.76 0 0 0 0-34.6ZM52 203V53l75 75Zm92-58l12.4 12.4l-58 33.2ZM98.41 65.43l58 33.2L144 111ZM178 145l-17-17l17-17l29.72 17Z"/>`),
		g.Group(children),
	)
}