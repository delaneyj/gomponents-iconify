package pepicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HeartPrint(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<g fill="currentColor" fill-rule="evenodd" clip-rule="evenodd"><path d="M11 6.722c1.69-3.023 7.5-1.968 7.5 2.4c0 2.918-2.5 5.582-7.5 7.993c-5-2.41-7.5-5.075-7.5-7.993c0-4.368 5.81-5.423 7.5-2.4Z" opacity=".8"/><path d="M5.742 4.634C4.256 5.02 3 6.207 3 8.122c0 2.56 2.168 5.065 7 7.437c4.832-2.372 7-4.877 7-7.437c0-1.915-1.256-3.102-2.742-3.488c-1.517-.394-3.12.078-3.821 1.332a.5.5 0 0 1-.873 0C8.863 4.712 7.259 4.24 5.742 4.634ZM10 4.856c-1.1-1.263-2.932-1.6-4.51-1.19C3.65 4.145 2 5.67 2 8.122c0 3.237 2.767 6.025 7.783 8.444a.5.5 0 0 0 .434 0C15.233 14.146 18 11.359 18 8.122c0-2.453-1.65-3.977-3.49-4.456c-1.578-.41-3.41-.073-4.51 1.19Z"/></g>`),
		g.Group(children),
	)
}