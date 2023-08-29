package uim

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ObjectUngroup(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 16a2 2 0 1 1 0-4a2 2 0 0 1 0 4zm0-2.002zM4 6a2 2 0 1 1 0-4a2 2 0 0 1 0 4zm0-2.002zM14 6a2 2 0 1 1 0-4a2 2 0 0 1 0 4zm0-2.002zM14 16a2 2 0 1 1 0-4a2 2 0 0 1 0 4zm0-2.002z" opacity=".5"/><path fill="currentColor" d="M10 22a2 2 0 1 1 0-4a2 2 0 0 1 0 4zm0-2.002zM10 12a2 2 0 1 1 0-4a2 2 0 0 1 0 4zm0-2.002zM20 12a2 2 0 1 1 0-4a2 2 0 0 1 0 4zm0-2.002zM20 22a2 2 0 1 1 0-4a2 2 0 0 1 0 4zm0-2.002z"/><path fill="currentColor" d="M12.278 5a1.936 1.936 0 0 1 0-2H5.722a1.936 1.936 0 0 1 0 2h6.556zM4 12a1.98 1.98 0 0 1 1 .278V5.722a1.936 1.936 0 0 1-2 0v6.556A1.98 1.98 0 0 1 4 12z" opacity=".25"/><path fill="currentColor" d="M20 18a1.98 1.98 0 0 1 1 .278v-6.556a1.936 1.936 0 0 1-2 0v6.556A1.98 1.98 0 0 1 20 18zm-10 0a1.98 1.98 0 0 1 1 .278v-6.556a1.936 1.936 0 0 1-2 0v6.556A1.98 1.98 0 0 1 10 18z" opacity=".5"/><path fill="currentColor" d="M12.278 13H11v2h1.278a1.936 1.936 0 0 1 0-2zM9 15v-2H5.722a1.936 1.936 0 0 1 0 2H9z" opacity=".25"/><path fill="currentColor" d="M18.278 19h-6.556a1.936 1.936 0 0 1 0 2h6.556a1.936 1.936 0 0 1 0-2zm0-8a1.936 1.936 0 0 1 0-2h-6.556a1.936 1.936 0 0 1 0 2h6.556z" opacity=".5"/><path fill="currentColor" d="M15 9V5.722a1.936 1.936 0 0 1-2 0V9h2zm-2 2v1.278a1.936 1.936 0 0 1 2 0V11h-2z" opacity=".25"/>`),
		g.Group(children),
	)
}