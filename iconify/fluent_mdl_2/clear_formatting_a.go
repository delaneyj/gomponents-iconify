package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ClearFormattingA(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="m640 0l329 988l-105 105l-65-197H353l-85 256H128L512 0h128zM396 768h360L576 228L396 768z"/>`),
		g.Group(children),
	)
}