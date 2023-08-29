package ls

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Compass(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 717 717"),
		g.Raw(`<path fill="currentColor" d="M358 0c198 0 359 161 359 359S556 717 358 717S0 557 0 359S160 0 358 0zm0 634c152 0 276-123 276-275S510 83 358 83S83 207 83 359s123 275 275 275zm59-217l-214 97l97-214l214-97zm-175 59l155-78l-78-78z"/>`),
		g.Group(children),
	)
}