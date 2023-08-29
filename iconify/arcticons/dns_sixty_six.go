package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DnsSixtySix(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<rect width="37" height="37" x="5.5" y="5.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="2"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M10.91 22.15V11.5h1.73A5.32 5.32 0 0 1 18 16.83h0a5.32 5.32 0 0 1-5.32 5.32Zm9.56-10.65v10.65m7.06 0V11.5m-7.06 0l7.06 10.65"/><circle cx="14.44" cy="32.97" r="3.53" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M17.65 27.15a3.44 3.44 0 0 0-3-1.3h-.24a3.53 3.53 0 0 0-3.53 3.53V33"/><circle cx="24" cy="32.97" r="3.53" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M27.22 27.15a3.46 3.46 0 0 0-3-1.3H24a3.53 3.53 0 0 0-3.53 3.53V33m9.78-12a3 3 0 0 0 2.61 1.17h1.58a2.66 2.66 0 0 0 2.65-2.66h0a2.66 2.66 0 0 0-2.65-2.66h-1.75A2.67 2.67 0 0 1 30 14.16h0a2.66 2.66 0 0 1 2.65-2.66h1.58a3 3 0 0 1 2.61 1.17"/>`),
		g.Group(children),
	)
}