package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GalnetDictionary(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<rect width="37" height="37" x="5.5" y="5.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="2"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M19.245 20.567v10.117a3.372 3.372 0 0 1-3.373 3.372h0a3.362 3.362 0 0 1-2.384-.988"/><rect width="6.745" height="8.937" x="12.5" y="20.567" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="3.372" transform="rotate(-180 15.872 25.035)"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M22.438 15.944v11.803a1.686 1.686 0 0 0 1.687 1.686h.505m3.148 0l2.215-12.562m2.169 12.562l2.215-12.562m-7.739 8.937h7.925m-6.988-5.311H35.5"/>`),
		g.Group(children),
	)
}