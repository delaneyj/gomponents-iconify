package pajamas

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func IssueClosed(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M9 2.5H3a.5.5 0 0 0-.5.5v10a.5.5 0 0 0 .5.5h6a.5.5 0 0 0 .5-.5V3a.5.5 0 0 0-.5-.5ZM3 1a2 2 0 0 0-2 2v10a2 2 0 0 0 2 2h6a2 2 0 0 0 1.651-.87a.753.753 0 0 0 .098-.145l3.878-7.45a1.75 1.75 0 0 0-.744-2.361l-2.912-1.516A2 2 0 0 0 9 1H3Zm10.297 4.841L11 10.254v-5.89l2.19 1.14a.25.25 0 0 1 .107.337ZM8.28 7.281A.75.75 0 0 0 7.22 6.22L5.25 8.19l-.47-.47a.75.75 0 0 0-1.06 1.06l1 1a.75.75 0 0 0 1.06 0l2.5-2.5Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}