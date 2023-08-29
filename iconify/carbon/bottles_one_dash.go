package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BottlesOneDash(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M26 9.37V3a1 1 0 0 0-1-1h-3v2h2v6.72s3 .507 3 4.28v13h-3v2h4a1 1 0 0 0 1-1V15c0-3.452-1.933-5.024-3-5.63zm-7 0V3a1 1 0 0 0-1-1h-3v2h2v6.72s3 .507 3 4.28v13h-3v2h4a1 1 0 0 0 1-1V15c0-3.452-1.933-5.024-3-5.63zM13 28h-3v2h4a1 1 0 0 0 1-1v-4h-2v3zm-8-3H3v4a1 1 0 0 0 1 1h4v-2H5v-3zm8-7h2v5h-2zM3 18h2v5H3zm9-8.63V3a1 1 0 0 0-1-1H7a1 1 0 0 0-1 1v6.37c-1.067.606-3 2.178-3 5.63v1h2v-1c0-3.772 3-4.28 3-4.28V4h2v6.72s3 .508 3 4.28v1h2v-1c0-3.452-1.933-5.024-3-5.63z"/>`),
		g.Group(children),
	)
}