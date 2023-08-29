package ls

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Sitemap(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 717 717"),
		g.Raw(`<path fill="currentColor" d="M604 338v142h76v181H499V480h75V368H355v112h76v181H249V480h76V368H106v112h75v181H0V480h76V338h249V228h-98V47h226v181h-98v110h249z"/>`),
		g.Group(children),
	)
}