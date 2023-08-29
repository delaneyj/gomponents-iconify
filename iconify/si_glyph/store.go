package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Store(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M14.021 7.028v5.014H4V7h-.965v8H15V7.028h-.979zM4 0h9.902v.929H4zm10.295 3.001l-.311-.955H4.021l-.336.955l-2.674 3.301l1.334.668h1.634V5.938h1.042V6.97h1.958V5.938h1.042V6.97h1.958V5.938h1.042V6.97h1.958V5.938h1.042V6.97h1.657l1.291-.668l-2.674-3.301z"/>`),
		g.Group(children),
	)
}