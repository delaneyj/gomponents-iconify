package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SplitBranch(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTSplitBranch0"><g fill="none"><path fill="#555" d="M44 44V4H24v13l9 9v18h11ZM4 4v40h21V30l-9-9V4H4Z"/><path stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M44 44V4H24v13l9 9v18h11ZM4 4v40h21V30l-9-9V4H4Z"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTSplitBranch0)"/>`),
		g.Group(children),
	)
}