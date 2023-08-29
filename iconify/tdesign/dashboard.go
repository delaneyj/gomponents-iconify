package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Dashboard(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 3a9 9 0 0 0-4.5 16.796l.865.5l-1.001 1.732l-.866-.5A10.996 10.996 0 0 1 1 12C1 5.925 5.925 1 12 1c1.203 0 2.362.193 3.448.552l.95.313l-.627 1.9l-.95-.314A8.991 8.991 0 0 0 12 3Zm8.981 1.414l-5.542 5.542a4 4 0 1 1-1.42-1.41L19.567 3l1.414 1.414Zm1.154 3.188l.313.95C22.807 9.638 23 10.797 23 12c0 4.071-2.212 7.625-5.496 9.526l-.865.5l-1.002-1.73l.865-.501A8.996 8.996 0 0 0 21 12a8.99 8.99 0 0 0-.45-2.822l-.314-.95l1.9-.626ZM12 10a2 2 0 1 0 0 4a2 2 0 0 0 0-4Z"/>`),
		g.Group(children),
	)
}