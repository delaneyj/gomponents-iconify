package guidance

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FishingAreaOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" d="M21.6 12.5s-1.6-1-1.6-2.25c0-.966.784-1.75 1.75-1.75m-.15 4h.3m-.3 0l-1.227 2m1.377-6c.966 0 1.746.784 1.746 1.75c0 1.25-1.596 2.25-1.596 2.25m-.15-4v-8h-.25C16.373 2 11 7.373 9.5 12.5m12.4 0l1.223 2M0 19.5h7m.5-5.5V6.5H6.328a3 3 0 0 0-2.906 2.255L1.5 16.25v.25h8V18c0 5.5 2.5 5.5 2.5 5.5m11.496-5.25s-2.464-1.688-5.748 0c-3.285 1.688-5.748 0-5.748 0m11.496 2.5s-2.464-1.688-5.748 0c-3.285 1.688-5.748 0-5.748 0M7.35 4.5s-1.6-1-1.6-2.25a1.747 1.747 0 1 1 3.496 0C9.246 3.5 7.65 4.5 7.65 4.5h-.3Z"/>`),
		g.Group(children),
	)
}