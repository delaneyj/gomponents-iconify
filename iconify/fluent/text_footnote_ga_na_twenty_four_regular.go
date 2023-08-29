package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TextFootnoteGaNaTwentyFourRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<g fill="none"><path d="M20.5 7.263a5.27 5.27 0 0 1-.607.475a.75.75 0 0 1-.832-1.248c.764-.51 1.33-1.321 1.456-1.9A.75.75 0 0 1 22 4.75v6.5a.75.75 0 0 1-1.5 0V7.263zM2 9.25a.75.75 0 0 1 .75-.75h3.5a.75.75 0 0 1 .75.75c0 .968-.166 2.241-.76 3.358c-.61 1.149-1.673 2.117-3.374 2.383a.75.75 0 1 1-.232-1.482c1.175-.184 1.862-.816 2.282-1.606c.312-.585.476-1.262.544-1.903H2.75A.75.75 0 0 1 2 9.25zm6-3v12a.75.75 0 0 0 1.5 0v-12a.75.75 0 0 0-1.5 0zm8 12v-12a.75.75 0 0 1 1.5 0v12a.75.75 0 0 1-1.5 0zM12.742 7.106a.75.75 0 0 0-1.485-.212l-1 7a.75.75 0 0 0 .85.848l3.5-.5a.75.75 0 0 0-.213-1.485l-2.51.36l.858-6.01z" fill="currentColor"/></g>`),
		g.Group(children),
	)
}