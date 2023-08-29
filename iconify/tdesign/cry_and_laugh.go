package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CryAndLaugh(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 3a9 9 0 1 0 0 18a9 9 0 0 0 0-18ZM1 12C1 5.925 5.925 1 12 1s11 4.925 11 11s-4.925 11-11 11S1 18.075 1 12Zm5-3h4v2H9v2h6v-2h-1V9h4v2h-1v3h-.98l-.04.195C15.582 16.198 14.09 18 12 18s-3.581-1.802-3.98-3.805L7.98 14H7v-3H6V9Zm4.445 6c.422.662 1.02 1 1.555 1c.535 0 1.133-.338 1.555-1h-3.11Z"/>`),
		g.Group(children),
	)
}