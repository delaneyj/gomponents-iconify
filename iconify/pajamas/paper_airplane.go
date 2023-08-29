package pajamas

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PaperAirplane(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M14.968 1.966a.75.75 0 0 0-.934-.934l-12.5 3.75a.75.75 0 0 0-.18 1.355L5.952 8.99l-1.731 1.73a.75.75 0 1 0 1.06 1.061l1.731-1.73l2.852 4.595a.75.75 0 0 0 1.355-.18l3.75-12.5ZM8.101 8.96l2.159 3.48l2.417-8.056L8.1 8.96Zm3.515-5.637L3.56 5.74L7.04 7.9l4.576-4.577Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}