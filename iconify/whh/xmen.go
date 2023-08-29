package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Xmen(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M512 1024q-139 0-257-68.5T68.5 769T0 512t68.5-257T255 68.5T512 0t257 68.5T955.5 255t68.5 257t-68.5 257T769 955.5T512 1024zm0-192q78 0 148-36L512 648L364 796q70 36 148 36zM228 660l148-148l-148-148q-36 70-36 148t36 148zm284-468q-78 0-148 36l148 148l148-148q-70-36-148-36zm284 172L648 512l148 148q36-70 36-148t-36-148z"/>`),
		g.Group(children),
	)
}