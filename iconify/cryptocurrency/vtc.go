package cryptocurrency

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Vtc(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M27.25 4.624a13.105 13.105 0 0 0-.312-.29L15.783 19.787l-4.263-5.438H4.582l-1.95 2.22H6.26l8.661 11.194c.29.272.576.408.862.408c.285 0 .557-.136.816-.408L30.294 8.82a15.158 15.158 0 0 0-.295-.574A15.928 15.928 0 0 1 32 16c0 8.837-7.163 16-16 16S0 24.837 0 16S7.163 0 16 0c4.386 0 8.36 1.765 11.25 4.624z"/>`),
		g.Group(children),
	)
}