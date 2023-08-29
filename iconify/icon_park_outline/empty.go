package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Empty(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M9.438 42h3.185a2 2 0 0 0 1.868-1.286l2.2-5.753l.025-5.526c0-.116.011-.233.032-.348L20.152 10.3a2 2 0 0 0-1.968-2.357h-7.321a2 2 0 0 0-2 2.056l.523 18.871l-1.77 3.33a2 2 0 0 0 .269 2.266l.561.633l-.963 4.48A2 2 0 0 0 9.438 42Zm7.279-12.74l-7.331-.39M38.578 42h-3.18a2 2 0 0 1-1.878-1.313l-2.093-5.726l-.025-5.526c0-.117-.011-.233-.032-.348L27.965 10.3a2 2 0 0 1 1.968-2.357h7.322a2 2 0 0 1 2 2.055l-.523 18.872l1.771 3.33a2 2 0 0 1-.27 2.266l-.561.633l.87 4.523A2 2 0 0 1 38.578 42ZM31.4 29.26l7.331-.39"/>`),
		g.Group(children),
	)
}