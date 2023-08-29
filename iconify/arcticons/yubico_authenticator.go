package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func YubicoAuthenticator(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="24" cy="24" r="21.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="36.728" cy="11.272" r=".75" fill="currentColor"/><circle cx="36.728" cy="11.272" r=".75" fill="currentColor"/><circle cx="42" cy="24" r=".75" fill="currentColor"/><circle cx="42" cy="24" r=".75" fill="currentColor"/><circle cx="36.728" cy="36.728" r=".75" fill="currentColor"/><circle cx="36.728" cy="36.728" r=".75" fill="currentColor"/><circle cx="24" cy="42" r=".75" fill="currentColor"/><circle cx="24" cy="42" r=".75" fill="currentColor"/><circle cx="11.272" cy="36.728" r=".75" fill="currentColor"/><circle cx="11.272" cy="36.728" r=".75" fill="currentColor"/><circle cx="6" cy="24" r=".75" fill="currentColor"/><circle cx="6" cy="24" r=".75" fill="currentColor"/><circle cx="11.272" cy="11.272" r=".75" fill="currentColor"/><circle cx="11.272" cy="11.272" r=".75" fill="currentColor"/><circle cx="24" cy="6" r=".75" fill="currentColor"/><circle cx="24" cy="6" r=".75" fill="currentColor"/><circle cx="24" cy="24" r="15.07" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m21.29 32.5l8.943-17M24 27.349L17.767 15.5"/>`),
		g.Group(children),
	)
}