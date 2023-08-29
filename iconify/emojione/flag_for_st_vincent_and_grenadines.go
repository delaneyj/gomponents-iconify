package emojione

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagForStVincentAndGrenadines(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="#ffe62e" d="M32 62c5.9 0 11.4-1.7 16-4.6V6.6C43.4 3.7 37.9 2 32 2S20.6 3.7 16 6.6v50.7c4.6 3 10.1 4.7 16 4.7"/><path fill="#2a5f9e" d="M16 6.6C7.6 11.9 2 21.3 2 32s5.6 20.1 14 25.4V6.6z"/><path fill="#699635" d="M48 57.4c8.4-5.3 14-14.7 14-25.4S56.4 11.9 48 6.6v50.8M24.8 17.9l-6 10.4l6 10.3l6-10.3zm14.4 0l-6 10.4l6 10.3l6-10.3zM26 41.7l6 10.4l6-10.4l-6-10.3z"/>`),
		g.Group(children),
	)
}