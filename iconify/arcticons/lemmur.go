package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Lemmur(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M19.091 38.204s-1.545-3.466 2.436-5.646m0 0s1.106-4.421 2.473-4.436M9.293 29.797l-3.22 1.602m2.394-4.006l-3.671.766m-.188-3.345l3.72.002m.042-1.006S2.008 19.438 5.6 8.185"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M15.42 12.421a24.01 24.01 0 0 0-9.821-4.235s5.14 2.758 5.37 7.932m4.745 7.051s-3.739.506-3.406 5.457s5.168 2.864 5.168 2.864s3.387-1.648 2.836-4.337a4.707 4.707 0 0 0-4.597-3.984Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M17.578 28.637a1.106 1.106 0 1 1-1.106-1.105a1.106 1.106 0 0 1 1.106 1.105Zm8.895 3.921S25.367 28.137 24 28.122m-2.473 4.436s2.295 6.246 4.946 0m.361 6.629c-.1.014-2.853-.864-2.853-3.446"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M21.128 39.187c.1.014 2.853-.864 2.853-3.446"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24 10.34c-.414-.01-15.4.097-15.4 14.307a8.16 8.16 0 0 0 1.32 5.56c1.822 2.55 3.077 4.843 9.608 8.3A8.54 8.54 0 0 0 24 39.815m4.909-1.611s1.545-3.466-2.436-5.646m0 0S25.367 28.137 24 28.122m14.708 1.675l3.22 1.602m-2.395-4.006l3.671.766m.188-3.345l-3.72.002m-.042-1.006s6.363-4.372 2.77-15.625"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M32.58 12.421a24.01 24.01 0 0 1 9.821-4.235s-5.14 2.758-5.37 7.932m-4.746 7.051s3.74.506 3.407 5.457s-5.168 2.864-5.168 2.864s-3.387-1.648-2.836-4.337a4.707 4.707 0 0 1 4.597-3.984Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M30.422 28.637a1.106 1.106 0 1 0 1.106-1.105a1.106 1.106 0 0 0-1.106 1.105Zm-8.895 3.921s1.106-4.421 2.473-4.436m2.473 4.436s-2.295 6.246-4.946 0m-.361 6.629c.1.014 2.853-.864 2.853-3.446"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M26.872 39.187c-.1.014-2.853-.864-2.853-3.446"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24 10.34c.414-.01 15.4.097 15.4 14.307a8.16 8.16 0 0 1-1.32 5.56c-1.822 2.55-3.077 4.843-9.608 8.3A8.54 8.54 0 0 1 24 39.815"/>`),
		g.Group(children),
	)
}