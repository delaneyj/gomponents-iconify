package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Pivo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<rect width="37" height="37" x="5.5" y="5.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="2" ry="2"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m28.353 20.761l-2.671 7.079l-2.672-7.079"/><rect width="5.343" height="7.08" x="30.157" y="20.761" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="2.672" ry="2.672"/><circle cx="20.381" cy="17.488" r=".75" fill="currentColor"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M20.381 20.761v7.079M12.5 25.169a2.672 2.672 0 0 0 2.672 2.671h0a2.672 2.672 0 0 0 2.671-2.671v-1.737a2.672 2.672 0 0 0-2.671-2.671h0a2.672 2.672 0 0 0-2.672 2.671m0-2.671v10.686"/>`),
		g.Group(children),
	)
}