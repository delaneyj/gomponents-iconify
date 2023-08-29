package prime

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SortNumericAltDown(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8.38 19.75a.75.75 0 0 1-.54-.22L5.34 17a.75.75 0 0 1 0-1.06a.77.77 0 0 1 1.07 0l2 2l2-2a.77.77 0 0 1 1.07 0a.75.75 0 0 1 0 1.06l-2.5 2.5a.74.74 0 0 1-.6.25Z"/><path fill="currentColor" d="M8.38 19.75a.76.76 0 0 1-.76-.75V5a.76.76 0 0 1 .76-.75a.75.75 0 0 1 .74.75v14a.75.75 0 0 1-.74.75Zm8.74-.5a.75.75 0 0 1-.74-.75v-4.06l-.39.22a.75.75 0 0 1-.73-1.32l.66-.36a1.19 1.19 0 0 1 1.22-.11a1.29 1.29 0 0 1 .74 1.18v4.45a.76.76 0 0 1-.76.75Zm-.5-10A2.25 2.25 0 1 1 18.88 7a2.24 2.24 0 0 1-2.26 2.25Zm0-3a.75.75 0 0 0 0 1.5a.75.75 0 1 0 0-1.5Z"/><path fill="currentColor" d="M16.11 11.25h-.49a.75.75 0 0 1 0-1.5h.49a1.28 1.28 0 0 0 1.25-1.19V7a.75.75 0 0 1 .74-.75a.76.76 0 0 1 .76.75v1.64a2.78 2.78 0 0 1-2.75 2.61Z"/>`),
		g.Group(children),
	)
}