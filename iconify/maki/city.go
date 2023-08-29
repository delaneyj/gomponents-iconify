package maki

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func City(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="M12.637 4h-1.639V2.36a.36.36 0 0 0-.36-.36h-.279a.36.36 0 0 0-.36.36V4H8.36a.36.36 0 0 0-.36.36v5.638H4.361c-.2 0-.363.163-.363.363v2.275c0 .2.163.362.363.362h8.275c.2 0 .36-.162.36-.361V4.36a.36.36 0 0 0-.36-.36Zm-6.638 7.998h-1v-1h1v1Zm2 0h-1v-1h1v1Zm2 0h-1v-1h1v1Zm0-2h-1v-1h1v1Zm0-2h-1v-1h1v1Zm0-2h-1v-1h1v1Zm2 6h-1v-1h1v1Zm0-2h-1v-1h1v1Zm0-2h-1v-1h1v1Zm0-2h-1v-1h1v1Zm-5-3.637A.36.36 0 0 0 6.638 2H4.36a.36.36 0 0 0-.36.36V4H2.36a.36.36 0 0 0-.36.36v8.287c0 .194.157.35.35.35H3V9h3.999V2.36ZM4 7.998H3v-1h1v1Zm0-2H3v-1h1v1Zm2 2H5v-1h1v1Zm0-2H5v-1h1v1Zm0-2H5V3h1v1v-.002Z"/>`),
		g.Group(children),
	)
}