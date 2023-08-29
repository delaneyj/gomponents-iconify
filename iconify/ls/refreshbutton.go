package ls

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Refreshbutton(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 717 717"),
		g.Raw(`<path fill="currentColor" d="M358 0c198 0 359 161 359 359S556 717 358 717S0 557 0 359S160 0 358 0zM198 368h-69c0 60 23 120 67 165c90 89 235 89 325 0c89-90 89-236 0-325c-42-41-94-63-148-66V65L198 179l175 113v-82c37 3 71 18 99 46c63 62 63 166 0 228c-62 63-165 63-227 0c-33-31-48-74-47-116z"/>`),
		g.Group(children),
	)
}