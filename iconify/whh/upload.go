package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Upload(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M512 1024q-139 0-257-68.5T68.5 769T0 512t68.5-257T255 68.5T512 0t257 68.5T955.5 255t68.5 257t-68.5 257T769 955.5T512 1024zm33-882q-13-14-32.5-14T480 142L138 487q-12 14-8.5 19.5T148 512h236v320q0 27 19 45.5t45 18.5h129q26 0 45-18.5t19-45.5V512h235q15 0 19-5.5t-8-19.5z"/>`),
		g.Group(children),
	)
}