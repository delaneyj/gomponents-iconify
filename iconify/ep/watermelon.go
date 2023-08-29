package ep

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Watermelon(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="m683.072 600.32l-43.648 162.816l-61.824-16.512l53.248-198.528L576 493.248l-158.4 158.4l-45.248-45.248l158.4-158.4l-55.616-55.616l-198.528 53.248l-16.512-61.824l162.816-43.648L282.752 200A384 384 0 0 0 824 741.248L683.072 600.32zm231.552 141.056a448 448 0 1 1-632-632l632 632z"/>`),
		g.Group(children),
	)
}