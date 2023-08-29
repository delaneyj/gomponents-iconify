package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Lotion(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTLotion0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path fill="#555" d="M17 37h14v7H17z"/><path d="M36 4H12s0 8 1 17s4 16 4 16h14s3-7 4-16s1-17 1-17Z"/><path d="M20.643 21.889c1.431-1.88 2.535-4.479 3.131-5.889c1.044 1.41 3.31 4.948 4.026 6.829c.894 2.35-1.342 5.171-4.026 5.171c-2.684 0-4.92-3.76-3.131-6.111ZM13 10h22"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTLotion0)"/>`),
		g.Group(children),
	)
}