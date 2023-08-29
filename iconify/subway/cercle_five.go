package subway

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CercleFive(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M256 0C114.6 0 0 114.6 0 256s114.6 256 256 256c70.7 0 134.7-28.6 181-75c46.3-46.4 75-110.4 75-181C512 114.6 397.4 0 256 0zm0 256c-96.7-96.7-160.9-96-200.8-71.5C84.6 101.9 163.3 42.7 256 42.7c117.8 0 213.3 95.5 213.3 213.3c0 25.1-4.6 49.1-12.5 71.5C416.9 352 352.7 352.7 256 256z"/>`),
		g.Group(children),
	)
}