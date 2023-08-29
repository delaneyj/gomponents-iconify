package codicon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FileCode(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="m10.57 1.14l3.28 3.3l.15.36v9.7l-.5.5h-11l-.5-.5v-13l.5-.5h7.72l.35.14zM10 5h3l-3-3v3zM3 2v12h10V6H9.5L9 5.5V2H3zm2.062 7.533l1.817-1.828L6.17 7L4 9.179v.707l2.171 2.174l.707-.707l-1.816-1.82zM8.8 7.714l.7-.709l2.189 2.175v.709L9.5 12.062l-.705-.709l1.831-1.82L8.8 7.714z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}