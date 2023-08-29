package cib

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Toggl(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M16 0C7.161 0 0 7.161 0 16s7.161 16 16 16s16-7.161 16-16S24.839 0 16 0zm-1.12 6.229h2.24v11.145h-2.24zM16 24.208a7.83 7.83 0 0 1-7.817-7.828a7.818 7.818 0 0 1 5.749-7.536v2.281a5.655 5.655 0 0 0-3.593 5.276a5.662 5.662 0 1 0 7.73-5.271V8.838a7.825 7.825 0 0 1 5.755 7.547a7.818 7.818 0 0 1-7.823 7.817z"/>`),
		g.Group(children),
	)
}