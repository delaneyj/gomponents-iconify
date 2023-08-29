package subway

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DownloadFour(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M302.5 93.1h-93.1v186.2H69.8L256 512l186.2-232.7H302.5V93.1zM0 0v46.5h512V0H0z"/>`),
		g.Group(children),
	)
}