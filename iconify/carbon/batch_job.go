package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BatchJob(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M32 26v-2h-2.101a4.968 4.968 0 0 0-.732-1.753l1.49-1.49l-1.414-1.414l-1.49 1.49A4.964 4.964 0 0 0 26 20.101V18h-2v2.101a4.968 4.968 0 0 0-1.753.732l-1.49-1.49l-1.414 1.414l1.49 1.49A4.964 4.964 0 0 0 20.101 24H18v2h2.101a4.97 4.97 0 0 0 .732 1.753l-1.49 1.49l1.414 1.414l1.49-1.49a4.964 4.964 0 0 0 1.753.732V32h2v-2.101a4.968 4.968 0 0 0 1.753-.732l1.49 1.49l1.414-1.414l-1.49-1.49A4.964 4.964 0 0 0 29.899 26H32zm-7 2c-1.654 0-3-1.346-3-3s1.346-3 3-3s3 1.346 3 3s-1.346 3-3 3zm-5-11h-8a2 2 0 0 1-2-2V4a2 2 0 0 1 2-2h8a2 2 0 0 1 2 2v11a2 2 0 0 1-2 2zm-8-2h8V4h-8v11z"/><path fill="currentColor" d="M17 21H8a2 2 0 0 1-2-2V7h2v12h9v2Z"/><path fill="currentColor" d="M13 25H4c-1.103 0-2-.897-2-2V11h2v12h9v2Z"/>`),
		g.Group(children),
	)
}