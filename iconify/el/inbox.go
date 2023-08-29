package el

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Inbox(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1200 1200"),
		g.Raw(`<path fill="currentColor" d="M478.125 0v280.005h-142.31L600 676.318l264.185-396.313h-142.31V0h-243.75zM0 621.753V1200h1200V621.753H805.811c0 113.627-92.184 205.811-205.811 205.811S394.189 735.38 394.189 621.753H0z"/>`),
		g.Group(children),
	)
}