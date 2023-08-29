package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Microsoftlens(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M4.5 16v16a2 2 0 0 0 2 2h16a2 2 0 0 0 2-2V16a2 2 0 0 0-2-2h-16a2 2 0 0 0-2 2Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M9.43 34a18.5 18.5 0 1 0 0-20m32.46 2.44L24.5 16.4m-2.46-2.47l4.81-8.34m-2.35 26l4.87.01M23.1 42.4l10.67-18.38m6.22 10.83L29.39 16.4"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M11.75 18.01v11h5.5"/>`),
		g.Group(children),
	)
}