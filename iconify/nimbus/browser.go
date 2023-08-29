package nimbus

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Browser(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<ellipse cx="3.72" cy="4.02" fill="currentColor" rx=".67" ry=".62"/><path fill="currentColor" d="M6.29 4.65A.65.65 0 0 0 7 4a.67.67 0 0 0-1.38 0a.65.65 0 0 0 .67.65z"/><ellipse cx="8.87" cy="4.02" fill="currentColor" rx=".67" ry=".63"/><path fill="currentColor" d="M14.25 1.5H1.75A1.25 1.25 0 0 0 .5 2.75v10.5a1.25 1.25 0 0 0 1.25 1.25h12.5a1.25 1.25 0 0 0 1.25-1.25V2.75a1.25 1.25 0 0 0-1.25-1.25zM1.75 2.75h12.5v2.5H1.75v-2.5zm0 10.5V6.5h12.5v6.75z"/>`),
		g.Group(children),
	)
}