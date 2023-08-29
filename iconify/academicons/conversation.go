package academicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Conversation(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 448 512"),
		g.Raw(`<path fill="currentColor" d="M324.199 96H215.277c-84.567 0-154.25 69.683-154.25 154.25v6.088c0 30.444 9.472 58.182 25.709 83.214L5.552 416l118.393-37.21c25.708 19.62 58.182 30.445 92.685 30.445h107.569c85.92 0 154.25-69.683 154.25-152.897v-6.089C478.448 165.683 408.764 96 324.198 96ZM406 276c0 46.68-35.234 75.67-81.238 75.67H213.134C166.454 351.67 132 322.68 132 276v-40c0-46.68 34.453-81.201 81.134-81.201h111.628c46.68 0 81.238 34.52 81.238 81.201Z" class="st1"/>`),
		g.Group(children),
	)
}