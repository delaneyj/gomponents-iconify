package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MotionMode(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5 4.85v14.3q-.425-.425-.8-.9t-.7-.975V6.725q.325-.5.7-.975t.8-.9Zm4-2.4v19.1q-.525-.175-1.025-.388T7 20.676V3.325q.475-.275.975-.488T9 2.45Zm7 18.725V2.825Q18.65 4 20.325 6.45T22 12q0 3.1-1.675 5.55T16 21.175ZM12 22q-.25 0-.5-.013t-.5-.037V2.05q.25-.025.5-.037T12 2q.5 0 1 .05t1 .15v19.6q-.5.1-1 .15T12 22Z"/>`),
		g.Group(children),
	)
}