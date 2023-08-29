package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func IbmDataReplication(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M20 20v2h5.22a11.016 11.016 0 0 1-11.97 4.653l-.499 1.937A13 13 0 0 0 26 24.293V28h2v-8zm5-17a4.005 4.005 0 0 0-4 4a3.954 3.954 0 0 0 .567 2.019L9.019 21.567A3.952 3.952 0 0 0 7 21a4 4 0 1 0 4 4a3.954 3.954 0 0 0-.567-2.019l12.548-12.548A3.952 3.952 0 0 0 25 11a4 4 0 0 0 0-8zM7 27a2 2 0 1 1 2-2a2.002 2.002 0 0 1-2 2zM25 9a2 2 0 1 1 2-2a2.002 2.002 0 0 1-2 2zm-9-6A13.04 13.04 0 0 0 6 7.707V4H4v8h8v-2H6.78a11.016 11.016 0 0 1 11.97-4.653l.499-1.937A13.036 13.036 0 0 0 16 3z"/>`),
		g.Group(children),
	)
}