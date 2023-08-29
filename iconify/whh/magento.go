package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Magento(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M768 841V320L448 160L128 320v521L0 768V256L448 0l449 256v512zm-384 18l64 37l64-37V325l129 75v512l-193 112l-192-112V400l128-75v534z"/>`),
		g.Group(children),
	)
}