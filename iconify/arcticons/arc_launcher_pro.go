package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArcLauncherPro(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24 5.747L5 42.253l13.004-6.765L24 24.505l5.975 11.027L43 42.253L24 5.747Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M17.452 42.218v-4.854h1.577c.91 0 1.639.728 1.639 1.638s-.728 1.638-1.639 1.638h-1.577m4.94 1.578v-4.854h1.578c.91 0 1.638.728 1.638 1.638s-.728 1.638-1.638 1.638h-1.578m1.648.01l1.508 1.507m3.362.061c-.91 0-1.578-.728-1.578-1.578v-1.638c0-.91.728-1.638 1.578-1.638c.91 0 1.638.728 1.638 1.638v1.578c0 .91-.728 1.638-1.638 1.638Z"/>`),
		g.Group(children),
	)
}