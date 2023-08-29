package map_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ConvenienceStore(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 50 50"),
		g.Raw(`<path fill="currentColor" d="M41 19h-.272a1.966 1.966 0 0 0-.266-.365L28.918 6.267c.051-.247.082-.503.082-.767C29 3.567 27.576 2 25.818 2h-3.637C20.424 2 19 3.567 19 5.5c0 .264.032.52.082.767L7.538 18.635a2.062 2.062 0 0 0-.266.365H2v7h2v15a5 5 0 0 0 5 5h32a5 5 0 0 0 5-5V26h2v-7h-7zM22.182 9h3.637l.163-.018L35.331 19H12.669l9.35-10.018l.163.018zM13 40c0 1.1-.9 2-2 2s-2-.9-2-2V28c0-1.1.9-2 2-2s2 .9 2 2v12zm9 0c0 1.1-.9 2-2 2s-2-.9-2-2V28c0-1.1.9-2 2-2s2 .9 2 2v12zm10 0c0 1.1-.9 2-2 2s-2-.9-2-2V28c0-1.1.9-2 2-2s2 .9 2 2v12zm9 0c0 1.1-.9 2-2 2s-2-.9-2-2V28c0-1.1.9-2 2-2s2 .9 2 2v12z"/>`),
		g.Group(children),
	)
}