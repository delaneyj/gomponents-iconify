package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Storm(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M817 763L704 480q0-13-9.5-22.5T672 448h-64q-13 0-22.5 9.5T576 480v160H448l-64-160q0-13-9.5-22.5T352 448h-64q-22 0-30 22q-66-80-66-182q0-119 84.5-203.5T480 0q88 0 159.5 48T744 174q-68 22-118 74t-70 121q35-52 91-82.5T768 256q106 0 181 75t75 181q0 93-59 163.5T817 763zM256 640h-96q-13 0-22.5 9.5T128 672v74q-58-28-93-82T0 544q0-57 27-106t73-80q51 103 156 153v129zm64-128l128 320H320v192L192 704h128V512zm320 0l128 320H640v192L512 704h128V512z"/>`),
		g.Group(children),
	)
}