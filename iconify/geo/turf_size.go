package geo

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TurfSize(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 100 100"),
		g.Raw(`<path fill="currentColor" d="M76.647 30.353a2 2 0 0 1-2-2v-3h-3a2 2 0 0 1 0-4h5a2 2 0 0 1 2 2v5a2 2 0 0 1-2 2zm-13.659-5h-8.659a2 2 0 0 1 0-4h8.659a2 2 0 0 1 0 4zm-17.318 0h-8.659a2 2 0 0 1 0-4h8.659a2 2 0 0 1 0 4zm-22.317 5a2 2 0 0 1-2-2v-5a2 2 0 0 1 2-2h5a2 2 0 0 1 0 4h-3v3a2 2 0 0 1-2 2zm0 34.635a2 2 0 0 1-2-2v-8.659a2 2 0 0 1 4 0v8.659a2 2 0 0 1-2 2zm0-17.318a2 2 0 0 1-2-2v-8.659a2 2 0 0 1 4 0v8.659a2 2 0 0 1-2 2zm5 30.977h-5a2 2 0 0 1-2-2v-5a2 2 0 0 1 4 0v3h3a2 2 0 0 1 0 4zm34.635 0h-8.659a2 2 0 0 1 0-4h8.659a2 2 0 0 1 0 4zm-17.318 0h-8.659a2 2 0 0 1 0-4h8.659a2 2 0 0 1 0 4zm30.977 0h-5a2 2 0 0 1 0-4h3v-3a2 2 0 0 1 4 0v5a2 2 0 0 1-2 2zm0-13.659a2 2 0 0 1-2-2v-8.659a2 2 0 0 1 4 0v8.659a2 2 0 0 1-2 2zm0-17.318a2 2 0 0 1-2-2v-8.659a2 2 0 0 1 4 0v8.659a2 2 0 0 1-2 2z"/><path fill="currentColor" d="M90.216 92.216H9.784a2 2 0 0 1-2-2V9.784a2 2 0 0 1 2-2h80.432a2 2 0 0 1 2 2v80.432a2 2 0 0 1-2 2zm-78.432-4h76.432V11.784H11.784v76.432z"/>`),
		g.Group(children),
	)
}