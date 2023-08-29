package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ToiletPaperOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4.27 4.28C3.502 5.55 3 7.639 3 10c0 3.866 1.343 7 3 7s3-3.134 3-7c0-.34-.01-.672-.03-1M21 10c0-3.866-1.343-7-3-7M7 3h11m3 7v7m-1.513 2.496L18 19l-3 2l-3-3l-3 2V10m-3 0h.01M3 3l18 18"/>`),
		g.Group(children),
	)
}