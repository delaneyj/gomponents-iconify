package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func IosLauncher(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 5.5h-33a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2Z"/><circle cx="11.685" cy="18.031" r=".75" fill="currentColor"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M11.685 22.18v8.975m17.108-1.485c.83 1.082 1.873 1.485 3.322 1.485h2.005a3.38 3.38 0 0 0 3.38-3.38v-.014a3.38 3.38 0 0 0-3.38-3.38h-2.212a3.383 3.383 0 0 1-3.383-3.383h0a3.39 3.39 0 0 1 3.39-3.39h1.995c1.45 0 2.49.402 3.322 1.484"/><rect width="8.975" height="13.547" x="15.909" y="17.607" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="4.487" ry="4.487"/>`),
		g.Group(children),
	)
}