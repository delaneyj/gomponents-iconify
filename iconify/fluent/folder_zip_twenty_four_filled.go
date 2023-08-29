package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FolderZipTwentyFourFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M2 6.25V8h6.129a.75.75 0 0 0 .53-.22l2.591-2.59l-1.53-1.531A2.25 2.25 0 0 0 8.129 3H5.25A3.25 3.25 0 0 0 2 6.25Zm0 11.5V9.5h6.129a2.25 2.25 0 0 0 1.59-.659L13.062 5.5h.439v3.75c0 .414.336.75.75.75H15v3h-.25a.75.75 0 0 0 0 1.5H15V16h-.25a.75.75 0 0 0 0 1.5H15V21H5.25A3.25 3.25 0 0 1 2 17.75ZM16.5 21h2.25A3.25 3.25 0 0 0 22 17.75v-9a3.25 3.25 0 0 0-3.25-3.25H18v3.75a.75.75 0 0 1-.75.75h-.75v4.5h.25a.75.75 0 0 1 0 1.5h-.25v5Zm0-15.5H15v3h1.5v-3Z"/>`),
		g.Group(children),
	)
}