package ls

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Adjust(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 717 717"),
		g.Raw(`<path fill="currentColor" d="M358 0c198 0 359 161 359 359S556 717 358 717S0 557 0 359S160 0 358 0zm0 63v592c164 0 296-132 296-296S522 63 358 63z"/>`),
		g.Group(children),
	)
}