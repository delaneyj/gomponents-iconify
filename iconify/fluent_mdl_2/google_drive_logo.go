package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GoogleDriveLogo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M1995 1261h-646L699 128h646l650 1133zm-1186 93h1239l-323 566H487l322-566zM619 270l323 566l-619 1084L0 1354L619 270z"/>`),
		g.Group(children),
	)
}