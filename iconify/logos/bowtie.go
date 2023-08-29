package logos

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bowtie(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path d="M256 0L128.001 106.493L0 0v249.817l128.001-106.492L256 249.817V0zM150.135 124.908l77.552-64.521v129.042l-77.552-64.521zM28.313 60.387l77.553 64.521l-77.553 64.521V60.387z" fill="#333"/>`),
		g.Group(children),
	)
}