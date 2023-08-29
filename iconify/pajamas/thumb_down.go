package pajamas

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ThumbDown(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="m10.72 13.402l-.27-.902L10 11h3.989a2 2 0 0 0 1.932-2.517l-1.19-4.448A4 4 0 0 0 9.852 1.2L3 3H0v8h3l4.39 4.39c1.47 1.47 3.928.002 3.33-1.988ZM3 9.5H1.5v-5H3v5Zm7.232-6.85L4.5 4.157v6.222l3.951 3.951c.12.119.22.149.296.155a.533.533 0 0 0 .314-.08a.533.533 0 0 0 .22-.238a.466.466 0 0 0 .003-.334l-.72-2.402l-.58-1.931h6.005a.5.5 0 0 0 .483-.63l-1.19-4.447a2.5 2.5 0 0 0-3.05-1.772Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}