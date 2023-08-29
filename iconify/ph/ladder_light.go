package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LadderLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M213.64 214L162.38 73l9.82-27H184a6 6 0 0 0 0-12H88a6 6 0 0 0 0 12h7.43L34.36 214a6 6 0 0 0 3.64 7.64a6.15 6.15 0 0 0 2 .36a6 6 0 0 0 5.64-3.95L55.84 190h51.23l-8.71 24a6 6 0 0 0 3.59 7.69a6.15 6.15 0 0 0 2 .36a6 6 0 0 0 5.64-3.95l18.98-52.1h54.86l18.93 52.05A6 6 0 0 0 208 222a6.15 6.15 0 0 0 2.05-.36a6 6 0 0 0 3.59-7.64Zm-84.75-84H77.66l13.09-36H142Zm30.54-84l-13.09 36H95.11l13.09-36ZM60.2 178l13.09-36h51.23l-13.09 36Zm72.73-24L156 90.56L179.07 154Z"/>`),
		g.Group(children),
	)
}