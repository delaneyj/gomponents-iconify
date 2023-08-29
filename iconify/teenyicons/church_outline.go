package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChurchOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="M3 8.5a.5.5 0 0 0-1 0h1Zm10 0a.5.5 0 0 0-1 0h1Zm-7.5 2V10H5v.5h.5Zm4 0h.5V10h-.5v.5ZM0 15h15v-1H0v1Zm7.252-9.934l-7 4l.496.868l7-4l-.496-.868Zm7.496 4l-7-4l-.496.868l7 4l.496-.868ZM7 0v2.5h1V0H7Zm0 2.5v3h1v-3H7ZM5 3h2.5V2H5v1Zm2.5 0H10V2H7.5v1ZM2 8.5v6h1v-6H2Zm10 0v6h1v-6h-1Zm-6 6v-4H5v4h1ZM5.5 11h4v-1h-4v1Zm3.5-.5v4h1v-4H9Z"/>`),
		g.Group(children),
	)
}