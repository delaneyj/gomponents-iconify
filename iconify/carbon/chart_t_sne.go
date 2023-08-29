package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChartTSne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<circle cx="10" cy="20" r="2" fill="currentColor"/><circle cx="10" cy="28" r="2" fill="currentColor"/><circle cx="10" cy="14" r="2" fill="currentColor"/><circle cx="28" cy="4" r="2" fill="currentColor"/><circle cx="22" cy="6" r="2" fill="currentColor"/><circle cx="28" cy="10" r="2" fill="currentColor"/><circle cx="20" cy="12" r="2" fill="currentColor"/><circle cx="28" cy="22" r="2" fill="currentColor"/><circle cx="26" cy="28" r="2" fill="currentColor"/><circle cx="20" cy="26" r="2" fill="currentColor"/><circle cx="22" cy="20" r="2" fill="currentColor"/><circle cx="16" cy="4" r="2" fill="currentColor"/><circle cx="4" cy="24" r="2" fill="currentColor"/><circle cx="4" cy="16" r="2" fill="currentColor"/>`),
		g.Group(children),
	)
}