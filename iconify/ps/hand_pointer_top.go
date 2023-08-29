package ps

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HandPointerTop(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M3 152v124q0 58 37 102t93 51q68 8 118-36t50-111v-43q0-17-12.5-29.5T259 197v85q0 47-35 79t-83 28q-41-7-68.5-39.5T45 276v-81q0-18-12.5-30.5T3 152zm234 85v-64q0-21-21-21t-21 21v64q0 22 21 22t21-22zm-64-21v-64q0-21-21-21t-21 21v64q0 21 21 21t21-21zM109 24q0-21-21-21T67 24v192h42V24z"/>`),
		g.Group(children),
	)
}