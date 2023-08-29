package line_md

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PaintDropHalfFilledTwotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-opacity="0" d="M6 15C6 9 12 4 12 4C12 4 14.9522 6.46019 16.715 10L6.8347 18C6.31173 17.2671 6 16.2894 6 15Z"><animate fill="freeze" attributeName="fill-opacity" begin="0.4s" dur="0.5s" values="0;1"/></path><path fill="currentColor" fill-opacity="0" d="M12 4C12 4 18 9 18 15C18 19 15 20 12 20C9 20 6 19 6 15C6 9 12 4 12 4Z"><animate fill="freeze" attributeName="fill-opacity" begin="0.4s" dur="0.15s" values="0;0.3"/></path><path fill="none" stroke="currentColor" stroke-dasharray="28" stroke-dashoffset="28" stroke-linecap="round" stroke-width="2" d="M12 3C12 3 19 9 19 15C19 17 18 21 12 21M12 3C12 3 5 9 5 15C5 17 6 21 12 21"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.4s" values="28;0"/></path>`),
		g.Group(children),
	)
}