package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CloudFogBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M120 200H72a12 12 0 0 1 0-24h48a12 12 0 0 1 0 24Zm64-24h-24a12 12 0 0 0 0 24h24a12 12 0 0 0 0-24Zm-24 36h-56a12 12 0 0 0 0 24h56a12 12 0 0 0 0-24Zm72-124a76.08 76.08 0 0 1-76 76H76a52 52 0 1 1 9-103.22A76 76 0 0 1 232 88Zm-24 0a52 52 0 0 0-104 0a12 12 0 0 1-24 0c0-1.24 0-2.48.09-3.71A29.28 29.28 0 0 0 76 84a28 28 0 0 0 0 56h80a52.06 52.06 0 0 0 52-52Z"/>`),
		g.Group(children),
	)
}