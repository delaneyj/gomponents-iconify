package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LineageGallery(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<rect width="39" height="26" x="4.5" y="11" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="2"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m4.5 28.501l11.8-11.8l20.241 20.241M43.5 16.701l-13.6 13.6M35.529 11l-5.657 5.657l-5.601-5.601"/>`),
		g.Group(children),
	)
}