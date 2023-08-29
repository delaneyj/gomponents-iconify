package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Samsungpaymini(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<rect width="39" height="31" x="4.5" y="8.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="2"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M36.348 20.4H43.5v7.2h0h-7.152a2.848 2.848 0 0 1-2.848-2.848v-1.504a2.848 2.848 0 0 1 2.848-2.848Z"/><circle cx="36.472" cy="24" r=".795" fill="currentColor"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M15.416 28.473c.758.636 1.576.928 3.414.928h.931a2.747 2.747 0 0 0 2.744-2.75h0a2.747 2.747 0 0 0-2.744-2.75H17.9a2.747 2.747 0 0 1-2.744-2.75h0a2.747 2.747 0 0 1 2.744-2.75h.931c1.838 0 2.656.292 3.414.928"/>`),
		g.Group(children),
	)
}