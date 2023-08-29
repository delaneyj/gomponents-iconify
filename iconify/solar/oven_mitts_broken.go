package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func OvenMittsBroken(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none"><path fill="currentColor" d="M9.133 17.025a.75.75 0 1 0 1.05-1.07l-1.05 1.07Zm3.364 1.192a.75.75 0 1 0-1.039 1.082l1.04-1.082Zm-2.314-2.262l-5.639-5.536l-1.05 1.07l5.639 5.536l1.05-1.07Zm3.634 3.53l-1.32-1.268l-1.039 1.082l1.32 1.268l1.039-1.082Z"/><path stroke="currentColor" stroke-linecap="round" stroke-width="1.5" d="M4.019 16.537C2.673 15.22 2 14.563 2 13.745c0-.536.29-1.004.87-1.634c.59-.643.886-.964 1.02-1.3c.133-.336.137-.714.144-1.47l.032-3.342C4.032 3.817 5.441 2.027 7.213 2c1.455-.022 2.702 1.152 3.121 2.78m9.744 8.616a6.314 6.314 0 0 0 0-9.072c-2.562-2.505-6.716-2.505-9.278 0l-.466.455m0 0l-.962.941M17.416 16l-.728.711l-3.39 3.315C11.952 21.342 11.278 22 10.443 22c-.837 0-1.51-.658-2.855-1.974L6.517 18.98"/></g>`),
		g.Group(children),
	)
}