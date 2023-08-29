package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Powdertoy(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.109" d="M41.996 42.415H6.004L24.26 26.497Z"/><circle cx="24.259" cy="10.701" r=".75" fill="currentColor"/><circle cx="24.259" cy="15.066" r=".75" fill="currentColor"/><circle cx="24.259" cy="19.431" r=".75" fill="currentColor"/><circle cx="28.922" cy="25.816" r=".75" fill="currentColor"/><circle cx="19.294" cy="26.736" r=".75" fill="currentColor"/><circle cx="24.259" cy="23.796" r=".75" fill="currentColor"/><circle cx="24.259" cy="6.335" r=".75" fill="currentColor"/><circle cx="31.14" cy="28.578" r=".75" fill="currentColor"/><circle cx="16.489" cy="28.578" r=".75" fill="currentColor"/>`),
		g.Group(children),
	)
}