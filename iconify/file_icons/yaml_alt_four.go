package file_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func YamlAltFour(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="m235.793 20.939l-91.815 137.674v87.275H87.702v-87.275L0 20.939h63.25l55.768 88.646l56.225-88.646h60.55zm94.501 174.925H228.433l-20.717 50.024H162.61l95.38-224.949h46.137l91.51 224.949h-48.196l-17.148-50.024zm-16.92-44.734l-31.226-82.55l-34.837 82.55h66.062zM87.701 270.59v220.47h47.303V338.982l49.505 102.22h37.234l51.196-105.813v155.626h45.379V270.59h-61.96l-54.977 99.706l-52.36-99.706h-61.32zM512 443.2H395.638V270.59h-48.196v219.522H512v-46.91z"/>`),
		g.Group(children),
	)
}