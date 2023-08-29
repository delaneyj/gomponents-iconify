package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StarlightLauncher(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M13.252 42.616a21.454 21.454 0 0 1-5.784-4.867m27.28-32.365c10.282 5.936 13.804 19.083 7.868 29.364a21.41 21.41 0 0 1-6.649 7.115c-4.432 3.127-6.567.312-7.458-1.522s-.384-5.743 1.488-7.58c1.12-1.099 1.727-1.311 3.198-3.452a10.618 10.618 0 0 0-.381-11.233"/><path d="M11.564 36.317a4.612 4.612 0 0 0 7.987 4.611l16.885-29.245m-7.987-4.611a4.612 4.612 0 0 1 7.987 4.611m-7.987-4.611L11.564 36.317"/></g><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M34.748 5.384a21.454 21.454 0 0 1 5.784 4.867m-27.28 32.365C2.97 36.68-.553 23.533 5.384 13.252a21.41 21.41 0 0 1 6.649-7.115c4.432-3.127 6.567-.312 7.458 1.522s.384 5.743-1.488 7.58c-1.12 1.098-1.727 1.311-3.198 3.452a10.618 10.618 0 0 0 .381 11.233"/>`),
		g.Group(children),
	)
}