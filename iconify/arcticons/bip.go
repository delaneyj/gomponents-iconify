package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bip(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m37.016 41.117l5.189 1.222l-1.093-5.33l.005.006c6.504-8.56 5.688-20.606-1.91-28.21C30.812.402 17.196.397 8.794 8.792C1.195 16.397.38 28.445 6.883 37.005c7.185 9.456 20.676 11.297 30.133 4.112ZM9.462 13.25v10.5"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M14.338 28.75h4.874c3.25 0 4.75-3 4.75-5v-5.5m0 5.5c0 2 1.508 4.987 4.75 5m0 0l4.748-.003a4.916 4.916 0 0 0 4.658-2.901a4.847 4.847 0 0 0-.973-5.371a4.94 4.94 0 0 0-5.385-1.11a4.875 4.875 0 0 0-3.048 4.543V34.75"/><circle cx="23.962" cy="15.317" r=".75" fill="currentColor"/><path fill="none" stroke="currentColor" stroke-linejoin="round" d="M19.212 23.875a4.875 4.875 0 1 1-9.75 0a4.875 4.875 0 0 1 9.75 0Z"/>`),
		g.Group(children),
	)
}