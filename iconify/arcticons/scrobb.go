package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Scrobb(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M17.027 20.747c.822-2.074 1.654-4.148 3.719-5.582c2.065-1.425 5.372-2.193 7.903-.457a9.993 9.993 0 0 1 3.718 8.826c-.557 2.85-3.417 4.294-5.116 6.039c-1.7 1.754-2.248 3.81-3.015 5.08a3.425 3.425 0 0 1-2.832 1.827a4.494 4.494 0 0 1-2.979-.859a5.048 5.048 0 0 1-1.398-2.33"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24.93 18.673a2.535 2.535 0 1 1-2.54 2.54a2.538 2.538 0 0 1 2.54-2.54Zm9.292-5.829q6.971 7.903 0 17.195"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M6.538 21.898a13.441 13.441 0 0 1 2.55-4.403c1.388-1.535 3.6-3.19 5.792-5.555c2.193-2.376 4.376-5.455 7.19-6.725c2.805-1.27 6.25-.73 9.036.694a17.753 17.753 0 0 1 6.487 6.03a21.627 21.627 0 0 1 3.243 7.648a28.427 28.427 0 0 1 .704 9.968a13.741 13.741 0 0 1-4.404 8.341a22.853 22.853 0 0 1-9.51 5.327a8.678 8.678 0 0 1-7.648-1.617c-2.12-1.6-3.938-4.057-6.021-6.03s-4.431-3.463-5.793-5.327a10.79 10.79 0 0 1-1.854-5.564a7.278 7.278 0 0 1 .229-2.787Z"/>`),
		g.Group(children),
	)
}