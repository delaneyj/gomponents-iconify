package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Markup(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 3a9 9 0 0 0-5.61 16.038L8.228 12H10V8.586l2-2l2 2V12h1.773l1.837 7.038A9 9 0 0 0 12 3Zm3.832 17.146L14.227 14H9.772l-1.604 6.146A8.963 8.963 0 0 0 12 21c1.372 0 2.67-.306 3.832-.854ZM1 12C1 5.925 5.925 1 12 1s11 4.925 11 11a10.991 10.991 0 0 1-5.5 9.528A10.954 10.954 0 0 1 12 23a10.954 10.954 0 0 1-6.013-1.787A10.991 10.991 0 0 1 1 12Z"/>`),
		g.Group(children),
	)
}