package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ReceiptXLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M216 42H40a14 14 0 0 0-14 14v152a6 6 0 0 0 8.68 5.37L64 198.71l29.32 14.66a6 6 0 0 0 5.36 0L128 198.71l29.32 14.66a6 6 0 0 0 5.36 0L192 198.71l29.32 14.66a6 6 0 0 0 2.68.63a5.93 5.93 0 0 0 3.15-.9A6 6 0 0 0 230 208V56a14 14 0 0 0-14-14Zm2 156.29l-23.32-11.66a6 6 0 0 0-5.36 0L160 201.29l-29.32-14.66a6 6 0 0 0-5.36 0L96 201.29l-29.32-14.66a6 6 0 0 0-5.36 0L38 198.29V56a2 2 0 0 1 2-2h176a2 2 0 0 1 2 2Zm-61.76-98L136.48 120l19.76 19.76a6 6 0 1 1-8.48 8.48L128 128.48l-19.76 19.76a6 6 0 0 1-8.48-8.48L119.52 120l-19.76-19.76a6 6 0 0 1 8.48-8.48L128 111.52l19.76-19.76a6 6 0 0 1 8.48 8.48Z"/>`),
		g.Group(children),
	)
}