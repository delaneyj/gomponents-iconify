package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Backup(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M12.073 35.927A16.867 16.867 0 1 0 7.133 24"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M7.133 30.127L1.867 24h10.532l-5.266 6.127z"/><circle cx="24" cy="24" r="4" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/>`),
		g.Group(children),
	)
}