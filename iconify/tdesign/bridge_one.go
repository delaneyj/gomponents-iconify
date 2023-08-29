package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BridgeOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M1 1.999h6V5.53c.208.154.524.363.947.575c.878.438 2.226.894 4.053.894s3.175-.456 4.053-.894c.423-.212.739-.42.947-.575V2h6v20h-6v-5H7v5H1V2Zm6 13h4v-6.04c-1.686-.134-3.004-.594-3.947-1.066A8.93 8.93 0 0 1 7 7.867v7.132Zm6-6.04V15h4V7.867a11.78 11.78 0 0 1-.053.026c-.943.472-2.26.932-3.947 1.067ZM3 4v16h2V4H3Zm16 0v16h2V4h-2Z"/>`),
		g.Group(children),
	)
}