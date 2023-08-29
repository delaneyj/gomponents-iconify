package typcn

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowForwardOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 19.999a.999.999 0 0 1-1-1v-1a9.98 9.98 0 0 1 8-9.796V6.499c0-.534.208-1.036.585-1.414c.756-.757 2.075-.756 2.829-.001l6.288 6.203a.996.996 0 0 1 0 1.424l-6.293 6.207c-.746.746-2.067.751-2.823-.005A1.986 1.986 0 0 1 11 17.499v-1.437c-2.495.201-4.523.985-6.164 3.484a1 1 0 0 1-.836.453zm8-5.989l1-.01v3.499l5.576-5.5L13 6.503V10s-.384-.004-.891.052a7.982 7.982 0 0 0-6.892 6.08C7.338 14.404 9.768 14.066 12 14.01z"/>`),
		g.Group(children),
	)
}