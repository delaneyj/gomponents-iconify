package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PencilRulerTwoFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m5.636 12.707l1.829 1.829l1.414-1.415l-1.829-1.828l1.415-1.414l1.828 1.828l1.414-1.414L9.88 8.465l1.414-1.415l1.828 1.829l1.415-1.414l-1.829-1.829l2.829-2.828a1 1 0 0 1 1.414 0l4.242 4.242a1 1 0 0 1 0 1.415L8.465 21.192a1 1 0 0 1-1.415 0L2.808 16.95a1 1 0 0 1 0-1.414l2.828-2.829Zm8.485 5.656l4.243-4.242L21 16.757V21h-4.242l-2.637-2.637ZM5.636 9.878L2.808 7.05a1 1 0 0 1 0-1.415l2.828-2.828a1 1 0 0 1 1.414 0l2.83 2.828l-4.244 4.243Z"/>`),
		g.Group(children),
	)
}