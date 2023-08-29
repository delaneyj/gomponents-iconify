package map_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MovieRental(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 50 50"),
		g.Raw(`<path fill="currentColor" d="M25.507 22.572a2.743 2.743 0 0 0-2.741 2.74a2.743 2.743 0 0 0 2.741 2.74a2.744 2.744 0 0 0 2.741-2.74c0-1.511-1.23-2.74-2.741-2.74zM41.824 4.073A2.073 2.073 0 0 0 39.752 2H10.073A2.073 2.073 0 0 0 8 4.073v41.856C8 47.073 8.928 48 10.073 48h29.679a2.07 2.07 0 0 0 2.072-2.071V4.073zM25.507 38.415c-7.225 0-13.103-5.877-13.103-13.103c0-7.224 5.878-13.102 13.103-13.102c7.224 0 13.101 5.878 13.101 13.102c-.001 7.226-5.878 13.103-13.101 13.103z"/>`),
		g.Group(children),
	)
}