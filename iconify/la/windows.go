package la

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Windows(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M27 5L5 7.992v16.016L27 27V5zm-2 2.29V15H15V8.65l10-1.36zM13 8.921V15H7V9.738l6-.816zM7 17h6v6.078l-6-.816V17zm8 0h10v7.71l-10-1.36V17z"/>`),
		g.Group(children),
	)
}