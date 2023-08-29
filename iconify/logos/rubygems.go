package logos

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Rubygems(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<g fill="#D34231"><path d="m76.748 97.434l-.163-.163l-36.11 36.11l87.674 87.512l36.11-35.948l51.564-51.563l-36.11-36.11v-.164H76.584l.163.326Z"/><path d="M127.823.976L.135 74.173v146.395l127.688 73.197l127.689-73.197V74.173L127.823.976Zm103.29 205.603l-103.29 59.534l-103.29-59.534V87.837l103.29-59.534l103.29 59.534v118.742Z"/></g>`),
		g.Group(children),
	)
}