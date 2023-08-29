package typcn

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DeviceLaptop(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M21.989 16.049c.009-.315.011-.657.011-1.049V6c0-1.654-1.346-3-3-3H5C3.346 3 2 4.346 2 6v9c0 .385.002.73.012 1.049A2.504 2.504 0 0 0 0 18.5C0 19.878 1.122 21 2.5 21h19c1.378 0 2.5-1.122 2.5-2.5a2.504 2.504 0 0 0-2.011-2.451zM4 6c0-.551.449-1 1-1h14c.551 0 1 .449 1 1v9c0 .388-.005.726-.014 1H19V7c0-.55-.45-1-1-1H6c-.55 0-1 .45-1 1v9h-.98c-.012-.264-.02-.599-.02-1V6zm14 10H6V7h12v9zm3.5 3h-19c-.271 0-.5-.229-.5-.5s.229-.5.5-.5h19c.271 0 .5.229.5.5s-.229.5-.5.5z"/>`),
		g.Group(children),
	)
}