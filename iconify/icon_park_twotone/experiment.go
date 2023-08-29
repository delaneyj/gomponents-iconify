package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Experiment(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTExperiment0"><g fill="none" stroke="#fff" stroke-width="4"><path stroke-linecap="round" d="M12 4h24"/><path stroke-linecap="round" stroke-linejoin="round" d="m10.777 30l7.242-14.961V4h12.01v11.039L37.245 30"/><path fill="#555" stroke-linejoin="round" d="M7.794 43.673a3.273 3.273 0 0 1-1.52-4.372L10.777 30S18 35 24 30c6-5 13.246 0 13.246 0l4.49 9.305A3.273 3.273 0 0 1 38.787 44H9.22c-.494 0-.981-.112-1.426-.327Z"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTExperiment0)"/>`),
		g.Group(children),
	)
}