package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MeipianPhotoAlbumMaker(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="24" cy="24" r="21.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m14.821 17.68l9.285 17.132L33.39 17.68Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M14.821 17.68c-11.762 8.299-1.559 23.182 9.285 17.132c10.508 5.697 20.723-7.464 9.284-17.133c-.778-12.865-19.541-10.317-18.569 0"/>`),
		g.Group(children),
	)
}