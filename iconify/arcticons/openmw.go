package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Openmw(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M11.356 24A20.503 20.503 0 0 1 27.587 3.952a20.5 20.5 0 1 0 0 40.096A20.503 20.503 0 0 1 11.357 24Zm21.74-5.188c2.456-.678 6.415-4.517 6.827-6.494c1.113 3.722 1.549 7.87-3.778 12.351m0 .001c1.815 1.787 7.12 3.296 9.037 2.665c-2.666 2.825-6.042 5.277-12.585 2.904"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M32.597 30.239c-.64 2.466.704 7.814 2.21 9.16c-3.78-.898-7.59-2.595-8.807-9.449m0 0c-2.457.679-6.416 4.518-6.828 6.494c-1.113-3.721-1.549-7.87 3.778-12.351m0 0c-1.815-1.788-7.12-3.297-9.037-2.666c2.666-2.824 6.041-5.276 12.585-2.903m0 0c.641-2.466-.704-7.815-2.21-9.16c3.78.897 7.59 2.594 8.808 9.448"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m33.096 18.812l-3.651 5.281l-2.946-5.569m-.5 11.427l3.446-5.858m3.152 6.146l-3.152-6.146m6.7.576l-6.7-.576m-6.495 0h6.495"/>`),
		g.Group(children),
	)
}