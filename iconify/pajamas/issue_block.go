package pajamas

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func IssueBlock(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M3 1a2 2 0 0 0-2 2v3.25a.75.75 0 0 0 1.5 0V3a.5.5 0 0 1 .5-.5h6a.5.5 0 0 1 .5.5v10.234a.753.753 0 0 0 .57.744a.752.752 0 0 0 .851-.393l3.7-7.047a1.75 1.75 0 0 0-.74-2.366l-2.91-1.514A2 2 0 0 0 9 1H3Zm8 3.364v5.844l2.293-4.367a.25.25 0 0 0-.106-.338L11 4.363Zm-4.783 8.792a2.5 2.5 0 0 0-3.373-3.373l3.373 3.373Zm-1.06 1.061a2.5 2.5 0 0 1-3.373-3.373l3.372 3.373ZM4 16a4 4 0 1 0 0-8a4 4 0 0 0 0 8Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}