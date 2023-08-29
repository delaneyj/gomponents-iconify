package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HuaweiHidisk(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M14.68 36.934a89.15 89.15 0 0 1 18.645.375A9.232 9.232 0 0 0 43.49 28.23a9.924 9.924 0 0 0-10.054-9.904a8.596 8.596 0 0 0-8.178 6.228c-1.527 6.326-6.141 9.256-10.58 12.38Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M33.57 18.326c-.975-4.062-4.726-7.69-11.424-7.69c-4.26 0-7.715 2.27-9.379 8.928c-.642 2.574-.179 5.02 1.726 8.854c1.62 3.26 1.546 6.118.188 8.515"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M12.477 22.489C9.174 21.678 4.55 24.72 4.5 29.77c-.056 5.698 5.647 7.53 10.18 7.164"/>`),
		g.Group(children),
	)
}