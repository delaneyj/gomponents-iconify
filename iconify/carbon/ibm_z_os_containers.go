package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func IbmZOsContainers(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M29 12h-9V3h9v9Zm-7-2h5V5h-5v5Z"/><path fill="currentColor" d="M20 15v2h7v10H17V3H5c-1.103 0-2 .897-2 2v22c0 1.103.897 2 2 2h22c1.103 0 2-.897 2-2V15h-9ZM6.414 17H15v8.586L6.414 17ZM15 15H6.414L15 6.414V15ZM13.586 5L5 13.586V5h8.586ZM5 18.414L13.586 27H5v-8.586Z"/>`),
		g.Group(children),
	)
}