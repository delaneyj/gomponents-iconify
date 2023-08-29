package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CommentOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTCommentOne0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path fill="#555" d="M44 6H4v30h9v5l10-5h21V6Z"/><path d="M14 21h20"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTCommentOne0)"/>`),
		g.Group(children),
	)
}