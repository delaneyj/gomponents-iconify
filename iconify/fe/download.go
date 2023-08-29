package fe

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Download(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g id="feDownload0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feDownload1" fill="currentColor"><path id="feDownload2" d="M5 19h14a1 1 0 0 1 0 2H5a1 1 0 0 1 0-2Zm8-5.825l3.243-3.242l1.414 1.414L12 17.004l-5.657-5.657l1.414-1.414L11 13.175V2h2v11.175Z"/></g></g>`),
		g.Group(children),
	)
}