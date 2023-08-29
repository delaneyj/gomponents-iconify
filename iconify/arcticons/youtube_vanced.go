package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func YoutubeVanced(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M43.112 14.394a5.006 5.006 0 0 0-3.533-3.533c-2.314-.894-24.732-1.332-31.236.025A5.006 5.006 0 0 0 4.81 14.42c-1.045 4.583-1.124 14.491.026 19.177a5.006 5.006 0 0 0 3.533 3.533c4.583 1.055 26.371 1.203 31.236 0a5.006 5.006 0 0 0 3.533-3.533c1.114-4.993 1.193-14.287-.026-19.203Z"/><g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="m19.471 31.495l10.052-5.765c1.34-.768 1.34-2.701 0-3.47l-10.052-5.765"/><path d="M25.102 23.475L20.37 20.76a.6.6 0 0 0-.899.521v5.428a.6.6 0 0 0 .899.52l4.732-2.713a.6.6 0 0 0 0-1.041Z"/></g>`),
		g.Group(children),
	)
}