package la

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Lightbulb(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M16 4c-4.957 0-9 4.043-9 9c0 1.918.844 3.906 2 5.688c.86 1.324 1.887 2.542 3 3.468V25c0 1.094.906 2 2 2l1 1h2l1-1c1.094 0 2-.906 2-2v-2.844c1.113-.926 2.14-2.144 3-3.468c1.156-1.782 2-3.77 2-5.688c0-4.957-4.043-9-9-9zm0 2c3.879 0 7 3.121 7 7c0 1.309-.645 3.035-1.656 4.594c-.903 1.39-2.09 2.629-3.188 3.406h-4.312c-1.098-.777-2.285-2.016-3.188-3.406C9.645 16.035 9 14.309 9 13c0-3.879 3.121-7 7-7zm-1.75 17h3.5c.078.055.16.094.25.125V25h-4v-1.875a.999.999 0 0 0 .25-.125z"/>`),
		g.Group(children),
	)
}