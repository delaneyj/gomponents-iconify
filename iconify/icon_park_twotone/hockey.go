package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Hockey(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTHockey0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="2" stroke-width="4"><path fill="#555" d="m36 4l-5.77 27.41c-.2.93-1.01 1.59-1.96 1.59H9c-1.66 0-3 1.34-3 3v1.55c0 3.48 2.95 6.23 6.43 5.98l17.06-1.22a5.996 5.996 0 0 0 5.44-4.75L42 4h-6Z"/><path d="M15 16c4.418 0 8-1.343 8-3s-3.582-3-8-3s-8 1.343-8 3s3.582 3 8 3Z"/><path d="M23 13v6c0 1.66-3.58 3-8 3s-8-1.34-8-3v-6"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTHockey0)"/>`),
		g.Group(children),
	)
}