package subway

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DownloadThree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M442.2 186.2H302.5V0h-93.1v186.2H69.8L256 418.9l186.2-232.7zM0 465.5V512h512v-46.5H0z"/>`),
		g.Group(children),
	)
}