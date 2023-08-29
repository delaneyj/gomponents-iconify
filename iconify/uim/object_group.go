package uim

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ObjectGroup(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11 10h5a1 1 0 0 1 1 1v5a1 1 0 0 1-1 1h-5a1 1 0 0 1-1-1v-5a1 1 0 0 1 1-1z" opacity=".5"/><path fill="currentColor" d="M10 11a1 1 0 0 1 1-1h3V8a1 1 0 0 0-1-1H8a1 1 0 0 0-1 1v5a1 1 0 0 0 1 1h2v-3zM4 22a2 2 0 1 1 0-4a2 2 0 0 1 0 4zM4 6a2 2 0 1 1 0-4a2 2 0 0 1 0 4zm16 0a2 2 0 1 1 0-4a2 2 0 0 1 0 4zm0 16a2 2 0 1 1 0-4a2 2 0 0 1 0 4z"/><path fill="currentColor" d="M18.278 5a1.936 1.936 0 0 1 0-2H5.722a1.936 1.936 0 0 1 0 2h12.556zM20 18a1.98 1.98 0 0 1 1 .278V5.722a1.936 1.936 0 0 1-2 0v12.556A1.98 1.98 0 0 1 20 18zM4 18a1.98 1.98 0 0 1 1 .278V5.722a1.936 1.936 0 0 1-2 0v12.556A1.98 1.98 0 0 1 4 18zm14.278 1H5.722a1.936 1.936 0 0 1 0 2h12.556a1.936 1.936 0 0 1 0-2z" opacity=".25"/>`),
		g.Group(children),
	)
}