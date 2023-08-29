package file_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Buck(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="m389.104 259.712l-40.21-40.21l46.19-45.983V68.77h-37.736v89.079l-35.054 35.054l-34.642-34.642V68.768h-37.735v105.164l85.78 85.78h-53.406L144.96 122.175V33.302h-37.736v104.544l46.19 46.396h-53.82l-61.86-62.067V33.302H0v104.544l84.13 84.13h107.226l37.735 37.736H121.866l88.666 88.667L80.42 478.699h53.406l130.32-130.32l-51.138-51.139h261.258v18.765L358.38 385.289l-93.615 93.41h52.994l61.86-62.067L512 337.656v-77.944H389.104zm-100.215 62.066l26.6 26.6l26.6-26.6h-53.2z"/>`),
		g.Group(children),
	)
}