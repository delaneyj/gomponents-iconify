package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ScreenRecorder(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<rect width="28.3" height="25.12" x="4.5" y="11.44" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="2.49"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m32.8 27.71l8.32 7.37A1.43 1.43 0 0 0 43.5 34V14a1.43 1.43 0 0 0-2.38-1.07l-8.32 7.36"/>`),
		g.Group(children),
	)
}