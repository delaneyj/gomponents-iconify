package emojione

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DisappointedFace(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<circle cx="32" cy="32" r="30" fill="#ffdd67"/><path fill="#664e27" d="M25.5 28.4c1.4 2.9-.4 6.6-3.9 8.3c-3.5 1.6-7.5.6-8.9-2.3c-.8-1.9 12-7.9 12.8-6m13 0c-1.4 2.9.4 6.6 3.9 8.3c3.5 1.6 7.5.6 8.9-2.3c.8-1.9-12-7.9-12.8-6"/><path fill="#917524" d="M22.7 19.8c-2.7 3.3-9.2 6.3-13.5 6.3c-.6 0-.7 2.2 0 2.2c4.9 0 12-3.3 15.2-7.1c.5-.5-1.3-1.8-1.7-1.4m18.6 0c2.7 3.3 9.2 6.3 13.5 6.3c.6 0 .7 2.2 0 2.2c-4.9 0-12-3.3-15.2-7.1c-.5-.5 1.3-1.8 1.7-1.4"/><path fill="#664e27" d="M40.6 46.4c-5.4-2.5-11.8-2.5-17.2 0c-1.3.6.3 4.2 1.7 3.5c3.6-1.7 8.9-2.3 13.9 0c1.3.6 3-2.8 1.6-3.5"/>`),
		g.Group(children),
	)
}