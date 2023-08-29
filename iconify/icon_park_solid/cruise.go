package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Cruise(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSCruise0"><g fill="none" stroke="#fff" stroke-width="4"><path stroke-linejoin="round" d="m38 42l3.39-9.325a2 2 0 0 0-.984-2.472l-15.512-7.756a2 2 0 0 0-1.788 0L7.594 30.203a2 2 0 0 0-.985 2.472L10 42"/><path fill="#fff" stroke-linecap="round" stroke-linejoin="round" d="M35 14H13a2 2 0 0 0-2 2v12l12.162-5.613a2 2 0 0 1 1.676 0L37 28V16a2 2 0 0 0-2-2Z"/><path stroke-linecap="round" stroke-linejoin="round" d="M29 14V6a2 2 0 0 0-2-2h-6a2 2 0 0 0-2 2v8"/><path stroke-linecap="round" d="M24 32v8"/><path stroke-linecap="round" stroke-linejoin="round" d="M4 44c4 0 4-2 7-2s3 2 6 2s3.5-2 7-2s4 2 7 2s3-2 6-2s3 2 7 2"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSCruise0)"/>`),
		g.Group(children),
	)
}