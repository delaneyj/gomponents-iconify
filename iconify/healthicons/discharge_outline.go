package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DischargeOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="currentColor"><path fill-rule="evenodd" d="M4.649 20.064a1 1 0 0 0-.585 1.287l2.89 7.709L5.22 36H5a1 1 0 1 0 0 2h38a1 1 0 1 0 0-2h-.22l-1.734-6.94l2.89-7.709a1 1 0 1 0-1.872-.702L40.432 25H30v2h2v2h-1a1 1 0 0 0-.312 1.95L29.246 36H18.754l-1.442-5.05A1 1 0 0 0 17 29h-1v-2h2v-2H7.568l-1.632-4.351a1 1 0 0 0-1.287-.585Zm34.415 8.585A1 1 0 0 0 39 29h-5v-2h5.682l-.618 1.649ZM40.719 36l-1.25-5h-6.715l-1.428 5h9.393ZM8.936 28.649L8.318 27H14v2H9a1 1 0 0 0-.064-.351ZM16.674 36l-1.428-5H8.53l-1.25 5h9.393Z" clip-rule="evenodd"/><path d="M23 13v-3h2v3h3v2h-3v3h-2v-3h-3v-2h3Z"/></g>`),
		g.Group(children),
	)
}